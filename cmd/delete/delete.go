package delete

import (
	"log"

	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/spf13/cobra"
)

// Options are options for configuring the delete command
type Options struct {
	Db    string
	Table string
	IDCol string
}

// CreateDeleteCmd generates the configuration for the delete subcommand
func CreateDeleteCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete rows from a table",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				return err
			}
			defer db.Close()

			ids := make([]interface{}, len(args))
			for i, arg := range args {
				ids[i] = arg
			}

			numDeleted, err := db.DeleteRowsByID(opts.Table, opts.IDCol, ids...)
			if err != nil {
				return err
			}

			log.Printf("Deleted [%d] rows", numDeleted)
			return nil
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *Options) {
	flags := cmd.Flags()
	flags.StringVarP(&opts.Db, "database", "d", sqlite.GetDbPath(), "Name of database")
	flags.StringVarP(&opts.IDCol, "id", "i", "id", "Name of id column")
	flags.StringVarP(&opts.Table, "table", "t", "", "table to delete from")
	cobra.MarkFlagRequired(flags, "table")
}
