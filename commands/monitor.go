package commands

import (
	"log"

	"github.com/chainpusher/chainpusher/config"
	"github.com/chainpusher/chainpusher/monitor"
	"github.com/fbsobreira/gotron-sdk/pkg/client"
	"google.golang.org/grpc"
)

type MonitorCommand struct {
	Client  *client.GrpcClient
	Monitor monitor.Monitor

	monitors []monitor.PlatformWatcher
}

func (m *MonitorCommand) Execute() error {

	m.Monitor.Start()

	return nil
}

func NewMonitorCommand(c *config.Config) *MonitorCommand {
	client := client.NewGrpcClient("")
	err := client.Start(grpc.WithInsecure())

	if err != nil {
		log.Panic(err)
		panic(err)
	}

	pff := monitor.NewPlatformWatcherDefaultFactory(c)
	m := monitor.NewDefaultMonitor(pff)

	return &MonitorCommand{
		Client:  client,
		Monitor: m,
	}

}
