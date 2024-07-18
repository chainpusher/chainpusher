package commands

import (
	"sync"

	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/blockchain/service"
	"github.com/chainpusher/chainpusher/monitor"
	"github.com/sirupsen/logrus"
)

type MonitorCommand struct {
	ctx *monitor.Ctx

	watchers []*monitor.PlatformWatcher

	running bool
}

func (m *MonitorCommand) AddWatcher(watcher *monitor.PlatformWatcher) {
	m.watchers = append(m.watchers, watcher)
}

func (m *MonitorCommand) StartPlatformWithWaitGroup(platform model.Platform, wg *sync.WaitGroup) {
	watcher, err := monitor.NewPlatformWatcher(m.ctx, platform)
	if err != nil {
		logrus.Errorf("Error creating watcher for platform %s: %v", platform, err)
		wg.Done()
		return
	}
	m.AddWatcher(watcher)

	go func(watcher *monitor.PlatformWatcher, wg *sync.WaitGroup) {
		logrus.Tracef("Starting watcher for platform %s", platform)
		watcher.Start()
		wg.Done()
	}(watcher, wg)

}

func (m *MonitorCommand) Start() error {
	logrus.Tracef("Starting monitor command")
	var wg sync.WaitGroup

	platforms := service.GetAllPlatform()
	for _, platform := range platforms {
		wg.Add(1)

		m.StartPlatformWithWaitGroup(platform, &wg)
	}

	m.running = true

	wg.Wait()

	return nil
}

func (m *MonitorCommand) Stop() error {
	return nil
}

func (m *MonitorCommand) Running() bool {
	return m.running
}

func NewMonitorCommand(ctx *monitor.Ctx) *MonitorCommand {
	c := ctx.Config
	channel := make(chan interface{}, 10000)

	w := monitor.NewBlockLoggingWatcher(channel, c.BlockLoggingFile)
	if w != nil {
		logrus.Debug("Block logging watcher created")
		w.Start()
	}

	ctx.Config = c
	ctx.Channel = channel

	return &MonitorCommand{
		ctx:      ctx,
		watchers: []*monitor.PlatformWatcher{},
	}

}
