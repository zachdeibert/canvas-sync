package apisync

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// MethodArgument for a method
type MethodArgument struct {
	Name        string
	Required    bool
	Type        string
	Description string
	EnumValues  []string
}

// Method for the API
type Method struct {
	Name        string
	Description string
	Arguments   []MethodArgument
	ReturnType  string
	EndPoint    string
}

var (
	methodTopRegex      = regexp.MustCompile("(?m:\\s@([^\\s]+) *([^\n]*)((?:\n\\s*[^@\\s][^\n]+)*))")
	methodArgumentRegex = regexp.MustCompile("^([^ ]+)(?: \\[([^\\]]+)\\]$)?")
	methodEndPointRegex = regexp.MustCompile("https://<canvas>/api/v1/([^\\s?]+)")
)

// ParseMethod parses a method from a string
func ParseMethod(str string) (*Method, error) {
	matches := methodTopRegex.FindAllStringSubmatch(fmt.Sprintf(" %s", str), -1)
	m := &Method{
		Name:        "",
		Description: "",
		Arguments:   []MethodArgument{},
		ReturnType:  "interface{}",
		EndPoint:    "",
	}
	for _, match := range matches {
		switch match[1] {
		case "API":
			m.Name = strings.TrimSpace(match[2])
			m.Description = strings.TrimSpace(apiFileLineTrim.ReplaceAllLiteralString(match[3], " "))
			break
		case "beta":
			if len(m.Description) == 0 {
				m.Description = strings.TrimSpace(apiFileLineTrim.ReplaceAllLiteralString(match[3], " "))
			}
			break
		case "argument":
			arg := methodArgumentRegex.FindStringSubmatch(match[2])
			if arg == nil {
				fmt.Println(match[2])
				return nil, errors.New("Unable to match argument regex")
			}
			found := false
			name := arg[1]
			if idx := strings.IndexRune(name, '['); idx >= 0 {
				name = name[0:idx]
			}
			for _, a := range m.Arguments {
				if a.Name == name {
					found = true
					break
				}
			}
			if found {
				break
			}
			a := MethodArgument{
				Name:        name,
				Required:    false,
				Type:        "interface{}",
				Description: strings.TrimSpace(apiFileLineTrim.ReplaceAllLiteralString(match[3], " ")),
				EnumValues:  nil,
			}
			array := false
			for _, attr := range strings.Split(arg[2], ",") {
				switch v := strings.TrimSpace(attr); strings.ToLower(v) {
				case "required":
					a.Required = true
					break
				case "optional":
					a.Required = false
					break
				case "date", "datetime":
					a.Type = "time.Time"
					break
				case "boolean":
					a.Type = "bool"
					break
				case "string", "hash", "url":
					a.Type = "string"
					break
				case "integer":
					a.Type = "int"
					break
				case "number":
					a.Type = "float64"
					break
				case "array":
					array = true
					break
				case "json":
					a.Type = "map[string]interface{}"
					break
				case "", "deprecated":
					break
				default:
					if strings.HasPrefix(v, "default") {
						break
					}
					if v[0] >= 'A' && v[0] <= 'Z' {
						a.Type = v
						break
					}
					if strings.ContainsRune(v, '|') || strings.ContainsRune(v, '"') {
						vals := strings.Split(v, "|")
						a.EnumValues = make([]string, len(vals))
						for i, val := range vals {
							if strings.HasPrefix(val, "\"") && strings.HasSuffix(val, "\"") {
								a.EnumValues[i] = val[1 : len(val)-1]
							} else {
								a.EnumValues[i] = val
							}
						}
						break
					}
					fmt.Fprintf(os.Stderr, "Error: unknown argument attribute '%s'\n", v)
					break
				}
			}
			if array {
				a.Type = fmt.Sprintf("[]%s", a.Type)
			}
			m.Arguments = append(m.Arguments, a)
			break
		case "example_request":
			endpoint := methodEndPointRegex.FindStringSubmatch(match[3])
			if endpoint != nil {
				m.EndPoint = endpoint[1]
			}
			break
		case "example_response", "response_field", "subtopic":
			break
		case "returns":
			t := strings.TrimSpace(match[2])
			switch {
			case strings.ContainsRune(t, '{'), strings.ContainsRune(t, ' '):
				break
			case strings.HasPrefix(t, "[") && strings.HasSuffix(t, "]"):
				m.ReturnType = fmt.Sprintf("[]%s", t[1:len(t)-1])
				break
			case t == "boolean":
				m.ReturnType = "bool"
				break
			default:
				m.ReturnType = t
				break
			}
			break
		default:
			fmt.Fprintf(os.Stderr, "Error: unknown tag '@%s'\n", match[1])
			break
		}
	}
	return m, nil
}
