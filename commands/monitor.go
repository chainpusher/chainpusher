package commands

import (
	"sync"

	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/blockchain/service"
	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/monitor"
	"github.com/sirupsen/logrus"
)

type MonitorCommand struct {
	ctx *monitor.Ctx

	watchers []*monitor.PlatformWatcher
}

func (m *MonitorCommand) Execute() error {
	logrus.Trace("Executing monitor command")
	m.Start()

	return nil
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

func (m *MonitorCommand) Start() {
	logrus.Tracef("Starting monitor command")
	var wg sync.WaitGroup

	platforms := service.GetAllPlatform()
	for _, platform := range platforms {
		wg.Add(1)

		m.StartPlatformWithWaitGroup(platform, &wg)
	}

	wg.Wait()
}

func (m *MonitorCommand) Stop() {

}

func NewMonitorCommand(c *config.Config) *MonitorCommand {
	channel := make(chan interface{}, 10000)

	w := monitor.NewBlockLoggingWatcher(channel, c.BlockLoggingFile)
	if w != nil {
		logrus.Debug("Block logging watcher created")
		w.Start()
	}

	ctx := &monitor.Ctx{
		Config:  c,
		Channel: channel,
	}

	return &MonitorCommand{
		ctx:      ctx,
		watchers: []*monitor.PlatformWatcher{},
	}

}
