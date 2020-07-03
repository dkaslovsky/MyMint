package row

import (
	"fmt"
	"strings"

	"github.com/dkaslovsky/MyMint/cmd/defaults"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/parse"
	"github.com/dkaslovsky/MyMint/pkg/source"
	"github.com/spf13/cobra"
)

// Options are options for configuring the row command
type Options struct {
	Path   string
	Db     string
	Source string
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

			ds, err := source.LoadDataSource(opts.Source)
			if err != nil {
				return err
			}
			fmt.Println(ds)

			csvRowParser, err := ds.GenerateCsvRowParser()
			if err != nil {
				return err
			}

			csvRows, err := parse.ReadCsvWithoutHeader(strings.NewReader(row), csvRowParser)
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
	flags.StringVarP(&opts.Db, "database", "d", defaults.DefaultDB, "Name of database")
	flags.StringVarP(&opts.Source, "source", "s", "", "Path to datasource definition file")
	cobra.MarkFlagRequired(flags, "source")
}
