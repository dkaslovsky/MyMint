package parse

import (
	"encoding/csv"
	"io"
)

// ReadCSV parses lines of a CSV file
func ReadCSV(r io.Reader, lineParser func([]string) (interface{}, error)) (rows []interface{}, err error) {
	reader := csv.NewReader(r)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return rows, err
		}
		row, err := lineParser(line)
		if err != nil {
			return rows, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}
