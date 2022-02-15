package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidAccessToken = errors.New("invalid access token")
)

type Auth struct {
	saltPassword string
	signedKey    string
	tokenExpires time.Duration
	UserDB       repository.UserDB
}

// GetUserFromRequest return user from DB
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

	claims, err := ParseToken(headerParts[1], []byte(a.signedKey))
	if err != nil {
		return nil, err
	}

	return a.UserDB.GetUserByID(ctx, claims.ID)
}

func ParseToken(accessToken string, signedKey []byte) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return signedKey, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidAccessToken
}

func NewAuth(cfg *config.Config, userDB *repository.UserDB) *Auth {
	return &Auth{
		saltPassword: cfg.UserPasswordSalt,
		tokenExpires: cfg.TokenExpires,
		signedKey:    cfg.TokenSignedKey,
		UserDB:       *userDB,
	}
}
