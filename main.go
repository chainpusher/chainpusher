package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chainpusher/chainpusher/commands"
	"github.com/chainpusher/chainpusher/config"
	"github.com/spf13/cobra"
)

func main() {
	var configArg string

	rootCmd := &cobra.Command{
		Use:   "chainpusher",
		Short: "A CLI tool for pushing blockchain data",
		Long: "Chainpusher is a CLI tool for pushing blockchain data to a remote server. " +
			"Chainpusher can also monitor blockchain data and push it to a remote server.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	monitorCmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor blockchain data",
		Run: func(cmd *cobra.Command, args []string) {

			defer func() {
				if r := recover(); r != nil {
					log.Println("Recovered in f", r)
				}
			}()

			c, err := config.ParseConfigFromYaml(configArg)
			if err != nil {
				log.Fatalf("failed to parse config: %v", err)

				return
			}

			monitor := commands.NewMonitorCommand(c)
			monitor.Execute()
		},
	}

	rootCmd.AddCommand(monitorCmd)
	rootCmd.PersistentFlags().StringVarP(
		&configArg,
		"config",
		"c",
		"",
		"config file (default is $HOME/.chainpusher.yaml)",
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("cmd error", err)
		os.Exit(1)
	}
}
