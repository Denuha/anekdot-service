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
	MetricsDB MetricsDB
}

type AnekdotDB interface {
	InsertAnekdotList(ctx context.Context, anekdotList []models.Anekdot) error
	GetRandomAnekdot(ctx context.Context, user *models.User) (*models.Anekdot, error)
	GetAnekdotByID(ctx context.Context, anekdotID int) (*models.Anekdot, error)

	GetUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64) (*models.AnekdotVote, error)
	UpdateUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64, value int) error
	PostUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64, value int) error
}

type UserDB interface {
	InsertUser(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)
	UpdateChatID(ctx context.Context, tx *sql.Tx, userID int64, chatID *int64) error
	GetUserByRealmAndExternalID(ctx context.Context, tx *sql.Tx, realm, externalID string) (*models.User, error)
	GetUserList(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	CreateUser(ctx context.Context, user *models.UserRegistation) (int, error)
	SelectLogin(ctx context.Context, username, realm, pass string) (int, error)
}

type CommonDB interface {
	BeginTransaction(ctx context.Context) (*sql.Tx, error)
	CommitTransaction(ctx context.Context, tx *sql.Tx) error
}

type MetricsDB interface {
	GetMetrics(ctx context.Context) (*models.Metrics, error)
}

func NewRepositories(pg clientRepo.PostgresClient) *Repositories {
	return &Repositories{
		AnekdotDB: NewAnekdotRepo(pg),
		CommonDB:  NewCommonRepo(pg),
		UserDB:    NewUserRepo(pg),
		MetricsDB: NewMetricsRepo(pg),
	}
}
