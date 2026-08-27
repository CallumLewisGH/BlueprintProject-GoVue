package routes

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/requests"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/responses"
	command "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/commands"
	query "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/queries"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/validation"
	"github.com/danielgtaylor/huma/v2"
	"github.com/google/uuid"
)

func RegisterUserRoutes(s *api.Server) {
	userGroup := huma.NewGroup(s.API, "/users")
	userGroup.UseSimpleModifier(huma.OperationTags("users"))
	userGroup.UseMiddleware(middleware.RequireAuthHuma(s.API))
	userGroup.UseSimpleModifier(func(o *huma.Operation) {
		o.Security = []map[string][]string{{"BearerAuth": {}}}
	})

	huma.Get(userGroup, "/me", getAuthenticatedUser)

	huma.Get(userGroup, "/{id}", getUserById)
	huma.Post(userGroup, "/search", getUsers)
	huma.Patch(userGroup, "/me", updateAuthenticatedUserById)
	huma.Delete(userGroup, "/me", deleteAuthenticatedUserById)

	// Deliberately no public POST /users: user creation only happens inside
	// the OAuth callback (auth_routes.go), which derives authId/email from
	// the verified provider response instead of trusting client input.

	adminUserGroup := huma.NewGroup(s.API, "/users")
	adminUserGroup.UseSimpleModifier(huma.OperationTags("users", "admin"))
	adminUserGroup.UseMiddleware(middleware.RequireAuthHuma(s.API), middleware.RequireAdminHuma(s.API))
	adminUserGroup.UseSimpleModifier(func(o *huma.Operation) {
		o.Security = []map[string][]string{{"BearerAuth": {}}}
	})

	huma.Patch(adminUserGroup, "/{id}/role", updateUserRole)
}

func getAuthenticatedUser(ctx context.Context, input *requests.EmptyRequestInput) (*responses.UserPrivateProfileResponse, error) {
	user, err := query.GetUserByIdQuery(ctx, input.UserID)

	if err != nil {
		return nil, huma.Error404NotFound("User not found in database")
	}

	return responses.ToUserPrivateProfileResponse(user), nil
}

func getUserById(ctx context.Context, input *requests.GuidIdRequestInput) (*responses.UserPublicProfileResponse, error) {
	userID, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("User ID must be a valid GUID")
	}

	user, err := query.GetUserByIdQuery(ctx, userID)
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	return responses.ToUserPublicProfileResponse(user), nil
}

func getUsers(ctx context.Context, input *requests.QueryRequestInput) (*responses.UserPublicProfileResponses, error) {
	users, err := query.GetUsersQuery(ctx, input.Body)
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}
	return responses.ToUserPublicProfileResponses(users), nil
}

func updateAuthenticatedUserById(ctx context.Context, input *requests.UpdateUserRequestInput) (*responses.UserPrivateProfileResponse, error) {
	if errs := validation.ValidateBody(input.Body); errs != nil {
		return nil, huma.Error400BadRequest("Validation failed", errs...)
	}

	updatedUser, err := command.UpdateUserByIdCommand(ctx, input.UserID, input.Body)
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	return responses.ToUserPrivateProfileResponse(updatedUser), nil
}

func deleteAuthenticatedUserById(ctx context.Context, input *requests.EmptyRequestInput) (*responses.UserPrivateProfileResponse, error) {
	user, err := command.DeleteUserByIdCommand(ctx, input.UserID)
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	return responses.ToUserPrivateProfileResponse(user), nil
}

func updateUserRole(ctx context.Context, input *requests.UpdateUserRoleRequestInput) (*responses.UserPrivateProfileResponse, error) {
	if errs := validation.ValidateBody(input.Body); errs != nil {
		return nil, huma.Error400BadRequest("Validation failed", errs...)
	}

	targetUserID, err := uuid.Parse(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("User ID must be a valid GUID")
	}

	updatedUser, err := command.UpdateUserRoleCommand(ctx, targetUserID, input.Body.Role)
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}

	return responses.ToUserPrivateProfileResponse(updatedUser), nil
}
