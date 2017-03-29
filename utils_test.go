package blackvoice

import (
	"testing"
)

func TestCsvToMapCollection(t *testing.T) {
	input := `"key1","key2","key3"
"value1","value2","value3"`
	var expected []map[string]string
	expected := append(expected, map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
}