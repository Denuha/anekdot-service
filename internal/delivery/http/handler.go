package http

import (
	"net/http"

	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Services
	log      *logrus.Logger
	cfg      *config.Config
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		h.initRoutesAnekdot(api)
		h.initRoutesUser(api)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}

func (h *Handler) Response(c *gin.Context,
	statusCode int,
	err error,
	message interface{}) {
	resp := models.Response{
		HasError: false,
		Resp:     message,
	}
	if err != nil {
		resp.HasError = true
		resp.ErrorText = err.Error()
		h.log.Errorf("%s", err.Error())
	}

	c.AbortWithStatusJSON(statusCode, resp)
}

func NewHandlers(services *service.Services, log *logrus.Logger, cfg *config.Config) *Handler {
	return &Handler{
		services: services,
		log:      log,
		cfg:      cfg,
	}
}
