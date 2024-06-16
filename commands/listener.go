package commands

import (
	"github.com/chainpusher/chainpusher/monitor"
)

type Listener interface {
	ConfigLoaded(ctx *monitor.Ctx)
}
