package telegram

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/Denuha/anekdot-service/internal/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) ProcessUpdates(updates *tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range *updates {
		if update.Message != nil {
			log.Printf("[%s] message %s", update.Message.From.String(), update.Message.Text)

			_, err := t.getSender(&update)
			if err != nil {
				log.Println(err)
			}

			var msg tgbotapi.MessageConfig

			switch update.Message.Text {
			case "/start":
				msg = t.processCommandStart(&update)
			case "/random":
				msg = t.processCommandRandom(&update)
			case "/help":
				msg = t.processCommandHelp(&update)
			case "/metrics":
				msg = t.processCommandMetrics(&update)
			case "/test":
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Тест ОК")
			default:
				msg = t.processCommandUnknown(&update)
			}
			bot.Send(msg)
		}

		if update.CallbackQuery != nil {
			log.Printf("[%s] callback %s", update.CallbackQuery.From.String(), update.CallbackQuery.Data)

			userDB, err := t.getSender(&update)
			if err != nil {
				log.Println(err)
			}

			ctx := context.Background()
			ctx, err = utils.PutUserToContext(ctx, userDB)
			if err != nil {
				log.Println(err)
			}

			t.callbackQueryHandler(ctx, update.CallbackQuery)
			msg := t.callbackRating(ctx, &update)
			bot.Send(msg)
		}
	}
}

// getSender Получает отправителя и добавляет его в локальную БД
func (t *Telegram) getSender(update *tgbotapi.Update) (*models.User, error) {
	var tgUser *tgbotapi.User

	if update.Message != nil {
		tgUser = update.Message.From
	}
	if update.CallbackQuery != nil {
		tgUser = update.CallbackQuery.From
	}

	ctx := context.Background()

	tx, err := t.CommonDB.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	var userDB *models.User

	// dont add group chat
	chatID := &update.FromChat().ID
	if *chatID < 0 {
		chatID = nil
	}

	user := models.User{
		UserName:   tgUser.String(),
		ExternalID: strconv.Itoa(int(tgUser.ID)),
		Realm:      "tg",
		ChatID:     chatID,
	}

	userDB, err = t.UserDB.GetUserByRealmAndExternalID(ctx, tx, user.Realm, user.ExternalID)
	if err != nil {
		if err == sql.ErrNoRows {
			userDB, err = t.UserDB.InsertUser(ctx, tx, &user)
			if err != nil {
				_ = tx.Rollback()
				return nil, err
			}
		} else {
			_ = tx.Rollback()
			return nil, err
		}
	}

	if userDB.ChatID == nil && chatID != nil {
		err = t.UserDB.UpdateChatID(ctx, tx, userDB.ID, chatID)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
	}

	err = t.CommonDB.CommitTransaction(ctx, tx)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	return userDB, nil
}
