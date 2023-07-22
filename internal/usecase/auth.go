package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/radium-rtf/radium-backend/internal/entity"
	"github.com/radium-rtf/radium-backend/internal/lib/session"
	"github.com/radium-rtf/radium-backend/internal/model"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/otp"
)

type AuthUseCase struct {
	userRepo     postgres.User
	sessionRepo  postgres.Session
	hasher       hash.Hasher
	session      session.Session
	otpGenerator otp.Generator
}

func NewAuthUseCase(userRepo postgres.User, sessionRepo postgres.Session, hasher hash.Hasher, session session.Session) AuthUseCase {
	otpGenerator := otp.NewOTPGenerator()
	return AuthUseCase{
		userRepo: userRepo, sessionRepo: sessionRepo,
		hasher: hasher, session: session,
		otpGenerator: otpGenerator,
	}
}

func (uc AuthUseCase) SignIn(ctx context.Context, email, password string) (model.Tokens, error) {
	var tokens model.Tokens
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return tokens, err
	}
	if !uc.hasher.Equals(user.Password, password) {
		return tokens, errors.New("неверный пароль")
	}
	return uc.createSession(ctx, user.Id)
}

func (uc AuthUseCase) SignUp(ctx context.Context, user *entity.User) (model.Tokens, error) {
	var tokens model.Tokens
	password, err := uc.hasher.Hash(user.Password)
	if err != nil {
		return tokens, err
	}
	user.Password = password
	err = uc.userRepo.Create(ctx, user)
	if err != nil {
		return tokens, err
	}
	user, err = uc.userRepo.GetByEmail(ctx, user.Email)
	if err != nil {
		return tokens, err
	}

	verificationCode := uc.otpGenerator.RandomSecret(16)
	uc.userRepo.SetVerificationCode(ctx, user.Id, verificationCode)
	return uc.createSession(ctx, user.Id)
}

func (uc AuthUseCase) RefreshToken(ctx context.Context, refreshToken string) (model.Tokens, error) {
	// добавить юзерайди
	user, err := uc.userRepo.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return model.Tokens{}, err
	}
	return uc.refreshSession(ctx, user.Id, refreshToken)
}

func (uc AuthUseCase) createSession(ctx context.Context, id uuid.UUID) (model.Tokens, error) {
	userSession, tokens, err := uc.session.Create(id)
	if err != nil {
		return model.Tokens{}, err
	}

	err = uc.sessionRepo.Create(ctx, userSession)
	return tokens, err
}

func (uc AuthUseCase) refreshSession(ctx context.Context, id uuid.UUID, refreshToken string) (model.Tokens, error) {
	tokens, refreshTime, err := uc.session.Refresh(id, refreshToken)
	if err != nil {
		return model.Tokens{}, err
	}
	err = uc.sessionRepo.Update(ctx, refreshToken, refreshTime)
	return tokens, err
}

func (uc AuthUseCase) VerifyEmail(ctx context.Context, verificationCode string) (bool, error) {
	user, err := uc.userRepo.GetByVerificationCode(ctx, verificationCode)
	fmt.Printf("user: %v\n", user)
	fmt.Printf("err: %v\n", err)
	if err != nil {
		return false, err
	}

	return uc.verifyUser(ctx, user.Id)
}

func (uc AuthUseCase) verifyUser(ctx context.Context, id uuid.UUID) (bool, error) {
	err := uc.userRepo.Verify(ctx, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
