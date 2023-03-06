package usecase

import (
	"context"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo"
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type AccountUseCase struct {
	userRepo repo.UserRepo
}

func NewAccountUseCase(pg *postgres.Postgres) AccountUseCase {
	return AccountUseCase{userRepo: repo.NewUserRepo(pg)}
}

func (uc AccountUseCase) Account(ctx context.Context, userId string) (entity.UserDto, error) {
	user, err := uc.userRepo.GetById(ctx, userId)
	if err != nil {
		return entity.UserDto{}, err
	}
	return entity.NewUserDto(user), nil
}
