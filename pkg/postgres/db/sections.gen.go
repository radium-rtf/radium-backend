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

func newSection(db *gorm.DB, opts ...gen.DOOption) section {
	_section := section{}

	_section.sectionDo.UseDB(db, opts...)
	_section.sectionDo.UseModel(&entity.Section{})

	tableName := _section.sectionDo.TableName()
	_section.ALL = field.NewAsterisk(tableName)
	_section.Id = field.NewField(tableName, "id")
	_section.CreatedAt = field.NewTime(tableName, "created_at")
	_section.UpdatedAt = field.NewTime(tableName, "updated_at")
	_section.DeletedAt = field.NewField(tableName, "deleted_at")
	_section.PageId = field.NewField(tableName, "page_id")
	_section.Order = field.NewUint(tableName, "order")
	_section.TextSection = sectionHasOneTextSection{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("TextSection", "entity.TextSection"),
	}

	_section.ChoiceSection = sectionHasOneChoiceSection{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ChoiceSection", "entity.ChoiceSection"),
	}

	_section.MultiChoiceSection = sectionHasOneMultiChoiceSection{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("MultiChoiceSection", "entity.MultiChoiceSection"),
	}

	_section.ShortAnswerSection = sectionHasOneShortAnswerSection{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ShortAnswerSection", "entity.ShortAnswerSection"),
	}

	_section.AnswerSection = sectionHasOneAnswerSection{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("AnswerSection", "entity.AnswerSection"),
	}

	_section.fillFieldMap()

	return _section
}

type section struct {
	sectionDo sectionDo

	ALL         field.Asterisk
	Id          field.Field
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	PageId      field.Field
	Order       field.Uint
	TextSection sectionHasOneTextSection

	ChoiceSection sectionHasOneChoiceSection

	MultiChoiceSection sectionHasOneMultiChoiceSection

	ShortAnswerSection sectionHasOneShortAnswerSection

	AnswerSection sectionHasOneAnswerSection

	fieldMap map[string]field.Expr
}

func (s section) Table(newTableName string) *section {
	s.sectionDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s section) As(alias string) *section {
	s.sectionDo.DO = *(s.sectionDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *section) updateTableName(table string) *section {
	s.ALL = field.NewAsterisk(table)
	s.Id = field.NewField(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.PageId = field.NewField(table, "page_id")
	s.Order = field.NewUint(table, "order")

	s.fillFieldMap()

	return s
}

func (s *section) WithContext(ctx context.Context) ISectionDo { return s.sectionDo.WithContext(ctx) }

func (s section) TableName() string { return s.sectionDo.TableName() }

func (s section) Alias() string { return s.sectionDo.Alias() }

func (s *section) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *section) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["id"] = s.Id
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["page_id"] = s.PageId
	s.fieldMap["order"] = s.Order

}

func (s section) clone(db *gorm.DB) section {
	s.sectionDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s section) replaceDB(db *gorm.DB) section {
	s.sectionDo.ReplaceDB(db)
	return s
}

type sectionHasOneTextSection struct {
	db *gorm.DB

	field.RelationField
}

func (a sectionHasOneTextSection) Where(conds ...field.Expr) *sectionHasOneTextSection {
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

func (a sectionHasOneTextSection) WithContext(ctx context.Context) *sectionHasOneTextSection {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a sectionHasOneTextSection) Session(session *gorm.Session) *sectionHasOneTextSection {
	a.db = a.db.Session(session)
	return &a
}

func (a sectionHasOneTextSection) Model(m *entity.Section) *sectionHasOneTextSectionTx {
	return &sectionHasOneTextSectionTx{a.db.Model(m).Association(a.Name())}
}

type sectionHasOneTextSectionTx struct{ tx *gorm.Association }

func (a sectionHasOneTextSectionTx) Find() (result *entity.TextSection, err error) {
	return result, a.tx.Find(&result)
}

func (a sectionHasOneTextSectionTx) Append(values ...*entity.TextSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a sectionHasOneTextSectionTx) Replace(values ...*entity.TextSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a sectionHasOneTextSectionTx) Delete(values ...*entity.TextSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a sectionHasOneTextSectionTx) Clear() error {
	return a.tx.Clear()
}

func (a sectionHasOneTextSectionTx) Count() int64 {
	return a.tx.Count()
}

type sectionHasOneChoiceSection struct {
	db *gorm.DB

	field.RelationField
}

func (a sectionHasOneChoiceSection) Where(conds ...field.Expr) *sectionHasOneChoiceSection {
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

func (a sectionHasOneChoiceSection) WithContext(ctx context.Context) *sectionHasOneChoiceSection {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a sectionHasOneChoiceSection) Session(session *gorm.Session) *sectionHasOneChoiceSection {
	a.db = a.db.Session(session)
	return &a
}

func (a sectionHasOneChoiceSection) Model(m *entity.Section) *sectionHasOneChoiceSectionTx {
	return &sectionHasOneChoiceSectionTx{a.db.Model(m).Association(a.Name())}
}

type sectionHasOneChoiceSectionTx struct{ tx *gorm.Association }

func (a sectionHasOneChoiceSectionTx) Find() (result *entity.ChoiceSection, err error) {
	return result, a.tx.Find(&result)
}

func (a sectionHasOneChoiceSectionTx) Append(values ...*entity.ChoiceSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a sectionHasOneChoiceSectionTx) Replace(values ...*entity.ChoiceSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a sectionHasOneChoiceSectionTx) Delete(values ...*entity.ChoiceSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a sectionHasOneChoiceSectionTx) Clear() error {
	return a.tx.Clear()
}

func (a sectionHasOneChoiceSectionTx) Count() int64 {
	return a.tx.Count()
}

type sectionHasOneMultiChoiceSection struct {
	db *gorm.DB

	field.RelationField
}

func (a sectionHasOneMultiChoiceSection) Where(conds ...field.Expr) *sectionHasOneMultiChoiceSection {
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

func (a sectionHasOneMultiChoiceSection) WithContext(ctx context.Context) *sectionHasOneMultiChoiceSection {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a sectionHasOneMultiChoiceSection) Session(session *gorm.Session) *sectionHasOneMultiChoiceSection {
	a.db = a.db.Session(session)
	return &a
}

func (a sectionHasOneMultiChoiceSection) Model(m *entity.Section) *sectionHasOneMultiChoiceSectionTx {
	return &sectionHasOneMultiChoiceSectionTx{a.db.Model(m).Association(a.Name())}
}

type sectionHasOneMultiChoiceSectionTx struct{ tx *gorm.Association }

func (a sectionHasOneMultiChoiceSectionTx) Find() (result *entity.MultiChoiceSection, err error) {
	return result, a.tx.Find(&result)
}

func (a sectionHasOneMultiChoiceSectionTx) Append(values ...*entity.MultiChoiceSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a sectionHasOneMultiChoiceSectionTx) Replace(values ...*entity.MultiChoiceSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a sectionHasOneMultiChoiceSectionTx) Delete(values ...*entity.MultiChoiceSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a sectionHasOneMultiChoiceSectionTx) Clear() error {
	return a.tx.Clear()
}

func (a sectionHasOneMultiChoiceSectionTx) Count() int64 {
	return a.tx.Count()
}

type sectionHasOneShortAnswerSection struct {
	db *gorm.DB

	field.RelationField
}

func (a sectionHasOneShortAnswerSection) Where(conds ...field.Expr) *sectionHasOneShortAnswerSection {
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

func (a sectionHasOneShortAnswerSection) WithContext(ctx context.Context) *sectionHasOneShortAnswerSection {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a sectionHasOneShortAnswerSection) Session(session *gorm.Session) *sectionHasOneShortAnswerSection {
	a.db = a.db.Session(session)
	return &a
}

func (a sectionHasOneShortAnswerSection) Model(m *entity.Section) *sectionHasOneShortAnswerSectionTx {
	return &sectionHasOneShortAnswerSectionTx{a.db.Model(m).Association(a.Name())}
}

type sectionHasOneShortAnswerSectionTx struct{ tx *gorm.Association }

func (a sectionHasOneShortAnswerSectionTx) Find() (result *entity.ShortAnswerSection, err error) {
	return result, a.tx.Find(&result)
}

func (a sectionHasOneShortAnswerSectionTx) Append(values ...*entity.ShortAnswerSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a sectionHasOneShortAnswerSectionTx) Replace(values ...*entity.ShortAnswerSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a sectionHasOneShortAnswerSectionTx) Delete(values ...*entity.ShortAnswerSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a sectionHasOneShortAnswerSectionTx) Clear() error {
	return a.tx.Clear()
}

func (a sectionHasOneShortAnswerSectionTx) Count() int64 {
	return a.tx.Count()
}

type sectionHasOneAnswerSection struct {
	db *gorm.DB

	field.RelationField
}

func (a sectionHasOneAnswerSection) Where(conds ...field.Expr) *sectionHasOneAnswerSection {
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

func (a sectionHasOneAnswerSection) WithContext(ctx context.Context) *sectionHasOneAnswerSection {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a sectionHasOneAnswerSection) Session(session *gorm.Session) *sectionHasOneAnswerSection {
	a.db = a.db.Session(session)
	return &a
}

func (a sectionHasOneAnswerSection) Model(m *entity.Section) *sectionHasOneAnswerSectionTx {
	return &sectionHasOneAnswerSectionTx{a.db.Model(m).Association(a.Name())}
}

type sectionHasOneAnswerSectionTx struct{ tx *gorm.Association }

func (a sectionHasOneAnswerSectionTx) Find() (result *entity.AnswerSection, err error) {
	return result, a.tx.Find(&result)
}

func (a sectionHasOneAnswerSectionTx) Append(values ...*entity.AnswerSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a sectionHasOneAnswerSectionTx) Replace(values ...*entity.AnswerSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a sectionHasOneAnswerSectionTx) Delete(values ...*entity.AnswerSection) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a sectionHasOneAnswerSectionTx) Clear() error {
	return a.tx.Clear()
}

func (a sectionHasOneAnswerSectionTx) Count() int64 {
	return a.tx.Count()
}

type sectionDo struct{ gen.DO }

type ISectionDo interface {
	gen.SubQuery
	Debug() ISectionDo
	WithContext(ctx context.Context) ISectionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISectionDo
	WriteDB() ISectionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISectionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISectionDo
	Not(conds ...gen.Condition) ISectionDo
	Or(conds ...gen.Condition) ISectionDo
	Select(conds ...field.Expr) ISectionDo
	Where(conds ...gen.Condition) ISectionDo
	Order(conds ...field.Expr) ISectionDo
	Distinct(cols ...field.Expr) ISectionDo
	Omit(cols ...field.Expr) ISectionDo
	Join(table schema.Tabler, on ...field.Expr) ISectionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISectionDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISectionDo
	Group(cols ...field.Expr) ISectionDo
	Having(conds ...gen.Condition) ISectionDo
	Limit(limit int) ISectionDo
	Offset(offset int) ISectionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISectionDo
	Unscoped() ISectionDo
	Create(values ...*entity.Section) error
	CreateInBatches(values []*entity.Section, batchSize int) error
	Save(values ...*entity.Section) error
	First() (*entity.Section, error)
	Take() (*entity.Section, error)
	Last() (*entity.Section, error)
	Find() ([]*entity.Section, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Section, err error)
	FindInBatches(result *[]*entity.Section, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Section) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISectionDo
	Assign(attrs ...field.AssignExpr) ISectionDo
	Joins(fields ...field.RelationField) ISectionDo
	Preload(fields ...field.RelationField) ISectionDo
	FirstOrInit() (*entity.Section, error)
	FirstOrCreate() (*entity.Section, error)
	FindByPage(offset int, limit int) (result []*entity.Section, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISectionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sectionDo) Debug() ISectionDo {
	return s.withDO(s.DO.Debug())
}

func (s sectionDo) WithContext(ctx context.Context) ISectionDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sectionDo) ReadDB() ISectionDo {
	return s.Clauses(dbresolver.Read)
}

func (s sectionDo) WriteDB() ISectionDo {
	return s.Clauses(dbresolver.Write)
}

func (s sectionDo) Session(config *gorm.Session) ISectionDo {
	return s.withDO(s.DO.Session(config))
}

func (s sectionDo) Clauses(conds ...clause.Expression) ISectionDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sectionDo) Returning(value interface{}, columns ...string) ISectionDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sectionDo) Not(conds ...gen.Condition) ISectionDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sectionDo) Or(conds ...gen.Condition) ISectionDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sectionDo) Select(conds ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sectionDo) Where(conds ...gen.Condition) ISectionDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sectionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ISectionDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sectionDo) Order(conds ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sectionDo) Distinct(cols ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sectionDo) Omit(cols ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sectionDo) Join(table schema.Tabler, on ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sectionDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISectionDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sectionDo) RightJoin(table schema.Tabler, on ...field.Expr) ISectionDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sectionDo) Group(cols ...field.Expr) ISectionDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sectionDo) Having(conds ...gen.Condition) ISectionDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sectionDo) Limit(limit int) ISectionDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sectionDo) Offset(offset int) ISectionDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sectionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISectionDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sectionDo) Unscoped() ISectionDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sectionDo) Create(values ...*entity.Section) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sectionDo) CreateInBatches(values []*entity.Section, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sectionDo) Save(values ...*entity.Section) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sectionDo) First() (*entity.Section, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Section), nil
	}
}

func (s sectionDo) Take() (*entity.Section, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Section), nil
	}
}

func (s sectionDo) Last() (*entity.Section, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Section), nil
	}
}

func (s sectionDo) Find() ([]*entity.Section, error) {
	result, err := s.DO.Find()
	return result.([]*entity.Section), err
}

func (s sectionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Section, err error) {
	buf := make([]*entity.Section, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sectionDo) FindInBatches(result *[]*entity.Section, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sectionDo) Attrs(attrs ...field.AssignExpr) ISectionDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sectionDo) Assign(attrs ...field.AssignExpr) ISectionDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sectionDo) Joins(fields ...field.RelationField) ISectionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sectionDo) Preload(fields ...field.RelationField) ISectionDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sectionDo) FirstOrInit() (*entity.Section, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Section), nil
	}
}

func (s sectionDo) FirstOrCreate() (*entity.Section, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Section), nil
	}
}

func (s sectionDo) FindByPage(offset int, limit int) (result []*entity.Section, count int64, err error) {
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

func (s sectionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sectionDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sectionDo) Delete(models ...*entity.Section) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sectionDo) withDO(do gen.Dao) *sectionDo {
	s.DO = *do.(*gen.DO)
	return s
}
