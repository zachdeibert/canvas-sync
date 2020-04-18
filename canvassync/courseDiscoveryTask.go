package canvassync

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

type courseDiscoveryResult struct {
	id   int
	name string
}

func courseDiscoveryTask(c *canvas.Canvas, coursesCh chan<- []courseDiscoveryResult) func(*task.Task, func()) {
	return func(t *task.Task, finish func()) {
		courses, err := c.CoursesListYourCourses(t.CreateProgress(1), nil, nil, nil, nil, nil, nil, nil)
		if err != nil {
			panic(err)
		}
		res := []courseDiscoveryResult{}
		for _, course := range courses {
			if course.Name != "" {
				res = append(res, courseDiscoveryResult{
					id:   course.ID,
					name: course.Name,
				})
			}
		}
		coursesCh <- res
		finish()
	}
}
