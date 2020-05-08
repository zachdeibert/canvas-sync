package coursetasks

import (
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

type moduleHandler struct {
	Types            []string
	FileExtension    string
	DetermineModTime func(*task.Task, *canvas.Canvas, canvas.ModuleItem) (*time.Time, interface{}, error)
	Download         func(*task.Task, *canvas.Canvas, canvas.ModuleItem, interface{}) ([]byte, error)
}

type moduleEntry struct {
	Filename string
	Item     canvas.ModuleItem
	Data     interface{}
	Handler  moduleHandler
}

var (
	fileIDRegex = regexp.MustCompile("/([^/]+)$")

	moduleHandlers = []*moduleHandler{
		{
			Types:         []string{},
			FileExtension: ".json",
			DetermineModTime: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem) (*time.Time, interface{}, error) {
				return nil, nil, nil
			},
			Download: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem, data interface{}) ([]byte, error) {
				return json.Marshal(item)
			},
		},
		{
			Types: []string{
				"Page",
				"Discussion",
				"Assignment",
				"Quiz",
				"SubHeader",
				"ExternalTool",
			},
			FileExtension: ".txt",
			DetermineModTime: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem) (*time.Time, interface{}, error) {
				return nil, nil, nil
			},
			Download: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem, data interface{}) ([]byte, error) {
				return []byte(fmt.Sprintf("%s %d\n%s\n", *item.Type, item.ContentID, item.URL)), nil
			},
		},
		{
			Types: []string{
				"File",
			},
			FileExtension: "",
			DetermineModTime: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem) (*time.Time, interface{}, error) {
				res := fileIDRegex.FindStringSubmatch(item.URL)
				if res == nil || len(res) != 2 {
					return nil, nil, fmt.Errorf("Invalid file URL '%s'", item.URL)
				}
				file, err := c.FilesGetFile(t.CreateProgress(0), nil, res[1])
				if err != nil {
					return nil, nil, err
				}
				if len(file.URL) > 0 {
					newMod := file.ModifiedAt
					if file.UpdatedAt.After(newMod) {
						newMod = file.UpdatedAt
					}
					return &newMod, []string{
						file.URL,
						file.ContentType,
					}, nil
				}
				return nil, nil, nil
			},
			Download: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem, data interface{}) ([]byte, error) {
				if data == nil {
					return nil, errFileLocked
				}
				d := data.([]string)
				res, _, err := c.RequestRaw(d[0], d[1], 10)
				return res, err
			},
		},
		{
			Types: []string{
				"ExternalUrl",
			},
			FileExtension: ".url",
			DetermineModTime: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem) (*time.Time, interface{}, error) {
				return nil, nil, nil
			},
			Download: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem, data interface{}) ([]byte, error) {
				return []byte(fmt.Sprintf("%s\n", item.ExternalURL)), nil
			},
		},
	}
)

func init() {
	registerFileStructure("Modules", func(p *task.Progress, c *canvas.Canvas, courseId int) ([]interface{}, error) {
		// apiGet
		modules, err := c.ModulesListModules(p, []canvas.ModulesListModulesInclude{
			canvas.ModulesListModulesIncludeItems,
		}, nil, nil, fmt.Sprint(courseId))
		if err != nil {
			return nil, err
		}
		res := []interface{}{}
		for _, module := range modules {
			modPart := fmt.Sprintf("%d - %s", module.ID, InvalidPathRunes.ReplaceAllLiteralString(module.Name, "_"))
			entries := make([]interface{}, len(module.Items))
			for i, item := range module.Items {
				var h *moduleHandler = nil
				if item.Type != nil {
					for _, handler := range moduleHandlers {
						for _, t := range handler.Types {
							if t == string(*item.Type) {
								h = handler
								break
							}
						}
						if h != nil {
							break
						}
					}
					itemPart := fmt.Sprintf("%d - %s%s", item.ID, InvalidPathRunes.ReplaceAllLiteralString(item.Title, "_"), h.FileExtension)
					entries[i] = &moduleEntry{
						Filename: path.Join(modPart, itemPart),
						Item:     item,
						Handler:  *h,
					}
				}
				if h == nil {
					h = moduleHandlers[0]
				}
			}
			res = append(res, entries...)
		}
		return res, nil
	}, func(f interface{}) string {
		// getFilename
		return f.(*moduleEntry).Filename
	}, func(t *task.Task, c *canvas.Canvas, f interface{}) (*time.Time, error) {
		// determineLastModTime
		entry := f.(*moduleEntry)
		tm, data, err := entry.Handler.DetermineModTime(t, c, entry.Item)
		entry.Data = data
		return tm, err
	}, func(t *task.Task, c *canvas.Canvas, f interface{}) ([]byte, error) {
		// downloadFile
		entry := f.(*moduleEntry)
		return entry.Handler.Download(t, c, entry.Item, entry.Data)
	})
}
