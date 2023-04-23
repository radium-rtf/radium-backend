package usecase

import (
	"context"
	"errors"
	"github.com/radium-rtf/radium-backend/config"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type AccountUseCase struct {
	userRepo       repo.UserRepo
	passwordHasher hash.PasswordHasher
	courseRepo     repo.CourseRepo
}

func NewAccountUseCase(pg *postgres.Postgres, cfg *config.Config) AccountUseCase {
	passwordHasher := hash.NewSHA1Hasher(cfg.PasswordSalt)
	return AccountUseCase{userRepo: repo.NewUserRepo(pg), courseRepo: repo.NewCourseRepo(pg), passwordHasher: passwordHasher}
}

func (uc AccountUseCase) Account(ctx context.Context, userId string) (entity.UserDto, error) {
	user, err := uc.userRepo.GetById(ctx, userId)
	if err != nil {
		return entity.UserDto{}, err
	}
	return entity.NewUserDto(user), nil
}

func (uc AccountUseCase) UpdateName(ctx context.Context, userid string, newName entity.UserName) error {
	return uc.userRepo.UpdateName(ctx, userid, newName)
}

func (uc AccountUseCase) UpdatePassword(ctx context.Context, userid string, password entity.PasswordUpdate) error {
	user, err := uc.userRepo.GetById(ctx, userid)
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
	return uc.userRepo.UpdatePassword(ctx, userid, hashedPassword)
}

func (uc AccountUseCase) GetStudentCourses(ctx context.Context, studentId string) ([]entity.Course, error) {
	return uc.courseRepo.GetByStudent(ctx, studentId)
}
