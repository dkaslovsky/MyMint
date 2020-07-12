package delete

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dkaslovsky/MyMint/cmd/constants"
	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/spf13/cobra"
)

// CreateDeleteCmd generates the configuration for the delete subcommand.
// It can be attached to any upstream cobra command
func CreateDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Subcommand for deleting categories",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			confDir := os.Getenv(constants.ConfEnvVar)
			path := filepath.Join(confDir, constants.ManualCategoryFile)
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
