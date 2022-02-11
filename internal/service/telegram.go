package service

import (
	"github.com/Denuha/anekdot-service/internal/config"
	"github.com/Denuha/anekdot-service/internal/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type telegram struct {
	log *logrus.Logger
	cfg config.Config
	bot *tgbotapi.BotAPI
}

func (t *telegram) TelegramBot() error {
	// bot, err := tgbotapi.NewBotAPI(t.cfg.TelegramToken)
	// t.bot = bot

	// if err != nil {
	// 	return fmt.Errorf("bot api: %s", err.Error())
	// }
	// t.log.Println("Telegram server is started")

	// //Set update timeout
	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	//Get updates from bot
	//updates := bot.GetUpdatesChan(u)

	return nil
}

func (t *telegram) GetBot() *tgbotapi.BotAPI {
	return t.bot
}

func NewTelegramService(repos *repository.Repositories, cfg *config.Config,
	log *logrus.Logger) Telegram {
	return &telegram{
		log: log,
		cfg: *cfg,
	}
}
