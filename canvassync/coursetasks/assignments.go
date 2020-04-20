package coursetasks

import (
	"fmt"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/canvassync/coursetasks/html"
	"github.com/zachdeibert/canvas-sync/htmlgen"
	"github.com/zachdeibert/canvas-sync/task"
)

func init() {
	registerHTMLWithFileAttachments("Assignments", html.AssignmentChildCtor, func(p *task.Progress, c *canvas.Canvas, courseId int) ([]interface{}, error) {
		// apiGet
		l, err := c.AssignmentsListAssignments(p, []canvas.AssignmentsListAssignmentsInclude{
			canvas.AssignmentsListAssignmentsIncludeSubmission,
		}, nil, nil, nil, nil, nil, nil, nil, fmt.Sprint(courseId))
		var o []interface{} = nil
		if l != nil {
			o = make([]interface{}, len(l))
			for i, v := range l {
				o[i] = v
			}
		}
		return o, err
	}, func(o interface{}) string {
		// getFilename
		a := o.(canvas.Assignment)
		return fmt.Sprintf("%d - %s", a.ID, a.Name)
	}, func(o interface{}) []canvas.FileAttachment {
		// getAttachments
		return o.(canvas.Assignment).Submission.Attachments
	}, func(o interface{}, doc *htmlgen.Document) bool {
		// isModified
		assignment := o.(canvas.Assignment)
		children := doc.Children()
		if len(children) > 0 {
			if a, ok := children[0].(*html.Assignment); ok {
				if a.Data.UpdatedAt == assignment.UpdatedAt {
					return false
				}
			}
		}
		return true
	}, func(o interface{}, doc *htmlgen.Document, c *canvas.Canvas, t *task.Task, courseId int) {
		// createDoc
		assignment := o.(canvas.Assignment)
		doc.Title = assignment.Name
		a := html.CreateAssignment()
		a.Data = assignment
		if assignment.Submission != nil {
			s := html.CreateAssignmentSubmission()
			s.Data = *assignment.Submission
			for _, comment := range assignment.Submission.SubmissionComments {
				c := html.CreateAssignmentSubmissionComment()
				c.Data = comment
				s.AppendChild(c)
			}
			for _, attachment := range assignment.Submission.Attachments {
				at := html.CreateAssignmentSubmissionAttachment()
				at.Data = attachment
				s.AppendChild(at)
			}
			a.AppendChild(s)
		}
		doc.AppendChild(a)
	})
}
