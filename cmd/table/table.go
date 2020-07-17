package table

import (
	"github.com/dkaslovsky/MyMint/cmd/table/create"
	"github.com/dkaslovsky/MyMint/cmd/table/csv"
	"github.com/spf13/cobra"
)

// CreateTableCmd generates the configuration for the table subcommand
func CreateTableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "table",
		Short: "Table operations",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	cmd.AddCommand(
		create.CreateCreateCmd(),
		csv.CreateCsvCmd(),
	)
	return cmd
}
