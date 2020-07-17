package category

import (
	"github.com/dkaslovsky/MyMint/cmd/ledger/category/add"
	"github.com/dkaslovsky/MyMint/cmd/ledger/category/delete"
	"github.com/dkaslovsky/MyMint/cmd/ledger/category/list"
	"github.com/spf13/cobra"
)

// CreateCategoryCmd generates the configuration for the category subcommand
func CreateCategoryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "category",
		Short: "Interact with ledger categories",
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
