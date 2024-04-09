package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var slugRegexp = regexp.MustCompile("^[a-z|0-9|_|-]+$")

func (v *validate) slugValidate(fl validator.FieldLevel) bool {
	return slugRegexp.MatchString(fl.Field().String())
}
