package blackvoice

import (
	"fmt"
	"io"
	"encoding/csv"
	"strings"
)

func csvToSpreadsheetValuesCollection(input string) (SpreadsheetValuesCollection, error) {
	var result SpreadsheetValuesCollection
	reader := csv.NewReader(strings.NewReader(input))
	header, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("Failed to read header(first) line %s", err)
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		spreadSheetValue := SpreadsheetValues{}
		for index, column := range header {
			spreadSheetValue[column] = record[index]
		}
		result = append(result, spreadSheetValue)
	}
	return result, nil
}