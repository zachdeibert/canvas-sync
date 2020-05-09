package canvassync

import (
	"fmt"
	"time"

	git "github.com/libgit2/git2go/v30"
	"github.com/zachdeibert/canvas-sync/task"
)

func writeDatabaseTask(dbPath string) func(*task.Task, func()) {
	return func(t *task.Task, finish func()) {
		p := t.CreateProgress(1)
		p.SetWork(2)
		// Open database
		db, err := git.OpenRepository(dbPath)
		if err != nil {
			panic(err)
		}
		defer db.Free()
		index, err := db.Index()
		if err != nil {
			panic(err)
		}
		defer index.Free()
		p.Finish(1)
		// Check for changes in the working tree and add them
		status, err := db.StatusList(&git.StatusOptions{
			Show:  git.StatusShowWorkdirOnly,
			Flags: git.StatusOptIncludeUntracked | git.StatusOptRecurseUntrackedDirs,
		})
		if err != nil {
			panic(err)
		}
		defer status.Free()
		count, err := status.EntryCount()
		if err != nil {
			panic(err)
		}
		p.AddWork(count)
		for i := 0; i < count; i++ {
			entry, err := status.ByIndex(i)
			if err != nil {
				panic(err)
			}
			changes := []string{}
			if len(entry.HeadToIndex.OldFile.Path) != 0 {
				if len(entry.HeadToIndex.NewFile.Path) != 0 && entry.HeadToIndex.OldFile.Path != entry.HeadToIndex.NewFile.Path {
					changes = []string{
						entry.HeadToIndex.NewFile.Path,
						entry.HeadToIndex.OldFile.Path,
					}
				} else {
					changes = []string{
						entry.HeadToIndex.OldFile.Path,
					}
				}
			} else if len(entry.HeadToIndex.NewFile.Path) != 0 {
				changes = []string{
					entry.HeadToIndex.NewFile.Path,
				}
			}
			index.AddAll(changes, 0, nil)
			p.Finish(1)
		}
		// Commit
		if count > 0 {
			treeID, err := index.WriteTree()
			if err != nil {
				panic(err)
			}
			if err = index.Write(); err != nil {
				panic(err)
			}
			tree, err := db.LookupTree(treeID)
			if err != nil {
				panic(err)
			}
			defer tree.Free()
			parent, _, err := db.RevparseExt("HEAD")
			if err != nil {
				panic(err)
			}
			parents := []*git.Commit{}
			if parent != nil {
				defer parent.Free()
				commit, err := db.LookupCommit(parent.Id())
				if err != nil {
					panic(err)
				}
				defer commit.Free()
				parents = []*git.Commit{
					commit,
				}
			}
			sig := &git.Signature{
				Name:  "Canvas Sync Utility",
				Email: "zachdeibert@gmail.com",
				When:  time.Now(),
			}
			if _, err = db.CreateCommit("HEAD", sig, sig, fmt.Sprintf("Canvas Sync at %s", sig.When.Format("Mon Jan 2 2006 15:04:05 MST")), tree, parents...); err != nil {
				panic(err)
			}
		}
		p.Finish(1)
		// Done!
		finish()
	}
}
