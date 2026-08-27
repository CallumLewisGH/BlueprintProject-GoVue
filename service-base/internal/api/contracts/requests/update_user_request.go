package requests

import (
	"fmt"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"

	"github.com/go-playground/validator/v10"
)

type UpdateUserRequest struct {
	Username       *string `json:"username,omitempty" validate:"omitempty,min=8,max=16,alphanum"`
	Email          *string `json:"email,omitempty" validate:"omitempty,email"`
	Timezone       *string `json:"timezone,omitempty" validate:"omitempty,timezone"`
	ProfilePicture *string `json:"profilePicture,omitempty" validate:"omitempty,http_url"`
	Bio            *string `json:"bio,omitempty" validate:"omitempty,max=500"`
}

func (r UpdateUserRequest) GetValidationError(err validator.FieldError) string {
	switch err.Field() {
	case "Username":
		return user.GetUsernameError(err)
	case "Email":
		return user.GetEmailError(err)
	case "Timezone":
		return user.GetTimezoneError(err)
	case "ProfilePicture":
		return user.GetProfilePictureError(err)
	case "Bio":
		return user.GetBioError(err)
	default:
		return fmt.Sprintf("Invalid value for %s", err.Field())
	}
}

type UpdateUserRequestInput struct {
	middleware.AuthenticatedInput
	Body UpdateUserRequest
}
