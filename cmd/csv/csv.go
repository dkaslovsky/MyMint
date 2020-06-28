package csv

import (
	"log"
	"os"

	"github.com/dkaslovsky/MyMint/pkg/data"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/parse"
	"github.com/spf13/cobra"
)

// Options are options for configuring the csv command
type Options struct {
	Path string
	Db   string
}

// CreateCsvCmd generates the configuration for the csv subcommand.
// It can be attached to any upstream cobra command
func CreateCsvCmd() *cobra.Command {
	opts := Options{}

	cmd := &cobra.Command{
		Use:   "csv",
		Short: "Persist records from a csv file",
		RunE: func(cmd *cobra.Command, args []string) error {
			csvFile, err := os.Open(opts.Path)
			if err != nil {
				return err
			}
			defer csvFile.Close()
			csvRows, err := parse.ReadCSV(csvFile, data.ExampleTableCSVParser)
			if err != nil {
				return err
			}

			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			table := sqlite.NewTable("mytable", data.ExampleTableSchema)
			err = db.CreateTable(table)
			if err != nil {
				log.Fatal(err)
			}

			err = db.InsertRows(table, csvRows)
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
	flags.StringVarP(&opts.Db, "database", "d", "mydb.db", "Name of database")
}
