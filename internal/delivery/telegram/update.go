package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) ProcessUpdates(updates *tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range *updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			var msg tgbotapi.MessageConfig

			switch update.Message.Text {
			case "/start":
				msg = t.processCommandStart(&update)
			case "/random":
				msg = t.processCommandRandom(&update)
			case "/help":
				msg = t.processCommandHelp(&update)
			default:
				msg = t.processCommandUnknown(&update)
			}
			bot.Send(msg)
		}
	}
}
