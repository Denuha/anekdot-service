package service

import (
	"context"

	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
)

type user struct {
	repository.UserDB
}

func (u *user) GetUserList(ctx context.Context) ([]models.User, error) {
	return u.UserDB.GetUserList(ctx)
}

func NewUserService(repos *repository.Repositories) User {
	return &user{
		UserDB: repos.UserDB,
	}
}
