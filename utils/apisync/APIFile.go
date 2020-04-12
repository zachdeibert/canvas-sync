package apisync

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
)

// APIFile contains documentation information about an API call
type APIFile struct {
	Topic          string
	Description    string
	ModelComments  []string
	MethodComments []string
}

var (
	apiFileTopRegex     = regexp.MustCompile("(?ms:^# @API ([^\n]+)\n((?:#[^\n]*\n)+)(.*))")
	apiFileMethodRegex  = regexp.MustCompile("(?m:^(\\s+# @API (?:[^\n]+)\n(?:\\s+#[^\n]*\n)+)\n)")
	apiFileCommentStart = regexp.MustCompile("(?m:^\\s*#)")
	apiFileModelStart   = regexp.MustCompile("(?m:^\\s*@model\\s+[^\\s]+\\s*?$)")
	apiFileLineTrim     = regexp.MustCompile("(?m:\\s*?\n\\s*)")
	apiFileMethodStart  = regexp.MustCompile("(?m:^[^#\n]*\n(\\s*#\\s*@API[^\n]*(?:\n\\s*#[^\n]*)*))")
)

// ParseFile reads an API file
func ParseFile(filename string) (*APIFile, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := string(bytes)
	topMatches := apiFileTopRegex.FindAllStringSubmatch(content, -1)
	if len(topMatches) != 1 {
		if len(topMatches) == 0 {
			return nil, nil
		}
		return nil, errors.New("Multiple APIs found in one file")
	}
	top := apiFileCommentStart.ReplaceAllLiteralString(topMatches[0][2], "")
	modelIdxs := apiFileModelStart.FindAllStringIndex(top, -1)
	models := make([]string, len(modelIdxs))
	modelIdxs = append(modelIdxs, []int{len(top)})
	for i := range models {
		models[i] = strings.TrimSpace(top[modelIdxs[i][0]:modelIdxs[i+1][0]])
	}
	methodComments := apiFileMethodStart.FindAllStringSubmatch(topMatches[0][3], -1)
	methods := make([]string, len(methodComments))
	for i, m := range methodComments {
		methods[i] = strings.TrimSpace(apiFileCommentStart.ReplaceAllLiteralString(m[1], ""))
	}
	return &APIFile{
		Topic:          strings.TrimSpace(topMatches[0][1]),
		Description:    strings.TrimSpace(apiFileLineTrim.ReplaceAllLiteralString(top[0:modelIdxs[0][0]], " ")),
		ModelComments:  models,
		MethodComments: methods,
	}, nil
}
