package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

type validate struct {
	v         *validator.Validate
	passwdErr error
}

func newValidate() (*validate, error) {
	v := validate{v: validator.New(validator.WithRequiredStructEnabled())}

	err := v.v.RegisterValidation("url", v.urlValidate)
	if err != nil {
		return nil, err
	}
	err = v.v.RegisterValidation("password", v.passwordValidate)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func Struct(i any) error {
	v, err := newValidate() // TODO: хотелось бы прокидывать ошибки не через структуру, тк библиотека содержит кэш
	if err != nil {
		return err
	}

	err = v.v.Struct(i)
	if err == nil {
		return nil
	}

	fieldErr := err.(validator.ValidationErrors)[0]
	return v.newValidationError(strings.ToLower(fieldErr.Field()), fieldErr.Tag(), fieldErr.Param(), err)
}
