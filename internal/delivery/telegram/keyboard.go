package telegram

import (
	"fmt"

	"github.com/Denuha/anekdot-service/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func createKeyboardRating(anekdot *models.Anekdot) tgbotapi.InlineKeyboardMarkup {
	btnLike := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👍 %d", anekdot.Likes),
		fmt.Sprintf("rating:%s:%d", btnRatingLike, anekdot.ID))

	btnDislike := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👎 %d", anekdot.Dislikes),
		fmt.Sprintf("rating:%s:%d", btnRatingDislike, anekdot.ID))

	btnSkip := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🤨 %d", anekdot.Skips),
		fmt.Sprintf("rating:%s:%d", btnRatingSkip, anekdot.ID))

	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	buttons = append(buttons, btnLike, btnDislike, btnSkip)

	return tgbotapi.NewInlineKeyboardMarkup(
		buttons,
	)
}

type btnRating string

const (
	btnRatingLike    btnRating = "like"
	btnRatingDislike btnRating = "dislike"
	btnRatingSkip    btnRating = "skip"
)
