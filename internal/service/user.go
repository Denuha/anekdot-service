package service

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
	"github.com/dgrijalva/jwt-go"
)

type user struct {
	repository.UserDB
	saltPassword string
	signedKey    string
	tokenExpires time.Duration
}

func (u *user) GetUserList(ctx context.Context) ([]models.User, error) {
	return u.UserDB.GetUserList(ctx)
}

func (u *user) Registration(ctx context.Context, user *models.UserRegistation) (int, error) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(u.saltPassword))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))
	user.Realm = "user"

	return u.UserDB.CreateUser(ctx, user)
}

func (u *user) Login(ctx context.Context, user *models.UserLogin) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(u.saltPassword))

	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))

	id, err := u.UserDB.SelectLogin(ctx, user.Username, user.Realm, user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("wrong credentials")
		}
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(u.tokenExpires.Seconds()),
			IssuedAt:  time.Now().Unix(),
		},
		UserName: user.Username,
		Realm:    user.Realm,
		ID:       id,
	})

	return token.SignedString([]byte(u.signedKey))
}

func NewUserService(repos *repository.Repositories, cfg *config.Config) User {
	return &user{
		UserDB:       repos.UserDB,
		saltPassword: cfg.UserPasswordSalt,
		tokenExpires: cfg.TokenExpires,
		signedKey:    cfg.TokenSignedKey,
	}
}
