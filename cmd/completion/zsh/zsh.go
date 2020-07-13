package zsh

import (
	"os"

	"github.com/spf13/cobra"
)

// CreateZshCmd generates the configuration for the zsh subcommand
func CreateZshCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "zsh",
		Short: "Generate completions for zsh",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.GenZshCompletion(os.Stdout)
			return nil
		},
	}
	return cmd
}
