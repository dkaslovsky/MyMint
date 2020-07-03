package insert

import (
	"github.com/dkaslovsky/MyMint/cmd/insert/csv"
	"github.com/dkaslovsky/MyMint/cmd/insert/row"
	"github.com/spf13/cobra"
)

// CreateInsertCmd generates the configuration for the insert subcommand.
// It can be attached to any upstream cobra command
func CreateInsertCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert",
		Short: "Subcommand for insert operations",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	cmd.AddCommand(
		csv.CreateCsvCmd(),
		row.CreateRowCmd(),
	)
	return cmd
}
