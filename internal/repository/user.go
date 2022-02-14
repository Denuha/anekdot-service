package repository

import (
	"context"
	"database/sql"

	"github.com/Denuha/anekdot-service/internal/models"
	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
)

type userDB struct {
	client clientRepo.PostgresClient
}

func (u *userDB) InsertUser(ctx context.Context, tx *sql.Tx, userInsert *models.User) (*models.User, error) {
	const queryInsertUser = `INSERT INTO anekdot.user ("username", external_id, realm)
	VALUES ($1, $2, $3)
	RETURNING id, username, external_id, realm, create_time;`

	var user models.User

	row := tx.QueryRowContext(ctx, queryInsertUser, userInsert.UserName, userInsert.ExternalID, userInsert.Realm)
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
	SELECT id, username, external_id, realm, create_time FROM anekdot."user" WHERE realm=$1 AND external_id=$2;`

	var user models.User
	row := tx.QueryRowContext(ctx, querySelectUser, realm, externalID)
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

func NewUserRepo(client clientRepo.PostgresClient) UserDB {
	return &userDB{
		client: client}
}
