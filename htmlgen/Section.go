package htmlgen

// ChildConstructor is a function that can create new child nodes
type ChildConstructor func() (Section, []ChildConstructor)

// Section represents a section of HTML code
type Section interface {
	AppendChild(Section)
	Children() []Section
	String() string
	Parse(string, []ChildConstructor) (string, bool)
}
