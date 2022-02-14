package repository

import (
	"context"
	"database/sql"

	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
)

type commonDBRepo struct {
	clientRepo.PostgresClient
}

func (c *commonDBRepo) BeginTransaction(ctx context.Context) (*sql.Tx, error) {
	client, err := c.GetClient()
	if err != nil {
		return nil, err
	}

	tx, err := client.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *commonDBRepo) CommitTransaction(_ context.Context, tx *sql.Tx) error {
	return tx.Commit()
}

func NewCommonRepo(postgresClient clientRepo.PostgresClient) CommonDB {
	return &commonDBRepo{
		PostgresClient: postgresClient,
	}
}
