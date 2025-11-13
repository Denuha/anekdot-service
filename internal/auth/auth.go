package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/Denuha/anekdot-service/internal/auth/token"
	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidAccessToken = errors.New("invalid access token")
)

type Auth struct {
	saltPassword string
	signedKey    string
	tokenExpires time.Duration
	UserDB       repository.User
	SessionDB    repository.Session
}

func NewAuth(cfg *config.Config, repos *repository.Repositories) *Auth {
	return &Auth{
		saltPassword: cfg.UserPasswordSalt,
		tokenExpires: cfg.TokenExpires,
		signedKey:    cfg.TokenSignedKey,
		UserDB:       repos.User,
		SessionDB:    repos.Session,
	}
}

// GetUserFromRequest return user from DB.
// Check token and session.
func (a *Auth) GetUserFromRequest(ctx *gin.Context) (*models.User, error) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("empty auth header")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return nil, errors.New("wrong auth header")
	}

	if headerParts[0] != "Bearer" {
		return nil, errors.New("auth method is not Bearer")
	}

	accessToken := headerParts[1]

	claims, err := token.ParseToken(accessToken, []byte(a.signedKey))
	if err != nil {
		return nil, err
	}

	session, err := a.SessionDB.GetSession(ctx, claims.ID)
	if err != nil {
		return nil, errors.New("no access")
	}
	if session.AccessToken != accessToken {
		return nil, errors.New("no access")
	}

	return a.UserDB.GetUserByID(ctx, claims.ID)
}
