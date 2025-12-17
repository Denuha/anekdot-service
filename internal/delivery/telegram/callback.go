package telegram

import (
	"context"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// callbackMsg изменяет рейтинг и возвращает новое сообщение для изменения клавиатуры
func (t *Telegram) callbackMsg(ctx context.Context, query *tgbotapi.CallbackQuery) *tgbotapi.EditMessageReplyMarkupConfig {
	split := strings.Split(query.Data, ":")

	// Если пришел не рейтинг
	if split[0] != "rating" {
		return nil
	}
	anekdotID, err := strconv.Atoi(split[1])
	if err != nil {
		t.log.Errorln(err)
		return nil
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
		t.log.Errorln(err)
		return nil
	}

	anekdot, err := t.services.Anekdot.GetAnekdotByID(ctx, anekdotID)
	if err != nil {
		t.log.Errorln(err)
		return nil
	}
	replyMarkupUpdate := createKeyboardRating(anekdot)
	updateMsg := tgbotapi.NewEditMessageReplyMarkup(
		query.Message.Chat.ID,
		query.Message.MessageID,
		replyMarkupUpdate,
	)

	return &updateMsg
}

// anekMsg анекдот и клавиатура с голосами
func (t *Telegram) anekMsg(ctx context.Context, chatID int64) tgbotapi.MessageConfig {
	anekdot, err := t.services.Anekdot.GetRandomAnekdot(ctx)
	if err != nil {
		t.log.Errorln(err)
	}
	message := anekdot.Text
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyMarkup = createKeyboardRating(anekdot)

	return msg
}
