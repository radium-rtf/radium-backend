package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/repoerr"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
	"time"
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
		Set("expires_in = ?", expiresIn).
		Exec(ctx)
	if err != nil {
		return err
	}

	rowsAffected, _ := exec.RowsAffected()
	if rowsAffected == 0 {
		return repoerr.NotFound
	}
	return err
}
