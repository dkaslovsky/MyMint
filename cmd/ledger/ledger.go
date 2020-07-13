package ledger

import (
	"github.com/dkaslovsky/MyMint/cmd/ledger/add"
	"github.com/dkaslovsky/MyMint/cmd/ledger/dump"
	"github.com/dkaslovsky/MyMint/cmd/ledger/initialize"
	"github.com/spf13/cobra"
)

// CreateLedgerCmd generates the configuration for the ledger subcommand
func CreateLedgerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ledger",
		Short: "Subcommands for interacting with the ledger",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	cmd.AddCommand(
		initialize.CreateInitCmd(),
		dump.CreateDumpCmd(),
		add.CreateAddCmd(),
	)
	return cmd
}
