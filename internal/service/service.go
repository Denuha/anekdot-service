package service

import (
	"context"

	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Services struct {
	Anekdot Anekdot
	//Telegram Telegram
}

type Anekdot interface {
	ParseAnekdots(ctx context.Context, source string) (int, error)
	GetRandomAnekdot(ctx context.Context) (*models.Anekdot, error)
}

type Telegram interface {
	TelegramBot() error
	GetBot() *tgbotapi.BotAPI
}

func NewServices(cfg *config.Config, repos *repository.Repositories, log *logrus.Logger) *Services {
	return &Services{
		Anekdot: NewAnekdotService(repos),
		//Telegram: NewTelegramService(repos, cfg, log),
	}
}
