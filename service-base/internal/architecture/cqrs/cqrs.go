package cqrs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/database"
	"gorm.io/gorm"
)

type QueryResult[R any] struct {
	Data R
	Err  error
}

func DbQuery[R any](ctx context.Context, queryFunc func(*gorm.DB, context.Context) (R, error)) (R, error) {
	db := database.GetDatabase().GetGormDatabase()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	tx := db.WithContext(ctx)
	return queryFunc(tx, ctx)
}

func DbExecute[R any](ctx context.Context, commandFunc func(*gorm.DB, context.Context) (R, error)) (result R, err error) {
	db := database.GetDatabase().GetGormDatabase()

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	tx := db.Begin().WithContext(ctx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Printf("Recovered from panic in Execute: %v", r)
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	result, err = commandFunc(tx, ctx)
	if err != nil {
		tx.Rollback()
		return result, fmt.Errorf("operation failed: %w", err)
	}

	if err := tx.Commit().Error; err != nil {
		return result, fmt.Errorf("commit failed: %w", err)
	}

	return result, nil
}
