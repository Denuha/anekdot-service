package http

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initRoutesAnekdot(rg *gin.RouterGroup) {
	anekdotGroup := rg.Group("/anekdot")
	anekdotGroup.Use(h.userVerify())

	anekdotGroup.GET("/parse", h.parseAnekdots)
	anekdotGroup.GET("/random", h.getRandomAnekdot)
	anekdotGroup.GET("/:id", h.getAnekdotByID)
	anekdotGroup.PUT("/:id/rating", h.updateRating)
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

func (h *Handler) updateRating(ctx *gin.Context) {
	var valueStr string
	var value int

	anekdotIDstr := ctx.Param("id")
	anekdotID, err := strconv.Atoi(anekdotIDstr)
	if err != nil {
		h.Response(ctx, http.StatusBadRequest, err, "")
		return
	}

	paramMap := ctx.Request.URL.Query()
	valueQuery := paramMap["value"]
	if len(valueQuery) != 0 {
		valueStr = valueQuery[0]
		switch valueStr {
		case "like":
			value = 1
		case "dislike":
			value = -1
		default:
			h.Response(ctx, http.StatusBadRequest, errors.New("method is not like/dislike"), "")
			return
		}
	}

	if value == 0 {
		h.Response(ctx, http.StatusOK, nil, map[string]int{"value": value})
		return
	}

	err = h.services.Anekdot.UpdateRating(ctx, anekdotID, value)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, "")
		return
	}

	h.Response(ctx, http.StatusOK, nil, map[string]int{"value": value})
}

func (h *Handler) getAnekdotByID(ctx *gin.Context) {
	anekdotIDstr := ctx.Param("id")
	anekdotID, err := strconv.Atoi(anekdotIDstr)
	if err != nil {
		h.Response(ctx, http.StatusBadRequest, err, "")
		return
	}

	annekdot, err := h.services.Anekdot.GetAnekdotByID(ctx, anekdotID)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	h.Response(ctx, http.StatusOK, nil, annekdot)
}
