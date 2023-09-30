package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/repoerr"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
	"time"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Session struct {
	db *bun.DB
}

func NewSessionRepo(pg *postgres.Postgres) Session {
	return Session{db: pg.DB}
}

func (r Session) Create(ctx context.Context, session entity.Session) error {
	_, err := r.db.NewInsert().Model(&session).Exec(ctx)
	return err
}

func (r Session) Update(ctx context.Context, refreshToken uuid.UUID, expiresIn time.Time) error {
	exec, err := r.db.NewUpdate().
		Table("sessions").
		Where("refresh_token = ?", refreshToken).
		Set("expires_in = ?", expiresIn).Exec(ctx)

	rowsAffected, _ := exec.RowsAffected()
	if err == nil && rowsAffected == 0 {
		return repoerr.SessionNotFound
	}
	return err
}
