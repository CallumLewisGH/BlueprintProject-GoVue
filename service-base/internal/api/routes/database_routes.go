package routes

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/responses"
	command "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/commands"
	query "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/handlers/queries"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"
	"github.com/danielgtaylor/huma/v2"
)

func RegisterDatabaseRoutes(s *api.Server) {
	databaseGroup := huma.NewGroup(s.API, "/database")
	databaseGroup.UseSimpleModifier(huma.OperationTags("database"))

	huma.Get(databaseGroup, "/health", getHealth)

	// Migrations touch schema/data and must be admin-only, not just logged-in.
	migrationsGroup := huma.NewGroup(s.API, "/database")
	migrationsGroup.UseSimpleModifier(huma.OperationTags("database"))
	migrationsGroup.UseMiddleware(middleware.RequireAuthHuma(s.API), middleware.RequireAdminHuma(s.API))
	migrationsGroup.UseSimpleModifier(func(o *huma.Operation) {
		o.Security = []map[string][]string{{"BearerAuth": {}}}
	})

	huma.Post(migrationsGroup, "/migrations", runMigrations)
}

func getHealth(ctx context.Context, input *struct{}) (*responses.HealthResponse, error) {
	health := query.CheckDatabaseHealth()
	return &responses.HealthResponse{Body: health}, nil
}

func runMigrations(ctx context.Context, input *struct{}) (*responses.MigrationsResponse, error) {
	err := command.RunDatabaseMigrations()
	if err != nil {
		return nil, huma.Error500InternalServerError(err.Error())
	}
	return &responses.MigrationsResponse{Body: "Migrations Successfull"}, nil
}
