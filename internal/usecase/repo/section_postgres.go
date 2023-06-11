package repo

import (
	"context"
	"github.com/google/uuid"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type SectionRepo struct {
	pg *db.Query
}

func NewSectionRepo(pg *db.Query) SectionRepo {
	return SectionRepo{pg: pg}
}

func (r SectionRepo) CreateSection(ctx context.Context, post entity.SectionPost) (*entity.Section, error) {
	section := entity.NewSectionPostToSection(post)
	err := r.pg.Section.
		WithContext(ctx).
		Preload(r.pg.Section.TextSection).
		Preload(r.pg.Section.ChoiceSection).
		Preload(r.pg.Section.MultiChoiceSection).
		Preload(r.pg.Section.ShortAnswerSection).
		Create(&section)

	if err != nil {
		return &entity.Section{}, nil
	}

	return &section, nil
}

func (r SectionRepo) GetSectionById(ctx context.Context, id uuid.UUID) (*entity.Section, error) {
	return r.pg.Section.
		WithContext(ctx).
		Preload(r.pg.Section.TextSection).
		Preload(r.pg.Section.ChoiceSection).
		Preload(r.pg.Section.MultiChoiceSection).
		Preload(r.pg.Section.ShortAnswerSection).
		Where(r.pg.Section.ID.Eq(id)).
		First()
}

// func (r SectionRepo) GetQuestionById(ctx context.Context, id uint) (entity.SectionQuestion, error) {
// 	var question entity.SectionQuestion
// 	sql, args, err := r.pg.Builder.Select("id", "cost", "case_sensitive",
// 		"slide_id", "order_by", "question", "answer").
// 		From("sections_question").
// 		Where(sq.Eq{"id": id}).
// 		ToSql()

// 	if err != nil {
// 		return question, err
// 	}
// 	rows, err := r.pg.Pool.Query(ctx, sql, args...)
// 	if err != nil {
// 		return entity.SectionQuestion{}, err
// 	}
// 	defer rows.Close()
// 	if !rows.Next() {
// 		return entity.SectionQuestion{}, errors.New("секция не найдена")
// 	}
// 	err = rows.Scan(&question.Id, &question.Cost, &question.CaseSensitive,
// 		&question.SlideId, &question.OrderBy, &question.Question, &question.Answer)
// 	return question, err
// }

// func (r SectionRepo) CreateQuestionAnswer(ctx context.Context, post entity.SectionQuestionAnswer) (uint, error) {
// 	var id uint
// 	sql, args, err := r.pg.Builder.Insert("sections_question_answers").
// 		Columns("answer", "verdict", "section_id", "user_id").
// 		Values(post.Answer, post.Verdict, post.SectionId, post.UserId).
// 		Suffix("returning id").ToSql()

// 	if err != nil {
// 		return id, err
// 	}

// 	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
// 	err = rows.Scan(&id)
// 	return id, err
// }

// func (r SectionRepo) CreateChoice(ctx context.Context, section entity.SectionChoice) (uint, error) {
// 	var id uint
// 	sql, args, err := r.pg.Builder.Insert("sections_choice").
// 		Columns("slide_id", "order_by", "answer", "question", "variants", "cost").
// 		Values(section.SlideId, section.OrderBy, section.Answer, section.Question, pq.Array(section.Variants), section.Cost).
// 		Suffix("returning id").ToSql()
// 	if err != nil {
// 		return id, err
// 	}

// 	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
// 	err = rows.Scan(&id)
// 	return id, err
// }

// func (r SectionRepo) CreateChoiceAnswer(ctx context.Context, answer entity.SectionChoiceAnswer) (uint, error) {
// 	var id uint
// 	sql, args, err := r.pg.Builder.Insert("sections_choice_answers").
// 		Columns("answer", "verdict", "section_id", "user_id").
// 		Values(answer.Answer, answer.Verdict, answer.SectionId, answer.UserId).
// 		Suffix("returning id").ToSql()

// 	if err != nil {
// 		return id, err
// 	}

// 	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
// 	err = rows.Scan(&id)
// 	return id, err
// }

// func (r SectionRepo) CreateMultiChoice(ctx context.Context, section entity.SectionMultiChoice) (uint, error) {
// 	var id uint
// 	sql, args, err := r.pg.Builder.Insert("sections_multi_choice").
// 		Columns("slide_id", "order_by", "answer", "question", "variants", "cost").
// 		Values(section.SlideId, section.OrderBy, pq.Array(section.Answer), section.Question, pq.Array(section.Variants), section.Cost).
// 		Suffix("returning id").ToSql()
// 	if err != nil {
// 		return id, err
// 	}

// 	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
// 	err = rows.Scan(&id)
// 	return id, err
// }

// func (r SectionRepo) GetMultiChoiceById(ctx context.Context, id uint) (entity.SectionMultiChoice, error) {
// 	var choice entity.SectionMultiChoice
// 	sql, args, err := r.pg.Builder.Select("id", "cost", "order_by", "variants",
// 		"question", "answer", "slide_id").
// 		From("sections_multi_choice").
// 		Where(sq.Eq{"id": id}).
// 		ToSql()
// 	if err != nil {
// 		return entity.SectionMultiChoice{}, err
// 	}
// 	rows, err := r.pg.Pool.Query(ctx, sql, args...)
// 	if err != nil {
// 		return entity.SectionMultiChoice{}, err
// 	}
// 	defer rows.Close()
// 	if !rows.Next() {
// 		return entity.SectionMultiChoice{}, errors.New("секция не найдена")
// 	}
// 	err = rows.Scan(&choice.Id, &choice.Cost, &choice.OrderBy, &choice.Variants,
// 		&choice.Question, &choice.Answer, &choice.SlideId)
// 	return choice, err
// }

// func (r SectionRepo) CreateMultiChoiceAnswer(ctx context.Context, answer entity.SectionMultiChoiceAnswer) (uint, error) {
// 	var id uint
// 	sql, args, err := r.pg.Builder.Insert("sections_multi_choice_answers").
// 		Columns("answer", "verdict", "section_id", "user_id").
// 		Values(pq.Array(answer.Answer), answer.Verdict, answer.SectionId, answer.UserId).
// 		Suffix("returning id").ToSql()

// 	if err != nil {
// 		return id, err
// 	}

// 	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
// 	err = rows.Scan(&id)
// 	return id, err
// }
