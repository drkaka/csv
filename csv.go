package csv

import (
	"encoding/csv"
	"fmt"
	"os"
)

// ParseFile to parse a CSV file.
func ParseFile(f string) ([][]string, error) {
	var records [][]string
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err = r.ReadAll()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(records); i++ {
		fmt.Printf("%20s\t%10s\t%10s\n", records[i][0], records[i][1], records[i][2])
	}
	return records, nil
}
