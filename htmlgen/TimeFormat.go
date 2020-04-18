package htmlgen

import "time"

// TimeFormat represents a field that contains a formatted time
type TimeFormat struct {
	Ptr    *time.Time
	Format string
}

// FormatHTML formats the TimeFormat for printing into the HTML document
func (t *TimeFormat) FormatHTML() string {
	return t.Ptr.Format(t.Format)
}

// ParseHTML is the opposite of FormatHTML
func (t *TimeFormat) ParseHTML(str string) error {
	var err error
	*t.Ptr, err = time.Parse(t.Format, str)
	return err
}

// CreateTimeFormat creates a new time format
func CreateTimeFormat(ptr *time.Time, format string) *TimeFormat {
	return &TimeFormat{
		Ptr:    ptr,
		Format: format,
	}
}

// CreateDateFormat creates a new time format for dates
func CreateDateFormat(ptr *time.Time) *TimeFormat {
	return CreateTimeFormat(ptr, "Mon Jan 2, 2006")
}

// CreateDateTimeFormat creates a new time format for dates with an associated time
func CreateDateTimeFormat(ptr *time.Time) *TimeFormat {
	return CreateTimeFormat(ptr, "Mon Jan 2, 2006 at 3:04:05 PM")
}
