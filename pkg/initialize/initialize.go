package initialize

import (
	"fmt"
	"os"

	"github.com/dkaslovsky/MyMint/pkg/conf"
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/ledger"
)

// Initialize creates the mymint data directory, subdirectories, and database with an empty ledger table
func Initialize(dbName string) (err error) {
	_, err = os.Stat(conf.Config.AppDir)
	if !os.IsNotExist(err) {
		return fmt.Errorf("Cannot initialize: directory [%s] already exists", conf.Config.AppDir)
	}

	err = os.MkdirAll(conf.Config.AppDir, 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(conf.Config.DataSourcePath, 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll(conf.Config.CategoryPath, 0755)
	if err != nil {
		return err
	}

	fileHandle, err := os.OpenFile(conf.Config.LedgerCategoryFilePath, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	fileHandle.Close()

	fileHandle, err = os.OpenFile(conf.Config.KeywordCategoryFilePath, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	fileHandle.Close()

	db, err := sqlite.NewDb(dbName)
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
