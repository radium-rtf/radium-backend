package decode

import (
	"encoding/json"
	"github.com/radium-rtf/radium-backend/pkg/validator"
	"io"
)

func Json(body io.ReadCloser, v any) error {
	err := json.NewDecoder(body).Decode(v)
	if err != nil {
		return err
	}

	return validator.Struct(v)
}
