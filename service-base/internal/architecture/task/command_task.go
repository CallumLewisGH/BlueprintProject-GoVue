package task

import (
	"context"

	cqrs "github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/architecture/cqrs"
	"gorm.io/gorm"
)

type CommandTask[R any] struct {
	Result      R
	err         error
	done        chan struct{}
	ctx         context.Context
	commandFunc func(*gorm.DB, context.Context) (R, error)
}

func NewCommandTask[R any](ctx context.Context, commandFunc func(*gorm.DB, context.Context) (R, error)) *CommandTask[R] {
	task := &CommandTask[R]{
		ctx:         ctx,
		done:        make(chan struct{}),
		commandFunc: commandFunc,
	}
	task.ExecuteCommandAsync()
	return task
}

func (t *CommandTask[R]) ExecuteCommandAsync() *CommandTask[R] {
	go func() {
		defer close(t.done)
		t.Result, t.err = cqrs.DbExecute(t.ctx, t.commandFunc)
	}()
	return t
}

func (t *CommandTask[R]) Await() (R, error) {
	<-t.done
	return t.Result, t.err
}
