package repo

import (
	"context"
	"errors"
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

func (r UserRepo) Create(ctx context.Context, signUp entity.SignUp, username string) error {
	sql, args, err := r.pg.Builder.
		Insert("users").
		Columns("name", "email", "password", "username").
		Values(signUp.Name, signUp.Email, signUp.Password, username).
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
		Select("id", "name", "email", "username", "password").
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
	err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	return r.get(ctx, sq.Eq{"email": email})
}

func (r UserRepo) GetById(ctx context.Context, userId string) (entity.User, error) {
	return r.get(ctx, sq.Eq{"id": userId})
}

func (r UserRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (entity.User, error) {
	var user entity.User
	sql, args, err := r.pg.Builder.
		Select("user_id", "name", "email", "username", "password", "expires_in").
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
	err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Username, &user.Password, &expiresIn)
	if err != nil {
		return user, err
	}
	if expiresIn.Before(time.Now()) {
		return user, errors.New("сессия истекла")
	}
	return user, nil
}

func (r UserRepo) GetByVerificationCode(ctx context.Context, verificationCode string) (entity.User, error) {
	var user entity.User
	sql, args, err := r.pg.Builder.
		Select("id").
		Where(sq.Eq{"verification_code": verificationCode}).
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
	err = rows.Scan(&user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r UserRepo) SetVerificationCode(ctx context.Context, id uint, code string) error {
	sql, args, err := r.pg.Builder.
		Update("users").
		Where(sq.Eq{"id": id}).
		Set("verification_code", code).
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

func (r UserRepo) VerifyUser(ctx context.Context, id uint) error {
	sql, args, err := r.pg.Builder.
		Update("users").
		Where(sq.Eq{"id": id}).
		Set("is_verified", true).
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
