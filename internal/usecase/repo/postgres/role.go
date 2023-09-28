package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"gorm.io/gen/field"
)

type Role struct {
	pg *db.Query
}

func NewRoleRepo(pg *postgres.Postgres) Role {
	return Role{pg: pg.Q}
}

func (r Role) AddTeacher(ctx context.Context, email string) error {
	return r.addRole(ctx, r.pg.User.IsTeacher, email)
}

func (r Role) AddAuthor(ctx context.Context, email string) error {
	return r.addRole(ctx, r.pg.User.IsAuthor, email)
}

func (r Role) addRole(ctx context.Context, field field.Expr, email string) error {
	u := r.pg.User
	user, err := u.WithContext(ctx).Where(u.Email.Eq(email)).First()
	if err != nil {
		return err
	}

	_, err = u.WithContext(ctx).
		Where(u.Id.Eq(user.Id)).
		UpdateColumn(field, true)

	return err
}
