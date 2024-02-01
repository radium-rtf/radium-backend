package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	Group struct {
		bun.BaseModel `bun:"table:groups"`
		DBModel

		Name       string
		InviteCode string

		Courses  []*Course `bun:"m2m:group_course,join:Group=Course"`
		Students []*User   `bun:"m2m:group_student,join:Group=User"`
	}

	GroupStudent struct {
		bun.BaseModel `bun:"table:group_student"`

		GroupId uuid.UUID `bun:",pk"`
		Group   *Group    `bun:"rel:belongs-to,join:group_id=id"`

		UserId uuid.UUID `bun:",pk"`
		User   *User     `bun:"rel:belongs-to,join:user_id=id"`
	}

	GroupCourse struct {
		bun.BaseModel `bun:"table:group_course"`

		GroupId uuid.UUID `bun:",pk"`
		Group   *Group    `bun:"rel:belongs-to,join:group_id=id"`

		CourseId uuid.UUID `bun:",pk"`
		Course   *Course   `bun:"rel:belongs-to,join:course_id=id"`
	}
)
