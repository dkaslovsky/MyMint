package csv

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dkaslovsky/MyMint/pkg/conf"

	"github.com/dkaslovsky/MyMint/pkg/category"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/parse"
	"github.com/dkaslovsky/MyMint/pkg/source"
	"github.com/doug-martin/goqu/v9"
	"github.com/spf13/cobra"
)

const (
	// categoryMatchField is the field to use for assigning categories
	categoryMatchField = "Description"
	// categoryField is the field for an assigned category
	categoryField = "Category"
)

// Options are options for configuring the csv command
type Options struct {
	Db     string
	Source string
}

// CreateCsvCmd generates the configuration for the csv subcommand
func CreateCsvCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "csv",
		Short: "Persist records from a csv file to a table",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			csvPath := args[0]

			sourcePath := filepath.Join(conf.Config.DataSourcePath, opts.Source)
			ext := filepath.Ext(sourcePath)
			if ext == "" {
				sourcePath += ".json"
			}
			ds, err := source.LoadDataSource(sourcePath)
			if err != nil {
				return err
			}

			csvRowParser, err := ds.GenerateCsvRowParser()
			if err != nil {
				return err
			}

			csvFile, err := os.Open(csvPath)
			if err != nil {
				return err
			}
			defer csvFile.Close()

			csvReader := parse.ReadCsvWithoutHeader
			if ds.Csv.Header {
				csvReader = parse.ReadCsvWithHeader
			}
			csvRows, err := csvReader(csvFile, csvRowParser)
			if err != nil {
				return err
			}

			keywordCatMap, err := category.LoadKeywordCatMap(conf.Config.KeywordCategoryFilePath)
			if err != nil {
				return err
			}
			for _, row := range csvRows {
				r := row.(goqu.Record)
				r[categoryField] = nil
				cat, err := keywordCatMap.GetFromRecord(r, categoryMatchField)
				if err != nil {
					return err
				}
				if cat != "" {
					r[categoryField] = cat
				}
				fmt.Println(row)
			}

			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				return err
			}
			defer db.Close()
			numInserted, lastID, err := db.InsertRows(ds.Table, csvRows)
			if err != nil {
				return err
			}

			log.Printf("Inserted [%d] rows ending with id [%d]", numInserted, lastID)
			return nil
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *Options) {
	flags := cmd.Flags()
	flags.StringVarP(&opts.Db, "database", "d", conf.Config.DefaultSqliteDbPath, "Name of database")
	flags.StringVarP(&opts.Source, "source", "s", "", "Path to datasource definition file")
	cobra.MarkFlagRequired(flags, "source")
}
