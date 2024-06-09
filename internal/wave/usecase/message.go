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
) ([]*entity.Message, error) {
	messages, err := uc.message.GetMessagesFrom(ctx, chatId.String(), page, pageSize, sort, order)
	return messages, err
}

func (uc MessageUseCase) publish(ctx context.Context, userId uuid.UUID, event string, message *model.Message) error {
	client := uc.centrifugo.GetClient(userId.String(), 0)

	jsonBytes, err := json.Marshal(model.CentrifugoEvent{
		Event:    event,
		ChatId:   message.Chat.Id,
		ChatType: message.Chat.Type,
		Message:  message,
	})

	if err != nil {
		return err
	}
	_, err = client.Publish(ctx, message.Chat.Id.String(), jsonBytes)
	if err != nil {
		return err
	}
	_, err = client.Publish(ctx, userId.String(), jsonBytes)
	return err
}

func (uc MessageUseCase) SendMessage(ctx context.Context, message *entity.Message, chatId uuid.UUID) (*model.Message, error) {
	err := uc.content.Create(ctx, message.Content)
	if err != nil {
		return nil, err
	}
	err = uc.message.Create(ctx, message)
	if err != nil {
		return nil, err
	}
	uc.message.LinkToChat(ctx, message.Id, chatId)

	messageModel := model.NewMessage(message)
	messageModel.Chat = model.NewChat(uc.message.GetChatFromMessage(ctx, message.Id))

	err = uc.publish(ctx, message.SenderId, "send", messageModel)

	return messageModel, err
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
	message.Chat = model.NewChat(chatObject)

	err = uc.publish(ctx, userId, "edit", message)

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
	message.Chat = model.NewChat(chatObject)

	err = uc.publish(ctx, userId, "remove", message)

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
	message.Chat = model.NewChat(chatObject)
	message.Pinned = status

	err = uc.publish(ctx, userId, "pin", message)

	return message, err
}

func (uc MessageUseCase) GetPinnedMessages(ctx context.Context, chatId uuid.UUID) ([]*model.Message, error) {
	messageObjects, err := uc.message.GetPinnedMessages(ctx, chatId)
	messages := model.NewMessages(messageObjects)
	for _, m := range messages {
		m.Pinned = true
	}
	return messages, err
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
