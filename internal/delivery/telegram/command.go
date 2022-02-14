package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
		return tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
	}

	message := anekdot.Text
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

	msg.ReplyMarkup = createKeyboardRating(anekdot)

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
