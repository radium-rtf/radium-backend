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
		Order("notification.read desc").
		Scan(ctx)

	if errors.Is(sql.ErrNoRows, err) {
		return nil, repoerr.NotFound
	}
	return *notifications, err
}

func (r Notification) Update(ctx context.Context, notification entity.Notification) (entity.Notification, error) {
	info, err := r.db.NewUpdate().
		Model(&notification).
		WherePK().
		OmitZero().
		Exec(ctx)

	if err != nil {
		return notification, err
	}
	if n, _ := info.RowsAffected(); n == 0 {
		return notification, repoerr.NotFound
	}

	return notification, nil
}

func (r Notification) Read(ctx context.Context, ids []uuid.UUID, userId uuid.UUID) (int64, error) {
	info, err := r.db.
		NewUpdate().
		Model(&entity.Notification{}).
		Where("user_id = ? and id in (?)", userId, bun.In(ids)).
		Set("read = ?", true).
		Exec(ctx)

	if err != nil {
		return 0, err
	}
	return info.RowsAffected()
}

func (r Notification) Delete(ctx context.Context, ids []uuid.UUID, userId uuid.UUID) (int64, error) {
	info, err := r.db.
		NewDelete().
		Model(&entity.Notification{}).
		Where("user_id = ? and id in (?)", userId, bun.In(ids)).
		Exec(ctx)

	if err != nil {
		return 0, err
	}
	return info.RowsAffected()
}
