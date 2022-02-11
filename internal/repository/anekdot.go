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
		queryInsert = queryInsert.Values(anekdot.Text, anekdot.Status, anekdot.ExternalID, anekdot.SenderID)
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
	return nil
}
func (a *anekdot) GetRandomAnekdot(ctx context.Context) (*models.Anekdot, error) {
	return nil, nil
}

func NewAnekdotRepo(client clientRepo.PostgresClient) AnekdotDB {
	return &anekdot{
		client: client,
	}
}
