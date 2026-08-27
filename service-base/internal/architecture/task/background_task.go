package task

import (
	"context"
	"log"
)

type BackgroundTask struct {
	ctx      context.Context
	taskFunc func(context.Context) error
}

func NewBackgroundTask(ctx context.Context, taskFunc func(context.Context) error) {
	bgCtx := context.WithoutCancel(ctx)

	t := &BackgroundTask{
		ctx:      bgCtx,
		taskFunc: taskFunc,
	}

	t.Execute()
}

func (t *BackgroundTask) Execute() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("PANIC in BackgroundTask: %v", r)
			}
		}()

		if err := t.taskFunc(t.ctx); err != nil {
			log.Printf("Background Task Failed: %v", err)
		}
	}()
}
