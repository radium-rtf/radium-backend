package usecase

import (
	postgres2 "github.com/radium-rtf/radium-backend/internal/wave/usecase/repo/postgres"
)

type PostUseCase struct {
	post postgres2.Post
}

func NewPostUseCase(postRepo postgres2.Post) PostUseCase {
	return PostUseCase{post: postRepo}
}
