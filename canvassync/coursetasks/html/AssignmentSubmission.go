package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	assignmentSubmissionTemplate *AssignmentSubmission
	// AssignmentSubmissionChildCtor for parsing a template
	AssignmentSubmissionChildCtor = func() (htmlgen.Section, []htmlgen.ChildConstructor) {
		return CreateAssignmentSubmission(), []htmlgen.ChildConstructor{
			AssignmentSubmissionAttachmentChildCtor,
			AssignmentSubmissionCommentChildCtor,
		}
	}
)

// AssignmentSubmission HTML template
type AssignmentSubmission struct {
	Data   canvas.Submission
	format *htmlgen.FormatSection
}

// CreateAssignmentSubmission creates a new template
func CreateAssignmentSubmission() *AssignmentSubmission {
	obj := &AssignmentSubmission{}
	args := []interface{}{
		htmlgen.CreateDateTimeFormat(&obj.Data.SubmittedAt),
		&obj.Data.SecondsLate,
		&obj.Data.Grade,
		&obj.Data.Attempt,
		htmlgen.CreateDateTimeFormat(&obj.Data.PostedAt),
		htmlgen.FormatSectionChild,
	}
	if assignmentSubmissionTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	<p>Submitted at %s (%.0f seconds late)</p>
	<p>Grade: %s (attempt #%d)</p>
	<p>Grade posted at %s</p>
	%s
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = assignmentSubmissionTemplate.format.Clone(args)
	}
	return obj
}

func init() {
	assignmentSubmissionTemplate = CreateAssignmentSubmission()
}

// AppendChild adds a child to the section
func (t *AssignmentSubmission) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *AssignmentSubmission) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *AssignmentSubmission) String() string {
	return t.format.String()
}

// Parse the template
func (t *AssignmentSubmission) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
