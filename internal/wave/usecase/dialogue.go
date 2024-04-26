package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type DialogueUseCase struct {
	dialogue postgres2.Dialogue
}

func (uc DialogueUseCase) GetDialogue(ctx context.Context) (*entity.Dialogue, error) {
	dialogue, err := uc.dialogue.Get(ctx)
	return dialogue, err
}

func (uc DialogueUseCase) CreateDialogue(ctx context.Context, userId uuid.UUID, recipientId uuid.UUID) (*entity.Dialogue, error) {
	dialogue, err := uc.dialogue.Get(ctx)
	return dialogue, err
}

func NewDialogueUseCase(dialogueRepo postgres2.Dialogue) DialogueUseCase {
	return DialogueUseCase{dialogue: dialogueRepo}
}
