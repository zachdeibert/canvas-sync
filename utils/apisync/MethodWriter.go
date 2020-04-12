package apisync

import (
	"fmt"
	"strings"
)

const (
	scalarMethodFormat = `%s
func (c *Canvas) %s(progress *task.Progress%s) (*%s, error) {
	endpoint := fmt.Sprintf("%s"%s)
	params := map[string]interface{}{}
%s
	responseCtor := %s
	var res *%s
	callback := func(obj interface{}) error {
		res = obj.(*%s)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}`
	vectorMethodFormat = `%s
func (c *Canvas) %s(progress *task.Progress%s) ([]%s, error) {
	endpoint := fmt.Sprintf("%s"%s)
	params := map[string]interface{}{}
	%s
	responseCtor := %s
	var res []%s
	callback := func(obj interface{}) error {
		arr := *obj.(*[]%s)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}`
	compositeResponseCtor = `func() interface{} {
		return &%s{}
	}`
	vectorResponseCtor = `func() interface{} {
		return &[]%s{}
	}`
	primitiveResponseCtor = `func() interface{} {
		var tmp %s
		return &tmp
	}`
)

func (m *Method) Write(apiName string, imports *[]string) (string, error) {
	addImport(imports, "fmt")
	addImport(imports, "github.com/zachdeibert/canvas-sync/task")
	name := toGoIdentifier(fmt.Sprintf("%s_%s", apiName, m.Name), true)
	comment := descComment(name, "API call", m.Description, 0, 120)
	paramArgs := "" // TODO
	resType := m.ReturnType
	format := scalarMethodFormat
	responseCtor := compositeResponseCtor
	if resType == "interface{}" {
		resType = "map[string]interface{}"
	} else if strings.HasPrefix(resType, "[]") {
		resType = resType[2:]
		format = vectorMethodFormat
		responseCtor = vectorResponseCtor
	} else if resType[0] >= 'a' && resType[0] <= 'z' {
		responseCtor = primitiveResponseCtor
	}
	endpointFormat := ""     // TODO
	endpointFormatArgs := "" // TODO
	paramsCode := ""         // TODO
	return fmt.Sprintf(format, comment, name, paramArgs, resType, endpointFormat, endpointFormatArgs, paramsCode, fmt.Sprintf(responseCtor, resType), resType, resType), nil
}
