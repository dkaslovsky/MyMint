package bash

import (
	"os"

	"github.com/spf13/cobra"
)

// CreateBashCmd generates the configuration for the zsh subcommand
func CreateBashCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bash",
		Short: "Generate completions for bash",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.GenBashCompletion(os.Stdout)
			return nil
		},
	}
	return cmd
}
