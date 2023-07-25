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
		Preload(s.TextSection, s.ChoiceSection, s.MultiChoiceSection, s.ShortAnswerSection, s.AnswerSection).
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

// TODO: перенос в module
// func (r SlideRepo) Get(ctx context.Context, moduleId uint) (entity.ModuleSlides, error) {
// 	slides := entity.ModuleSlides{}
// 	sql, args, err := r.pg.Builder.
// 		Select("row_to_json(row)").
// 		From("module_slides_view as row").Where(sq.Eq{"id": moduleId}).ToSql()
// 	if err != nil {
// 		return slides, err
// 	}
// 	rows, err := r.pg.Pool.Query(ctx, sql, args...)
// 	if err != nil {
// 		return slides, err
// 	}
// 	defer rows.Close()
// 	if !rows.Next() {
// 		return slides, entity.ModulesNotFoundErr
// 	}
// 	slidesJson := ""
// 	err = rows.Scan(&slidesJson)
// 	if err != nil {
// 		return slides, err
// 	}
// 	return slides, json.NewDecoder(strings.NewReader(slidesJson)).Decode(&slides)
// }

// TODO: полиморфы
// func (r SlideRepo) GetSections(ctx context.Context, slideId uint) (entity.SlideSections, error) {
// 	slideSections := entity.SlideSections{}
// 	sql, args, err := r.pg.Builder.
// 		Select("row_to_json(row)").
// 		From("slide_sections_view as row").Where(sq.Eq{"id": slideId}).ToSql()
// 	if err != nil {
// 		return slideSections, err
// 	}
// 	rows, err := r.pg.Pool.Query(ctx, sql, args...)
// 	if err != nil {
// 		return slideSections, err
// 	}
// 	defer rows.Close()
// 	if !rows.Next() {
// 		return slideSections, entity.ModulesNotFoundErr
// 	}
// 	slideSectionsJson := ""
// 	err = rows.Scan(&slideSectionsJson)
// 	if err != nil {
// 		return slideSections, err
// 	}
// 	return slideSections, json.NewDecoder(strings.NewReader(slideSectionsJson)).Decode(&slideSections)
// }
