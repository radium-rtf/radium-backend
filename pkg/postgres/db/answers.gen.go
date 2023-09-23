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

func newAnswer(db *gorm.DB, opts ...gen.DOOption) answer {
	_answer := answer{}

	_answer.answerDo.UseDB(db, opts...)
	_answer.answerDo.UseModel(&entity.Answer{})

	tableName := _answer.answerDo.TableName()
	_answer.ALL = field.NewAsterisk(tableName)
	_answer.Id = field.NewField(tableName, "id")
	_answer.CreatedAt = field.NewTime(tableName, "created_at")
	_answer.UpdatedAt = field.NewTime(tableName, "updated_at")
	_answer.DeletedAt = field.NewField(tableName, "deleted_at")
	_answer.Verdict = field.NewString(tableName, "verdict")
	_answer.UserId = field.NewField(tableName, "user_id")
	_answer.SectionId = field.NewField(tableName, "section_id")
	_answer.Choice = answerHasOneChoice{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Choice", "entity.ChoiceSectionAnswer"),
	}

	_answer.MultiChoice = answerHasOneMultiChoice{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("MultiChoice", "entity.MultichoiceSectionAnswer"),
	}

	_answer.ShortAnswer = answerHasOneShortAnswer{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("ShortAnswer", "entity.ShortAnswerSectionAnswer"),
	}

	_answer.Answer = answerHasOneAnswer{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Answer", "entity.AnswerSectionAnswer"),
	}

	_answer.Code = answerHasOneCode{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Code", "entity.CodeSectionAnswer"),
	}

	_answer.Permutation = answerHasOnePermutation{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Permutation", "entity.PermutationSectionAnswer"),
	}

	_answer.Review = answerHasOneReview{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Review", "entity.Review"),
	}

	_answer.fillFieldMap()

	return _answer
}

type answer struct {
	answerDo answerDo

	ALL       field.Asterisk
	Id        field.Field
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Verdict   field.String
	UserId    field.Field
	SectionId field.Field
	Choice    answerHasOneChoice

	MultiChoice answerHasOneMultiChoice

	ShortAnswer answerHasOneShortAnswer

	Answer answerHasOneAnswer

	Code answerHasOneCode

	Permutation answerHasOnePermutation

	Review answerHasOneReview

	fieldMap map[string]field.Expr
}

func (a answer) Table(newTableName string) *answer {
	a.answerDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a answer) As(alias string) *answer {
	a.answerDo.DO = *(a.answerDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *answer) updateTableName(table string) *answer {
	a.ALL = field.NewAsterisk(table)
	a.Id = field.NewField(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Verdict = field.NewString(table, "verdict")
	a.UserId = field.NewField(table, "user_id")
	a.SectionId = field.NewField(table, "section_id")

	a.fillFieldMap()

	return a
}

func (a *answer) WithContext(ctx context.Context) IAnswerDo { return a.answerDo.WithContext(ctx) }

func (a answer) TableName() string { return a.answerDo.TableName() }

func (a answer) Alias() string { return a.answerDo.Alias() }

func (a answer) Columns(cols ...field.Expr) gen.Columns { return a.answerDo.Columns(cols...) }

func (a *answer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *answer) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 14)
	a.fieldMap["id"] = a.Id
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["verdict"] = a.Verdict
	a.fieldMap["user_id"] = a.UserId
	a.fieldMap["section_id"] = a.SectionId

}

func (a answer) clone(db *gorm.DB) answer {
	a.answerDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a answer) replaceDB(db *gorm.DB) answer {
	a.answerDo.ReplaceDB(db)
	return a
}

type answerHasOneChoice struct {
	db *gorm.DB

	field.RelationField
}

func (a answerHasOneChoice) Where(conds ...field.Expr) *answerHasOneChoice {
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

func (a answerHasOneChoice) WithContext(ctx context.Context) *answerHasOneChoice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a answerHasOneChoice) Session(session *gorm.Session) *answerHasOneChoice {
	a.db = a.db.Session(session)
	return &a
}

func (a answerHasOneChoice) Model(m *entity.Answer) *answerHasOneChoiceTx {
	return &answerHasOneChoiceTx{a.db.Model(m).Association(a.Name())}
}

type answerHasOneChoiceTx struct{ tx *gorm.Association }

func (a answerHasOneChoiceTx) Find() (result *entity.ChoiceSectionAnswer, err error) {
	return result, a.tx.Find(&result)
}

func (a answerHasOneChoiceTx) Append(values ...*entity.ChoiceSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a answerHasOneChoiceTx) Replace(values ...*entity.ChoiceSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a answerHasOneChoiceTx) Delete(values ...*entity.ChoiceSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a answerHasOneChoiceTx) Clear() error {
	return a.tx.Clear()
}

func (a answerHasOneChoiceTx) Count() int64 {
	return a.tx.Count()
}

type answerHasOneMultiChoice struct {
	db *gorm.DB

	field.RelationField
}

func (a answerHasOneMultiChoice) Where(conds ...field.Expr) *answerHasOneMultiChoice {
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

func (a answerHasOneMultiChoice) WithContext(ctx context.Context) *answerHasOneMultiChoice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a answerHasOneMultiChoice) Session(session *gorm.Session) *answerHasOneMultiChoice {
	a.db = a.db.Session(session)
	return &a
}

func (a answerHasOneMultiChoice) Model(m *entity.Answer) *answerHasOneMultiChoiceTx {
	return &answerHasOneMultiChoiceTx{a.db.Model(m).Association(a.Name())}
}

type answerHasOneMultiChoiceTx struct{ tx *gorm.Association }

func (a answerHasOneMultiChoiceTx) Find() (result *entity.MultichoiceSectionAnswer, err error) {
	return result, a.tx.Find(&result)
}

func (a answerHasOneMultiChoiceTx) Append(values ...*entity.MultichoiceSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a answerHasOneMultiChoiceTx) Replace(values ...*entity.MultichoiceSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a answerHasOneMultiChoiceTx) Delete(values ...*entity.MultichoiceSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a answerHasOneMultiChoiceTx) Clear() error {
	return a.tx.Clear()
}

func (a answerHasOneMultiChoiceTx) Count() int64 {
	return a.tx.Count()
}

type answerHasOneShortAnswer struct {
	db *gorm.DB

	field.RelationField
}

func (a answerHasOneShortAnswer) Where(conds ...field.Expr) *answerHasOneShortAnswer {
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

func (a answerHasOneShortAnswer) WithContext(ctx context.Context) *answerHasOneShortAnswer {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a answerHasOneShortAnswer) Session(session *gorm.Session) *answerHasOneShortAnswer {
	a.db = a.db.Session(session)
	return &a
}

func (a answerHasOneShortAnswer) Model(m *entity.Answer) *answerHasOneShortAnswerTx {
	return &answerHasOneShortAnswerTx{a.db.Model(m).Association(a.Name())}
}

type answerHasOneShortAnswerTx struct{ tx *gorm.Association }

func (a answerHasOneShortAnswerTx) Find() (result *entity.ShortAnswerSectionAnswer, err error) {
	return result, a.tx.Find(&result)
}

func (a answerHasOneShortAnswerTx) Append(values ...*entity.ShortAnswerSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a answerHasOneShortAnswerTx) Replace(values ...*entity.ShortAnswerSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a answerHasOneShortAnswerTx) Delete(values ...*entity.ShortAnswerSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a answerHasOneShortAnswerTx) Clear() error {
	return a.tx.Clear()
}

func (a answerHasOneShortAnswerTx) Count() int64 {
	return a.tx.Count()
}

type answerHasOneAnswer struct {
	db *gorm.DB

	field.RelationField
}

func (a answerHasOneAnswer) Where(conds ...field.Expr) *answerHasOneAnswer {
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

func (a answerHasOneAnswer) WithContext(ctx context.Context) *answerHasOneAnswer {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a answerHasOneAnswer) Session(session *gorm.Session) *answerHasOneAnswer {
	a.db = a.db.Session(session)
	return &a
}

func (a answerHasOneAnswer) Model(m *entity.Answer) *answerHasOneAnswerTx {
	return &answerHasOneAnswerTx{a.db.Model(m).Association(a.Name())}
}

type answerHasOneAnswerTx struct{ tx *gorm.Association }

func (a answerHasOneAnswerTx) Find() (result *entity.AnswerSectionAnswer, err error) {
	return result, a.tx.Find(&result)
}

func (a answerHasOneAnswerTx) Append(values ...*entity.AnswerSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a answerHasOneAnswerTx) Replace(values ...*entity.AnswerSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a answerHasOneAnswerTx) Delete(values ...*entity.AnswerSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a answerHasOneAnswerTx) Clear() error {
	return a.tx.Clear()
}

func (a answerHasOneAnswerTx) Count() int64 {
	return a.tx.Count()
}

type answerHasOneCode struct {
	db *gorm.DB

	field.RelationField
}

func (a answerHasOneCode) Where(conds ...field.Expr) *answerHasOneCode {
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

func (a answerHasOneCode) WithContext(ctx context.Context) *answerHasOneCode {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a answerHasOneCode) Session(session *gorm.Session) *answerHasOneCode {
	a.db = a.db.Session(session)
	return &a
}

func (a answerHasOneCode) Model(m *entity.Answer) *answerHasOneCodeTx {
	return &answerHasOneCodeTx{a.db.Model(m).Association(a.Name())}
}

type answerHasOneCodeTx struct{ tx *gorm.Association }

func (a answerHasOneCodeTx) Find() (result *entity.CodeSectionAnswer, err error) {
	return result, a.tx.Find(&result)
}

func (a answerHasOneCodeTx) Append(values ...*entity.CodeSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a answerHasOneCodeTx) Replace(values ...*entity.CodeSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a answerHasOneCodeTx) Delete(values ...*entity.CodeSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a answerHasOneCodeTx) Clear() error {
	return a.tx.Clear()
}

func (a answerHasOneCodeTx) Count() int64 {
	return a.tx.Count()
}

type answerHasOnePermutation struct {
	db *gorm.DB

	field.RelationField
}

func (a answerHasOnePermutation) Where(conds ...field.Expr) *answerHasOnePermutation {
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

func (a answerHasOnePermutation) WithContext(ctx context.Context) *answerHasOnePermutation {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a answerHasOnePermutation) Session(session *gorm.Session) *answerHasOnePermutation {
	a.db = a.db.Session(session)
	return &a
}

func (a answerHasOnePermutation) Model(m *entity.Answer) *answerHasOnePermutationTx {
	return &answerHasOnePermutationTx{a.db.Model(m).Association(a.Name())}
}

type answerHasOnePermutationTx struct{ tx *gorm.Association }

func (a answerHasOnePermutationTx) Find() (result *entity.PermutationSectionAnswer, err error) {
	return result, a.tx.Find(&result)
}

func (a answerHasOnePermutationTx) Append(values ...*entity.PermutationSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a answerHasOnePermutationTx) Replace(values ...*entity.PermutationSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a answerHasOnePermutationTx) Delete(values ...*entity.PermutationSectionAnswer) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a answerHasOnePermutationTx) Clear() error {
	return a.tx.Clear()
}

func (a answerHasOnePermutationTx) Count() int64 {
	return a.tx.Count()
}

type answerHasOneReview struct {
	db *gorm.DB

	field.RelationField
}

func (a answerHasOneReview) Where(conds ...field.Expr) *answerHasOneReview {
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

func (a answerHasOneReview) WithContext(ctx context.Context) *answerHasOneReview {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a answerHasOneReview) Session(session *gorm.Session) *answerHasOneReview {
	a.db = a.db.Session(session)
	return &a
}

func (a answerHasOneReview) Model(m *entity.Answer) *answerHasOneReviewTx {
	return &answerHasOneReviewTx{a.db.Model(m).Association(a.Name())}
}

type answerHasOneReviewTx struct{ tx *gorm.Association }

func (a answerHasOneReviewTx) Find() (result *entity.Review, err error) {
	return result, a.tx.Find(&result)
}

func (a answerHasOneReviewTx) Append(values ...*entity.Review) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a answerHasOneReviewTx) Replace(values ...*entity.Review) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a answerHasOneReviewTx) Delete(values ...*entity.Review) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a answerHasOneReviewTx) Clear() error {
	return a.tx.Clear()
}

func (a answerHasOneReviewTx) Count() int64 {
	return a.tx.Count()
}

type answerDo struct{ gen.DO }

type IAnswerDo interface {
	gen.SubQuery
	Debug() IAnswerDo
	WithContext(ctx context.Context) IAnswerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAnswerDo
	WriteDB() IAnswerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAnswerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAnswerDo
	Not(conds ...gen.Condition) IAnswerDo
	Or(conds ...gen.Condition) IAnswerDo
	Select(conds ...field.Expr) IAnswerDo
	Where(conds ...gen.Condition) IAnswerDo
	Order(conds ...field.Expr) IAnswerDo
	Distinct(cols ...field.Expr) IAnswerDo
	Omit(cols ...field.Expr) IAnswerDo
	Join(table schema.Tabler, on ...field.Expr) IAnswerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAnswerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAnswerDo
	Group(cols ...field.Expr) IAnswerDo
	Having(conds ...gen.Condition) IAnswerDo
	Limit(limit int) IAnswerDo
	Offset(offset int) IAnswerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAnswerDo
	Unscoped() IAnswerDo
	Create(values ...*entity.Answer) error
	CreateInBatches(values []*entity.Answer, batchSize int) error
	Save(values ...*entity.Answer) error
	First() (*entity.Answer, error)
	Take() (*entity.Answer, error)
	Last() (*entity.Answer, error)
	Find() ([]*entity.Answer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Answer, err error)
	FindInBatches(result *[]*entity.Answer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.Answer) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAnswerDo
	Assign(attrs ...field.AssignExpr) IAnswerDo
	Joins(fields ...field.RelationField) IAnswerDo
	Preload(fields ...field.RelationField) IAnswerDo
	FirstOrInit() (*entity.Answer, error)
	FirstOrCreate() (*entity.Answer, error)
	FindByPage(offset int, limit int) (result []*entity.Answer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAnswerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a answerDo) Debug() IAnswerDo {
	return a.withDO(a.DO.Debug())
}

func (a answerDo) WithContext(ctx context.Context) IAnswerDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a answerDo) ReadDB() IAnswerDo {
	return a.Clauses(dbresolver.Read)
}

func (a answerDo) WriteDB() IAnswerDo {
	return a.Clauses(dbresolver.Write)
}

func (a answerDo) Session(config *gorm.Session) IAnswerDo {
	return a.withDO(a.DO.Session(config))
}

func (a answerDo) Clauses(conds ...clause.Expression) IAnswerDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a answerDo) Returning(value interface{}, columns ...string) IAnswerDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a answerDo) Not(conds ...gen.Condition) IAnswerDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a answerDo) Or(conds ...gen.Condition) IAnswerDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a answerDo) Select(conds ...field.Expr) IAnswerDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a answerDo) Where(conds ...gen.Condition) IAnswerDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a answerDo) Order(conds ...field.Expr) IAnswerDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a answerDo) Distinct(cols ...field.Expr) IAnswerDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a answerDo) Omit(cols ...field.Expr) IAnswerDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a answerDo) Join(table schema.Tabler, on ...field.Expr) IAnswerDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a answerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAnswerDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a answerDo) RightJoin(table schema.Tabler, on ...field.Expr) IAnswerDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a answerDo) Group(cols ...field.Expr) IAnswerDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a answerDo) Having(conds ...gen.Condition) IAnswerDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a answerDo) Limit(limit int) IAnswerDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a answerDo) Offset(offset int) IAnswerDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a answerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAnswerDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a answerDo) Unscoped() IAnswerDo {
	return a.withDO(a.DO.Unscoped())
}

func (a answerDo) Create(values ...*entity.Answer) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a answerDo) CreateInBatches(values []*entity.Answer, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a answerDo) Save(values ...*entity.Answer) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a answerDo) First() (*entity.Answer, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Answer), nil
	}
}

func (a answerDo) Take() (*entity.Answer, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Answer), nil
	}
}

func (a answerDo) Last() (*entity.Answer, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Answer), nil
	}
}

func (a answerDo) Find() ([]*entity.Answer, error) {
	result, err := a.DO.Find()
	return result.([]*entity.Answer), err
}

func (a answerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.Answer, err error) {
	buf := make([]*entity.Answer, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a answerDo) FindInBatches(result *[]*entity.Answer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a answerDo) Attrs(attrs ...field.AssignExpr) IAnswerDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a answerDo) Assign(attrs ...field.AssignExpr) IAnswerDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a answerDo) Joins(fields ...field.RelationField) IAnswerDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a answerDo) Preload(fields ...field.RelationField) IAnswerDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a answerDo) FirstOrInit() (*entity.Answer, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Answer), nil
	}
}

func (a answerDo) FirstOrCreate() (*entity.Answer, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.Answer), nil
	}
}

func (a answerDo) FindByPage(offset int, limit int) (result []*entity.Answer, count int64, err error) {
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

func (a answerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a answerDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a answerDo) Delete(models ...*entity.Answer) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *answerDo) withDO(do gen.Dao) *answerDo {
	a.DO = *do.(*gen.DO)
	return a
}
