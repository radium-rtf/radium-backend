package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"time"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Session struct {
}

func NewSessionRepo(pg *postgres.Postgres) Session {
	return Session{}
}

func (r Session) Create(ctx context.Context, session entity.Session) error {
	panic("not implemented")
}

func (r Session) Update(ctx context.Context, refreshToken string, expiresIn time.Time) error {
	panic("not implemented")
}
