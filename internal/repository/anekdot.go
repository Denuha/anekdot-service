package repository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"

	"github.com/Denuha/anekdot-service/internal/models"
	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
)

type anekdot struct {
	client clientRepo.PostgresClient
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
	querySelect := `
	SELECT a.id, a."text", a.external_id, a.create_time, a.status, a.sender_id, s."name",
		(SELECT count(u.value) 
				FROM user_votes u
				WHERE u.anekdot_id= a.id AND u.value > 0) AS likes ,
		(SELECT count(u.value) 
				FROM user_votes u
				WHERE u.anekdot_id= a.id AND u.value < 0) AS dislikes 
	FROM anekdot.anekdot a
	LEFT JOIN anekdot.sender s ON a.sender_id=s.id 
	WHERE a.id = $1;`

	cl, err := a.client.GetClient()
	if err != nil {
		return nil, err
	}

	row := cl.QueryRowContext(ctx, querySelect, anekdotID)

	var anekdot models.Anekdot
	err = row.Scan(
		&anekdot.ID,
		&anekdot.Text,
		&anekdot.ExternalID,
		&anekdot.CreateTime,
		&anekdot.Status,
		&anekdot.Sender.ID,
		&anekdot.Sender.Name,
		&anekdot.Likes,
		&anekdot.Dislikes,
	)

	if err != nil {
		return nil, err
	}

	return &anekdot, nil
}

func (a *anekdot) GetRandomAnekdot(ctx context.Context) (*models.Anekdot, error) {
	querySelect := `
	SELECT a.id, a."text", a.external_id, a.create_time, a.status, a.sender_id, s."name",  
	(SELECT count(u.value) 
			FROM anekdot.user_votes u
			WHERE u.anekdot_id= a.id AND u.value > 0) AS likes ,
	(SELECT count(u.value) 
			FROM anekdot.user_votes u
			WHERE u.anekdot_id= a.id AND u.value < 0) AS dislikes 
	FROM anekdot.anekdot a
	LEFT JOIN anekdot.sender s ON a.sender_id=s.id
	LEFT JOIN anekdot.user_votes u ON a.id=u.anekdot_id
	ORDER BY random() limit 1`

	cl, err := a.client.GetClient()
	if err != nil {
		return nil, err
	}

	row := cl.QueryRowContext(ctx, querySelect)

	var anekdot models.Anekdot
	err = row.Scan(
		&anekdot.ID,
		&anekdot.Text,
		&anekdot.ExternalID,
		&anekdot.CreateTime,
		&anekdot.Status,
		&anekdot.Sender.ID,
		&anekdot.Sender.Name,
		&anekdot.Likes,
		&anekdot.Dislikes,
	)

	if err != nil {
		return nil, err
	}

	return &anekdot, nil
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
	const queyUpdate = `UPDATE anekdot.user_votes
	SET value=$1
	WHERE user_id=$2 AND anekdot_id=$3;`

	cl, err := a.client.GetClient()
	if err != nil {
		return err
	}

	res, err := cl.ExecContext(ctx, queyUpdate, value, userID, anekdotID)
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

func NewAnekdotRepo(client clientRepo.PostgresClient) AnekdotDB {
	return &anekdot{
		client: client,
	}
}
