package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type GroupUseCase struct {
	group postgres2.Group
	chat  ChatUseCase
}

func (uc GroupUseCase) GetGroup(ctx context.Context, chatId uuid.UUID) (*entity.Group, error) {
	group, err := uc.group.Get(ctx, chatId)
	return group, err
}

func (uc GroupUseCase) CreateGroup(ctx context.Context, group *entity.Group) error {
	err := uc.group.Create(ctx, group)
	if err != nil {
		return err
	}
	err = uc.chat.CreateFromGroup(ctx, group)
	return err
}

func (uc GroupUseCase) AddMember(ctx context.Context, groupId, userId uuid.UUID, admin bool) error {
	return uc.group.AddMember(ctx, groupId, userId, admin)
}

func (uc GroupUseCase) RemoveMember(ctx context.Context, groupId, userId uuid.UUID) error {
	return uc.group.RemoveMember(ctx, groupId, userId)
}

func NewGroupUseCase(groupRepo postgres2.Group, chatUC ChatUseCase) GroupUseCase {
	return GroupUseCase{group: groupRepo, chat: chatUC}
}
