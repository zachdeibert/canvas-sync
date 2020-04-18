package html

import (
	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/htmlgen"
)

var (
	announcementTemplate *Announcement
	// AnnouncementChildCtors for parsing a template
	AnnouncementChildCtors = []htmlgen.ChildConstructor{
		func() (htmlgen.Section, []htmlgen.ChildConstructor) {
			return CreateAnnouncementAttachment(), []htmlgen.ChildConstructor{}
		},
	}
)

// Announcement HTML template
type Announcement struct {
	Data   canvas.DiscussionTopic
	format *htmlgen.FormatSection
}

// CreateAnnouncement creates a new template
func CreateAnnouncement() *Announcement {
	obj := &Announcement{}
	args := []interface{}{
		&obj.Data.Title,
		&obj.Data.UserName,
		&obj.Data.Message,
		htmlgen.FormatSectionChild,
		htmlgen.CreateDateTimeFormat(&obj.Data.LastReplyAt),
	}
	if announcementTemplate == nil {
		var err error
		if obj.format, err = htmlgen.CreateFormatSection(`
<div>
	<main>
		<h1>%s</h1>
		<h3>From %s:</h3>
		<div>
			%s
		</div>
	</main>
	%s
	<footer>
		Last updated %s
	</footer>
</div>
`, args); err != nil {
			panic(err)
		}
	} else {
		obj.format = announcementTemplate.format.Clone(args)
	}
	return obj
}

// ParseAnnouncement parses a string to a template
func ParseAnnouncement(str string) *Announcement {
	o := CreateAnnouncement()
	if _, ok := o.Parse(str, AnnouncementChildCtors); ok {
		return o
	}
	return nil
}

func init() {
	announcementTemplate = CreateAnnouncement()
}

// AppendChild adds a child to the section
func (t *Announcement) AppendChild(child htmlgen.Section) {
	t.format.AppendChild(child)
}

// Children gets the child elements
func (t *Announcement) Children() []htmlgen.Section {
	return t.format.Children()
}

func (t *Announcement) String() string {
	return t.format.String()
}

// Parse the template
func (t *Announcement) Parse(str string, childCtors []htmlgen.ChildConstructor) (string, bool) {
	return t.format.Parse(str, childCtors)
}
