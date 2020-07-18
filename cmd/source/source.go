package source

import (
	"github.com/dkaslovsky/MyMint/cmd/source/category"
	"github.com/dkaslovsky/MyMint/cmd/source/dump"
	"github.com/dkaslovsky/MyMint/cmd/source/list"
	"github.com/spf13/cobra"
)

// CreateSourceCmd generates the configuration for the source subcommand
func CreateSourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source",
		Short: "Datasource file operations",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	cmd.AddCommand(
		list.CreateListCmd(),
		dump.CreateDumpCmd(),
		category.CreateCategoryCmd(),
	)
	return cmd
}
