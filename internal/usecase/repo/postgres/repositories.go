package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres/db"
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
}

func NewRepositories(pg *db.Query) Repositories {
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
	}
}
