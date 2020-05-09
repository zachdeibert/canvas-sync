package csvgen

import (
	"io/ioutil"
	"strings"
)

// CSV represents a CSV file
type CSV interface {
	AddRow(data ...interface{})
	AddSection(title []interface{}, data ...interface{}) *Section
}

// TopCSV represents the top-level of a CSV document
type TopCSV struct {
	Columns []Column
	Rows    []Row
}

// CreateCSV creates a new CSV
func CreateCSV() *TopCSV {
	return &TopCSV{
		Columns: []Column{},
		Rows:    []Row{},
	}
}

// AddColumn adds a column to the CSV file
func (c *TopCSV) AddColumn(name, format string) {
	c.Columns = append(c.Columns, CreateColumn(name, format))
}

// AddRow adds a row to the CSV file
func (c *TopCSV) AddRow(data ...interface{}) {
	c.Rows = append(c.Rows, CreateRow(data...))
}

// AddSection creates a new subsection in the section
func (c *TopCSV) AddSection(title []interface{}, data ...interface{}) *Section {
	section := CreateSection(title, data...)
	c.Rows = append(c.Rows, section)
	return section
}

// CSV converts the CSV file into fields
func (c TopCSV) CSV(columns []Column) [][]string {
	res := [][]string{make([]string, len(columns))}
	for i, col := range columns {
		res[0][i] = col.HeaderString()
	}
	for _, row := range c.Rows {
		res = append(res, row.CSV(columns)...)
	}
	return res
}

// Format converts the CSV file into fields
func (c TopCSV) Format() [][]string {
	return c.CSV(c.Columns)
}

func (c TopCSV) String() string {
	csv := c.Format()
	lines := make([]string, len(csv))
	for i, fields := range csv {
		lines[i] = strings.Join(fields, ",")
	}
	return strings.Join(lines, "\n")
}

// WriteFile writes the CSV to a file
func (c TopCSV) WriteFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(c.String()), 0644)
}
