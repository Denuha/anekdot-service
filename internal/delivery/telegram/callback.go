package telegram

import (
	"context"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) callbackQueryHandler(ctx context.Context, query *tgbotapi.CallbackQuery) {
	split := strings.Split(query.Data, ":")
	if split[0] == "rating" {

		anekdotID, err := strconv.Atoi(split[1])
		if err != nil {
			log.Println(err)
			return
		}

		var value int
		valueStr := split[2]
		switch valueStr {
		case "like":
			value = 1
		case "dislike":
			value = -1
		default:
			value = 0
		}

		err = t.services.Anekdot.UpdateRating(ctx, anekdotID, value)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
}

func (t *Telegram) processCommandRandomCallback(ctx context.Context, update *tgbotapi.Update) tgbotapi.MessageConfig {
	anekdot, err := t.services.Anekdot.GetRandomAnekdot(ctx)
	if err != nil {
		t.log.Println(err)
	}
	message := anekdot.Text
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
	msg.ReplyMarkup = createKeyboardRandomAnekdot(anekdot)

	return msg
}
