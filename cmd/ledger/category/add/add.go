package add

import (
	"log"

	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/spf13/cobra"
)

// CreateAddCmd generates the configuration for the add subcommand
func CreateAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a category to the ledger categories",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := conf.Config.LedgerCategoryFilePath
			categories, err := category.LoadLedgerCategories(path)
			if err != nil {
				return err
			}
			category := args[0]
			if categories.Contains(category) {
				log.Printf("[%s] already in categories, nothing to do", category)
				return nil
			}
			categories.Add(category)
			err = categories.Write(path)
			if err != nil {
				log.Printf("error writing categories, check contents of file [%s] manually", path)
				return err
			}
			log.Printf("added [%s] to categories", category)
			return nil
		},
	}
	return cmd
}
