package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsv(file string, skipFirstLine bool) ([][]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("ReadCsv: unable to read csv from: " + file)
	}
	defer f.Close()

	r := csv.NewReader(f)
	if skipFirstLine {
		if _, err := r.Read(); err != nil {
			return nil, err
		}
	}

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
