package postgres

import (
	"context"
	"database/sql/driver"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"gorm.io/gen/field"
)

type Answer struct {
	pg *db.Query
}

func NewAnswerRepo(pg *db.Query) Answer {
	return Answer{pg: pg}
}

func (r Answer) Create(ctx context.Context, answer *entity.Answer) error {
	return r.pg.Answer.WithContext(ctx).Create(answer)
}

func (r Answer) Get(ctx context.Context, userId uuid.UUID, sectionsIds []uuid.UUID) (map[uuid.UUID]*entity.Answer, error) {
	values := make([]driver.Valuer, 0, len(sectionsIds))
	for _, id := range sectionsIds {
		values = append(values, id)
	}

	q := r.pg.Answer
	// TODO: хз как написать норм запрос на этом, потом ещё раз попробую.....
	answers, err := q.WithContext(ctx).
		Where(q.UserId.Eq(userId)).
		Where(q.SectionId.In(values...)).
		Preload(field.Associations).
		Preload(q.Review).
		Find()

	if err != nil {
		return nil, err
	}

	result := make(map[uuid.UUID]*entity.Answer)
	for _, answer := range answers {
		prev, ok := result[answer.SectionId]
		if !ok || prev.CreatedAt.Before(answer.CreatedAt) {
			result[answer.SectionId] = answer
		}
	}

	return result, nil
}
