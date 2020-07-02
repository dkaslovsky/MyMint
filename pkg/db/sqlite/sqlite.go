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

// DbType is a sqlite3 data type
type DbType string

// sqlite3 data types
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
func (db *Db) CreateTable(name string, schema Schema) (err error) {
	query := fmt.Sprintf("CREATE TABLE %s (%s)", name, schema)
	_, err = db.Exec(query)
	return err
}

// InsertRows inserts rows into a table
func (db *Db) InsertRows(table string, rows []interface{}) (err error) {
	_, err = db.
		Insert(table).
		Rows(rows...).
		Executor().
		Exec()
	return err
}

// Schema represents a table's schema as a map of column name to type
type Schema map[string]string

func (sc Schema) String() (st string) {
	s := []string{}
	for col, colType := range sc {
		s = append(s, fmt.Sprintf("`%s` %s", col, colType))
	}
	return strings.Join(s, ",")
}
