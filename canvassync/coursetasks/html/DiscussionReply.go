package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	discussionReplyTemplate *DiscussionReply
	// DiscussionReplyChildCtor for parsing a template
	DiscussionReplyChildCtor = func() (htmlgen.Section, []htmlgen.ChildConstructor) {
		return CreateDiscussionReply(), []htmlgen.ChildConstructor{
			func() (htmlgen.Section, []htmlgen.ChildConstructor) {
				return CreateDiscussionReply(), []htmlgen.ChildConstructor{}
			},
		}
	}
)

// DiscussionReply HTML template
type DiscussionReply struct {
	Data   canvas.DiscussionTopicFullViewEntry
	User   canvas.DiscussionTopicFullViewUser
	format *htmlgen.FormatSection
}

// CreateDiscussionReply creates a new template
func CreateDiscussionReply() *DiscussionReply {
	obj := &DiscussionReply{}
	args := []interface{}{
		&obj.User.DisplayName,
		htmlgen.CreateDateTimeFormat(&obj.Data.UpdatedAt),
		&obj.Data.Message,
		htmlgen.FormatSectionChild,
	}
	if discussionReplyTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	<h3>Reply from %s at %s:</h3>
	<div>
		%s
	</div>
	<div style="padding-left: 40px">
		%s
	</div>
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = discussionReplyTemplate.format.Clone(args)
	}
	return obj
}

func init() {
	discussionReplyTemplate = CreateDiscussionReply()
}

// AppendChild adds a child to the section
func (t *DiscussionReply) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *DiscussionReply) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *DiscussionReply) String() string {
	return t.format.String()
}

// Parse the template
func (t *DiscussionReply) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
