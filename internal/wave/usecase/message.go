package usecase

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/model"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/centrifugo"
)

type MessageUseCase struct {
	message    postgres2.Message
	content    postgres2.Content
	dialogue   postgres2.Dialogue
	centrifugo centrifugo.Centrifugo
}

func (uc MessageUseCase) GetMessage(ctx context.Context, messageId uuid.UUID) (*entity.Message, error) {
	message, _, err := uc.message.Get(ctx, messageId)
	return message, err
}

func (uc MessageUseCase) GetMessagesFrom(
	ctx context.Context, chatId uuid.UUID, page, pageSize int, sort, order string,
) ([]*entity.Message, error) {
	messages, err := uc.message.GetMessagesFrom(ctx, chatId.String(), page, pageSize, sort, order)
	return messages, err
}

func (uc MessageUseCase) GetLastMessage(ctx context.Context, chatId uuid.UUID) (*entity.Message, error) {
	messages, err := uc.message.GetMessagesFrom(ctx, chatId.String(), 1, 1, "date", "desc")
	return messages[0], err
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
	uc.dialogue.LinkMessage(ctx, chatId, messageObject.Id)

	message := model.NewMessage(messageObject)
	message.SetChat(model.NewChat(
		chatId,
		chatId.String(), // TODO: change name
		"dialogue",
		nil,
	))

	client := uc.centrifugo.GetClient(userId.String(), 0)

	jsonBytes, err := json.Marshal(model.CentrifugoEvent{
		Event:   "create",
		Message: message,
	})
	if err != nil {
		return nil, err
	}
	_, err = client.Publish(ctx, chatId.String(), jsonBytes)
	if err != nil {
		return nil, err
	}

	return message, err
}

func NewMessageUseCase(
	messageRepo postgres2.Message,
	contentRepo postgres2.Content,
	dialogueRepo postgres2.Dialogue,
	centrifugo centrifugo.Centrifugo,
) MessageUseCase {
	return MessageUseCase{
		message:    messageRepo,
		content:    contentRepo,
		dialogue:   dialogueRepo,
		centrifugo: centrifugo,
	}
}

func (uc MessageUseCase) EditMessage(ctx context.Context, userId, messageId uuid.UUID, content model.Content) (*model.Message, error) {
	messageObject, chatModel, err := uc.message.Get(ctx, messageId)
	if err != nil {
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
	message.SetChat(*chatModel)

	client := uc.centrifugo.GetClient(userId.String(), 0)

	jsonBytes, err := json.Marshal(model.CentrifugoEvent{
		Event:   "edit",
		Message: message,
	})
	if err != nil {
		return nil, err
	}
	_, err = client.Publish(ctx, chatModel.Id.String(), jsonBytes)
	if err != nil {
		return nil, err
	}

	return message, err
}

func (uc MessageUseCase) RemoveMessage(ctx context.Context, userId, messageId uuid.UUID) (*model.Message, error) {
	messageObject, chatModel, err := uc.message.Get(ctx, messageId)
	if err != nil {
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
	message.SetChat(*chatModel)

	client := uc.centrifugo.GetClient(userId.String(), 0)

	jsonBytes, err := json.Marshal(model.CentrifugoEvent{
		Event:   "remove",
		Message: message,
	})
	if err != nil {
		return nil, err
	}
	_, err = client.Publish(ctx, chatModel.Id.String(), jsonBytes)
	if err != nil {
		return nil, err
	}

	return message, err
}
