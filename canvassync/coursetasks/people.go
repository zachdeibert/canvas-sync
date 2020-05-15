package coursetasks

import (
	"fmt"
	"strings"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/csvgen"
	"github.com/zachdeibert/canvas-sync/task"
)

func init() {
	registerCSV("People", func(t *task.Task, c *canvas.Canvas, courseId int, csv csvgen.CSV) {
		users, err := c.CoursesListUsersInCourse(t.CreateProgress(1), nil, nil, nil, nil, nil, []canvas.CoursesListUsersInCourseInclude{
			canvas.CoursesListUsersInCourseIncludeEnrollments,
		}, nil, nil, nil, fmt.Sprint(courseId))
		if err != nil {
			if e, ok := err.(canvas.InvalidStatusCodeError); ok && e.Code == 401 {
				return
			}
			panic(err)
		}
		for _, user := range users {
			es := make([]string, len(user.Enrollments))
			for i, e := range user.Enrollments {
				if strings.HasSuffix(e.Role, "Enrollment") {
					es[i] = e.Role[0 : len(e.Role)-len("Enrollment")]
				} else {
					es[i] = e.Role
				}
			}
			role := strings.Join(es, ", ")
			if sortableParts := strings.Split(user.SortableName, ", "); len(sortableParts) == 2 {
				fmName := strings.Split(sortableParts[1], " ")
				csv.AddRow(user.ID, fmName[0], strings.Join(fmName[1:], " "), sortableParts[0], role)
			} else {
				csv.AddRow(user.ID, user.Name, "", "", role)
			}
		}
	}, "ID", "%d", "First Name", "%s", "Middle Name", "%s", "Last Name", "%s", "Role", "%s")
}
