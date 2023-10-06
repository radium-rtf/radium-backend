package postgres

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/repoerr"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	db *bun.DB
}

func NewUserRepo(pg *postgres.Postgres) User {
	return User{db: pg.DB}
}

func (r User) Create(ctx context.Context, user *entity.User) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(user).Exec(ctx)
		if err != nil {
			return err
		}

		roles := &entity.Roles{UserId: user.Id}
		_, err = tx.NewInsert().Model(roles).Exec(ctx)
		return err
	})
}

func (r User) get(ctx context.Context, value columnValue) (*entity.User, error) {
	var user = new(entity.User)
	err := r.db.NewSelect().Model(user).Relation("Roles").
		Where(value.column+" = ?", value.value).
		Limit(1).
		Scan(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.UserNotFound
	}
	return user, err
}

func (r User) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return r.get(ctx, columnValue{column: "email", value: email})
}

func (r User) GetById(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return r.get(ctx, columnValue{column: "id", value: id})
}

func (r User) GetByIds(ctx context.Context, ids []uuid.UUID) ([]*entity.User, error) {
	var users []*entity.User
	err := r.db.NewSelect().Model(&users).Relation("Roles").
		Where("id in (?)", bun.In(ids)).
		Scan(ctx)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.UserNotFound
	}
	return users, err
}

func (r User) GetByRefreshToken(ctx context.Context, refreshToken uuid.UUID) (*entity.User, error) {
	var session = new(entity.Session)

	err := r.db.NewSelect().
		Model(session).
		Where("refresh_token = ?", refreshToken).
		Order("expires_in desc").
		Limit(1).
		Relation("User").
		Relation("User.Roles").
		Scan(ctx)

	if err == nil && session.ExpiresIn.Before(time.Now()) {
		return nil, repoerr.SessionIsExpired
	}
	if errors.Is(pgx.ErrNoRows, err) {
		return nil, repoerr.SessionNotFound
	}

	return session.User, err

}

func (r User) updateColumn(ctx context.Context, value columnValue, where columnValue) error {
	exec, err := r.db.NewUpdate().
		Table("users").
		Where(where.column+" = ?", where.value).
		Set(value.column+" = ?", value.value).Exec(ctx)

	rowsAffected, _ := exec.RowsAffected()
	if err == nil && rowsAffected == 0 {
		return repoerr.UserNotFound
	}
	return err
}

func (r User) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	value := columnValue{column: "password", value: password}
	where := columnValue{column: "id", value: id}
	return r.updateColumn(ctx, value, where)
}

func (r User) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := r.db.NewUpdate().
		Model(user).
		WherePK().
		OmitZero().
		Exec(ctx)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.UserNotFound
	}
	if err != nil {
		return nil, err
	}
	return r.GetById(ctx, user.Id)
}
