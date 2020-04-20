package coursetasks

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
	"github.com/zachdeibert/canvas-sync/task"
)

func registerHTMLWithFileAttachments(name string, docType htmlgen.ChildConstructor,
	apiGet func(*task.Progress, *canvas.Canvas, int) ([]interface{}, error),
	getFilename func(interface{}) string,
	getAttachments func(interface{}) []canvas.FileAttachment,
	isModified func(interface{}, *htmlgen.Document) bool,
	createDoc func(interface{}, *htmlgen.Document, *canvas.Canvas, *task.Task, int)) {
	registerHTMLWithAttachments(name, docType, apiGet, getFilename, func(o interface{}) []interface{} {
		a := getAttachments(o)
		b := make([]interface{}, len(a))
		for i, v := range a {
			b[i] = v
		}
		return b
	}, func(a interface{}) string {
		return a.(canvas.FileAttachment).Filename
	}, func(a interface{}, filename string) {
		// TODO
	}, func(a interface{}, filename string) bool {
		return false
	}, isModified, createDoc)
}

func registerHTMLWithAttachments(name string, docType htmlgen.ChildConstructor,
	apiGet func(*task.Progress, *canvas.Canvas, int) ([]interface{}, error),
	getFilename func(interface{}) string,
	getAttachments func(interface{}) []interface{},
	getAttachmentFilename func(interface{}) string,
	downloadAttachment func(interface{}, string),
	attachmentChanged func(interface{}, string) bool,
	isModified func(interface{}, *htmlgen.Document) bool,
	createDoc func(interface{}, *htmlgen.Document, *canvas.Canvas, *task.Task, int)) {

	register(name, func(t *task.Task, c *canvas.Canvas, db string, courseId int, finish func()) {
		list, err := apiGet(t.CreateProgress(1), c, courseId)
		if err != nil {
			panic(err)
		}
		fileWrites := t.CreateProgress(1)
		fileWrites.SetWork(len(list))
		for _, obj := range list {
			fileBaseName := path.Join(db, InvalidPathRunes.ReplaceAllLiteralString(getFilename(obj), ""))
			standaloneFile := fmt.Sprintf("%s.html", fileBaseName)
			attachedHTMLFile := path.Join(fileBaseName, "index.html")
			var outFile string
			var attachments []interface{}
			if attachments = getAttachments(obj); len(attachments) > 0 {
				if _, err := os.Stat(standaloneFile); err == nil {
					if err := os.Remove(standaloneFile); err != nil {
						panic(err)
					}
				}
				if err := os.Mkdir(fileBaseName, 0755); err != nil && !os.IsExist(err) {
					panic(err)
				}
				outFile = attachedHTMLFile
			} else {
				if _, err := os.Stat(fileBaseName); err == nil {
					if err := os.RemoveAll(fileBaseName); err != nil {
						panic(err)
					}
				}
				outFile = standaloneFile
			}
			content, err := ioutil.ReadFile(outFile)
			if err == nil {
				doc := htmlgen.ParseDocument(string(content), []htmlgen.ChildConstructor{docType})
				if doc != nil {
					if !isModified(obj, doc) {
						fileWrites.Finish(1)
						continue
					}
				}
			} else if !os.IsNotExist(err) {
				panic(err)
			}
			doc := htmlgen.CreateDocument()
			createDoc(obj, doc, c, t, courseId)
			if err := ioutil.WriteFile(outFile, []byte(doc.String()), 0644); err != nil {
				panic(err)
			}
			if len(attachments) > 0 {
				filenames := make([]string, len(attachments))
				for i, a := range attachments {
					filenames[i] = getAttachmentFilename(a)
				}
				dir, err := ioutil.ReadDir(fileBaseName)
				if err != nil {
					panic(err)
				}
				for _, f := range dir {
					if f.Name() != "index.html" {
						found := false
						for i, af := range filenames {
							if f.Name() == af {
								found = true
								if attachmentChanged(attachments[i], af) {
									downloadAttachment(attachments[i], af)
								}
							}
						}
						if !found {
							if err := os.RemoveAll(path.Join(fileBaseName, f.Name())); err != nil {
								panic(err)
							}
						}
					}
				}
				for i, af := range filenames {
					found := false
					for _, f := range dir {
						if f.Name() == af {
							found = true
						}
					}
					if !found {
						downloadAttachment(attachments[i], af)
					}
				}
			}
			fileWrites.Finish(1)
		}
		finish()
	})
}
