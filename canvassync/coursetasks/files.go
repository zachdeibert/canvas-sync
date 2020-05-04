package coursetasks

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

type fileFolderMapData struct {
	RealFolder string
	MetaFolder string
}

func init() {
	register("Files", func(t *task.Task, c *canvas.Canvas, db string, courseId int, finish func()) {
		files, err := c.FilesListFiles(t.CreateProgress(0.1), nil, nil, nil, nil, nil, nil, nil, fmt.Sprint(courseId))
		if err != nil {
			if e, ok := err.(canvas.InvalidStatusCodeError); ok && e.Code == 401 {
				finish()
			} else {
				panic(err)
			}
		}
		downloadProg := t.CreateProgress(1)
		downloadProg.SetWork(len(files))
		folders, err := c.FilesListAllFolders(t.CreateProgress(0.1), fmt.Sprint(courseId))
		if err != nil {
			if e, ok := err.(canvas.InvalidStatusCodeError); ok && e.Code == 401 {
				finish()
			} else {
				panic(err)
			}
		}
		metaFolderRoot := path.Join(db, ".syncmeta")
		if err := os.MkdirAll(metaFolderRoot, 0755); err != nil {
			panic(err)
		}
		folderMap := make(map[int]fileFolderMapData)
		for _, folder := range folders {
			if folder.FullName == "course files" {
				folderMap[folder.ID] = fileFolderMapData{
					RealFolder: db,
					MetaFolder: metaFolderRoot,
				}
			} else {
				if strings.HasPrefix(folder.FullName, "course files/") {
					folder.FullName = folder.FullName[len("course files/"):]
				}
				part := path.Join(strings.Split(folder.FullName, "/")...)
				dirname := path.Join(db, part)
				if err := os.MkdirAll(dirname, 0755); err != nil {
					panic(err)
				}
				metaDirname := path.Join(metaFolderRoot, part)
				if err := os.MkdirAll(metaDirname, 0755); err != nil {
					panic(err)
				}
				folderMap[folder.ID] = fileFolderMapData{
					RealFolder: dirname,
					MetaFolder: metaDirname,
				}
			}
		}
		for _, file := range files {
			basename, err := url.QueryUnescape(file.Filename)
			if err != nil {
				panic(err)
			}
			folder := folderMap[file.FolderID]
			filename := path.Join(folder.RealFolder, basename)
			modFilename := path.Join(folder.MetaFolder, fmt.Sprintf("%s.txt", basename))
			newMod := file.ModifiedAt
			if file.UpdatedAt.After(newMod) {
				newMod = file.UpdatedAt
			}
			if _, err1 := os.Stat(filename); err1 == nil {
				if _, err2 := os.Stat(modFilename); err2 == nil {
					modRaw, err := ioutil.ReadFile(modFilename)
					if err != nil {
						panic(err)
					}
					mod, err := time.Parse(time.RFC3339, string(modRaw))
					if math.Abs(mod.Sub(newMod).Seconds()) < 2 {
						downloadProg.Finish(1)
						continue
					}
				} else if !os.IsNotExist(err2) {
					panic(err2)
				}
			} else if !os.IsNotExist(err1) {
				panic(err1)
			}
			data, _, err := c.RequestRaw(file.URL, file.ContentType, 10)
			if err != nil {
				panic(err)
			}
			if err := ioutil.WriteFile(filename, data, 0644); err != nil {
				panic(err)
			}
			if err := ioutil.WriteFile(modFilename, []byte(newMod.Format(time.RFC3339)), 0644); err != nil {
				panic(err)
			}
			downloadProg.Finish(1)
		}
		finish()
	})
}
