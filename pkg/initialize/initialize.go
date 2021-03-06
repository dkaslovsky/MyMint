package initialize

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/ledger"
)

// Initialize creates the mymint data directory, subdirectories, and database with an empty ledger table
func Initialize(dbPath string) (err error) {
	_, err = os.Stat(conf.Config.AppDir)
	if !os.IsNotExist(err) {
		return fmt.Errorf("Cannot initialize: directory [%s] already exists", conf.Config.AppDir)
	}

	// create top level directory
	err = os.MkdirAll(conf.Config.AppDir, 0755)
	if err != nil {
		return err
	}

	// create datasource subdir
	err = os.MkdirAll(conf.Config.DataSourcePath, 0755)
	if err != nil {
		return err
	}

	// create category subdir
	err = os.MkdirAll(conf.Config.CategoryPath, 0755)
	if err != nil {
		return err
	}

	// write empty ledger category file
	fileHandle, err := os.OpenFile(conf.Config.LedgerCategoryFilePath, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	fileHandle.Close()

	// write keyword category file with empty json
	err = ioutil.WriteFile(conf.Config.KeywordCategoryFilePath, []byte("{}"), 0644)
	if err != nil {
		return err
	}

	// create db and ledger table
	db, err := sqlite.NewDb(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.CreateTable(ledger.TableName, ledger.TableSchema)
	if err != nil {
		return err
	}

	return nil
}
