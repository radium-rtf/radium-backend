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

func newPermutationSection(db *gorm.DB, opts ...gen.DOOption) permutationSection {
	_permutationSection := permutationSection{}

	_permutationSection.permutationSectionDo.UseDB(db, opts...)
	_permutationSection.permutationSectionDo.UseModel(&entity.PermutationSection{})

	tableName := _permutationSection.permutationSectionDo.TableName()
	_permutationSection.ALL = field.NewAsterisk(tableName)
	_permutationSection.Id = field.NewField(tableName, "id")
	_permutationSection.CreatedAt = field.NewTime(tableName, "created_at")
	_permutationSection.UpdatedAt = field.NewTime(tableName, "updated_at")
	_permutationSection.DeletedAt = field.NewField(tableName, "deleted_at")
	_permutationSection.Question = field.NewString(tableName, "question")
	_permutationSection.Answer = field.NewField(tableName, "answer")
	_permutationSection.OwnerID = field.NewField(tableName, "owner_id")
	_permutationSection.OwnerType = field.NewString(tableName, "owner_type")

	_permutationSection.fillFieldMap()

	return _permutationSection
}

type permutationSection struct {
	permutationSectionDo permutationSectionDo

	ALL       field.Asterisk
	Id        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Question  field.String
	Answer    field.Field
	OwnerID   field.Field
	OwnerType field.String

	fieldMap map[string]field.Expr
}

func (p permutationSection) Table(newTableName string) *permutationSection {
	p.permutationSectionDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p permutationSection) As(alias string) *permutationSection {
	p.permutationSectionDo.DO = *(p.permutationSectionDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *permutationSection) updateTableName(table string) *permutationSection {
	p.ALL = field.NewAsterisk(table)
	p.Id = field.NewField(table, "id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")
	p.Question = field.NewString(table, "question")
	p.Answer = field.NewField(table, "answer")
	p.OwnerID = field.NewField(table, "owner_id")
	p.OwnerType = field.NewString(table, "owner_type")

	p.fillFieldMap()

	return p
}

func (p *permutationSection) WithContext(ctx context.Context) IPermutationSectionDo {
	return p.permutationSectionDo.WithContext(ctx)
}

func (p permutationSection) TableName() string { return p.permutationSectionDo.TableName() }

func (p permutationSection) Alias() string { return p.permutationSectionDo.Alias() }

func (p permutationSection) Columns(cols ...field.Expr) gen.Columns {
	return p.permutationSectionDo.Columns(cols...)
}

func (p *permutationSection) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *permutationSection) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 8)
	p.fieldMap["id"] = p.Id
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
	p.fieldMap["question"] = p.Question
	p.fieldMap["answer"] = p.Answer
	p.fieldMap["owner_id"] = p.OwnerID
	p.fieldMap["owner_type"] = p.OwnerType
}

func (p permutationSection) clone(db *gorm.DB) permutationSection {
	p.permutationSectionDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p permutationSection) replaceDB(db *gorm.DB) permutationSection {
	p.permutationSectionDo.ReplaceDB(db)
	return p
}

type permutationSectionDo struct{ gen.DO }

type IPermutationSectionDo interface {
	gen.SubQuery
	Debug() IPermutationSectionDo
	WithContext(ctx context.Context) IPermutationSectionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPermutationSectionDo
	WriteDB() IPermutationSectionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPermutationSectionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPermutationSectionDo
	Not(conds ...gen.Condition) IPermutationSectionDo
	Or(conds ...gen.Condition) IPermutationSectionDo
	Select(conds ...field.Expr) IPermutationSectionDo
	Where(conds ...gen.Condition) IPermutationSectionDo
	Order(conds ...field.Expr) IPermutationSectionDo
	Distinct(cols ...field.Expr) IPermutationSectionDo
	Omit(cols ...field.Expr) IPermutationSectionDo
	Join(table schema.Tabler, on ...field.Expr) IPermutationSectionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPermutationSectionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPermutationSectionDo
	Group(cols ...field.Expr) IPermutationSectionDo
	Having(conds ...gen.Condition) IPermutationSectionDo
	Limit(limit int) IPermutationSectionDo
	Offset(offset int) IPermutationSectionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPermutationSectionDo
	Unscoped() IPermutationSectionDo
	Create(values ...*entity.PermutationSection) error
	CreateInBatches(values []*entity.PermutationSection, batchSize int) error
	Save(values ...*entity.PermutationSection) error
	First() (*entity.PermutationSection, error)
	Take() (*entity.PermutationSection, error)
	Last() (*entity.PermutationSection, error)
	Find() ([]*entity.PermutationSection, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.PermutationSection, err error)
	FindInBatches(result *[]*entity.PermutationSection, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.PermutationSection) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPermutationSectionDo
	Assign(attrs ...field.AssignExpr) IPermutationSectionDo
	Joins(fields ...field.RelationField) IPermutationSectionDo
	Preload(fields ...field.RelationField) IPermutationSectionDo
	FirstOrInit() (*entity.PermutationSection, error)
	FirstOrCreate() (*entity.PermutationSection, error)
	FindByPage(offset int, limit int) (result []*entity.PermutationSection, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPermutationSectionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p permutationSectionDo) Debug() IPermutationSectionDo {
	return p.withDO(p.DO.Debug())
}

func (p permutationSectionDo) WithContext(ctx context.Context) IPermutationSectionDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p permutationSectionDo) ReadDB() IPermutationSectionDo {
	return p.Clauses(dbresolver.Read)
}

func (p permutationSectionDo) WriteDB() IPermutationSectionDo {
	return p.Clauses(dbresolver.Write)
}

func (p permutationSectionDo) Session(config *gorm.Session) IPermutationSectionDo {
	return p.withDO(p.DO.Session(config))
}

func (p permutationSectionDo) Clauses(conds ...clause.Expression) IPermutationSectionDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p permutationSectionDo) Returning(value interface{}, columns ...string) IPermutationSectionDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p permutationSectionDo) Not(conds ...gen.Condition) IPermutationSectionDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p permutationSectionDo) Or(conds ...gen.Condition) IPermutationSectionDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p permutationSectionDo) Select(conds ...field.Expr) IPermutationSectionDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p permutationSectionDo) Where(conds ...gen.Condition) IPermutationSectionDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p permutationSectionDo) Order(conds ...field.Expr) IPermutationSectionDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p permutationSectionDo) Distinct(cols ...field.Expr) IPermutationSectionDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p permutationSectionDo) Omit(cols ...field.Expr) IPermutationSectionDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p permutationSectionDo) Join(table schema.Tabler, on ...field.Expr) IPermutationSectionDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p permutationSectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPermutationSectionDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p permutationSectionDo) RightJoin(table schema.Tabler, on ...field.Expr) IPermutationSectionDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p permutationSectionDo) Group(cols ...field.Expr) IPermutationSectionDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p permutationSectionDo) Having(conds ...gen.Condition) IPermutationSectionDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p permutationSectionDo) Limit(limit int) IPermutationSectionDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p permutationSectionDo) Offset(offset int) IPermutationSectionDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p permutationSectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPermutationSectionDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p permutationSectionDo) Unscoped() IPermutationSectionDo {
	return p.withDO(p.DO.Unscoped())
}

func (p permutationSectionDo) Create(values ...*entity.PermutationSection) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p permutationSectionDo) CreateInBatches(values []*entity.PermutationSection, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p permutationSectionDo) Save(values ...*entity.PermutationSection) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p permutationSectionDo) First() (*entity.PermutationSection, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PermutationSection), nil
	}
}

func (p permutationSectionDo) Take() (*entity.PermutationSection, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PermutationSection), nil
	}
}

func (p permutationSectionDo) Last() (*entity.PermutationSection, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PermutationSection), nil
	}
}

func (p permutationSectionDo) Find() ([]*entity.PermutationSection, error) {
	result, err := p.DO.Find()
	return result.([]*entity.PermutationSection), err
}

func (p permutationSectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.PermutationSection, err error) {
	buf := make([]*entity.PermutationSection, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p permutationSectionDo) FindInBatches(result *[]*entity.PermutationSection, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p permutationSectionDo) Attrs(attrs ...field.AssignExpr) IPermutationSectionDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p permutationSectionDo) Assign(attrs ...field.AssignExpr) IPermutationSectionDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p permutationSectionDo) Joins(fields ...field.RelationField) IPermutationSectionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p permutationSectionDo) Preload(fields ...field.RelationField) IPermutationSectionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p permutationSectionDo) FirstOrInit() (*entity.PermutationSection, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PermutationSection), nil
	}
}

func (p permutationSectionDo) FirstOrCreate() (*entity.PermutationSection, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.PermutationSection), nil
	}
}

func (p permutationSectionDo) FindByPage(offset int, limit int) (result []*entity.PermutationSection, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p permutationSectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p permutationSectionDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p permutationSectionDo) Delete(models ...*entity.PermutationSection) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *permutationSectionDo) withDO(do gen.Dao) *permutationSectionDo {
	p.DO = *do.(*gen.DO)
	return p
}
