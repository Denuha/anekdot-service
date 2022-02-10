package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Denuha/anekdot-service/internal/config"
)

type Server interface {
	Run() error
	Stop(ctx context.Context) error
}

type server struct {
	httpServer *http.Server
}

func (s *server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func NewServer(cfg config.Config, handler http.Handler) Server {
	return &server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: handler,
		},
	}
}
