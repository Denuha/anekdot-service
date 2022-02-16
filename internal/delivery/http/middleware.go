package http

import (
	"errors"
	"net/http"

	"github.com/Denuha/anekdot-service/internal/auth"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) userVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *models.User
		var err error

		if h.cfg.Debug {
			user = &auth.DebugUser
		} else {
			user, err = h.auth.GetUserFromRequest(ctx)
			if err != nil {
				h.Response(ctx, http.StatusUnauthorized, err, "")
				return
			}
		}

		if user == nil {
			h.Response(ctx, http.StatusUnauthorized, errors.New("user is nil"), "user is nil")
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}

func (h *Handler) adminVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *models.User
		var err error

		if h.cfg.Debug {
			user = &auth.DebugUser
		} else {
			user, err = h.auth.GetUserFromRequest(ctx)
			if err != nil {
				h.Response(ctx, http.StatusUnauthorized, err, "")
				return
			}
		}

		if user == nil {
			h.Response(ctx, http.StatusUnauthorized, errors.New("user is nil"), "user is nil")
			return
		}

		if !user.IsAdmin {
			h.Response(ctx, http.StatusForbidden, errors.New("no access"), "no access")
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
