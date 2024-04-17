package usecase

import (
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type MessageUseCase struct {
	message postgres2.Message
}

func NewMessageUseCase(messageRepo postgres2.Message) MessageUseCase {
	return MessageUseCase{message: messageRepo}
}
