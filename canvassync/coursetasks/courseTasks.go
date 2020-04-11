package coursetasks

import (
	"os"
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

func createTask(d courseTask, c *canvas.Canvas, db string, courseID int) func(*task.Task, func()) {
	return func(t *task.Task, finish func()) {
		dir := path.Join(db, d.name)
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
		d.f(t, c, dir, courseID, finish)
	}
}

// CreateTasks creates all the tasks for a course under a parent task
func CreateTasks(parent *task.Task, c *canvas.Canvas, db string, courseID int) []*task.Task {
	res := make([]*task.Task, len(tasks))
	for i, d := range tasks {
		res[i] = parent.CreateSubtask(d.name, createTask(d, c, db, courseID))
	}
	return res
}
