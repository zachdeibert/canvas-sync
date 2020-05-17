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
		model("Assignment").property("rubric_settings").setType("string", "interface{}").done().done().
		method("FilesListFiles").setMethodEndPoint("folders/<folder_id>/files", "courses/<course_id>/files").done().
		method("AssignmentGroupsListAssignmentGroups").setMethodEndPoint("", "courses/<course_id>/assignment_groups").
		arg("include").setType("string", "[]string").done().done().
		model("AssignmentGroup").property("assignments").setType("[]int", "[]Assignment").done().
		property("group_weight").setType("int", "float64").done().
		property("integration_data").setType("map[interface{}]interface{}", "map[string]interface{}").done().done().
		method("ModulesListModules").setMethodEndPoint("courses/222/modules", "courses/<course_id>/modules").
		arg("include").setType("string", "[]string").done().done().
		method("CoursesListUsersInCourse").setMethodEndPoint("", "courses/<course_id>/users").
		arg("include").setType("string", "[]string").done().done().
		model("Grade").property("current_score").setType("string", "float64").done().
		property("final_score").setType("string", "float64").done().done().
		method("PagesListPages").setMethodEndPoint("courses/123/pages", "courses/<course_id>/pages").done().
		model("User").addProperties(apisync.ModelProperty{
		Name:        "display_name",
		Description: "",
		Example:     "",
		Type:        "string",
		EnumValues:  []string{},
	}).done().
		method("PagesShowPage").setMethodEndPoint("courses/123/pages/my-page-url", "courses/<course_id>/pages/<url>").done().
		method("SubmissionsGetASingleSubmission").setMethodEndPoint("", "courses/<course_id>/assignments/<assignment_id>/submissions/<user_id>").
		setMethodReturnType("interface{}", "Submission").
		arg("include").setType("string", "[]string").done().done().
		model("SubmissionComment").addProperties(apisync.ModelProperty{
		Name:        "attachments",
		Description: "",
		Example:     "",
		Type:        "[]FileAttachment",
		EnumValues:  []string{},
	}).property("author").setType("string", "User").done()
}
