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

func newCourse(db *gorm.DB, opts ...gen.DOOption) course {
	_course := course{}

	_course.courseDo.UseDB(db, opts...)
	_course.courseDo.UseModel(&entity.Course{})

	tableName := _course.courseDo.TableName()
	_course.ALL = field.NewAsterisk(tableName)
	_course.Id = field.NewField(tableName, "id")
	_course.CreatedAt = field.NewTime(tableName, "created_at")
	_course.UpdatedAt = field.NewTime(tableName, "updated_at")
	_course.DeletedAt = field.NewField(tableName, "deleted_at")
	_course.Name = field.NewString(tableName, "name")
	_course.Slug = field.NewString(tableName, "slug")
	_course.ShortDescription = field.NewString(tableName, "short_description")
	_course.Description = field.NewString(tableName, "description")
	_course.Logo = field.NewString(tableName, "logo")
	_course.Banner = field.NewString(tableName, "banner")
	_course.Links = courseHasManyLinks{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Links", "entity.Link"),
	}

	_course.Modules = courseHasManyModules{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Modules", "entity.Module"),
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
			}
		}{
			RelationField: field.NewRelation("Modules.Pages", "entity.Page"),
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
			}{
				RelationField: field.NewRelation("Modules.Pages.Sections", "entity.Section"),
				TextSection: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Modules.Pages.Sections.TextSection", "entity.TextSection"),
				},
				ChoiceSection: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Modules.Pages.Sections.ChoiceSection", "entity.ChoiceSection"),
				},
				MultiChoiceSection: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Modules.Pages.Sections.MultiChoiceSection", "entity.MultiChoiceSection"),
				},
				ShortAnswerSection: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Modules.Pages.Sections.ShortAnswerSection", "entity.ShortAnswerSection"),
				},
			},
		},
	}

	_course.Authors = courseManyToManyAuthors{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Authors", "entity.User"),
		Sessions: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Authors.Sessions", "entity.Session"),
		},
		Courses: struct {
			field.RelationField
			Links struct {
				field.RelationField
			}
			Modules struct {
				field.RelationField
			}
			Authors struct {
				field.RelationField
			}
			Students struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Authors.Courses", "entity.Course"),
			Links: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Authors.Courses.Links", "entity.Link"),
			},
			Modules: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Authors.Courses.Modules", "entity.Module"),
			},
			Authors: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Authors.Courses.Authors", "entity.User"),
			},
			Students: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Authors.Courses.Students", "entity.User"),
			},
		},
	}

	_course.Students = courseManyToManyStudents{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Students", "entity.User"),
	}

	_course.fillFieldMap()

	return _course
}

type course struct {
	courseDo courseDo

	ALL              field.Asterisk
	Id               field.Field
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	Name             field.String
	Slug             field.String
	ShortDescription field.String
	Description      field.String
	Logo             field.String
	Banner           field.String
	Links            courseHasManyLinks

	Modules courseHasManyModules

	Authors courseManyToManyAuthors

	Students courseManyToManyStudents

	fieldMap map[string]field.Expr
}

func (c course) Table(newTableName string) *course {
	c.courseDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c course) As(alias string) *course {
	c.courseDo.DO = *(c.courseDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *course) updateTableName(table string) *course {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewField(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Name = field.NewString(table, "name")
	c.Slug = field.NewString(table, "slug")
	c.ShortDescription = field.NewString(table, "short_description")
	c.Description = field.NewString(table, "description")
	c.Logo = field.NewString(table, "logo")
	c.Banner = field.NewString(table, "banner")

	c.fillFieldMap()

	return c
}

func (c *course) WithContext(ctx context.Context) ICourseDo { return c.courseDo.WithContext(ctx) }

func (c course) TableName() string { return c.courseDo.TableName() }

func (c course) Alias() string { return c.courseDo.Alias() }

func (c *course) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *course) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 14)
	c.fieldMap["id"] = c.Id
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["name"] = c.Name
	c.fieldMap["slug"] = c.Slug
	c.fieldMap["short_description"] = c.ShortDescription
	c.fieldMap["description"] = c.Description
	c.fieldMap["logo"] = c.Logo
	c.fieldMap["banner"] = c.Banner

}

func (c course) clone(db *gorm.DB) course {
	c.courseDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c course) replaceDB(db *gorm.DB) course {
	c.courseDo.ReplaceDB(db)
	return c
}

type courseHasManyLinks struct {
	db *gorm.DB

	field.RelationField
}

func (a courseHasManyLinks) Where(conds ...field.Expr) *courseHasManyLinks {
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

func (a courseHasManyLinks) WithContext(ctx context.Context) *courseHasManyLinks {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a courseHasManyLinks) Session(session *gorm.Session) *courseHasManyLinks {
	a.db = a.db.Session(session)
	return &a
}

func (a courseHasManyLinks) Model(m *entity.Course) *courseHasManyLinksTx {
	return &courseHasManyLinksTx{a.db.Model(m).Association(a.Name())}
}

type courseHasManyLinksTx struct{ tx *gorm.Association }

func (a courseHasManyLinksTx) Find() (result []*entity.Link, err error) {
	return result, a.tx.Find(&result)
}

func (a courseHasManyLinksTx) Append(values ...*entity.Link) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a courseHasManyLinksTx) Replace(values ...*entity.Link) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a courseHasManyLinksTx) Delete(values ...*entity.Link) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a courseHasManyLinksTx) Clear() error {
	return a.tx.Clear()
}

func (a courseHasManyLinksTx) Count() int64 {
	return a.tx.Count()
}

type courseHasManyModules struct {
	db *gorm.DB

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
		}
	}
}

func (a courseHasManyModules) Where(conds ...field.Expr) *courseHasManyModules {
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

func (a courseHasManyModules) WithContext(ctx context.Context) *courseHasManyModules {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a courseHasManyModules) Session(session *gorm.Session) *courseHasManyModules {
	a.db = a.db.Session(session)
	return &a
}

func (a courseHasManyModules) Model(m *entity.Course) *courseHasManyModulesTx {
	return &courseHasManyModulesTx{a.db.Model(m).Association(a.Name())}
}

type courseHasManyModulesTx struct{ tx *gorm.Association }

func (a courseHasManyModulesTx) Find() (result []*entity.Module, err error) {
	return result, a.tx.Find(&result)
}

func (a courseHasManyModulesTx) Append(values ...*entity.Module) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a courseHasManyModulesTx) Replace(values ...*entity.Module) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a courseHasManyModulesTx) Delete(values ...*entity.Module) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a courseHasManyModulesTx) Clear() error {
	return a.tx.Clear()
}

func (a courseHasManyModulesTx) Count() int64 {
	return a.tx.Count()
}

type courseManyToManyAuthors struct {
	db *gorm.DB

	field.RelationField

	Sessions struct {
		field.RelationField
	}
	Courses struct {
		field.RelationField
		Links struct {
			field.RelationField
		}
		Modules struct {
			field.RelationField
		}
		Authors struct {
			field.RelationField
		}
		Students struct {
			field.RelationField
		}
	}
}

func (a courseManyToManyAuthors) Where(conds ...field.Expr) *courseManyToManyAuthors {
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

func (a courseManyToManyAuthors) WithContext(ctx context.Context) *courseManyToManyAuthors {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a courseManyToManyAuthors) Session(session *gorm.Session) *courseManyToManyAuthors {
	a.db = a.db.Session(session)
	return &a
}

func (a courseManyToManyAuthors) Model(m *entity.Course) *courseManyToManyAuthorsTx {
	return &courseManyToManyAuthorsTx{a.db.Model(m).Association(a.Name())}
}

type courseManyToManyAuthorsTx struct{ tx *gorm.Association }

func (a courseManyToManyAuthorsTx) Find() (result []*entity.User, err error) {
	return result, a.tx.Find(&result)
}

func (a courseManyToManyAuthorsTx) Append(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a courseManyToManyAuthorsTx) Replace(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a courseManyToManyAuthorsTx) Delete(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a courseManyToManyAuthorsTx) Clear() error {
	return a.tx.Clear()
}

func (a courseManyToManyAuthorsTx) Count() int64 {
	return a.tx.Count()
}

type courseManyToManyStudents struct {
	db *gorm.DB

	field.RelationField
}

func (a courseManyToManyStudents) Where(conds ...field.Expr) *courseManyToManyStudents {
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

func (a courseManyToManyStudents) WithContext(ctx context.Context) *courseManyToManyStudents {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a courseManyToManyStudents) Session(session *gorm.Session) *courseManyToManyStudents {
	a.db = a.db.Session(session)
	return &a
}

func (a courseManyToManyStudents) Model(m *entity.Course) *courseManyToManyStudentsTx {
	return &courseManyToManyStudentsTx{a.db.Model(m).Association(a.Name())}
}

type courseManyToManyStudentsTx struct{ tx *gorm.Association }

func (a courseManyToManyStudentsTx) Find() (result []*entity.User, err error) {
	return result, a.tx.Find(&result)
}

func (a courseManyToManyStudentsTx) Append(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a courseManyToManyStudentsTx) Replace(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a courseManyToManyStudentsTx) Delete(values ...*entity.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a courseManyToManyStudentsTx) Clear() error {
	return a.tx.Clear()
}

func (a courseManyToManyStudentsTx) Count() int64 {
	return a.tx.Count()
}

type courseDo struct{ gen.DO }

type ICourseDo interface {
	gen.SubQuery
	Debug() ICourseDo
	WithContext(ctx context.Context) ICourseDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICourseDo
	WriteDB() ICourseDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICourseDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICourseDo
	Not(conds ...gen.Condition) ICourseDo
	Or(conds ...gen.Condition) ICourseDo
	Select(conds ...field.Expr) ICourseDo
	Where(conds ...gen.Condition) ICourseDo
	Order(conds ...field.Expr) ICourseDo
	Distinct(cols ...field.Expr) ICourseDo
	Omit(cols ...field.Expr) ICourseDo
	Join(table schema.Tabler, on ...field.Expr) ICourseDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICourseDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICourseDo
	Group(cols ...field.Expr) ICourseDo
	Having(conds ...gen.Condition) ICourseDo
	Limit(limit int) ICourseDo
	Offset(offset int) ICourseDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICourseDo
	Unscoped() ICourseDo
	Create(values ...*entity.Course) error
	CreateInBatches(values []*entity.Course, batchSize int) error
	Save(values ...*entity.Course) error
	First() (*entity.Course, error)
	Take() (*entity.Course, error)
	Last() (*entity.Course, error)
	Find() ([]*entity.Course, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Course, err error)
	FindInBatches(result *[]*entity.Course, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Course) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICourseDo
	Assign(attrs ...field.AssignExpr) ICourseDo
	Joins(fields ...field.RelationField) ICourseDo
	Preload(fields ...field.RelationField) ICourseDo
	FirstOrInit() (*entity.Course, error)
	FirstOrCreate() (*entity.Course, error)
	FindByPage(offset int, limit int) (result []*entity.Course, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICourseDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c courseDo) Debug() ICourseDo {
	return c.withDO(c.DO.Debug())
}

func (c courseDo) WithContext(ctx context.Context) ICourseDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c courseDo) ReadDB() ICourseDo {
	return c.Clauses(dbresolver.Read)
}

func (c courseDo) WriteDB() ICourseDo {
	return c.Clauses(dbresolver.Write)
}

func (c courseDo) Session(config *gorm.Session) ICourseDo {
	return c.withDO(c.DO.Session(config))
}

func (c courseDo) Clauses(conds ...clause.Expression) ICourseDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c courseDo) Returning(value interface{}, columns ...string) ICourseDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c courseDo) Not(conds ...gen.Condition) ICourseDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c courseDo) Or(conds ...gen.Condition) ICourseDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c courseDo) Select(conds ...field.Expr) ICourseDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c courseDo) Where(conds ...gen.Condition) ICourseDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c courseDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICourseDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c courseDo) Order(conds ...field.Expr) ICourseDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c courseDo) Distinct(cols ...field.Expr) ICourseDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c courseDo) Omit(cols ...field.Expr) ICourseDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c courseDo) Join(table schema.Tabler, on ...field.Expr) ICourseDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c courseDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICourseDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c courseDo) RightJoin(table schema.Tabler, on ...field.Expr) ICourseDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c courseDo) Group(cols ...field.Expr) ICourseDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c courseDo) Having(conds ...gen.Condition) ICourseDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c courseDo) Limit(limit int) ICourseDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c courseDo) Offset(offset int) ICourseDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c courseDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICourseDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c courseDo) Unscoped() ICourseDo {
	return c.withDO(c.DO.Unscoped())
}

func (c courseDo) Create(values ...*entity.Course) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c courseDo) CreateInBatches(values []*entity.Course, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c courseDo) Save(values ...*entity.Course) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c courseDo) First() (*entity.Course, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Course), nil
	}
}

func (c courseDo) Take() (*entity.Course, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Course), nil
	}
}

func (c courseDo) Last() (*entity.Course, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Course), nil
	}
}

func (c courseDo) Find() ([]*entity.Course, error) {
	result, err := c.DO.Find()
	return result.([]*entity.Course), err
}

func (c courseDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Course, err error) {
	buf := make([]*entity.Course, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c courseDo) FindInBatches(result *[]*entity.Course, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c courseDo) Attrs(attrs ...field.AssignExpr) ICourseDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c courseDo) Assign(attrs ...field.AssignExpr) ICourseDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c courseDo) Joins(fields ...field.RelationField) ICourseDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c courseDo) Preload(fields ...field.RelationField) ICourseDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c courseDo) FirstOrInit() (*entity.Course, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Course), nil
	}
}

func (c courseDo) FirstOrCreate() (*entity.Course, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Course), nil
	}
}

func (c courseDo) FindByPage(offset int, limit int) (result []*entity.Course, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c courseDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c courseDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c courseDo) Delete(models ...*entity.Course) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *courseDo) withDO(do gen.Dao) *courseDo {
	c.DO = *do.(*gen.DO)
	return c
}
