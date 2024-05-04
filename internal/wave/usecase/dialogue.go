package usecase

import (
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type DialogueUseCase struct {
	dialogue postgres2.Dialogue
}

func NewDialogueUseCase(dialogueRepo postgres2.Dialogue) DialogueUseCase {
	return DialogueUseCase{dialogue: dialogueRepo}
}
