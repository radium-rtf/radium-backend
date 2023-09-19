package validator

import (
	"github.com/go-playground/validator/v10"
	"net/url"
	"reflect"
)

func (v *validate) urlValidate(fl validator.FieldLevel) bool {
	if !fl.Field().CanConvert(reflect.TypeOf("")) {
		return false
	}

	s := fl.Field().String()
	if s == "" {
		return true
	}

	url, err := url.Parse(s)
	if err != nil || url.Scheme == "" {
		return false
	}
	if url.Host == "" && url.Fragment == "" && url.Opaque == "" {
		return false
	}

	return true
}
