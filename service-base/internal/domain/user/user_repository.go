package user

import (
	repository "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	*repository.BaseRepo[User]
}

func NewRepo(Db *gorm.DB) *UserRepo {
	return &UserRepo{BaseRepo: repository.NewBaseRepo[User](Db)}
}

func (repo *UserRepo) WithAuthId(authId string) *UserRepo {
	return &UserRepo{BaseRepo: repo.Where("auth_id = ?", authId)}
}

func (repo *UserRepo) WithUsernames(usernames []string) *UserRepo {
	return &UserRepo{BaseRepo: repo.Where("username IN ?", usernames)}
}

func (repo *UserRepo) WithUsername(username string) *UserRepo {
	return &UserRepo{BaseRepo: repo.Where("username = ?", username)}
}

func (repo *UserRepo) WithIds(ids []uuid.UUID) *UserRepo {
	return &UserRepo{BaseRepo: repo.Where("id IN ?", ids)}
}

func (repo *UserRepo) WithId(id uuid.UUID) *UserRepo {
	return &UserRepo{BaseRepo: repo.Where("id = ?", id)}
}

func (repo *UserRepo) WithRefreshTokenHash(hash string) *UserRepo {
	return &UserRepo{BaseRepo: repo.Where("refresh_token_hash = ?", hash)}
}

func (repo *UserRepo) IsActive() *UserRepo {
	return &UserRepo{BaseRepo: repo.Where("is_active = ?", true)}
}

func (repo *UserRepo) IsInactive() *UserRepo {
	return &UserRepo{BaseRepo: repo.Where("is_active = ?", false)}
}
