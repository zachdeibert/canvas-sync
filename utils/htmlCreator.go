package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main2() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <struct name> [model name]\n", os.Args[0])
		os.Exit(1)
	}
	filename := path.Join("..", "canvassync", "coursetasks", "html", fmt.Sprintf("%s.go", os.Args[1]))
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: file %s already exists.\n", filename)
		os.Exit(1)
	}
	uppercase := strings.ToUpper(os.Args[1][0:1]) + os.Args[1][1:]
	lowercase := strings.ToLower(os.Args[1][0:1]) + os.Args[1][1:]
	var model string
	if len(os.Args) == 3 {
		model = os.Args[2]
	} else {
		model = uppercase
	}
	ioutil.WriteFile(filename, []byte(fmt.Sprintf(`package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	%sTemplate *%s
	// %sChildCtors for parsing a template
	%sChildCtors = []htmlgen.ChildConstructor{
		// TODO
	}
)

// %s HTML template
type %s struct {
	Data   canvas.%s
	format *htmlgen.FormatSection
}

// Create%s creates a new template
func Create%s() *%s {
	obj := &%s{}
	args := []interface{}{
		// TODO
	}
	if %sTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(%c
<!-- TODO -->
%c, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = %sTemplate.format.Clone(args)
	}
	return obj
}

// Parse%s parses a string to a template
func Parse%s(str string) *%s {
	o := Create%s()
	if _, ok := o.Parse(str, %sChildCtors); ok {
		return o
	}
	return nil
}

func init() {
	%sTemplate = Create%s()
}

// AppendChild adds a child to the section
func (t *%s) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *%s) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *%s) String() string {
	return t.format.String()
}

// Parse the template
func (t *%s) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
`, lowercase, uppercase, uppercase, uppercase, uppercase, uppercase, model, uppercase, uppercase, uppercase, uppercase,
		lowercase, '`', '`', lowercase, uppercase, uppercase, uppercase, uppercase, uppercase, lowercase, uppercase,
		uppercase, uppercase, uppercase, uppercase)), 0644)
}
