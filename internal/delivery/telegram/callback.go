package telegram

import (
	"context"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// callbackQueryHandler обрабатывает запрос из callback
func (t *Telegram) callbackQueryHandler(ctx context.Context, query *tgbotapi.CallbackQuery) {
	split := strings.Split(query.Data, ":")

	// Если пришел рейтинг
	if split[0] == "rating" {

		anekdotID, err := strconv.Atoi(split[1])
		if err != nil {
			log.Println(err)
			return
		}

		var value int
		valueStr := split[2]
		switch btnRating(valueStr) {
		case btnRatingLike:
			value = 1
		case btnRatingDislike:
			value = -1
		case btnRatingSkip:
			value = 0
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

// callback на like/dislike/skip
func (t *Telegram) callbackRating(ctx context.Context, update *tgbotapi.Update) tgbotapi.MessageConfig {
	anekdot, err := t.services.Anekdot.GetRandomAnekdot(ctx)
	if err != nil {
		t.log.Println(err)
	}
	message := anekdot.Text
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message)
	msg.ReplyMarkup = createKeyboardRating(anekdot)

	return msg
}
