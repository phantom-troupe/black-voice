package blackvoice

import (
	"fmt"
	"strconv"
	"github.com/parnurzeal/gorequest"
)

type PublicSpreadsheet struct {
	key string
	gid int
	lineNumberOfHeaders int
}

type SpreadSheetValues map[string]string
type SpreadSheetValuesCollection []SpreadSheetValues

func (x PublicSpreadsheet) query(q string) (SpreadSheetValuesCollection, error) {
	res, body, errs := gorequest.New().Get(
		"https://docs.google.com/a/google.com/spreadsheets/d/" +
		x.key +
		"/gviz/tq?gid=" + strconv.Itoa(x.gid) +
		"&headers=" + strconv.Itoa(x.lineNumberOfHeaders) +
		"&tq=" + q +
		"&tqx=out:csv").
	    End()

	if errs != nil {
		return make(SpreadSheetValuesCollection, 0), fmt.Errorf("Failed to fetch from public spreadsheet: %s", errs[0])
	}

	if res.StatusCode >= 400 {
		return make(SpreadSheetValuesCollection, 0), fmt.Errorf("Returned status code is %d", res.StatusCode)
	}
	collection, err := csvToSpreadSheetValuesCollection(body)

	if err != nil {
		return make(SpreadSheetValuesCollection, 0), err
	}
	return collection, nil
}
