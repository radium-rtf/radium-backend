package repo

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type ModuleRepo struct {
	pg *db.Query
}

func NewModuleRepo(pg *db.Query) ModuleRepo {
	return ModuleRepo{pg: pg}
}

func (r ModuleRepo) Create(ctx context.Context, module entity.Module) error {
	err := r.pg.Module.WithContext(ctx).Create(&module)
	return err
}

// func (r ModuleRepo) GetModules(ctx context.Context, id int) (entity.CourseModules, error) {
// 	// var modules entity.CourseModules
// 	// sql, args, err := r.pg.Builder.
// 	// 	Select("row_to_json(row)").From("courses_with_modules_view as row").
// 	// 	Where(sq.Eq{"id": id}).ToSql()
// 	// if err != nil {
// 	// 	return modules, err
// 	// }
// 	// rows, err := r.pg.Pool.Query(ctx, sql, args...)
// 	// if err != nil {
// 	// 	return modules, err
// 	// }
// 	// defer rows.Close()
// 	// if !rows.Next() {
// 	// 	return modules, entity.CourseNotFoundErr
// 	// }
// 	// modulesJson := ""
// 	// err = rows.Scan(&modulesJson)
// 	// if err != nil {
// 	// 	return modules, err
// 	// }
// 	// return modules, json.NewDecoder(strings.NewReader(modulesJson)).Decode(&modules)
// }

// func (r ModuleRepo) GetModuleId(ctx context.Context, courseId uint, nameEng string) (uint, error) {
// 	var moduleId uint
// 	sql, args, err := r.pg.Builder.
// 		Select("id").From("modules").
// 		Where(sq.And{sq.Eq{"course_id": courseId}, sq.Eq{"name_eng": nameEng}}).Limit(1).
// 		ToSql()
// 	if err != nil {
// 		return moduleId, err
// 	}
// 	rows, err := r.pg.Pool.Query(ctx, sql, args...)
// 	if err != nil {
// 		return 0, err
// 	}
// 	defer rows.Close()
// 	if !rows.Next() {
// 		return 0, entity.ModulesNotFoundErr
// 	}
// 	err = rows.Scan(&moduleId)
// 	return moduleId, err
// }

// func (r ModuleRepo) Get(ctx context.Context, moduleId entity.SlidesRequest) (entity.ModuleSlides, error) {
// 	slides := entity.ModuleSlides{}
// 	sql, args, err := r.pg.Builder.
// 		Select("row_to_json(row)").
// 		From("module_slides_view as row").Where(sq.Eq{"module_id": moduleId}).ToSql()
// 	if err != nil {
// 		return slides, err
// 	}
// 	rows, err := r.pg.Pool.Query(ctx, sql, args...)
// 	if err != nil {
// 		return slides, err
// 	}
// 	defer rows.Close()
// 	if !rows.Next() {
// 		return slides, entity.CourseNotFoundErr
// 	}
// 	slidesJson := ""
// 	err = rows.Scan(&slidesJson)
// 	if err != nil {
// 		return slides, err
// 	}
// 	return slides, json.NewDecoder(strings.NewReader(slidesJson)).Decode(&slides)
// }
