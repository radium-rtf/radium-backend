package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type MessageUseCase struct {
	message postgres2.Message
}

func (uc MessageUseCase) GetMessage(ctx context.Context) (*entity.Message, error) {
	message, err := uc.message.Get(ctx)
	return message, err
}

func (uc MessageUseCase) GetMessagesFrom(ctx context.Context, chatId uuid.UUID) ([]*entity.Message, error) {
	message, err := uc.message.Get(ctx)
	return []*entity.Message{message}, err
}

func (uc MessageUseCase) SendMessage(ctx context.Context, chatId uuid.UUID, content model.Content) (*entity.Message, error) {
	message, err := uc.message.Get(ctx)
	return message, err
}

func NewMessageUseCase(messageRepo postgres2.Message) MessageUseCase {
	return MessageUseCase{message: messageRepo}
}
