package repository

import (
	"context"

	"github.com/Denuha/anekdot-service/internal/models"
	service "github.com/Denuha/anekdot-service/internal/service/parser"
)

type RepositoryDB interface {
	InsertAnekdotList(ctx context.Context, anekdotList []models.Anekdot) error
	ChangeRating(ctx context.Context, anekdotID int, method service.MethodRaitng) error
}
