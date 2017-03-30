package blackvoice

import (
	"reflect"
	"testing"
)

func TestPublicSpreadsheetQuery_gid0(t *testing.T) {
	spreadsheet := &PublicSpreadsheet{
		"14aayP76anHyRJyeVcTBJMTvqwyPeWZFFBpGffhko9HU",
		0,
		1,
	}
	actual, err := spreadsheet.query(`SELECT A,B,C WHERE C = "spreadsheet-sql-public001@example.com"`)
	var expected SpreadSheetValuesCollection
	expected = append(expected, SpreadSheetValues{
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

