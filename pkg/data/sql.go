package data

import (
	"fmt"
	"strconv"

	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
)

var ExampleTableSchema = sqlite.Schema{
	"id":   "INTEGER PRIMARY KEY",
	"name": "TEXT",
	"val":  "REAL",
}

type ExampleTableRow struct {
	Name string  `db:"name"`
	Val  float64 `db:"val"`
}

func ExampleTableCSVParser(record []string) (row interface{}, err error) {
	// this check is for safety but not needed since record will come from csv.Reader.Read()
	expectedRecordLen := 2
	if len(record) != expectedRecordLen {
		return row, fmt.Errorf(
			"malformed record [%v]: expected [%d] columns, found [%d]",
			record,
			expectedRecordLen,
			len(record),
		)
	}
	name := record[0]
	val, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return row, fmt.Errorf("could not parse record [%v]: %s", record, err)
	}
	return ExampleTableRow{
		Name: name,
		Val:  val,
	}, nil
}
