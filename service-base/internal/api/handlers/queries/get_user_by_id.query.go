package query

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserByIdQuery(ctx context.Context, id uuid.UUID) (*user.UserDTO, error) {
	queryFunc := func(db *gorm.DB, ctx context.Context) (*user.UserDTO, error) {
		repo := user.NewRepo(db)
		user, err := repo.WithId(id).First()
		if err != nil {
			return nil, err
		}

		return user.ToUserDTO(), nil
	}
	return task.NewQueryTask(ctx, queryFunc).Await()
}
