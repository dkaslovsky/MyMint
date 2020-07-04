package source

import (
	"github.com/dkaslovsky/MyMint/cmd/source/cat"
	"github.com/dkaslovsky/MyMint/cmd/source/list"
	"github.com/spf13/cobra"
)

// CreateSourceCmd generates the configuration for the source subcommand.
// It can be attached to any upstream cobra command
func CreateSourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source",
		Short: "Subcommand for source operations",
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
