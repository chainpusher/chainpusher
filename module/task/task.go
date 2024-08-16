package task

type Task interface {
	Start() error

	Stop() error

	Running() bool
}
