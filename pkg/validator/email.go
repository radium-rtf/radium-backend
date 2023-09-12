package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var emailRegexp = regexp.MustCompile("[a-zA-Z.]@urfu.(me|ru)")

func (v *validate) emailValidate(fl validator.FieldLevel) bool {
	return emailRegexp.MatchString(fl.Field().String())
}
