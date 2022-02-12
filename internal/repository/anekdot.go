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
func (a *anekdot) ChangeRating(ctx context.Context, anekdotID int, method MethodRaitng) error {
	queryUpdate := `
	UPDATE anekdot.anekdot 
	SET rating = rating + $1
	WHERE id = $2;`

	cl, err := a.client.GetClient()
	if err != nil {
		return err
	}

	deltaRating := 0

	switch method {
	case MethodIncr:
		deltaRating = 1
	case MethodDecr:
		deltaRating = -1
	default:
		return errors.New("wrong method")
	}

	res, err := cl.ExecContext(ctx, queryUpdate, deltaRating, anekdotID)
	if err != nil {
		return err
	}

	if count, _ := res.RowsAffected(); count == 0 {
		return errors.New("no change")
	}

	return nil
}

func (a *anekdot) GetAnekdotByID(ctx context.Context, anekdotID int) (*models.Anekdot, error) {
	querySelect := `
	SELECT a.id, a."text", a.rating, a.external_id, a.create_time, a.status, a.sender_id, s."name" 
	FROM anekdot.anekdot a
	LEFT JOIN anekdot.sender s on a.sender_id=s.id 
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
		&anekdot.Rating,
		&anekdot.ExternalID,
		&anekdot.CreateTime,
		&anekdot.Status,
		&anekdot.Sender.ID,
		&anekdot.Sender.Name,
	)

	if err != nil {
		return nil, err
	}

	return &anekdot, nil

	return nil, nil
}

func (a *anekdot) GetRandomAnekdot(ctx context.Context) (*models.Anekdot, error) {
	querySelect := `
		SELECT a.id, a."text", a.rating, a.external_id, a.create_time, a.status, a.sender_id, s."name" 
		FROM anekdot.anekdot a
		LEFT JOIN anekdot.sender s on a.sender_id=s.id 
		ORDER BY random() limit 1;`

	cl, err := a.client.GetClient()
	if err != nil {
		return nil, err
	}

	row := cl.QueryRowContext(ctx, querySelect)

	var anekdot models.Anekdot
	err = row.Scan(
		&anekdot.ID,
		&anekdot.Text,
		&anekdot.Rating,
		&anekdot.ExternalID,
		&anekdot.CreateTime,
		&anekdot.Status,
		&anekdot.Sender.ID,
		&anekdot.Sender.Name,
	)

	if err != nil {
		return nil, err
	}

	return &anekdot, nil
}

func NewAnekdotRepo(client clientRepo.PostgresClient) AnekdotDB {
	return &anekdot{
		client: client,
	}
}
