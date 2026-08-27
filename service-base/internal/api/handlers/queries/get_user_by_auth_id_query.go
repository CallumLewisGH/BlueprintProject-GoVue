package query

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"gorm.io/gorm"
)

func GetUserByAuthIdQuery(ctx context.Context, userAuthId string) (*user.UserDTO, error) {
	queryFunc := func(db *gorm.DB, ctx context.Context) (*user.UserDTO, error) {
		repo := user.NewRepo(db)
		user, err := repo.WithAuthId(userAuthId).First()
		if err != nil {
			return nil, err
		}

		return user.ToUserDTO(), nil
	}

	return task.NewQueryTask(ctx, queryFunc).Await()
}
