package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	pageTemplate *Page
	// PageChildCtor for parsing a template
	PageChildCtor = func() (htmlgen.Section, []htmlgen.ChildConstructor) {
		return CreatePage(), []htmlgen.ChildConstructor{}
	}
)

// Page HTML template
type Page struct {
	Data   canvas.Page
	Editor string
	format *htmlgen.FormatSection
}

// CreatePage creates a new template
func CreatePage() *Page {
	obj := &Page{}
	args := []interface{}{
		&obj.Data.Body,
		&obj.Editor,
		htmlgen.CreateDateTimeFormat(&obj.Data.UpdatedAt),
	}
	if pageTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	<main>
		%s
	</main>
	<footer>
		<p>Last edited by %s on %s</p>
	</footer>
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = pageTemplate.format.Clone(args)
	}
	return obj
}

func init() {
	pageTemplate = CreatePage()
}

// AppendChild adds a child to the section
func (t *Page) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *Page) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *Page) String() string {
	return t.format.String()
}

// Parse the template
func (t *Page) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
