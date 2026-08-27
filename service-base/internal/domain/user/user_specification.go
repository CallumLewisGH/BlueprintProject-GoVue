package user

import (
	"fmt"

	queryFilter "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/query"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/specification"
	"github.com/google/uuid"
)

type UserSpecification func(repo *UserRepo) *UserRepo

var userSpecifications = map[specification.SpecificationKey]func([]string) (UserSpecification, error){
	"withusernames": withUsernamesSpecification,
	"withids":       withIdsSpecification,
	"isactive":      isActiveSpecification,
	"":              emptySpecification,
}

func GetUserSpecification(specificationName specification.SpecificationKey, value []string) (UserSpecification, error) {
	factory, exists := userSpecifications[specificationName]
	if !exists {
		return nil, fmt.Errorf("unknown specification: %s", specificationName)
	}
	return factory(value)
}

func (repo *UserRepo) ApplyQuery(query queryFilter.Query) (*UserRepo, error) {
	for _, param := range query {
		spec, err := GetUserSpecification(param.Name, param.Value)
		if err != nil {
			return nil, err
		}
		repo = spec(repo)
	}
	return repo, nil
}

func emptySpecification(data []string) (UserSpecification, error) {
	return func(repo *UserRepo) *UserRepo {
		return repo
	}, nil
}

func withUsernamesSpecification(data []string) (UserSpecification, error) {
	return func(repo *UserRepo) *UserRepo {
		return repo.WithUsernames(data)
	}, nil
}

func withIdsSpecification(data []string) (UserSpecification, error) {
	var ids []uuid.UUID
	for _, idStr := range data {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, fmt.Errorf("invalid uuid: %v", err)
		}
		ids = append(ids, id)
	}

	return func(repo *UserRepo) *UserRepo {
		return repo.WithIds(ids)
	}, nil
}

func isActiveSpecification(data []string) (UserSpecification, error) {
	if len(data) != 1 {
		return nil, fmt.Errorf("isactive expects exactly one value, got %d", len(data))
	}

	switch data[0] {
	case "true", "1", "yes":
		return func(repo *UserRepo) *UserRepo { return repo.IsActive() }, nil
	case "false", "0", "no":
		return func(repo *UserRepo) *UserRepo { return repo.IsInactive() }, nil
	default:
		return nil, fmt.Errorf("invalid active flag: %s", data[0])
	}
}
