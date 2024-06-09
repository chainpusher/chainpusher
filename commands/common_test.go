package commands_test

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestCommandParameter(t *testing.T) {

	expectSubCommandExecuted := true
	expectLoggingToBeTrue := true

	wasSubcommandExecute := false
	wasLoggingTrue := false

	rootCmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	subCommand := &cobra.Command{
		Use: "subcommand",
		Run: func(cmd *cobra.Command, args []string) {
			wasSubcommandExecute = true
			wasLoggingTrue, _ = cmd.Flags().GetBool("logging")
		},
	}
	subCommand.PersistentFlags().Bool("logging", false, "Enable logging")

	rootCmd.AddCommand(subCommand)
	rootCmd.SetArgs([]string{"subcommand", "--logging"})
	rootCmd.Execute()

	if expectSubCommandExecuted != wasSubcommandExecute {
		t.Errorf("Expected subcommand to be executed")
	}

	if expectLoggingToBeTrue != wasLoggingTrue {
		t.Errorf("Expected logging to be true")
	}
}
