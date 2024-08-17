package task_test

import (
	"github.com/chainpusher/chainpusher/module/task"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFunctionTask(t *testing.T) {
	executed := false
	task1 := task.NewFunctionTask(func(stop chan bool) error {
		time.Sleep(100 * time.Millisecond)
		executed = true
		return nil
	})
	err := task1.Start()
	assert.Nil(t, err)
	assert.True(t, executed)
}

func TestFunctionTaskStop(t *testing.T) {
	executed := false
	task1 := task.NewFunctionTask(func(stop chan bool) error {
		time.Sleep(100 * time.Millisecond)

		select {
		case <-stop:
			return nil
		default:
			executed = true
			return nil
		}
	})
	go func() {
		time.Sleep(10 * time.Millisecond)
		_ = task1.Stop()
	}()
	err := task1.Start()
	assert.Nil(t, err)
	assert.False(t, executed)
}
