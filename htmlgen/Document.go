package htmlgen

var documentTemplate *Document = nil

// Document represents an HTML document
type Document struct {
	Title  string
	format *FormatSection
}

// CreateDocument creates a new Document
func CreateDocument() *Document {
	doc := &Document{
		Title: "Canvas Sync",
	}
	args := []interface{}{
		&doc.Title,
		FormatSectionChild,
	}
	if documentTemplate == nil {
		var err error
		if doc.format, err = CreateFormatSection(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8" />
		<title>%s</title>
	</head>
	<body>
		%s
	</body>
</html>
`, args); err != nil {
			panic(err)
		}
	} else {
		doc.format = documentTemplate.format.Clone(args)
	}
	return doc
}

func init() {
	documentTemplate = CreateDocument()
}

// AppendChild adds a section to the document body
func (d *Document) AppendChild(child Section) {
	d.format.AppendChild(child)
}

// Children gets the elements in the document body
func (d *Document) Children() []Section {
	return d.format.Children()
}

func (d *Document) String() string {
	return d.format.String()
}

// Parse a document
func (d *Document) Parse(str string, childCtors []ChildConstructor) (string, bool) {
	return d.format.Parse(str, childCtors)
}

// ParseDocument parses an entire document
func ParseDocument(str string, childCtors []ChildConstructor) *Document {
	d := CreateDocument()
	if str, ok := d.Parse(str, childCtors); len(str) == 0 && ok {
		return d
	}
	return nil
}
