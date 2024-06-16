package main

import (
	"github.com/chainpusher/chainpusher/commands"
	"github.com/chainpusher/chainpusher/monitor"
	"github.com/sirupsen/logrus"
)

type DefaultCommandListener struct {
}

func (l *DefaultCommandListener) ConfigLoaded(ctx *monitor.Ctx) {
	logrus.Tracef("Config loaded %v", *ctx)
}

func main() {
	commands.RunCommandWithOptions(commands.MonitorCommandOptions{
		Listener: &DefaultCommandListener{},
	})
}
