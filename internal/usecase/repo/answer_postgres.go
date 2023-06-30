package repo

import (
	"context"
	"database/sql/driver"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

type AnswerRepo struct {
	pg *db.Query
}

func NewAnswerRepo(pg *db.Query) AnswerRepo {
	return AnswerRepo{pg: pg}
}

func (r AnswerRepo) CreateOrUpdate(ctx context.Context, answer *entity.Answer) error {
	q := r.pg.Answer

	columns := []clause.Column{{Name: q.UserId.ColumnName().String()}, {Name: q.SectionId.ColumnName().String()}}
	return q.WithContext(ctx).
		Preload(field.Associations).
		Clauses(clause.OnConflict{UpdateAll: true, Columns: columns}).Create(answer)
}

func (r AnswerRepo) Get(ctx context.Context, userId uuid.UUID, sectionsIds []uuid.UUID) (map[uuid.UUID]*entity.Answer, error) {
	values := make([]driver.Valuer, 0, len(sectionsIds))
	for _, id := range sectionsIds {
		values = append(values, id)
	}

	q := r.pg.Answer

	answers, err := q.WithContext(ctx).
		Preload(field.Associations).
		Where(q.UserId.Eq(userId)).
		Where(q.SectionId.In(values...)).
		Find()

	if err != nil {
		return nil, err
	}

	result := make(map[uuid.UUID]*entity.Answer)
	for _, answer := range answers {
		result[answer.SectionId] = answer
	}

	return result, nil
}
