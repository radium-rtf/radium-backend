package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type User struct {
}

func NewUserRepo(pg *postgres.Postgres) User {
	return User{}
}

func (r User) Create(ctx context.Context, user *entity.User) error {
	panic("not implemented")
}

func (r User) get(ctx context.Context, cond ...any) (*entity.User, error) {
	panic("not implemented")
}

func (r User) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	panic("not implemented")
}

func (r User) GetById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	panic("not implemented")
}

func (r User) GetByVerificationCode(ctx context.Context, verificationCode string) (*entity.User, error) {
	panic("not implemented")
}

func (r User) GetByRefreshToken(ctx context.Context, refreshToken string) (*entity.User, error) {
	panic("not implemented")
}

func (r User) updateColumn(ctx context.Context, where any, column any, value interface{}) error {
	panic("not implemented")
}

func (r User) SetVerificationCode(ctx context.Context, id uuid.UUID, code string) error {
	panic("not implemented")
}

func (r User) Verify(ctx context.Context, id uuid.UUID) error {
	panic("not implemented")
}

func (r User) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	panic("not implemented")
}

func (r User) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	panic("not implemented")
}
