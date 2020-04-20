package overridegen

import (
	"fmt"
	"strings"
)

func equals(a, b, typeName string) string {
	if (typeName[0] >= 'A' && typeName[0] <= 'Z') || strings.ContainsRune(typeName, '.') {
		return fmt.Sprintf("reflect.DeepEqual(%s, %s)", a, b)
	}
	return fmt.Sprintf("%s == %s", a, b)
}

func (l *sliceLookup) writeCode(s *slice, ctx string, idx int) string {
	props := make([]string, len(l.TargetContext.Slices)+len(l.TargetContext.Properties))
	for i, slice := range l.TargetContext.Slices {
		props[i] = fmt.Sprintf("\t\t\t\ts%d: &(*o.s%d)[i].%s,", i, idx, slice.QualifiedName)
	}
	for i, prop := range l.TargetContext.Properties {
		props[i+len(l.TargetContext.Slices)] = fmt.Sprintf("\t\t\t\tp%d: &(*o.s%d)[i].%s,", i, idx, prop.QualifiedName)
	}
	return fmt.Sprintf(`func (o *%s) %s(search %s) *%s {
	for i, obj := range *o.s%d {
		if %s {
			return &%s {
				parent: o,
%s
			}
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to find key '%%v'", search))
}`, ctx, l.FuncName, l.AttributeType, l.TargetContext.ContextStructName, idx, equals(fmt.Sprintf("obj.%s", l.AttributeName), "search", l.AttributeType), l.TargetContext.ContextStructName, strings.Join(props, "\n"))
}

func (p *primitive) writeCode(ctx string, idx int) string {
	return fmt.Sprintf(`func (o *%s) set%s(old, new %s) *%s {
	if !(%s) {
		panic(fmt.Errorf("Error applying override patch: expected '%%v' but got '%%v'", old, *o.p%d))
	}
	*o.p%d = new
	return o
}`, ctx, p.IdentifierName, p.TypeName, ctx, equals(fmt.Sprintf("*o.p%d", idx), "old", p.TypeName), idx, idx)
}

func (s *slice) writeCode(ctx string, idx int) []string {
	parts := []string{
		fmt.Sprintf(`func (o *%s) add%s(obj %s) *%s {
	for _, e := range *o.s%d {
		if %s {
			panic(fmt.Errorf("Error applying override patch: unable to add '%%v' to list due to it already existing", obj))
		}
	}
	*o.s%d = append(*o.s%d, obj)
	return o
}`, ctx, s.IdentifierName, s.ElementType, ctx, idx, equals("e", "obj", s.ElementType), idx, idx),
		fmt.Sprintf(`func (o *%s) remove%s(obj %s) *%s {
	for i, e := range *o.s%d {
		if %s {
			*o.s%d = append((*o.s%d)[:i], (*o.s%d)[i+1:]...)
			return o
		}
	}
	panic(fmt.Errorf("Error applying override patch: unable to remove '%%v' from list due to it not existing", obj))
}`, ctx, s.IdentifierName, s.ElementType, ctx, idx, equals("e", "obj", s.ElementType), idx, idx, idx),
	}
	for _, l := range s.LookupFuncs {
		parts = append(parts, l.writeCode(s, ctx, idx))
	}
	if s.SubContext != nil {
		parts = append(parts, s.SubContext.writeSubcontextCode()...)
	}
	return parts
}

func (c *context) writeSubcontextCode() []string {
	parentCtxName := "interface{}"
	if c.Parent != nil {
		parentCtxName = c.Parent.ContextStructName
	}
	fields := make([]string, len(c.Slices)+len(c.Properties))
	for i, slice := range c.Slices {
		fields[i] = fmt.Sprintf("\ts%-5d %s", i, ptrType(slice.FieldType))
	}
	for i, prop := range c.Properties {
		fields[i+len(c.Slices)] = fmt.Sprintf("\tp%-5d *%s", i, prop.TypeName)
	}
	parts := []string{
		fmt.Sprintf(`type %s struct {
	parent *%s
%s
}`, c.ContextStructName, parentCtxName, strings.Join(fields, "\n")),
	}
	if c.Parent == nil {
		args := make([]string, len(fields))
		assignments := make([]string, len(args))
		for i, slice := range c.Slices {
			args[i] = fmt.Sprintf("%s %s", slice.QualifiedName, ptrType(slice.FieldType))
			assignments[i] = fmt.Sprintf("\t\ts%d: %s,", i, slice.QualifiedName)
		}
		for i, prop := range c.Properties {
			args[i+len(c.Slices)] = fmt.Sprintf("%s %s", prop.QualifiedName, ptrType(prop.TypeName))
			assignments[i+len(c.Slices)] = fmt.Sprintf("\t\tp%d: %s,", i, prop.QualifiedName)
		}
		parts = append(parts, fmt.Sprintf(`func createContext(%s) *%s {
	return &%s {
%s
	}
}`, strings.Join(args, ", "), c.ContextStructName, c.ContextStructName, strings.Join(assignments, "\n")))
	} else {
		parts = append(parts, fmt.Sprintf(`func (o *%s) done() *%s {
	return o.parent
}`, c.ContextStructName, c.Parent.ContextStructName))
	}
	for i, slice := range c.Slices {
		parts = append(parts, slice.writeCode(c.ContextStructName, i)...)
	}
	for i, prop := range c.Properties {
		parts = append(parts, prop.writeCode(c.ContextStructName, i))
	}
	return parts
}

func (c *context) writeCode() string {
	return fmt.Sprintf(`package overrides

import (
	"fmt"
	"reflect"

	"github.com/zachdeibert/canvas-sync/utils/apisync"
)

%s
`, strings.Join(c.writeSubcontextCode(), "\n\n"))
}
