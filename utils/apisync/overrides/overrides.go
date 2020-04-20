package overrides

import "github.com/zachdeibert/canvas-sync/utils/apisync"

// ApplyOverrides applies overrides to the API file (fixes for when the documentation comments are wrong)
func ApplyOverrides(models *[]*apisync.Model, methods *[]apisync.MethodAPIPair) {
	createContext(models, methods).
		model("Submission").addProperties(apisync.ModelProperty{
		Name:        "attachments",
		Description: "",
		Example:     "",
		Type:        "[]FileAttachment",
		EnumValues:  []string{},
	}).done().
		method("AnnouncementsListAnnouncements").arg("context_codes").setType("interface{}", "[]string").done().done().
		method("AssignmentsListAssignments").
		setMethodEndPoint("", "courses/<course_id>/assignments").
		arg("include").setType("string", "[]string").done().done().
		method("CoursesListYourCourses").setMethodEndPoint("", "courses").done().
		method("DiscussionTopicsGetTheFullTopic").setMethodReturnType("interface{}", "DiscussionTopicFullView").done().
		model("LockInfo").property("context_module").setType("string", "interface{}").done().done().
		model("RubricCriteria").property("points").setType("int", "float64").done().done().
		model("RubricRating").property("points").setType("int", "float64").done().done().
		model("Assignment").property("rubric_settings").setType("string", "interface{}").done().done()
}