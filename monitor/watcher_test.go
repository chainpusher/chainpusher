package monitor_test

import (
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/monitor"
	"github.com/sirupsen/logrus"
)

func TestWatcher_Move(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	now := time.Now()
	advanceTo := now.Add(2 * time.Second)

	watcher := monitor.NewWatcher(monitor.NewWatcherOptions{
		TimeRequiredForOneCycle: 10,
		IsOneTime:               true,
		TickHandler:             &MockTickHandler{},
	})

	go func() {
		time.Sleep(1 * time.Second)
		watcher.AdvanceTimeTo(advanceTo)
	}()

	watcher.StartContinuousMovement()
}

// test the Watcher normal flow
func TestWatcher_NormalFlow(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)

	watcher := monitor.NewWatcher(monitor.NewWatcherOptions{
		TimeRequiredForOneCycle: 1,
		IsOneTime:               true,
		TickHandler:             &MockTickHandler{},
	})
	watcher.StartContinuousMovement()

	if watcher.TheLastMovement > 1000*time.Millisecond {
		t.Error("Watcher should have taken less than 1 second to complete")
	}
}

// test the Watcher acceleration flow
func TestWatcher_AccelerationFlow(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	start := time.Now()

	advanceTo := start.Add(-8000 * time.Millisecond)

	watcher := monitor.NewWatcher(monitor.NewWatcherOptions{
		TimeRequiredForOneCycle: 10,
		IsOneTime:               true,
		TickHandler:             &MockTickHandler{advanceTo: advanceTo},
	})

	watcher.StartContinuousMovement()

	logrus.Tracef("The last movement took %v", watcher.TheLastMovement)

	if watcher.TheLastMovement < 1000*time.Millisecond {
		t.Error("Watcher should have taken greater than 1 second to complete")
	}
}

type MockTickHandler struct {
	advanceTo time.Time
}

func (m *MockTickHandler) Tick() {
	time.Sleep(1 * time.Second)

}

type MockTickHandlerResult struct {
	advanceTo time.Time
}

func (m *MockTickHandlerResult) GetBlockCreatedAt() time.Time {
	if !m.advanceTo.IsZero() {
		return m.advanceTo
	}
	return time.Now()
}

func (m *MockTickHandlerResult) GetBlock() interface{} {
	return nil
}
