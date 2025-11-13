package service

import (
	"context"

	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
)

type metrics struct {
	metricsDB repository.Metrics
}

func (m *metrics) GetMetrics(ctx context.Context) (*models.Metrics, error) {
	return m.metricsDB.GetMetrics(ctx)
}

func NewMetricsService(repos *repository.Repositories) Metrics {
	return &metrics{
		metricsDB: repos.Metrics,
	}
}
