package commands

import (
	"github.com/chainpusher/chainpusher/module/task"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func GatewayCommand(cmd *cobra.Command, args []string) {
	cfg, err := GetConfig(cmd)
	if err != nil {
		logrus.Errorf("failed to load config: %v", err)
		return
	}

	m := task.NewManager()
	m.Start()
}

func NewGatewayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gateway",
		Short: "Starts the gateway",
		Run:   GatewayCommand,
	}

	SetupConfigPathFlag(cmd)

	return cmd
}
