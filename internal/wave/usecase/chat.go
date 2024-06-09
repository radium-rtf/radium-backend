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

type ChatUseCase struct {
	chat       postgres2.Chat
	centrifugo centrifugo.Centrifugo
}

func (uc ChatUseCase) GetChat(ctx context.Context, chatId uuid.UUID) (*entity.Chat, error) {
	chat, err := uc.chat.Get(ctx, chatId)
	return chat, err
}

func (uc ChatUseCase) GetChatByMessageId(ctx context.Context, messageId uuid.UUID) (*entity.Chat, error) {
	chat, err := uc.chat.GetByMessageId(ctx, messageId)
	return chat, err
}

func (uc ChatUseCase) GetChats(ctx context.Context, userId uuid.UUID) ([]*entity.Chat, error) {
	chats, err := uc.chat.GetAllByUserId(ctx, userId)
	return chats, err
}

func (uc ChatUseCase) GetChatToken(ctx context.Context, chatId, userId uuid.UUID) (string, error) {
	chatToken, err := uc.centrifugo.GetSubscriptionToken(chatId.String(), userId.String(), 0)
	return chatToken, err
}

func (uc ChatUseCase) CreateChat(ctx context.Context, Name string, Type string) (*entity.Chat, error) {
	chat := &entity.Chat{
		Id:   uuid.New(),
		Name: Name,
		Type: Type,
	}
	err := uc.chat.Create(ctx, chat)
	return chat, err
}

func (uc ChatUseCase) publish(ctx context.Context, event string, dialogue *model.Dialogue, group *model.Group) (*[]byte, error) {
	client := uc.centrifugo.GetClient(uuid.Nil.String(), 0)

	var chatId uuid.UUID
	var chatType string
	if dialogue != nil {
		chatId = dialogue.Id
		chatType = "dialogue"
	} else if group != nil {
		chatId = group.Id
		chatType = "group"
	} else {
		return nil, fmt.Errorf("dialogue and group are nil")
	}

	jsonBytes, err := json.Marshal(model.CentrifugoEvent{
		Event:    event + "-chat",
		ChatId:   chatId,
		ChatType: chatType,
		Dialogue: dialogue,
		Group:    group,
	})
	if err != nil {
		return nil, err
	}

	_, err = client.Publish(ctx, chatId.String(), jsonBytes)
	if err != nil {
		return nil, err
	}
	return &jsonBytes, err
}

func (uc ChatUseCase) publishDialogue(ctx context.Context, event string, dialogue *model.Dialogue) error {
	data, err := uc.publish(ctx, event, dialogue, nil)
	if err != nil {
		return err
	}
	userId := dialogue.FirstUserId.String()
	_, _ = uc.centrifugo.GetClient(userId, 0).Publish(ctx, userId, *data)
	userId = dialogue.SecondUserId.String()
	_, _ = uc.centrifugo.GetClient(userId, 0).Publish(ctx, userId, *data)
	return nil
}

func (uc ChatUseCase) publishGroup(ctx context.Context, event string, group *model.Group) error {
	data, err := uc.publish(ctx, event, nil, group)
	if err != nil {
		return err
	}
	for _, userId := range group.Members {
		_, _ = uc.centrifugo.GetClient(userId.String(), 0).Publish(ctx, userId.String(), *data)
	}
	return nil
}

func (uc ChatUseCase) CreateFromDialogue(ctx context.Context, dialogue *entity.Dialogue) error {
	chat := &entity.Chat{
		Id:   dialogue.Id,
		Name: dialogue.FirstUserId.String() + " / " + dialogue.SecondUserId.String(),
		Type: "dialogue",
	}

	err := uc.chat.Create(ctx, chat)
	if err != nil {
		return err
	}

	Model := model.NewDialogue(dialogue)
	return uc.publishDialogue(ctx, "create", &Model)
}

func (uc ChatUseCase) ModifyFromDialogue(ctx context.Context, dialogue *entity.Dialogue) error {
	chat := &entity.Chat{
		Id:   dialogue.Id,
		Name: dialogue.FirstUserId.String() + " / " + dialogue.SecondUserId.String(),
		Type: "dialogue",
	}

	err := uc.chat.Update(ctx, chat)
	if err != nil {
		return err
	}

	Model := model.NewDialogue(dialogue)
	return uc.publishDialogue(ctx, "modify", &Model)
}

func (uc ChatUseCase) CreateFromGroup(ctx context.Context, group *entity.Group) error {
	chat := &entity.Chat{
		Id:   group.Id,
		Name: group.Name,
		Type: "group",
	}

	err := uc.chat.Create(ctx, chat)
	if err != nil {
		return err
	}

	Model := model.NewGroup(group)
	Model.Members = append(Model.Members, group.OwnerId)
	return uc.publishGroup(ctx, "create", &Model)
}

func (uc ChatUseCase) ModifyFromGroup(ctx context.Context, group *entity.Group) error {
	chat := &entity.Chat{
		Id:   group.Id,
		Name: group.Name,
		Type: "group",
	}

	err := uc.chat.Update(ctx, chat)
	if err != nil {
		return err
	}

	Model := model.NewGroup(group)
	return uc.publishGroup(ctx, "modify", &Model)
}

func NewChatUseCase(chatRepo postgres2.Chat, centrifugo centrifugo.Centrifugo) ChatUseCase {
	return ChatUseCase{chat: chatRepo, centrifugo: centrifugo}
}
