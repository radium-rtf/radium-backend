package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

const (
	passwordMinLength = 8
	passwordMaxLength = 32
	passwordMinLower  = 1
	passwordMinUpper  = 1
	passwordMinDigit  = 1
	passwordMinSymbol = 1
)

var (
	lengthRegexp    = regexp.MustCompile(fmt.Sprintf(`^.{%d,%d}$`, passwordMinLength, passwordMaxLength))
	lowerCaseRegexp = regexp.MustCompile(fmt.Sprintf(`[a-z]{%d,}`, passwordMinLower))
	upperCaseRegexp = regexp.MustCompile(fmt.Sprintf(`[A-Z]{%d,}`, passwordMinUpper))
	digitRegexp     = regexp.MustCompile(fmt.Sprintf(`[0-9]{%d,}`, passwordMinDigit))
	symbolRegexp    = regexp.MustCompile(fmt.Sprintf(`[^A-Za-z0-9]{%d,}`, passwordMinSymbol))
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
	if ok := lowerCaseRegexp.MatchString(fieldValue); !ok {
		v.passwdErr = fmt.Errorf("field %s must contain at least %d lowercase letter(s)", fl.FieldName(), passwordMinLower)
		return false
	}
	if ok := upperCaseRegexp.MatchString(fieldValue); !ok {
		v.passwdErr = fmt.Errorf("field %s must contain at least %d uppercase letter(s)", fl.FieldName(), passwordMinUpper)
		return false
	}
	if ok := digitRegexp.MatchString(fieldValue); !ok {
		v.passwdErr = fmt.Errorf("field %s must contain at least %d digit(s)", fl.FieldName(), passwordMinDigit)
		return false
	}
	if ok := symbolRegexp.MatchString(fieldValue); !ok {
		v.passwdErr = fmt.Errorf("field %s must contain at least %d special character(s)", fl.FieldName(), passwordMinSymbol)
		return false
	}

	return true
}
