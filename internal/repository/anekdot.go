package repository

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"

	"github.com/Denuha/anekdot-service/internal/models"
	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
	"github.com/Denuha/anekdot-service/internal/utils"
)

type anekdot struct {
	client clientRepo.PostgresClient
}

func NewAnekdotRepo(client clientRepo.PostgresClient) AnekdotDB {
	return &anekdot{
		client: client,
	}
}

func (a *anekdot) InsertAnekdotList(ctx context.Context, anekdotList []models.Anekdot) error {
	queryInsert := sq.Insert("anekdot.anekdot").Columns("text", "status", "external_id", "sender_id")

	for _, anekdot := range anekdotList {
		queryInsert = queryInsert.Values(anekdot.Text, anekdot.Status, anekdot.ExternalID, anekdot.Sender.ID)
	}

	query, args, err := queryInsert.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	cl, err := a.client.GetClient()
	if err != nil {
		return err
	}

	result, err := cl.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	if aff, _ := result.RowsAffected(); aff < 1 {
		return errors.New("affected 0")
	}

	return nil
}

func (a *anekdot) GetAnekdotByID(ctx context.Context, anekdotID int) (*models.Anekdot, error) {
	cl, err := a.client.GetClient()
	if err != nil {
		return nil, err
	}

	anekdot, err := a.getAnekdot(ctx, cl, anekdotID)
	if err != nil {
		return nil, err
	}

	return anekdot, nil
}

func (a *anekdot) getAnekdot(ctx context.Context, tx *sql.DB, anekdotID int) (*models.Anekdot, error) {
	querySelect := sq.Select(`a.id`, `a."text"`, `a.external_id`, `a.create_time`, `a.status`, `a.sender_id`, `s."name"`).
		From(`anekdot.anekdot a`).
		Join(`anekdot.sender s ON a.sender_id=s.id`)

	if anekdotID == 0 {
		// random anekdot
		user, err := utils.GetUserFromContext(ctx)
		if err != nil {
			return nil, err
		}

		querySelect = querySelect.LeftJoin(`anekdot.user_votes uv on uv.anekdot_id = a.id`).
			Where(sq.Or{
				sq.NotEq{`uv.user_id`: user.ID},
				sq.Eq{`uv.user_id`: nil},
			},
			)

		querySelect = querySelect.OrderBy(`random()`).Limit(1)
	} else {
		querySelect = querySelect.Where(sq.Eq{"a.id": anekdotID})
	}

	query, args, err := querySelect.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var anekdot models.Anekdot
	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&anekdot.ID,
		&anekdot.Text,
		&anekdot.ExternalID,
		&anekdot.CreateTime,
		&anekdot.Status,
		&anekdot.Sender.ID,
		&anekdot.Sender.Name,
	)
	if err != nil {
		return nil, err
	}

	likes, err := a.getCountLikes(ctx, tx, anekdotID)
	if err != nil {
		return nil, err
	}

	dislikes, err := a.getCountDisLikes(ctx, tx, anekdotID)
	if err != nil {
		return nil, err
	}

	skips, err := a.getCountSkips(ctx, tx, anekdotID)
	if err != nil {
		return nil, err
	}

	anekdot.Likes = likes
	anekdot.Dislikes = dislikes
	anekdot.Skips = skips
	return &anekdot, nil
}

func (a *anekdot) GetRandomAnekdot(ctx context.Context) (*models.Anekdot, error) {
	cl, err := a.client.GetClient()
	if err != nil {
		return nil, err
	}
	anekdot, err := a.getAnekdot(ctx, cl, 0)
	if err != nil {
		return nil, err
	}

	return anekdot, nil
}

func (a *anekdot) GetUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64) (*models.AnekdotVote, error) {
	const querySelect = `SELECT user_id, anekdot_id, value 
	FROM anekdot.user_votes
	WHERE user_id=$1 AND anekdot_id=$2;`

	cl, err := a.client.GetClient()
	if err != nil {
		return nil, err
	}

	var vote models.AnekdotVote
	row := cl.QueryRowContext(ctx, querySelect, userID, anekdotID)
	err = row.Scan(
		&vote.UserID,
		&vote.AnekdotID,
		&vote.Value,
	)

	if err != nil {
		return nil, err
	}
	return &vote, nil
}

func (a *anekdot) UpdateUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64, value int) error {
	const queryUpdate = `UPDATE anekdot.user_votes
	SET value=$1
	WHERE user_id=$2 AND anekdot_id=$3;`

	cl, err := a.client.GetClient()
	if err != nil {
		return err
	}

	res, err := cl.ExecContext(ctx, queryUpdate, value, userID, anekdotID)
	if err != nil {
		return err
	}
	if count, _ := res.RowsAffected(); count < 1 {
		return errors.New("no change")
	}

	return nil
}

func (a *anekdot) PostUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64, value int) error {
	const queyUpdate = `INSERT INTO anekdot.user_votes (user_id, anekdot_id, value)
	VALUES ($1,$2,$3);`

	cl, err := a.client.GetClient()
	if err != nil {
		return err
	}

	res, err := cl.ExecContext(ctx, queyUpdate, userID, anekdotID, value)
	if err != nil {
		return err
	}
	if count, _ := res.RowsAffected(); count < 1 {
		return errors.New("no change")
	}

	return nil
}

// getCountLikes get count rating > 0
func (a *anekdot) getCountLikes(ctx context.Context, tx *sql.DB, anekdotID int) (int, error) {
	querySelect := sq.Select(`count(u.value)`).
		From(`anekdot.user_votes u`).
		Where(sq.And{
			sq.Eq{"u.anekdot_id": anekdotID},
			sq.Gt{"u.value": 0},
		})

	query, args, err := querySelect.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	countLikes := 0
	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&countLikes,
	)
	if err != nil {
		return 0, nil
	}

	return countLikes, nil
}

// getCountLikes get count rating < 0
func (a *anekdot) getCountDisLikes(ctx context.Context, tx *sql.DB, anekdotID int) (int, error) {
	querySelect := sq.Select(`count(u.value)`).
		From(`anekdot.user_votes u`).
		Where(sq.And{
			sq.Eq{"u.anekdot_id": anekdotID},
			sq.Lt{"u.value": 0},
		})

	query, args, err := querySelect.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	countDisLikes := 0
	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&countDisLikes,
	)
	if err != nil {
		return 0, nil
	}

	return countDisLikes, nil
}

// getCountLikes get count rating = 0
func (a *anekdot) getCountSkips(ctx context.Context, tx *sql.DB, anekdotID int) (int, error) {
	querySelect := sq.Select(`count(u.value)`).
		From(`anekdot.user_votes u`).
		Where(sq.And{
			sq.Eq{"u.anekdot_id": anekdotID},
			sq.Eq{"u.value": 0},
		})

	query, args, err := querySelect.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return 0, err
	}

	countSkips := 0
	row := tx.QueryRowContext(ctx, query, args...)
	err = row.Scan(
		&countSkips,
	)
	if err != nil {
		return 0, nil
	}

	return countSkips, nil
}
