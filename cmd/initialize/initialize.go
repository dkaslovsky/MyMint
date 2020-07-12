package initialize

import (
	"github.com/dkaslovsky/MyMint/cmd/constants"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/spf13/cobra"
)

// Options are command options
type Options struct {
	Db string
}

// CreateInitCmd generates the configuration for the init subcommand
func CreateInitCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a database with a ledger",
		//Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				return err
			}
			defer db.Close()

			err = db.CreateTable(table, schema)
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
	flags.StringVarP(&opts.Db, "database", "", constants.DefaultDb, "Name of database")
}

var (
	// table name for ledger
	table = "ledger"
	// table schema for ledger
	schema = sqlite.Schema{
		"id":          "INTEGER PRIMARY KEY",
		"Category":    "TEXT",
		"Date":        "TEXT",
		"Amount":      "REAL",
		"Description": "TEXT",
	}
)
