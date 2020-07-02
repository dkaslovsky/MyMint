package dump

import (
	"fmt"

	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/spf13/cobra"
)

// Options are options for configuring the dump command
type Options struct {
	Db    string
	Table string
}

// CreateDumpCmd generates the configuration for the dump subcommand.
// It can be attached to any upstream cobra command
func CreateDumpCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "dump",
		Short: "Dump table contents to console",
		RunE: func(cmd *cobra.Command, args []string) error {

			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				return err
			}
			defer db.Close()

			scanner, err := db.
				From(opts.Table).
				Executor().
				Scanner()
			if err != nil {
				return err
			}
			defer scanner.Close()

			type Row struct {
				ID   int64  `db:"id"`
				Name string `db:"name"`
				Val  int64  `db:"val"`
			}

			for scanner.Next() {
				row := Row{}
				err = scanner.ScanStruct(&row)
				if err != nil {
					return err
				}
				fmt.Println(row)
			}

			if scanner.Err() != nil {
				fmt.Println(scanner.Err())
			}

			return nil
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *Options) {
	flags := cmd.Flags()
	flags.StringVarP(&opts.Db, "database", "d", "mydb.db", "Name of database")
	flags.StringVarP(&opts.Table, "table", "t", "", "Table to validate")
	cobra.MarkFlagRequired(flags, "table")
}
