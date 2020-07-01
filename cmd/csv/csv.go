package csv

import (
	"os"

	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/parse"
	"github.com/dkaslovsky/MyMint/pkg/source"
	"github.com/spf13/cobra"
)

// Options are options for configuring the csv command
type Options struct {
	Path   string
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
		RunE: func(cmd *cobra.Command, args []string) error {

			ds, err := source.LoadDataSource(opts.Source)
			if err != nil {
				return err
			}

			csvRowParser, err := ds.GenerateCsvRowParser()
			if err != nil {
				return err
			}
			csvFile, err := os.Open(opts.Path)
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
			err = db.InsertRows(ds.Table, csvRows)
			if err != nil {
				return err
			}

			return nil
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *Options) {
	flags := cmd.Flags()
	flags.StringVarP(&opts.Path, "path", "p", "", "Path to csv file")
	flags.StringVarP(&opts.Source, "source", "s", "", "Path to datasource definition file")
	flags.StringVarP(&opts.Db, "database", "d", "mydb.db", "Name of database")
	cobra.MarkFlagRequired(flags, "path")
	cobra.MarkFlagRequired(flags, "source")
}
