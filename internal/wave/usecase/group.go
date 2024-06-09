package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type GroupUseCase struct {
	group postgres2.Group
	chat  postgres2.Chat
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
	chat := &entity.Chat{
		Id:   group.Id,
		Name: "Group " + group.Id.String(),
		Type: "group",
	}
	err = uc.chat.Create(ctx, chat)
	return err
}

func (uc GroupUseCase) AddMember(ctx context.Context, groupId, userId uuid.UUID, admin bool) error {
	return uc.group.AddMember(ctx, groupId, userId, admin)
}

func NewGroupUseCase(groupRepo postgres2.Group, chatRepo postgres2.Chat) GroupUseCase {
	return GroupUseCase{group: groupRepo, chat: chatRepo}
}
