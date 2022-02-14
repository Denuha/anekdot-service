package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initRoutesUser(rg *gin.RouterGroup) {
	userGroup := rg.Group("/user")
	userGroup.Use(h.userVerify())

	userGroup.GET("", h.getUserList)
}

func (h *Handler) getUserList(ctx *gin.Context) {
	users, err := h.services.User.GetUserList(ctx)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	h.Response(ctx, http.StatusOK, nil, users)
}
