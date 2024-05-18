package usecase

import (
	"context"

	"github.com/radium-rtf/radium-backend/internal/wave/lib/centrifugo"
)

type UserUseCase struct {
	centrifugo centrifugo.Centrifugo
}

func (uc UserUseCase) GetClientToken(ctx context.Context) (string, error) {
	token, err := uc.centrifugo.GetConnectionToken("testUser", 0)
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