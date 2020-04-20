package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	assignmentTemplate *Assignment
	// AssignmentChildCtor for parsing a template
	AssignmentChildCtor = func() (htmlgen.Section, []htmlgen.ChildConstructor) {
		return CreateAssignment(), []htmlgen.ChildConstructor{
			AssignmentSubmissionChildCtor,
		}
	}
)

// Assignment HTML template
type Assignment struct {
	Data   canvas.Assignment
	format *htmlgen.FormatSection
}

// CreateAssignment creates a new template
func CreateAssignment() *Assignment {
	obj := &Assignment{}
	args := []interface{}{
		&obj.Data.Name,
		&obj.Data.Description,
		htmlgen.FormatSectionChild,
		&obj.Data.PointsPossible,
		htmlgen.CreateDateTimeFormat(&obj.Data.CreatedAt),
		htmlgen.CreateDateTimeFormat(&obj.Data.UnlockAt),
		htmlgen.CreateDateTimeFormat(&obj.Data.DueAt),
		htmlgen.CreateDateTimeFormat(&obj.Data.LockAt),
		htmlgen.CreateDateTimeFormat(&obj.Data.UpdatedAt),
	}
	if assignmentTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	<h1>%s</h1>
	<main>
		%s
	</main>
	<div>
		%s
	</div>
	<footer>
		<p>Points possible: %.2f</p>
		<p>Created at: %s</p>
		<p>Unlocked at: %s</p>
		<p>Due at: %s</p>
		<p>Locked at: %s</p>
		<p>Last modified at: %s</p>
	</footer>
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = assignmentTemplate.format.Clone(args)
	}
	return obj
}

func init() {
	assignmentTemplate = CreateAssignment()
}

// AppendChild adds a child to the section
func (t *Assignment) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *Assignment) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *Assignment) String() string {
	return t.format.String()
}

// Parse the template
func (t *Assignment) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
