package repository

import (
	"context"
	"database/sql"

	"github.com/Denuha/anekdot-service/internal/models"
	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
)

type MethodRaitng int

type Repositories struct {
	AnekdotDB AnekdotDB
	CommonDB  CommonDB
	UserDB    UserDB
}

type AnekdotDB interface {
	InsertAnekdotList(ctx context.Context, anekdotList []models.Anekdot) error
	GetRandomAnekdot(ctx context.Context) (*models.Anekdot, error)
	GetAnekdotByID(ctx context.Context, anekdotID int) (*models.Anekdot, error)

	GetUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64) (*models.AnekdotVote, error)
	UpdateUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64, value int) error
	PostUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64, value int) error
}

type UserDB interface {
	InsertUser(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)
	GetUserByRealmAndExternalID(ctx context.Context, tx *sql.Tx, realm, externalID string) (*models.User, error)
	GetUserList(ctx context.Context) ([]models.User, error)
}

type CommonDB interface {
	BeginTransaction(ctx context.Context) (*sql.Tx, error)
	CommitTransaction(ctx context.Context, tx *sql.Tx) error
}

func NewRepositories(pg clientRepo.PostgresClient) *Repositories {
	return &Repositories{
		AnekdotDB: NewAnekdotRepo(pg),
		CommonDB:  NewCommonRepo(pg),
		UserDB:    NewUserRepo(pg),
	}
}
