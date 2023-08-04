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

func newAnswerSection(db *gorm.DB, opts ...gen.DOOption) answerSection {
	_answerSection := answerSection{}

	_answerSection.answerSectionDo.UseDB(db, opts...)
	_answerSection.answerSectionDo.UseModel(&entity.AnswerSection{})

	tableName := _answerSection.answerSectionDo.TableName()
	_answerSection.ALL = field.NewAsterisk(tableName)
	_answerSection.Id = field.NewField(tableName, "id")
	_answerSection.CreatedAt = field.NewTime(tableName, "created_at")
	_answerSection.UpdatedAt = field.NewTime(tableName, "updated_at")
	_answerSection.DeletedAt = field.NewField(tableName, "deleted_at")
	_answerSection.Question = field.NewString(tableName, "question")
	_answerSection.OwnerID = field.NewField(tableName, "owner_id")
	_answerSection.OwnerType = field.NewString(tableName, "owner_type")

	_answerSection.fillFieldMap()

	return _answerSection
}

type answerSection struct {
	answerSectionDo answerSectionDo

	ALL       field.Asterisk
	Id        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Question  field.String
	OwnerID   field.Field
	OwnerType field.String

	fieldMap map[string]field.Expr
}

func (a answerSection) Table(newTableName string) *answerSection {
	a.answerSectionDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a answerSection) As(alias string) *answerSection {
	a.answerSectionDo.DO = *(a.answerSectionDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *answerSection) updateTableName(table string) *answerSection {
	a.ALL = field.NewAsterisk(table)
	a.Id = field.NewField(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Question = field.NewString(table, "question")
	a.OwnerID = field.NewField(table, "owner_id")
	a.OwnerType = field.NewString(table, "owner_type")

	a.fillFieldMap()

	return a
}

func (a *answerSection) WithContext(ctx context.Context) IAnswerSectionDo {
	return a.answerSectionDo.WithContext(ctx)
}

func (a answerSection) TableName() string { return a.answerSectionDo.TableName() }

func (a answerSection) Alias() string { return a.answerSectionDo.Alias() }

func (a *answerSection) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *answerSection) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.Id
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["question"] = a.Question
	a.fieldMap["owner_id"] = a.OwnerID
	a.fieldMap["owner_type"] = a.OwnerType
}

func (a answerSection) clone(db *gorm.DB) answerSection {
	a.answerSectionDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a answerSection) replaceDB(db *gorm.DB) answerSection {
	a.answerSectionDo.ReplaceDB(db)
	return a
}

type answerSectionDo struct{ gen.DO }

type IAnswerSectionDo interface {
	gen.SubQuery
	Debug() IAnswerSectionDo
	WithContext(ctx context.Context) IAnswerSectionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAnswerSectionDo
	WriteDB() IAnswerSectionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAnswerSectionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAnswerSectionDo
	Not(conds ...gen.Condition) IAnswerSectionDo
	Or(conds ...gen.Condition) IAnswerSectionDo
	Select(conds ...field.Expr) IAnswerSectionDo
	Where(conds ...gen.Condition) IAnswerSectionDo
	Order(conds ...field.Expr) IAnswerSectionDo
	Distinct(cols ...field.Expr) IAnswerSectionDo
	Omit(cols ...field.Expr) IAnswerSectionDo
	Join(table schema.Tabler, on ...field.Expr) IAnswerSectionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAnswerSectionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAnswerSectionDo
	Group(cols ...field.Expr) IAnswerSectionDo
	Having(conds ...gen.Condition) IAnswerSectionDo
	Limit(limit int) IAnswerSectionDo
	Offset(offset int) IAnswerSectionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAnswerSectionDo
	Unscoped() IAnswerSectionDo
	Create(values ...*entity.AnswerSection) error
	CreateInBatches(values []*entity.AnswerSection, batchSize int) error
	Save(values ...*entity.AnswerSection) error
	First() (*entity.AnswerSection, error)
	Take() (*entity.AnswerSection, error)
	Last() (*entity.AnswerSection, error)
	Find() ([]*entity.AnswerSection, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.AnswerSection, err error)
	FindInBatches(result *[]*entity.AnswerSection, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.AnswerSection) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAnswerSectionDo
	Assign(attrs ...field.AssignExpr) IAnswerSectionDo
	Joins(fields ...field.RelationField) IAnswerSectionDo
	Preload(fields ...field.RelationField) IAnswerSectionDo
	FirstOrInit() (*entity.AnswerSection, error)
	FirstOrCreate() (*entity.AnswerSection, error)
	FindByPage(offset int, limit int) (result []*entity.AnswerSection, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAnswerSectionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a answerSectionDo) Debug() IAnswerSectionDo {
	return a.withDO(a.DO.Debug())
}

func (a answerSectionDo) WithContext(ctx context.Context) IAnswerSectionDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a answerSectionDo) ReadDB() IAnswerSectionDo {
	return a.Clauses(dbresolver.Read)
}

func (a answerSectionDo) WriteDB() IAnswerSectionDo {
	return a.Clauses(dbresolver.Write)
}

func (a answerSectionDo) Session(config *gorm.Session) IAnswerSectionDo {
	return a.withDO(a.DO.Session(config))
}

func (a answerSectionDo) Clauses(conds ...clause.Expression) IAnswerSectionDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a answerSectionDo) Returning(value interface{}, columns ...string) IAnswerSectionDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a answerSectionDo) Not(conds ...gen.Condition) IAnswerSectionDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a answerSectionDo) Or(conds ...gen.Condition) IAnswerSectionDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a answerSectionDo) Select(conds ...field.Expr) IAnswerSectionDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a answerSectionDo) Where(conds ...gen.Condition) IAnswerSectionDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a answerSectionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAnswerSectionDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a answerSectionDo) Order(conds ...field.Expr) IAnswerSectionDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a answerSectionDo) Distinct(cols ...field.Expr) IAnswerSectionDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a answerSectionDo) Omit(cols ...field.Expr) IAnswerSectionDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a answerSectionDo) Join(table schema.Tabler, on ...field.Expr) IAnswerSectionDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a answerSectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAnswerSectionDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a answerSectionDo) RightJoin(table schema.Tabler, on ...field.Expr) IAnswerSectionDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a answerSectionDo) Group(cols ...field.Expr) IAnswerSectionDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a answerSectionDo) Having(conds ...gen.Condition) IAnswerSectionDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a answerSectionDo) Limit(limit int) IAnswerSectionDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a answerSectionDo) Offset(offset int) IAnswerSectionDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a answerSectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAnswerSectionDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a answerSectionDo) Unscoped() IAnswerSectionDo {
	return a.withDO(a.DO.Unscoped())
}

func (a answerSectionDo) Create(values ...*entity.AnswerSection) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a answerSectionDo) CreateInBatches(values []*entity.AnswerSection, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a answerSectionDo) Save(values ...*entity.AnswerSection) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a answerSectionDo) First() (*entity.AnswerSection, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerSection), nil
	}
}

func (a answerSectionDo) Take() (*entity.AnswerSection, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerSection), nil
	}
}

func (a answerSectionDo) Last() (*entity.AnswerSection, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerSection), nil
	}
}

func (a answerSectionDo) Find() ([]*entity.AnswerSection, error) {
	result, err := a.DO.Find()
	return result.([]*entity.AnswerSection), err
}

func (a answerSectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.AnswerSection, err error) {
	buf := make([]*entity.AnswerSection, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a answerSectionDo) FindInBatches(result *[]*entity.AnswerSection, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a answerSectionDo) Attrs(attrs ...field.AssignExpr) IAnswerSectionDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a answerSectionDo) Assign(attrs ...field.AssignExpr) IAnswerSectionDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a answerSectionDo) Joins(fields ...field.RelationField) IAnswerSectionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a answerSectionDo) Preload(fields ...field.RelationField) IAnswerSectionDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a answerSectionDo) FirstOrInit() (*entity.AnswerSection, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerSection), nil
	}
}

func (a answerSectionDo) FirstOrCreate() (*entity.AnswerSection, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerSection), nil
	}
}

func (a answerSectionDo) FindByPage(offset int, limit int) (result []*entity.AnswerSection, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a answerSectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a answerSectionDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a answerSectionDo) Delete(models ...*entity.AnswerSection) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *answerSectionDo) withDO(do gen.Dao) *answerSectionDo {
	a.DO = *do.(*gen.DO)
	return a
}
