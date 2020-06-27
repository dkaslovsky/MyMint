package main

import (
	"database/sql"
	"fmt"
	"log"
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

	insertRows := []interface{}{
		ExampleTableRow{Name: "foo", Val: 10.1},
		ExampleTableRow{Name: "bar", Val: 13.5},
		ExampleTableRow{Name: "baz", Val: 11},
	}
	_, err = db.
		Insert(exampleTable.Name).
		Rows(insertRows...).
		Executor().
		Exec()
	if err != nil {
		log.Fatal(err)
	}

	rows := []struct {
		ID int64 `db:"id"`
		ExampleTableRow
	}{}
	err = db.
		From(exampleTable.Name).
		Where(goqu.C("val").Gt(10.5)).
		ScanStructs(&rows)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)
}
