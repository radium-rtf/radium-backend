package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
)

type AccountUseCase struct {
	userRepo       repo.UserRepo
	passwordHasher hash.PasswordHasher
	courseRepo     repo.CourseRepo
}

func NewAccountUseCase(pg *db.Query, cfg *config.Config) AccountUseCase {
	passwordHasher := hash.NewSHA1Hasher(cfg.PasswordSalt)
	return AccountUseCase{userRepo: repo.NewUserRepo(pg), courseRepo: repo.NewCourseRepo(pg), passwordHasher: passwordHasher}
}

func (uc AccountUseCase) Account(ctx context.Context, userId string) (entity.UserDto, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return entity.UserDto{}, err
	}
	fmt.Print(uid)
	user, err := uc.userRepo.GetById(ctx, uid)
	if err != nil {
		return entity.UserDto{}, err
	}
	return entity.NewUserDto(user), nil
}

func (uc AccountUseCase) UpdateUser(ctx context.Context, userId string, update entity.UpdateUserRequest) (entity.UserDto, error) {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return entity.UserDto{}, err
	}
	result, err := uc.userRepo.UpdateUser(ctx, uid, update)
	if err != nil {
		return entity.UserDto{}, err
	}
	return entity.NewUserDto(result), nil
}

func (uc AccountUseCase) UpdatePassword(ctx context.Context, userId string, password entity.PasswordUpdate) error {
	uid, err := uuid.Parse(userId)
	if err != nil {
		return err
	}
	user, err := uc.userRepo.GetById(ctx, uid)
	if err != nil {
		return err
	}
	if !uc.passwordHasher.Equals(user.Password, password.Current) {
		return errors.New("неверный пароль")
	}
	hashedPassword, err := uc.passwordHasher.Hash(password.New)
	if err != nil {
		return err
	}
	return uc.userRepo.UpdatePassword(ctx, uid, hashedPassword)
}

func (uc AccountUseCase) GetStudentCourses(ctx context.Context, studentId string) ([]*entity.Course, error) {
	uid, err := uuid.Parse(studentId)
	if err != nil {
		return nil, err
	}

	return uc.courseRepo.GetByStudent(ctx, uid)
}
