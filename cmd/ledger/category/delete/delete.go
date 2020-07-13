package delete

import (
	"log"

	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/spf13/cobra"
)

// CreateDeleteCmd generates the configuration for the delete subcommand
func CreateDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a category from the ledger categories",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := conf.GetLedgerCategoryPath()
			categories, err := category.LoadCategories(path)
			if err != nil {
				return err
			}
			category := args[0]
			if !categories.Contains(category) {
				log.Printf("[%s] not in categories, nothing to do", category)
				return nil
			}
			categories.Delete(category)
			err = categories.Write(path)
			if err != nil {
				log.Printf("error writing categories, check contents of file [%s] manually", path)
				return err
			}
			log.Printf("deleted [%s] from categories", category)
			return nil
		},
	}
	return cmd
}
