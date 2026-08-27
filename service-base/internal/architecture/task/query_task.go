package task

import (
	"context"

	cqrs "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/cqrs"
	"gorm.io/gorm"
)

type QueryTask[R any] struct {
	Result    R
	err       error
	done      chan struct{}
	ctx       context.Context
	queryFunc func(*gorm.DB, context.Context) (R, error)
}

func NewQueryTask[R any](ctx context.Context, queryFunc func(*gorm.DB, context.Context) (R, error)) *QueryTask[R] {
	task := &QueryTask[R]{
		ctx:       ctx,
		done:      make(chan struct{}),
		queryFunc: queryFunc,
	}
	task.ExecuteQueryAsync()
	return task
}

func (t *QueryTask[R]) ExecuteQueryAsync() *QueryTask[R] {
	go func() {
		defer close(t.done)
		t.Result, t.err = cqrs.DbQuery(t.ctx, t.queryFunc)
	}()
	return t
}

func (t *QueryTask[R]) Await() (R, error) {
	<-t.done
	return t.Result, t.err
}
