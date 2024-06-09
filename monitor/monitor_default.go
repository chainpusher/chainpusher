package monitor

import (
	"sync"

	"github.com/chainpusher/chainpusher/chain"
	"github.com/chainpusher/chainpusher/model"
)

type MonitorDefault struct {
	PlatformWatcherFactory PlatformWatcherFactory
}

func (md *MonitorDefault) OnTransaction(transaction model.Transaction) {
}

func (md *MonitorDefault) StartPlatform(platform model.Platform) {
	pfw := md.PlatformWatcherFactory.CreatePlatformWatcher(platform)
	pfw.Start()
}

func (md *MonitorDefault) StartPlatformWithWaitGroup(platform model.Platform, wg *sync.WaitGroup) {
	md.StartPlatform(platform)
	wg.Done()
}

func (md *MonitorDefault) Start() {
	var wg sync.WaitGroup

	platforms := chain.GetAllPlatform()
	for _, platform := range platforms {
		wg.Add(1)

		go md.StartPlatformWithWaitGroup(platform, &wg)
	}

	wg.Wait()
}

func (md *MonitorDefault) Stop() {

}

func NewDefaultMonitor(factory PlatformWatcherFactory) Monitor {
	return &MonitorDefault{
		PlatformWatcherFactory: factory,
	}
}
