package cmd

import (
	"github.com/dkaslovsky/MyMint/cmd/createtable"
	"github.com/dkaslovsky/MyMint/cmd/csv"
	"github.com/dkaslovsky/MyMint/cmd/dumptable"
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
		dumptable.CreateDumpTableCmd(),
		createtable.CreateCreateTableCmd(),
	)
	return cmd.Execute()
}
