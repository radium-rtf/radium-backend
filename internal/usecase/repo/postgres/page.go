package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type Page struct {
	pg *db.Query
}

func NewPageRepo(pg *db.Query) Page {
	return Page{pg: pg}
}

func (r Page) Create(ctx context.Context, page *entity.Page) (*entity.Page, error) {
	err := r.pg.Page.WithContext(ctx).Create(page)
	if err != nil {
		return &entity.Page{}, err
	}
	return page, err
}

func (r Page) GetByID(ctx context.Context, id uuid.UUID) (*entity.Page, error) {
	p := r.pg.Page
	s := p.Sections
	return p.WithContext(ctx).Debug().
		Order(p.Order).
		Preload(s, s.Order(r.pg.Section.Order)).
		Preload(s.TextSection, s.ChoiceSection, s.MultiChoiceSection, s.ShortAnswerSection, s.AnswerSection, s.CodeSection).
		Where(p.Id.Eq(id)).
		Take()
}

func (r Page) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	p := r.pg.Page.WithContext(ctx)
	if !isSoft {
		p = p.Unscoped()
	}
	info, err := p.Where(r.pg.Page.Id.Eq(id)).Delete()
	if err == nil && info.RowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}
