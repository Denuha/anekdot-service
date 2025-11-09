package service

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
)

type user struct {
	userDB repository.User

	saltPassword string
}

func NewUserService(repos *repository.Repositories, cfg *config.Config) User {
	return &user{
		userDB:       repos.User,
		saltPassword: cfg.UserPasswordSalt,
	}
}

func (u *user) Registration(ctx context.Context, user *models.UserRegistation) (int, error) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(u.saltPassword))
	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))
	user.Realm = "user"

	return u.userDB.CreateUser(ctx, user)
}

func (u *user) GetUserList(ctx context.Context) ([]models.User, error) {
	return u.userDB.GetUserList(ctx)
}
