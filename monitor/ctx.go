package monitor

import (
	"github.com/chainpusher/blockchain/model"
	"github.com/chainpusher/blockchain/service"
	"github.com/chainpusher/chainpusher/config"
)

type Ctx struct {
	Config *config.Config

	Channel chan interface{}

	Listeners []service.BlockListener

	Movement Movement

	Platforms []model.Platform
}

func (c *Ctx) AddBlockListener(listener service.BlockListener) {
	c.Listeners = append(c.Listeners, listener)
}

func (c *Ctx) GetPlatforms() []model.Platform {
	if len(c.Platforms) == 0 {
		c.Platforms = []model.Platform{
			model.PlatformEthereum,
			model.PlatformTron,
		}
	}

	return c.Platforms
}

func NewContext(cfg *config.Config) *Ctx {
	return &Ctx{
		Config:  cfg,
		Channel: make(chan interface{}),
	}
}
