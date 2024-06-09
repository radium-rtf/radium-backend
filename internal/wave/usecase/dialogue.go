package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type DialogueUseCase struct {
	dialogue postgres2.Dialogue
	chat     ChatUseCase
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
	err = uc.chat.CreateFromDialogue(ctx, dialogue)
	return err
}

func NewDialogueUseCase(dialogueRepo postgres2.Dialogue, chatUC ChatUseCase) DialogueUseCase {
	return DialogueUseCase{dialogue: dialogueRepo, chat: chatUC}
}
