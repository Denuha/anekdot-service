package http

import (
	"net/http"
	"time"

	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initRoutesAuth(rg *gin.RouterGroup) {
	authRoutes := rg.Group("")

	authRoutes.POST("/login", h.login)
	authRoutes.POST("/registration", h.registration)
}

func (h *Handler) registration(ctx *gin.Context) {
	var user *models.UserRegistation

	err := ctx.BindJSON(&user)
	if err != nil {
		h.Response(ctx, http.StatusBadRequest, err, nil)
		return
	}

	id, err := h.services.User.Registration(ctx, user)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusCreated, nil, map[string]int{"user_id": id})
}

func (h *Handler) login(ctx *gin.Context) {
	var user *models.UserLogin

	type login struct {
		AccessToken string        `json:"access_token"`
		Expires     time.Duration `json:"expires_seconds"`
	}

	err := ctx.BindJSON(&user)
	if err != nil {
		h.Response(ctx, http.StatusBadRequest, err, nil)
		return
	}

	token, err := h.services.User.Login(ctx, user)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, login{
		AccessToken: token,
		Expires:     time.Duration(h.cfg.TokenExpires.Seconds()),
	})
}
