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
		Insert("course_modules").
		Columns("id", "course_id", "name").
		Values(module.Id, module.CourseId, module.Name).
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
	defer rows.Close()
	if err != nil {
		return modules, err
	}
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
