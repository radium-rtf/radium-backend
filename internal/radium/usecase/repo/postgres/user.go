package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"fmt"
	"github.com/google/uuid"
	entity "github.com/radium-rtf/radium-backend/internal/radium/entity"
	repoerr2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/repoerr"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type User struct {
	db             *bun.DB
	defaultGroupId uuid.UUID

	group Group
}

func NewUserRepo(pg *postgres.Postgres) User {
	return User{db: pg.DB, defaultGroupId: pg.DefaultGroupId, group: NewGroupRepo(pg)}
}

func (r User) Create(ctx context.Context, user *entity.User) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(user).Exec(ctx)
		if err != nil {
			return err
		}

		roles := entity.GetAllRoles(user.Id)
		user.Roles = roles
		_, err = tx.NewInsert().Model(roles).Exec(ctx)
		if err != nil {
			return err
		}

		return err
	})
}

func (r User) get(ctx context.Context, value columnValue) (*entity.User, error) {
	var user = new(entity.User)
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

func (r User) GetFull(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user = new(entity.User)
	err := r.db.NewSelect().Model(user).
		Relation("Contact").
		Relation("Roles").
		Relation("Author.Modules.Pages.Sections.UsersAnswers.Review").
		Relation("Author.Modules.Pages.Sections.UsersAnswers.File").
		Relation("Author.LastVisitedPage", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("user_id =?", id)
		}).
		Relation("Author.Authors").
		Relation("Coauthor.Modules.Pages.Sections.UsersAnswers.Review").
		Relation("Coauthor.LastVisitedPage", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("user_id =?", id)
		}).
		Relation("Coauthor.Modules.Pages.Sections.UsersAnswers.File").
		Relation("LastVisitedPage", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("last_visited_page.updated_at desc").Limit(1)
		}).
		Relation("LastVisitedPage.Course.Modules.Pages.Sections.UsersAnswers.Review").
		Relation("LastVisitedPage.Course.Modules.Pages.Sections.UsersAnswers.File").
		Relation("LastVisitedPage.Course.Authors").
		Relation("LastVisitedPage.Course.Coauthors").
		Relation("Courses.Modules.Pages.Sections.UsersAnswers.Review").
		Relation("Courses.Modules.Pages.Sections.UsersAnswers.File").
		Limit(1).
		Scan(ctx)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, repoerr2.NotFound
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
		Relation("Contact").
		Where("id in (?)", bun.In(ids)).
		Scan(ctx)
	if errors.Is(sql.ErrNoRows, err) {
		return nil, repoerr2.NotFound
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

func (r User) Update(ctx context.Context, user *entity.User) (*entity.User, error) {

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

func (r User) CreateUnverifiedUser(ctx context.Context, user *entity.UnverifiedUser) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r User) GetUnverifiedUser(ctx context.Context, email, verificationCode string) (*entity.UnverifiedUser, error) {
	var user = new(entity.UnverifiedUser)
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

func (r User) SaveLastVisitedPage(ctx context.Context, page *entity.Page, userId uuid.UUID) error {
	module := new(entity.Module)
	err := r.db.NewSelect().
		Model(module).
		Where("id = ?", page.ModuleId).
		Scan(ctx)
	if err != nil {
		return err
	}

	lastVisitedPage := &entity.LastVisitedPage{
		CourseId:  module.CourseId,
		UserId:    userId,
		PageId:    page.Id,
		UpdatedAt: time.Now(),
	}

	set := `course_id = excluded.course_id, page_id = excluded.page_id, 
			user_id = excluded.user_id, updated_at = excluded.updated_at`
	_, err = r.db.NewInsert().
		Model(lastVisitedPage).On("conflict (user_id, course_id) do update").
		Set(set).
		Exec(ctx)

	return err
}

func (r User) Search(ctx context.Context, query string, limit int) ([]*entity.User, error) {
	var (
		users   []*entity.User
		tsquery = fmt.Sprintf("(CONCAT(CAST('%v'::tsquery as text), ':*'))::tsquery", query)
	)

	where := `tsvector_name @@ ? or tsvector_email @@ ?`
	order := `(tsvector_name <=> ?) + (tsvector_email <=> ?)`

	err := r.db.NewSelect().
		Model(&users).
		Where(where, bun.Safe(tsquery), bun.Safe(tsquery)).
		OrderExpr(order, bun.Safe(tsquery), bun.Safe(tsquery)).
		Limit(limit).
		Scan(ctx)

	return users, err
}
