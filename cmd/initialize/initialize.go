package initialize

import (
	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/dkaslovsky/MyMint/pkg/initialize"
	"github.com/spf13/cobra"
)

// Options are options for configuring the initialize command
type Options struct {
	Db string
}

// CreateInitCmd generates the configuration for the init subcommand
func CreateInitCmd() *cobra.Command {
	opts := Options{}
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize MyMint",
		RunE: func(cmd *cobra.Command, args []string) error {
			return initialize.Initialize(opts.Db)
		},
	}
	attachOpts(cmd, &opts)
	return cmd
}

func attachOpts(cmd *cobra.Command, opts *Options) {
	flags := cmd.Flags()
	flags.StringVarP(&opts.Db, "database", "d", conf.Config.DefaultSqliteDbPath, "Name of database")
}
