package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type DialogueUseCase struct {
	dialogue postgres2.Dialogue
	chat     postgres2.Chat
}

func (uc DialogueUseCase) GetDialogue(ctx context.Context, chatId uuid.UUID) (*entity.Dialogue, error) {
	dialogue, err := uc.dialogue.Get(ctx, chatId)
	return dialogue, err
}

func (uc DialogueUseCase) GetDialogueByUsers(ctx context.Context, firstUser, secondUser uuid.UUID) (*entity.Dialogue, error) {
	dialogue, err := uc.dialogue.GetByUsers(ctx, firstUser, secondUser)
	return dialogue, err
}

func (uc DialogueUseCase) CreateDialogue(ctx context.Context, dialogue *entity.Dialogue) error {
	err := uc.dialogue.Create(ctx, dialogue)
	if err != nil {
		return err
	}
	userId := dialogue.FirstUserId
	recipientId := dialogue.SecondUserId
	chat := &entity.Chat{
		Id:   dialogue.Id,
		Name: userId.String() + " / " + recipientId.String(),
		Type: "dialogue",
	}
	err = uc.chat.Create(ctx, chat)
	return err
}

func NewDialogueUseCase(dialogueRepo postgres2.Dialogue, chatRepo postgres2.Chat) DialogueUseCase {
	return DialogueUseCase{dialogue: dialogueRepo, chat: chatRepo}
}
