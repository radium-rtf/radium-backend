package usecase

import (
	"context"
	"strings"

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
	messages, err := uc.message.GetMessagesFrom(ctx, chatId.String())
	return messages, err
}

func (uc MessageUseCase) SendMessage(ctx context.Context, chatId uuid.UUID, content model.Content) (*model.Message, error) {
	// test user: 65ff1149-f306-4d35-8b7b-a58c2781d4be
	user := "65ff1149-f306-4d35-8b7b-a58c2781d4be"
	uc.centrifugo.GetClient(user, 0)
	var err error
	sub, err := uc.centrifugo.GetSubscription(chatId.String(), user, 0)
	if err != nil {
		return nil, err
	}
	sub.Subscribe()
	// defer sub.Unsubscribe()

	text := strings.ReplaceAll(content.Text, `"`, `\"`)

	json_data := []byte(`{"value":"` + text + `", "userId": "` + user + `", "chatId": "` + chatId.String() + `"}`)
	_, err = sub.Publish(ctx, json_data)
	if err != nil {
		return nil, err
	}
	message := model.Message{
		ChatId:  chatId,
		Content: content,
	}
	// err = uc.message.Create(ctx, &entity.Message{
	// 	ChatId: chatId,
	// 	Content: &entity.Content{
	// 		Text: content.Text,
	// 	},
	// })
	// if err != nil {
	// 	return nil, err
	// }
	return &message, nil
}

func NewMessageUseCase(messageRepo postgres2.Message, centrifugo centrifugo.Centrifugo) MessageUseCase {
	return MessageUseCase{message: messageRepo, centrifugo: centrifugo}
}
