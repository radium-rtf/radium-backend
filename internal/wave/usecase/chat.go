package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
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

func NewChatUseCase(chatRepo postgres2.Chat, centrifugo centrifugo.Centrifugo) ChatUseCase {
	return ChatUseCase{chat: chatRepo, centrifugo: centrifugo}
}
