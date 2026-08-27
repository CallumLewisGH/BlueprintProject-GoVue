package command

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteUserByIdCommand(ctx context.Context, id uuid.UUID) (*user.UserDTO, error) {
	commandFunc := func(db *gorm.DB, ctx context.Context) (*user.UserDTO, error) {
		repo := user.NewRepo(db)
		delted_user, err := repo.WithId(id).Unscoped().DeleteOne()

		if err != nil {
			return nil, err
		}

		return delted_user.ToUserDTO(), nil
	}
	return task.NewCommandTask(ctx, commandFunc).Await()
}
