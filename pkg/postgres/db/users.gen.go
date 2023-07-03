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

func newUser(db *gorm.DB, opts ...gen.DOOption) user {
	_user := user{}

	_user.userDo.UseDB(db, opts...)
	_user.userDo.UseModel(&entity.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewAsterisk(tableName)
	_user.Id = field.NewField(tableName, "id")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")
	_user.DeletedAt = field.NewField(tableName, "deleted_at")
	_user.Avatar = field.NewString(tableName, "avatar")
	_user.Email = field.NewString(tableName, "email")
	_user.Name = field.NewString(tableName, "name")
	_user.Password = field.NewString(tableName, "password")
	_user.VerificationCode = field.NewString(tableName, "verification_code")
	_user.IsVerified = field.NewBool(tableName, "is_verified")
	_user.Sessions = userHasManySessions{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Sessions", "entity.Session"),
	}

	_user.Courses = userManyToManyCourses{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Courses", "entity.Course"),
		Links: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Courses.Links", "entity.Link"),
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
				}
			}
		}{
			RelationField: field.NewRelation("Courses.Modules", "entity.Module"),
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
				RelationField: field.NewRelation("Courses.Modules.Pages", "entity.Page"),
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
					RelationField: field.NewRelation("Courses.Modules.Pages.Sections", "entity.Section"),
					TextSection: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Courses.Modules.Pages.Sections.TextSection", "entity.TextSection"),
					},
					ChoiceSection: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Courses.Modules.Pages.Sections.ChoiceSection", "entity.ChoiceSection"),
					},
					MultiChoiceSection: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Courses.Modules.Pages.Sections.MultiChoiceSection", "entity.MultiChoiceSection"),
					},
					ShortAnswerSection: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("Courses.Modules.Pages.Sections.ShortAnswerSection", "entity.ShortAnswerSection"),
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
		}{
			RelationField: field.NewRelation("Courses.Authors", "entity.User"),
			Sessions: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Courses.Authors.Sessions", "entity.Session"),
			},
			Courses: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Courses.Authors.Courses", "entity.Course"),
			},
		},
		Students: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Courses.Students", "entity.User"),
		},
	}

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo userDo

	ALL              field.Asterisk
	Id               field.Field
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	Avatar           field.String
	Email            field.String
	Name             field.String
	Password         field.String
	VerificationCode field.String
	IsVerified       field.Bool
	Sessions         userHasManySessions

	Courses userManyToManyCourses

	fieldMap map[string]field.Expr
}

func (u user) Table(newTableName string) *user {
	u.userDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *user) updateTableName(table string) *user {
	u.ALL = field.NewAsterisk(table)
	u.Id = field.NewField(table, "id")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")
	u.Avatar = field.NewString(table, "avatar")
	u.Email = field.NewString(table, "email")
	u.Name = field.NewString(table, "name")
	u.Password = field.NewString(table, "password")
	u.VerificationCode = field.NewString(table, "verification_code")
	u.IsVerified = field.NewBool(table, "is_verified")

	u.fillFieldMap()

	return u
}

func (u *user) WithContext(ctx context.Context) IUserDo { return u.userDo.WithContext(ctx) }

func (u user) TableName() string { return u.userDo.TableName() }

func (u user) Alias() string { return u.userDo.Alias() }

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 12)
	u.fieldMap["id"] = u.Id
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
	u.fieldMap["avatar"] = u.Avatar
	u.fieldMap["email"] = u.Email
	u.fieldMap["name"] = u.Name
	u.fieldMap["password"] = u.Password
	u.fieldMap["verification_code"] = u.VerificationCode
	u.fieldMap["is_verified"] = u.IsVerified

}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u user) replaceDB(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userHasManySessions struct {
	db *gorm.DB

	field.RelationField
}

func (a userHasManySessions) Where(conds ...field.Expr) *userHasManySessions {
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

func (a userHasManySessions) WithContext(ctx context.Context) *userHasManySessions {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userHasManySessions) Session(session *gorm.Session) *userHasManySessions {
	a.db = a.db.Session(session)
	return &a
}

func (a userHasManySessions) Model(m *entity.User) *userHasManySessionsTx {
	return &userHasManySessionsTx{a.db.Model(m).Association(a.Name())}
}

type userHasManySessionsTx struct{ tx *gorm.Association }

func (a userHasManySessionsTx) Find() (result []*entity.Session, err error) {
	return result, a.tx.Find(&result)
}

func (a userHasManySessionsTx) Append(values ...*entity.Session) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userHasManySessionsTx) Replace(values ...*entity.Session) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userHasManySessionsTx) Delete(values ...*entity.Session) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userHasManySessionsTx) Clear() error {
	return a.tx.Clear()
}

func (a userHasManySessionsTx) Count() int64 {
	return a.tx.Count()
}

type userManyToManyCourses struct {
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
	}
	Students struct {
		field.RelationField
	}
}

func (a userManyToManyCourses) Where(conds ...field.Expr) *userManyToManyCourses {
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

func (a userManyToManyCourses) WithContext(ctx context.Context) *userManyToManyCourses {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a userManyToManyCourses) Session(session *gorm.Session) *userManyToManyCourses {
	a.db = a.db.Session(session)
	return &a
}

func (a userManyToManyCourses) Model(m *entity.User) *userManyToManyCoursesTx {
	return &userManyToManyCoursesTx{a.db.Model(m).Association(a.Name())}
}

type userManyToManyCoursesTx struct{ tx *gorm.Association }

func (a userManyToManyCoursesTx) Find() (result []*entity.Course, err error) {
	return result, a.tx.Find(&result)
}

func (a userManyToManyCoursesTx) Append(values ...*entity.Course) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a userManyToManyCoursesTx) Replace(values ...*entity.Course) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a userManyToManyCoursesTx) Delete(values ...*entity.Course) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a userManyToManyCoursesTx) Clear() error {
	return a.tx.Clear()
}

func (a userManyToManyCoursesTx) Count() int64 {
	return a.tx.Count()
}

type userDo struct{ gen.DO }

type IUserDo interface {
	gen.SubQuery
	Debug() IUserDo
	WithContext(ctx context.Context) IUserDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserDo
	WriteDB() IUserDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserDo
	Not(conds ...gen.Condition) IUserDo
	Or(conds ...gen.Condition) IUserDo
	Select(conds ...field.Expr) IUserDo
	Where(conds ...gen.Condition) IUserDo
	Order(conds ...field.Expr) IUserDo
	Distinct(cols ...field.Expr) IUserDo
	Omit(cols ...field.Expr) IUserDo
	Join(table schema.Tabler, on ...field.Expr) IUserDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserDo
	Group(cols ...field.Expr) IUserDo
	Having(conds ...gen.Condition) IUserDo
	Limit(limit int) IUserDo
	Offset(offset int) IUserDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo
	Unscoped() IUserDo
	Create(values ...*entity.User) error
	CreateInBatches(values []*entity.User, batchSize int) error
	Save(values ...*entity.User) error
	First() (*entity.User, error)
	Take() (*entity.User, error)
	Last() (*entity.User, error)
	Find() ([]*entity.User, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.User, err error)
	FindInBatches(result *[]*entity.User, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*entity.User) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserDo
	Assign(attrs ...field.AssignExpr) IUserDo
	Joins(fields ...field.RelationField) IUserDo
	Preload(fields ...field.RelationField) IUserDo
	FirstOrInit() (*entity.User, error)
	FirstOrCreate() (*entity.User, error)
	FindByPage(offset int, limit int) (result []*entity.User, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userDo) Debug() IUserDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) IUserDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) ReadDB() IUserDo {
	return u.Clauses(dbresolver.Read)
}

func (u userDo) WriteDB() IUserDo {
	return u.Clauses(dbresolver.Write)
}

func (u userDo) Session(config *gorm.Session) IUserDo {
	return u.withDO(u.DO.Session(config))
}

func (u userDo) Clauses(conds ...clause.Expression) IUserDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Returning(value interface{}, columns ...string) IUserDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userDo) Not(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IUserDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userDo) Order(conds ...field.Expr) IUserDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) IUserDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) IUserDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) IUserDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) IUserDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() IUserDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*entity.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*entity.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*entity.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*entity.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) Take() (*entity.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) Last() (*entity.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) Find() ([]*entity.User, error) {
	result, err := u.DO.Find()
	return result.([]*entity.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.User, err error) {
	buf := make([]*entity.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*entity.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) IUserDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userDo) Preload(fields ...field.RelationField) IUserDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userDo) FirstOrInit() (*entity.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) FirstOrCreate() (*entity.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*entity.User, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userDo) Delete(models ...*entity.User) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}
