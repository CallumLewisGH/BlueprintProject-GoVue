package queryFilter

import "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/specification"

type QueryParameter struct {
	Name  specification.SpecificationKey
	Value []string
}

type Query []QueryParameter
