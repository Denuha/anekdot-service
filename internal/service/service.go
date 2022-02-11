package service

import (
	"context"

	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
)

type Services struct {
	Anekdot Anekdot
}

type Anekdot interface {
	ParseAnekdots(ctx context.Context, source string) (int, error)
	GetRandomAnekdot(ctx context.Context) (*models.Anekdot, error)
}

func NewServices(cfg *config.Config, repos *repository.Repositories) *Services {
	return &Services{
		Anekdot: NewAnekdotService(repos),
	}
}
