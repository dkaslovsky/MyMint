package csv

import (
	"os"

	"github.com/dkaslovsky/MyMint/pkg/data"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/parse"
	"github.com/spf13/cobra"
)

// Options are options for configuring the csv command
type Options struct {
	Path  string
	Db    string
	Table string
}

// CreateCsvCmd generates the configuration for the csv subcommand.
// It can be attached to any upstream cobra command
func CreateCsvCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "csv",
		Short: "Persist records from a csv file",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				return err
			}
			defer db.Close()

			csvFile, err := os.Open(opts.Path)
			if err != nil {
				return err
			}
			defer csvFile.Close()

			csvRows, err := parse.ReadCSV(csvFile, data.ExampleTableCSVParser) // FIX ME!
			if err != nil {
				return err
			}

			err = db.InsertRows(opts.Table, csvRows)
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
	flags.StringVarP(&opts.Db, "database", "d", "", "Name of database")
	flags.StringVarP(&opts.Table, "table", "t", "", "Name of table")
	cobra.MarkFlagRequired(flags, "database")
	cobra.MarkFlagRequired(flags, "table")
}
