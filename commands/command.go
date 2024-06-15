package commands

import (
	"fmt"
	"github.com/chainpusher/blockchain/service"
	monitor2 "github.com/chainpusher/chainpusher/monitor"
	"github.com/sirupsen/logrus"
	"log"
	"os"

	"github.com/chainpusher/chainpusher/config"
	"github.com/spf13/cobra"
)

type MonitorCommandOptions struct {
	listeners []service.BlockListener
}

func NewMonitorCobraCommand(options MonitorCommandOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor blockchain data",
		Run: func(cmd *cobra.Command, args []string) {

			var cfg *config.Config

			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered in f", r)
				}
			}()

			p, err := cmd.Flags().GetString("config")

			if err != nil {
				cfg = &config.Config{}
			} else {
				cfg, err = config.ParseConfigFromYaml(p)
				if err != nil {
					cfg = &config.Config{}
					logrus.Errorf("failed to parse config: %v", err)
				}
			}

			isTesting, err := cmd.Flags().GetBool("test")
			if err == nil {
				cfg.IsTesting = isTesting
			}

			cfg.BlockLoggingFile, _ = cmd.Flags().GetString("block-file")

			SetupLogger(cfg)

			ctx := monitor2.Ctx{
				Config:    cfg,
				Listeners: options.listeners,
			}

			monitor := NewMonitorCommand(&ctx)
			err = monitor.Execute()
			if err != nil {
				return
			}
		},
	}

	cmd.PersistentFlags().StringP("block-file", "b", "", "File to write raw blockchain data to")
	cmd.PersistentFlags().StringP("trx-file", "t", "", "File to write transactions to")
	cmd.PersistentFlags().BoolP("test", "x", false, "Test mode")

	return cmd
}

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chainpusher",
		Short: "A CLI tool for pushing blockchain data",
		Long: "Chainpusher is a CLI tool for pushing blockchain data to a remote server. " +
			"Chainpusher can also monitor blockchain data and push it to a remote server.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.PersistentFlags().String(
		"config",
		"c",
		"config file (default is $HOME/.chainpusher.yaml)",
	)

	return cmd
}

func RunCommand() {

	RunCommandWithOptions(MonitorCommandOptions{
		listeners: []service.BlockListener{},
	})
}

func RunCommandWithOptions(options MonitorCommandOptions) {

	rootCmd := NewRootCommand()
	monitorCmd := NewMonitorCobraCommand(options)

	rootCmd.AddCommand(monitorCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
