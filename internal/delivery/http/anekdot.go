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
	anekdotGroup.GET("/:anekdotID", h.getAnekdotByID)
	anekdotGroup.PUT("/:anekdotID/rating", h.updateRating)
}

// @Summary Parse anekdots to db
// @Description source="anekdotme"
// @Security ApiKeyAuth
// @Tags Anekdot
// @ID parseAnekdots
// @Accept json
// @Produce json
// @Param source query string true "source"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 200 {object} models.Response "OK"
// @Router /anekdot/parse [get]
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

// @Summary Get random anekdot
// @Description
// @Security ApiKeyAuth
// @Tags Anekdot
// @ID getRandomAnekdot
// @Accept json
// @Produce json
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 200 {object} models.Response{resp=models.Anekdot} "OK"
// @Router /anekdot/random [get]
func (h *Handler) getRandomAnekdot(ctx *gin.Context) {
	anekdot, err := h.services.Anekdot.GetRandomAnekdot(ctx)
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, anekdot)
}

// @Summary Update rating
// @Description value="like/dislike"
// @Security ApiKeyAuth
// @Tags Anekdot
// @ID updateRating
// @Accept json
// @Produce json
// @Param anekdotID path int true "anekdot ID"
// @Param value query string true "value"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 200 {object} models.Response "OK"
// @Router /anekdot/{anekdotID}/rating [put]
func (h *Handler) updateRating(ctx *gin.Context) {
	var valueStr string
	var value int

	anekdotIDstr := ctx.Param("anekdotID")
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
			h.Response(ctx, http.StatusBadRequest, errors.New("value is not like/dislike"), "")
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

// @Summary Get random by ID
// @Description
// @Security ApiKeyAuth
// @Tags Anekdot
// @ID getAnekdotByID
// @Accept json
// @Produce json
// @Param anekdotID path int true "anekdot ID"
// @Failure 500 {object} models.Response "Internal Server Error"
// @Failure 400 {object} models.Response "Bad Request"
// @Failure 401 {object} models.Response "Unauthorized"
// @Failure 403 {object} models.Response "Forbidden"
// @Success 200 {object} models.Response{resp=models.Anekdot} "OK"
// @Router /anekdot/{anekdotID} [get]
func (h *Handler) getAnekdotByID(ctx *gin.Context) {
	anekdotIDstr := ctx.Param("anekdotID")
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
