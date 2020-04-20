package overridegen

import (
	"fmt"
	"reflect"
)

type primitive struct {
	IdentifierName string
	QualifiedName  string
	TypeName       string
}

func (p *primitive) String() string {
	return tabsToSpaces(fmt.Sprintf(`&primitive (%p) {
	IdentifierName: "%s",
	QualifiedName: "%s",
	TypeName: "%s",
}`, p, p.IdentifierName, p.QualifiedName, p.TypeName))
}

func createPrimitive(qualifiedName string, t reflect.Type) *primitive {
	return &primitive{
		IdentifierName: createIdentifierName(qualifiedName),
		QualifiedName:  qualifiedName,
		TypeName:       typeName(t),
	}
}
