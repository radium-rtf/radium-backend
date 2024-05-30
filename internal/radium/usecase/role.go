package usecase

import (
	"context"
	"errors"
	"slices"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
)

type RoleUseCase struct {
	role   postgres2.Role
	course postgres2.Course
}

func NewRoleUseCase(role postgres2.Role, course postgres2.Course) RoleUseCase {
	return RoleUseCase{role: role, course: course}
}

func (uc RoleUseCase) UpdateRole(ctx context.Context, user *entity.User) error {
	return uc.role.UpdateRole(ctx, user)
}

func (uc RoleUseCase) AddTeacher(ctx context.Context, email string) error {
	return uc.role.AddTeacher(ctx, email)
}

func (uc RoleUseCase) AddAdmin(ctx context.Context, email string) error {
	return uc.role.AddAdmin(ctx, email)
}

func (uc RoleUseCase) AddAuthor(ctx context.Context, email string) error {
	return uc.role.AddAuthor(ctx, email)
}

func (uc RoleUseCase) AddCoauthor(ctx context.Context, email string, courseId, authorId uuid.UUID) error {
	course, err := uc.course.GetFullById(ctx, courseId)
	if err != nil {
		return err
	}
	var canAddCoauthor, isAlreadyAuthor bool
	for _, author := range course.Authors {
		canAddCoauthor = canAddCoauthor || author.Id == authorId
		isAlreadyAuthor = isAlreadyAuthor || author.Email == email
	}

	if !canAddCoauthor {
		return errors.New("только автор курса может добавлять соавторов")
	}
	if isAlreadyAuthor {
		return errors.New("нельзя быть автором и соавтором одновременно")
	}

	return uc.role.AddCoauthor(ctx, email, courseId)
}

func (uc RoleUseCase) DeleteCoAuthor(ctx context.Context, id uuid.UUID, courseId uuid.UUID, deleter uuid.UUID) error {
	course, err := uc.course.GetFullById(ctx, courseId)
	if err != nil {
		return err
	}
	isAuthor := slices.ContainsFunc(course.Authors, func(user *entity.User) bool {
		return user.Id == deleter
	})
	if !isAuthor {
		return errors.New("только автор может удалять соавторов")
	}
	return uc.role.DeleteCoauthor(ctx, id, courseId)
}
