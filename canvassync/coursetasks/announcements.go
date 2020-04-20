package coursetasks

import (
	"fmt"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/canvassync/coursetasks/html"
	"github.com/zachdeibert/canvas-sync/htmlgen"
	"github.com/zachdeibert/canvas-sync/task"
)

func init() {
	registerHTMLWithFileAttachments("Announcements", html.AssignmentChildCtor, func(p *task.Progress, c *canvas.Canvas, courseId int) ([]interface{}, error) {
		// apiGet
		startDate := time.Unix(0, 0)
		endDate := time.Now().Add(time.Hour * 24)
		a, err := c.AnnouncementsListAnnouncements(p, []string{
			fmt.Sprintf("course_%d", courseId),
		}, &startDate, &endDate, nil, nil)
		var o []interface{} = nil
		if a != nil {
			o = make([]interface{}, len(a))
			for i, v := range a {
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
		announcement := o.(canvas.DiscussionTopic)
		children := doc.Children()
		if len(children) > 0 {
			if a, ok := children[0].(*html.Announcement); ok {
				if a.Data.LastReplyAt == announcement.LastReplyAt {
					return false
				}
			}
		}
		return true
	}, func(o interface{}, doc *htmlgen.Document, c *canvas.Canvas, t *task.Task, courseId int) {
		// createDoc
		announcement := o.(canvas.DiscussionTopic)
		doc.Title = announcement.Title
		a := html.CreateAnnouncement()
		a.Data = announcement
		for _, attachment := range announcement.Attachments {
			at := html.CreateAnnouncementAttachment()
			at.Data = attachment
			a.AppendChild(at)
		}
		doc.AppendChild(a)
		if announcement.DiscussionSubentryCount > 0 {
			view, err := c.DiscussionTopicsGetTheFullTopic(t.CreateProgress(0.01), fmt.Sprint(courseId), fmt.Sprint(announcement.ID))
			if err != nil {
				panic(err)
			}
			for _, rv1 := range view.View {
				r1 := html.CreateAnnouncementReply()
				r1.Data = rv1
				for _, u := range view.Participants {
					if u.ID == rv1.UserID {
						r1.User = u
						break
					}
				}
				for _, rv2 := range rv1.Replies {
					r2 := html.CreateAnnouncementReply()
					r2.Data = rv2
					for _, u := range view.Participants {
						if u.ID == rv2.UserID {
							r2.User = u
							break
						}
					}
					r1.AppendChild(r2)
				}
				a.AppendChild(r1)
			}
		}
	})
}
