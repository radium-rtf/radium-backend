package model

import (
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/entity"
)

type (
	Post struct {
		Id uuid.UUID `json:"id"`
	}
)

func NewPost(post *entity.Post) Post {
	return Post{
		Id: post.Id,
	}
}
