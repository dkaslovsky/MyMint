package initialize

import (
	"os"

	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/spf13/cobra"
)

// CreateInitCmd generates the configuration for the init subcommand
func CreateInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Create ledger categories file",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := conf.Config.LedgerCategoryPath
			fileHandle, err := os.OpenFile(path, os.O_CREATE, 0644)
			if err != nil {
				return err
			}
			return fileHandle.Close()
		},
	}
	return cmd
}
