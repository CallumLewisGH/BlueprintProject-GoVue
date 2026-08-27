package query

import (
	"context"
	"fmt"

	queryFilter "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/query"
	task "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"gorm.io/gorm"
)

func GetUsersQuery(ctx context.Context, query queryFilter.Query) ([]user.UserDTO, error) {
	queryFunc := func(db *gorm.DB, ctx context.Context) ([]user.UserDTO, error) {
		userRepo := user.NewRepo(db)

		userRepo, err := userRepo.ApplyQuery(query)
		if err != nil {
			return nil, fmt.Errorf("apply query: %w", err)
		}

		users, err := userRepo.Find()
		if err != nil {
			return nil, fmt.Errorf("find users: %w", err)
		}

		return user.ToUserDTOs(users), nil
	}

	return task.NewQueryTask(ctx, queryFunc).Await()
}
