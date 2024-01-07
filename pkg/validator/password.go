package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

const (
	passwordMinLength   = 8
	passwordMaxLength   = 32
	passwordMinDigit    = 1
	passwordMinNonDigit = 1
)

var (
	lengthRegexp   = regexp.MustCompile(fmt.Sprintf(`^.{%d,%d}$`, passwordMinLength, passwordMaxLength))
	digitRegexp    = regexp.MustCompile(fmt.Sprintf(`[0-9]{%d,}`, passwordMinDigit))
	nonDigitRegexp = regexp.MustCompile(fmt.Sprintf(`[^0-9]{%d,}`, passwordMinNonDigit))
)

func (v *validate) passwordValidate(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		v.passwdErr = fmt.Errorf("field %s must be a string", fl.FieldName())
		return false
	}

	fieldValue := fl.Field().String()

	if ok := lengthRegexp.MatchString(fieldValue); !ok {
		v.passwdErr = fmt.Errorf("field %s must be between %d and %d characters", fl.FieldName(), passwordMinLength, passwordMaxLength)
		return false
	}

	if ok := digitRegexp.MatchString(fieldValue); !ok {
		v.passwdErr = fmt.Errorf("field %s must contain at least %d digit letter(s)", fl.FieldName(), passwordMinDigit)
		return false
	}

	if ok := nonDigitRegexp.MatchString(fieldValue); !ok {
		v.passwdErr = fmt.Errorf("field %s must contain at least %d non-digit letter(s)", fl.FieldName(), passwordMinNonDigit)
		return false
	}

	return true
}
