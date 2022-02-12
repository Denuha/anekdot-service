package repository

import (
	"context"

	"github.com/Denuha/anekdot-service/internal/models"
	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
)

type MethodRaitng int

const (
	MethodIncr MethodRaitng = iota
	MethodDecr
)

type Repositories struct {
	AnekdotDB AnekdotDB
}

type AnekdotDB interface {
	InsertAnekdotList(ctx context.Context, anekdotList []models.Anekdot) error
	ChangeRating(ctx context.Context, anekdotID int, method MethodRaitng) error
	GetRandomAnekdot(ctx context.Context) (*models.Anekdot, error)
	GetAnekdotByID(ctx context.Context, anekdotID int) (*models.Anekdot, error)
}

func NewRepositories(pg clientRepo.PostgresClient) *Repositories {
	return &Repositories{
		AnekdotDB: NewAnekdotRepo(pg),
	}
}
