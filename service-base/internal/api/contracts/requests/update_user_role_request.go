package requests

import (
	"fmt"

	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"

	"github.com/go-playground/validator/v10"
)

type UpdateUserRoleRequest struct {
	Role string `json:"role" validate:"required,oneof=user admin"`
}

func (r UpdateUserRoleRequest) GetValidationError(err validator.FieldError) string {
	switch err.Field() {
	case "Role":
		switch err.Tag() {
		case "required":
			return "Role is required"
		case "oneof":
			return fmt.Sprintf("Role must be one of: %s, %s", user.RoleUser, user.RoleAdmin)
		default:
			return "Invalid role"
		}
	default:
		return fmt.Sprintf("Invalid value for %s", err.Field())
	}
}

type UpdateUserRoleRequestInput struct {
	ID   string `path:"id" doc:"ID (UUID) of the user whose role is being changed"`
	Body UpdateUserRoleRequest
}
