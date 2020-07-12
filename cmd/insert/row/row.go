package row

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dkaslovsky/MyMint/cmd/constants"
	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/parse"
	"github.com/dkaslovsky/MyMint/pkg/source"
	"github.com/spf13/cobra"
)

// Options are options for configuring the row command
type Options struct {
	Db       string
	Source   string
	Category string
}

// CreateRowCmd generates the configuration for the row subcommand.
// It can be attached to any upstream cobra command
func CreateRowCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "row",
		Short: "Persist a single row from commandline input",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			row := args[0]

			confDir := os.Getenv(constants.ConfEnvVar)

			sourcePath := filepath.Join(confDir, constants.DataSourceDir, opts.Source)
			ext := filepath.Ext(sourcePath)
			if ext == "" {
				sourcePath += ".json"
			}
			ds, err := source.LoadDataSource(sourcePath)
			if err != nil {
				return err
			}

			categoryPath := filepath.Join(confDir, constants.ManualCategoryFile)
			categories, err := category.LoadCategories(categoryPath)
			if err != nil {
				return err
			}

			csvRowParser, err := ds.GenerateCsvRowParser()
			if err != nil {
				return err
			}

			if opts.Category != "" {
				if !categories.Contains(opts.Category) {
					return fmt.Errorf("unknown category [%s] must be added before it can be used", opts.Category)
				}
				row = fmt.Sprintf("%s,%s", row, opts.Category)
			}

			csvRow, err := parse.ReadCsvWithoutHeader(strings.NewReader(row), csvRowParser)
			if err != nil {
				return err
			}

			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				return err
			}
			defer db.Close()
			_, id, err := db.InsertRows(ds.Table, csvRow)
			if err != nil {
				return err
			}

			log.Printf("Inserted row with id [%d]", id)
			return nil
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *Options) {
	flags := cmd.Flags()
	flags.StringVarP(&opts.Db, "database", "d", constants.DefaultDb, "Name of database")
	flags.StringVarP(&opts.Source, "source", "s", "", "Datasource definition file name")
	flags.StringVarP(&opts.Category, "category", "c", "", "Category to associate with row")
	cobra.MarkFlagRequired(flags, "source")
}
