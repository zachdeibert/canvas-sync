package coursetasks

import (
	"fmt"
	"net/url"
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

type fileEntry struct {
	Filename    string
	ModTime     *time.Time
	ContentType string
	URL         string
}

func init() {
	registerFileStructure("Files", func(p *task.Progress, c *canvas.Canvas, courseId int) ([]interface{}, error) {
		// apiGet
		files, err := c.FilesListFiles(p, nil, nil, nil, nil, nil, nil, nil, fmt.Sprint(courseId))
		if err != nil {
			return nil, err
		}
		folders, err := c.FilesListAllFolders(p, fmt.Sprint(courseId))
		if err != nil {
			return nil, err
		}
		folderMap := make(map[int]string)
		for _, folder := range folders {
			if folder.FullName == "course files" {
				folderMap[folder.ID] = ""
			} else {
				if strings.HasPrefix(folder.FullName, "course files/") {
					folder.FullName = folder.FullName[len("course files/"):]
				}
				folderMap[folder.ID] = path.Join(strings.Split(folder.FullName, "/")...)
			}
		}
		entries := make([]interface{}, len(files))
		for i, file := range files {
			basename, err := url.QueryUnescape(file.Filename)
			if err != nil {
				return nil, err
			}
			folder := folderMap[file.FolderID]
			filename := path.Join(folder, basename)
			newMod := file.ModifiedAt
			if file.UpdatedAt.After(newMod) {
				newMod = file.UpdatedAt
			}
			entries[i] = fileEntry{
				Filename:    filename,
				ModTime:     &newMod,
				ContentType: file.ContentType,
				URL:         file.URL,
			}
		}
		return entries, nil
	}, func(f interface{}) string {
		// getFilename
		return f.(fileEntry).Filename
	}, func(t *task.Task, c *canvas.Canvas, f interface{}) (*time.Time, error) {
		// determineLastModTime
		return f.(fileEntry).ModTime, nil
	}, func(t *task.Task, c *canvas.Canvas, f interface{}) ([]byte, error) {
		// downloadFile
		file := f.(fileEntry)
		data, _, err := c.RequestRaw(file.URL, file.ContentType, 10)
		return data, err
	})
}
