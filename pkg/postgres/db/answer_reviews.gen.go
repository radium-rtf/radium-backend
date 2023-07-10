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

func newAnswerReview(db *gorm.DB, opts ...gen.DOOption) answerReview {
	_answerReview := answerReview{}

	_answerReview.answerReviewDo.UseDB(db, opts...)
	_answerReview.answerReviewDo.UseModel(&entity.AnswerReview{})

	tableName := _answerReview.answerReviewDo.TableName()
	_answerReview.ALL = field.NewAsterisk(tableName)
	_answerReview.Id = field.NewField(tableName, "id")
	_answerReview.CreatedAt = field.NewTime(tableName, "created_at")
	_answerReview.UpdatedAt = field.NewTime(tableName, "updated_at")
	_answerReview.DeletedAt = field.NewField(tableName, "deleted_at")
	_answerReview.OwnerId = field.NewField(tableName, "owner_id")
	_answerReview.Score = field.NewFloat32(tableName, "score")

	_answerReview.fillFieldMap()

	return _answerReview
}

type answerReview struct {
	answerReviewDo answerReviewDo

	ALL       field.Asterisk
	Id        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	OwnerId   field.Field
	Score     field.Float32

	fieldMap map[string]field.Expr
}

func (a answerReview) Table(newTableName string) *answerReview {
	a.answerReviewDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a answerReview) As(alias string) *answerReview {
	a.answerReviewDo.DO = *(a.answerReviewDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *answerReview) updateTableName(table string) *answerReview {
	a.ALL = field.NewAsterisk(table)
	a.Id = field.NewField(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.OwnerId = field.NewField(table, "owner_id")
	a.Score = field.NewFloat32(table, "score")

	a.fillFieldMap()

	return a
}

func (a *answerReview) WithContext(ctx context.Context) IAnswerReviewDo {
	return a.answerReviewDo.WithContext(ctx)
}

func (a answerReview) TableName() string { return a.answerReviewDo.TableName() }

func (a answerReview) Alias() string { return a.answerReviewDo.Alias() }

func (a *answerReview) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *answerReview) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.Id
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["owner_id"] = a.OwnerId
	a.fieldMap["score"] = a.Score
}

func (a answerReview) clone(db *gorm.DB) answerReview {
	a.answerReviewDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a answerReview) replaceDB(db *gorm.DB) answerReview {
	a.answerReviewDo.ReplaceDB(db)
	return a
}

type answerReviewDo struct{ gen.DO }

type IAnswerReviewDo interface {
	gen.SubQuery
	Debug() IAnswerReviewDo
	WithContext(ctx context.Context) IAnswerReviewDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAnswerReviewDo
	WriteDB() IAnswerReviewDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAnswerReviewDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAnswerReviewDo
	Not(conds ...gen.Condition) IAnswerReviewDo
	Or(conds ...gen.Condition) IAnswerReviewDo
	Select(conds ...field.Expr) IAnswerReviewDo
	Where(conds ...gen.Condition) IAnswerReviewDo
	Order(conds ...field.Expr) IAnswerReviewDo
	Distinct(cols ...field.Expr) IAnswerReviewDo
	Omit(cols ...field.Expr) IAnswerReviewDo
	Join(table schema.Tabler, on ...field.Expr) IAnswerReviewDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAnswerReviewDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAnswerReviewDo
	Group(cols ...field.Expr) IAnswerReviewDo
	Having(conds ...gen.Condition) IAnswerReviewDo
	Limit(limit int) IAnswerReviewDo
	Offset(offset int) IAnswerReviewDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAnswerReviewDo
	Unscoped() IAnswerReviewDo
	Create(values ...*entity.AnswerReview) error
	CreateInBatches(values []*entity.AnswerReview, batchSize int) error
	Save(values ...*entity.AnswerReview) error
	First() (*entity.AnswerReview, error)
	Take() (*entity.AnswerReview, error)
	Last() (*entity.AnswerReview, error)
	Find() ([]*entity.AnswerReview, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.AnswerReview, err error)
	FindInBatches(result *[]*entity.AnswerReview, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.AnswerReview) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAnswerReviewDo
	Assign(attrs ...field.AssignExpr) IAnswerReviewDo
	Joins(fields ...field.RelationField) IAnswerReviewDo
	Preload(fields ...field.RelationField) IAnswerReviewDo
	FirstOrInit() (*entity.AnswerReview, error)
	FirstOrCreate() (*entity.AnswerReview, error)
	FindByPage(offset int, limit int) (result []*entity.AnswerReview, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAnswerReviewDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a answerReviewDo) Debug() IAnswerReviewDo {
	return a.withDO(a.DO.Debug())
}

func (a answerReviewDo) WithContext(ctx context.Context) IAnswerReviewDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a answerReviewDo) ReadDB() IAnswerReviewDo {
	return a.Clauses(dbresolver.Read)
}

func (a answerReviewDo) WriteDB() IAnswerReviewDo {
	return a.Clauses(dbresolver.Write)
}

func (a answerReviewDo) Session(config *gorm.Session) IAnswerReviewDo {
	return a.withDO(a.DO.Session(config))
}

func (a answerReviewDo) Clauses(conds ...clause.Expression) IAnswerReviewDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a answerReviewDo) Returning(value interface{}, columns ...string) IAnswerReviewDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a answerReviewDo) Not(conds ...gen.Condition) IAnswerReviewDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a answerReviewDo) Or(conds ...gen.Condition) IAnswerReviewDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a answerReviewDo) Select(conds ...field.Expr) IAnswerReviewDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a answerReviewDo) Where(conds ...gen.Condition) IAnswerReviewDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a answerReviewDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAnswerReviewDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a answerReviewDo) Order(conds ...field.Expr) IAnswerReviewDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a answerReviewDo) Distinct(cols ...field.Expr) IAnswerReviewDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a answerReviewDo) Omit(cols ...field.Expr) IAnswerReviewDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a answerReviewDo) Join(table schema.Tabler, on ...field.Expr) IAnswerReviewDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a answerReviewDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAnswerReviewDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a answerReviewDo) RightJoin(table schema.Tabler, on ...field.Expr) IAnswerReviewDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a answerReviewDo) Group(cols ...field.Expr) IAnswerReviewDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a answerReviewDo) Having(conds ...gen.Condition) IAnswerReviewDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a answerReviewDo) Limit(limit int) IAnswerReviewDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a answerReviewDo) Offset(offset int) IAnswerReviewDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a answerReviewDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAnswerReviewDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a answerReviewDo) Unscoped() IAnswerReviewDo {
	return a.withDO(a.DO.Unscoped())
}

func (a answerReviewDo) Create(values ...*entity.AnswerReview) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a answerReviewDo) CreateInBatches(values []*entity.AnswerReview, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a answerReviewDo) Save(values ...*entity.AnswerReview) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a answerReviewDo) First() (*entity.AnswerReview, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerReview), nil
	}
}

func (a answerReviewDo) Take() (*entity.AnswerReview, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerReview), nil
	}
}

func (a answerReviewDo) Last() (*entity.AnswerReview, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerReview), nil
	}
}

func (a answerReviewDo) Find() ([]*entity.AnswerReview, error) {
	result, err := a.DO.Find()
	return result.([]*entity.AnswerReview), err
}

func (a answerReviewDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.AnswerReview, err error) {
	buf := make([]*entity.AnswerReview, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a answerReviewDo) FindInBatches(result *[]*entity.AnswerReview, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a answerReviewDo) Attrs(attrs ...field.AssignExpr) IAnswerReviewDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a answerReviewDo) Assign(attrs ...field.AssignExpr) IAnswerReviewDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a answerReviewDo) Joins(fields ...field.RelationField) IAnswerReviewDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a answerReviewDo) Preload(fields ...field.RelationField) IAnswerReviewDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a answerReviewDo) FirstOrInit() (*entity.AnswerReview, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerReview), nil
	}
}

func (a answerReviewDo) FirstOrCreate() (*entity.AnswerReview, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AnswerReview), nil
	}
}

func (a answerReviewDo) FindByPage(offset int, limit int) (result []*entity.AnswerReview, count int64, err error) {
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

func (a answerReviewDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a answerReviewDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a answerReviewDo) Delete(models ...*entity.AnswerReview) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *answerReviewDo) withDO(do gen.Dao) *answerReviewDo {
	a.DO = *do.(*gen.DO)
	return a
}
