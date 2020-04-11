package coursetasks

import (
	"path"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

type courseTask struct {
	name string
	f    func(*task.Task, *canvas.Canvas, string, int, func())
}

var tasks []courseTask

func register(name string, f func(*task.Task, *canvas.Canvas, string, int, func())) {
	t := courseTask{
		name: name,
		f:    f,
	}
	if tasks == nil {
		tasks = []courseTask{t}
	} else {
		tasks = append(tasks, t)
	}
}

// CreateTasks creates all the tasks for a course under a parent task
func CreateTasks(parent *task.Task, c *canvas.Canvas, db string, courseID int) []*task.Task {
	res := make([]*task.Task, len(tasks))
	for i, d := range tasks {
		res[i] = parent.CreateSubtask(d.name, func(t *task.Task, finish func()) {
			d.f(t, c, path.Join(db, d.name), courseID, finish)
		})
	}
	return res
}
