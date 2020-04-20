package apisync

import (
	"fmt"
	"strings"
)

const (
	modelFormat = `%s
type %s struct {
%s
}`
	modelPropFormat = `%s
    %s %s %cjson:"%s"%c`
)

func (m *Model) Write(imports *[]string) (string, error) {
	props := make([]string, len(m.Properties))
	for i, p := range m.Properties {
		pkg := strings.Split(p.Type, ".")
		if len(pkg) == 2 {
			typeName := pkg[0]
			if strings.HasPrefix(typeName, "[]") {
				typeName = typeName[2:]
			}
			addImport(imports, typeName)
		}
		t := p.Type
		if t[0] >= 'A' && t[0] <= 'Z' {
			t = fmt.Sprintf("*%s", t)
		}
		name := ToGoIdentifier(p.Name, true)
		props[i] = fmt.Sprintf(modelPropFormat, descComment(name, "field", p.Description, 4, 120), name, t, '`', p.Name, '`')
	}
	return fmt.Sprintf(modelFormat, descComment(m.Name, "model object", m.Description, 0, 120), m.Name, strings.Join(props, "\n")), nil
}
