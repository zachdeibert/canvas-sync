package coursetasks

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

var errFileLocked = fmt.Errorf("File locked")

func registerFileStructure(name string,
	apiGet func(*task.Progress, *canvas.Canvas, int) ([]interface{}, error),
	getFilename func(interface{}) string,
	determineLastModTime func(*task.Task, *canvas.Canvas, interface{}) (*time.Time, error),
	downloadFile func(*task.Task, *canvas.Canvas, interface{}) ([]byte, error)) {
	register(name, func(t *task.Task, c *canvas.Canvas, db string, courseId int, finish func()) {
		files, err := apiGet(t.CreateProgress(0.1), c, courseId)
		if err != nil {
			if e, ok := err.(canvas.InvalidStatusCodeError); ok && e.Code == 401 {
				finish()
				return
			}
			panic(err)
		}
		p := t.CreateProgress(1)
		p.SetWork(len(files))
		metaFolderRoot := path.Join(db, ".syncmeta")
		for _, file := range files {
			name := getFilename(file)
			filename := path.Join(db, name)
			modFile := path.Join(metaFolderRoot, fmt.Sprintf("%s.txt", name))
			if err := os.MkdirAll(path.Dir(filename), 0755); err != nil {
				panic(err)
			}
			if err := os.MkdirAll(path.Dir(modFile), 0755); err != nil {
				panic(err)
			}
			newMod, err := determineLastModTime(t, c, file)
			if err != nil {
				panic(err)
			}
			if newMod != nil {
				if _, err1 := os.Stat(filename); err1 == nil {
					if _, err2 := os.Stat(modFile); err2 == nil {
						modRaw, err := ioutil.ReadFile(modFile)
						if err != nil {
							panic(err)
						}
						mod, err := time.Parse(time.RFC3339, string(modRaw))
						if math.Abs(mod.Sub(*newMod).Seconds()) < 2 {
							p.Finish(1)
							continue
						}
					} else if !os.IsNotExist(err2) {
						panic(err2)
					}
				} else if !os.IsNotExist(err1) {
					panic(err1)
				}
			}
			data, err := downloadFile(t, c, file)
			if err != errFileLocked {
				if err != nil {
					panic(err)
				}
				if err = ioutil.WriteFile(filename, data, 0644); err != nil {
					panic(err)
				}
				if newMod != nil {
					if err = ioutil.WriteFile(modFile, []byte(newMod.Format(time.RFC3339)), 0644); err != nil {
						panic(err)
					}
				}
			}
			p.Finish(1)
		}
		finish()
	})
}
