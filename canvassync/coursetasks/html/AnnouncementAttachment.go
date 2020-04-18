package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	announcementAttachmentTemplate *AnnouncementAttachment
	// AnnouncementAttachmentChildCtor for parsing a template
	AnnouncementAttachmentChildCtor = func() (htmlgen.Section, []htmlgen.ChildConstructor) {
		return CreateAnnouncementAttachment(), []htmlgen.ChildConstructor{}
	}
)

// AnnouncementAttachment HTML template
type AnnouncementAttachment struct {
	Data   canvas.FileAttachment
	format *htmlgen.FormatSection
}

// CreateAnnouncementAttachment creates a new template
func CreateAnnouncementAttachment() *AnnouncementAttachment {
	obj := &AnnouncementAttachment{}
	args := []interface{}{
		&obj.Data.DisplayName,
	}
	if announcementAttachmentTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	Attached file %s
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = announcementAttachmentTemplate.format.Clone(args)
	}
	return obj
}

func init() {
	announcementAttachmentTemplate = CreateAnnouncementAttachment()
}

// AppendChild adds a child to the section
func (t *AnnouncementAttachment) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *AnnouncementAttachment) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *AnnouncementAttachment) String() string {
	return t.format.String()
}

// Parse the template
func (t *AnnouncementAttachment) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
