package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
)

type field struct {
	Name string
	Type string
	JSON string
}

type fields []field

func (f fields) Len() int {
	return len(f)
}

func (f fields) Less(i, j int) bool {
	return strings.Compare(f[i].Name, f[j].Name) < 0
}

func (f fields) Swap(i, j int) {
	tmp := f[i]
	f[i] = f[j]
	f[j] = tmp
}

type replacement struct {
	Regex       *regexp.Regexp
	Replacement string
}

var (
	replacements = []replacement{
		{
			Regex:       regexp.MustCompile("([a-z]|\\A)Id([A-Z]|\\z)"),
			Replacement: "${1}ID${2}",
		},
		{
			Regex:       regexp.MustCompile("([a-z]|\\A)Uuid([A-Z]|\\z)"),
			Replacement: "${1}UUID${2}",
		},
		{
			Regex:       regexp.MustCompile("([a-z]|\\A)Url([A-Z]|\\z)"),
			Replacement: "${1}URL${2}",
		},
		{
			Regex:       regexp.MustCompile("([a-z]|\\A)Html([A-Z]|\\z)"),
			Replacement: "${1}HTML${2}",
		},
	}
	commentRegex = regexp.MustCompile("(?m:^\\s*//.+)")
)

func main3() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <struct name>\n", os.Args[0])
		os.Exit(1)
	}
	filename := path.Join("..", "canvas", "model", fmt.Sprintf("%s.go", os.Args[1]))
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: file %s already exists.\n", filename)
		os.Exit(1)
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	buf = commentRegex.ReplaceAllLiteral(buf, []byte{})
	var data map[string]interface{}
	if err = json.Unmarshal(buf, &data); err != nil {
		panic(err)
	}
	fields := fields{}
	maxName := 0
	maxType := len("interface{}")
	for k, v := range data {
		t := ""
		if v == nil {
			t = "interface{}"
		} else {
			switch val := v.(type) {
			case float64:
				switch {
				case math.Floor(val) == val:
					t = "int"
					break
				default:
					t = "float64"
					break
				}
				break
			case string:
				t = "string"
				break
			case bool:
				t = "bool"
				break
			default:
				fmt.Fprintf(os.Stderr, "Error: unknown type %T\n", v)
				break
			}
		}
		name := ""
		capital := true
		for _, c := range k {
			switch {
			case c == '_':
				capital = true
				break
			case c >= 'a' && c <= 'z':
				if capital {
					name += string(c + 'A' - 'a')
					capital = false
				} else {
					name += string(c)
				}
				break
			case c >= 'A' && c <= 'Z':
				name += string(c)
				capital = false
				break
			default:
				fmt.Fprintf(os.Stderr, "Error: unknown rune %c\n", c)
				break
			}
		}
		for _, r := range replacements {
			name = r.Regex.ReplaceAllString(name, r.Replacement)
		}
		fields = append(fields, field{
			Name: name,
			Type: t,
			JSON: k,
		})
		if len(name) > maxName {
			maxName = len(name)
		}
		if len(t) > maxType {
			maxType = len(t)
		}
	}
	sort.Sort(fields)
	fmt.Fprintf(f, "package model\n\n// %s object\ntype %s struct {\n", os.Args[1], os.Args[1])
	format := fmt.Sprintf("    %%-%ds %%-%ds `json:\"%%s\"`\n", maxName, maxType)
	formatTodo := fmt.Sprintf("    %%-%ds %%-%ds `json:\"%%s\"` // TODO\n", maxName, maxType)
	for _, field := range fields {
		if field.Type == "" {
			fmt.Fprintf(f, formatTodo, field.Name, "interface{}", field.JSON)
		} else {
			fmt.Fprintf(f, format, field.Name, field.Type, field.JSON)
		}
	}
	fmt.Fprintln(f, "}")
}
