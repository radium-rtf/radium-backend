package usecase

import (
	"github.com/radium-rtf/radium-backend/internal/radium/lib/auth"
	"github.com/radium-rtf/radium-backend/internal/radium/lib/session"
	"github.com/radium-rtf/radium-backend/internal/radium/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/email"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/hash"
)

type Dependencies struct {
	Repos   postgres.Repositories
	Storage filestorage.Storage

	PasswordHasher         hash.Hasher
	TokenManager           auth.TokenManager
	Session                session.Session
	Smtp                   *email.SMTPSender
	LengthVerificationCode int
}

type UseCases struct {
	Account AccountUseCase
	Answer  AnswerUseCase
	Auth    AuthUseCase
	Course  CourseUseCase
	File    FileUseCase
	Group   GroupUseCase
	Module  ModuleUseCase
	Page    PageUseCase
	Review  ReviewUseCase
	Section SectionUseCase
	Teacher TeacherUseCase
	Author  AuthorUseCase
	Role    RoleUseCase

	Deps Dependencies
}

func NewUseCases(deps Dependencies) UseCases {
	repos := deps.Repos

	return UseCases{
		Deps: deps,

		Account: NewAccountUseCase(repos.User, repos.Course, deps.PasswordHasher),
		Answer:  NewAnswerUseCase(repos.Section, repos.Answer, repos.File),
		Auth:    NewAuthUseCase(repos.User, repos.Session, deps.PasswordHasher, deps.Session, deps.Smtp, deps.LengthVerificationCode),
		Course:  NewCourseUseCase(repos.Course, repos.User),
		File:    NewFileUseCase(deps.Storage, repos.File),
		Group:   NewGroupUseCase(repos.Group, repos.Course, repos.Answer, repos.Teacher),
		Module:  NewModuleUseCase(repos.Module, repos.Course),
		Page:    NewPageUseCase(repos.Page, repos.Module),
		Review:  NewReviewUseCase(repos.Review, repos.Section),
		Section: NewSectionUseCase(repos.Section, repos.Page),
		Teacher: NewTeacherUseCase(repos.Teacher),
		Author:  NewAuthorUseCase(repos.Course),
		Role:    NewRoleUseCase(repos.Role, repos.Course),
	}
}
