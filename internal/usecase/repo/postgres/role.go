package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Role struct {
}

func NewRoleRepo(pg *postgres.Postgres) Role {
	return Role{}
}

func (r Role) AddTeacher(ctx context.Context, email string) error {
	panic("not implemented")
}

func (r Role) AddAuthor(ctx context.Context, email string) error {
	panic("not implemented")
}

func (r Role) addRole(ctx context.Context, field any, email string) error {
	panic("not implemented")
}
