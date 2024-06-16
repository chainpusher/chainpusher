package monitor

import (
	"github.com/chainpusher/blockchain/service"
	"github.com/chainpusher/chainpusher/config"
)

type Ctx struct {
	Config *config.Config

	Channel chan interface{}

	Listeners []service.BlockListener

	Movement Movement
}
