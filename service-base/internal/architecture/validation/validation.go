package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ValidationMessenger interface {
	GetValidationError(err validator.FieldError) string
}

var (
	validate *validator.Validate = validator.New()
)

func ValidateBody[messenger ValidationMessenger](input messenger) []error {
	var validationErrors []error

	err := validate.Struct(input)
	if err == nil {
		return nil
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, err := range validatorErrs {
		validationErrors = append(validationErrors, errors.New(input.GetValidationError(err)))
	}

	return validationErrors
}
