package table

import (
	"github.com/dkaslovsky/MyMint/cmd/table/create"
	"github.com/dkaslovsky/MyMint/cmd/table/dump"
	"github.com/spf13/cobra"
)

// CreateTableCmd generates the configuration for the table subcommand.
// It can be attached to any upstream cobra command
func CreateTableCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "table",
		Short: "Subcommand for table operations",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	cmd.AddCommand(
		dump.CreateDumpCmd(),
		create.CreateCreateCmd(),
	)
	return cmd
}
