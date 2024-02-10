package order

type (
	Order struct {
		Order uint `json:"order" validate:"min=1"`
	}
)
