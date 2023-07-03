//go:build ignore
// +build ignore

package repo

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type GroupRepo struct {
	pg *db.Query
}

func NewGroupRepo(pg *db.Query) GroupRepo {
	return GroupRepo{pg: pg}
}

func (r GroupRepo) Create(ctx context.Context, group entity.Group) error {
	err := r.pg.Group.WithContext(ctx).Create(&group)
	return err
}

// TODO: проверить жестко
func (r GroupRepo) JoinStudent(ctx context.Context, group entity.GroupJoin) error {
	g, err := r.pg.Group.WithContext(ctx).Where(r.pg.Group.Id.Eq(group.GroupId)).Take()
	if err != nil {
		return err
	}
	err = r.pg.Group.Students.Model(g).Append(&entity.User{Id: group.UserId})
	return err
}

// func (r GroupRepo) CreateGroupTeacher(ctx context.Context, teacher entity.GroupTeacher) error {
// 	sql, args, err := r.pg.Builder.
// 		Insert("group_teacher").Columns("id", "user_id", "group_id").
// 		Values(teacher.Id, teacher.UserId, teacher.GroupId).ToSql()
// 	if err != nil {
// 		return err
// 	}
// 	_, err = r.pg.Pool.Exec(ctx, sql, args...)
// 	return err
// }
