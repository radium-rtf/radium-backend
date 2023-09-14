package entity

type (
	User struct {
		DBModel
		Avatar           string
		Email            string `gorm:"unique; not null"`
		Name             string `gorm:"not null"`
		Password         string `gorm:"not null"`
		VerificationCode string
		IsVerified       bool
		Courses          []*Course `gorm:"many2many:course_students"`
		Groups           []*Group  `gorm:"many2many:group_student"`
		Sessions         []Session

		IsTeacher bool `gorm:"not null; default:false"`
		IsAuthor  bool `gorm:"not null; default:false"`
	}
)
