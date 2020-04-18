package htmlgen

// CustomFormat formats a custom type (with %s)
type CustomFormat interface {
	FormatHTML() string
}
