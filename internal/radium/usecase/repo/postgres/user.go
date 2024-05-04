package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	repoerr2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/repoerr"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type User struct {
	db             *bun.DB
	defaultGroupId uuid.UUID
}

func NewUserRepo(pg *postgres.Postgres) User {
	defaultGroupId := uuid.MustParse("81af02da-bf9e-4769-aa07-36903517733d")

	return User{db: pg.DB, defaultGroupId: defaultGroupId}
}

func (r User) Create(ctx context.Context, user *entity2.User) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(user).Exec(ctx)
		if err != nil {
			return err
		}

		roles := entity2.GetAllRoles(user.Id)
		user.Roles = roles
		_, err = tx.NewInsert().Model(roles).Exec(ctx)
		if err != nil {
			return err
		}

		studentGroup := &entity2.GroupStudent{GroupId: r.defaultGroupId, UserId: user.Id}
		_, err = tx.NewInsert().Model(studentGroup).Exec(ctx)
		return err
	})
}

func (r User) get(ctx context.Context, value columnValue) (*entity2.User, error) {
	var user = new(entity2.User)
	err := r.db.NewSelect().Model(user).
		Relation("Contact").
		Relation("Roles").
		Where(value.column+" = ?", value.value).
		Limit(1).
		Scan(ctx)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, repoerr2.NotFound
	}
	return user, err
}

func (r User) GetFull(ctx context.Context, id uuid.UUID) (*entity2.User, error) {
	var user = new(entity2.User)
	err := r.db.NewSelect().Model(user).
		Relation("Contact").
		Relation("Roles").
		Relation("Author").
		Relation("Author.Authors").
		Relation("Coauthor").
		Relation("Courses").
		Relation("Courses.Modules", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Relation("Courses.Modules.Pages", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Relation("Courses.Modules.Pages.Sections", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Relation("Courses.Modules.Pages.Sections.UsersAnswers", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("answer.created_at desc").Where("user_id = ?", id)
		}).
		Relation("Courses.Modules.Pages.Sections.UsersAnswers.Review").
		Relation("Courses.Modules.Pages.Sections.UsersAnswers.File").
		Where("id = ?", id).
		Limit(1).
		Scan(ctx)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, repoerr2.NotFound
	}
	return user, err
}

func (r User) GetByEmail(ctx context.Context, email string) (*entity2.User, error) {
	return r.get(ctx, columnValue{column: "email", value: email})
}

func (r User) GetById(ctx context.Context, id uuid.UUID) (*entity2.User, error) {
	return r.get(ctx, columnValue{column: "id", value: id})
}

func (r User) GetByIds(ctx context.Context, ids []uuid.UUID) ([]*entity2.User, error) {
	var users []*entity2.User
	err := r.db.NewSelect().Model(&users).Relation("Roles").
		Relation("Contact").
		Where("id in (?)", bun.In(ids)).
		Scan(ctx)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, repoerr2.NotFound
	}
	return users, err
}

func (r User) GetByRefreshToken(ctx context.Context, refreshToken uuid.UUID) (*entity2.User, error) {
	var session = new(entity2.Session)

	err := r.db.NewSelect().
		Model(session).
		Where("refresh_token = ?", refreshToken).
		Order("expires_in desc").
		Limit(1).
		Relation("User").
		Relation("User.Roles").
		Scan(ctx)

	if err == nil && session.ExpiresIn.Before(time.Now()) {
		return nil, repoerr2.SessionIsExpired
	}
	if errors.Is(sql.ErrNoRows, err) {
		return nil, repoerr2.NotFound
	}

	return session.User, err

}

func (r User) updateColumn(ctx context.Context, value columnValue, where columnValue) error {
	exec, err := r.db.NewUpdate().
		Table("users").
		Where(where.column+" = ?", where.value).
		Set(value.column+" = ?", value.value).Exec(ctx)
	if err != nil {
		return err
	}

	rowsAffected, _ := exec.RowsAffected()
	if rowsAffected == 0 {
		return repoerr2.NotFound
	}

	return err
}

func (r User) UpdatePassword(ctx context.Context, id uuid.UUID, password string) error {
	value := columnValue{column: "password", value: password}
	where := columnValue{column: "id", value: id}
	return r.updateColumn(ctx, value, where)
}

func (r User) Update(ctx context.Context, user *entity2.User) (*entity2.User, error) {

	err := r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

		if user.Contact != nil {
			_, err := tx.NewInsert().
				Model(user.Contact).
				On("conflict (user_id) do update").
				Set("name = EXCLUDED.name, link = EXCLUDED.link").
				Exec(ctx)
			if err != nil {
				return err
			}
		}

		exec, err := tx.NewUpdate().
			Model(user).
			WherePK().
			OmitZero().
			Exec(ctx)
		if err != nil {
			return err
		}

		rowsAffected, err := exec.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return repoerr2.NotFound
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	updatedUser, err := r.GetById(ctx, user.Id)

	return updatedUser, err
}

func (r User) CreateUnverifiedUser(ctx context.Context, user *entity2.UnverifiedUser) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r User) GetUnverifiedUser(ctx context.Context, email, verificationCode string) (*entity2.UnverifiedUser, error) {
	var user = new(entity2.UnverifiedUser)
	err := r.db.NewSelect().
		Model(user).
		Where("email = ? and verification_code = ?", email, verificationCode).
		Limit(1).
		Scan(ctx)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, repoerr2.NotFound
	}
	return user, err
}
