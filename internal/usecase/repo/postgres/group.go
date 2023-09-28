package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres"

	"github.com/google/uuid"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Group struct {
	pg *db.Query
}

func NewGroupRepo(pg *postgres.Postgres) Group {
	return Group{pg: pg.Q}
}

func (r Group) Create(ctx context.Context, group *entity.Group) (*entity.Group, error) {
	g := r.pg.Group
	err := g.WithContext(ctx).
		Preload(g.Courses, g.Students).
		Create(group)
	if err != nil {
		return nil, err
	}
	return r.GetById(ctx, group.Id)
}

func (r Group) GetById(ctx context.Context, id uuid.UUID) (*entity.Group, error) {
	g := r.pg.Group
	return g.WithContext(ctx).
		Preload(g.Students, g.Courses).
		Where(g.Id.Eq(id)).
		First()
}

func (r Group) GetByInviteCode(ctx context.Context, code string) (*entity.Group, error) {
	g := r.pg.Group
	return g.WithContext(ctx).
		Where(g.InviteCode.Eq(code)).
		First()
}

func (r Group) JoinStudent(ctx context.Context, studentId uuid.UUID, code string) error {
	g, err := r.GetByInviteCode(ctx, code)
	if err != nil {
		return err
	}

	err = r.pg.Group.Students.Model(g).
		Append(&entity.User{DBModel: entity.DBModel{Id: studentId}})
	return err
}

func (r Group) Get(ctx context.Context) ([]*entity.Group, error) {
	return r.pg.Group.WithContext(ctx).Find()
}
