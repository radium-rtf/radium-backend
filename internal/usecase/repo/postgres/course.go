package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/repoerr"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"
)

type Course struct {
	db *bun.DB
}

func NewCourseRepo(pg *postgres.Postgres) Course {
	return Course{db: pg.DB}
}

func (r Course) Create(ctx context.Context, course *entity.Course) (*entity.Course, error) {
	err := r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(course).Exec(ctx)
		if err != nil {
			return err
		}

		var courseAuthor []*entity.CourseAuthor
		for _, author := range course.Authors {
			courseAuthor = append(courseAuthor, &entity.CourseAuthor{CourseId: course.Id, UserId: author.Id})
		}
		_, err = tx.NewInsert().Model(&courseAuthor).Exec(ctx)
		if err != nil {
			return err
		}

		if len(course.Links) != 0 {
			_, err = tx.NewInsert().Model(&course.Links).Exec(ctx)
		}
		return err
	})
	if err != nil {
		return nil, err
	}
	return r.GetFullById(ctx, course.Id)
}

func (r Course) Get(ctx context.Context) ([]*entity.Course, error) {
	var courses []*entity.Course
	err := r.db.NewSelect().
		Model(&courses).
		Relation("Students").
		Where("is_published = ?", true).
		Scan(ctx)
	return courses, err
}

func (r Course) GetFullById(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	where := columnValue{column: "id", value: id}
	return r.getFull(ctx, where)
}

func (r Course) GetFullBySlug(ctx context.Context, slug string) (*entity.Course, error) {
	where := columnValue{column: "slug", value: slug}
	return r.getFull(ctx, where)
}

func (r Course) getFull(ctx context.Context, where columnValue) (*entity.Course, error) {
	var course = new(entity.Course)
	err := r.db.NewSelect().
		Model(course).
		Where(where.column+" = ?", where.value).
		Relation("Authors").
		Relation("Authors.Roles").
		Relation("Coauthors").
		Relation("Coauthors.Roles").
		Relation("Links").
		Relation("Students").
		Relation("Modules", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Relation("Modules.Pages", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Relation("Modules.Pages.Sections", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("order")
		}).
		Scan(ctx)
	return course, err
}

func (r Course) Join(ctx context.Context, userId, courseId uuid.UUID) error {
	student := &entity.CourseStudent{CourseId: courseId, UserId: userId}
	_, err := r.db.NewInsert().Model(student).Exec(ctx)
	return err
}

func (r Course) GetByStudent(ctx context.Context, userId uuid.UUID) ([]*entity.Course, error) {
	var user = new(entity.User)

	err := r.db.NewSelect().
		Model(user).
		Where("id = ?", userId).
		Relation("Courses", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("is_published = ?", true)
		}).Relation("Courses.Authors").Relation("Courses.Links").
		Relation("Courses.Authors.Roles").
		Relation("Courses.Students").
		Relation("Courses.Coauthors").
		Relation("Courses.Coauthors.Roles").
		Scan(ctx)

	return user.Courses, err
}

func (r Course) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	var query = r.db.NewDelete().
		Model(&entity.Course{}).
		Where("id = ?", id)
	if !isSoft {
		query = query.ForceDelete()
	}
	_, err := query.Exec(ctx)
	return err
}

func (r Course) Update(ctx context.Context, course *entity.Course) (*entity.Course, error) {
	info, err := r.db.NewUpdate().
		Model(course).
		WherePK().
		OmitZero().
		Exec(ctx)

	n, _ := info.RowsAffected()
	if err == nil && n == 0 {
		return nil, repoerr.CourseNotFound
	}
	if err != nil {
		return nil, err
	}
	return r.GetFullById(ctx, course.Id)
}

func (r Course) UpdatePublish(ctx context.Context, id uuid.UUID, status bool) (*entity.Course, error) {
	info, err := r.db.NewUpdate().
		Model(&entity.Course{}).
		Where("id = ?", id).
		Set("is_published = ?", status).
		Exec(ctx)

	n, _ := info.RowsAffected()
	if err == nil && n == 0 {
		return nil, repoerr.CourseNotFound
	}
	if err != nil {
		return nil, err
	}
	return r.GetFullById(ctx, id)
}

func (r Course) GetByAuthorId(ctx context.Context, id uuid.UUID) ([]*entity.Course, error) {
	coursesIds := r.db.NewSelect().
		Column("course_id").
		Model(&entity.CourseAuthor{}).
		Where("user_id = ?", id)

	var courses []*entity.Course
	err := r.db.NewSelect().
		Model(&courses).
		Relation("Students").
		Where("id in (?)", coursesIds).
		Scan(ctx)

	return courses, err
}

func (r Course) GetLinkById(ctx context.Context, id uuid.UUID) (*entity.Link, error) {
	var link = new(entity.Link)

	err := r.db.NewSelect().Model(link).
		Where("link.id = ?", id).
		Relation("Course").
		Relation("Course.Authors").
		Relation("Course.Coauthors").
		Scan(ctx)

	return link, err
}

func (r Course) DeleteLink(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.NewDelete().
		Model(&entity.Link{}).
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func (r Course) CreateLink(ctx context.Context, link *entity.Link) error {
	_, err := r.db.NewInsert().Model(link).Exec(ctx)
	return err
}

func (r Course) GetRecommendations(ctx context.Context, userId uuid.UUID, limit int) ([]*entity.Course, error) {
	coursesIds := r.db.NewSelect().
		Column("course_id").
		Model(&entity.CourseStudent{}).
		Where("user_id = ?", userId)

	var courses []*entity.Course
	err := r.db.NewSelect().
		Model(&courses).
		Where("id not in (?) and is_published", coursesIds).
		Limit(limit).
		Scan(ctx)

	return courses, err
}
