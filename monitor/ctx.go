package monitor

import "github.com/chainpusher/chainpusher/config"

type Ctx struct {
	Config *config.Config

	Channel chan interface{}
}
