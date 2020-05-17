package coursetasks

import (
	"fmt"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/canvassync/coursetasks/html"
	"github.com/zachdeibert/canvas-sync/htmlgen"
	"github.com/zachdeibert/canvas-sync/task"
)

type assignmentData struct {
	Assignment canvas.Assignment
	Submission *canvas.Submission
	LastUpdate time.Time
}

func init() {
	registerHTMLWithFileAttachments("Assignments", html.AssignmentChildCtor, func(p *task.Progress, c *canvas.Canvas, courseId int) ([]interface{}, error) {
		// apiGet
		l, err := c.AssignmentsListAssignments(p, nil, nil, nil, nil, nil, nil, nil, nil, fmt.Sprint(courseId))
		var o []interface{} = nil
		if l != nil {
			o = make([]interface{}, len(l))
			p.AddWork(len(l))
			np := task.CreateProgress()
			for i, v := range l {
				s, err := c.SubmissionsGetASingleSubmission(np, []canvas.SubmissionsGetASingleSubmissionInclude{
					canvas.SubmissionsGetASingleSubmissionIncludeSubmissionComments,
				}, fmt.Sprint(courseId), fmt.Sprint(v.ID), "self")
				if err != nil {
					return nil, err
				}
				update := v.UpdatedAt
				if s != nil {
					if s.GradedAt.After(update) {
						update = s.GradedAt
					}
					if s.PostedAt.After(update) {
						update = s.PostedAt
					}
					if s.SubmittedAt.After(update) {
						update = s.SubmittedAt
					}
					for _, c := range s.SubmissionComments {
						if c.CreatedAt.After(update) {
							update = c.CreatedAt
						}
						if c.EditedAt.After(update) {
							update = c.EditedAt
						}
					}
				}
				o[i] = assignmentData{
					Assignment: v,
					Submission: s,
					LastUpdate: update,
				}
				p.Finish(1)
			}
		}
		return o, err
	}, func(o interface{}) string {
		// getFilename
		a := o.(assignmentData)
		return fmt.Sprintf("%d - %s", a.Assignment.ID, a.Assignment.Name)
	}, func(o interface{}) []canvas.FileAttachment {
		// getAttachments
		a := o.(assignmentData)
		attachments := a.Submission.Attachments
		for _, comment := range a.Submission.SubmissionComments {
			attachments = append(attachments, comment.Attachments...)
		}
		return attachments
	}, func(o interface{}, doc *htmlgen.Document) bool {
		// isModified
		assignment := o.(assignmentData)
		children := doc.Children()
		if len(children) > 0 {
			if a, ok := children[0].(*html.Assignment); ok {
				if a.LastUpdate == assignment.LastUpdate {
					return false
				}
			}
		}
		return true
	}, func(o interface{}, doc *htmlgen.Document, c *canvas.Canvas, t *task.Task, courseId int) {
		// createDoc
		assignment := o.(assignmentData)
		doc.Title = assignment.Assignment.Name
		a := html.CreateAssignment()
		a.LastUpdate = assignment.LastUpdate
		a.Data = assignment.Assignment
		if assignment.Submission != nil {
			s := html.CreateAssignmentSubmission()
			s.Data = *assignment.Submission
			for _, comment := range assignment.Submission.SubmissionComments {
				c := html.CreateAssignmentSubmissionComment()
				c.Data = comment
				for _, attachment := range comment.Attachments {
					at := html.CreateAssignmentSubmissionAttachment()
					at.Data = attachment
					c.AppendChild(at)
				}
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
