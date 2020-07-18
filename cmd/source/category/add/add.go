package add

import (
	"log"

	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/spf13/cobra"
)

// Options are options for configuring the add command
type Options struct {
	Key string
	Val string
}

// CreateAddCmd generates the configuration for the add subcommand
func CreateAddCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a category to the keyword category mappings",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := conf.Config.KeywordCategoryFilePath
			categories, err := category.LoadKeywordCategories(path)
			if err != nil {
				return err
			}

			if categories.Contains(opts.Key) {
				log.Printf("key [%s] already in categories, delete and re-add to change the value", opts.Key)
				return nil
			}
			categories.Add(opts.Key, opts.Val)
			err = categories.Write(path)
			if err != nil {
				log.Printf("error writing categories, check contents of file [%s] manually", path)
				return err
			}
			log.Printf("added [%s: %s] to categories", opts.Key, opts.Val)
			return nil
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *Options) {
	flags := cmd.Flags()
	flags.StringVarP(&opts.Key, "key", "k", "", "Key (substring to match)")
	flags.StringVarP(&opts.Val, "value", "v", "", "Value (category to assign")
	cobra.MarkFlagRequired(flags, "key")
	cobra.MarkFlagRequired(flags, "value")
}
