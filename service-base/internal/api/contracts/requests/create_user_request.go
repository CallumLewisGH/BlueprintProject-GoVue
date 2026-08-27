package requests

import (
	"fmt"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=8,max=16,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	AuthId   string `json:"authId" validate:"required"`
}

func (r CreateUserRequest) GetValidationError(err validator.FieldError) string {
	switch err.Field() {
	case "Username":
		return user.GetUsernameError(err)
	case "Email":
		return user.GetEmailError(err)
	case "AuthId":
		return user.GetAuthIdError(err)
	default:
		return fmt.Sprintf("Invalid value for %s", err.Field())
	}
}
