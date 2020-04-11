package canvassync

import (
	"../canvas"
	"../canvas/model"
	"../task"
)

type courseDiscoveryResult struct {
	id   int
	name string
}

func courseDiscoveryTask(c *canvas.Canvas, coursesCh chan<- []courseDiscoveryResult) func(*task.Task, func()) {
	return func(t *task.Task, finish func()) {
		res := []courseDiscoveryResult{}
		if err := c.Request("courses", nil, func() interface{} {
			return &[]model.Course{}
		}, func(obj interface{}) error {
			courses := *obj.(*[]model.Course)
			for _, course := range courses {
				if course.Name != "" {
					res = append(res, courseDiscoveryResult{
						id:   course.ID,
						name: course.Name,
					})
				}
			}
			return nil
		}); err != nil {
			panic(err)
		}
		coursesCh <- res
		finish()
	}
}
