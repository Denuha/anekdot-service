package utils

import (
	"context"
	"errors"

	"github.com/Denuha/anekdot-service/internal/models"
)

// GetUserFromContext получает пользователя из контекста
func GetUserFromContext(ctx context.Context) (*models.User, error) {
	userCtx := ctx.Value("user")
	if userCtx == nil {
		return nil, errors.New("can't get user from context")
	}

	user := userCtx.(*models.User)
	return user, nil
}

func PutUserToContext(ctx context.Context, user *models.User) (context.Context, error) {
	if user == nil {
		return ctx, errors.New("user is nil")
	}

	resCtx := context.WithValue(ctx, "user", user)
	return resCtx, nil
}
