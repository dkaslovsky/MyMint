package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	*goqu.Database
	driver *sql.DB
}

func NewDb(name string) (db *Db, err error) {
	driver, err := sql.Open("sqlite3", name)
	if err != nil {
		return db, err
	}
	return &Db{
		Database: goqu.New("sqlite3", driver),
		driver:   driver,
	}, nil
}

func (db *Db) Close() (err error) {
	return db.driver.Close()
}

type Schema map[string]string

type Table struct {
	Name   string
	Schema Schema
}

type ExampleTableRow struct {
	Name string  `db:"name"`
	Val  float64 `db:"val"`
}

var dbName = "mydb.db"

var exampleTable = Table{
	Name: "mytable",
	Schema: Schema{
		"id":   "INTEGER PRIMARY KEY",
		"name": "TEXT",
		"val":  "REAL",
	},
}

func buildCreateTableQuery(table Table) (query string) {
	var s []string
	for col, colType := range table.Schema {
		s = append(s, fmt.Sprintf("%s %s", col, colType))
	}
	return fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (%s)",
		table.Name,
		strings.Join(s, ", "),
	)
}

func readCSV(path string) (rows []interface{}, err error) {
	recordLen := 2

	csvFile, err := os.Open(path)
	if err != nil {
		return rows, err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return rows, err
		}
		if len(record) != recordLen {
			return rows, fmt.Errorf(
				"malformed record [%v]: expected [%d] columns, found [%d]",
				record,
				recordLen,
				len(record),
			)
		}
		val, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return rows, fmt.Errorf("could not parse record [%v]: %s", record, err)
		}
		row := ExampleTableRow{
			Name: record[0],
			Val:  val,
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func main() {

	db, err := NewDb(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(buildCreateTableQuery(exampleTable))
	if err != nil {
		log.Fatal(err)
	}

	csvRows, err := readCSV("./example.csv")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.
		Insert(exampleTable.Name).
		Rows(csvRows...).
		Executor().
		Exec()
	if err != nil {
		log.Fatal(err)
	}

	queryRows := []struct {
		ID int64 `db:"id"`
		ExampleTableRow
	}{}
	err = db.
		From(exampleTable.Name).
		Where(goqu.C("val").Gt(101)).
		ScanStructs(&queryRows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queryRows)
}
