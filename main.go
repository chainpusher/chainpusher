package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chainpusher/chainpusher/commands"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "chainpusher",
		Short: "A CLI tool for pushing blockchain data",
	}

	monitorCmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor blockchain data",
		Run: func(cmd *cobra.Command, args []string) {
			// Add your monitor logic here
			fmt.Println("Monitoring blockchain data...")

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
				}
			}()

			monitor := commands.NewMonitorCommand()
			monitor.Execute()

			log.Println("End of monitoring")
		},
	}

	rootCmd.AddCommand(monitorCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("cmd error", err)
		os.Exit(1)
	}
}
