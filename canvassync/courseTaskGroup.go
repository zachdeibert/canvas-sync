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
		t.InheritProgress()
		var done int = 0
		children := coursetasks.CreateTasks(t, c, path.Join(db, fmt.Sprintf("%d - %s", course.id, coursetasks.InvalidPathRunes.ReplaceAllLiteralString(course.name, "_"))), course.id)
		listener := func(_ *task.Task) {
			if done++; done == len(children) {
				finish()
			}
		}
		for _, child := range children {
			child.AddFinishListener(listener)
		}
	}
}
