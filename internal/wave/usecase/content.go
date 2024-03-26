package usecase

import (
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type ContentUseCase struct {
	content postgres2.Content
}

func NewContentUseCase(contentRepo postgres2.Content) ContentUseCase {
	return ContentUseCase{content: contentRepo}
}
