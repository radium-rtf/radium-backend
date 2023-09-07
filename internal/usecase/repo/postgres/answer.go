package postgres

import (
	"context"
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
	values := uuids(sectionsIds).toValuers()

	q := r.pg.Answer
	// TODO: должен быть запрос, который достает последние ответы, потом доку горма еще раз почитаю
	answers, err := q.WithContext(ctx).
		Where(q.UserId.Eq(userId), q.SectionId.In(values...)).
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

func (r Answer) GetByUsers(ctx context.Context, usersIds, sectionsIds []uuid.UUID) (*entity.UsersAnswersCollection, error) {
	sectionsV := uuids(sectionsIds).toValuers()
	usersV := uuids(usersIds).toValuers()

	q := r.pg.Answer

	// TODO: должен быть запрос, который достает последние ответы, потом доку горма еще раз почитаю
	answers, err := q.WithContext(ctx).
		Where(q.UserId.In(usersV...), q.SectionId.In(sectionsV...)).
		Preload(field.Associations, q.Review).
		Find()

	if err != nil {
		return nil, err
	}

	users, err := r.pg.User.WithContext(ctx).Where(r.pg.User.Id.In(usersV...)).Find()
	if err != nil {
		return nil, err
	}

	answersMap := make(map[uuid.UUID]*entity.Answer)
	for _, answer := range answers {
		prev, ok := answersMap[answer.SectionId]
		if !ok || prev.CreatedAt.Before(answer.CreatedAt) {
			answersMap[answer.SectionId] = answer
		}
	}

	answers = make([]*entity.Answer, 0, len(answersMap))
	for _, answer := range answersMap {
		answers = append(answers, answer)
	}

	return entity.NewUsersAnswersCollection(users, answers), nil
}
