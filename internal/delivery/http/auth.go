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

// @Summary Registration of new user
// @Description
// @Security ApiKeyAuth
// @Tags Users
// @ID registration
// @Accept json
// @Produce json
// @Param data body models.UserRegistation true "Body request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 201 {object} models.Response "OK"
// @Router /registration [post]
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

// @Summary Login
// @Description
// @Security ApiKeyAuth
// @Tags Users
// @ID login
// @Accept json
// @Produce json
// @Param data body models.UserLogin true "Body request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 201 {object} models.Response{resp=models.Login} "OK"
// @Router /login [post]
func (h *Handler) login(ctx *gin.Context) {
	var user *models.UserLogin

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

	h.Response(ctx, http.StatusOK, nil, models.Login{
		AccessToken: token,
		Expires:     time.Duration(h.cfg.TokenExpires.Seconds()),
	})
}
