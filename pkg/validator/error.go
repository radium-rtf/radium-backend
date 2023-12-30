package validator

type ValidationError struct {
	Tag     string `json:"tag"`
	Param   string `json:"param"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (err ValidationError) Error() string {
	return err.Message
}

func (v *validate) newValidationError(field, tag, param string, err error) error {
	validationErr := ValidationError{Tag: tag, Field: field, Message: err.Error(), Param: param}

	switch tag {
	case "password":
		validationErr.Message = v.passwdErr.Error()
	default:
		validationErr.Message = err.Error()
	}

	return validationErr
}
