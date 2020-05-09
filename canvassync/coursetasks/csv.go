package coursetasks

import (
	"fmt"
	"path"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/csvgen"
	"github.com/zachdeibert/canvas-sync/task"
)

func registerCSV(name string, genCSV func(*task.Task, *canvas.Canvas, int, csvgen.CSV), cols ...string) {
	if len(cols)%2 != 0 {
		panic("Invalid column spec for CSV")
	}
	register(name, func(t *task.Task, c *canvas.Canvas, db string, courseId int, finish func()) {
		csv := csvgen.CreateCSV()
		for i := 0; i < len(cols); i += 2 {
			csv.AddColumn(cols[i], cols[i+1])
		}
		genCSV(t, c, courseId, csv)
		if err := csv.WriteFile(path.Join(db, fmt.Sprintf("%s.csv", name))); err != nil {
			panic(err)
		}
		finish()
	})
}
