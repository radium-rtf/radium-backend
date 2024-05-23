package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	User struct {
		bun.BaseModel `bun:"table:users"`
		DBModel

		Avatar   sql.NullString
		Email    string
		Name     string
		Password string

		Contact *Contact  `bun:"rel:has-one,join:id=user_id"`
		Roles   *Roles    `bun:"rel:has-one,join:id=user_id"`
		Courses []*Course `bun:"m2m:students,join:User=Course"`
		Answers []*Answer `bun:"rel:has-many,join:id=user_id"`

		Author   []*Course `bun:"m2m:course_author,join:User=Course"`
		Coauthor []*Course `bun:"m2m:course_coauthor,join:User=Course"`

		LastVisitedPage *LastVisitedPage `bun:"rel:has-one,join:id=user_id"`
	}

	UnverifiedUser struct {
		Id       uuid.UUID `bun:",pk"`
		Avatar   sql.NullString
		Email    string
		Name     string
		Password string

		VerificationCode string
		ExpiresAt        time.Time
	}
)
