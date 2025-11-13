package repository

import (
	"context"

	"github.com/Denuha/anekdot-service/internal/models"
	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
)

type metrics struct {
	client clientRepo.PostgresClient
}

func (m *metrics) GetMetrics(ctx context.Context) (*models.Metrics, error) {
	const querySelect = `
	SELECT (SELECT count(*) FROM anekdot."user" u) AS number_users,
(SELECT count(*) FROM anekdot.anekdot a) AS number_anekdots,
(SELECT count(*) FROM anekdot.user_votes a) AS number_user_votes`

	cl, err := m.client.GetClient()
	if err != nil {
		return nil, err
	}

	row := cl.QueryRowContext(ctx, querySelect)

	var metrics models.Metrics
	err = row.Scan(
		&metrics.NumberUsers,
		&metrics.NumberAnekdots,
		&metrics.NumberUserVotes,
	)

	if err != nil {
		return nil, err
	}

	return &metrics, nil
}

func NewMetricsRepo(client clientRepo.PostgresClient) Metrics {
	return &metrics{
		client: client,
	}
}
