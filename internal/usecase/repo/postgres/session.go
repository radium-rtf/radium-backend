package postgres

import (
	"context"
	"time"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Session struct {
	pg *db.Query
}

func NewSessionRepo(pg *db.Query) Session {
	return Session{pg: pg}
}

func (r Session) Create(ctx context.Context, session entity.Session) error {
	err := r.pg.Session.WithContext(ctx).Create(&session)
	if err != nil {
		return err
	}
	return nil
}

func (r Session) Update(ctx context.Context, refreshToken string, expiresIn time.Time) error {
	s := r.pg.Session
	_, err := s.WithContext(ctx).Where(s.RefreshToken.Eq(refreshToken)).Update(s.ExpiresIn, expiresIn)
	if err != nil {
		return err
	}
	return nil
}
