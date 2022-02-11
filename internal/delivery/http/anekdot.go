package http

import (
	"errors"
	"net/http"

	"github.com/Denuha/anekdot-service/internal/service/parser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initRoutesAnekdot(rg *gin.RouterGroup) {
	anekdotGroup := rg.Group("/anekdot")

	anekdotGroup.GET("/parse", h.parseAnekdots)
}

func (h *Handler) parseAnekdots(ctx *gin.Context) {
	var source string
	var parserClient parser.Parser

	paramMap := ctx.Request.URL.Query()
	sourceQuery := paramMap["source"]
	if len(sourceQuery) != 0 {
		source = sourceQuery[0]

		if source == "" {
			source = "anekdotme"
		}
	}

	switch source {
	case "anekdotme":
		pAnekdotme := parser.NewParserAnekdotme()
		ps := parser.NewParserService(pAnekdotme)
		parserClient = ps.Parser
	default:
		h.Response(ctx, http.StatusBadRequest, errors.New("source is wrong"), nil)
		return
	}

	anekdots, err := parserClient.ParseAnekdots()
	if err != nil {
		h.Response(ctx, http.StatusInternalServerError, err, nil)
		return
	}

	h.Response(ctx, http.StatusOK, nil, anekdots)
}
