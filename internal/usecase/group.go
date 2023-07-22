package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type GroupUseCase struct {
	groupRepo postgres.Group
}

func NewGroupUseCase(groupRepo postgres.Group) GroupUseCase {
	return GroupUseCase{groupRepo: groupRepo}
}

func (uc GroupUseCase) Create(ctx context.Context, group *entity.Group) (*entity.Group, error) {
	return uc.groupRepo.Create(ctx, group)
}

func (uc GroupUseCase) Join(ctx context.Context, studentId uuid.UUID, code string) error {
	return uc.groupRepo.JoinStudent(ctx, studentId, code)
}

func (uc GroupUseCase) GetById(ctx context.Context, id uuid.UUID) (*entity.Group, error) {
	return uc.groupRepo.GetById(ctx, id)
}

func (uc GroupUseCase) Get(ctx context.Context) ([]*entity.Group, error) {
	return uc.groupRepo.Get(ctx)
}

// func (uc GroupUseCase) CreateTeacher(ctx context.Context, teacher entity.GroupTeacher) error {
// 	return uc.groupRepo.CreateGroupTeacher(ctx, teacher)
// }
