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
}

type RepositoryDB interface {
	InsertAnekdotList(ctx context.Context, anekdotList []models.Anekdot) error
	ChangeRating(ctx context.Context, anekdotID int, method MethodRaitng) error
}

func NewRepositories(pg clientRepo.PostgresClient) *Repositories {
	return &Repositories{}
}
