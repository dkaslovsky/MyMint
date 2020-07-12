package constants

const (
	// DefaultDb is the default sqlite database file
	DefaultDb = "mydb.db"
	// DefaultIDColumn is the default name of the primary key in a SQL table
	DefaultIDColumn = "id"
	// ConfEnvVar is the environment variable for the path to configuration files
	ConfEnvVar = "MYMINT_CONF_DIR"
	// DataSourceDir is the name of the directory where datasource files are stored
	DataSourceDir = "datasources"
	// ManualCategoryFile is the name of the file where manual categories are stored
	ManualCategoryFile = "categories/manual"
	// KeywordCategoryFile is the name of the file where keyword categories are stored
	KeywordCategoryFile = "categories/keyword"
)
