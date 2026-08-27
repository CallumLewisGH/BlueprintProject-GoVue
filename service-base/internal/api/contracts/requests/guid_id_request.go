package requests

import "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"

type GuidIdRequestInput struct {
	middleware.AuthenticatedInput
	ID string `path:"id" doc:"ID (UUID)"`
}
