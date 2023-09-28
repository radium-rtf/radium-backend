package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Course struct {
}

func NewCourseRepo(pg *postgres.Postgres) Course {
	return Course{}
}

func (r Course) Create(ctx context.Context, course *entity.Course) (*entity.Course, error) {
	panic("not implemented")
}

func (r Course) GetByName(ctx context.Context, name string) (*entity.Course, error) {
	panic("not implemented")
}

func (r Course) GetCourses(ctx context.Context) ([]*entity.Course, error) {
	panic("not implemented")
}

// TODO: спросить про необходимость показывать курс полностью в списке всех курсов
func (r Course) get(ctx context.Context, where ...any) ([]*entity.Course, error) {
	panic("not implemented")
}

func (r Course) GetById(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	panic("not implemented")
}

func (r Course) GetFullById(ctx context.Context, id uuid.UUID) (*entity.Course, error) {
	panic("not implemented")
}

func (r Course) GetFullBySlug(ctx context.Context, slug string) (*entity.Course, error) {
	panic("not implemented")
}

func (r Course) getFull(ctx context.Context, where ...any) (*entity.Course, error) {
	panic("not implemented")
}

func (r Course) Join(ctx context.Context, userId, courseId uuid.UUID) error {
	panic("not implemented")
}

func (r Course) GetByStudent(ctx context.Context, userId uuid.UUID) ([]*entity.Course, error) {
	panic("not implemented")
}

func (r Course) Delete(ctx context.Context, id uuid.UUID, isSoft bool) error {
	panic("not implemented")
}

func (r Course) Update(ctx context.Context, course *entity.Course) (*entity.Course, error) {
	panic("not implemented")
}
