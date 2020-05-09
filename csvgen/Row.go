package csvgen

import (
	"fmt"
	"strings"
)

// Row represents a row in the CSV file
type Row interface {
	CSV([]Column) [][]string
}

type basicRow struct {
	Data []interface{}
}

// CreateRow creates a new Row
func CreateRow(data ...interface{}) Row {
	return basicRow{
		Data: data,
	}
}

// CSV converts the row into CSV fields
func (b basicRow) CSV(columns []Column) [][]string {
	res := make([]string, len(columns))
	d := b.Data
	for i, col := range columns {
		res[i] = col.DataString(d[0:col.NumFormatArgs])
		if strings.ContainsRune(res[i], ',') {
			res[i] = fmt.Sprintf("\"%s\"", strings.ReplaceAll(res[i], "\"", "\"\""))
		}
		d = d[col.NumFormatArgs:]
	}
	return [][]string{res}
}
