package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Review struct {
}

func NewReviewRepo(pg *postgres.Postgres) Review {
	return Review{}
}

func (r Review) Create(ctx context.Context, review *entity.Review) (*entity.Review, error) {
	panic("not implemented")
}
