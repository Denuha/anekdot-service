package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initRoutesAnekdot(rg *gin.RouterGroup) {
	anekdotGroup := rg.Group("/anekdot")

	anekdotGroup.GET("/parse", h.parseAnekdots)
	anekdotGroup.GET("/random", h.getRandomAnekdot)
}

func (h *Handler) parseAnekdots(ctx *gin.Context) {
	var source string

	paramMap := ctx.Request.URL.Query()
	sourceQuery := paramMap["source"]
	if len(sourceQuery) != 0 {
		source = sourceQuery[0]

		if source == "" {
			source = "anekdotme"
		}
	}

	count, err := h.services.Anekdot.ParseAnekdots(ctx, source)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, map[string]int{"count": count})
}

func (h *Handler) getRandomAnekdot(ctx *gin.Context) {
	anekdot, err := h.services.Anekdot.GetRandomAnekdot(ctx)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, anekdot)
}
