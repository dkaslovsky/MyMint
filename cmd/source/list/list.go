package list

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/spf13/cobra"
)

// CreateListCmd generates the configuration for the list subcommand
func CreateListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List datasources",
		RunE: func(cmd *cobra.Command, args []string) error {

			path := conf.GetDataSourcePath()
			fileInfo, err := os.Stat(path)
			if err != nil {
				return err
			}
			if !fileInfo.Mode().IsDir() {
				return fmt.Errorf("[%s] is not a directory", path)
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
