package sys

type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) Add(t Task) {
	tm.tasks = append(tm.tasks, t)
}

func (tm *TaskManager) Start() {
	for _, t := range tm.tasks {
		go tm.start(t)
	}
}

func (tm *TaskManager) start(task Task) {
	_ = task.Start()
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}
