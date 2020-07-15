package conf

import (
	"log"
	"os"
	"path/filepath"
)

const (
	// confEnvVar is the environment variable for the path to configuration files
	confEnvVar = "MYMINT_CONF_DIR"
	// dataSourceDir is the name of the directory where datasource files are stored
	dataSourceDir = "datasources"
	// categoryDir is the name of the directory where category files are stored
	categoryDir = "categories"
	// ledgerCategoryFile is the name of the file where ledger categories are stored
	ledgerCategoryFile = "ledger"
	// keywordCategoryFile is the name of the file where keyword categories are stored
	keywordCategoryFile = "keyword"
	// defaultSqliteDb is the name of the default sqlite database file
	defaultSqliteDb = "mymint.db"
)

// appDir is the directory containing files neccessary for the application to run
var appDir = os.Getenv(confEnvVar)

func init() {
	if appDir == "" {
		log.Fatalf("environment variable [%s] is not set!", confEnvVar)
	}
}

type config struct {
	// AppDir is the directory containing files neccessary for the application to run
	AppDir string
	// DataSourcePath is the path to the datasource files
	DataSourcePath string
	// CategoryPath is the path to the category files
	CategoryPath string
	// LedgerCategoryFilePath is the path to the ledger category file
	LedgerCategoryFilePath string
	// KeywordCategoryFilePath is the path to the keyword category file
	KeywordCategoryFilePath string
	// DefaultSqliteDbPath is the path to the database file using the default name
	DefaultSqliteDbPath string
}

// Config is the configuration necessary to run the application
var Config = config{
	AppDir:                  appDir,
	DataSourcePath:          filepath.Join(appDir, dataSourceDir),
	CategoryPath:            filepath.Join(appDir, categoryDir),
	LedgerCategoryFilePath:  filepath.Join(appDir, categoryDir, ledgerCategoryFile),
	KeywordCategoryFilePath: filepath.Join(appDir, categoryDir, keywordCategoryFile),
	DefaultSqliteDbPath:     filepath.Join(appDir, defaultSqliteDb),
}
