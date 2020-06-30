package conf

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dkaslovsky/MyMint/pkg/db/sqlite"
	"github.com/doug-martin/goqu/v9"
)

type Config struct {
	DataType  string      `json:"data_type"`
	TableName string      `json:"table_name"`
	Schema    interface{} `json:"schema"`
	CsvHeader []string    `json:"csv_header"`
}

var parsers = map[sqlite.DbType]func(string) (interface{}, error){
	sqlite.DbInteger: func(s string) (interface{}, error) {
		return strconv.ParseInt(s, 10, 64)
	},
	sqlite.DbFloat: func(s string) (interface{}, error) {
		return strconv.ParseFloat(s, 64)
	},
	sqlite.DbString: func(s string) (interface{}, error) {
		return s, nil
	},
}

func (c *Config) GenerateSchema() (schema sqlite.Schema, err error) {
	schema = sqlite.Schema{}
	for field, fieldType := range c.Schema.(map[string]interface{}) {
		schema[field] = fieldType.(string)
	}
	return schema, nil
}

func (c *Config) GenerateCsvParser() (lineParser func([]string) (interface{}, error), err error) {
	schema, _ := c.GenerateSchema()

	return func(vals []string) (record interface{}, err error) {
		var ok bool
		row := goqu.Record{}

		for i, val := range vals {
			field := c.CsvHeader[i]
			fieldType := sqlite.DbType(strings.Split(schema[field], " ")[0])
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
