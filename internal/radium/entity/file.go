package entity

type (
	File struct {
		Url  string `bun:",pk"`
		Name string
		Type string
		Size int64
	}
)
