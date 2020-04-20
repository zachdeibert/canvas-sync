package overridegen

import (
	"fmt"
	"reflect"
	"strings"
)

type slice struct {
	IdentifierName string
	QualifiedName  string
	FieldType      string
	ElementType    string
	LookupFuncs    []*sliceLookup
	SubContext     *context
}

func (s *slice) String() string {
	funcs := make([]string, len(s.LookupFuncs))
	for i, f := range s.LookupFuncs {
		funcs[i] = indent(f.String(), 8, 0)
	}
	funcStr := strings.Join(funcs, ",\n")
	if len(funcStr) > 0 {
		funcStr = fmt.Sprintf("\n%s,\n\t", funcStr)
	}
	subCtx := "nil"
	if s.SubContext != nil {
		subCtx = indent(s.SubContext.String(), 4, 1)
	}
	return tabsToSpaces(fmt.Sprintf(`&slice (%p) {
	IdentifierName: "%s",
	QualifiedName: "%s",
	FieldType: "%s",
	ElementType: "%s",
	LookupFuncs: []*sliceLookup{%s},
	SubContext: %s,
}`, s, s.IdentifierName, s.QualifiedName, s.FieldType, s.ElementType, funcStr, subCtx))
}

func createSlice(qualifiedName string, t reflect.Type) *slice {
	return &slice{
		IdentifierName: createIdentifierName(qualifiedName),
		QualifiedName:  qualifiedName,
		FieldType:      typeName(t),
		ElementType:    typeName(dereference(t).Elem()),
		LookupFuncs:    []*sliceLookup{},
		SubContext:     nil,
	}
}

func (s *slice) addLookupFunc(funcName, attrName string, t reflect.Type, target *context) {
	s.LookupFuncs = append(s.LookupFuncs, createSliceLookup(funcName, attrName, t, target))
}
