package usecase

import (
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type GroupChatUseCase struct {
	groupChat postgres2.GroupChat
}

func NewGroupChatUseCase(groupChatRepo postgres2.GroupChat) GroupChatUseCase {
	return GroupChatUseCase{groupChat: groupChatRepo}
}
