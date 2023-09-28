package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"time"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Session struct {
	pg *db.Query
}

func NewSessionRepo(pg *postgres.Postgres) Session {
	return Session{pg: pg.Q}
}

func (r Session) Create(ctx context.Context, session entity.Session) error {
	return r.pg.Session.
		WithContext(ctx).
		Create(&session)
}

func (r Session) Update(ctx context.Context, refreshToken string, expiresIn time.Time) error {
	s := r.pg.Session
	_, err := s.WithContext(ctx).Where(s.RefreshToken.Eq(refreshToken)).Update(s.ExpiresIn, expiresIn)
	return err
}
