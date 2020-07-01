package parse

import (
	"encoding/csv"
	"io"
)

// CsvRowParser is a function that parses a row of a CSV
type CsvRowParser func([]string) (interface{}, error)

// ReadCsvWithoutHeader parses lines of a CSV file that does not have a header
func ReadCsvWithoutHeader(r io.Reader, rowParser CsvRowParser) (rows []interface{}, err error) {
	reader := csv.NewReader(r)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return rows, err
		}
		row, err := rowParser(line)
		if err != nil {
			return rows, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}

// ReadCsvWithHeader parses lines of a CSV file that has a header
func ReadCsvWithHeader(r io.Reader, rowParser CsvRowParser) (rows []interface{}, err error) {
	reader := csv.NewReader(r)
	seenHeader := false
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return rows, err
		}
		// skip the header
		if !seenHeader {
			seenHeader = true
			continue
		}
		row, err := rowParser(line)
		if err != nil {
			return rows, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}
