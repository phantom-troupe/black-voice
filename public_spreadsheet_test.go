package blackvoice

import (
	"reflect"
	"testing"
)

func TestPublicSpreadsheetQuery_gid0_1object(t *testing.T) {
	spreadsheet := &PublicSpreadsheet{
		"14aayP76anHyRJyeVcTBJMTvqwyPeWZFFBpGffhko9HU",
		0,
		1,
	}
	actual, err := spreadsheet.Query(`SELECT A,B,C WHERE C = "spreadsheet-sql-public001@example.com"`)
	var expected SpreadsheetValuesCollection
	expected = append(expected, SpreadsheetValues{
		"name": "spreadsheet-sql-public001",
		"url": "https://spreadsheet-sql-public001.example.com",
		"email": "spreadsheet-sql-public001@example.com",
	})
	if err != nil {
		t.Fatalf("This test must be no error. Error: %#v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%s != %s", actual, expected)
	}
}
