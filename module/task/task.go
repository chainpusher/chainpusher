package task

type Task interface {
	Start() error

	Stop() error

	Running() bool
}

type FunctionTask struct {
	handler   FunctionTaskHandler
	stop      chan bool
	stopped   chan bool
	isRunning bool
}

func (f *FunctionTask) Running() bool {
	return f.isRunning
}

func (f *FunctionTask) Stop() error {
	f.stop <- true

	return nil
}

func (f *FunctionTask) Start() error {
	f.isRunning = true
	var err error
	if err = f.handler(f.stop); err != nil {

	}
	f.isRunning = false
	//f.stopped <- true
	return err
}

type FunctionTaskHandler func(stop chan bool) error

func NewFunctionTask(handler FunctionTaskHandler) *FunctionTask {
	stop := make(chan bool, 1)
	stopped := make(chan bool, 1)
	return &FunctionTask{handler: handler, stop: stop, stopped: stopped}
}
