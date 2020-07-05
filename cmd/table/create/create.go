package create

import (
	"os"
	"path/filepath"

	"github.com/dkaslovsky/MyMint/cmd/constants"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/source"
	"github.com/spf13/cobra"
)

// Options are options for configuring the create command
type Options struct {
	Db string
}

// CreateCreateCmd generates the configuration for the create subcommand.
// It can be attached to any upstream cobra command
func CreateCreateCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a table from a datasource definition",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			confDir := os.Getenv(constants.ConfEnvVar)
			sourcePath := filepath.Join(confDir, constants.DataSourceDir, args[0])
			ext := filepath.Ext(sourcePath)
			if ext == "" {
				sourcePath += ".json"
			}

			ds, err := source.LoadDataSource(sourcePath)
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
	flags.StringVarP(&opts.Db, "database", "d", constants.DefaultDb, "Name of database")
}
