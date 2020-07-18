package category

import (
	"github.com/dkaslovsky/MyMint/cmd/source/category/add"
	"github.com/dkaslovsky/MyMint/cmd/source/category/delete"
	"github.com/dkaslovsky/MyMint/cmd/source/category/list"
	"github.com/spf13/cobra"
)

// CreateCategoryCmd generates the configuration for the category subcommand
func CreateCategoryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "category",
		Short: "Interact with the keyword category mapping for datasources",
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
