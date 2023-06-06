package repo

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type SlideRepo struct {
	pg *db.Query
}

func NewSlideRepo(pg *db.Query) SlideRepo {
	return SlideRepo{pg: pg}
}

func (r SlideRepo) Create(ctx context.Context, page entity.Page) (*entity.Page, error) {
	err := r.pg.Page.WithContext(ctx).Create(&page)
	if err != nil {
		return &entity.Page{}, err
	}
	return &page, err
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
