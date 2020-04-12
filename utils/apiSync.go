package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/zachdeibert/canvas-sync/utils/apisync"
)

const (
	noImports = `package canvas

%s
`
	oneImport = `package canvas

import "%s"

%s
`
	multiImport = `package canvas

import (
%s
)

%s
`
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <Canvas source dir>\n", os.Args[0])
		os.Exit(1)
	}
	dir := path.Join(os.Args[1], "app", "controllers")
	outFile := path.Join("..", "canvas", "api.go")
	backup := fmt.Sprintf("%s.bk", outFile)
	if _, err := os.Stat(backup); !os.IsNotExist(err) {
		if err = os.Remove(backup); err != nil {
			panic(err)
		}
	}
	os.Rename(outFile, backup)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	models := []*apisync.Model{}
	methods := []apisync.MethodAPIPair{}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".rb") {
			inputFile := path.Join(dir, file.Name())
			fmt.Printf("Parsing file %s...\n", inputFile)
			api, err := apisync.ParseFile(inputFile)
			if err != nil {
				panic(err)
			}
			if api != nil {
				for _, s := range api.ModelComments {
					m, err := apisync.ParseModel(s)
					if err != nil {
						panic(err)
					}
					models = append(models, m)
				}
				for _, s := range api.MethodComments {
					m, err := apisync.ParseMethod(s)
					if err != nil {
						panic(err)
					}
					methods = append(methods, apisync.MethodAPIPair{
						Method:  m,
						APIName: api.Topic,
					})
				}
			}
		}
	}
	apisync.DedupModels(&models)
	apisync.DedupMethods(&methods)
	apisync.FixMissingTypes(models, methods)
	imports := []string{}
	parts := []string{}
	for _, m := range models {
		part, err := m.Write(&imports)
		if err != nil {
			panic(err)
		}
		parts = append(parts, part)
	}
	for _, m := range methods {
		part, err := m.Method.Write(m.APIName, &imports)
		if err != nil {
			panic(err)
		}
		parts = append(parts, part)
	}
	txt := ""
	switch len(imports) {
	case 0:
		txt = fmt.Sprintf(noImports, strings.Join(parts, "\n\n"))
		break
	case 1:
		txt = fmt.Sprintf(oneImport, imports[0], strings.Join(parts, "\n\n"))
		break
	default:
		stdImports := []string{}
		nonstdImports := []string{}
		for _, im := range imports {
			if strings.ContainsRune(im, '/') {
				nonstdImports = append(nonstdImports, fmt.Sprintf(`    "%s"`, im))
			} else {
				stdImports = append(stdImports, fmt.Sprintf(`    "%s"`, im))
			}
		}
		sort.Strings(stdImports)
		sort.Strings(nonstdImports)
		if len(stdImports) == 0 {
			imports = nonstdImports
		} else if len(nonstdImports) == 0 {
			imports = stdImports
		} else {
			imports = make([]string, len(stdImports)+len(nonstdImports)+1)
			copy(imports, stdImports)
			copy(imports[len(stdImports)+1:], nonstdImports)
		}
		txt = fmt.Sprintf(multiImport, strings.Join(imports, "\n"), strings.Join(parts, "\n\n"))
		break
	}
	if err = ioutil.WriteFile(outFile, []byte(txt), 0644); err != nil {
		panic(err)
	}
}
