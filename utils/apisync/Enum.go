package apisync

import (
	"fmt"
	"strings"
)

const (
	enumFormat = `// %s enumeration
type %s string

const (
%s
)`
	enumValFormat = `    // %s enum value ("%s")
    %s %s = "%s"`
)

func defineEnum(values []string, typeName *string, propName string) string {
	*typeName = toGoIdentifier(propName, true)
	vals := make([]string, len(values))
	for i, val := range values {
		name := toGoIdentifier(fmt.Sprintf("%s_%s", propName, val), true)
		if val == "" {
			name = fmt.Sprintf("%sNone", name)
		}
		vals[i] = fmt.Sprintf(enumValFormat, name, val, name, *typeName, val)
	}
	return fmt.Sprintf(enumFormat, *typeName, *typeName, strings.Join(vals, "\n"))
}

// DefineEnums creates enums for the string parameters that have specific allowed values
func DefineEnums(models []*Model, methods []MethodAPIPair) []string {
	enums := []string{}
	for i, model := range models {
		for j, prop := range model.Properties {
			if len(prop.EnumValues) > 0 {
				enums = append(enums, defineEnum(prop.EnumValues, &models[i].Properties[j].Type, fmt.Sprintf("%s_%s", model.Name, prop.Name)))
			}
		}
	}
	for i, method := range methods {
		for j, arg := range method.Method.Arguments {
			if len(arg.EnumValues) > 0 {
				enums = append(enums, defineEnum(arg.EnumValues, &methods[i].Method.Arguments[j].Type, fmt.Sprintf("%s_%s_%s", method.APIName, method.Method.Name, arg.Name)))
			}
		}
	}
	return enums
}
