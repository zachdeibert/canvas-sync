package canvassync

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/zachdeibert/canvas-sync/task"
)

func writeDatabaseTask(dbPath string) func(*task.Task, func()) {
	return func(t *task.Task, finish func()) {
		p := t.CreateProgress(1)
		p.SetWork(4)
		// Open database
		db, err := git.PlainOpen(dbPath)
		if err != nil {
			panic(err)
		}
		tree, err := db.Worktree()
		if err != nil {
			panic(err)
		}
		p.Finish(1)
		// Check for changes in the working tree
		status, err := tree.Status()
		if err != nil {
			panic(err)
		}
		hasChanges := false
		for change, s := range status {
			if !strings.HasPrefix(change, ".git") && (s.Worktree != git.Unmodified || s.Staging != git.Unmodified) {
				hasChanges = true
			}
		}
		p.Finish(1)
		// Add new files
		if hasChanges {
			if err = tree.AddGlob("*"); err != nil {
				panic(err)
			}
		}
		p.Finish(1)
		// Commit
		if hasChanges {
			if _, err = tree.Commit(fmt.Sprintf("Canvas Sync at %s", time.Now().Format("Mon Jan 2 2006 15:04:05 MST")), &git.CommitOptions{
				All: true,
				Author: &object.Signature{
					Name:  "Canvas Sync Utility",
					Email: "zachdeibert@gmail.com",
				},
				Committer: nil,
			}); err != nil {
				panic(err)
			}
		}
		p.Finish(1)
		// Done!
		finish()
	}
}
