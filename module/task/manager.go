package task

import (
	"github.com/sirupsen/logrus"
	"sync"
)

type Manager struct {
	tasks []Task
	wg    *sync.WaitGroup
}

func (m *Manager) Add(t Task) {
	m.tasks = append(m.tasks, t)
}

func (m *Manager) Start() {
	for _, t := range m.tasks {
		m.wg.Add(1)
		go func(t Task) {
			if err := t.Start(); err != nil {
				logrus.Errorf("failed to start task: %v", err)
			}
			m.wg.Done()
		}(t)
	}
}

func (m *Manager) Wait() {
	m.wg.Wait()
}

func (m *Manager) Stop() {
	wg := sync.WaitGroup{}

	for _, t := range m.tasks {
		wg.Add(1)

		go func(t Task, wg *sync.WaitGroup) {
			if err := t.Stop(); err != nil {
				logrus.Errorf("failed to stop task: %v", err)
			}
			wg.Done()
		}(t, &wg)
	}

	wg.Wait()
}

func NewManager(tasks ...Task) *Manager {
	return &Manager{tasks: tasks, wg: &sync.WaitGroup{}}
}
