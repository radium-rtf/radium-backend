package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	postgres2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/hash"
)

type AccountUseCase struct {
	userRepo       postgres2.User
	passwordHasher hash.Hasher
	courseRepo     postgres2.Course
}

func (uc AccountUseCase) GetFullUser(ctx context.Context, id uuid.UUID) (*entity2.User, error) {
	return uc.userRepo.GetFull(ctx, id)
}

func (uc AccountUseCase) GetUser(ctx context.Context, id uuid.UUID) (*entity2.User, error) {
	return uc.userRepo.GetById(ctx, id)
}

func NewAccountUseCase(userRepo postgres2.User, courseRepo postgres2.Course, passwordHasher hash.Hasher) AccountUseCase {
	return AccountUseCase{userRepo: userRepo, courseRepo: courseRepo, passwordHasher: passwordHasher}
}

func (uc AccountUseCase) UpdateUser(ctx context.Context, update *entity2.User) (*entity2.User, error) {
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

func (uc AccountUseCase) GetStudentCourses(ctx context.Context, studentId uuid.UUID) ([]*entity2.Course, error) {
	return uc.courseRepo.GetByStudent(ctx, studentId)
}

func (uc AccountUseCase) GetRecommendations(ctx context.Context, userId uuid.UUID, limit int) ([]*entity2.Course, error) {
	return uc.courseRepo.GetRecommendations(ctx, userId, limit)
}

func (uc AccountUseCase) Search(ctx context.Context, name string, limit int) ([]*entity2.User, error) {
	return uc.userRepo.Search(ctx, name, limit)
}
