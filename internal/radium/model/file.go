package model

import (
	"github.com/radium-rtf/radium-backend/internal/radium/entity"
)

type (
	File struct {
		Location  string  `json:"location"`
		Name      string  `json:"name"`
		SizeInKiB float64 `json:"sizeInKiB"`
		Type      string  `json:"type"`
	}
)

func NewFile(file *entity.File) *File {
	if file == nil {
		return nil
	}
	return &File{
		Type:      file.Type,
		Location:  file.Url,
		Name:      file.Name,
		SizeInKiB: float64(file.Size) / 1024,
	}
}
