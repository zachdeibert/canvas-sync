package overridegen

import (
	"fmt"
	"reflect"
)

type sliceLookup struct {
	FuncName      string
	AttributeName string
	AttributeType string
	TargetContext *context
}

func (l *sliceLookup) String() string {
	return tabsToSpaces(fmt.Sprintf(`&sliceLookup (%p) {
	FuncName: "%s",
	AttributeName: "%s",
	AttributeType: "%s",
	TargetContext: %p,
}`, l, l.FuncName, l.AttributeName, l.AttributeType, l.TargetContext))
}

func createSliceLookup(funcName, attrName string, t reflect.Type, target *context) *sliceLookup {
	return &sliceLookup{
		FuncName:      funcName,
		AttributeName: attrName,
		AttributeType: typeName(t),
		TargetContext: target,
	}
}
