package command

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateUserRoleCommand(ctx context.Context, targetUserID uuid.UUID, role string) (*user.UserDTO, error) {
	commandFunc := func(db *gorm.DB, ctx context.Context) (*user.UserDTO, error) {
		repo := user.NewRepo(db)
		updatedUser, err := repo.WithId(targetUserID).UpdateOne(map[string]any{"role": role})
		if err != nil {
			return nil, err
		}

		return updatedUser.ToUserDTO(), nil
	}
	return task.NewCommandTask(ctx, commandFunc).Await()
}
