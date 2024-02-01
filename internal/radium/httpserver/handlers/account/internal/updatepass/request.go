package updatepass

type Password struct {
	New     string `json:"new" validate:"password"`
	Current string `json:"current" validate:"required"`
}
