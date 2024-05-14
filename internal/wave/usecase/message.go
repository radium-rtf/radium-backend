package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/lib/centrifugo"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type MessageUseCase struct {
	message    postgres2.Message
	centrifugo centrifugo.Centrifugo
}

func (uc MessageUseCase) GetMessage(ctx context.Context) (*entity.Message, error) {
	message, err := uc.message.Get(ctx)
	return message, err
}

func (uc MessageUseCase) GetMessagesFrom(ctx context.Context, chatId uuid.UUID) ([]*entity.Message, error) {
	message, err := uc.message.Get(ctx)
	return []*entity.Message{message}, err
}

func (uc MessageUseCase) SendMessage(ctx context.Context, chatId uuid.UUID, content model.Content) (*model.Message, error) {
	var err error
	client := uc.centrifugo.GetClient("testUser", 0)
	client.Connect()
	sub, err := uc.centrifugo.GetSubscription(chatId.String(), "testUser", 0)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()
	sub.Subscribe()

	json_data := []byte(`{"value":"` + content.Text + `"}`)
	_, err = sub.Publish(ctx, json_data)
	if err != nil {
		return nil, err
	}
	message := model.Message{
		ChatId:  chatId,
		Content: content,
	}
	return &message, nil
}

func NewMessageUseCase(messageRepo postgres2.Message, centrifugo centrifugo.Centrifugo) MessageUseCase {
	return MessageUseCase{message: messageRepo, centrifugo: centrifugo}
}
