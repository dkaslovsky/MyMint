package initialize

import (
	"github.com/dkaslovsky/MyMint/pkg/initialize"
	"github.com/spf13/cobra"
)

func CreateInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize MyMint",
		RunE: func(cmd *cobra.Command, args []string) error {

			return initialize.Initialize("mymint.db")

		},
	}
	return cmd
}
