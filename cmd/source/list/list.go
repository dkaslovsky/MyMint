package list

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dkaslovsky/MyMint/cmd/constants"
	"github.com/spf13/cobra"
)

// CreateListCmd generates the configuration for the list subcommand.
// It can be attached to any upstream cobra command
func CreateListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List datasources",
		RunE: func(cmd *cobra.Command, args []string) error {

			path := os.Getenv(constants.DataSourceEnvVar)
			if path == "" {
				return fmt.Errorf("could not read path to datasource files from environment variable [%s]", constants.DataSourceEnvVar)
			}

			fileInfo, err := os.Stat(path)
			if err != nil {
				return err
			}
			if !fileInfo.Mode().IsDir() {
				return fmt.Errorf("expected directory, received [%s]", path)
			}
			files, err := ioutil.ReadDir(path)
			if err != nil {
				return err
			}
			for _, file := range files {
				fileName := file.Name()
				fileExt := filepath.Ext(fileName)
				fmt.Println(strings.TrimSuffix(fileName, fileExt))
			}

			return nil
		},
	}
	return cmd
}
