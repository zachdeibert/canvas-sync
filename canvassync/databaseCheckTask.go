package canvassync

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

func databaseCheckTask(c *canvas.Canvas, name chan<- string, dbCh chan<- string) func(*task.Task, func()) {
	return func(t *task.Task, finish func()) {
		p := t.CreateProgress(1)
		p.SetWork(4)
		user, err := c.UsersShowUserDetails(t.CreateProgress(1), nil)
		if err != nil {
			panic(err)
		}
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
		for change, s := range status {
			if !strings.HasPrefix(change, ".git") && (s.Worktree != git.Unmodified || s.Staging != git.Unmodified) {
				panic(fmt.Sprintf("Database is not clean; Previous run of program did not exit cleanly.\n%s", status.String()))
			}
		}
		p.Finish(1)
		// Done!
		finish()
	}
}
