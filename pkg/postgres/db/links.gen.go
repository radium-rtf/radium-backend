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

func newLink(db *gorm.DB, opts ...gen.DOOption) link {
	_link := link{}

	_link.linkDo.UseDB(db, opts...)
	_link.linkDo.UseModel(&entity.Link{})

	tableName := _link.linkDo.TableName()
	_link.ALL = field.NewAsterisk(tableName)
	_link.Id = field.NewField(tableName, "id")
	_link.CreatedAt = field.NewTime(tableName, "created_at")
	_link.UpdatedAt = field.NewTime(tableName, "updated_at")
	_link.DeletedAt = field.NewField(tableName, "deleted_at")
	_link.Name = field.NewString(tableName, "name")
	_link.Link = field.NewString(tableName, "link")
	_link.CourseId = field.NewField(tableName, "course_id")

	_link.fillFieldMap()

	return _link
}

type link struct {
	linkDo linkDo

	ALL       field.Asterisk
	Id        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Name      field.String
	Link      field.String
	CourseId  field.Field

	fieldMap map[string]field.Expr
}

func (l link) Table(newTableName string) *link {
	l.linkDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l link) As(alias string) *link {
	l.linkDo.DO = *(l.linkDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *link) updateTableName(table string) *link {
	l.ALL = field.NewAsterisk(table)
	l.Id = field.NewField(table, "id")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")
	l.DeletedAt = field.NewField(table, "deleted_at")
	l.Name = field.NewString(table, "name")
	l.Link = field.NewString(table, "link")
	l.CourseId = field.NewField(table, "course_id")

	l.fillFieldMap()

	return l
}

func (l *link) WithContext(ctx context.Context) ILinkDo { return l.linkDo.WithContext(ctx) }

func (l link) TableName() string { return l.linkDo.TableName() }

func (l link) Alias() string { return l.linkDo.Alias() }

func (l *link) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *link) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 7)
	l.fieldMap["id"] = l.Id
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
	l.fieldMap["deleted_at"] = l.DeletedAt
	l.fieldMap["name"] = l.Name
	l.fieldMap["link"] = l.Link
	l.fieldMap["course_id"] = l.CourseId
}

func (l link) clone(db *gorm.DB) link {
	l.linkDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l link) replaceDB(db *gorm.DB) link {
	l.linkDo.ReplaceDB(db)
	return l
}

type linkDo struct{ gen.DO }

type ILinkDo interface {
	gen.SubQuery
	Debug() ILinkDo
	WithContext(ctx context.Context) ILinkDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILinkDo
	WriteDB() ILinkDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILinkDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILinkDo
	Not(conds ...gen.Condition) ILinkDo
	Or(conds ...gen.Condition) ILinkDo
	Select(conds ...field.Expr) ILinkDo
	Where(conds ...gen.Condition) ILinkDo
	Order(conds ...field.Expr) ILinkDo
	Distinct(cols ...field.Expr) ILinkDo
	Omit(cols ...field.Expr) ILinkDo
	Join(table schema.Tabler, on ...field.Expr) ILinkDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILinkDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILinkDo
	Group(cols ...field.Expr) ILinkDo
	Having(conds ...gen.Condition) ILinkDo
	Limit(limit int) ILinkDo
	Offset(offset int) ILinkDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILinkDo
	Unscoped() ILinkDo
	Create(values ...*entity.Link) error
	CreateInBatches(values []*entity.Link, batchSize int) error
	Save(values ...*entity.Link) error
	First() (*entity.Link, error)
	Take() (*entity.Link, error)
	Last() (*entity.Link, error)
	Find() ([]*entity.Link, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Link, err error)
	FindInBatches(result *[]*entity.Link, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Link) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILinkDo
	Assign(attrs ...field.AssignExpr) ILinkDo
	Joins(fields ...field.RelationField) ILinkDo
	Preload(fields ...field.RelationField) ILinkDo
	FirstOrInit() (*entity.Link, error)
	FirstOrCreate() (*entity.Link, error)
	FindByPage(offset int, limit int) (result []*entity.Link, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILinkDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l linkDo) Debug() ILinkDo {
	return l.withDO(l.DO.Debug())
}

func (l linkDo) WithContext(ctx context.Context) ILinkDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l linkDo) ReadDB() ILinkDo {
	return l.Clauses(dbresolver.Read)
}

func (l linkDo) WriteDB() ILinkDo {
	return l.Clauses(dbresolver.Write)
}

func (l linkDo) Session(config *gorm.Session) ILinkDo {
	return l.withDO(l.DO.Session(config))
}

func (l linkDo) Clauses(conds ...clause.Expression) ILinkDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l linkDo) Returning(value interface{}, columns ...string) ILinkDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l linkDo) Not(conds ...gen.Condition) ILinkDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l linkDo) Or(conds ...gen.Condition) ILinkDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l linkDo) Select(conds ...field.Expr) ILinkDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l linkDo) Where(conds ...gen.Condition) ILinkDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l linkDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ILinkDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l linkDo) Order(conds ...field.Expr) ILinkDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l linkDo) Distinct(cols ...field.Expr) ILinkDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l linkDo) Omit(cols ...field.Expr) ILinkDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l linkDo) Join(table schema.Tabler, on ...field.Expr) ILinkDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l linkDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILinkDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l linkDo) RightJoin(table schema.Tabler, on ...field.Expr) ILinkDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l linkDo) Group(cols ...field.Expr) ILinkDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l linkDo) Having(conds ...gen.Condition) ILinkDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l linkDo) Limit(limit int) ILinkDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l linkDo) Offset(offset int) ILinkDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l linkDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILinkDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l linkDo) Unscoped() ILinkDo {
	return l.withDO(l.DO.Unscoped())
}

func (l linkDo) Create(values ...*entity.Link) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l linkDo) CreateInBatches(values []*entity.Link, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l linkDo) Save(values ...*entity.Link) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l linkDo) First() (*entity.Link, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) Take() (*entity.Link, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) Last() (*entity.Link, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) Find() ([]*entity.Link, error) {
	result, err := l.DO.Find()
	return result.([]*entity.Link), err
}

func (l linkDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Link, err error) {
	buf := make([]*entity.Link, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l linkDo) FindInBatches(result *[]*entity.Link, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l linkDo) Attrs(attrs ...field.AssignExpr) ILinkDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l linkDo) Assign(attrs ...field.AssignExpr) ILinkDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l linkDo) Joins(fields ...field.RelationField) ILinkDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l linkDo) Preload(fields ...field.RelationField) ILinkDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l linkDo) FirstOrInit() (*entity.Link, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) FirstOrCreate() (*entity.Link, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Link), nil
	}
}

func (l linkDo) FindByPage(offset int, limit int) (result []*entity.Link, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l linkDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l linkDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l linkDo) Delete(models ...*entity.Link) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *linkDo) withDO(do gen.Dao) *linkDo {
	l.DO = *do.(*gen.DO)
	return l
}
