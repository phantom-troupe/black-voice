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

func (x PublicSpreadsheet) query(q string) (string, error) {
	res, body, err := gorequest.New().Get(
		"https://docs.google.com/a/google.com/spreadsheets/d/" +
		x.key +
		"/gviz/tq?gid=" + strconv.Itoa(x.gid) +
		"&headers=" + strconv.Itoa(x.lineNumberOfHeaders) +
		"&tq=" + q +
		"&tqx=out:csv").
	    End()

	if err != nil {
		return "", fmt.Errorf("Failed to fetch from public spreadsheet: %s", err[0])
	}

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("Returned status code is %d", res.StatusCode)
	}
	return body, nil
}
