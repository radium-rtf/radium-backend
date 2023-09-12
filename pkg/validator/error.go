package validator

import "fmt"

func (v *validate) newValidationError(field string, tag string, err error) error {
	switch tag {
	case "email":
		return fmt.Errorf("field %s must be a valid email address: pattern - %s", field, emailRegexp.String())
	case "password":
		return v.passwdErr
	default:
		return err
	}
}
