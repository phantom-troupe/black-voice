black-voice ()
==============

google spreadsheet server side client SQL library created by golang

Outline
-------

This library can extract data from google spreadsheet by SQL like syntax.
e.g: `SELECT * WHERE A = "Bob"`
Of course you can extract data from private spreadsheet.

Installing
----------

```bash
go get -u github.com/phantom-troupe/blackvoice
```

Example
-------

### Public Spreadsheet

```golang
package main

import (
    "fmt"
    "github.com/phantom-troupe/blackvoice"
)

func main() {
    sheet := blackvoice.NewPublicSpreadsheet(
        "<spreadsheet key>",
        <GID (your spreadsheet table number)>,
        <line number of header>,
    )
    query := `SELECT * WHERE A = "Bob"`
    fmt.Println("%#v", sheet.Query(query))
}
```
