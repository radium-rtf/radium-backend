package postgres

import (
	"context"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"

	"github.com/google/uuid"

	"github.com/radium-rtf/radium-backend/internal/entity"
)

type Group struct {
	db *bun.DB
}

func NewGroupRepo(pg *postgres.Postgres) Group {
	return Group{db: pg.DB}
}

func (r Group) Create(ctx context.Context, group *entity.Group) (*entity.Group, error) {
	err := r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().Model(group).Exec(ctx)
		if err != nil {
			return err
		}

		var groupCourse []*entity.GroupCourse
		for _, course := range group.Courses {
			groupCourse = append(groupCourse, &entity.GroupCourse{CourseId: course.Id, GroupId: group.Id})
		}

		if len(groupCourse) != 0 {
			_, err = tx.NewInsert().Model(&groupCourse).Exec(ctx)
		}
		if err != nil {
			return err
		}

		var groupStudent []*entity.GroupStudent
		for _, student := range group.Students {
			groupStudent = append(groupStudent, &entity.GroupStudent{UserId: student.Id, GroupId: group.Id})
		}
		if len(groupStudent) != 0 {
			_, err = tx.NewInsert().Model(&groupStudent).Exec(ctx)
		}
		if err != nil {
			return err
		}

		return err
	})

	return group, err
}

func (r Group) GetById(ctx context.Context, id uuid.UUID) (*entity.Group, error) {
	return r.get(ctx, columnValue{column: "id", value: id})
}

func (r Group) GetByInviteCode(ctx context.Context, code string) (*entity.Group, error) {
	return r.get(ctx, columnValue{column: "invite_code", value: code})
}

func (r Group) get(ctx context.Context, where columnValue) (*entity.Group, error) {
	var group = new(entity.Group)
	err := r.db.NewSelect().
		Model(group).
		Where(where.column+" = ?", where.value).
		Relation("Courses").Relation("Students").
		Scan(ctx)
	return group, err
}

func (r Group) JoinStudent(ctx context.Context, studentId uuid.UUID, code string) error {
	group, err := r.GetByInviteCode(ctx, code)
	if err != nil {
		return err
	}
	groupStudent := &entity.GroupStudent{UserId: studentId, GroupId: group.Id}
	_, err = r.db.NewInsert().Model(groupStudent).Exec(ctx)
	return err
}

func (r Group) Get(ctx context.Context) ([]*entity.Group, error) {
	var groups []*entity.Group
	err := r.db.NewSelect().
		Model(&groups).
		Scan(ctx)
	return groups, err
}

func (r Group) GetWithAnswers(ctx context.Context, groupId uuid.UUID, courseId uuid.UUID) (*entity.Group, error) {
	var group = new(entity.Group)
	sectionsIds := r.db.NewSelect().
		Model(&entity.Course{}).
		ColumnExpr("sections.id").
		Join("join modules on course.id = modules.course_id").
		Join("join pages on pages.module_id = modules.id").
		Join("join sections on sections.page_id = pages.id").
		Where("course.id = ?", courseId)

	students := r.db.NewSelect().
		ColumnExpr("user_id").
		Model(&entity.CourseStudent{}).
		Where("course_id = ?", courseId)

	err := r.db.NewSelect().
		Model(group).
		Where("id = ?", groupId).
		Relation("Courses", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("course.id = ?", courseId)
		}).
		Relation("Students", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("group_student.user_id in (?)", students)
		}).
		Relation("Students.Answers", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.
				Where("answer.section_id in (?) and answer.verdict in ('WAIT', 'REVIEWED')", sectionsIds).
				Order("answer.created_at desc")
		}).
		Relation("Students.Answers.Section").
		Relation("Students.Answers.Review", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Order("review.created_at desc")
		}).
		Relation("Students.Answers.File").
		Limit(1).
		Scan(ctx)

	return group, err
}
