package delete

import (
	"log"
	"strings"

	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/spf13/cobra"
)

// CreateDeleteCmd generates the configuration for the delete subcommand
func CreateDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a category from the keyword category mappings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := conf.Config.KeywordCategoryFilePath
			categories, err := category.LoadKeywordCategories(path)
			if err != nil {
				return err
			}
			key := strings.ToLower(args[0])
			if !categories.Contains(key) {
				log.Printf("[%s] not in categories, nothing to do", key)
				return nil
			}
			categories.Delete(key)
			err = categories.Write(path)
			if err != nil {
				log.Printf("error writing categories, check contents of file [%s] manually", path)
				return err
			}
			log.Printf("deleted [%s] from categories", key)
			return nil
		},
	}
	return cmd
}
