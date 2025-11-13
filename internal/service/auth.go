package service

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Denuha/anekdot-service/internal/auth/token"
	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
	"github.com/sirupsen/logrus"
)

type auth struct {
	userDB    repository.User
	sessionDB repository.Session

	signedKeyAccess     string
	signedKeyRefresh    string
	tokenExpiresAccess  time.Duration
	tokenExpiresRefresh time.Duration

	saltPassword string
}

func NewAuth(repos *repository.Repositories, cfg *config.Config) *auth {
	return &auth{
		userDB:              repos.User,
		sessionDB:           repos.Session,
		signedKeyAccess:     cfg.TokenSignedKey,
		tokenExpiresAccess:  cfg.TokenExpires,
		signedKeyRefresh:    cfg.TokenRefreshSignedKey,
		tokenExpiresRefresh: cfg.TokenRefreshExpires,
		saltPassword:        cfg.UserPasswordSalt,
	}
}

func (a *auth) Login(ctx context.Context, user *models.UserLogin) (*models.Login, error) {
	pwd := sha1.New()
	pwd.Write([]byte(user.Password))
	pwd.Write([]byte(a.saltPassword))

	user.Password = fmt.Sprintf("%x", pwd.Sum(nil))

	userId, err := a.userDB.SelectLogin(ctx, user.Username, user.Realm, user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("wrong credentials")
		}
		return nil, err
	}

	login, err := a.generateLogin(userId, user)
	if err != nil {
		return nil, err
	}

	err = a.sessionDB.InsertSession(ctx, &models.SessionInsert{
		UserID:       userId,
		AccessToken:  login.AccessToken,
		RefreshToken: login.RefreshToken,
	})
	if err != nil {
		return nil, err
	}

	return login, nil
}

func (a *auth) generateLogin(userID int, user *models.UserLogin) (*models.Login, error) {
	now := time.Now()

	accessToken, err := token.GenerateToken(now, userID, user, a.signedKeyAccess, a.tokenExpiresAccess)
	if err != nil {
		return nil, err
	}

	refreshToken, err := token.GenerateToken(now, userID, user, a.signedKeyRefresh, a.tokenExpiresRefresh)
	if err != nil {
		return nil, err
	}

	return &models.Login{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *auth) RefreshToken(ctx context.Context, refreshToken string) (*models.Login, error) {
	claims, err := token.ParseToken(refreshToken, []byte(a.signedKeyRefresh))
	if err != nil {
		return nil, err
	}

	session, err := a.sessionDB.GetSession(ctx, claims.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no access")
		}
		return nil, err
	}

	if session.RefreshToken != refreshToken {
		return nil, errors.New("no access")
	}

	login, err := a.generateLogin(claims.ID, &models.UserLogin{
		Username: claims.UserName,
		Realm:    claims.Realm,
	})
	if err != nil {
		return nil, err
	}

	err = a.sessionDB.InsertSession(ctx, &models.SessionInsert{
		UserID:       claims.ID,
		AccessToken:  login.AccessToken,
		RefreshToken: login.RefreshToken,
	})

	return login, err
}

func (a *auth) Logout(ctx context.Context, accessToken string) error {
	claims, err := token.ParseToken(accessToken, []byte(a.signedKeyAccess))
	if err != nil {
		return err
	}

	session, err := a.sessionDB.GetSession(ctx, claims.ID)
	if err != nil {
		logrus.Errorln(err)
		return errors.New("no access")
	}

	if session.AccessToken != accessToken {
		return errors.New("no access")
	}

	return a.sessionDB.DeleteSession(ctx, claims.ID)
}
