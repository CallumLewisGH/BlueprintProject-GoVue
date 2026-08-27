package task

import (
	"context"
)

type WorkTask[R any] struct {
	Result   R
	err      error
	done     chan struct{}
	ctx      context.Context
	workFunc func(context.Context) (R, error)
}

func NewWorkTask[R any](ctx context.Context, workFunc func(context.Context) (R, error)) *WorkTask[R] {
	task := &WorkTask[R]{
		ctx:      ctx,
		done:     make(chan struct{}),
		workFunc: workFunc,
	}
	task.ExecuteAsync()
	return task
}

func (t *WorkTask[R]) ExecuteAsync() *WorkTask[R] {
	go func() {
		defer close(t.done)
		t.Result, t.err = t.workFunc(t.ctx)
	}()
	return t
}

func (t *WorkTask[R]) Await() (R, error) {
	<-t.done
	return t.Result, t.err
}
