package repo

import (
	"context"
	"time"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type SessionRepo struct {
	pg *db.Query
}

func NewSessionRepo(pg *db.Query) SessionRepo {
	return SessionRepo{pg: pg}
}

func (r SessionRepo) Create(ctx context.Context, session entity.Session) error {
	err := r.pg.Session.WithContext(ctx).Create(&session)
	if err != nil {
		return err
	}
	return nil
}

func (r SessionRepo) Update(ctx context.Context, refreshToken string, expiresIn time.Time) error {
	s := r.pg.Session
	_, err := s.WithContext(ctx).Where(s.RefreshToken.Eq(refreshToken)).Update(s.ExpiresIn, expiresIn)
	if err != nil {
		return err
	}
	return nil
}
