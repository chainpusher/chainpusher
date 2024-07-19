package commands_test

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCommandParameter(t *testing.T) {

	rootCmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	subCommand := &cobra.Command{
		Use: "subcommand",
		Run: func(cmd *cobra.Command, args []string) {

			_, _ = cmd.Flags().GetBool("logging")
		},
	}
	subCommand.PersistentFlags().Bool("logging", false, "Enable logging")

	rootCmd.AddCommand(subCommand)
	rootCmd.SetArgs([]string{"subcommand", "--logging"})
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
