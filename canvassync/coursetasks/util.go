package coursetasks

import "regexp"

// InvalidPathRunes matches runes that are invalid in a path
var InvalidPathRunes = regexp.MustCompile("[^-_a-zA-Z0-9 ()]")
