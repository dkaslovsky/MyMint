package source

import (
	"github.com/dkaslovsky/MyMint/cmd/source/cat"
	"github.com/dkaslovsky/MyMint/cmd/source/list"
	"github.com/spf13/cobra"
)

// CreateSourceCmd generates the configuration for the source subcommand
func CreateSourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source",
		Short: "Subcommand for interacting with data source files",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	cmd.AddCommand(
		list.CreateListCmd(),
		cat.CreateCatCmd(),
	)
	return cmd
}
