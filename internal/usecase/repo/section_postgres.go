package repo

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type SectionRepo struct {
	pg *postgres.Postgres
}

func NewSectionRepo(pg *postgres.Postgres) SectionRepo {
	return SectionRepo{pg: pg}
}

func (r SectionRepo) CreateText(ctx context.Context, post entity.SectionTextPost) (uint, error) {
	var sectionId uint
	sql, args, err := r.pg.Builder.Insert("sections_text").
		Columns("order_by", "markdown", "slide_id").
		Values(post.OrderBy, post.Markdown, post.SlideId).
		Suffix("returning id").ToSql()
	if err != nil {
		return sectionId, err
	}
	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
	err = rows.Scan(&sectionId)
	return sectionId, err
}

func (r SectionRepo) CreateQuestion(ctx context.Context, post entity.SectionQuestionPost) (uint, error) {
	var id uint
	sql, args, err := r.pg.Builder.Insert("sections_question").
		Columns("question", "case_sensitive", "cost", "slide_id", "answer", "order_by").
		Values(post.Question, post.CaseSensitive, post.Cost, post.SlideId, post.Answer, post.OrderBy).
		Suffix("returning id").ToSql()

	if err != nil {
		return id, err
	}

	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
	err = rows.Scan(&id)
	return id, err
}

func (r SectionRepo) GetQuestionById(ctx context.Context, id uint) (entity.SectionQuestion, error) {
	var question entity.SectionQuestion
	sql, args, err := r.pg.Builder.Select("id", "cost", "case_sensitive",
		"slide_id", "order_by", "question", "answer").
		From("sections_question").
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return question, err
	}
	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.SectionQuestion{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return entity.SectionQuestion{}, errors.New("секция не найдена")
	}
	err = rows.Scan(&question.Id, &question.Cost, &question.CaseSensitive,
		&question.SlideId, &question.OrderBy, &question.Question, &question.Answer)
	return question, err
}

func (r SectionRepo) CreateQuestionAnswer(ctx context.Context, post entity.SectionQuestionAnswer) (uint, error) {
	var id uint
	sql, args, err := r.pg.Builder.Insert("sections_question_answers").
		Columns("answer", "verdict", "section_id", "user_id").
		Values(post.Answer, post.Verdict, post.SectionId, post.UserId).
		Suffix("returning id").ToSql()

	if err != nil {
		return id, err
	}

	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
	err = rows.Scan(&id)
	return id, err
}

func (r SectionRepo) CreateChoice(ctx context.Context, section entity.SectionChoice) (uint, error) {
	var id uint
	sql, args, err := r.pg.Builder.Insert("sections_choice").
		Columns("slide_id", "order_by", "answer", "question", "variants", "cost").
		Values(section.SlideId, section.OrderBy, section.Answer, section.Question, pq.Array(section.Variants), section.Cost).
		Suffix("returning id").ToSql()
	if err != nil {
		return id, err
	}

	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
	err = rows.Scan(&id)
	return id, err
}

func (r SectionRepo) GetChoiceById(ctx context.Context, id uint) (entity.SectionChoice, error) {
	var choice entity.SectionChoice
	sql, args, err := r.pg.Builder.Select("id", "cost", "order_by", "variants",
		"question", "answer", "slide_id").
		From("sections_choice").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return entity.SectionChoice{}, err
	}
	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return entity.SectionChoice{}, err
	}
	defer rows.Close()
	if !rows.Next() {
		return entity.SectionChoice{}, errors.New("секция не найдена")
	}
	err = rows.Scan(&choice.Id, &choice.Cost, &choice.OrderBy, &choice.Variants,
		&choice.Question, &choice.Answer, &choice.SlideId)
	return choice, err
}

func (r SectionRepo) CreateChoiceAnswer(ctx context.Context, answer entity.SectionChoiceAnswer) (uint, error) {
	var id uint
	sql, args, err := r.pg.Builder.Insert("sections_choice_answers").
		Columns("answer", "verdict", "section_id", "user_id").
		Values(answer.Answer, answer.Verdict, answer.SectionId, answer.UserId).
		Suffix("returning id").ToSql()

	if err != nil {
		return id, err
	}

	rows := r.pg.Pool.QueryRow(ctx, sql, args...)
	err = rows.Scan(&id)
	return id, err
}
