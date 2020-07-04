package add

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dkaslovsky/MyMint/cmd/constants"
	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/spf13/cobra"
)

// CreateAddCmd generates the configuration for the add subcommand.
// It can be attached to any upstream cobra command
func CreateAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Subcommand for adding categories",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			confDir := os.Getenv(constants.ConfEnvVar)
			path := filepath.Join(confDir, constants.CategoryFile)
			categories, err := category.LoadCategories(path)
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
