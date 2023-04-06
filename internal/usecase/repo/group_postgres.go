package repo

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type GroupRepo struct {
	pg *postgres.Postgres
}

func NewGroupRepo(pg *postgres.Postgres) GroupRepo {
	return GroupRepo{pg: pg}
}

func (r GroupRepo) Create(ctx context.Context, group entity.Group) error {
	sql, args, err := r.pg.Builder.
		Insert("groups").Columns("id", "name").
		Values(group.Id, group.Name).ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (r GroupRepo) JoinStudent(ctx context.Context, group entity.GroupJoin) error {
	sql, args, err := r.pg.Builder.
		Insert("group_student").Columns("user_id", "group_id").
		Values(group.UserId, group.GroupId).ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (r GroupRepo) CreateGroupTeacher(ctx context.Context, teacher entity.GroupTeacher) error {
	sql, args, err := r.pg.Builder.
		Insert("group_teacher").Columns("id", "user_id", "group_id").
		Values(teacher.Id, teacher.UserId, teacher.GroupId).ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	return err
}
