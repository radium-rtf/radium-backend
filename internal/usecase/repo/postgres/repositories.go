package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Repositories struct {
	Course
	Answer
	Group
	Module
	Page
	Review
	Section
	Session
	Teacher
	User
	Role
	File
}

func NewRepositories(pg *postgres.Postgres) Repositories {
	return Repositories{
		Course:  NewCourseRepo(pg),
		Answer:  NewAnswerRepo(pg),
		Group:   NewGroupRepo(pg),
		Module:  NewModuleRepo(pg),
		Page:    NewPageRepo(pg),
		Review:  NewReviewRepo(pg),
		Section: NewSectionRepo(pg),
		Session: NewSessionRepo(pg),
		Teacher: NewTeacherRepo(pg),
		User:    NewUserRepo(pg),
		Role:    NewRoleRepo(pg),
		File:    NewFileRepo(pg),
	}
}
