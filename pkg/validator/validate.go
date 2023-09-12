package validator

import (
	"github.com/go-playground/validator/v10"
)

type validate struct {
	v         *validator.Validate
	passwdErr error
}

func newValidate() (*validate, error) {
	v := validate{v: validator.New()}

	err := v.v.RegisterValidation("email", v.emailValidate)
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
	v, err := newValidate()
	if err != nil {
		return err
	}

	err = v.v.Struct(i)
	if err == nil {
		return nil
	}

	fieldErr := err.(validator.ValidationErrors)[0]
	return v.newValidationError(fieldErr.Field(), fieldErr.Tag(), err)
}
