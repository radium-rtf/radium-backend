package model

type (
	Link struct {
		Name string `json:"name" validate:"required,min=1,max=15"`
		Link string `json:"link" validate:"required,url"`
	}
)
