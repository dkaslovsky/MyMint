package cmd

import (
	"github.com/dkaslovsky/MyMint/cmd/initialize"
	"github.com/spf13/cobra"
)

// Run starts, configures, and executes the cli interface
func Run() error {
	cmd := &cobra.Command{
		Use:           "mymint",
		Short:         "Root mymint command",
		Long:          `mymint persists personal finance data`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.AddCommand(
		initialize.CreateInitCmd(),
	)
	return cmd.Execute()
}
