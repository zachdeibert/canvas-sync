package canvassync

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	git "github.com/libgit2/git2go/v30"
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

func databaseCheckTask(c *canvas.Canvas, name chan<- string, dbCh chan<- string) func(*task.Task, func()) {
	return func(t *task.Task, finish func()) {
		p := t.CreateProgress(1)
		p.SetWork(5)
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
		g, err := git.OpenRepository(db)
		if err != nil {
			if g, err = git.InitRepository(db, false); err != nil {
				panic(err)
			}
		}
		defer g.Free()
		p.Finish(1)
		// Check for working directory changes
		status, err := g.StatusList(&git.StatusOptions{
			Show:  git.StatusShowIndexAndWorkdir,
			Flags: git.StatusOptIncludeUntracked,
		})
		if err != nil {
			panic(err)
		}
		defer status.Free()
		count, err := status.EntryCount()
		if err != nil {
			panic(err)
		}
		if count != 0 {
			str := &strings.Builder{}
			str.WriteString("Database is not clean; Previous run of program did not exit cleanly.\n")
			for i := 0; i < count; i++ {
				entry, err := status.ByIndex(i)
				if err == nil {
					path := entry.IndexToWorkdir.NewFile.Path
					if len(path) == 0 {
						path = entry.IndexToWorkdir.OldFile.Path
					}
					if len(path) == 0 {
						path = entry.HeadToIndex.NewFile.Path
					}
					if len(path) == 0 {
						path = entry.HeadToIndex.OldFile.Path
					}
					fmt.Fprintf(str, "  %s\n", path)
				}
			}
			panic(str.String())
		}
		p.Finish(1)
		// Copy old raw request cache over to new folder
		oldDir := c.RawSaveFolder
		c.RawSaveFolder = path.Join(db, ".raw")
		files := []string{}
		if err = filepath.Walk(oldDir, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				files = append(files, p[len(oldDir):])
			}
			return nil
		}); err != nil {
			panic(err)
		}
		for _, file := range files {
			dest := path.Join(c.RawSaveFolder, file)
			if err = os.MkdirAll(path.Dir(dest), 0755); err != nil {
				panic(err)
			}
			if err = os.Rename(path.Join(oldDir, file), dest); err != nil {
				panic(err)
			}
		}
		p.Finish(1)
		// Done!
		finish()
	}
}
