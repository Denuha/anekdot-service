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

func (h *Handler) getMetrics(ctx *gin.Context) {
	metrics, err := h.services.Metrics.GetMetrics(ctx)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, metrics)
}
