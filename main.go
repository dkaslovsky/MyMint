package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dkaslovsky/MyMint/pkg/data" // PICK A BETTER NAME THAN DATA!
	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/parse"

	"github.com/doug-martin/goqu/v9"
)

var dbName = "mydb.db"

func main() {

	db, err := sqlite.NewDb(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	table := sqlite.NewTable("mytable", data.ExampleTableSchema)
	err = db.CreateTable(table)
	if err != nil {
		log.Fatal(err)
	}

	csvFile, err := os.Open("./example.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	csvRows, err := parse.ReadCSV(csvFile, data.ExampleTableCSVParser)
	if err != nil {
		log.Fatal(err)
	}

	err = db.InsertRows(table, csvRows)
	if err != nil {
		log.Fatal(err)
	}

	// ------------------------------------------------
	// validate that data was inserted
	queryRows := []struct {
		ID int64 `db:"id"`
		data.ExampleTableRow
	}{}
	err = db.
		From(table.Name).
		Where(goqu.C("val").Gt(101)).
		ScanStructs(&queryRows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queryRows)
}
