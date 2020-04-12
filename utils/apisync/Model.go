package apisync

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type modelPropertyValuesJSON struct {
	Values []string `json:"values"`
}

type modelPropertyJSON struct {
	Descritpion     string                  `json:"description"`
	Example         interface{}             `json:"example"`
	Type            string                  `json:"type"`
	AllowableValues modelPropertyValuesJSON `json:"allowableValues"`
	Ref             string                  `json:"$ref"`
	Items           *modelPropertyJSON      `json:"items"`
	Key             *modelPropertyJSON      `json:"key"`
	Value           *modelPropertyJSON      `json:"value"`
}

type modelJSON struct {
	ID          string                       `json:"id"`
	Description string                       `json:"description"`
	Properties  map[string]modelPropertyJSON `json:"properties"`
}

// ModelProperty is a property of a model
type ModelProperty struct {
	Name        string
	Description string
	Example     string
	Type        string
	EnumValues  []string
}

// Model of an object
type Model struct {
	Name        string
	Description string
	Properties  []ModelProperty
}

var (
	modelCommentRegex = regexp.MustCompile("(?ms:@model\\s+[^\n]+\n(.*))")
)

func modelGetType(prop modelPropertyJSON) string {
	if prop.Ref != "" {
		return prop.Ref
	}
	switch prop.Type {
	case "string":
		return "string"
	case "integer":
		return "int"
	case "datetime":
		return "time.Time"
	case "boolean":
		return "bool"
	case "number":
		return "float64"
	case "array":
		if prop.Items != nil {
			return fmt.Sprintf("[]%s", modelGetType(*prop.Items))
		}
		return "[]interface{}"
	case "object":
		k := "interface{}"
		v := "interface{}"
		if prop.Key != nil {
			k = modelGetType(*prop.Key)
		}
		if prop.Value != nil {
			v = modelGetType(*prop.Value)
		}
		return fmt.Sprintf("map[%s]%s", k, v)
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown type '%s'\n", prop.Type)
		return "interface{}"
	}
}

// ParseModel parses a Model from a string
func ParseModel(str string) (*Model, error) {
	match := modelCommentRegex.FindStringSubmatch(str)
	if match == nil {
		return nil, errors.New("Unable to match model comment regex")
	}
	var data modelJSON
	if err := json.Unmarshal([]byte(match[1]), &data); err != nil {
		return nil, err
	}
	m := &Model{
		Name:        data.ID,
		Description: data.Description,
		Properties:  []ModelProperty{},
	}
	for k, v := range data.Properties {
		m.Properties = append(m.Properties, ModelProperty{
			Name:        k,
			Description: v.Descritpion,
			Example:     fmt.Sprint(v.Example),
			Type:        modelGetType(v),
			EnumValues:  v.AllowableValues.Values,
		})
	}
	sort.Sort(m)
	return m, nil
}

func (m *Model) Len() int {
	return len(m.Properties)
}

func (m *Model) Less(i, j int) bool {
	return strings.Compare(m.Properties[i].Name, m.Properties[j].Name) < 0
}

func (m *Model) Swap(i, j int) {
	tmp := m.Properties[i]
	m.Properties[i] = m.Properties[j]
	m.Properties[j] = tmp
}
