package postauthor

type Email struct {
	Email string `json:"email" validate:"email"`
}
