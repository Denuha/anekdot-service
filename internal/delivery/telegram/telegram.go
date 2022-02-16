package telegram

import (
	"github.com/Denuha/anekdot-service/internal/repository"
	"github.com/Denuha/anekdot-service/internal/service"
	"github.com/Denuha/anekdot-service/internal/utils"
	"github.com/sirupsen/logrus"
)

type Telegram struct {
	services *service.Services
	log      *logrus.Logger
	userUtls utils.UtilsUser

	CommonDB repository.CommonDB
	UserDB   repository.UserDB
}

func NewTelegramDelivery(services *service.Services, log *logrus.Logger, repos *repository.Repositories) *Telegram {
	return &Telegram{
		services: services,
		log:      log,
		userUtls: utils.NewUtilsUser(),
		CommonDB: repos.CommonDB,
		UserDB:   repos.UserDB,
	}
}
