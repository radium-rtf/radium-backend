package usecase

import (
	"context"
	"encoding/json"
	"fmt"

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

func (uc MessageUseCase) GetMessagesFrom(
	ctx context.Context, chatId uuid.UUID, page, pageSize int, sort, order string,
) ([]*model.Message, error) {
	messageObjects, err := uc.message.GetMessagesFrom(ctx, chatId.String(), page, pageSize, sort, order)
	messages := model.NewMessages(messageObjects)
	for _, m := range messages {
		isPinned, _ := uc.message.IsPinned(ctx, m.Id) // TODO: make it more efficient
		m.SetPinned(isPinned)
	}
	return messages, err
}

func (uc MessageUseCase) GetLastMessage(ctx context.Context, chatId uuid.UUID) (*model.Message, error) {
	messages, err := uc.GetMessagesFrom(ctx, chatId, 1, 1, "date", "desc")
	if len(messages) == 0 {
		return nil, err
	}
	return messages[0], err
}

func (uc MessageUseCase) publish(ctx context.Context, userId uuid.UUID, chatId uuid.UUID, event string, message *model.Message) error {
	client := uc.centrifugo.GetClient(userId.String(), 0)

	jsonBytes, err := json.Marshal(model.CentrifugoEvent{
		Event:   event,
		Message: message,
	})
	if err != nil {
		return err
	}
	_, err = client.Publish(ctx, chatId.String(), jsonBytes)
	if err != nil {
		return err
	}
	return nil
}

func (uc MessageUseCase) SendMessage(ctx context.Context, userId, chatId uuid.UUID, content model.Content) (*model.Message, error) {
	contentObject := &entity.Content{
		DBModel: entity.DBModel{
			Id: uuid.New(),
		},
		Text: content.Text,
	}
	err := uc.content.Create(ctx, contentObject)
	if err != nil {
		return nil, err
	}
	messageObject := &entity.Message{
		DBModel: entity.DBModel{
			Id: uuid.New(),
		},
		SenderId:  userId,
		ContentId: contentObject.Id,
		Content:   contentObject,
	}
	err = uc.message.Create(ctx, messageObject)
	if err != nil {
		return nil, err
	}
	uc.message.LinkToChat(ctx, messageObject.Id, chatId)

	message := model.NewMessage(messageObject)
	message.SetChat(*model.NewChat(uc.message.GetChatFromMessage(ctx, messageObject.Id), nil))

	err = uc.publish(ctx, userId, chatId, "send", message)

	return message, err
}

func (uc MessageUseCase) EditMessage(ctx context.Context, userId, messageId uuid.UUID, content model.Content) (*model.Message, error) {
	messageObject, chatObject, err := uc.message.Get(ctx, messageId)
	if err != nil {
		return nil, err
	}
	if messageObject.SenderId != userId {
		err = fmt.Errorf("unauthorized")
		return nil, err
	}
	contentObject := messageObject.Content
	if content.Text != "" {
		contentObject.Text = content.Text
	}
	err = uc.content.Update(ctx, contentObject)
	if err != nil {
		return nil, err
	}
	message := model.NewMessage(messageObject)
	message.SetChat(*model.NewChat(chatObject, nil))
	pinned, _ := uc.message.IsPinned(ctx, messageId)
	message.SetPinned(pinned)

	err = uc.publish(ctx, userId, message.Chat.Id, "edit", message)

	return message, err
}

func (uc MessageUseCase) RemoveMessage(ctx context.Context, userId, messageId uuid.UUID) (*model.Message, error) {
	messageObject, chatObject, err := uc.message.Get(ctx, messageId)
	if err != nil {
		return nil, err
	}
	if messageObject.SenderId != userId {
		err = fmt.Errorf("unauthorized")
		return nil, err
	}
	err = uc.content.Delete(ctx, messageObject.ContentId)
	if err != nil {
		return nil, err
	}
	err = uc.message.Delete(ctx, messageId)
	if err != nil {
		return nil, err
	}
	message := model.NewMessage(messageObject)
	message.SetChat(*model.NewChat(chatObject, nil))

	err = uc.publish(ctx, userId, message.Chat.Id, "remove", message)

	return message, err
}

func (uc MessageUseCase) PinMessage(ctx context.Context, userId, messageId uuid.UUID, status bool) (*model.Message, error) {
	messageObject, chatObject, err := uc.message.Get(ctx, messageId)
	if err != nil {
		return nil, err
	}
	if messageObject.SenderId != userId {
		err = fmt.Errorf("unauthorized")
		return nil, err
	}
	err = uc.message.Pin(ctx, messageId, chatObject.Id, chatObject.Type, status)
	if err != nil {
		return nil, err
	}
	message := model.NewMessage(messageObject)
	message.SetChat(*model.NewChat(chatObject, nil))
	message.SetPinned(status)

	err = uc.publish(ctx, userId, message.Chat.Id, "pin", message)

	return message, err
}

func NewMessageUseCase(
	messageRepo postgres2.Message,
	contentRepo postgres2.Content,
	centrifugo centrifugo.Centrifugo,
) MessageUseCase {
	return MessageUseCase{
		message:    messageRepo,
		content:    contentRepo,
		centrifugo: centrifugo,
	}
}
