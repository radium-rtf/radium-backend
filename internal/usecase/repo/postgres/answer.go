package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Answer struct {
}

func NewAnswerRepo(pg *postgres.Postgres) Answer {
	return Answer{}
}

func (r Answer) Create(ctx context.Context, answer *entity.Answer) error {
	panic("not implemented")
}

func (r Answer) Get(ctx context.Context, userId uuid.UUID, sectionsIds []uuid.UUID) (map[uuid.UUID]*entity.Answer, error) {
	panic("not implemented")
}

func (r Answer) GetByUsers(ctx context.Context, userId []uuid.UUID, sectionsIds []uuid.UUID) (
	*entity.UsersAnswersCollection, error) {
	panic("not implemented")
}
