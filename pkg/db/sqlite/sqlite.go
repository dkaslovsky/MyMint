package sqlite

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/doug-martin/goqu/v9"
	// blank import for drivers
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

type DbType string

const (
	DbInteger DbType = "INTEGER"
	DbFloat   DbType = "REAL"
	DbString  DbType = "TEXT"
)

// Db wraps goqu.Database to expose a Close method for the underlying database driver
type Db struct {
	*goqu.Database
	driver *sql.DB
}

// NewDb creates a new Db
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

// Close closes the database driver
func (db *Db) Close() (err error) {
	return db.driver.Close()
}

// CreateTable creates a table
func (db *Db) CreateTable(table *Table) (err error) {
	_, err = db.Exec(table.GetCreateQuery())
	return err
}

// InsertRows inserts rows into a table
func (db *Db) InsertRows(tableName string, rows []interface{}) (err error) {
	_, err = db.
		Insert(tableName).
		Rows(rows...).
		Executor().
		Exec()
	return err
}

// Schema represents a table's schema as a map of column name to type
type Schema map[string]string

// Table is a SQL table
type Table struct {
	Name   string
	Schema Schema
}

// NewTable creates a new Table
func NewTable(name string, schema Schema) (table *Table) {
	return &Table{
		Name:   name,
		Schema: schema,
	}
}

// GetCreateQuery builds a query to create a table
func (table *Table) GetCreateQuery() (query string) {
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
