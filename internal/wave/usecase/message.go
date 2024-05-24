package usecase

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/centrifugo"
)

type MessageUseCase struct {
	message    postgres2.Message
	content    postgres2.Content
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
	userId, ok := ctx.Value("userId").(uuid.UUID)
	if !ok {
		userId = uuid.Nil
	}
	uc.centrifugo.GetClient(userId.String(), 0)
	var err error
	sub, err := uc.centrifugo.GetSubscription(chatId.String(), userId.String(), 0)
	if err != nil {
		return nil, err
	}
	sub.Subscribe()
	// defer sub.Unsubscribe()

	text := strings.ReplaceAll(content.Text, `"`, `\"`)

	json_data := []byte(`{"value":"` + text + `", "userId": "` + userId.String() + `", "chatId": "` + chatId.String() + `"}`)
	_, err = sub.Publish(ctx, json_data)
	if err != nil {
		return nil, err
	}
	contentObject := &entity.Content{
		DBModel: entity.DBModel{
			Id: uuid.New(),
		},
		Text: content.Text,
	}
	err = uc.content.Create(ctx, contentObject)
	if err != nil {
		return nil, err
	}
	message := model.Message{
		ChatId:  chatId,
		Content: content,
	}
	err = uc.message.Create(ctx, &entity.Message{
		DBModel: entity.DBModel{
			Id: uuid.New(),
		},
		ChatId:    chatId,
		ContentId: contentObject.Id,
	})
	return &message, err
}

func NewMessageUseCase(messageRepo postgres2.Message, contentRepo postgres2.Content, centrifugo centrifugo.Centrifugo) MessageUseCase {
	return MessageUseCase{message: messageRepo, content: contentRepo, centrifugo: centrifugo}
}
