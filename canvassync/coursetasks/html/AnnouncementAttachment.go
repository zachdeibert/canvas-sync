package html

import (
	"github.com/zachdeibert/canvas-sync/canvas/model"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	announcementAttachmentTemplate *AnnouncementAttachment
	// AnnouncementAttachmentChildCtors for parsing a template
	AnnouncementAttachmentChildCtors = []htmlgen.ChildConstructor{}
)

// AnnouncementAttachment HTML template
type AnnouncementAttachment struct {
	Data   model.File
	format *htmlgen.FormatSection
}

// CreateAnnouncementAttachment creates a new template
func CreateAnnouncementAttachment() *AnnouncementAttachment {
	obj := &AnnouncementAttachment{}
	args := []interface{}{
		&obj.Data.DisplayName,
		&obj.Data.ID,
	}
	if announcementAttachmentTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	Attached file %s (id = %d)
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = announcementAttachmentTemplate.format.Clone(args)
	}
	return obj
}

// ParseAnnouncementAttachment parses a string to a template
func ParseAnnouncementAttachment(str string) *AnnouncementAttachment {
	o := CreateAnnouncementAttachment()
	if _, ok := o.Parse(str, AnnouncementAttachmentChildCtors); ok {
		return o
	}
	return nil
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
