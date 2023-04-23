package repo

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type UserRepo struct {
	pg *postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) UserRepo {
	return UserRepo{pg: pg}
}

func (r UserRepo) Create(ctx context.Context, signUp entity.SignUp) error {
	sql, args, err := r.pg.Builder.
		Insert("users").
		Columns("id", "name", "email", "password").
		Values(uuid.NewString(), signUp.Name, signUp.Email, signUp.Password).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r UserRepo) get(ctx context.Context, where sq.Eq) (entity.User, error) {
	user := entity.User{}
	sql, args, err := r.pg.Builder.
		Select("id", "name", "email", "password").
		Where(where).
		Limit(1).
		From("users").
		ToSql()
	if err != nil {
		return user, err
	}
	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	if !rows.Next() {
		return user, errors.New("пользователь не найден")
	}
	err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	return r.get(ctx, sq.Eq{"email": email})
}

func (r UserRepo) GetById(ctx context.Context, id string) (entity.User, error) {
	return r.get(ctx, sq.Eq{"id": id})
}

func (r UserRepo) GetByVerificationCode(ctx context.Context, verificationCode string) (entity.User, error) {
	return r.get(ctx, sq.Eq{"verification_code": verificationCode})
}

func (r UserRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (entity.User, error) {
	var user entity.User
	sql, args, err := r.pg.Builder.
		Select("user_id", "name", "email", "password", "expires_in").
		Where(sq.Eq{"refresh_token": refreshToken}).
		Limit(1).
		From("sessions").
		Join("users ON users.id = sessions.user_id").
		ToSql()
	if err != nil {
		return user, err
	}
	rows, err := r.pg.Pool.Query(ctx, sql, args...)
	if err != nil {
		return user, err
	}
	defer rows.Close()
	var expiresIn time.Time
	if !rows.Next() {
		return user, errors.New("сессия не найдена")
	}
	err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &expiresIn)
	if err != nil {
		return user, err
	}
	if expiresIn.Before(time.Now()) {
		return user, errors.New("сессия истекла")
	}
	return user, nil
}

func (r UserRepo) updateColumn(ctx context.Context, where sq.Eq, column string, value interface{}) error {
	sql, args, err := r.pg.Builder.
		Update("users").
		Where(where).
		Set(column, value).
		ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Pool.Exec(ctx, sql, args...)
	return err
}

func (r UserRepo) SetVerificationCode(ctx context.Context, id string, code string) error {
	return r.updateColumn(ctx, sq.Eq{"id": id}, "verification_code", code)
}

func (r UserRepo) VerifyUser(ctx context.Context, id string) error {
	return r.updateColumn(ctx, sq.Eq{"id": id}, "is_verified", true)
}

func (r UserRepo) UpdateName(ctx context.Context, id string, name entity.UserName) error {
	return r.updateColumn(ctx, sq.Eq{"id": id}, "name", name.Name)
}

func (r UserRepo) UpdatePassword(ctx context.Context, id string, password string) error {
	return r.updateColumn(ctx, sq.Eq{"id": id}, "password", password)
}
