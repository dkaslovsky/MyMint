package cmd

import (
	"github.com/dkaslovsky/MyMint/cmd/category"
	"github.com/dkaslovsky/MyMint/cmd/delete"
	"github.com/dkaslovsky/MyMint/cmd/insert"
	"github.com/dkaslovsky/MyMint/cmd/source"
	"github.com/dkaslovsky/MyMint/cmd/table"
	"github.com/spf13/cobra"
)

// Run starts, configures, and executes the cli interface
func Run() error {
	cmd := &cobra.Command{
		Use:           "mymint",
		Short:         "root mymint command",
		Long:          `mymint persists personal finance data`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.AddCommand(
		insert.CreateInsertCmd(),
		delete.CreateDeleteCmd(),
		source.CreateSourceCmd(),
		table.CreateTableCmd(),
		category.CreateCategoryCmd(),
	)
	return cmd.Execute()
}
