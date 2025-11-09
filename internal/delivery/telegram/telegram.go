package telegram

import (
	"github.com/Denuha/anekdot-service/internal/repository"
	"github.com/Denuha/anekdot-service/internal/service"
	"github.com/sirupsen/logrus"
)

type Telegram struct {
	services *service.Services
	log      *logrus.Logger

	CommonDB repository.Common
	UserDB   repository.User
}

func NewTelegramDelivery(services *service.Services, log *logrus.Logger, repos *repository.Repositories) *Telegram {
	return &Telegram{
		services: services,
		log:      log,
		CommonDB: repos.Common,
		UserDB:   repos.User,
	}
}
