package csvgen

// Section represents a section header
type Section struct {
	Row
	Title   []interface{}
	SubRows []Row
}

// CreateSection creates a new Section
func CreateSection(title []interface{}, data ...interface{}) *Section {
	return &Section{
		Title: title,
		Row:   CreateRow(data...),
	}
}

// AddRow creates a new row in the section
func (s *Section) AddRow(data ...interface{}) {
	s.SubRows = append(s.SubRows, CreateRow(data...))
}

// AddSection creates a new subsection in the section
func (s *Section) AddSection(title []interface{}, data ...interface{}) *Section {
	section := CreateSection(title, data...)
	s.SubRows = append(s.SubRows, section)
	return section
}

// CSV converts the section into CSV fields
func (s Section) CSV(columns []Column) [][]string {
	rows := [][]string{{}}
	if len(columns) > 0 {
		rows[0] = append([]string{columns[0].DataString(s.Title)}, s.Row.CSV(columns[1:])[0]...)
		for _, row := range s.SubRows {
			for _, fmtRow := range row.CSV(columns[1:]) {
				rows = append(rows, append([]string{""}, fmtRow...))
			}
		}
	}
	return rows
}
