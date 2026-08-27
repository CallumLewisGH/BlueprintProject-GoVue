package task

type Task[R any] interface {
	Await() (R, error)
}
