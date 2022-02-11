package service

import (
	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/repository"
	p "github.com/Denuha/anekdot-service/internal/service/parser"
)

type Services struct {
	ParserService p.Parser
}

func NewServices(cfg *config.Config, repos *repository.Repositories) *Services {
	return &Services{}
}
