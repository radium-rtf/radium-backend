package mapper

import "github.com/radium-rtf/radium-backend/internal/entity"

type Review struct {
}

func (r Review) PostToReview(post entity.AnswerReviewPost) *entity.AnswerReview {
	return &entity.AnswerReview{
		OwnerId: post.Id,
		Score:   post.Score,
	}
}
