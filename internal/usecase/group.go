package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type GroupUseCase struct {
	groupRepo repo.GroupRepo
}

func NewGroupUseCase(pg *postgres.Postgres) GroupUseCase {
	return GroupUseCase{groupRepo: repo.NewGroupRepo(pg)}
}

func (uc GroupUseCase) Create(ctx context.Context, name entity.GroupName) (entity.Group, error) {
	var group = entity.Group{Id: uuid.NewString(), Name: name.Name}
	return group, uc.groupRepo.Create(ctx, group)
}

func (uc GroupUseCase) Join(ctx context.Context, joinGroup entity.GroupJoin) error {
	return uc.groupRepo.JoinStudent(ctx, joinGroup)
}

func (uc GroupUseCase) CreateTeacher(ctx context.Context, teacher entity.GroupTeacher) error {
	return uc.groupRepo.CreateGroupTeacher(ctx, teacher)
}
