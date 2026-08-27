package command

import (
	"context"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/api/contracts/requests"
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/task"
	user "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"gorm.io/gorm"
)

func CreateUserCommand(ctx context.Context, userReq requests.CreateUserRequest) (*user.UserDTO, error) {
	commandFunc := func(db *gorm.DB, ctx context.Context) (*user.UserDTO, error) {
		userObj := user.User{
			Username: userReq.Username,
			Email:    userReq.Email,
			AuthId:   userReq.AuthId,
		}
		repo := user.NewRepo(db)

		created_user, err := repo.CreateOne(userObj)

		if err != nil {
			return nil, err
		}

		return created_user.ToUserDTO(), nil
	}

	return task.NewCommandTask(ctx, commandFunc).Await()
}
