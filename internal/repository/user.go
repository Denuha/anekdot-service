package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Denuha/anekdot-service/internal/models"
	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
)

type userDB struct {
	client clientRepo.PostgresClient
}

func (u *userDB) InsertUser(ctx context.Context, tx *sql.Tx, userInsert *models.User) (*models.User, error) {
	const queryInsertUser = `INSERT INTO anekdot.user ("username", external_id, realm, chat_id)
	VALUES ($1, $2, $3, $4)
	RETURNING id, username, external_id, realm, create_time;`

	var user models.User

	row := tx.QueryRowContext(ctx, queryInsertUser, userInsert.UserName, userInsert.ExternalID,
		userInsert.Realm, userInsert.ChatID)
	err := row.Scan(
		&user.ID,
		&user.UserName,
		&user.ExternalID,
		&user.Realm,
		&user.CreateTime,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (u *userDB) GetUserByRealmAndExternalID(ctx context.Context, tx *sql.Tx, realm, externalID string) (*models.User, error) {
	const querySelectUser = `
	SELECT id, username, external_id, realm, create_time, chat_id 
	FROM anekdot."user" 
	WHERE realm=$1 AND external_id=$2;`

	var user models.User
	row := tx.QueryRowContext(ctx, querySelectUser, realm, externalID)
	err := row.Scan(
		&user.ID,
		&user.UserName,
		&user.ExternalID,
		&user.Realm,
		&user.CreateTime,
		&user.ChatID,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userDB) GetUserList(ctx context.Context) ([]models.User, error) {
	const querySelect = `SELECT id, username, external_id, realm, create_time, chat_id 
	FROM anekdot."user";`

	users := make([]models.User, 0)

	cl, err := u.client.GetClient()
	if err != nil {
		return users, err
	}

	rows, err := cl.QueryContext(ctx, querySelect)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var tmp models.User
		err = rows.Scan(
			&tmp.ID,
			&tmp.UserName,
			&tmp.ExternalID,
			&tmp.Realm,
			&tmp.CreateTime,
			&tmp.ChatID,
		)

		if err != nil {
			return users, err
		}
		users = append(users, tmp)
	}

	return users, nil
}

func (u *userDB) GetUserByID(ctx context.Context, userID int) (*models.User, error) {
	const querySelect = `SELECT id, username, external_id, realm, create_time, is_admin, chat_id
	FROM anekdot.user u
	WHERE u.id = $1;`

	cl, err := u.client.GetClient()
	if err != nil {
		return nil, err
	}

	var user models.User

	row := cl.QueryRowContext(ctx, querySelect, userID)
	err = row.Scan(
		&user.ID,
		&user.UserName,
		&user.ExternalID,
		&user.Realm,
		&user.CreateTime,
		&user.IsAdmin,
		&user.ChatID,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userDB) CreateUser(ctx context.Context, user *models.UserRegistation) (int, error) {
	const queryInsertUser = `INSERT INTO anekdot.user ("username", realm, password)
	VALUES ($1, $2, $3)
	RETURNING id;`

	cl, err := u.client.GetClient()
	if err != nil {
		return 0, err
	}

	var id int

	row := cl.QueryRowContext(ctx, queryInsertUser, user.Username, user.Realm, user.Password)
	err = row.Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userDB) SelectLogin(ctx context.Context, username, realm, pass string) (int, error) {
	const querySelect = `SELECT id
	FROM anekdot.user u
	WHERE u.username = $1 AND u.realm = $2 AND u.password = $3;`

	cl, err := u.client.GetClient()
	if err != nil {
		return 0, err
	}

	var id int

	row := cl.QueryRowContext(ctx, querySelect, username, realm, pass)
	err = row.Scan(
		&id,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u userDB) UpdateChatID(ctx context.Context, tx *sql.Tx, userID int64, chatID *int64) error {
	const queryUpdate = `UPDATE anekdot.user
	SET chat_id = $1
	WHERE id=$2;`

	res, err := tx.ExecContext(ctx, queryUpdate, chatID, userID)
	if err != nil {
		return err
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		return errors.New("no change")
	}

	return nil
}

func NewUserRepo(client clientRepo.PostgresClient) UserDB {
	return &userDB{
		client: client}
}
