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

// @Summary Get user list
// @Description
// @Security ApiKeyAuth
// @Tags Users
// @ID getUserList
// @Accept json
// @Produce json
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 200 {object} models.Response{resp=[]models.User} "OK"
// @Router /user [get]
func (h *Handler) getUserList(ctx *gin.Context) {
	users, err := h.services.User.GetUserList(ctx)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	h.Response(ctx, http.StatusOK, nil, users)
}
