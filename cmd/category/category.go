package category

import (
	"github.com/dkaslovsky/MyMint/cmd/category/add"
	"github.com/dkaslovsky/MyMint/cmd/category/delete"
	"github.com/dkaslovsky/MyMint/cmd/category/list"
	"github.com/spf13/cobra"
)

// CreateCategoryCmd generates the configuration for the category subcommand.
// It can be attached to any upstream cobra command
func CreateCategoryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "category",
		Short: "Subcommand for managing categories",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	cmd.AddCommand(
		list.CreateListCmd(),
		add.CreateAddCmd(),
		delete.CreateDeleteCmd(),
	)
	return cmd
}
