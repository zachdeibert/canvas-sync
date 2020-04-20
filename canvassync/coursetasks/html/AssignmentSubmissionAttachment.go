package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	assignmentSubmissionAttachmentTemplate *AssignmentSubmissionAttachment
	// AssignmentSubmissionAttachmentChildCtor for parsing a template
	AssignmentSubmissionAttachmentChildCtor = func() (htmlgen.Section, []htmlgen.ChildConstructor) {
		return CreateAssignmentSubmissionAttachment(), []htmlgen.ChildConstructor{}
	}
)

// AssignmentSubmissionAttachment HTML template
type AssignmentSubmissionAttachment struct {
	Data   canvas.FileAttachment
	format *htmlgen.FormatSection
}

// CreateAssignmentSubmissionAttachment creates a new template
func CreateAssignmentSubmissionAttachment() *AssignmentSubmissionAttachment {
	obj := &AssignmentSubmissionAttachment{}
	args := []interface{}{
		&obj.Data.Filename,
		&obj.Data.DisplayName,
	}
	if assignmentSubmissionAttachmentTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	<p>Attached file <a href="%s">%s</a>.</p>
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = assignmentSubmissionAttachmentTemplate.format.Clone(args)
	}
	return obj
}

func init() {
	assignmentSubmissionAttachmentTemplate = CreateAssignmentSubmissionAttachment()
}

// AppendChild adds a child to the section
func (t *AssignmentSubmissionAttachment) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *AssignmentSubmissionAttachment) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *AssignmentSubmissionAttachment) String() string {
	return t.format.String()
}

// Parse the template
func (t *AssignmentSubmissionAttachment) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
