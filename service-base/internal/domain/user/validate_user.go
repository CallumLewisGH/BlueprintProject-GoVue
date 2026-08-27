package user

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetUsernameError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "min":
		return fmt.Sprintf("%s must be at least 8 characters long", err.Field())
	case "max":
		return fmt.Sprintf("%s must not exceed 16 characters", err.Field())
	case "alphanum":
		return fmt.Sprintf("%s can only contain letters and numbers", err.Field())
	default:
		return fmt.Sprintf("Invalid %s format", err.Field())
	}
}

func GetAuthIdError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	default:
		return fmt.Sprintf("Invalid %s format", err.Field())
	}
}

func GetEmailError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	default:
		return fmt.Sprintf("Invalid %s format", err.Field())
	}
}

func GetTimezoneError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "timezone":
		return fmt.Sprintf("%s must be a valid IANA timezone", err.Field())
	default:
		return fmt.Sprintf("Invalid %s format", err.Field())
	}
}

func GetBioError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "max":
		return fmt.Sprintf("%s must not exceed 500 characters", err.Field())
	default:
		return fmt.Sprintf("Invalid %s format", err.Field())
	}
}

func GetProfilePictureError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", err.Field())
	case "http_url":
		return fmt.Sprintf("%s must be a valid URL", err.Field())
	default:
		return fmt.Sprintf("Invalid %s format", err.Field())
	}
}
