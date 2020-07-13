package create

import (
	"path/filepath"

	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/source"
	"github.com/spf13/cobra"
)

// Options are options for configuring the create command
type Options struct {
	Db string
}

// CreateCreateCmd generates the configuration for the create subcommand
func CreateCreateCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new table from a datasource",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			path := filepath.Join(conf.GetDataSourcePath(), args[0])
			ext := filepath.Ext(path)
			if ext == "" {
				path += ".json"
			}

			ds, err := source.LoadDataSource(path)
			if err != nil {
				return err
			}

			db, err := sqlite.NewDb(opts.Db)
			if err != nil {
				return err
			}
			defer db.Close()

			err = db.CreateTable(ds.Table, ds.Schema)
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
	flags.StringVarP(&opts.Db, "database", "d", sqlite.GetDbPath(), "Name of database")
}
