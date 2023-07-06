package usecase

import (
	"context"
	"github.com/google/uuid"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type GroupUseCase struct {
	groupRepo repo.GroupRepo
}

func NewGroupUseCase(pg *db.Query) GroupUseCase {
	return GroupUseCase{groupRepo: repo.NewGroupRepo(pg)}
}

func (uc GroupUseCase) Create(ctx context.Context, group *entity.Group) (*entity.Group, error) {
	return uc.groupRepo.Create(ctx, group)
}

func (uc GroupUseCase) Join(ctx context.Context, joinGroup entity.GroupJoin) error {
	return uc.groupRepo.JoinStudent(ctx, joinGroup)
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
