package requests

import (
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/middleware"
	queryFilter "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/query"
)

type QueryRequestInput struct {
	middleware.AuthenticatedInput
	Body queryFilter.Query
}
