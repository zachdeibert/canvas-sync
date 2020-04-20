package overridegen

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var (
	illegalIdentiferChars = regexp.MustCompile("[.*\\[\\]]+")
)

func indent(str string, n int, startLine int) string {
	indentC := make([]rune, n)
	for i := range indentC {
		indentC[i] = ' '
	}
	indent := string(indentC)
	lines := strings.Split(str, "\n")
	for i, l := range lines {
		if i >= startLine {
			lines[i] = fmt.Sprintf("%s%s", indent, l)
		}
	}
	return strings.Join(lines, "\n")
}

func tabsToSpaces(str string) string {
	return strings.ReplaceAll(str, "\t", "    ")
}

func createIdentifierName(qualifiedName string) string {
	str := illegalIdentiferChars.ReplaceAllLiteralString(qualifiedName, "")
	if str[0] >= 'a' && str[0] <= 'z' {
		str = fmt.Sprintf("%c%s", (str[0] - 'a' + 'A'), str[1:])
	}
	return str
}

func typeName(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Ptr:
		return fmt.Sprintf("*%s", typeName(t.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", typeName(t.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", typeName(t.Key()), typeName(t.Elem()))
	default:
		name := t.Name()
		if name[0] >= 'A' && name[0] <= 'Z' {
			return fmt.Sprintf("apisync.%s", name)
		}
		return name
	}
}

func dereference(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func ptrType(typename string) string {
	if strings.HasPrefix(typename, "*") {
		return typename
	}
	return fmt.Sprintf("*%s", typename)
}
