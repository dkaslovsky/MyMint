package cmd

import (
	"github.com/dkaslovsky/MyMint/cmd/csv"
	"github.com/dkaslovsky/MyMint/cmd/table"
	"github.com/spf13/cobra"
)

// Run starts, configures, and executes the cli interface
func Run() error {
	cmd := &cobra.Command{
		Use:           "mymint",
		Short:         "root mymint command",
		Long:          `mymint persists personal finance data`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.AddCommand(
		csv.CreateCsvCmd(),
		table.CreateTableCmd(),
	)
	return cmd.Execute()
}
