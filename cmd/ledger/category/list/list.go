package list

import (
	"fmt"

	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/spf13/cobra"
)

// CreateListCmd generates the configuration for the list subcommand
func CreateListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List the ledger categories",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := conf.Config.LedgerCategoryPath
			categories, err := category.LoadCategories(path)
			if err != nil {
				return err
			}
			fmt.Println(categories)
			return nil
		},
	}
	return cmd
}
