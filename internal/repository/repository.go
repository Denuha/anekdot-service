package repository

import (
	"context"
	"database/sql"

	"github.com/Denuha/anekdot-service/internal/models"
	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
)

type MethodRaitng int

type Repositories struct {
	Anekdot Anekdot
	Common  Common
	User    User
	Metrics Metrics
	Session Session
}

type Anekdot interface {
	InsertAnekdotList(ctx context.Context, anekdotList []models.Anekdot) error
	GetRandomAnekdot(ctx context.Context, user *models.User) (*models.Anekdot, error)
	GetAnekdotByID(ctx context.Context, anekdotID int) (*models.Anekdot, error)

	GetUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64) (*models.AnekdotVote, error)
	UpdateUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64, value int) error
	PostUserVoteByAnekdotID(ctx context.Context, anekdotID int, userID int64, value int) error
}

type User interface {
	InsertUser(ctx context.Context, tx *sql.Tx, user *models.User) (*models.User, error)
	UpdateChatID(ctx context.Context, tx *sql.Tx, userID int64, chatID *int64) error
	GetUserByRealmAndExternalID(ctx context.Context, tx *sql.Tx, realm, externalID string) (*models.User, error)
	GetUserList(ctx context.Context) ([]models.User, error)
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	CreateUser(ctx context.Context, user *models.UserRegistation) (int, error)
	SelectLogin(ctx context.Context, username, realm, pass string) (int, error)
}

type Common interface {
	BeginTransaction(ctx context.Context) (*sql.Tx, error)
	CommitTransaction(ctx context.Context, tx *sql.Tx) error
}

type Metrics interface {
	GetMetrics(ctx context.Context) (*models.Metrics, error)
}

type Session interface {
	GetSession(ctx context.Context, userID int) (*models.Session, error)
	InsertSession(ctx context.Context, session *models.SessionInsert) error
	DeleteSession(ctx context.Context, userID int) error
}

func NewRepositories(pg clientRepo.PostgresClient) *Repositories {
	return &Repositories{
		Anekdot: NewAnekdotRepo(pg),
		Common:  NewCommonRepo(pg),
		User:    NewUserRepo(pg),
		Metrics: NewMetricsRepo(pg),
		Session: NewSessionRepo(pg),
	}
}
