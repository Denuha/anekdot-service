package service

import "github.com/Denuha/anekdot-service/internal/models"

type Parser interface {
	ParseAnekdots() ([]models.Anekdot, error)
}
