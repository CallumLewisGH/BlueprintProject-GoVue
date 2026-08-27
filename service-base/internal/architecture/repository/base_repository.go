package repository

import (
	"gorm.io/gorm"
)

type BaseRepo[T any] struct {
	Db *gorm.DB
}

func NewBaseRepo[T any](Db *gorm.DB) *BaseRepo[T] {
	return &BaseRepo[T]{Db: Db}
}

func (repo *BaseRepo[T]) Joins(query string, args ...any) *BaseRepo[T] {
	return &BaseRepo[T]{Db: repo.Db.Joins(query, args...)}
}

func (repo *BaseRepo[T]) Unscoped() *BaseRepo[T] {
	return &BaseRepo[T]{Db: repo.Db.Unscoped()}
}

func (repo *BaseRepo[T]) Limit(limit int) *BaseRepo[T] {
	return &BaseRepo[T]{Db: repo.Db.Limit(limit)}
}

func (repo *BaseRepo[T]) Offset(offset int) *BaseRepo[T] {
	return &BaseRepo[T]{Db: repo.Db.Offset(offset)}
}

func (repo *BaseRepo[T]) Order(order string) *BaseRepo[T] {
	return &BaseRepo[T]{Db: repo.Db.Order(order)}
}

func (repo *BaseRepo[T]) Where(query any, args ...any) *BaseRepo[T] {
	return &BaseRepo[T]{Db: repo.Db.Where(query, args...)}
}

func (repo *BaseRepo[T]) Preload(query string, args ...any) *BaseRepo[T] {
	return &BaseRepo[T]{Db: repo.Db.Preload(query, args...)}
}

// Terminating Operations
func (repo *BaseRepo[T]) Find() ([]T, error) {
	var results []T
	err := repo.Db.Find(&results).Error
	return results, err
}

func (repo *BaseRepo[T]) First() (T, error) {
	var result T
	err := repo.Db.First(&result).Error
	return result, err
}

func (repo *BaseRepo[T]) Count() (int64, error) {
	var count int64
	err := repo.Db.Model(new(T)).Count(&count).Error
	return count, err
}

func (repo *BaseRepo[T]) CreateOne(dest T) (T, error) {
	if err := repo.Db.Create(&dest).Error; err != nil {
		return *new(T), err
	}
	return dest, nil
}

func (repo *BaseRepo[T]) CreateMany(dest []T) ([]T, error) {
	if err := repo.Db.Create(&dest).Error; err != nil {
		return nil, err
	}
	return dest, nil
}

func (repo *BaseRepo[T]) DeleteOne() (T, error) {
	var model T
	if err := repo.Db.Unscoped().First(&model).Error; err != nil {
		return *new(T), err
	}
	if err := repo.Db.Delete(&model).Error; err != nil {
		return *new(T), err
	}
	return model, nil
}

func (repo *BaseRepo[T]) DeleteMany() ([]T, error) {
	var models []T
	if err := repo.Db.Unscoped().Find(&models).Error; err != nil {
		return nil, err
	}
	if err := repo.Db.Delete(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (repo *BaseRepo[T]) UpdateOne(updates any) (T, error) {
	result := repo.Db.Model(new(T)).Updates(updates)
	if result.Error != nil {
		return *new(T), result.Error
	}
	var updatedModel T
	if err := repo.Db.First(&updatedModel).Error; err != nil {
		return *new(T), err
	}
	return updatedModel, nil
}

func (repo *BaseRepo[T]) UpdateMany(updates any) ([]T, error) {
	result := repo.Db.Model(new(T)).Updates(updates)
	if result.Error != nil {
		return nil, result.Error
	}

	var updatedModels []T
	if err := repo.Db.Find(&updatedModels).Error; err != nil {
		return nil, err
	}
	return updatedModels, nil
}
