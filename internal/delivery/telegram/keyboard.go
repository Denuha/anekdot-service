package telegram

import (
	"fmt"

	"github.com/Denuha/anekdot-service/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func createKeyboardRating(anekdot *models.Anekdot) tgbotapi.InlineKeyboardMarkup {
	btnLike := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👍 %d", anekdot.Likes),
		fmt.Sprintf("rating:%d:like", anekdot.ID))

	btnDislike := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👎 %d", anekdot.Dislikes),
		fmt.Sprintf("rating:%d:dislike", anekdot.ID))

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	buttons = append(buttons, btnLike, btnDislike)

	return tgbotapi.NewInlineKeyboardMarkup(
		buttons,
	)
}
