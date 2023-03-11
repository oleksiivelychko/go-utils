package validator_helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	validator.FieldError
}

type ValidationErrors []ValidationError

type Validation struct {
	ValidatorValidate *validator.Validate
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s'. Error: Field validation for '%s' failed on the '%s' tag",
		validationError.Namespace(),
		validationError.Field(),
		validationError.Tag(),
	)
}

func (validationErrors ValidationErrors) Errors() []string {
	var validationErrArr []string
	for _, err := range validationErrors {
		validationErrArr = append(validationErrArr, err.Error())
	}

	return validationErrArr
}

func (validation *Validation) Validate(i interface{}) ValidationErrors {
	validationStruct := validation.ValidatorValidate.Struct(i)
	if validationStruct == nil {
		return ValidationErrors{}
	}

	validationErrors := validationStruct.(validator.ValidationErrors)
	if len(validationErrors) == 0 {
		return nil
	}

	var validationErrArr []ValidationError
	for _, err := range validationErrors {
		ve := ValidationError{err.(validator.FieldError)}
		validationErrArr = append(validationErrArr, ve)
	}

	return validationErrArr
}
