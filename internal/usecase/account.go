package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/hash"
)

type AccountUseCase struct {
	userRepo       postgres.User
	passwordHasher hash.Hasher
	courseRepo     postgres.Course
}

func (uc AccountUseCase) GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return uc.userRepo.GetById(ctx, id)
}

func NewAccountUseCase(userRepo postgres.User, courseRepo postgres.Course, passwordHasher hash.Hasher) AccountUseCase {
	return AccountUseCase{userRepo: userRepo, courseRepo: courseRepo, passwordHasher: passwordHasher}
}

func (uc AccountUseCase) UpdateUser(ctx context.Context, update *entity.User) (*entity.User, error) {
	return uc.userRepo.Update(ctx, update)
}

func (uc AccountUseCase) UpdatePassword(ctx context.Context, userId uuid.UUID, current, new string) error {
	user, err := uc.userRepo.GetById(ctx, userId)
	if err != nil {
		return err
	}

	if !uc.passwordHasher.Equals(user.Password, current) {
		return errors.New("неверный пароль")
	}

	hashedPassword, err := uc.passwordHasher.Hash(new)
	if err != nil {
		return err
	}

	return uc.userRepo.UpdatePassword(ctx, userId, hashedPassword)
}

func (uc AccountUseCase) GetStudentCourses(ctx context.Context, studentId uuid.UUID) ([]*entity.Course, error) {
	return uc.courseRepo.GetByStudent(ctx, studentId)
}
