package monitor

import (
	"github.com/chainpusher/chainpusher/application"
	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/model"
)

type PlatformWatcherDefaultFactory struct {
	Config *config.Config
}

func (p *PlatformWatcherDefaultFactory) CreatePlatformWatcher(platform model.Platform) PlatformWatcher {
	transactionService := application.NewTransactionService(p.Config)

	switch platform {
	case model.PlatformTron:
		return NewPlatformWatcherTron(p.Config, transactionService)
	default:
		panic("Platform not supported")
	}

}

func NewPlatformWatcherDefaultFactory(config *config.Config) *PlatformWatcherDefaultFactory {
	return &PlatformWatcherDefaultFactory{
		Config: config,
	}
}
