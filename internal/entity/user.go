package entity

type (
	UserDto struct {
		Id    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	User struct {
		Id               string
		Email            string
		Name             string
		Password         string
		VerificationCode string
		IsVerified       bool
	}

	UserName struct {
		Name string `json:"name"`
	}

	PasswordUpdate struct {
		New     string `json:"new"`
		Current string `json:"current"`
	}
)

func NewUserDto(user User) UserDto {
	return UserDto{
		Id:    user.Id,
		Email: user.Email,
		Name:  user.Name,
	}
}
