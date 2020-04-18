package coursetasks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/canvassync/coursetasks/html"
	"github.com/zachdeibert/canvas-sync/htmlgen"
	"github.com/zachdeibert/canvas-sync/task"
)

func init() {
	register("Announcements", func(t *task.Task, c *canvas.Canvas, db string, courseId int, finish func()) {
		defer finish()
		startDate := time.Unix(0, 0)
		endDate := time.Now().Add(time.Hour * 24)
		announcements, err := c.AnnouncementsListAnnouncements(t.CreateProgress(1), []string{
			fmt.Sprintf("course_%d", courseId),
		}, &startDate, &endDate, nil, nil)
		if err != nil {
			panic(err)
		}
		fileWrites := t.CreateProgress(1)
		fileWrites.SetWork(len(announcements))
		for _, announcement := range announcements {
			filename := path.Join(db, fmt.Sprintf("%d - %s.html", announcement.ID, InvalidPathRunes.ReplaceAllLiteralString(announcement.Title, "_")))
			content, err := ioutil.ReadFile(filename)
			if err == nil {
				doc := htmlgen.ParseDocument(string(content), []htmlgen.ChildConstructor{html.AnnouncementChildCtor})
				if doc != nil {
					children := doc.Children()
					if len(children) > 0 {
						if a, ok := children[0].(*html.Announcement); ok {
							if a.Data.LastReplyAt == announcement.LastReplyAt {
								fileWrites.Finish(1)
								continue
							}
						}
					}
				}
			} else if !os.IsNotExist(err) {
				panic(err)
			}
			doc := htmlgen.CreateDocument()
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
			if err := ioutil.WriteFile(filename, []byte(doc.String()), 0644); err != nil {
				panic(err)
			}
			fileWrites.Finish(1)
		}
	})
}
