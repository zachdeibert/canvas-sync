package canvas

import "strings"

const (
	// ResponseType is the MIME type of the HTTP response data
	ResponseType = "application/x-http-response"
	// RedirectType is the MIME type of a redirected HTTP request
	RedirectType = "application/x-http-redirect"
	// DefaultType is the type to default to when no other associations were found
	DefaultType = "application/octet-stream"
)

// FileAssociations is a map of MIME types to file extensions
var FileAssociations = map[string]string{
	"application/json": ".json",
	"text/html":        ".html",
	RedirectType:       ".url",
	DefaultType:        ".bin",
}

func mimeToExt(mime string) string {
	mimeParts := strings.Split(mime, "+")
	extParts := make([]string, len(mimeParts))
	for i, p := range mimeParts {
		if ext, ok := FileAssociations[p]; ok {
			extParts[i] = ext
		} else {
			extParts[i] = FileAssociations[DefaultType]
		}
	}
	return strings.Join(extParts, "")
}
