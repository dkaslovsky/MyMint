package dump

import (
	"fmt"

	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/ledger"
	"github.com/spf13/cobra"
)

// Options are command options
type Options struct {
	Db string
}

// CreateDumpCmd generates the configuration for the dump subcommand
func CreateDumpCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "dump",
		Short: "Prints the ledger to the console",
		RunE: func(cmd *cobra.Command, args []string) error {

			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				return err
			}
			defer db.Close()

			scanner, err := db.
				From(ledger.TableName).
				Executor().
				Scanner()
			if err != nil {
				return err
			}
			defer scanner.Close()

			for scanner.Next() {
				r := &ledger.Row{}
				err = scanner.ScanStruct(r)
				if err != nil {
					return err
				}
				fmt.Println(r)
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
	flags.StringVarP(&opts.Db, "database", "", sqlite.GetDbPath(), "Name of database")
}
