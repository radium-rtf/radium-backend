package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
)

type RoleUseCase struct {
	role postgres.Role
}

func NewRoleUseCase(role postgres.Role) RoleUseCase {
	return RoleUseCase{role: role}
}

func (uc RoleUseCase) AddTeacher(ctx context.Context, email string) error {
	return uc.role.AddTeacher(ctx, email)
}

func (uc RoleUseCase) AddAuthor(ctx context.Context, email string) error {
	return uc.role.AddAuthor(ctx, email)
}
