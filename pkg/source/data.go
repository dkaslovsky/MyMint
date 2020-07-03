package source

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/dkaslovsky/MyMint/pkg/parse"
	"github.com/doug-martin/goqu/v9"
)

// DataSource contains metadata for reading and writing a datasource
type DataSource struct {
	Name   string        `json:"name"`
	Table  string        `json:"table"`
	Schema sqlite.Schema `json:"schema"`
	Csv    struct {
		Fields []string `json:"fields"`
		Header bool     `json:"header"`
	} `json:"csv"`
}

// LoadDataSource unmarshals a JSON definition for a datasource from a file
func LoadDataSource(path string) (ds *DataSource, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return ds, err
	}
	ds = &DataSource{}
	err = json.Unmarshal(bytes, ds)
	if err != nil {
		return ds, err
	}
	return ds, nil
}

// GenerateCsvRowParser returns a function for parsing lines of a CSV data source
func (ds *DataSource) GenerateCsvRowParser() (rowParser parse.CsvRowParser, err error) {
	return func(vals []string) (record interface{}, err error) {
		var ok bool
		row := goqu.Record{}

		for i, val := range vals {
			field := ds.Csv.Fields[i]
			fieldType := sqlite.DbType(strings.Split(ds.Schema[field], " ")[0])
			parser, found := parsers[fieldType]
			if !found {
				return row, fmt.Errorf("No parser found for type [%v]", fieldType)
			}
			parsed, err := parser(val)
			if err != nil {
				return row, err
			}

			switch fieldType {
			case sqlite.DbInteger:
				row[field], ok = parsed.(int64)
				if !ok {
					return row, fmt.Errorf("Could not cast value [%v] to int64", parsed)
				}
			case sqlite.DbFloat:
				row[field], ok = parsed.(float64)
				if !ok {
					return row, fmt.Errorf("Could not cast value [%v] to float64", parsed)
				}
			case sqlite.DbString:
				row[field], ok = parsed.(string)
				if !ok {
					return row, fmt.Errorf("Could not cast value [%v] to string", parsed)
				}
			default:
				return row, fmt.Errorf("Unknown type [%v]", fieldType)
			}
		}
		return row, nil
	}, nil
}

func (ds *DataSource) String() (s string) {
	bytes, _ := json.MarshalIndent(ds, "", "\t") // safe to ignore err since struct came from json
	return string(bytes)
}

type parser func(string) (interface{}, error)

var (
	integerParser = func(s string) (interface{}, error) {
		if s == "" {
			return 0, nil
		}
		return strconv.ParseInt(s, 10, 64)
	}
	floatParser = func(s string) (interface{}, error) {
		if s == "" {
			return 0.0, nil
		}
		return strconv.ParseFloat(s, 64)
	}
	stringParser = func(s string) (interface{}, error) {
		return s, nil
	}
	parsers = map[sqlite.DbType]parser{
		sqlite.DbInteger: integerParser,
		sqlite.DbFloat:   floatParser,
		sqlite.DbString:  stringParser,
	}
)
