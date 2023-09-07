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

func newModule(db *gorm.DB, opts ...gen.DOOption) module {
	_module := module{}

	_module.moduleDo.UseDB(db, opts...)
	_module.moduleDo.UseModel(&entity.Module{})

	tableName := _module.moduleDo.TableName()
	_module.ALL = field.NewAsterisk(tableName)
	_module.Id = field.NewField(tableName, "id")
	_module.CreatedAt = field.NewTime(tableName, "created_at")
	_module.UpdatedAt = field.NewTime(tableName, "updated_at")
	_module.DeletedAt = field.NewField(tableName, "deleted_at")
	_module.Slug = field.NewString(tableName, "slug")
	_module.Name = field.NewString(tableName, "name")
	_module.CourseId = field.NewField(tableName, "course_id")
	_module.Order = field.NewFloat64(tableName, "order")
	_module.Pages = moduleHasManyPages{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Pages", "entity.Page"),
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
			CodeSection struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Pages.Sections", "entity.Section"),
			TextSection: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Pages.Sections.TextSection", "entity.TextSection"),
			},
			ChoiceSection: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Pages.Sections.ChoiceSection", "entity.ChoiceSection"),
			},
			MultiChoiceSection: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Pages.Sections.MultiChoiceSection", "entity.MultiChoiceSection"),
			},
			ShortAnswerSection: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Pages.Sections.ShortAnswerSection", "entity.ShortAnswerSection"),
			},
			AnswerSection: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Pages.Sections.AnswerSection", "entity.AnswerSection"),
			},
			CodeSection: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Pages.Sections.CodeSection", "entity.CodeSection"),
			},
		},
	}

	_module.fillFieldMap()

	return _module
}

type module struct {
	moduleDo moduleDo

	ALL       field.Asterisk
	Id        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Slug      field.String
	Name      field.String
	CourseId  field.Field
	Order     field.Float64
	Pages     moduleHasManyPages

	fieldMap map[string]field.Expr
}

func (m module) Table(newTableName string) *module {
	m.moduleDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m module) As(alias string) *module {
	m.moduleDo.DO = *(m.moduleDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *module) updateTableName(table string) *module {
	m.ALL = field.NewAsterisk(table)
	m.Id = field.NewField(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.Slug = field.NewString(table, "slug")
	m.Name = field.NewString(table, "name")
	m.CourseId = field.NewField(table, "course_id")
	m.Order = field.NewFloat64(table, "order")

	m.fillFieldMap()

	return m
}

func (m *module) WithContext(ctx context.Context) IModuleDo { return m.moduleDo.WithContext(ctx) }

func (m module) TableName() string { return m.moduleDo.TableName() }

func (m module) Alias() string { return m.moduleDo.Alias() }

func (m module) Columns(cols ...field.Expr) gen.Columns { return m.moduleDo.Columns(cols...) }

func (m *module) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *module) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 9)
	m.fieldMap["id"] = m.Id
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["slug"] = m.Slug
	m.fieldMap["name"] = m.Name
	m.fieldMap["course_id"] = m.CourseId
	m.fieldMap["order"] = m.Order

}

func (m module) clone(db *gorm.DB) module {
	m.moduleDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m module) replaceDB(db *gorm.DB) module {
	m.moduleDo.ReplaceDB(db)
	return m
}

type moduleHasManyPages struct {
	db *gorm.DB

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
		CodeSection struct {
			field.RelationField
		}
	}
}

func (a moduleHasManyPages) Where(conds ...field.Expr) *moduleHasManyPages {
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

func (a moduleHasManyPages) WithContext(ctx context.Context) *moduleHasManyPages {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a moduleHasManyPages) Session(session *gorm.Session) *moduleHasManyPages {
	a.db = a.db.Session(session)
	return &a
}

func (a moduleHasManyPages) Model(m *entity.Module) *moduleHasManyPagesTx {
	return &moduleHasManyPagesTx{a.db.Model(m).Association(a.Name())}
}

type moduleHasManyPagesTx struct{ tx *gorm.Association }

func (a moduleHasManyPagesTx) Find() (result []*entity.Page, err error) {
	return result, a.tx.Find(&result)
}

func (a moduleHasManyPagesTx) Append(values ...*entity.Page) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a moduleHasManyPagesTx) Replace(values ...*entity.Page) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a moduleHasManyPagesTx) Delete(values ...*entity.Page) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a moduleHasManyPagesTx) Clear() error {
	return a.tx.Clear()
}

func (a moduleHasManyPagesTx) Count() int64 {
	return a.tx.Count()
}

type moduleDo struct{ gen.DO }

type IModuleDo interface {
	gen.SubQuery
	Debug() IModuleDo
	WithContext(ctx context.Context) IModuleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IModuleDo
	WriteDB() IModuleDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IModuleDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IModuleDo
	Not(conds ...gen.Condition) IModuleDo
	Or(conds ...gen.Condition) IModuleDo
	Select(conds ...field.Expr) IModuleDo
	Where(conds ...gen.Condition) IModuleDo
	Order(conds ...field.Expr) IModuleDo
	Distinct(cols ...field.Expr) IModuleDo
	Omit(cols ...field.Expr) IModuleDo
	Join(table schema.Tabler, on ...field.Expr) IModuleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IModuleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IModuleDo
	Group(cols ...field.Expr) IModuleDo
	Having(conds ...gen.Condition) IModuleDo
	Limit(limit int) IModuleDo
	Offset(offset int) IModuleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IModuleDo
	Unscoped() IModuleDo
	Create(values ...*entity.Module) error
	CreateInBatches(values []*entity.Module, batchSize int) error
	Save(values ...*entity.Module) error
	First() (*entity.Module, error)
	Take() (*entity.Module, error)
	Last() (*entity.Module, error)
	Find() ([]*entity.Module, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Module, err error)
	FindInBatches(result *[]*entity.Module, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Module) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IModuleDo
	Assign(attrs ...field.AssignExpr) IModuleDo
	Joins(fields ...field.RelationField) IModuleDo
	Preload(fields ...field.RelationField) IModuleDo
	FirstOrInit() (*entity.Module, error)
	FirstOrCreate() (*entity.Module, error)
	FindByPage(offset int, limit int) (result []*entity.Module, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IModuleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m moduleDo) Debug() IModuleDo {
	return m.withDO(m.DO.Debug())
}

func (m moduleDo) WithContext(ctx context.Context) IModuleDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m moduleDo) ReadDB() IModuleDo {
	return m.Clauses(dbresolver.Read)
}

func (m moduleDo) WriteDB() IModuleDo {
	return m.Clauses(dbresolver.Write)
}

func (m moduleDo) Session(config *gorm.Session) IModuleDo {
	return m.withDO(m.DO.Session(config))
}

func (m moduleDo) Clauses(conds ...clause.Expression) IModuleDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m moduleDo) Returning(value interface{}, columns ...string) IModuleDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m moduleDo) Not(conds ...gen.Condition) IModuleDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m moduleDo) Or(conds ...gen.Condition) IModuleDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m moduleDo) Select(conds ...field.Expr) IModuleDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m moduleDo) Where(conds ...gen.Condition) IModuleDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m moduleDo) Order(conds ...field.Expr) IModuleDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m moduleDo) Distinct(cols ...field.Expr) IModuleDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m moduleDo) Omit(cols ...field.Expr) IModuleDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m moduleDo) Join(table schema.Tabler, on ...field.Expr) IModuleDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m moduleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IModuleDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m moduleDo) RightJoin(table schema.Tabler, on ...field.Expr) IModuleDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m moduleDo) Group(cols ...field.Expr) IModuleDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m moduleDo) Having(conds ...gen.Condition) IModuleDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m moduleDo) Limit(limit int) IModuleDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m moduleDo) Offset(offset int) IModuleDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m moduleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IModuleDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m moduleDo) Unscoped() IModuleDo {
	return m.withDO(m.DO.Unscoped())
}

func (m moduleDo) Create(values ...*entity.Module) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m moduleDo) CreateInBatches(values []*entity.Module, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m moduleDo) Save(values ...*entity.Module) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m moduleDo) First() (*entity.Module, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Module), nil
	}
}

func (m moduleDo) Take() (*entity.Module, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Module), nil
	}
}

func (m moduleDo) Last() (*entity.Module, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Module), nil
	}
}

func (m moduleDo) Find() ([]*entity.Module, error) {
	result, err := m.DO.Find()
	return result.([]*entity.Module), err
}

func (m moduleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Module, err error) {
	buf := make([]*entity.Module, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m moduleDo) FindInBatches(result *[]*entity.Module, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m moduleDo) Attrs(attrs ...field.AssignExpr) IModuleDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m moduleDo) Assign(attrs ...field.AssignExpr) IModuleDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m moduleDo) Joins(fields ...field.RelationField) IModuleDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m moduleDo) Preload(fields ...field.RelationField) IModuleDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m moduleDo) FirstOrInit() (*entity.Module, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Module), nil
	}
}

func (m moduleDo) FirstOrCreate() (*entity.Module, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Module), nil
	}
}

func (m moduleDo) FindByPage(offset int, limit int) (result []*entity.Module, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m moduleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m moduleDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m moduleDo) Delete(models ...*entity.Module) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *moduleDo) withDO(do gen.Dao) *moduleDo {
	m.DO = *do.(*gen.DO)
	return m
}
