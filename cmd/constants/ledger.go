package constants

import "github.com/dkaslovsky/MyMint/pkg/db/sqlite"

var (
	// TableName is the name of the ledger table
	TableName = "ledger"
	// TableSchema is the schema for the ledger table
	TableSchema = sqlite.Schema{
		"id":          "INTEGER PRIMARY KEY",
		"Category":    "TEXT",
		"Date":        "TEXT",
		"Amount":      "REAL",
		"Description": "TEXT",
	}
)
