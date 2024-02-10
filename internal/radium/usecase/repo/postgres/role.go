package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Role struct {
	db *bun.DB

	user User
}

func NewRoleRepo(pg *postgres.Postgres) Role {
	return Role{db: pg.DB, user: NewUserRepo(pg)}
}

func (r Role) AddTeacher(ctx context.Context, email string) error {
	set := columnValue{column: "is_teacher", value: true}
	return r.setRole(ctx, set, email)
}

func (r Role) AddAuthor(ctx context.Context, email string) error {
	set := columnValue{column: "is_author", value: true}
	return r.setRole(ctx, set, email)
}

func (r Role) setRole(ctx context.Context, set columnValue, email string) error {
	user, err := r.user.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	_, err = r.db.NewUpdate().
		Table("roles").
		Where("user_id = ?", user.Id).
		Set(set.column+" = ?", set.value).
		Exec(ctx)
	return err
}

func (r Role) AddCoauthor(ctx context.Context, email string, courseId uuid.UUID) error {
	user, err := r.user.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	err = r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err = tx.NewUpdate().
			Table("roles").
			Where("user_id = ?", user.Id).
			Set("is_coauthor = ?", true).
			Exec(ctx)
		courseCoauthor := &entity.CourseCoauthor{UserId: user.Id, CourseId: courseId}
		_, err = tx.NewInsert().Model(courseCoauthor).Exec(ctx)
		return err
	})
	return err
}

func (r Role) DeleteCoauthor(ctx context.Context, id uuid.UUID, courseId uuid.UUID) error {
	_, err := r.db.NewDelete().
		Model(&entity.CourseCoauthor{}).
		Where("user_id = ? and course_id = ?", id, courseId).
		Exec(ctx)
	return err
}
