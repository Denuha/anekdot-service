package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initMetricsRoutes(rg *gin.RouterGroup) {
	metricsGroup := rg.Group("/metrics")
	metricsGroup.Use(h.adminVerify())

	metricsGroup.GET("", h.getMetrics)
}

// @Summary Get metrics of app
// @Description
// @Security ApiKeyAuth
// @Tags Metrics
// @ID getMetrics
// @Accept json
// @Produce json
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 200 {object} models.Response{resp=models.Metrics} "OK"
// @Router /metrics [get]
func (h *Handler) getMetrics(ctx *gin.Context) {
	metrics, err := h.services.Metrics.GetMetrics(ctx)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, metrics)
}
