package repo

import (
	"context"
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
