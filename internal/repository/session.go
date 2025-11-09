package repository

import (
	"context"

	"github.com/Denuha/anekdot-service/internal/models"
	clientRepo "github.com/Denuha/anekdot-service/internal/repository/client"
	sq "github.com/Masterminds/squirrel"
)

type session struct {
	client clientRepo.PostgresClient
}

func NewSessionRepo(client clientRepo.PostgresClient) Session {
	return &session{
		client: client,
	}
}

func (s *session) GetSession(ctx context.Context, userID int) (*models.Session, error) {
	query, args, err := sq.Select(`user_id`,
		`access_token`, `access_token_create_time`,
		`refresh_token`, `refresh_token_create_time`,
	).From(`anekdot."session"`).
		Where(sq.Eq{`user_id`: userID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	cl, err := s.client.GetClient()
	if err != nil {
		return nil, err
	}
	row := cl.QueryRowContext(ctx, query, args...)

	var session models.Session
	err = row.Scan(
		&session.UserID,
		&session.AccessToken,
		&session.AccessTokenCreateTime,
		&session.RefreshToken,
		&session.RefreshTokenCreateTime,
	)

	return &session, err
}

func (s *session) InsertSession(ctx context.Context, session *models.SessionInsert) error {
	query, args, err := sq.Insert(`anekdot."session"`).
		Columns(`user_id`, `access_token`, `refresh_token`).
		Values(session.UserID, session.AccessToken, session.RefreshToken).
		Suffix(`on conflict (user_id)
				do update set
					user_id = EXCLUDED.user_id,
					access_token = EXCLUDED.access_token,
					access_token_create_time = now(),
					refresh_token  = EXCLUDED.refresh_token,
					refresh_token_create_time = now()`).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	cl, err := s.client.GetClient()
	if err != nil {
		return err
	}
	_, err = cl.ExecContext(ctx, query, args...)
	return err
}

func (s *session) DeleteSession(ctx context.Context, userID int) error {
	query, args, err := sq.Delete(`anekdot."session"`).
		Where(sq.Eq{`user_id`: userID}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	cl, err := s.client.GetClient()
	if err != nil {
		return err
	}
	_, err = cl.ExecContext(ctx, query, args...)
	return err
}
