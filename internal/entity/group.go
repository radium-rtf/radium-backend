package entity

type (
	Group struct {
		DBModel
		Name       string    `gorm:"not null"`
		InviteCode string    `gorm:"not null"`
		Courses    []*Course `gorm:"many2many:group_course"`
		Students   []*User   `gorm:"many2many:group_student"`
	}
)
