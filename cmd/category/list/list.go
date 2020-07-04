package list

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dkaslovsky/MyMint/cmd/constants"
	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/spf13/cobra"
)

// CreateListCmd generates the configuration for the list subcommand.
// It can be attached to any upstream cobra command
func CreateListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "Subcommand for listing categories",
		RunE: func(cmd *cobra.Command, args []string) error {
			confDir := os.Getenv(constants.ConfEnvVar)
			path := filepath.Join(confDir, constants.CategoryFile)
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
