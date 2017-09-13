package client

import (
	"bufio"
	"encoding/csv"
	"os"
)

// Csv is struct for CSV files
type Csv struct {
	path string
}

// NewCsv initializes
func NewCsv(path string) Csv {
	return Csv{
		path: path,
	}
}

// Fetch returns data from this provider
func (o Csv) Fetch() ([][]string, error) {
	return o.read()
}

func (o Csv) read() (data [][]string, err error) {
	// Readers and other needed
	csvFile, _ := os.Open(o.path)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, err
}
