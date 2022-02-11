package server

import (
	"context"
	"fmt"
	"log"

	"github.com/Denuha/anekdot-service/internal/config"
	delivery "github.com/Denuha/anekdot-service/internal/delivery/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegranServer interface {
	Run() error
	Stop(ctx context.Context) error
}

type serverTelegram struct {
	bot *tgbotapi.BotAPI
	tg  *delivery.Telegram
	cfg config.Config
}

func (s *serverTelegram) Run() error {
	bot, err := tgbotapi.NewBotAPI(s.cfg.TelegramToken)
	s.bot = bot

	if err != nil {
		return fmt.Errorf("bot api: %s", err.Error())
	}
	log.Println("Telegram server is started")

	//Set update timeout
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := s.bot.GetUpdatesChan(u)
	s.tg.ProcessUpdates(&updates, s.bot)

	return fmt.Errorf("telegram server is off: %s", err.Error())
}

func (s *serverTelegram) Stop(ctx context.Context) error {
	return nil
}

func NewTelegramServer(cfg config.Config, tg *delivery.Telegram) TelegranServer {
	return &serverTelegram{
		cfg: cfg,
		tg:  tg,
	}
}
