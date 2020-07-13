package conf

import (
	"os"
	"path/filepath"
)

const (
	// ConfEnvVar is the environment variable for the path to configuration files
	ConfEnvVar = "MYMINT_CONF_DIR"
	// DataSourceDir is the name of the directory where datasource files are stored
	DataSourceDir = "datasources"
	// CategoryDir is the name of the directory where category files are stored
	CategoryDir = "categories"
	// ManualCategoryFile is the name of the file where manual categories are stored
	ManualCategoryFile = "manual"
	// KeywordCategoryFile is the name of the file where keyword categories are stored
	KeywordCategoryFile = "keyword"
)

// AppDir is the directory containing files neccessary for the application to run
var AppDir = os.Getenv(ConfEnvVar)

// GetDataSourcePath returns the path to the datasource files
func GetDataSourcePath() (path string) {
	return filepath.Join(AppDir, DataSourceDir)
}

// GetManualCategoryPath returns the path to the manual category file
func GetManualCategoryPath() (path string) {
	return filepath.Join(AppDir, CategoryDir, ManualCategoryFile)
}

// GetKeywordCategoryPath returns the path to the keyword category file
func GetKeywordCategoryPath() (path string) {
	return filepath.Join(AppDir, CategoryDir, KeywordCategoryFile)
}
