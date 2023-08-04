// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package db

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

func newTeacherCourse(db *gorm.DB, opts ...gen.DOOption) teacherCourse {
	_teacherCourse := teacherCourse{}

	_teacherCourse.teacherCourseDo.UseDB(db, opts...)
	_teacherCourse.teacherCourseDo.UseModel(&entity.TeacherCourse{})

	tableName := _teacherCourse.teacherCourseDo.TableName()
	_teacherCourse.ALL = field.NewAsterisk(tableName)
	_teacherCourse.TeacherId = field.NewField(tableName, "teacher_id")
	_teacherCourse.CourseId = field.NewField(tableName, "course_id")
	_teacherCourse.GroupId = field.NewField(tableName, "group_id")
	_teacherCourse.Course = teacherCourseBelongsToCourse{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Course", "entity.Course"),
		Links: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Course.Links", "entity.Link"),
		},
		Modules: struct {
			field.RelationField
			Pages struct {
				field.RelationField
				Sections struct {
					field.RelationField
					TextSection struct {
						field.RelationField
					}
					ChoiceSection struct {
						field.RelationField
					}
					MultiChoiceSection struct {
						field.RelationField
					}
					ShortAnswerSection struct {
						field.RelationField
					}
					AnswerSection struct {
						field.RelationField
					}
				}
			}
		}{
			RelationField: field.NewRelation("Course.Modules", "entity.Module"),
			Pages: struct {
				field.RelationField
				Sections struct {
					field.RelationField
					TextSection struct {
						field.RelationField
					}
					ChoiceSection struct {
						field.RelationField
					}
					MultiChoiceSection struct {
						field.RelationField
					}
					ShortAnswerSection struct {
						field.RelationField
					}
					AnswerSection struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("Course.Modules.Pages", "entity.Page"),
				Sections: struct {
					field.RelationField
					TextSection struct {
						field.RelationField
					}
					ChoiceSection struct {
						field.RelationField
					}
					MultiChoiceSection struct {
						field.RelationField
					}
					ShortAnswerSection struct {
						field.RelationField
					}
					AnswerSection struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("Course.Modules.Pages.Sections", "entity.Section"),
					TextSection: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Course.Modules.Pages.Sections.TextSection", "entity.TextSection"),
					},
					ChoiceSection: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Course.Modules.Pages.Sections.ChoiceSection", "entity.ChoiceSection"),
					},
					MultiChoiceSection: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Course.Modules.Pages.Sections.MultiChoiceSection", "entity.MultiChoiceSection"),
					},
					ShortAnswerSection: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Course.Modules.Pages.Sections.ShortAnswerSection", "entity.ShortAnswerSection"),
					},
					AnswerSection: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Course.Modules.Pages.Sections.AnswerSection", "entity.AnswerSection"),
					},
				},
			},
		},
		Authors: struct {
			field.RelationField
			Sessions struct {
				field.RelationField
			}
			Courses struct {
				field.RelationField
			}
			Groups struct {
				field.RelationField
				Courses struct {
					field.RelationField
				}
				Students struct {
					field.RelationField
				}
			}
		}{
			RelationField: field.NewRelation("Course.Authors", "entity.User"),
			Sessions: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Course.Authors.Sessions", "entity.Session"),
			},
			Courses: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Course.Authors.Courses", "entity.Course"),
			},
			Groups: struct {
				field.RelationField
				Courses struct {
					field.RelationField
				}
				Students struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Course.Authors.Groups", "entity.Group"),
				Courses: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Course.Authors.Groups.Courses", "entity.Course"),
				},
				Students: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Course.Authors.Groups.Students", "entity.User"),
				},
			},
		},
		Students: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Course.Students", "entity.User"),
		},
		Groups: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Course.Groups", "entity.Group"),
		},
	}

	_teacherCourse.Group = teacherCourseBelongsToGroup{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Group", "entity.Group"),
	}

	_teacherCourse.fillFieldMap()

	return _teacherCourse
}

type teacherCourse struct {
	teacherCourseDo teacherCourseDo

	ALL       field.Asterisk
	TeacherId field.Field
	CourseId  field.Field
	GroupId   field.Field
	Course    teacherCourseBelongsToCourse

	Group teacherCourseBelongsToGroup

	fieldMap map[string]field.Expr
}

func (t teacherCourse) Table(newTableName string) *teacherCourse {
	t.teacherCourseDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t teacherCourse) As(alias string) *teacherCourse {
	t.teacherCourseDo.DO = *(t.teacherCourseDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *teacherCourse) updateTableName(table string) *teacherCourse {
	t.ALL = field.NewAsterisk(table)
	t.TeacherId = field.NewField(table, "teacher_id")
	t.CourseId = field.NewField(table, "course_id")
	t.GroupId = field.NewField(table, "group_id")

	t.fillFieldMap()

	return t
}

func (t *teacherCourse) WithContext(ctx context.Context) ITeacherCourseDo {
	return t.teacherCourseDo.WithContext(ctx)
}

func (t teacherCourse) TableName() string { return t.teacherCourseDo.TableName() }

func (t teacherCourse) Alias() string { return t.teacherCourseDo.Alias() }

func (t *teacherCourse) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *teacherCourse) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 5)
	t.fieldMap["teacher_id"] = t.TeacherId
	t.fieldMap["course_id"] = t.CourseId
	t.fieldMap["group_id"] = t.GroupId

}

func (t teacherCourse) clone(db *gorm.DB) teacherCourse {
	t.teacherCourseDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t teacherCourse) replaceDB(db *gorm.DB) teacherCourse {
	t.teacherCourseDo.ReplaceDB(db)
	return t
}

type teacherCourseBelongsToCourse struct {
	db *gorm.DB

	field.RelationField

	Links struct {
		field.RelationField
	}
	Modules struct {
		field.RelationField
		Pages struct {
			field.RelationField
			Sections struct {
				field.RelationField
				TextSection struct {
					field.RelationField
				}
				ChoiceSection struct {
					field.RelationField
				}
				MultiChoiceSection struct {
					field.RelationField
				}
				ShortAnswerSection struct {
					field.RelationField
				}
				AnswerSection struct {
					field.RelationField
				}
			}
		}
	}
	Authors struct {
		field.RelationField
		Sessions struct {
			field.RelationField
		}
		Courses struct {
			field.RelationField
		}
		Groups struct {
			field.RelationField
			Courses struct {
				field.RelationField
			}
			Students struct {
				field.RelationField
			}
		}
	}
	Students struct {
		field.RelationField
	}
	Groups struct {
		field.RelationField
	}
}

func (a teacherCourseBelongsToCourse) Where(conds ...field.Expr) *teacherCourseBelongsToCourse {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a teacherCourseBelongsToCourse) WithContext(ctx context.Context) *teacherCourseBelongsToCourse {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a teacherCourseBelongsToCourse) Session(session *gorm.Session) *teacherCourseBelongsToCourse {
	a.db = a.db.Session(session)
	return &a
}

func (a teacherCourseBelongsToCourse) Model(m *entity.TeacherCourse) *teacherCourseBelongsToCourseTx {
	return &teacherCourseBelongsToCourseTx{a.db.Model(m).Association(a.Name())}
}

type teacherCourseBelongsToCourseTx struct{ tx *gorm.Association }

func (a teacherCourseBelongsToCourseTx) Find() (result *entity.Course, err error) {
	return result, a.tx.Find(&result)
}

func (a teacherCourseBelongsToCourseTx) Append(values ...*entity.Course) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a teacherCourseBelongsToCourseTx) Replace(values ...*entity.Course) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a teacherCourseBelongsToCourseTx) Delete(values ...*entity.Course) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a teacherCourseBelongsToCourseTx) Clear() error {
	return a.tx.Clear()
}

func (a teacherCourseBelongsToCourseTx) Count() int64 {
	return a.tx.Count()
}

type teacherCourseBelongsToGroup struct {
	db *gorm.DB

	field.RelationField
}

func (a teacherCourseBelongsToGroup) Where(conds ...field.Expr) *teacherCourseBelongsToGroup {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a teacherCourseBelongsToGroup) WithContext(ctx context.Context) *teacherCourseBelongsToGroup {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a teacherCourseBelongsToGroup) Session(session *gorm.Session) *teacherCourseBelongsToGroup {
	a.db = a.db.Session(session)
	return &a
}

func (a teacherCourseBelongsToGroup) Model(m *entity.TeacherCourse) *teacherCourseBelongsToGroupTx {
	return &teacherCourseBelongsToGroupTx{a.db.Model(m).Association(a.Name())}
}

type teacherCourseBelongsToGroupTx struct{ tx *gorm.Association }

func (a teacherCourseBelongsToGroupTx) Find() (result *entity.Group, err error) {
	return result, a.tx.Find(&result)
}

func (a teacherCourseBelongsToGroupTx) Append(values ...*entity.Group) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a teacherCourseBelongsToGroupTx) Replace(values ...*entity.Group) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a teacherCourseBelongsToGroupTx) Delete(values ...*entity.Group) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a teacherCourseBelongsToGroupTx) Clear() error {
	return a.tx.Clear()
}

func (a teacherCourseBelongsToGroupTx) Count() int64 {
	return a.tx.Count()
}

type teacherCourseDo struct{ gen.DO }

type ITeacherCourseDo interface {
	gen.SubQuery
	Debug() ITeacherCourseDo
	WithContext(ctx context.Context) ITeacherCourseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITeacherCourseDo
	WriteDB() ITeacherCourseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITeacherCourseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITeacherCourseDo
	Not(conds ...gen.Condition) ITeacherCourseDo
	Or(conds ...gen.Condition) ITeacherCourseDo
	Select(conds ...field.Expr) ITeacherCourseDo
	Where(conds ...gen.Condition) ITeacherCourseDo
	Order(conds ...field.Expr) ITeacherCourseDo
	Distinct(cols ...field.Expr) ITeacherCourseDo
	Omit(cols ...field.Expr) ITeacherCourseDo
	Join(table schema.Tabler, on ...field.Expr) ITeacherCourseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITeacherCourseDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITeacherCourseDo
	Group(cols ...field.Expr) ITeacherCourseDo
	Having(conds ...gen.Condition) ITeacherCourseDo
	Limit(limit int) ITeacherCourseDo
	Offset(offset int) ITeacherCourseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITeacherCourseDo
	Unscoped() ITeacherCourseDo
	Create(values ...*entity.TeacherCourse) error
	CreateInBatches(values []*entity.TeacherCourse, batchSize int) error
	Save(values ...*entity.TeacherCourse) error
	First() (*entity.TeacherCourse, error)
	Take() (*entity.TeacherCourse, error)
	Last() (*entity.TeacherCourse, error)
	Find() ([]*entity.TeacherCourse, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.TeacherCourse, err error)
	FindInBatches(result *[]*entity.TeacherCourse, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.TeacherCourse) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITeacherCourseDo
	Assign(attrs ...field.AssignExpr) ITeacherCourseDo
	Joins(fields ...field.RelationField) ITeacherCourseDo
	Preload(fields ...field.RelationField) ITeacherCourseDo
	FirstOrInit() (*entity.TeacherCourse, error)
	FirstOrCreate() (*entity.TeacherCourse, error)
	FindByPage(offset int, limit int) (result []*entity.TeacherCourse, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITeacherCourseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t teacherCourseDo) Debug() ITeacherCourseDo {
	return t.withDO(t.DO.Debug())
}

func (t teacherCourseDo) WithContext(ctx context.Context) ITeacherCourseDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t teacherCourseDo) ReadDB() ITeacherCourseDo {
	return t.Clauses(dbresolver.Read)
}

func (t teacherCourseDo) WriteDB() ITeacherCourseDo {
	return t.Clauses(dbresolver.Write)
}

func (t teacherCourseDo) Session(config *gorm.Session) ITeacherCourseDo {
	return t.withDO(t.DO.Session(config))
}

func (t teacherCourseDo) Clauses(conds ...clause.Expression) ITeacherCourseDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t teacherCourseDo) Returning(value interface{}, columns ...string) ITeacherCourseDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t teacherCourseDo) Not(conds ...gen.Condition) ITeacherCourseDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t teacherCourseDo) Or(conds ...gen.Condition) ITeacherCourseDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t teacherCourseDo) Select(conds ...field.Expr) ITeacherCourseDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t teacherCourseDo) Where(conds ...gen.Condition) ITeacherCourseDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t teacherCourseDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITeacherCourseDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t teacherCourseDo) Order(conds ...field.Expr) ITeacherCourseDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t teacherCourseDo) Distinct(cols ...field.Expr) ITeacherCourseDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t teacherCourseDo) Omit(cols ...field.Expr) ITeacherCourseDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t teacherCourseDo) Join(table schema.Tabler, on ...field.Expr) ITeacherCourseDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t teacherCourseDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITeacherCourseDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t teacherCourseDo) RightJoin(table schema.Tabler, on ...field.Expr) ITeacherCourseDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t teacherCourseDo) Group(cols ...field.Expr) ITeacherCourseDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t teacherCourseDo) Having(conds ...gen.Condition) ITeacherCourseDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t teacherCourseDo) Limit(limit int) ITeacherCourseDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t teacherCourseDo) Offset(offset int) ITeacherCourseDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t teacherCourseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITeacherCourseDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t teacherCourseDo) Unscoped() ITeacherCourseDo {
	return t.withDO(t.DO.Unscoped())
}

func (t teacherCourseDo) Create(values ...*entity.TeacherCourse) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t teacherCourseDo) CreateInBatches(values []*entity.TeacherCourse, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t teacherCourseDo) Save(values ...*entity.TeacherCourse) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t teacherCourseDo) First() (*entity.TeacherCourse, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TeacherCourse), nil
	}
}

func (t teacherCourseDo) Take() (*entity.TeacherCourse, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TeacherCourse), nil
	}
}

func (t teacherCourseDo) Last() (*entity.TeacherCourse, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TeacherCourse), nil
	}
}

func (t teacherCourseDo) Find() ([]*entity.TeacherCourse, error) {
	result, err := t.DO.Find()
	return result.([]*entity.TeacherCourse), err
}

func (t teacherCourseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.TeacherCourse, err error) {
	buf := make([]*entity.TeacherCourse, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t teacherCourseDo) FindInBatches(result *[]*entity.TeacherCourse, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t teacherCourseDo) Attrs(attrs ...field.AssignExpr) ITeacherCourseDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t teacherCourseDo) Assign(attrs ...field.AssignExpr) ITeacherCourseDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t teacherCourseDo) Joins(fields ...field.RelationField) ITeacherCourseDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t teacherCourseDo) Preload(fields ...field.RelationField) ITeacherCourseDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t teacherCourseDo) FirstOrInit() (*entity.TeacherCourse, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TeacherCourse), nil
	}
}

func (t teacherCourseDo) FirstOrCreate() (*entity.TeacherCourse, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.TeacherCourse), nil
	}
}

func (t teacherCourseDo) FindByPage(offset int, limit int) (result []*entity.TeacherCourse, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t teacherCourseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t teacherCourseDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t teacherCourseDo) Delete(models ...*entity.TeacherCourse) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *teacherCourseDo) withDO(do gen.Dao) *teacherCourseDo {
	t.DO = *do.(*gen.DO)
	return t
}