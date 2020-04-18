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
		repliedAnnouncements := []canvas.DiscussionTopic{}
		for _, announcement := range announcements {
			filename := path.Join(db, fmt.Sprintf("%d - %s.html", announcement.ID, InvalidPathRunes.ReplaceAllLiteralString(announcement.Title, "_")))
			content, err := ioutil.ReadFile(filename)
			if err == nil {
				doc := htmlgen.CreateDocument()
				str, ok := doc.Parse(string(content), html.AnnouncementChildCtors)
				if len(str) == 0 && ok {
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
			if announcement.DiscussionSubentryCount > 0 {
				repliedAnnouncements = append(repliedAnnouncements, announcement)
			} else {
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
				if err := ioutil.WriteFile(filename, []byte(doc.String()), 0644); err != nil {
					panic(err)
				}
				fileWrites.Finish(1)
			}
		}
	})
}
