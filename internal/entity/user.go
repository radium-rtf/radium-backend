package entity

import (
	"database/sql"
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

		Roles   *Roles    `bun:"rel:has-one,join:id=user_id"`
		Courses []*Course `bun:"m2m:course_student,join:User=Course"`
		Answers []*Answer `bun:"rel:has-many,join:id=user_id"`

		Author   []*Course `bun:"m2m:course_author,join:User=Course"`
		Coauthor []*Course `bun:"m2m:course_coauthor,join:User=Course"`
	}
)
