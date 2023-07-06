package repo

import (
	"context"
	"github.com/google/uuid"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type GroupRepo struct {
	pg *db.Query
}

func NewGroupRepo(pg *db.Query) GroupRepo {
	return GroupRepo{pg: pg}
}

func (r GroupRepo) Create(ctx context.Context, group *entity.Group) (*entity.Group, error) {
	g := r.pg.Group
	err := g.WithContext(ctx).
		Preload(g.Courses, g.Students).
		Create(group)
	if err != nil {
		return nil, err
	}
	return r.GetById(ctx, group.Id)
}

func (r GroupRepo) GetById(ctx context.Context, id uuid.UUID) (*entity.Group, error) {
	g := r.pg.Group
	return g.WithContext(ctx).
		Preload(g.Students).
		Where(g.Id.Eq(id)).
		First()
}

func (r GroupRepo) GetByInviteCode(ctx context.Context, code string) (*entity.Group, error) {
	g := r.pg.Group
	return g.WithContext(ctx).
		Where(g.InviteCode.Eq(code)).
		First()
}

func (r GroupRepo) JoinStudent(ctx context.Context, invite entity.GroupJoin) error {
	g, err := r.GetByInviteCode(ctx, invite.InviteCode)
	if err != nil {
		return err
	}

	err = r.pg.Group.Students.Model(g).
		Append(&entity.User{DBModel: entity.DBModel{Id: invite.UserId}})
	return err
}

func (r GroupRepo) Get(ctx context.Context) ([]*entity.Group, error) {
	return r.pg.Group.WithContext(ctx).Find()
}
