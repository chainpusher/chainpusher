package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/chainpusher/chainpusher/config"
	"github.com/spf13/cobra"
)

func NewMonitorCobraCommand() *cobra.Command {
	return &cobra.Command{
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
					log.Fatalf("failed to parse config: %v", err)
					cfg = &config.Config{}
				}
			}

			SetupLogger(cfg)

			monitor := NewMonitorCommand(cfg)
			monitor.Execute()
		},
	}
}

func NewRootCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "chainpusher",
		Short: "A CLI tool for pushing blockchain data",
		Long: "Chainpusher is a CLI tool for pushing blockchain data to a remote server. " +
			"Chainpusher can also monitor blockchain data and push it to a remote server.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
}

func RunCommand() {

	rootCmd := NewRootCommand()
	monitorCmd := NewMonitorCobraCommand()

	rootCmd.AddCommand(monitorCmd)
	rootCmd.PersistentFlags().String(
		"config",
		"c",
		"config file (default is $HOME/.chainpusher.yaml)",
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
