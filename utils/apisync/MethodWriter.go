package apisync

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	scalarMethodFormat = `%s
func (c *Canvas) %s(progress *task.Progress%s) (*%s, error) {
	endpoint := fmt.Sprintf("%s"%s)
	params := map[string]interface{}{}%s
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
	params := map[string]interface{}{}%s
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
	paramCodeScalar = `
	if %s != nil {
		params["%s"] = *%s
	}`
	paramCodeVector = `
	if %s != nil && len(%s) > 0 {
		params["%s"] = %s
	}`
)

var (
	endpointArgRegex = regexp.MustCompile("<([^>]+)>")
)

func (m *Method) Write(apiName string, imports *[]string) (string, error) {
	addImport(imports, "fmt")
	addImport(imports, "github.com/zachdeibert/canvas-sync/task")
	name := ToGoIdentifier(fmt.Sprintf("%s_%s", apiName, m.Name), true)
	comment := descComment(name, "API call", m.Description, 0, 120)
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
	endpointMatches := endpointArgRegex.FindAllStringSubmatchIndex(m.EndPoint, -1)
	paramsCodes := make([]string, len(m.Arguments))
	paramsArgs := make([]string, len(paramsCodes)+len(endpointMatches)+1)
	for i, arg := range m.Arguments {
		name := ToGoIdentifier(arg.Name, false)
		if strings.HasPrefix(arg.Type, "[]") || strings.HasPrefix(arg.Type, "map[") {
			paramsArgs[i+1] = fmt.Sprintf("%s %s", name, arg.Type)
			paramsCodes[i] = fmt.Sprintf(paramCodeVector, name, name, arg.Name, name)
		} else {
			paramsArgs[i+1] = fmt.Sprintf("%s *%s", name, arg.Type)
			paramsCodes[i] = fmt.Sprintf(paramCodeScalar, name, arg.Name, name)
		}
	}
	start := 0
	endpointConstants := make([]string, len(endpointMatches)+1)
	endpointFormatArgs := make([]string, len(endpointConstants))
	for i, match := range endpointMatches {
		endpointConstants[i] = m.EndPoint[start:match[0]]
		start = match[1]
		name := ToGoIdentifier(m.EndPoint[match[2]:match[3]], false)
		paramsArgs[len(m.Arguments)+1+i] = fmt.Sprintf("%s string", name)
		endpointFormatArgs[i+1] = name
	}
	endpointConstants[len(endpointMatches)] = m.EndPoint[start:]
	paramArgsStr := strings.Join(paramsArgs, ", ")
	paramsCode := strings.Join(paramsCodes, "")
	endpointFormat := strings.Join(endpointConstants, "%s")
	endpointFormatArgsStr := strings.Join(endpointFormatArgs, ", ")
	return fmt.Sprintf(format, comment, name, paramArgsStr, resType, endpointFormat, endpointFormatArgsStr, paramsCode, fmt.Sprintf(responseCtor, resType), resType, resType), nil
}
