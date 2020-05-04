package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	discussionRootTemplate *DiscussionRoot
	// DiscussionRootChildCtor for parsing a template
	DiscussionRootChildCtor = func() (htmlgen.Section, []htmlgen.ChildConstructor) {
		return CreateDiscussionRoot(), []htmlgen.ChildConstructor{
			DiscussionReplyChildCtor,
		}
	}
)

// DiscussionRoot HTML template
type DiscussionRoot struct {
	Data   canvas.DiscussionTopic
	format *htmlgen.FormatSection
}

// CreateDiscussionRoot creates a new template
func CreateDiscussionRoot() *DiscussionRoot {
	obj := &DiscussionRoot{}
	args := []interface{}{
		&obj.Data.Title,
		&obj.Data.UserName,
		htmlgen.CreateDateTimeFormat(&obj.Data.PostedAt),
		&obj.Data.Message,
		htmlgen.FormatSectionChild,
		htmlgen.CreateDateTimeFormat(&obj.Data.LastReplyAt),
	}
	if discussionRootTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	<h1>%s</h1>
	<h3>Posted by %s at %s:</h3>
	<main>
		%s
	</main>
	<div>
		%s
	</div>
	<footer>Last updated at %s</footer>
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = discussionRootTemplate.format.Clone(args)
	}
	return obj
}

func init() {
	discussionRootTemplate = CreateDiscussionRoot()
}

// AppendChild adds a child to the section
func (t *DiscussionRoot) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *DiscussionRoot) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *DiscussionRoot) String() string {
	return t.format.String()
}

// Parse the template
func (t *DiscussionRoot) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
