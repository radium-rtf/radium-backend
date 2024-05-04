package usecase

import (
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type GroupUseCase struct {
	groupChat postgres2.Group
}

func NewGroupUseCase(groupChatRepo postgres2.Group) GroupUseCase {
	return GroupUseCase{groupChat: groupChatRepo}
}
