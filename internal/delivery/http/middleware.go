package http

import (
	"errors"
	"time"

	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/gin-gonic/gin"
)

var debugUser = models.User{
	ID:         2,
	UserName:   "quest",
	ExternalID: "",
	Realm:      "anekdot",
	CreateTime: time.Now(),
}

func (h *Handler) userVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *models.User

		if h.cfg.Debug {
			user = &debugUser
		}

		if user == nil {
			h.Response(ctx, 403, errors.New("user is nil"), "user is nil")
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
