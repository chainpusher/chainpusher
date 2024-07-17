package sys

type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) AddTask(task Task) {
	tm.tasks = append(tm.tasks, task)
}

func (tm *TaskManager) StartAll() {
	for _, task := range tm.tasks {
		go func() {
			_ = task.Start()
		}()
	}
}

func (tm *TaskManager) StopAll() {
	for _, task := range tm.tasks {
		_ = task.Stop()
	}
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}
