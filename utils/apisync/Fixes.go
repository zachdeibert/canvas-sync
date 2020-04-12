package apisync

import (
	"strings"
)

func fixCheckType(models []*Model, typename *string) {
	t := *typename
	array := false
	if strings.HasPrefix(t, "[]") {
		t = t[2:]
		array = true
	}
	if t[0] >= 'A' && t[0] <= 'Z' {
		found := false
		for _, model := range models {
			if model.Name == t {
				found = true
				break
			}
		}
		if !found {
			if array {
				*typename = "[]map[string]interface{}"
			} else {
				*typename = "map[string]interface{}"
			}
		}
	}
}

// FixMissingTypes replaces any missing types with generic ones
func FixMissingTypes(models []*Model, methods []MethodAPIPair) {
	for i, method := range methods {
		for j := range method.Method.Arguments {
			fixCheckType(models, &methods[i].Method.Arguments[j].Name)
		}
		fixCheckType(models, &methods[i].Method.ReturnType)
	}
}
