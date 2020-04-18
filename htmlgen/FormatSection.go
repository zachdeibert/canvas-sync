package htmlgen

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// FormatSection represents a section of HTML defined by a format string
type FormatSection struct {
	format   string
	args     []interface{}
	children []Section
}

var (
	formatSectionChild interface{}
	// FormatSectionChild is a formatting argument that represents the children nodes
	FormatSectionChild interface{} = &formatSectionChild
	errNotEnoughArgs               = errors.New("Not enough arguments for format string")
	errInvalidType                 = errors.New("Invalid argument type")
	errChildFormat                 = errors.New("Child format must be %s")
	errFormatString                = errors.New("Invalid format string")
	formatRegex                    = regexp.MustCompile("%[^a-zA-Z%]*[a-zA-Z]")
)

// CreateFormatSection creates a new FormatSection
func CreateFormatSection(format string, args []interface{}) (*FormatSection, error) {
	isFormat := false
	argI := 0
	for _, r := range format {
		if isFormat {
			isFormat = false
			if args[argI] == FormatSectionChild {
				if r == 's' {
					argI++
					continue
				}
				return nil, errChildFormat
			}
			switch r {
			case '%':
				break
			case 'd':
				if argI < len(args) {
					switch (args[argI]).(type) {
					case *int8, *uint8, *int16, *uint16, *int32, *uint32, *int64, *uint64, *int, *uint:
						argI++
						break
					default:
						return nil, errInvalidType
					}
				} else {
					return nil, errNotEnoughArgs
				}
				break
			case 'e', 'E', 'f', 'F', 'g', 'G':
				if argI < len(args) {
					switch (args[argI]).(type) {
					case *float32:
					case *float64:
						argI++
						break
					default:
						return nil, errInvalidType
					}
				} else {
					return nil, errNotEnoughArgs
				}
				break
			case 's':
				if argI < len(args) {
					switch (args[argI]).(type) {
					case *string, *fmt.Stringer, *[]byte, CustomFormat:
						argI++
						break
					default:
						return nil, errInvalidType
					}
				} else {
					return nil, errNotEnoughArgs
				}
				break
			case '#', '+', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
				isFormat = true
				break
			default:
				return nil, errFormatString
			}
		} else {
			isFormat = r == '%'
		}
	}
	return &FormatSection{
		format:   format,
		args:     args,
		children: []Section{},
	}, nil
}

// Clone creates a new copy of the FormatSection without doing all the checks that CreateFormatSection does
func (s *FormatSection) Clone(args []interface{}) *FormatSection {
	if len(args) == len(s.args) {
		for i, v := range args {
			if reflect.TypeOf(v) != reflect.TypeOf(s.args[i]) {
				return nil
			}
		}
		return &FormatSection{
			format:   s.format,
			args:     args,
			children: []Section{},
		}
	}
	return nil
}

// AppendChild adds a section to the format section
func (s *FormatSection) AppendChild(child Section) {
	s.children = append(s.children, child)
}

// Children gets all of the section's children
func (s *FormatSection) Children() []Section {
	return s.children
}

func (s *FormatSection) String() string {
	args := make([]interface{}, len(s.args))
	for i, p := range s.args {
		if p == FormatSectionChild {
			children := make([]string, len(s.children))
			for j, c := range s.children {
				children[j] = c.String()
			}
			args[i] = strings.Join(children, "")
		} else {
			switch v := s.args[i].(type) {
			case *int8:
				args[i] = *v
				break
			case *uint8:
				args[i] = *v
				break
			case *int16:
				args[i] = *v
				break
			case *uint16:
				args[i] = *v
				break
			case *int32:
				args[i] = *v
				break
			case *uint32:
				args[i] = *v
				break
			case *int64:
				args[i] = *v
				break
			case *uint64:
				args[i] = *v
				break
			case *int:
				args[i] = *v
				break
			case *uint:
				args[i] = *v
				break
			case *float32:
				args[i] = *v
				break
			case *float64:
				args[i] = *v
				break
			case *string:
				args[i] = *v
				break
			case *fmt.Stringer:
				args[i] = *v
				break
			case *[]byte:
				args[i] = *v
				break
			case CustomFormat:
				args[i] = v.FormatHTML()
				break
			}
		}
	}
	return fmt.Sprintf(s.format, args...)
}

// Parse a format section
func (s *FormatSection) Parse(str string, childCtors []ChildConstructor) (string, bool) {
	newChildren := make([]Section, len(childCtors))
	childChildCtors := make([][]ChildConstructor, len(childCtors))
	for i, ctor := range childCtors {
		newChildren[i], childChildCtors[i] = ctor()
	}
	formats := formatRegex.FindAllStringIndex(s.format, -1)
	start := 0
	strLeft := str
	for i, f := range formats {
		prefix := s.format[start:f[0]]
		start = f[1]
		if len(prefix) > len(strLeft) {
			return "", false
		}
		if !strings.HasPrefix(strLeft, prefix) {
			return "", false
		}
		strLeft = strLeft[len(prefix):]
		if s.args[i] == FormatSectionChild {
			for found := true; found; {
				found = false
				for j, child := range newChildren {
					left, ok := child.Parse(strLeft, childChildCtors[j])
					if ok {
						strLeft = left
						s.children = append(s.children, child)
						newChildren[j], _ = childCtors[j]()
						found = true
						break
					}
				}
			}
		} else if s.format[f[1]-1] == 's' {
			suffix := ""
			if i == len(formats)-1 {
				suffix = s.format[f[1]:]
			} else {
				suffix = s.format[f[1]:formats[i+1][0]]
			}
			fieldWidth := strings.Index(strLeft, suffix)
			*(s.args[i].(*string)) = strLeft[0:fieldWidth]
			strLeft = strLeft[fieldWidth:]
		} else {
			suffix := ""
			if n, err := fmt.Sscanf(strLeft, fmt.Sprintf("%s%%s", s.format[f[0]:f[1]]), s.args[i], &suffix); n != 2 || err != nil {
				s.children = []Section{}
				return "", false
			}
			fieldWidth := strings.Index(strLeft, suffix)
			strLeft = strLeft[fieldWidth:]
		}
	}
	suffix := s.format[start:]
	if len(suffix) > len(strLeft) {
		s.children = []Section{}
		return "", false
	}
	if !strings.HasPrefix(strLeft, suffix) {
		s.children = []Section{}
		return "", false
	}
	strLeft = strLeft[len(suffix):]
	return strLeft, true
}
