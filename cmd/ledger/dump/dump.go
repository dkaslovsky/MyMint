package dump

import (
	"fmt"
	"strings"

	"github.com/dkaslovsky/MyMint/cmd/constants"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
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
				From(constants.TableName).
				Executor().
				Scanner()
			if err != nil {
				return err
			}
			defer scanner.Close()

			for scanner.Next() {
				r := &row{}
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
	flags.StringVarP(&opts.Db, "database", "", constants.DefaultDb, "Name of database")
}

type row struct {
	ID          int64       `db:"id"`
	Date        string      `db:"Date"`
	Amount      float64     `db:"Amount"`
	Description string      `db:"Description"`
	Category    interface{} `db:"Category"`
}

func (r *row) String() (s string) {
	fields := []string{
		fmt.Sprintf("ID: %d", r.ID),
		fmt.Sprintf("Date: %s", r.Date),
		fmt.Sprintf("Amount: %f", r.Amount),
		fmt.Sprintf("Description: %s", r.Description),
		fmt.Sprintf("Category: %v", r.Category),
	}
	return strings.Join(fields, " | ")
}
