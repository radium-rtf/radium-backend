package usecase

import (
	"github.com/radium-rtf/radium-backend/internal/lib/session"
	"github.com/radium-rtf/radium-backend/internal/usecase/repo/postgres"
	"github.com/radium-rtf/radium-backend/pkg/auth"
	"github.com/radium-rtf/radium-backend/pkg/filestorage"
	"github.com/radium-rtf/radium-backend/pkg/hash"
)

type Dependencies struct {
	Repos   postgres.Repositories
	Storage filestorage.Storage

	PasswordHasher hash.Hasher
	TokenManager   auth.TokenManager
	Session        session.Session
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

	Deps Dependencies
}

func NewUseCases(deps Dependencies) UseCases {
	repos := deps.Repos

	return UseCases{
		Deps: deps,

		Account: NewAccountUseCase(repos.User, repos.Course, deps.PasswordHasher),
		Answer:  NewAnswerUseCase(repos.Section, repos.Answer),
		Auth:    NewAuthUseCase(repos.User, repos.Session, deps.PasswordHasher, deps.Session),
		Course:  NewCourseUseCase(repos.Course),
		File:    NewFileUseCase(deps.Storage),
		Group:   NewGroupUseCase(repos.Group, repos.Course, repos.Answer),
		Module:  NewModuleUseCase(repos.Module),
		Page:    NewPageUseCase(repos.Page),
		Review:  NewReviewUseCase(repos.Review),
		Section: NewSectionUseCase(repos.Section),
		Teacher: NewTeacherUseCase(repos.Teacher),
	}
}
