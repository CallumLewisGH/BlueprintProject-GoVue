package database

import (
	"sync"

	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
)

var (
	registryInstance *ModelRegistry
	registryOnce     sync.Once
)

type ModelRegistry struct {
	models []any
}

func (registry *ModelRegistry) RegisterAllModels() {
	// Register all models here
	// This is where you would add all your models to the registry.
	// For example, if you have a User model, you would add &user.User{} to the slice
	// You cannot forget this if you want a table with the model in the database

	dbModels := [...]any{
		&user.User{},
	}

	for _, model := range dbModels {
		registry.models = append(registry.models, model)
	}
}

func NewModelRegistry() *ModelRegistry {
	return &ModelRegistry{
		models: make([]any, 0),
	}
}

func GetModelRegistry() *ModelRegistry {
	registryOnce.Do(func() {
		registryInstance = &ModelRegistry{
			models: make([]any, 0),
		}
		registryInstance.RegisterAllModels()
	})
	return registryInstance
}

func ResetModelRegistry() {
	registryInstance = nil
	registryOnce = sync.Once{}
}
