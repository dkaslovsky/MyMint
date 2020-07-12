package ledger

import (
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/spf13/cobra"
)

var (
	// table name for ledger
	table = "ledger"
	// table schema for ledger
	schema = sqlite.Schema{
		"id":          "INTEGER PRIMARY KEY",
		"Category":    "TEXT",
		"Date":        "TEXT",
		"Amount":      "REAL",
		"Description": "TEXT",
	}
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
		CreateInitCmd(),
	)
	return cmd
}
