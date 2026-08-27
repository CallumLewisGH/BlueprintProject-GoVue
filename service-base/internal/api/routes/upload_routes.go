package routes

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/requests"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/responses"
	command "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/commands"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/validation"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterUploadRoutes(s *api.Server) {
	uploadGroup := huma.NewGroup(s.API, "/uploads")
	uploadGroup.UseSimpleModifier(huma.OperationTags("uploads"))
	uploadGroup.UseMiddleware(middleware.RequireAuthHuma(s.API))
	uploadGroup.UseSimpleModifier(func(o *huma.Operation) {
		o.Security = []map[string][]string{{"BearerAuth": {}}}
	})

	huma.Post(uploadGroup, "", createUpload)
}

func createUpload(ctx context.Context, input *requests.CreateUploadRequestInput) (*responses.UploadResponse, error) {
	if errs := validation.ValidateBody(input.Body); errs != nil {
		return nil, huma.Error400BadRequest("Validation failed", errs...)
	}

	uploadURL, publicURL, err := command.CreateUploadURLCommand(input.UserID, input.Body.Purpose, input.Body.ContentType)
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	return responses.ToUploadResponse(uploadURL, publicURL), nil
}
