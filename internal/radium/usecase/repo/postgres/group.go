package postgres

import (
	"context"

	entity "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
	"github.com/uptrace/bun"

	"github.com/google/uuid"
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

		var groupStudent []*entity.Student
		for _, student := range group.Students {
			for _, course := range group.Courses {
				s := &entity.Student{UserId: student.Id, GroupId: group.Id, CourseId: course.Id}
				groupStudent = append(groupStudent, s)
			}
		}

		if len(groupStudent) != 0 {
			_, err = tx.NewInsert().Model(&groupStudent).
				On("CONFLICT (user_id, course_id) DO UPDATE").
				Set("group_id = ?", group.Id).
				Exec(ctx)
		}

		return err
	})

	return group, err
}

func (r Group) Update(ctx context.Context, group *entity.Group) (*entity.Group, error) {
	err := r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewUpdate().Model(group).WherePK().OmitZero().Exec(ctx)

		return err
	})

	return group, err
}

func (r Group) AddStudent(ctx context.Context, groupId uuid.UUID, studentId uuid.UUID) (*entity.Group, error) {
	group, err := r.get(ctx, columnValue{column: "id", value: groupId})

	//err = r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
	//	_, err := tx.NewUpdate().Model(group).WherePK().OmitZero().Exec(ctx)

	//	return err
	//})

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

	insert := `
	insert into students (user_id, course_id, group_id) 
	select ? as user_id, course_id, group_id from group_course 
	where group_id = ?
	on conflict (user_id, course_id) do update set group_id = excluded.group_id
` // todo: хотелось бы делать этот разпрос через орм, но я не нашел как это написать через orm

	_, err = r.db.NewRaw(insert, studentId, group.Id).Exec(ctx)

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
	sectionsIds := r.db.NewSelect().
		Model(&entity.Course{}).
		ColumnExpr("sections.id").
		Join("join modules on course.id = modules.course_id").
		Join("join pages on pages.module_id = modules.id").
		Join("join sections on sections.page_id = pages.id").
		Where("course.id = ?", courseId)

	students := r.db.NewSelect().
		ColumnExpr("user_id").
		Model(&entity.Student{}).
		Where("course_id = ? and group_id = ?", courseId, groupId)

	var group = new(entity.Group)
	err := r.db.NewSelect().
		Model(group).
		Where("id = ?", groupId).
		Relation("Courses", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("course.id = ?", courseId)
		}).
		Relation("Students", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.Where("student.user_id in (?) and course_id = ?", students, courseId)
		}).
		Relation("Students.Answers", func(query *bun.SelectQuery) *bun.SelectQuery {
			return query.
				Where("answer.section_id in (?) and answer.verdict in ('WAIT', 'REVIEWED')", sectionsIds).
				OrderExpr("answer.created_at desc, answer.verdict = 'REVIEWED'")
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
