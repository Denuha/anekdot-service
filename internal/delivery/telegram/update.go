package telegram

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/Denuha/anekdot-service/internal/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) ProcessUpdates(updates *tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) {
	for update := range *updates {
		if update.Message != nil {
			log.Printf("[%s] message %s", update.Message.From.String(), update.Message.Text)

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

		if update.CallbackQuery != nil {
			log.Printf("[%s] callback %s", update.CallbackQuery.From.String(), update.CallbackQuery.Data)

			ctx := context.Background()
			userTg := processUser(&update)
			userDB, err := t.checkUser(ctx, userTg)
			if err != nil {
				log.Println(err)
			}

			ctx, err = t.userUtls.PutUserToContext(ctx, userDB)
			if err != nil {
				log.Println(err)
			}

			t.callbackQueryHandler(ctx, update.CallbackQuery)
			msg := t.processCommandRandomCallback(ctx, &update)
			bot.Send(msg)
		}
	}
}

// processUser tgUser2serviceUser
func processUser(update *tgbotapi.Update) *models.User {
	var tgUSer *tgbotapi.User

	if update.Message != nil {
		tgUSer = update.Message.From
	}
	if update.CallbackQuery != nil {
		tgUSer = update.CallbackQuery.From
	}

	user := models.User{
		UserName:   tgUSer.String(),
		ExternalID: strconv.Itoa(int(tgUSer.ID)),
		Realm:      "tg",
	}

	return &user
}

// tg middleware
// return user from DB
func (t *Telegram) checkUser(ctx context.Context, user *models.User) (*models.User, error) {
	tx, err := t.CommonDB.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	var userDB *models.User

	userDB, err = t.UserDB.GetUserByRealmAndExternalID(ctx, tx, user.Realm, user.ExternalID)
	if err != nil {
		if err == sql.ErrNoRows {
			userDB, err = t.UserDB.InsertUser(ctx, tx, user)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	err = t.CommonDB.CommitTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}
	return userDB, nil
}
