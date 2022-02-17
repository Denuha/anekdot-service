package http

import (
	"net/http"

	swag "github.com/Denuha/anekdot-service/docs" //
	"github.com/Denuha/anekdot-service/internal/auth"
	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Services
	log      *logrus.Logger
	cfg      *config.Config
	auth     *auth.Auth
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		h.initRoutesAnekdot(api)
		h.initRoutesUser(api)
		h.initRoutesAuth(api)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	swag.SwaggerInfo_swagger.BasePath = h.cfg.SwaggerBasePath
	swag.SwaggerInfo_swagger.Host = h.cfg.SwaggerHost
	swag.SwaggerInfo_swagger.Version = h.cfg.SwaggerVersion

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

func NewHandlers(services *service.Services, log *logrus.Logger, cfg *config.Config, auth *auth.Auth) *Handler {
	return &Handler{
		services: services,
		log:      log,
		cfg:      cfg,
		auth:     auth,
	}
}
