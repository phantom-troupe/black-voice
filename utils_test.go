package blackvoice

import (
	"testing"
	"reflect"
)

func TestCsvToSpreadSheetValuesCollection_validCsv(t *testing.T) {
	input := `"key1","key2","key3"
"value1","value2","value3"`
	actual, err := csvToSpreadsheetValuesCollection(input)
	var expected SpreadsheetValuesCollection
	expected = append(expected, SpreadsheetValues{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	if err != nil {
		t.Fatalf("This test must not be occurred error. Raw: %s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("%#v != %#v", actual, expected)
	}
}

func TestCsvToSpreadSheetValuesCollection_headerOnly(t *testing.T) {
	input :=`"key1","key2","key3"`
	actual, err := csvToSpreadsheetValuesCollection(input)
	var expected SpreadsheetValuesCollection
	if err != nil {
		t.Fatalf("This test must not be occurred error. Raw: %s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Error("%#v != %#v", actual, expected)
	}
}