package coursetasks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/canvas/model"
	"github.com/zachdeibert/canvas-sync/canvassync/coursetasks/html"
	"github.com/zachdeibert/canvas-sync/htmlgen"
	"github.com/zachdeibert/canvas-sync/task"
)

func init() {
	register("Announcements", func(t *task.Task, c *canvas.Canvas, db string, courseId int, finish func()) {
		defer finish()
		announcements := []model.Announcement{}
		if err := c.Request("announcements", map[string]interface{}{
			"context_codes": []string{
				fmt.Sprintf("course_%d", courseId),
			},
			"start_date": "1970-01-01",
			"end_date":   time.Now().Add(time.Hour * 24).Format("2006-01-02"),
		}, t.CreateProgress(1), func() interface{} {
			return &[]model.Announcement{}
		}, func(obj interface{}) error {
			part := *obj.(*[]model.Announcement)
			announcements = append(announcements, part...)
			return nil
		}); err != nil {
			panic(err)
		}
		fileWrites := t.CreateProgress(1)
		fileWrites.SetWork(len(announcements))
		repliedAnnouncements := []model.Announcement{}
		for _, announcement := range announcements {
			filename := path.Join(db, fmt.Sprintf("%d - %s.html", announcement.ID, InvalidPathRunes.ReplaceAllLiteralString(announcement.Title, "_")))
			content, err := ioutil.ReadFile(filename)
			if err == nil {
				doc := htmlgen.CreateDocument()
				str, ok := doc.Parse(string(content), html.AnnouncementChildCtors)
				if len(str) == 0 && ok {
					children := doc.Children()
					if len(children) > 0 {
						switch a := children[0].(type) {
						case *html.Announcement:
							if a.Data.LastReplyAt == announcement.LastReplyAt {
								fileWrites.Finish(1)
								continue
							}
							break
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
