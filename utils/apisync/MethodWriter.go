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
)

var (
	endpointArgRegex = regexp.MustCompile("<([^>]+)>")
)

func (m *Method) Write(apiName string, imports *[]string) (string, error) {
	addImport(imports, "fmt")
	addImport(imports, "github.com/zachdeibert/canvas-sync/task")
	name := toGoIdentifier(fmt.Sprintf("%s_%s", apiName, m.Name), true)
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
	paramsCodes := make([]string, len(m.Arguments)+1)
	paramsArgs := make([]string, len(paramsCodes)+len(endpointMatches))
	for i, arg := range m.Arguments {
		name := toGoIdentifier(arg.Name, false)
		paramsArgs[i+1] = fmt.Sprintf("%s %s", name, arg.Type)
		paramsCodes[i+1] = fmt.Sprintf(`    params["%s"] = %s`, arg.Name, name)
	}
	start := 0
	endpointConstants := make([]string, len(endpointMatches)+1)
	endpointFormatArgs := make([]string, len(endpointConstants))
	for i, match := range endpointMatches {
		endpointConstants[i] = m.EndPoint[start:match[0]]
		start = match[1]
		name := toGoIdentifier(m.EndPoint[match[2]:match[3]], false)
		paramsArgs[len(m.Arguments)+1+i] = fmt.Sprintf("%s string", name)
		endpointFormatArgs[i+1] = name
	}
	endpointConstants[len(endpointMatches)] = m.EndPoint[start:]
	paramArgsStr := strings.Join(paramsArgs, ", ")
	paramsCode := strings.Join(paramsCodes, "\n")
	endpointFormat := strings.Join(endpointConstants, "%s")
	endpointFormatArgsStr := strings.Join(endpointFormatArgs, ", ")
	return fmt.Sprintf(format, comment, name, paramArgsStr, resType, endpointFormat, endpointFormatArgsStr, paramsCode, fmt.Sprintf(responseCtor, resType), resType, resType), nil
}
