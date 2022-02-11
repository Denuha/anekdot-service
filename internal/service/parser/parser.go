package parser

import "github.com/Denuha/anekdot-service/internal/models"

type Parser interface {
	ParseAnekdots() ([]models.Anekdot, error)
}

type ParserService struct {
	Parser Parser
}

func NewParserService(parser Parser) ParserService {
	p := ParserService{
		Parser: parser,
	}

	return p
}
