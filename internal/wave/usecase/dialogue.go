package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	"github.com/radium-rtf/radium-backend/internal/wave/lib/centrifugo"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type DialogueUseCase struct {
	dialogue   postgres2.Dialogue
	centrifugo centrifugo.Centrifugo
}

func (uc DialogueUseCase) GetDialogue(ctx context.Context, chatId uuid.UUID) (*entity.Dialogue, error) {
	dialogue, err := uc.dialogue.Get(ctx)
	return dialogue, err
}

func (uc DialogueUseCase) GetDialogueToken(ctx context.Context, chatId uuid.UUID) (string, error) {
	chatToken, err := uc.centrifugo.GetSubscriptionToken(chatId.String(), "65ff1149-f306-4d35-8b7b-a58c2781d4be", 0)
	if err != nil {
		return "", err
	}
	return chatToken, nil
}

func (uc DialogueUseCase) CreateDialogue(ctx context.Context, userId uuid.UUID, recipientId uuid.UUID) (*entity.Dialogue, error) {
	dialogue, err := uc.dialogue.Get(ctx)
	return dialogue, err
}

func NewDialogueUseCase(dialogueRepo postgres2.Dialogue, centrifugo centrifugo.Centrifugo) DialogueUseCase {
	return DialogueUseCase{dialogue: dialogueRepo, centrifugo: centrifugo}
}
