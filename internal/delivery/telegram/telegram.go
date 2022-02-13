package telegram

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/Denuha/anekdot-service/internal/repository"
	"github.com/Denuha/anekdot-service/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Telegram struct {
	services *service.Services
	log      *logrus.Logger
}

// /start
func (t *Telegram) processCommandStart(update *tgbotapi.Update) tgbotapi.MessageConfig {
	message := update.Message.From.UserName + `, привет. Я бот с анекдотами.
Отправь команду /random для получение случайного анекдота.`
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	return msg
}

// /random
func (t *Telegram) processCommandRandom(update *tgbotapi.Update) tgbotapi.MessageConfig {
	anekdot, err := t.services.Anekdot.GetRandomAnekdot(context.Background())
	if err != nil {
		t.log.Println(err)
	}
	message := anekdot.Text + "\n<b>Рейтинг</b>: " + strconv.Itoa(anekdot.Rating)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	msg.ParseMode = "html"

	btnLike := tgbotapi.NewInlineKeyboardButtonData("Like", fmt.Sprintf("rating:%d:inc", anekdot.ID))
	btnDislike := tgbotapi.NewInlineKeyboardButtonData("Dislike", fmt.Sprintf("rating:%d:dec", anekdot.ID))

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	buttons = append(buttons, btnLike, btnDislike)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		buttons,
	)

	return msg
}

func (t *Telegram) processCommandRandomCallback(update *tgbotapi.Update) tgbotapi.MessageConfig {
	anekdot, err := t.services.Anekdot.GetRandomAnekdot(context.Background())
	if err != nil {
		t.log.Println(err)
	}
	message := anekdot.Text + "--------------\n<b>Рейтинг</b>: " + strconv.Itoa(anekdot.Rating)

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
	msg.ParseMode = "html"

	btnLike := tgbotapi.NewInlineKeyboardButtonData("Like", fmt.Sprintf("rating:%d:inc", anekdot.ID))
	btnDislike := tgbotapi.NewInlineKeyboardButtonData("Dislike", fmt.Sprintf("rating:%d:dec", anekdot.ID))

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	buttons = append(buttons, btnLike, btnDislike)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		buttons,
	)

	return msg
}

// /help
func (t *Telegram) processCommandHelp(update *tgbotapi.Update) tgbotapi.MessageConfig {
	message := "Спиок комманд:\n/start\n/random\n/help"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	return msg
}

// unknown command
func (t *Telegram) processCommandUnknown(update *tgbotapi.Update) tgbotapi.MessageConfig {
	message := "Я не знаю такую команду. Отправь /help для того, чтобы узнать возможности"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
	return msg
}

func (t *Telegram) callbackQueryHandler(query *tgbotapi.CallbackQuery) {
	split := strings.Split(query.Data, ":")
	if split[0] == "rating" {
		anekdotID, err := strconv.Atoi(split[1])
		if err != nil {
			return
		}

		var method repository.MethodRaitng

		methodStr := split[2]
		switch methodStr {
		case "inc":
			method = repository.MethodIncr
		case "dec":
			method = repository.MethodDecr
		default:
			method = repository.MethodIncr
		}

		err = t.services.Anekdot.UpdateRating(context.Background(), anekdotID, method)
		if err != nil {
			t.log.Println(err)
		}
		return
	}
}

func NewTelegramDelivery(services *service.Services, log *logrus.Logger) *Telegram {
	return &Telegram{
		services: services,
		log:      log,
	}
}
