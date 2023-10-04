package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	TeacherCourseGroup struct {
		bun.BaseModel `bun:"table:teacher_course_group"`

		UserId   uuid.UUID `bun:",pk"`
		CourseId uuid.UUID `bun:",pk"`
		GroupId  uuid.UUID `bun:",pk"`

		User   *User   `bun:"rel:belongs-to,join:user_id=id"`
		Course *Course `bun:"rel:belongs-to,join:course_id=id"`
		Group  *Group  `bun:"rel:belongs-to,join:group_id=id"`
	}
)
