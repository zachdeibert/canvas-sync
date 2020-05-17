package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	assignmentSubmissionCommentTemplate *AssignmentSubmissionComment
	// AssignmentSubmissionCommentChildCtor for parsing a template
	AssignmentSubmissionCommentChildCtor = func() (htmlgen.Section, []htmlgen.ChildConstructor) {
		return CreateAssignmentSubmissionComment(), []htmlgen.ChildConstructor{
			AssignmentSubmissionAttachmentChildCtor,
		}
	}
)

// AssignmentSubmissionComment HTML template
type AssignmentSubmissionComment struct {
	Data   canvas.SubmissionComment
	format *htmlgen.FormatSection
}

// CreateAssignmentSubmissionComment creates a new template
func CreateAssignmentSubmissionComment() *AssignmentSubmissionComment {
	obj := &AssignmentSubmissionComment{}
	args := []interface{}{
		&obj.Data.AuthorName,
		htmlgen.CreateDateTimeFormat(&obj.Data.EditedAt),
		&obj.Data.Comment,
		htmlgen.FormatSectionChild,
	}
	if assignmentSubmissionCommentTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	<h3>Comment from %s at %s:</h3>
	<div>
		%s
	</div>
	<div>
		%s
	</div>
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = assignmentSubmissionCommentTemplate.format.Clone(args)
	}
	return obj
}

func init() {
	assignmentSubmissionCommentTemplate = CreateAssignmentSubmissionComment()
}

// AppendChild adds a child to the section
func (t *AssignmentSubmissionComment) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *AssignmentSubmissionComment) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *AssignmentSubmissionComment) String() string {
	return t.format.String()
}

// Parse the template
func (t *AssignmentSubmissionComment) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
