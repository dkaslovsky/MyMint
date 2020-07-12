package constants

import (
	"os"
	"path/filepath"
)

var (
	// DefaultDb is the default sqlite database file
	DefaultDb = filepath.Join(os.Getenv(ConfEnvVar), "mymint.db")
	// DefaultIDColumn is the default name of the primary key in a SQL table
	DefaultIDColumn = "id"
)
