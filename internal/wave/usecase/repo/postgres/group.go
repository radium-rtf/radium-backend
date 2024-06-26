package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Group struct {
	db *bun.DB
}

func NewGroupRepo(pg *postgres.Postgres) Group {
	return Group{db: pg.DB}
}

func (r Group) Get(ctx context.Context, groupId uuid.UUID) (*entity.Group, error) {
	var group entity.Group
	err := r.db.NewSelect().Model(&group).
		Relation("Members").
		Where("id = ?", groupId).
		Scan(ctx)
	return &group, err
}

func (r Group) Create(ctx context.Context, group *entity.Group) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(group).Exec(ctx)
		return err
	})
}

func (r Group) Update(ctx context.Context, group *entity.Group) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewUpdate().Model(group).Where("id = ?", group.Id).Exec(ctx)
		return err
	})
}

func (r Group) Delete(ctx context.Context, groupId uuid.UUID) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewDelete().Model(&entity.Group{DBModel: entity.DBModel{Id: groupId}}).Exec(ctx)
		return err
	})
}

func (r Group) AddMember(ctx context.Context, groupId, userId uuid.UUID, admin bool) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(&entity.GroupMember{
			GroupId: groupId,
			UserId:  userId,
			Admin:   admin,
		}).Exec(ctx)
		return err
	})
}

func (r Group) UpdateMember(ctx context.Context, groupId, userId uuid.UUID, admin bool) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewUpdate().Model(&entity.GroupMember{
			GroupId: groupId,
			UserId:  userId,
			Admin:   admin,
		}).Where("group_id = ? AND user_id = ?", groupId, userId).Exec(ctx)
		return err
	})
}

func (r Group) RemoveMember(ctx context.Context, groupId uuid.UUID, userId uuid.UUID) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewDelete().Model(&entity.GroupMember{
			GroupId: groupId,
			UserId:  userId,
		}).Exec(ctx)
		return err
	})
}
