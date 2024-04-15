package postgres

import (
	"github.com/radium-rtf/radium-backend/pkg/postgres"
)

type Repositories struct {
	Course       Course
	Answer       Answer
	Group        Group
	Module       Module
	Page         Page
	Review       Review
	Section      Section
	Session      Session
	Teacher      Teacher
	User         User
	Role         Role
	File         File
	Notification Notification
}

func NewRepositories(pg *postgres.Postgres) Repositories {
	return Repositories{
		Course:       NewCourseRepo(pg),
		Answer:       NewAnswerRepo(pg),
		Group:        NewGroupRepo(pg),
		Module:       NewModuleRepo(pg),
		Page:         NewPageRepo(pg),
		Review:       NewReviewRepo(pg),
		Section:      NewSectionRepo(pg),
		Session:      NewSessionRepo(pg),
		Teacher:      NewTeacherRepo(pg),
		User:         NewUserRepo(pg),
		Role:         NewRoleRepo(pg),
		File:         NewFileRepo(pg),
		Notification: NewNotificationRepo(pg),
	}
}
