package completion

import (
	"github.com/dkaslovsky/MyMint/cmd/completion/bash"
	"github.com/dkaslovsky/MyMint/cmd/completion/zsh"
	"github.com/spf13/cobra"
)

// CreateCompletionCmd generates the configuration for the completion subcommand
func CreateCompletionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Subcommand for creating shell completions",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	cmd.AddCommand(
		bash.CreateBashCmd(),
		zsh.CreateZshCmd(),
	)
	return cmd
}
