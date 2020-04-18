package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	announcementReplyTemplate *AnnouncementReply
	// AnnouncementReplyChildCtor for parsing a template
	AnnouncementReplyChildCtor = func() (htmlgen.Section, []htmlgen.ChildConstructor) {
		return CreateAnnouncementReply(), []htmlgen.ChildConstructor{
			func() (htmlgen.Section, []htmlgen.ChildConstructor) {
				return CreateAnnouncementReply(), []htmlgen.ChildConstructor{}
			},
		}
	}
)

// AnnouncementReply HTML template
type AnnouncementReply struct {
	Data   canvas.DiscussionTopicFullViewEntry
	User   canvas.DiscussionTopicFullViewUser
	format *htmlgen.FormatSection
}

// CreateAnnouncementReply creates a new template
func CreateAnnouncementReply() *AnnouncementReply {
	obj := &AnnouncementReply{}
	args := []interface{}{
		&obj.User.DisplayName,
		htmlgen.CreateDateTimeFormat(&obj.Data.UpdatedAt),
		&obj.Data.Message,
		htmlgen.FormatSectionChild,
	}
	if announcementReplyTemplate == nil {
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
		obj.format = announcementReplyTemplate.format.Clone(args)
	}
	return obj
}

func init() {
	announcementReplyTemplate = CreateAnnouncementReply()
}

// AppendChild adds a child to the section
func (t *AnnouncementReply) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *AnnouncementReply) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *AnnouncementReply) String() string {
	return t.format.String()
}

// Parse the template
func (t *AnnouncementReply) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
