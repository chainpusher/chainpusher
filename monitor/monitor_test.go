package monitor_test

import (
	"testing"
	"time"

	"github.com/chainpusher/chainpusher/model"
	"github.com/chainpusher/chainpusher/monitor"
)

func TestMonitor_Start(t *testing.T) {
	start := time.Now()
	m := monitor.NewDefaultMonitor(&PlatformWatcherTestingFactory{})
	m.Start()

	elapsed := time.Since(start)

	if elapsed < 200*time.Millisecond {
		t.Error("Monitor start should take at least 200ms")
	}
}

type PlatformWatcherTesting struct {
}

func (p *PlatformWatcherTesting) Start() {
	time.Sleep(200 * time.Millisecond)
}

func (p *PlatformWatcherTesting) Stop() {
}

type PlatformWatcherTestingFactory struct {
}

func (p *PlatformWatcherTestingFactory) CreatePlatformWatcher(platform model.Platform) monitor.PlatformWatcher {
	return &PlatformWatcherTesting{}
}
