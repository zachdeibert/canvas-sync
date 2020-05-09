package csvgen

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	formatArgCounterRegex = regexp.MustCompile("(?:^|[^%])%[^%]")
	// ColumnNil represents a blank column
	ColumnNil = CreateColumn("", "")
)

// Column represents one column in the CSV file
type Column struct {
	Name          string
	FormatString  string
	NumFormatArgs int
}

// CreateColumn creates a new Column
func CreateColumn(name, format string) Column {
	return Column{
		Name:          name,
		FormatString:  format,
		NumFormatArgs: len(formatArgCounterRegex.FindAllString(format, -1)),
	}
}

// HeaderString gets the string to put in the header for this Column
func (c Column) HeaderString() string {
	if strings.ContainsRune(c.Name, ',') {
		return fmt.Sprintf("\"%s\"", strings.ReplaceAll(c.Name, "\"", "\"\""))
	}
	return c.Name
}

// DataString gets the string to put in a cell in the body of the CSV for this Column
func (c Column) DataString(args []interface{}) string {
	return fmt.Sprintf(c.FormatString, args...)
}
