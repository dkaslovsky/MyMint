package dump

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/dkaslovsky/MyMint/pkg/source"
	"github.com/spf13/cobra"
)

// CreateDumpCmd generates the configuration for the dump subcommand
func CreateDumpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dump",
		Short: "Print a datasource file to the console",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			path := filepath.Join(conf.Config.DataSourcePath, args[0])
			ext := filepath.Ext(path)
			if ext == "" {
				path += ".json"
			}

			fileInfo, err := os.Stat(path)
			if err != nil {
				return err
			}
			if !fileInfo.Mode().IsRegular() {
				return fmt.Errorf("[%s] is not a file", path)
			}

			ds, err := source.LoadDataSource(path)
			if err != nil {
				return err
			}
			fmt.Println(ds)
			return nil
		},
	}
	return cmd
}
