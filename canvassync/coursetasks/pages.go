package coursetasks

import (
	"fmt"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/canvassync/coursetasks/html"
	"github.com/zachdeibert/canvas-sync/htmlgen"
	"github.com/zachdeibert/canvas-sync/task"
)

func init() {
	registerHTML("Pages", html.PageChildCtor, func(p *task.Progress, c *canvas.Canvas, courseId int) ([]interface{}, error) {
		// apiGet
		pages, err := c.PagesListPages(p, nil, nil, nil, nil, fmt.Sprint(courseId))
		var o []interface{} = nil
		if pages != nil {
			o = make([]interface{}, len(pages))
			for i, v := range pages {
				o[i] = v
			}
		}
		return o, err
	}, func(o interface{}) string {
		// getFilename
		return o.(canvas.Page).URL
	}, func(o interface{}, doc *htmlgen.Document) bool {
		// isModified
		page := o.(canvas.Page)
		children := doc.Children()
		if len(children) > 0 {
			if a, ok := children[0].(*html.Page); ok {
				if a.Data.UpdatedAt == page.UpdatedAt {
					return false
				}
			}
		}
		return true
	}, func(o interface{}, doc *htmlgen.Document, c *canvas.Canvas, t *task.Task, courseId int) {
		// createDoc
		page, err := c.PagesShowPage(t.CreateProgress(1), fmt.Sprint(courseId), o.(canvas.Page).URL)
		if err != nil {
			panic(err)
		}
		doc.Title = page.Title
		p := html.CreatePage()
		p.Data = *page
		if page.LastEditedBy == nil {
			p.Editor = "unknown"
		} else {
			p.Editor = page.LastEditedBy.DisplayName
		}
		doc.AppendChild(p)
	})
}
