package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Role struct {
	pg *db.Query
}

func NewRoleRepo(pg *db.Query) Role {
	return Role{pg: pg}
}

func (r Role) AddTeacher(ctx context.Context, email string) error {
	u := r.pg.User
	user, err := u.WithContext(ctx).Where(u.Email.Eq(email)).First()
	if err != nil {
		return err
	}

	_, err = u.WithContext(ctx).
		Where(u.Id.Eq(user.Id)).
		UpdateColumn(u.IsTeacher, true)

	return err
}
