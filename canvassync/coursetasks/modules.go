package coursetasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"regexp"
	"time"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

type moduleHandler struct {
	Types         []string
	FileExtension string
	Handler       func(*task.Task, *canvas.Canvas, canvas.ModuleItem, string, string) error
}

var (
	fileIDRegex = regexp.MustCompile("/([^/]+)$")

	moduleHandlers = []*moduleHandler{
		{
			Types:         []string{},
			FileExtension: ".json",
			Handler: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem, outFile string, metaFile string) error {
				data, err := json.Marshal(item)
				if err != nil {
					return err
				}
				return ioutil.WriteFile(outFile, data, 0644)
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
			Handler: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem, outFile string, metaFile string) error {
				return ioutil.WriteFile(outFile, []byte(fmt.Sprintf("%s %d\n%s\n", *item.Type, item.ContentID, item.URL)), 0644)
			},
		},
		{
			Types: []string{
				"File",
			},
			FileExtension: "",
			Handler: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem, outFile string, metaFile string) error {
				res := fileIDRegex.FindStringSubmatch(item.URL)
				if res == nil || len(res) != 2 {
					return fmt.Errorf("Invalid file URL '%s'", item.URL)
				}
				file, err := c.FilesGetFile(t.CreateProgress(0), nil, res[1])
				if err != nil {
					return err
				}
				if len(file.URL) > 0 {
					newMod := file.ModifiedAt
					if file.UpdatedAt.After(newMod) {
						newMod = file.UpdatedAt
					}
					if _, err1 := os.Stat(outFile); err1 == nil {
						if _, err2 := os.Stat(metaFile); err2 == nil {
							modRaw, err := ioutil.ReadFile(metaFile)
							if err != nil {
								return err
							}
							mod, err := time.Parse(time.RFC3339, string(modRaw))
							if math.Abs(mod.Sub(newMod).Seconds()) < 2 {
								return nil
							}
						} else if !os.IsNotExist(err2) {
							return err2
						}
					} else if !os.IsNotExist(err1) {
						return err1
					}
					data, _, err := c.RequestRaw(file.URL, file.ContentType, 10)
					if err != nil {
						return err
					}
					if err = ioutil.WriteFile(outFile, data, 0644); err != nil {
						return err
					}
					if err = ioutil.WriteFile(metaFile, []byte(newMod.Format(time.RFC3339)), 0644); err != nil {
						return err
					}
				}
				return nil
			},
		},
		{
			Types: []string{
				"ExternalUrl",
			},
			FileExtension: ".url",
			Handler: func(t *task.Task, c *canvas.Canvas, item canvas.ModuleItem, outFile string, metaFile string) error {
				return ioutil.WriteFile(outFile, []byte(fmt.Sprintf("%s\n", item.ExternalURL)), 0644)
			},
		},
	}
)

func init() {
	register("Modules", func(t *task.Task, c *canvas.Canvas, db string, courseId int, finish func()) {
		modules, err := c.ModulesListModules(t.CreateProgress(0.1), []canvas.ModulesListModulesInclude{
			canvas.ModulesListModulesIncludeItems,
		}, nil, nil, fmt.Sprint(courseId))
		if err != nil {
			if e, ok := err.(canvas.InvalidStatusCodeError); ok && e.Code == 401 {
				finish()
				return
			}
			panic(err)
		}
		p := t.CreateProgress(1)
		for _, module := range modules {
			p.AddWork(len(module.Items))
		}
		metaFolderRoot := path.Join(db, ".syncmeta")
		if err := os.MkdirAll(metaFolderRoot, 0755); err != nil {
			panic(err)
		}
		for _, module := range modules {
			modPart := fmt.Sprintf("%d - %s", module.ID, InvalidPathRunes.ReplaceAllLiteralString(module.Name, "_"))
			modDir := path.Join(db, modPart)
			modMetaDir := path.Join(metaFolderRoot, modPart)
			if err := os.Mkdir(modDir, 0755); err != nil && !os.IsExist(err) {
				panic(err)
			}
			if err := os.Mkdir(modMetaDir, 0755); err != nil && !os.IsExist(err) {
				panic(err)
			}
			for _, item := range module.Items {
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
				}
				if h == nil {
					h = moduleHandlers[0]
				}
				itemPart := fmt.Sprintf("%d - %s%s", item.ID, InvalidPathRunes.ReplaceAllLiteralString(item.Title, "_"), h.FileExtension)
				filename := path.Join(modDir, itemPart)
				metaFile := path.Join(modMetaDir, fmt.Sprintf("%s.txt", itemPart))
				if err := h.Handler(t, c, item, filename, metaFile); err != nil {
					panic(err)
				}
				p.Finish(1)
			}
		}
		finish()
	})
}
