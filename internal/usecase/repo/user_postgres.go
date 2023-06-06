package repo

import (
	"context"
	"errors"
	"time"

	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
	"github.com/radium-rtf/radium-backend/pkg/utils"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type UserRepo struct {
	pg *db.Query
}

func NewUserRepo(pg *db.Query) UserRepo {
	return UserRepo{pg: pg}
}

func (r UserRepo) Create(ctx context.Context, signUp entity.SignUp) error {
	u := r.pg.User
	user := entity.User{Email: signUp.Email, Password: signUp.Password, Name: signUp.Name}
	err := r.pg.User.WithContext(ctx).Select(u.Email, u.Password, u.Name, u.Id).Create(&user)

	if err != nil {
		return err
	}

	return nil
}

func (r UserRepo) get(ctx context.Context, cond ...gen.Condition) (*entity.User, error) {
	user, err := r.pg.User.WithContext(ctx).Where(cond...).First()
	if err != nil {
		return &entity.User{}, err
	}

	return user, nil
}

func (r UserRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return r.get(ctx, r.pg.User.Email.Eq(email))
}

func (r UserRepo) GetById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return r.get(ctx, r.pg.User.Id.Eq(id))
}

func (r UserRepo) GetByVerificationCode(ctx context.Context, verificationCode string) (*entity.User, error) {
	return r.get(ctx, r.pg.User.VerificationCode.Eq(verificationCode))
}

func (r UserRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (*entity.User, error) {
	u := r.pg.User
	s := r.pg.Session
	user, err := u.WithContext(ctx).Preload(u.Sessions.On(s.RefreshToken.Eq(refreshToken))).Take()
	if err != nil {
		return nil, err
	}

	if user.Sessions[0].ExpiresIn.Before(time.Now()) {
		return user, errors.New("сессия истекла")
	}

	return user, nil
}

func (r UserRepo) updateColumn(ctx context.Context, where gen.Condition, column field.Expr, value interface{}) error {
	_, err := r.pg.User.WithContext(ctx).Where(where).Update(column, value)
	if err != nil {
		return err
	}
	return err
}

func (r UserRepo) SetVerificationCode(ctx context.Context, id uuid.UUID, code string) error {
	return r.updateColumn(ctx, r.pg.User.Id.Eq(id), r.pg.User.VerificationCode, code)
}

func (r UserRepo) VerifyUser(ctx context.Context, id uuid.UUID) error {
	return r.updateColumn(ctx, r.pg.User.Id.Eq(id), r.pg.User.IsVerified, true)
}

func (r UserRepo) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	return r.updateColumn(ctx, r.pg.User.Id.Eq(id), r.pg.User.Password, password)
}

func (r UserRepo) UpdateUser(ctx context.Context, id uuid.UUID, update entity.UpdateUserRequest) (*entity.User, error) {
	m := structs.Map(update)
	utils.RemoveEmptyMapFields(m)

	r.pg.User.WithContext(ctx).Debug().Select(r.pg.User.Name, r.pg.User.Avatar).Where(r.pg.User.Id.Eq(id)).Updates(m)
	return r.GetById(ctx, id)
}
