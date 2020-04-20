package overridegen

import (
	"fmt"
	"reflect"
	"strings"
)

type context struct {
	ContextStructName string
	Parent            *context
	ParentSlice       *slice
	Slices            []*slice
	Properties        []*primitive
}

func (c *context) String() string {
	slices := make([]string, len(c.Slices))
	for i, s := range c.Slices {
		slices[i] = indent(s.String(), 8, 0)
	}
	sliceStr := strings.Join(slices, ",\n")
	if len(sliceStr) > 0 {
		sliceStr = fmt.Sprintf("\n%s,\n\t", sliceStr)
	}
	props := make([]string, len(c.Properties))
	for i, p := range c.Properties {
		props[i] = indent(p.String(), 8, 0)
	}
	propStr := strings.Join(props, ",\n")
	if len(propStr) > 0 {
		propStr = fmt.Sprintf("\n%s,\n\t", propStr)
	}
	return tabsToSpaces(fmt.Sprintf(`&context (%p) {
	ContextStructName: "%s",
	Parent: %p,
	ParentSlice: %p,
	Slices: []*slice{%s},
	Properties: []*primitive{%s},
}`, c, c.ContextStructName, c.Parent, c.ParentSlice, sliceStr, propStr))
}

func createContext() *context {
	c := &context{
		ContextStructName: "overrideContext",
		Parent:            nil,
		ParentSlice:       nil,
		Slices:            []*slice{},
		Properties:        []*primitive{},
	}
	return c
}

func (c *context) reflect(t reflect.Type, prefix string) {
	it := dereference(t)
	for i := it.NumField() - 1; i >= 0; i-- {
		field := it.Field(i)
		ft := dereference(field.Type)
		switch ft.Kind() {
		case reflect.Slice:
			c.addSlice(fmt.Sprintf("%s%s", prefix, field.Name), field.Type)
			break
		case reflect.Struct:
			if dereference(field.Type) != it {
				c.reflect(field.Type, fmt.Sprintf("%s%s.", prefix, field.Name))
				break
			}
			fallthrough
		default:
			if tag, ok := field.Tag.Lookup("overridegen"); ok {
				s := c
				for s.ParentSlice == nil && s.Parent != nil {
					s = s.Parent
				}
				if s.ParentSlice != nil {
					s.ParentSlice.addLookupFunc(tag, fmt.Sprintf("%s%s", prefix, field.Name), ft, s)
				}
			}
			c.addPrimitive(fmt.Sprintf("%s%s", prefix, field.Name), field.Type)
			break
		}
	}
}

func (c *context) addSlice(name string, t reflect.Type) {
	s := createSlice(name, t)
	sliceType := dereference(t)
	structType := dereference(sliceType.Elem())
	if structType.Kind() == reflect.Struct {
		s.SubContext = &context{
			ContextStructName: fmt.Sprintf("%s%s", c.ContextStructName, createIdentifierName(name)),
			Parent:            c,
			ParentSlice:       s,
			Slices:            []*slice{},
			Properties:        []*primitive{},
		}
		s.SubContext.reflect(structType, "")
	}
	c.Slices = append(c.Slices, s)
}

func (c *context) addPrimitive(name string, t reflect.Type) {
	c.Properties = append(c.Properties, createPrimitive(name, t))
}
