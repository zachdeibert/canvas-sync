package canvassync

import (
	"fmt"
	"path"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/canvassync/coursetasks"
	"github.com/zachdeibert/canvas-sync/task"
)

func courseTaskGroup(c *canvas.Canvas, db string, course courseDiscoveryResult) func(*task.Task, func()) {
	return func(t *task.Task, finish func()) {
		p := t.CreateProgress(1)
		children := coursetasks.CreateTasks(t, c, path.Join(db, fmt.Sprintf("%d - %s", course.id, course.name)), course.id)
		p.SetWork(len(children))
		listener := func(_ *task.Task) {
			p.Finish(1)
			if p.GetStatus() == 1 {
				finish()
			}
		}
		for _, child := range children {
			child.AddFinishListener(listener)
		}
	}
}
