package ledger

import (
	"fmt"
	"strings"

	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
)

// TableName is the name of the ledger table
const TableName = "ledger"

// TableSchema is the schema for the ledger table
var TableSchema = sqlite.Schema{
	"id":          "INTEGER PRIMARY KEY",
	"Date":        "TEXT",
	"Amount":      "REAL",
	"Description": "TEXT",
	"Category":    "TEXT",
}

// Row is a ledger entry
type Row struct {
	ID          int64       `db:"id"`
	Date        string      `db:"Date"`
	Amount      float64     `db:"Amount"`
	Description string      `db:"Description"`
	Category    interface{} `db:"Category"`
}

func (r *Row) String() (s string) {
	fields := []string{
		fmt.Sprintf("ID: %d", r.ID),
		fmt.Sprintf("Date: %s", r.Date),
		fmt.Sprintf("Amount: %0.2f", r.Amount),
		fmt.Sprintf("Description: %s", r.Description),
		fmt.Sprintf("Category: %v", r.Category),
	}
	return strings.Join(fields, " | ")
}
