package http

import (
	"net/http"

	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Services
}

func (h *Handler) Init(_ *config.Config, log *logrus.Logger) *gin.Engine {
	router := gin.Default()

	// Init router
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}

func NewHandlers(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}
