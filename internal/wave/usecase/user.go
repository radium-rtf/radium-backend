package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/wave/lib/centrifugo"
)

type UserUseCase struct {
	centrifugo centrifugo.Centrifugo
}

func (uc UserUseCase) GetClientToken(ctx context.Context) (string, error) {
	userId, ok := ctx.Value("userId").(uuid.UUID)
	if !ok {
		userId = uuid.Nil
	}
	token, err := uc.centrifugo.GetConnectionToken(userId.String(), 0)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewUserUseCase(centrifugo centrifugo.Centrifugo) UserUseCase {
	return UserUseCase{
		centrifugo: centrifugo,
	}
}
