package postgres

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

type User struct {
	pg *db.Query
}

func NewUserRepo(pg *db.Query) User {
	return User{pg: pg}
}

func (r User) Create(ctx context.Context, user *entity.User) error {
	u := r.pg.User
	err := r.pg.User.WithContext(ctx).Select(u.Email, u.Password, u.Name, u.Id).Create(user)

	if err != nil {
		return err
	}

	return nil
}

func (r User) get(ctx context.Context, cond ...gen.Condition) (*entity.User, error) {
	user, err := r.pg.User.WithContext(ctx).Where(cond...).First()
	if err != nil {
		return &entity.User{}, err
	}

	return user, nil
}

func (r User) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return r.get(ctx, r.pg.User.Email.Eq(email))
}

func (r User) GetById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return r.get(ctx, r.pg.User.Id.Eq(id))
}

func (r User) GetByVerificationCode(ctx context.Context, verificationCode string) (*entity.User, error) {
	return r.get(ctx, r.pg.User.VerificationCode.Eq(verificationCode))
}

func (r User) GetByRefreshToken(ctx context.Context, refreshToken string) (*entity.User, error) {
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

func (r User) updateColumn(ctx context.Context, where gen.Condition, column field.Expr, value interface{}) error {
	_, err := r.pg.User.WithContext(ctx).Where(where).Update(column, value)
	if err != nil {
		return err
	}
	return err
}

func (r User) SetVerificationCode(ctx context.Context, id uuid.UUID, code string) error {
	return r.updateColumn(ctx, r.pg.User.Id.Eq(id), r.pg.User.VerificationCode, code)
}

func (r User) Verify(ctx context.Context, id uuid.UUID) error {
	return r.updateColumn(ctx, r.pg.User.Id.Eq(id), r.pg.User.IsVerified, true)
}

func (r User) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	return r.updateColumn(ctx, r.pg.User.Id.Eq(id), r.pg.User.Password, password)
}

func (r User) Update(ctx context.Context, update *entity.User) (*entity.User, error) {
	m := structs.Map(update)
	utils.RemoveEmptyMapFields(m)

	info, err := r.pg.User.WithContext(ctx).Debug().Where(r.pg.User.Id.Eq(update.Id)).Updates(m)
	if err != nil {
		return nil, err
	}
	if info.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return r.GetById(ctx, update.Id)
}

func (r User) Delete(ctx context.Context, destroy *entity.Destroy) error {
	s := r.pg.User.WithContext(ctx)
	if !destroy.IsSoft {
		s = s.Unscoped()
	}
	info, err := s.Where(r.pg.User.Id.Eq(destroy.Id)).Delete()
	if err == nil && info.RowsAffected == 0 {
		return errors.New("not found")
	}
	return err
}
