package source

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/dkaslovsky/MyMint/pkg/source"
	"github.com/spf13/cobra"
)

// CreateSourceCmd generates the configuration for the source subcommand.
// It can be attached to any upstream cobra command
func CreateSourceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "source",
		Short: "Subcommand for source operations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			path := args[0]
			fileInfo, err := os.Stat(path)
			if err != nil {
				return err
			}
			switch mode := fileInfo.Mode(); {
			case mode.IsDir():
				err = listFiles(path)
				if err != nil {
					return err
				}
			case mode.IsRegular():
				err = showSource(path)
				if err != nil {
					return err
				}
			}

			return nil
		},
	}
	return cmd
}

func listFiles(path string) (err error) {
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
}

func showSource(path string) (err error) {
	ds, err := source.LoadDataSource(path)
	if err != nil {
		return err
	}
	fmt.Println(ds)
	return nil
}
