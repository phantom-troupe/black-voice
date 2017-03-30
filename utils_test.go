package blackvoice

import (
	"testing"
	"reflect"
)

func TestCsvToSpreadSheetValuesCollection_validCsv(t *testing.T) {
	input := `"key1","key2","key3"
"value1","value2","value3"`
	actual, err := csvToSpreadSheetValuesCollection(input)
	var expected SpreadSheetValuesCollection
	expected = append(expected, SpreadSheetValues{
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
