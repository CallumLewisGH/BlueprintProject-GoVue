package requests

import "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"

type EmptyRequestInput struct {
	middleware.AuthenticatedInput
}
