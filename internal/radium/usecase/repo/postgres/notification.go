package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/repoerr"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Notification struct {
	db *bun.DB
}

func NewNotificationRepo(pg *postgres.Postgres) Notification {
	return Notification{db: pg.DB}
}

func (r Notification) Get(ctx context.Context, userId uuid.UUID) ([]entity.Notification, error) {
	var notifications = new([]entity.Notification)

	err := r.db.NewSelect().
		Model(notifications).
		Relation("Answer").
		Relation("Answer.Review").
		Relation("Answer.Review.Reviewer").
		Relation("Answer.Section").
		Relation("Answer.Section.Page").
		Relation("Answer.Section.Page.Module").
		Relation("Answer.Section.Page.Module.Course").
		Where("notification.user_id = ?", userId).
		Scan(ctx)

	if errors.Is(sql.ErrNoRows, err) {
		return nil, repoerr.NotFound
	}
	return *notifications, err
}
