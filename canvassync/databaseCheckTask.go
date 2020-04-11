package canvassync

import (
	"fmt"
	"os"
	"path"

	"github.com/go-git/go-git/v5"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/canvas/model"
	"github.com/zachdeibert/canvas-sync/task"
)

func databaseCheckTask(c *canvas.Canvas, name chan<- string, dbCh chan<- string) func(*task.Task, func()) {
	return func(t *task.Task, finish func()) {
		p := t.CreateProgress(1)
		p.SetWork(4)
		if err := c.Request("users/self", nil, t.CreateProgress(1), func() interface{} {
			return &model.User{}
		}, func(obj interface{}) error {
			// Get name and find database path
			user := obj.(*model.User)
			name <- user.Name
			db := path.Join("db", c.GetSubdomain(), fmt.Sprintf("%d - %s", user.ID, user.Name))
			dbCh <- db
			p.Finish(1)
			// Ensure database folder exists
			if err := os.MkdirAll(db, 0755); err != nil {
				panic(err)
			}
			p.Finish(1)
			// Ensure Git exists
			g, err := git.PlainOpen(db)
			if err != nil {
				if err == git.ErrRepositoryNotExists {
					if g, err = git.PlainInit(db, false); err != nil {
						panic(err)
					}
				} else {
					panic(err)
				}
			}
			p.Finish(1)
			// Check for working directory changes
			tree, err := g.Worktree()
			if err != nil {
				panic(err)
			}
			status, err := tree.Status()
			if err != nil {
				panic(err)
			}
			if !status.IsClean() {
				panic("Database is not clean; Previous run of program did not exit cleanly.")
			}
			p.Finish(1)
			// Done!
			finish()
			return nil
		}); err != nil {
			panic(err)
		}
	}
}
