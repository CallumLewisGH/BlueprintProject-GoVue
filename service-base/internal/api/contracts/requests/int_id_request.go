package requests

import "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"

type IntIdRequestInput struct {
	middleware.AuthenticatedInput
	ID int `path:"id" doc:"ID (integer)"`
}
