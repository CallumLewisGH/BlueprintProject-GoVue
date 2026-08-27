package command

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/requests"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateUserByIdCommand(ctx context.Context, id uuid.UUID, req requests.UpdateUserRequest) (*user.UserDTO, error) {
	commandFunc := func(db *gorm.DB, ctx context.Context) (*user.UserDTO, error) {
		repo := user.NewRepo(db)
		updated_user, err := repo.WithId(id).UpdateOne(req)

		if err != nil {
			return nil, err
		}

		return updated_user.ToUserDTO(), nil
	}
	return task.NewCommandTask(ctx, commandFunc).Await()
}
