package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"time"
)

type SessionRepo struct {
	pg *postgres.Postgres
}

func NewSessionRepo(pg *postgres.Postgres) SessionRepo {
	return SessionRepo{pg: pg}
}

func (r SessionRepo) Create(ctx context.Context, session entity.Session) error {
	sql, args, err := r.pg.Builder.
		Insert("sessions").
		Columns("refresh_token", "expires_in", "user_id").
		Values(session.RefreshToken, session.ExpiresIn, session.UserId).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r SessionRepo) Update(ctx context.Context, refreshToken string, expiresIn time.Time) error {
	sql, args, err := r.pg.Builder.
		Update("sessions").
		Where(sq.Eq{"refresh_token": refreshToken}).
		Set("expires_in", expiresIn).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
