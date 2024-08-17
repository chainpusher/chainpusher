package task_test

import (
	"github.com/chainpusher/chainpusher/module/task"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	executed := false
	m := task.NewManager(task.NewFunctionTask(func(stop chan bool) error {
		time.Sleep(10 * time.Millisecond)
		executed = true
		return nil
	}))

	m.Start()
	m.Wait()
	assert.True(t, executed)
}

func TestManagerStop(t *testing.T) {
	executed := false
	m := task.NewManager(task.NewFunctionTask(func(stop chan bool) error {
		time.Sleep(10 * time.Millisecond)
		executed = true
		return nil
	}))

	m.Start()
	m.Stop()
	assert.False(t, executed)
}
