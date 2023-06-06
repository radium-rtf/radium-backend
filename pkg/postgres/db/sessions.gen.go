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

func newSession(db *gorm.DB, opts ...gen.DOOption) session {
	_session := session{}

	_session.sessionDo.UseDB(db, opts...)
	_session.sessionDo.UseModel(&entity.Session{})

	tableName := _session.sessionDo.TableName()
	_session.ALL = field.NewAsterisk(tableName)
	_session.RefreshToken = field.NewString(tableName, "refresh_token")
	_session.ExpiresIn = field.NewTime(tableName, "expires_in")
	_session.UserId = field.NewField(tableName, "user_id")

	_session.fillFieldMap()

	return _session
}

type session struct {
	sessionDo sessionDo

	ALL          field.Asterisk
	RefreshToken field.String
	ExpiresIn    field.Time
	UserId       field.Field

	fieldMap map[string]field.Expr
}

func (s session) Table(newTableName string) *session {
	s.sessionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s session) As(alias string) *session {
	s.sessionDo.DO = *(s.sessionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *session) updateTableName(table string) *session {
	s.ALL = field.NewAsterisk(table)
	s.RefreshToken = field.NewString(table, "refresh_token")
	s.ExpiresIn = field.NewTime(table, "expires_in")
	s.UserId = field.NewField(table, "user_id")

	s.fillFieldMap()

	return s
}

func (s *session) WithContext(ctx context.Context) *sessionDo { return s.sessionDo.WithContext(ctx) }

func (s session) TableName() string { return s.sessionDo.TableName() }

func (s session) Alias() string { return s.sessionDo.Alias() }

func (s *session) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *session) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["refresh_token"] = s.RefreshToken
	s.fieldMap["expires_in"] = s.ExpiresIn
	s.fieldMap["user_id"] = s.UserId
}

func (s session) clone(db *gorm.DB) session {
	s.sessionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s session) replaceDB(db *gorm.DB) session {
	s.sessionDo.ReplaceDB(db)
	return s
}

type sessionDo struct{ gen.DO }

func (s sessionDo) Debug() *sessionDo {
	return s.withDO(s.DO.Debug())
}

func (s sessionDo) WithContext(ctx context.Context) *sessionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sessionDo) ReadDB() *sessionDo {
	return s.Clauses(dbresolver.Read)
}

func (s sessionDo) WriteDB() *sessionDo {
	return s.Clauses(dbresolver.Write)
}

func (s sessionDo) Session(config *gorm.Session) *sessionDo {
	return s.withDO(s.DO.Session(config))
}

func (s sessionDo) Clauses(conds ...clause.Expression) *sessionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sessionDo) Returning(value interface{}, columns ...string) *sessionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sessionDo) Not(conds ...gen.Condition) *sessionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sessionDo) Or(conds ...gen.Condition) *sessionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sessionDo) Select(conds ...field.Expr) *sessionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sessionDo) Where(conds ...gen.Condition) *sessionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sessionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sessionDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sessionDo) Order(conds ...field.Expr) *sessionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sessionDo) Distinct(cols ...field.Expr) *sessionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sessionDo) Omit(cols ...field.Expr) *sessionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sessionDo) Join(table schema.Tabler, on ...field.Expr) *sessionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sessionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sessionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sessionDo) RightJoin(table schema.Tabler, on ...field.Expr) *sessionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sessionDo) Group(cols ...field.Expr) *sessionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sessionDo) Having(conds ...gen.Condition) *sessionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sessionDo) Limit(limit int) *sessionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sessionDo) Offset(offset int) *sessionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sessionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sessionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sessionDo) Unscoped() *sessionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sessionDo) Create(values ...*entity.Session) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sessionDo) CreateInBatches(values []*entity.Session, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sessionDo) Save(values ...*entity.Session) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sessionDo) First() (*entity.Session, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Session), nil
	}
}

func (s sessionDo) Take() (*entity.Session, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Session), nil
	}
}

func (s sessionDo) Last() (*entity.Session, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Session), nil
	}
}

func (s sessionDo) Find() ([]*entity.Session, error) {
	result, err := s.DO.Find()
	return result.([]*entity.Session), err
}

func (s sessionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Session, err error) {
	buf := make([]*entity.Session, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sessionDo) FindInBatches(result *[]*entity.Session, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sessionDo) Attrs(attrs ...field.AssignExpr) *sessionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sessionDo) Assign(attrs ...field.AssignExpr) *sessionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sessionDo) Joins(fields ...field.RelationField) *sessionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sessionDo) Preload(fields ...field.RelationField) *sessionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sessionDo) FirstOrInit() (*entity.Session, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Session), nil
	}
}

func (s sessionDo) FirstOrCreate() (*entity.Session, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Session), nil
	}
}

func (s sessionDo) FindByPage(offset int, limit int) (result []*entity.Session, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sessionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sessionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sessionDo) Delete(models ...*entity.Session) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sessionDo) withDO(do gen.Dao) *sessionDo {
	s.DO = *do.(*gen.DO)
	return s
}
