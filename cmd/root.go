package cmd

import (
	"github.com/dkaslovsky/MyMint/cmd/ledger"
	"github.com/dkaslovsky/MyMint/cmd/source"
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
		ledger.CreateLedgerCmd(),
		source.CreateSourceCmd(),
	)
	return cmd.Execute()
}
