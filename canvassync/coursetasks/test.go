package coursetasks

import (
	"math/rand"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

func init() {
	f := func(t *task.Task, c *canvas.Canvas, db string, courseId int, finish func()) {
		defer finish()
		p := t.CreateProgress(1)
		p.SetWork(20)
		timer := time.NewTicker(time.Millisecond * time.Duration(50+rand.Float32()*100))
		defer timer.Stop()
		for p.GetStatus() < 1 {
			<-timer.C
			p.Finish(1)
		}
	}
	register("Test 1", f)
	register("Test 2", f)
	register("Test 3", f)
	register("Test 4", f)
}
