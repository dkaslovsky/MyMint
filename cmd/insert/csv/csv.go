package csv

import (
	"log"
	"os"

	"github.com/dkaslovsky/MyMint/cmd/constants"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/parse"
	"github.com/dkaslovsky/MyMint/pkg/source"
	"github.com/spf13/cobra"
)

// Options are options for configuring the csv command
type Options struct {
	Db     string
	Source string
}

// CreateCsvCmd generates the configuration for the csv subcommand.
// It can be attached to any upstream cobra command
func CreateCsvCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "csv",
		Short: "Persist records from a csv file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			csvPath := args[0]

			ds, err := source.LoadDataSource(opts.Source)
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
	flags.StringVarP(&opts.Db, "database", "d", constants.DefaultDb, "Name of database")
	flags.StringVarP(&opts.Source, "source", "s", "", "Path to datasource definition file")
	cobra.MarkFlagRequired(flags, "source")
}
