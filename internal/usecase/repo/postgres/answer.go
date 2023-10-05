package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Answer struct {
	db    *bun.DB
	users User
}

func NewAnswerRepo(pg *postgres.Postgres) Answer {
	return Answer{db: pg.DB, users: NewUserRepo(pg)}
}

func (r Answer) Create(ctx context.Context, answer *entity.Answer) error {
	_, err := r.db.NewInsert().Model(answer).Exec(ctx)
	return err
}

func (r Answer) Get(ctx context.Context, userId uuid.UUID, sectionsIds []uuid.UUID) (*entity.AnswersCollection, error) {
	answers, err := r.get(ctx, []uuid.UUID{userId}, sectionsIds)
	return entity.NewAnswersCollection(answers), err
}

func (r Answer) get(ctx context.Context, usersIds []uuid.UUID, sectionsIds []uuid.UUID) ([]*entity.Answer, error) {
	var answers []*entity.Answer

	subq := r.db.NewSelect().
		TableExpr("answers as a").
		ColumnExpr("max(a.created_at)").
		Where("a.user_id = answer.user_id and a.section_id = answer.section_id")
	err := r.db.NewSelect().
		Model(&answers).
		Where("user_id = (?) and section_id in (?) and answer.created_at = (?)", bun.In(usersIds), bun.In(sectionsIds), subq).
		Relation("Review").
		Scan(ctx)

	return answers, err
}

func (r Answer) GetByUsers(ctx context.Context, usersIds []uuid.UUID, sectionsIds []uuid.UUID) (
	*entity.UsersAnswersCollection, error) {
	answers, err := r.get(ctx, usersIds, sectionsIds)
	if err != nil {
		return nil, err
	}
	users, err := r.users.GetByIds(ctx, usersIds)
	if err != nil {
		return nil, err
	}

	return entity.NewUsersAnswersCollection(users, answers), nil
}
