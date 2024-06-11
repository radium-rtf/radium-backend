package create

type GroupCreate struct {
	Name string `json:"name" validate:"required"`
}
