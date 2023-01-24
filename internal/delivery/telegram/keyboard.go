package telegram

import (
	"fmt"

	"github.com/Denuha/anekdot-service/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func createKeyboardRating(anekdot *models.Anekdot) tgbotapi.InlineKeyboardMarkup {
	btnLike := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👍 %d", anekdot.Likes),
		fmt.Sprintf("rating:%d:%s", anekdot.ID, btnRatingLike))

	btnDislike := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("👎 %d", anekdot.Dislikes),
		fmt.Sprintf("rating:%d:%s", anekdot.ID, btnRatingDislike))

	btnSkip := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("🤨 %d", anekdot.Skips),
		fmt.Sprintf("rating:%d:%s", anekdot.ID, btnRatingSkip))

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
