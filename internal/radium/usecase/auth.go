package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	entity2 "github.com/radium-rtf/radium-backend/internal/radium/entity"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/session"
	"github.com/radium-rtf/radium-backend/internal/radium/model"
	postgres2 "github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/email"
	"github.com/radium-rtf/radium-backend/pkg/hash"
	"github.com/radium-rtf/radium-backend/pkg/str"
)

type AuthUseCase struct {
	user                   postgres2.User
	sessionRepo            postgres2.Session
	hasher                 hash.Hasher
	session                session.Session
	emailSender            *email.SMTPSender
	lengthVerificationCode int
}

func NewAuthUseCase(userRepo postgres2.User, sessionRepo postgres2.Session,
	hasher hash.Hasher, session session.Session, smtp *email.SMTPSender, lengthVerificationCode int) AuthUseCase {
	return AuthUseCase{
		user:                   userRepo,
		sessionRepo:            sessionRepo,
		hasher:                 hasher,
		session:                session,
		emailSender:            smtp,
		lengthVerificationCode: lengthVerificationCode,
	}
}

func (uc AuthUseCase) SignIn(ctx context.Context, email, password string) (model.Tokens, error) {
	var tokens model.Tokens
	user, err := uc.user.GetByEmail(ctx, email)
	if err != nil {
		return tokens, err
	}
	if !uc.hasher.Equals(user.Password, password) {
		return tokens, errors.New("неверный пароль")
	}
	return uc.createSession(ctx, user)
}

func (uc AuthUseCase) SignUp(ctx context.Context, user *entity2.User) (*entity2.UnverifiedUser, error) {
	var unverifiedUser *entity2.UnverifiedUser

	_, err := uc.user.GetByEmail(ctx, user.Email)
	if err == nil {
		return nil, errors.New("пользователь с такой почтой уже существует")
	}

	password, err := uc.hasher.Hash(user.Password)
	if err != nil {
		return unverifiedUser, err
	}
	user.Password = password

	code := str.Random(uc.lengthVerificationCode)

	unverifiedUser = &entity2.UnverifiedUser{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Password: user.Password,

		ExpiresAt:        time.Now().Add(time.Hour * 100),
		VerificationCode: code,
	}

	err = uc.user.CreateUnverifiedUser(ctx, unverifiedUser)
	if err != nil {
		return unverifiedUser, err
	}
	err = uc.emailSender.SendVerificationEmail(user.Email, user.Name, code)
	return unverifiedUser, err
}

func (uc AuthUseCase) signUp(ctx context.Context, user *entity2.User) (model.Tokens, error) {
	var tokens model.Tokens
	password, err := uc.hasher.Hash(user.Password)
	if err != nil {
		return tokens, err
	}
	user.Password = password
	err = uc.user.Create(ctx, user)
	if err != nil {
		return tokens, err
	}
	user, err = uc.user.GetByEmail(ctx, user.Email)
	if err != nil {
		return tokens, err
	}

	return uc.createSession(ctx, user)
}

func (uc AuthUseCase) RefreshToken(ctx context.Context, refreshToken uuid.UUID) (model.Tokens, error) {
	user, err := uc.user.GetByRefreshToken(ctx, refreshToken)
	if err != nil {
		return model.Tokens{}, err
	}
	return uc.refreshSession(ctx, user, refreshToken)
}

func (uc AuthUseCase) createSession(ctx context.Context, user *entity2.User) (model.Tokens, error) {
	userSession, tokens, err := uc.session.Create(model.NewUser(user))
	if err != nil {
		return model.Tokens{}, err
	}

	err = uc.sessionRepo.Create(ctx, userSession)
	return tokens, err
}

func (uc AuthUseCase) refreshSession(ctx context.Context, user *entity2.User, refreshToken uuid.UUID) (model.Tokens, error) {
	tokens, refreshTime, err := uc.session.Refresh(model.NewUser(user), refreshToken)
	if err != nil {
		return model.Tokens{}, err
	}
	err = uc.sessionRepo.Update(ctx, refreshToken, refreshTime)
	return tokens, err
}

func (uc AuthUseCase) VerifyEmail(ctx context.Context, email, verificationCode string) (model.Tokens, error) {
	var tokens model.Tokens

	unverifiedUser, err := uc.user.GetUnverifiedUser(ctx, email, verificationCode)
	if err != nil {
		return tokens, err
	}
	if unverifiedUser.VerificationCode != verificationCode {
		return tokens, errors.New("неправильный код подтверждения")
	}
	if time.Now().After(unverifiedUser.ExpiresAt) {
		return tokens, errors.New("время по подтверждение истекло")
	}

	user := &entity2.User{
		DBModel:  entity2.DBModel{Id: unverifiedUser.Id},
		Name:     unverifiedUser.Name,
		Password: unverifiedUser.Password,
		Avatar:   unverifiedUser.Avatar,
		Email:    unverifiedUser.Email,
	}
	err = uc.user.Create(ctx, user)
	if err != nil {
		return tokens, err
	}

	return uc.createSession(ctx, user)
}
