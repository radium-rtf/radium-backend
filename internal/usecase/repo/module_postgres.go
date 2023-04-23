package repo

import (
	"context"
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"strings"
)

type ModuleRepo struct {
	pg *postgres.Postgres
}

func NewModuleRepo(pg *postgres.Postgres) ModuleRepo {
	return ModuleRepo{pg: pg}
}

func (r ModuleRepo) Create(ctx context.Context, module entity.Module) error {
	sql, args, err := r.pg.Builder.
		Insert("modules").
		Columns("name_eng", "course_id", "name").
		Values(module.NameEng, module.CourseId, module.Name).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (r ModuleRepo) GetModules(ctx context.Context, id int) (entity.CourseModules, error) {
	var modules entity.CourseModules
	sql, args, err := r.pg.Builder.
		Select("row_to_json(row)").From("courses_with_modules_view as row").
		Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return modules, err
	}
	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return modules, err
	}
	defer rows.Close()
	if !rows.Next() {
		return modules, entity.CourseNotFoundErr
	}
	modulesJson := ""
	err = rows.Scan(&modulesJson)
	if err != nil {
		return modules, err
	}
	return modules, json.NewDecoder(strings.NewReader(modulesJson)).Decode(&modules)
}

func (r ModuleRepo) GetModuleId(ctx context.Context, courseId uint, nameEng string) (uint, error) {
	var moduleId uint
	sql, args, err := r.pg.Builder.
		Select("id").From("modules").
		Where(sq.And{sq.Eq{"course_id": courseId}, sq.Eq{"name_eng": nameEng}}).Limit(1).
		ToSql()
	if err != nil {
		return moduleId, err
	}
	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	if !rows.Next() {
		return 0, entity.ModulesNotFoundErr
	}
	err = rows.Scan(&moduleId)
	return moduleId, err
}

func (r ModuleRepo) Get(ctx context.Context, moduleId entity.SlidesRequest) (entity.ModuleSlides, error) {
	slides := entity.ModuleSlides{}
	sql, args, err := r.pg.Builder.
		Select("row_to_json(row)").
		From("module_slides_view as row").Where(sq.Eq{"module_id": moduleId}).ToSql()
	if err != nil {
		return slides, err
	}
	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return slides, err
	}
	defer rows.Close()
	if !rows.Next() {
		return slides, entity.CourseNotFoundErr
	}
	slidesJson := ""
	err = rows.Scan(&slidesJson)
	if err != nil {
		return slides, err
	}
	return slides, json.NewDecoder(strings.NewReader(slidesJson)).Decode(&slides)
}
