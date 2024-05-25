package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/centrifugo"
)

type DialogueUseCase struct {
	dialogue   postgres2.Dialogue
	centrifugo centrifugo.Centrifugo
}

func (uc DialogueUseCase) GetDialogue(ctx context.Context, chatId uuid.UUID) (*entity.Dialogue, error) {
	dialogue, err := uc.dialogue.Get(ctx, chatId)
	return dialogue, err
}

func (uc DialogueUseCase) GetDialogues(ctx context.Context, userId uuid.UUID) ([]*entity.Dialogue, error) {
	dialogues, err := uc.dialogue.GetAllByUserId(ctx, userId)
	return dialogues, err
}

func (uc DialogueUseCase) GetDialogueByUsers(ctx context.Context, firstUser, secondUser uuid.UUID) (*entity.Dialogue, error) {
	dialogue, err := uc.dialogue.GetByUsers(ctx, firstUser, secondUser)
	return dialogue, err
}

func (uc DialogueUseCase) GetDialogueToken(ctx context.Context, chatId uuid.UUID) (string, error) {
	userId, ok := ctx.Value("userId").(uuid.UUID)
	if !ok {
		userId = uuid.Nil
	}
	chatToken, err := uc.centrifugo.GetSubscriptionToken(chatId.String(), userId.String(), 0)
	return chatToken, err
}

func (uc DialogueUseCase) CreateDialogue(ctx context.Context, userId uuid.UUID, recipientId uuid.UUID) (*entity.Dialogue, error) {
	dialogue := &entity.Dialogue{
		Id:           uuid.New(),
		FirstUserId:  userId,
		SecondUserId: recipientId,
	}
	err := uc.dialogue.Create(ctx, dialogue)
	if err != nil {
		dialogue, _ = uc.dialogue.GetByUsers(ctx, userId, recipientId)
	}
	return dialogue, err
}

func NewDialogueUseCase(dialogueRepo postgres2.Dialogue, centrifugo centrifugo.Centrifugo) DialogueUseCase {
	return DialogueUseCase{dialogue: dialogueRepo, centrifugo: centrifugo}
}
