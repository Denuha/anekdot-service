package http

import (
	"net/http"

	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initRoutesAuth(rg *gin.RouterGroup) {
	authRoutes := rg.Group("/auth")

	authRoutes.POST("/login", h.login)
	authRoutes.POST("logout", h.logout)
	authRoutes.POST("/registration", h.registration)
	authRoutes.POST("/refresh", h.refreshToken)
}

// @Summary Registration of new user
// @Description
// @Tags Auth
// @ID registration
// @Accept json
// @Produce json
// @Param data body models.UserRegistation true "Body request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 201 {object} models.Response "OK"
// @Router /auth/registration [post]
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
// @Tags Auth
// @ID login
// @Accept json
// @Produce json
// @Param data body models.UserLogin true "Body request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 200 {object} models.Response{resp=models.Login} "OK"
// @Router /auth/login [post]
func (h *Handler) login(ctx *gin.Context) {
	var user *models.UserLogin

	err := ctx.BindJSON(&user)
	if err != nil {
		h.Response(ctx, http.StatusBadRequest, err, nil)
		return
	}

	token, err := h.services.Auth.Login(ctx, user)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, token)
}

// @Summary Logout
// @Description
// @Tags Auth
// @ID logout
// @Accept json
// @Produce json
// @Param data body models.AccessToken true "Body request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 200 {object} models.Response "OK"
// @Router /auth/logout [post]
func (h *Handler) logout(ctx *gin.Context) {
	var access *models.AccessToken

	err := ctx.BindJSON(&access)
	if err != nil {
		h.Response(ctx, http.StatusBadRequest, err, nil)
		return
	}

	err = h.services.Auth.Logout(ctx, access.AccessToken)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, nil)
}

// @Summary Refresh token
// @Description
// @Tags Auth
// @ID refreshToken
// @Accept json
// @Produce json
// @Param data body models.RefreshToken true "Body request"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 200 {object} models.Response{resp=models.Login} "OK"
// @Router /auth/refresh [post]
func (h *Handler) refreshToken(ctx *gin.Context) {
	var refresh *models.RefreshToken

	err := ctx.BindJSON(&refresh)
	if err != nil {
		h.Response(ctx, http.StatusBadRequest, err, nil)
		return
	}

	login, err := h.services.Auth.RefreshToken(ctx, refresh.RefreshToken)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, login)
}
