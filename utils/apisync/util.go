package apisync

import (
	"fmt"
	"regexp"
	"strings"
)

var nonGoIdentifier = regexp.MustCompile("[^a-zA-Z0-9]+")

func splitText(str string, width int) []string {
	words := strings.Split(str, " ")
	lines := []string{}
	lastLine := []string{}
	lastLineLen := 0
	newline := false
	for _, word := range words {
		if newline {
			newline = false
			lines = append(lines, strings.Join(lastLine, " "))
			lastLine = []string{word[1:]}
			lastLineLen = len(word)
			continue
		}
		if strings.HasSuffix(word, "\n") {
			newline = true
			word = word[0 : len(word)-1]
		}
		if strings.HasPrefix(word, "\n") {
			lines = append(lines, strings.Join(lastLine, " "))
			lastLine = []string{word[1:]}
			lastLineLen = len(word)
		} else if lastLineLen+len(word) > width {
			lines = append(lines, strings.Join(lastLine, " "))
			lastLine = []string{word}
			lastLineLen = len(word) + 1
		} else {
			lastLine = append(lastLine, word)
			lastLineLen += len(word) + 1
		}
	}
	lines = append(lines, strings.Join(lastLine, " "))
	return lines
}

func makeComment(str string, indent int, width int) string {
	a := splitText(str, width-indent-3)
	prefixB := make([]byte, indent+3)
	for i := 0; i < indent; i++ {
		prefixB[i] = ' '
	}
	prefixB[indent] = '/'
	prefixB[indent+1] = '/'
	prefixB[indent+2] = ' '
	prefix := string(prefixB)
	for i, v := range a {
		a[i] = fmt.Sprintf("%s%s", prefix, v)
	}
	return strings.Join(a, "\n")
}

func importString(pkgs []string) string {
	switch len(pkgs) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("\nimport \"%s\"\n", pkgs[0])
	default:
		imports := make([]string, len(pkgs))
		for i, pkg := range pkgs {
			imports[i] = fmt.Sprintf("    \"%s\"", pkg)
		}
		return fmt.Sprintf("\nimport (\n%s\n)\n", strings.Join(imports, "\n"))
	}
}

func toGoIdentifier(jsonIdentifier string, exported bool) string {
	words := strings.Split(strings.ToLower(nonGoIdentifier.ReplaceAllLiteralString(jsonIdentifier, "_")), "_")
	for i, word := range words {
		switch word {
		case "api":
			word = "API"
			break
		case "guid":
			word = "GUID"
			break
		case "html":
			word = "HTML"
			break
		case "http":
			word = "HTTP"
			break
		case "id":
			word = "ID"
			break
		case "ip":
			word = "IP"
			break
		case "tls":
			word = "TLS"
			break
		case "url":
			word = "URL"
			break
		case "uuid":
			word = "UUID"
			break
		case "":
			break
		default:
			word = fmt.Sprintf("%c%s", word[0]+'A'-'a', word[1:])
			break
		}
		words[i] = word
	}
	if !exported {
		words[0] = strings.ToLower(words[0])
	}
	str := strings.Join(words, "")
	switch str {
	case "type":
		return "typeName"
	case "select":
		return "selectField"
	case "error":
		return "err"
	default:
		return str
	}
}

func addImport(imports *[]string, pkg string) {
	for _, im := range *imports {
		if im == pkg {
			return
		}
	}
	*imports = append(*imports, pkg)
}

func descComment(name, t, desc string, indent, width int) string {
	str := fmt.Sprintf("%s %s", name, t)
	if len(desc) > 0 {
		str = fmt.Sprintf("%s: %s", str, desc)
	}
	return makeComment(str, indent, width)
}
