package requests

import (
	"fmt"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"
	"github.com/go-playground/validator/v10"
)

type CreateUploadRequest struct {
	Purpose     string `json:"purpose" validate:"required,oneof=profile"`
	ContentType string `json:"contentType" validate:"required,oneof=image/jpeg image/png image/webp"`
}

func (r CreateUploadRequest) GetValidationError(err validator.FieldError) string {
	switch err.Field() {
	case "Purpose":
		return "Purpose must be one of: profile"
	case "ContentType":
		return "ContentType must be one of: image/jpeg, image/png, image/webp"
	default:
		return fmt.Sprintf("Invalid value for %s", err.Field())
	}
}

type CreateUploadRequestInput struct {
	middleware.AuthenticatedInput
	Body CreateUploadRequest
}
