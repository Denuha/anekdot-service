package service

import (
	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/repository"
)

type Services struct {
}

func NewServices(cfg *config.Config, repos *repository.Repositories) *Services {
	return &Services{}
}
