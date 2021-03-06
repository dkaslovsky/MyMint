package ledger

import (
	"github.com/dkaslovsky/MyMint/cmd/ledger/add"
	"github.com/dkaslovsky/MyMint/cmd/ledger/category"
	"github.com/dkaslovsky/MyMint/cmd/ledger/dump"
	"github.com/spf13/cobra"
)

// CreateLedgerCmd generates the configuration for the ledger subcommand
func CreateLedgerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ledger",
		Short: "Ledger operations",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	cmd.AddCommand(
		dump.CreateDumpCmd(),
		add.CreateAddCmd(),
		category.CreateCategoryCmd(),
	)
	return cmd
}
