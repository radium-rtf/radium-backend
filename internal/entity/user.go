package entity

type (
	UserDto struct {
		Id       uint   `json:"id"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Username string `json:"username"`
	}
	User struct {
		Id       uint
		Email    string
		Name     string
		Username string
		Password string
	}
)

func NewUserDto(user User) UserDto {
	return UserDto{
		Id:       user.Id,
		Email:    user.Email,
		Name:     user.Name,
		Username: user.Username,
	}
}
