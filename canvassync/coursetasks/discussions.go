package coursetasks

import (
	"fmt"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/canvassync/coursetasks/html"
	"github.com/zachdeibert/canvas-sync/htmlgen"
	"github.com/zachdeibert/canvas-sync/task"
)

func init() {
	registerHTMLWithFileAttachments("Discussions", html.DiscussionRootChildCtor, func(p *task.Progress, c *canvas.Canvas, courseId int) ([]interface{}, error) {
		// apiGet
		l, err := c.DiscussionTopicsListDiscussionTopics(p, nil, nil, nil, nil, nil, nil, nil, fmt.Sprint(courseId))
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
		a := o.(canvas.DiscussionTopic)
		return fmt.Sprintf("%d - %s", a.ID, a.Title)
	}, func(o interface{}) []canvas.FileAttachment {
		// getAttachments
		return o.(canvas.DiscussionTopic).Attachments
	}, func(o interface{}, doc *htmlgen.Document) bool {
		// isModified
		topic := o.(canvas.DiscussionTopic)
		children := doc.Children()
		if len(children) > 0 {
			if a, ok := children[0].(*html.DiscussionRoot); ok {
				if a.Data.LastReplyAt == topic.LastReplyAt {
					return false
				}
			}
		}
		return true
	}, func(o interface{}, doc *htmlgen.Document, c *canvas.Canvas, t *task.Task, courseId int) {
		// createDoc
		topic := o.(canvas.DiscussionTopic)
		doc.Title = topic.Title
		d := html.CreateDiscussionRoot()
		d.Data = topic
		if topic.DiscussionSubentryCount > 0 && topic.UserCanSeePosts {
			view, err := c.DiscussionTopicsGetTheFullTopic(t.CreateProgress(0.01), fmt.Sprint(courseId), fmt.Sprint(topic.ID))
			if err != nil {
				panic(err)
			}
			for _, rv1 := range view.View {
				r1 := html.CreateDiscussionReply()
				r1.Data = rv1
				for _, u := range view.Participants {
					if u.ID == rv1.UserID {
						r1.User = u
						break
					}
				}
				for _, rv2 := range rv1.Replies {
					r2 := html.CreateDiscussionReply()
					r2.Data = rv2
					for _, u := range view.Participants {
						if u.ID == rv2.UserID {
							r2.User = u
							break
						}
					}
					r1.AppendChild(r2)
				}
				d.AppendChild(r1)
			}
		}
		doc.AppendChild(d)
	})
}
