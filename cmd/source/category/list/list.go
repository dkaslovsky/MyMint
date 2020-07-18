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
		Short: "List the keyword category mapping",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := conf.Config.KeywordCategoryFilePath
			categories, err := category.LoadKeywordCategories(path)
			if err != nil {
				return err
			}
			fmt.Println(categories)
			return nil
		},
	}
	return cmd
}
