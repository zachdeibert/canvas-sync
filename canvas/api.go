package canvas

import (
	"fmt"
	"time"

	"github.com/zachdeibert/canvas-sync/task"
)

// AccountNotificationIcon enumeration
type AccountNotificationIcon string

const (
	// AccountNotificationIconWarning enum value ("warning")
	AccountNotificationIconWarning AccountNotificationIcon = "warning"
	// AccountNotificationIconInformation enum value ("information")
	AccountNotificationIconInformation AccountNotificationIcon = "information"
	// AccountNotificationIconQuestion enum value ("question")
	AccountNotificationIconQuestion AccountNotificationIcon = "question"
	// AccountNotificationIconError enum value ("error")
	AccountNotificationIconError AccountNotificationIcon = "error"
	// AccountNotificationIconCalendar enum value ("calendar")
	AccountNotificationIconCalendar AccountNotificationIcon = "calendar"
)

// ReportParametersOrder enumeration
type ReportParametersOrder string

const (
	// ReportParametersOrderUsers enum value ("users")
	ReportParametersOrderUsers ReportParametersOrder = "users"
	// ReportParametersOrderCourses enum value ("courses")
	ReportParametersOrderCourses ReportParametersOrder = "courses"
	// ReportParametersOrderOutcomes enum value ("outcomes")
	ReportParametersOrderOutcomes ReportParametersOrder = "outcomes"
)

// AppointmentGroupParticipantType enumeration
type AppointmentGroupParticipantType string

const (
	// AppointmentGroupParticipantTypeUser enum value ("User")
	AppointmentGroupParticipantTypeUser AppointmentGroupParticipantType = "User"
	// AppointmentGroupParticipantTypeGroup enum value ("Group")
	AppointmentGroupParticipantTypeGroup AppointmentGroupParticipantType = "Group"
)

// AppointmentGroupParticipantVisibility enumeration
type AppointmentGroupParticipantVisibility string

const (
	// AppointmentGroupParticipantVisibilityPrivate enum value ("private")
	AppointmentGroupParticipantVisibilityPrivate AppointmentGroupParticipantVisibility = "private"
	// AppointmentGroupParticipantVisibilityProtected enum value ("protected")
	AppointmentGroupParticipantVisibilityProtected AppointmentGroupParticipantVisibility = "protected"
)

// AppointmentGroupWorkflowState enumeration
type AppointmentGroupWorkflowState string

const (
	// AppointmentGroupWorkflowStatePending enum value ("pending")
	AppointmentGroupWorkflowStatePending AppointmentGroupWorkflowState = "pending"
	// AppointmentGroupWorkflowStateActive enum value ("active")
	AppointmentGroupWorkflowStateActive AppointmentGroupWorkflowState = "active"
	// AppointmentGroupWorkflowStateDeleted enum value ("deleted")
	AppointmentGroupWorkflowStateDeleted AppointmentGroupWorkflowState = "deleted"
)

// AssignmentGradingType enumeration
type AssignmentGradingType string

const (
	// AssignmentGradingTypePassFail enum value ("pass_fail")
	AssignmentGradingTypePassFail AssignmentGradingType = "pass_fail"
	// AssignmentGradingTypePercent enum value ("percent")
	AssignmentGradingTypePercent AssignmentGradingType = "percent"
	// AssignmentGradingTypeLetterGrade enum value ("letter_grade")
	AssignmentGradingTypeLetterGrade AssignmentGradingType = "letter_grade"
	// AssignmentGradingTypeGpaScale enum value ("gpa_scale")
	AssignmentGradingTypeGpaScale AssignmentGradingType = "gpa_scale"
	// AssignmentGradingTypePoints enum value ("points")
	AssignmentGradingTypePoints AssignmentGradingType = "points"
)

// AssignmentSubmissionTypes enumeration
type AssignmentSubmissionTypes string

const (
	// AssignmentSubmissionTypesDiscussionTopic enum value ("discussion_topic")
	AssignmentSubmissionTypesDiscussionTopic AssignmentSubmissionTypes = "discussion_topic"
	// AssignmentSubmissionTypesOnlineQuiz enum value ("online_quiz")
	AssignmentSubmissionTypesOnlineQuiz AssignmentSubmissionTypes = "online_quiz"
	// AssignmentSubmissionTypesOnPaper enum value ("on_paper")
	AssignmentSubmissionTypesOnPaper AssignmentSubmissionTypes = "on_paper"
	// AssignmentSubmissionTypesNotGraded enum value ("not_graded")
	AssignmentSubmissionTypesNotGraded AssignmentSubmissionTypes = "not_graded"
	// AssignmentSubmissionTypesNone enum value ("none")
	AssignmentSubmissionTypesNone AssignmentSubmissionTypes = "none"
	// AssignmentSubmissionTypesExternalTool enum value ("external_tool")
	AssignmentSubmissionTypesExternalTool AssignmentSubmissionTypes = "external_tool"
	// AssignmentSubmissionTypesOnlineTextEntry enum value ("online_text_entry")
	AssignmentSubmissionTypesOnlineTextEntry AssignmentSubmissionTypes = "online_text_entry"
	// AssignmentSubmissionTypesOnlineURL enum value ("online_url")
	AssignmentSubmissionTypesOnlineURL AssignmentSubmissionTypes = "online_url"
	// AssignmentSubmissionTypesOnlineUpload enum value ("online_upload")
	AssignmentSubmissionTypesOnlineUpload AssignmentSubmissionTypes = "online_upload"
	// AssignmentSubmissionTypesMediaRecording enum value ("media_recording")
	AssignmentSubmissionTypesMediaRecording AssignmentSubmissionTypes = "media_recording"
)

// AuthenticationEventEventType enumeration
type AuthenticationEventEventType string

const (
	// AuthenticationEventEventTypeLogin enum value ("login")
	AuthenticationEventEventTypeLogin AuthenticationEventEventType = "login"
	// AuthenticationEventEventTypeLogout enum value ("logout")
	AuthenticationEventEventTypeLogout AuthenticationEventEventType = "logout"
)

// AssignmentEventWorkflowState enumeration
type AssignmentEventWorkflowState string

const (
	// AssignmentEventWorkflowStatePublished enum value ("published")
	AssignmentEventWorkflowStatePublished AssignmentEventWorkflowState = "published"
	// AssignmentEventWorkflowStateDeleted enum value ("deleted")
	AssignmentEventWorkflowStateDeleted AssignmentEventWorkflowState = "deleted"
)

// CollaboratorType enumeration
type CollaboratorType string

const (
	// CollaboratorTypeUser enum value ("user")
	CollaboratorTypeUser CollaboratorType = "user"
	// CollaboratorTypeGroup enum value ("group")
	CollaboratorTypeGroup CollaboratorType = "group"
)

// CommMessageWorkflowState enumeration
type CommMessageWorkflowState string

const (
	// CommMessageWorkflowStateCreated enum value ("created")
	CommMessageWorkflowStateCreated CommMessageWorkflowState = "created"
	// CommMessageWorkflowStateStaged enum value ("staged")
	CommMessageWorkflowStateStaged CommMessageWorkflowState = "staged"
	// CommMessageWorkflowStateSending enum value ("sending")
	CommMessageWorkflowStateSending CommMessageWorkflowState = "sending"
	// CommMessageWorkflowStateSent enum value ("sent")
	CommMessageWorkflowStateSent CommMessageWorkflowState = "sent"
	// CommMessageWorkflowStateBounced enum value ("bounced")
	CommMessageWorkflowStateBounced CommMessageWorkflowState = "bounced"
	// CommMessageWorkflowStateDashboard enum value ("dashboard")
	CommMessageWorkflowStateDashboard CommMessageWorkflowState = "dashboard"
	// CommMessageWorkflowStateCancelled enum value ("cancelled")
	CommMessageWorkflowStateCancelled CommMessageWorkflowState = "cancelled"
	// CommMessageWorkflowStateClosed enum value ("closed")
	CommMessageWorkflowStateClosed CommMessageWorkflowState = "closed"
)

// CommunicationChannelType enumeration
type CommunicationChannelType string

const (
	// CommunicationChannelTypeEmail enum value ("email")
	CommunicationChannelTypeEmail CommunicationChannelType = "email"
	// CommunicationChannelTypePush enum value ("push")
	CommunicationChannelTypePush CommunicationChannelType = "push"
	// CommunicationChannelTypeSms enum value ("sms")
	CommunicationChannelTypeSms CommunicationChannelType = "sms"
	// CommunicationChannelTypeTwitter enum value ("twitter")
	CommunicationChannelTypeTwitter CommunicationChannelType = "twitter"
)

// CommunicationChannelWorkflowState enumeration
type CommunicationChannelWorkflowState string

const (
	// CommunicationChannelWorkflowStateUnconfirmed enum value ("unconfirmed")
	CommunicationChannelWorkflowStateUnconfirmed CommunicationChannelWorkflowState = "unconfirmed"
	// CommunicationChannelWorkflowStateActive enum value ("active")
	CommunicationChannelWorkflowStateActive CommunicationChannelWorkflowState = "active"
)

// ContentExportExportType enumeration
type ContentExportExportType string

const (
	// ContentExportExportTypeCommonCartridge enum value ("common_cartridge")
	ContentExportExportTypeCommonCartridge ContentExportExportType = "common_cartridge"
	// ContentExportExportTypeQti enum value ("qti")
	ContentExportExportTypeQti ContentExportExportType = "qti"
)

// ContentExportWorkflowState enumeration
type ContentExportWorkflowState string

const (
	// ContentExportWorkflowStateCreated enum value ("created")
	ContentExportWorkflowStateCreated ContentExportWorkflowState = "created"
	// ContentExportWorkflowStateExporting enum value ("exporting")
	ContentExportWorkflowStateExporting ContentExportWorkflowState = "exporting"
	// ContentExportWorkflowStateExported enum value ("exported")
	ContentExportWorkflowStateExported ContentExportWorkflowState = "exported"
	// ContentExportWorkflowStateFailed enum value ("failed")
	ContentExportWorkflowStateFailed ContentExportWorkflowState = "failed"
)

// ContentMigrationWorkflowState enumeration
type ContentMigrationWorkflowState string

const (
	// ContentMigrationWorkflowStatePreProcessing enum value ("pre_processing")
	ContentMigrationWorkflowStatePreProcessing ContentMigrationWorkflowState = "pre_processing"
	// ContentMigrationWorkflowStatePreProcessed enum value ("pre_processed")
	ContentMigrationWorkflowStatePreProcessed ContentMigrationWorkflowState = "pre_processed"
	// ContentMigrationWorkflowStateRunning enum value ("running")
	ContentMigrationWorkflowStateRunning ContentMigrationWorkflowState = "running"
	// ContentMigrationWorkflowStateWaitingForSelect enum value ("waiting_for_select")
	ContentMigrationWorkflowStateWaitingForSelect ContentMigrationWorkflowState = "waiting_for_select"
	// ContentMigrationWorkflowStateCompleted enum value ("completed")
	ContentMigrationWorkflowStateCompleted ContentMigrationWorkflowState = "completed"
	// ContentMigrationWorkflowStateFailed enum value ("failed")
	ContentMigrationWorkflowStateFailed ContentMigrationWorkflowState = "failed"
)

// CompletionRequirementType enumeration
type CompletionRequirementType string

const (
	// CompletionRequirementTypeMustView enum value ("must_view")
	CompletionRequirementTypeMustView CompletionRequirementType = "must_view"
	// CompletionRequirementTypeMustSubmit enum value ("must_submit")
	CompletionRequirementTypeMustSubmit CompletionRequirementType = "must_submit"
	// CompletionRequirementTypeMustContribute enum value ("must_contribute")
	CompletionRequirementTypeMustContribute CompletionRequirementType = "must_contribute"
	// CompletionRequirementTypeMinScore enum value ("min_score")
	CompletionRequirementTypeMinScore CompletionRequirementType = "min_score"
)

// ModuleItemType enumeration
type ModuleItemType string

const (
	// ModuleItemTypeFile enum value ("File")
	ModuleItemTypeFile ModuleItemType = "File"
	// ModuleItemTypePage enum value ("Page")
	ModuleItemTypePage ModuleItemType = "Page"
	// ModuleItemTypeDiscussion enum value ("Discussion")
	ModuleItemTypeDiscussion ModuleItemType = "Discussion"
	// ModuleItemTypeAssignment enum value ("Assignment")
	ModuleItemTypeAssignment ModuleItemType = "Assignment"
	// ModuleItemTypeQuiz enum value ("Quiz")
	ModuleItemTypeQuiz ModuleItemType = "Quiz"
	// ModuleItemTypeSubHeader enum value ("SubHeader")
	ModuleItemTypeSubHeader ModuleItemType = "SubHeader"
	// ModuleItemTypeExternalURL enum value ("ExternalUrl")
	ModuleItemTypeExternalURL ModuleItemType = "ExternalUrl"
	// ModuleItemTypeExternalTool enum value ("ExternalTool")
	ModuleItemTypeExternalTool ModuleItemType = "ExternalTool"
)

// ModuleState enumeration
type ModuleState string

const (
	// ModuleStateLocked enum value ("locked")
	ModuleStateLocked ModuleState = "locked"
	// ModuleStateUnlocked enum value ("unlocked")
	ModuleStateUnlocked ModuleState = "unlocked"
	// ModuleStateStarted enum value ("started")
	ModuleStateStarted ModuleState = "started"
	// ModuleStateCompleted enum value ("completed")
	ModuleStateCompleted ModuleState = "completed"
)

// ModuleWorkflowState enumeration
type ModuleWorkflowState string

const (
	// ModuleWorkflowStateActive enum value ("active")
	ModuleWorkflowStateActive ModuleWorkflowState = "active"
	// ModuleWorkflowStateDeleted enum value ("deleted")
	ModuleWorkflowStateDeleted ModuleWorkflowState = "deleted"
)

// CourseDefaultView enumeration
type CourseDefaultView string

const (
	// CourseDefaultViewFeed enum value ("feed")
	CourseDefaultViewFeed CourseDefaultView = "feed"
	// CourseDefaultViewWiki enum value ("wiki")
	CourseDefaultViewWiki CourseDefaultView = "wiki"
	// CourseDefaultViewModules enum value ("modules")
	CourseDefaultViewModules CourseDefaultView = "modules"
	// CourseDefaultViewSyllabus enum value ("syllabus")
	CourseDefaultViewSyllabus CourseDefaultView = "syllabus"
	// CourseDefaultViewAssignments enum value ("assignments")
	CourseDefaultViewAssignments CourseDefaultView = "assignments"
)

// CourseWorkflowState enumeration
type CourseWorkflowState string

const (
	// CourseWorkflowStateUnpublished enum value ("unpublished")
	CourseWorkflowStateUnpublished CourseWorkflowState = "unpublished"
	// CourseWorkflowStateAvailable enum value ("available")
	CourseWorkflowStateAvailable CourseWorkflowState = "available"
	// CourseWorkflowStateCompleted enum value ("completed")
	CourseWorkflowStateCompleted CourseWorkflowState = "completed"
	// CourseWorkflowStateDeleted enum value ("deleted")
	CourseWorkflowStateDeleted CourseWorkflowState = "deleted"
)

// DiscussionTopicDiscussionType enumeration
type DiscussionTopicDiscussionType string

const (
	// DiscussionTopicDiscussionTypeSideComment enum value ("side_comment")
	DiscussionTopicDiscussionTypeSideComment DiscussionTopicDiscussionType = "side_comment"
	// DiscussionTopicDiscussionTypeThreaded enum value ("threaded")
	DiscussionTopicDiscussionTypeThreaded DiscussionTopicDiscussionType = "threaded"
)

// DiscussionTopicReadState enumeration
type DiscussionTopicReadState string

const (
	// DiscussionTopicReadStateRead enum value ("read")
	DiscussionTopicReadStateRead DiscussionTopicReadState = "read"
	// DiscussionTopicReadStateUnread enum value ("unread")
	DiscussionTopicReadStateUnread DiscussionTopicReadState = "unread"
)

// DiscussionTopicSubscriptionHold enumeration
type DiscussionTopicSubscriptionHold string

const (
	// DiscussionTopicSubscriptionHoldInitialPostRequired enum value ("initial_post_required")
	DiscussionTopicSubscriptionHoldInitialPostRequired DiscussionTopicSubscriptionHold = "initial_post_required"
	// DiscussionTopicSubscriptionHoldNotInGroupSet enum value ("not_in_group_set")
	DiscussionTopicSubscriptionHoldNotInGroupSet DiscussionTopicSubscriptionHold = "not_in_group_set"
	// DiscussionTopicSubscriptionHoldNotInGroup enum value ("not_in_group")
	DiscussionTopicSubscriptionHoldNotInGroup DiscussionTopicSubscriptionHold = "not_in_group"
	// DiscussionTopicSubscriptionHoldTopicIsAnnouncement enum value ("topic_is_announcement")
	DiscussionTopicSubscriptionHoldTopicIsAnnouncement DiscussionTopicSubscriptionHold = "topic_is_announcement"
)

// EpubExportWorkflowState enumeration
type EpubExportWorkflowState string

const (
	// EpubExportWorkflowStateCreated enum value ("created")
	EpubExportWorkflowStateCreated EpubExportWorkflowState = "created"
	// EpubExportWorkflowStateExporting enum value ("exporting")
	EpubExportWorkflowStateExporting EpubExportWorkflowState = "exporting"
	// EpubExportWorkflowStateExported enum value ("exported")
	EpubExportWorkflowStateExported EpubExportWorkflowState = "exported"
	// EpubExportWorkflowStateGenerating enum value ("generating")
	EpubExportWorkflowStateGenerating EpubExportWorkflowState = "generating"
	// EpubExportWorkflowStateGenerated enum value ("generated")
	EpubExportWorkflowStateGenerated EpubExportWorkflowState = "generated"
	// EpubExportWorkflowStateFailed enum value ("failed")
	EpubExportWorkflowStateFailed EpubExportWorkflowState = "failed"
)

// ExternalFeedVerbosity enumeration
type ExternalFeedVerbosity string

const (
	// ExternalFeedVerbosityLinkOnly enum value ("link_only")
	ExternalFeedVerbosityLinkOnly ExternalFeedVerbosity = "link_only"
	// ExternalFeedVerbosityTruncate enum value ("truncate")
	ExternalFeedVerbosityTruncate ExternalFeedVerbosity = "truncate"
	// ExternalFeedVerbosityFull enum value ("full")
	ExternalFeedVerbosityFull ExternalFeedVerbosity = "full"
)

// FavoriteContextType enumeration
type FavoriteContextType string

const (
	// FavoriteContextTypeCourse enum value ("Course")
	FavoriteContextTypeCourse FavoriteContextType = "Course"
)

// FeatureAppliesTo enumeration
type FeatureAppliesTo string

const (
	// FeatureAppliesToCourse enum value ("Course")
	FeatureAppliesToCourse FeatureAppliesTo = "Course"
	// FeatureAppliesToRootAccount enum value ("RootAccount")
	FeatureAppliesToRootAccount FeatureAppliesTo = "RootAccount"
	// FeatureAppliesToAccount enum value ("Account")
	FeatureAppliesToAccount FeatureAppliesTo = "Account"
	// FeatureAppliesToUser enum value ("User")
	FeatureAppliesToUser FeatureAppliesTo = "User"
)

// FeatureFlagContextType enumeration
type FeatureFlagContextType string

const (
	// FeatureFlagContextTypeCourse enum value ("Course")
	FeatureFlagContextTypeCourse FeatureFlagContextType = "Course"
	// FeatureFlagContextTypeAccount enum value ("Account")
	FeatureFlagContextTypeAccount FeatureFlagContextType = "Account"
	// FeatureFlagContextTypeUser enum value ("User")
	FeatureFlagContextTypeUser FeatureFlagContextType = "User"
)

// FeatureFlagState enumeration
type FeatureFlagState string

const (
	// FeatureFlagStateOff enum value ("off")
	FeatureFlagStateOff FeatureFlagState = "off"
	// FeatureFlagStateAllowed enum value ("allowed")
	FeatureFlagStateAllowed FeatureFlagState = "allowed"
	// FeatureFlagStateOn enum value ("on")
	FeatureFlagStateOn FeatureFlagState = "on"
)

// GroupCategoryAutoLeader enumeration
type GroupCategoryAutoLeader string

const (
	// GroupCategoryAutoLeaderFirst enum value ("first")
	GroupCategoryAutoLeaderFirst GroupCategoryAutoLeader = "first"
	// GroupCategoryAutoLeaderRandom enum value ("random")
	GroupCategoryAutoLeaderRandom GroupCategoryAutoLeader = "random"
)

// GroupCategorySelfSignup enumeration
type GroupCategorySelfSignup string

const (
	// GroupCategorySelfSignupRestricted enum value ("restricted")
	GroupCategorySelfSignupRestricted GroupCategorySelfSignup = "restricted"
	// GroupCategorySelfSignupEnabled enum value ("enabled")
	GroupCategorySelfSignupEnabled GroupCategorySelfSignup = "enabled"
)

// GroupMembershipWorkflowState enumeration
type GroupMembershipWorkflowState string

const (
	// GroupMembershipWorkflowStateAccepted enum value ("accepted")
	GroupMembershipWorkflowStateAccepted GroupMembershipWorkflowState = "accepted"
	// GroupMembershipWorkflowStateInvited enum value ("invited")
	GroupMembershipWorkflowStateInvited GroupMembershipWorkflowState = "invited"
	// GroupMembershipWorkflowStateRequested enum value ("requested")
	GroupMembershipWorkflowStateRequested GroupMembershipWorkflowState = "requested"
)

// GroupJoinLevel enumeration
type GroupJoinLevel string

const (
	// GroupJoinLevelParentContextAutoJoin enum value ("parent_context_auto_join")
	GroupJoinLevelParentContextAutoJoin GroupJoinLevel = "parent_context_auto_join"
	// GroupJoinLevelParentContextRequest enum value ("parent_context_request")
	GroupJoinLevelParentContextRequest GroupJoinLevel = "parent_context_request"
	// GroupJoinLevelInvitationOnly enum value ("invitation_only")
	GroupJoinLevelInvitationOnly GroupJoinLevel = "invitation_only"
)

// GroupRole enumeration
type GroupRole string

const (
	// GroupRoleCommunities enum value ("communities")
	GroupRoleCommunities GroupRole = "communities"
	// GroupRoleStudentOrganized enum value ("student_organized")
	GroupRoleStudentOrganized GroupRole = "student_organized"
	// GroupRoleImported enum value ("imported")
	GroupRoleImported GroupRole = "imported"
)

// MigrationIssueIssueType enumeration
type MigrationIssueIssueType string

const (
	// MigrationIssueIssueTypeTodo enum value ("todo")
	MigrationIssueIssueTypeTodo MigrationIssueIssueType = "todo"
	// MigrationIssueIssueTypeWarning enum value ("warning")
	MigrationIssueIssueTypeWarning MigrationIssueIssueType = "warning"
	// MigrationIssueIssueTypeError enum value ("error")
	MigrationIssueIssueTypeError MigrationIssueIssueType = "error"
)

// MigrationIssueWorkflowState enumeration
type MigrationIssueWorkflowState string

const (
	// MigrationIssueWorkflowStateActive enum value ("active")
	MigrationIssueWorkflowStateActive MigrationIssueWorkflowState = "active"
	// MigrationIssueWorkflowStateResolved enum value ("resolved")
	MigrationIssueWorkflowStateResolved MigrationIssueWorkflowState = "resolved"
)

// NotificationPreferenceFrequency enumeration
type NotificationPreferenceFrequency string

const (
	// NotificationPreferenceFrequencyImmediately enum value ("immediately")
	NotificationPreferenceFrequencyImmediately NotificationPreferenceFrequency = "immediately"
	// NotificationPreferenceFrequencyDaily enum value ("daily")
	NotificationPreferenceFrequencyDaily NotificationPreferenceFrequency = "daily"
	// NotificationPreferenceFrequencyWeekly enum value ("weekly")
	NotificationPreferenceFrequencyWeekly NotificationPreferenceFrequency = "weekly"
	// NotificationPreferenceFrequencyNever enum value ("never")
	NotificationPreferenceFrequencyNever NotificationPreferenceFrequency = "never"
)

// OutcomeImportWorkflowState enumeration
type OutcomeImportWorkflowState string

const (
	// OutcomeImportWorkflowStateCreated enum value ("created")
	OutcomeImportWorkflowStateCreated OutcomeImportWorkflowState = "created"
	// OutcomeImportWorkflowStateImporting enum value ("importing")
	OutcomeImportWorkflowStateImporting OutcomeImportWorkflowState = "importing"
	// OutcomeImportWorkflowStateSucceeded enum value ("succeeded")
	OutcomeImportWorkflowStateSucceeded OutcomeImportWorkflowState = "succeeded"
	// OutcomeImportWorkflowStateFailed enum value ("failed")
	OutcomeImportWorkflowStateFailed OutcomeImportWorkflowState = "failed"
)

// OutcomeCalculationMethod enumeration
type OutcomeCalculationMethod string

const (
	// OutcomeCalculationMethodDecayingAverage enum value ("decaying_average")
	OutcomeCalculationMethodDecayingAverage OutcomeCalculationMethod = "decaying_average"
	// OutcomeCalculationMethodNMastery enum value ("n_mastery")
	OutcomeCalculationMethodNMastery OutcomeCalculationMethod = "n_mastery"
	// OutcomeCalculationMethodLatest enum value ("latest")
	OutcomeCalculationMethodLatest OutcomeCalculationMethod = "latest"
	// OutcomeCalculationMethodHighest enum value ("highest")
	OutcomeCalculationMethodHighest OutcomeCalculationMethod = "highest"
)

// ProgressWorkflowState enumeration
type ProgressWorkflowState string

const (
	// ProgressWorkflowStateQueued enum value ("queued")
	ProgressWorkflowStateQueued ProgressWorkflowState = "queued"
	// ProgressWorkflowStateRunning enum value ("running")
	ProgressWorkflowStateRunning ProgressWorkflowState = "running"
	// ProgressWorkflowStateCompleted enum value ("completed")
	ProgressWorkflowStateCompleted ProgressWorkflowState = "completed"
	// ProgressWorkflowStateFailed enum value ("failed")
	ProgressWorkflowStateFailed ProgressWorkflowState = "failed"
)

// SisAssignmentSubmissionTypes enumeration
type SisAssignmentSubmissionTypes string

const (
	// SisAssignmentSubmissionTypesDiscussionTopic enum value ("discussion_topic")
	SisAssignmentSubmissionTypesDiscussionTopic SisAssignmentSubmissionTypes = "discussion_topic"
	// SisAssignmentSubmissionTypesOnlineQuiz enum value ("online_quiz")
	SisAssignmentSubmissionTypesOnlineQuiz SisAssignmentSubmissionTypes = "online_quiz"
	// SisAssignmentSubmissionTypesOnPaper enum value ("on_paper")
	SisAssignmentSubmissionTypesOnPaper SisAssignmentSubmissionTypes = "on_paper"
	// SisAssignmentSubmissionTypesNotGraded enum value ("not_graded")
	SisAssignmentSubmissionTypesNotGraded SisAssignmentSubmissionTypes = "not_graded"
	// SisAssignmentSubmissionTypesNone enum value ("none")
	SisAssignmentSubmissionTypesNone SisAssignmentSubmissionTypes = "none"
	// SisAssignmentSubmissionTypesExternalTool enum value ("external_tool")
	SisAssignmentSubmissionTypesExternalTool SisAssignmentSubmissionTypes = "external_tool"
	// SisAssignmentSubmissionTypesOnlineTextEntry enum value ("online_text_entry")
	SisAssignmentSubmissionTypesOnlineTextEntry SisAssignmentSubmissionTypes = "online_text_entry"
	// SisAssignmentSubmissionTypesOnlineURL enum value ("online_url")
	SisAssignmentSubmissionTypesOnlineURL SisAssignmentSubmissionTypes = "online_url"
	// SisAssignmentSubmissionTypesOnlineUpload enum value ("online_upload")
	SisAssignmentSubmissionTypesOnlineUpload SisAssignmentSubmissionTypes = "online_upload"
	// SisAssignmentSubmissionTypesMediaRecording enum value ("media_recording")
	SisAssignmentSubmissionTypesMediaRecording SisAssignmentSubmissionTypes = "media_recording"
)

// SisImportWorkflowState enumeration
type SisImportWorkflowState string

const (
	// SisImportWorkflowStateInitializing enum value ("initializing")
	SisImportWorkflowStateInitializing SisImportWorkflowState = "initializing"
	// SisImportWorkflowStateCreated enum value ("created")
	SisImportWorkflowStateCreated SisImportWorkflowState = "created"
	// SisImportWorkflowStateImporting enum value ("importing")
	SisImportWorkflowStateImporting SisImportWorkflowState = "importing"
	// SisImportWorkflowStateCleanupBatch enum value ("cleanup_batch")
	SisImportWorkflowStateCleanupBatch SisImportWorkflowState = "cleanup_batch"
	// SisImportWorkflowStateImported enum value ("imported")
	SisImportWorkflowStateImported SisImportWorkflowState = "imported"
	// SisImportWorkflowStateImportedWithMessages enum value ("imported_with_messages")
	SisImportWorkflowStateImportedWithMessages SisImportWorkflowState = "imported_with_messages"
	// SisImportWorkflowStateAborted enum value ("aborted")
	SisImportWorkflowStateAborted SisImportWorkflowState = "aborted"
	// SisImportWorkflowStateFailed enum value ("failed")
	SisImportWorkflowStateFailed SisImportWorkflowState = "failed"
	// SisImportWorkflowStateFailedWithMessages enum value ("failed_with_messages")
	SisImportWorkflowStateFailedWithMessages SisImportWorkflowState = "failed_with_messages"
	// SisImportWorkflowStateRestoring enum value ("restoring")
	SisImportWorkflowStateRestoring SisImportWorkflowState = "restoring"
	// SisImportWorkflowStatePartiallyRestored enum value ("partially_restored")
	SisImportWorkflowStatePartiallyRestored SisImportWorkflowState = "partially_restored"
	// SisImportWorkflowStateRestored enum value ("restored")
	SisImportWorkflowStateRestored SisImportWorkflowState = "restored"
)

// SubmissionSubmissionType enumeration
type SubmissionSubmissionType string

const (
	// SubmissionSubmissionTypeOnlineTextEntry enum value ("online_text_entry")
	SubmissionSubmissionTypeOnlineTextEntry SubmissionSubmissionType = "online_text_entry"
	// SubmissionSubmissionTypeOnlineURL enum value ("online_url")
	SubmissionSubmissionTypeOnlineURL SubmissionSubmissionType = "online_url"
	// SubmissionSubmissionTypeOnlineUpload enum value ("online_upload")
	SubmissionSubmissionTypeOnlineUpload SubmissionSubmissionType = "online_upload"
	// SubmissionSubmissionTypeMediaRecording enum value ("media_recording")
	SubmissionSubmissionTypeMediaRecording SubmissionSubmissionType = "media_recording"
)

// SubmissionWorkflowState enumeration
type SubmissionWorkflowState string

const (
	// SubmissionWorkflowStateGraded enum value ("graded")
	SubmissionWorkflowStateGraded SubmissionWorkflowState = "graded"
	// SubmissionWorkflowStateSubmitted enum value ("submitted")
	SubmissionWorkflowStateSubmitted SubmissionWorkflowState = "submitted"
	// SubmissionWorkflowStateUnsubmitted enum value ("unsubmitted")
	SubmissionWorkflowStateUnsubmitted SubmissionWorkflowState = "unsubmitted"
	// SubmissionWorkflowStatePendingReview enum value ("pending_review")
	SubmissionWorkflowStatePendingReview SubmissionWorkflowState = "pending_review"
)

// AccountsListAccountsInclude enumeration
type AccountsListAccountsInclude string

const (
	// AccountsListAccountsIncludeLtiGUID enum value ("lti_guid")
	AccountsListAccountsIncludeLtiGUID AccountsListAccountsInclude = "lti_guid"
	// AccountsListAccountsIncludeRegistrationSettings enum value ("registration_settings")
	AccountsListAccountsIncludeRegistrationSettings AccountsListAccountsInclude = "registration_settings"
	// AccountsListAccountsIncludeServices enum value ("services")
	AccountsListAccountsIncludeServices AccountsListAccountsInclude = "services"
)

// AccountsListActiveCoursesInAnAccountEnrollmentType enumeration
type AccountsListActiveCoursesInAnAccountEnrollmentType string

const (
	// AccountsListActiveCoursesInAnAccountEnrollmentTypeTeacher enum value ("teacher")
	AccountsListActiveCoursesInAnAccountEnrollmentTypeTeacher AccountsListActiveCoursesInAnAccountEnrollmentType = "teacher"
	// AccountsListActiveCoursesInAnAccountEnrollmentTypeStudent enum value ("student")
	AccountsListActiveCoursesInAnAccountEnrollmentTypeStudent AccountsListActiveCoursesInAnAccountEnrollmentType = "student"
	// AccountsListActiveCoursesInAnAccountEnrollmentTypeTa enum value ("ta")
	AccountsListActiveCoursesInAnAccountEnrollmentTypeTa AccountsListActiveCoursesInAnAccountEnrollmentType = "ta"
	// AccountsListActiveCoursesInAnAccountEnrollmentTypeObserver enum value ("observer")
	AccountsListActiveCoursesInAnAccountEnrollmentTypeObserver AccountsListActiveCoursesInAnAccountEnrollmentType = "observer"
	// AccountsListActiveCoursesInAnAccountEnrollmentTypeDesigner enum value ("designer")
	AccountsListActiveCoursesInAnAccountEnrollmentTypeDesigner AccountsListActiveCoursesInAnAccountEnrollmentType = "designer"
)

// AccountsListActiveCoursesInAnAccountState enumeration
type AccountsListActiveCoursesInAnAccountState string

const (
	// AccountsListActiveCoursesInAnAccountStateCreated enum value ("created")
	AccountsListActiveCoursesInAnAccountStateCreated AccountsListActiveCoursesInAnAccountState = "created"
	// AccountsListActiveCoursesInAnAccountStateClaimed enum value ("claimed")
	AccountsListActiveCoursesInAnAccountStateClaimed AccountsListActiveCoursesInAnAccountState = "claimed"
	// AccountsListActiveCoursesInAnAccountStateAvailable enum value ("available")
	AccountsListActiveCoursesInAnAccountStateAvailable AccountsListActiveCoursesInAnAccountState = "available"
	// AccountsListActiveCoursesInAnAccountStateCompleted enum value ("completed")
	AccountsListActiveCoursesInAnAccountStateCompleted AccountsListActiveCoursesInAnAccountState = "completed"
	// AccountsListActiveCoursesInAnAccountStateDeleted enum value ("deleted")
	AccountsListActiveCoursesInAnAccountStateDeleted AccountsListActiveCoursesInAnAccountState = "deleted"
	// AccountsListActiveCoursesInAnAccountStateAll enum value ("all")
	AccountsListActiveCoursesInAnAccountStateAll AccountsListActiveCoursesInAnAccountState = "all"
)

// AccountsListActiveCoursesInAnAccountInclude enumeration
type AccountsListActiveCoursesInAnAccountInclude string

const (
	// AccountsListActiveCoursesInAnAccountIncludeSyllabusBody enum value ("syllabus_body")
	AccountsListActiveCoursesInAnAccountIncludeSyllabusBody AccountsListActiveCoursesInAnAccountInclude = "syllabus_body"
	// AccountsListActiveCoursesInAnAccountIncludeTerm enum value ("term")
	AccountsListActiveCoursesInAnAccountIncludeTerm AccountsListActiveCoursesInAnAccountInclude = "term"
	// AccountsListActiveCoursesInAnAccountIncludeCourseProgress enum value ("course_progress")
	AccountsListActiveCoursesInAnAccountIncludeCourseProgress AccountsListActiveCoursesInAnAccountInclude = "course_progress"
	// AccountsListActiveCoursesInAnAccountIncludeStorageQuotaUsedMb enum value ("storage_quota_used_mb")
	AccountsListActiveCoursesInAnAccountIncludeStorageQuotaUsedMb AccountsListActiveCoursesInAnAccountInclude = "storage_quota_used_mb"
	// AccountsListActiveCoursesInAnAccountIncludeTotalStudents enum value ("total_students")
	AccountsListActiveCoursesInAnAccountIncludeTotalStudents AccountsListActiveCoursesInAnAccountInclude = "total_students"
	// AccountsListActiveCoursesInAnAccountIncludeTeachers enum value ("teachers")
	AccountsListActiveCoursesInAnAccountIncludeTeachers AccountsListActiveCoursesInAnAccountInclude = "teachers"
	// AccountsListActiveCoursesInAnAccountIncludeAccountName enum value ("account_name")
	AccountsListActiveCoursesInAnAccountIncludeAccountName AccountsListActiveCoursesInAnAccountInclude = "account_name"
	// AccountsListActiveCoursesInAnAccountIncludeConcluded enum value ("concluded")
	AccountsListActiveCoursesInAnAccountIncludeConcluded AccountsListActiveCoursesInAnAccountInclude = "concluded"
)

// AccountsListActiveCoursesInAnAccountSort enumeration
type AccountsListActiveCoursesInAnAccountSort string

const (
	// AccountsListActiveCoursesInAnAccountSortCourseName enum value ("course_name")
	AccountsListActiveCoursesInAnAccountSortCourseName AccountsListActiveCoursesInAnAccountSort = "course_name"
	// AccountsListActiveCoursesInAnAccountSortSisCourseID enum value ("sis_course_id")
	AccountsListActiveCoursesInAnAccountSortSisCourseID AccountsListActiveCoursesInAnAccountSort = "sis_course_id"
	// AccountsListActiveCoursesInAnAccountSortTeacher enum value ("teacher")
	AccountsListActiveCoursesInAnAccountSortTeacher AccountsListActiveCoursesInAnAccountSort = "teacher"
	// AccountsListActiveCoursesInAnAccountSortAccountName enum value ("account_name")
	AccountsListActiveCoursesInAnAccountSortAccountName AccountsListActiveCoursesInAnAccountSort = "account_name"
)

// AccountsListActiveCoursesInAnAccountOrder enumeration
type AccountsListActiveCoursesInAnAccountOrder string

const (
	// AccountsListActiveCoursesInAnAccountOrderAsc enum value ("asc")
	AccountsListActiveCoursesInAnAccountOrderAsc AccountsListActiveCoursesInAnAccountOrder = "asc"
	// AccountsListActiveCoursesInAnAccountOrderDesc enum value ("desc")
	AccountsListActiveCoursesInAnAccountOrderDesc AccountsListActiveCoursesInAnAccountOrder = "desc"
)

// AccountsListActiveCoursesInAnAccountSearchBy enumeration
type AccountsListActiveCoursesInAnAccountSearchBy string

const (
	// AccountsListActiveCoursesInAnAccountSearchByCourse enum value ("course")
	AccountsListActiveCoursesInAnAccountSearchByCourse AccountsListActiveCoursesInAnAccountSearchBy = "course"
	// AccountsListActiveCoursesInAnAccountSearchByTeacher enum value ("teacher")
	AccountsListActiveCoursesInAnAccountSearchByTeacher AccountsListActiveCoursesInAnAccountSearchBy = "teacher"
)

// AppointmentGroupsListAppointmentGroupsScope enumeration
type AppointmentGroupsListAppointmentGroupsScope string

const (
	// AppointmentGroupsListAppointmentGroupsScopeReservable enum value ("reservable")
	AppointmentGroupsListAppointmentGroupsScopeReservable AppointmentGroupsListAppointmentGroupsScope = "reservable"
	// AppointmentGroupsListAppointmentGroupsScopeManageable enum value ("manageable")
	AppointmentGroupsListAppointmentGroupsScopeManageable AppointmentGroupsListAppointmentGroupsScope = "manageable"
)

// AppointmentGroupsListAppointmentGroupsInclude enumeration
type AppointmentGroupsListAppointmentGroupsInclude string

const (
	// AppointmentGroupsListAppointmentGroupsIncludeAppointments enum value ("appointments")
	AppointmentGroupsListAppointmentGroupsIncludeAppointments AppointmentGroupsListAppointmentGroupsInclude = "appointments"
	// AppointmentGroupsListAppointmentGroupsIncludeChildEvents enum value ("child_events")
	AppointmentGroupsListAppointmentGroupsIncludeChildEvents AppointmentGroupsListAppointmentGroupsInclude = "child_events"
	// AppointmentGroupsListAppointmentGroupsIncludeParticipantCount enum value ("participant_count")
	AppointmentGroupsListAppointmentGroupsIncludeParticipantCount AppointmentGroupsListAppointmentGroupsInclude = "participant_count"
	// AppointmentGroupsListAppointmentGroupsIncludeReservedTimes enum value ("reserved_times")
	AppointmentGroupsListAppointmentGroupsIncludeReservedTimes AppointmentGroupsListAppointmentGroupsInclude = "reserved_times"
	// AppointmentGroupsListAppointmentGroupsIncludeAllContextCodes enum value ("all_context_codes")
	AppointmentGroupsListAppointmentGroupsIncludeAllContextCodes AppointmentGroupsListAppointmentGroupsInclude = "all_context_codes"
)

// AppointmentGroupsGetASingleAppointmentGroupInclude enumeration
type AppointmentGroupsGetASingleAppointmentGroupInclude string

const (
	// AppointmentGroupsGetASingleAppointmentGroupIncludeChildEvents enum value ("child_events")
	AppointmentGroupsGetASingleAppointmentGroupIncludeChildEvents AppointmentGroupsGetASingleAppointmentGroupInclude = "child_events"
	// AppointmentGroupsGetASingleAppointmentGroupIncludeAppointments enum value ("appointments")
	AppointmentGroupsGetASingleAppointmentGroupIncludeAppointments AppointmentGroupsGetASingleAppointmentGroupInclude = "appointments"
	// AppointmentGroupsGetASingleAppointmentGroupIncludeAllContextCodes enum value ("all_context_codes")
	AppointmentGroupsGetASingleAppointmentGroupIncludeAllContextCodes AppointmentGroupsGetASingleAppointmentGroupInclude = "all_context_codes"
)

// AppointmentGroupsListUserParticipantsRegistrationStatus enumeration
type AppointmentGroupsListUserParticipantsRegistrationStatus string

const (
	// AppointmentGroupsListUserParticipantsRegistrationStatusAll enum value ("all")
	AppointmentGroupsListUserParticipantsRegistrationStatusAll AppointmentGroupsListUserParticipantsRegistrationStatus = "all"
)

// AppointmentGroupsListStudentGroupParticipantsRegistrationStatus enumeration
type AppointmentGroupsListStudentGroupParticipantsRegistrationStatus string

const (
	// AppointmentGroupsListStudentGroupParticipantsRegistrationStatusAll enum value ("all")
	AppointmentGroupsListStudentGroupParticipantsRegistrationStatusAll AppointmentGroupsListStudentGroupParticipantsRegistrationStatus = "all"
)

// AssignmentGroupsGetAnAssignmentGroupInclude enumeration
type AssignmentGroupsGetAnAssignmentGroupInclude string

const (
	// AssignmentGroupsGetAnAssignmentGroupIncludeAssignments enum value ("assignments")
	AssignmentGroupsGetAnAssignmentGroupIncludeAssignments AssignmentGroupsGetAnAssignmentGroupInclude = "assignments"
	// AssignmentGroupsGetAnAssignmentGroupIncludeDiscussionTopic enum value ("discussion_topic")
	AssignmentGroupsGetAnAssignmentGroupIncludeDiscussionTopic AssignmentGroupsGetAnAssignmentGroupInclude = "discussion_topic"
	// AssignmentGroupsGetAnAssignmentGroupIncludeAssignmentVisibility enum value ("assignment_visibility")
	AssignmentGroupsGetAnAssignmentGroupIncludeAssignmentVisibility AssignmentGroupsGetAnAssignmentGroupInclude = "assignment_visibility"
	// AssignmentGroupsGetAnAssignmentGroupIncludeSubmission enum value ("submission")
	AssignmentGroupsGetAnAssignmentGroupIncludeSubmission AssignmentGroupsGetAnAssignmentGroupInclude = "submission"
)

// AssignmentGroupsListAssignmentGroupsInclude enumeration
type AssignmentGroupsListAssignmentGroupsInclude string

const (
	// AssignmentGroupsListAssignmentGroupsIncludeAssignments enum value ("assignments")
	AssignmentGroupsListAssignmentGroupsIncludeAssignments AssignmentGroupsListAssignmentGroupsInclude = "assignments"
	// AssignmentGroupsListAssignmentGroupsIncludeDiscussionTopic enum value ("discussion_topic")
	AssignmentGroupsListAssignmentGroupsIncludeDiscussionTopic AssignmentGroupsListAssignmentGroupsInclude = "discussion_topic"
	// AssignmentGroupsListAssignmentGroupsIncludeAllDates enum value ("all_dates")
	AssignmentGroupsListAssignmentGroupsIncludeAllDates AssignmentGroupsListAssignmentGroupsInclude = "all_dates"
	// AssignmentGroupsListAssignmentGroupsIncludeAssignmentVisibility enum value ("assignment_visibility")
	AssignmentGroupsListAssignmentGroupsIncludeAssignmentVisibility AssignmentGroupsListAssignmentGroupsInclude = "assignment_visibility"
	// AssignmentGroupsListAssignmentGroupsIncludeOverrides enum value ("overrides")
	AssignmentGroupsListAssignmentGroupsIncludeOverrides AssignmentGroupsListAssignmentGroupsInclude = "overrides"
	// AssignmentGroupsListAssignmentGroupsIncludeSubmission enum value ("submission")
	AssignmentGroupsListAssignmentGroupsIncludeSubmission AssignmentGroupsListAssignmentGroupsInclude = "submission"
	// AssignmentGroupsListAssignmentGroupsIncludeObservedUsers enum value ("observed_users")
	AssignmentGroupsListAssignmentGroupsIncludeObservedUsers AssignmentGroupsListAssignmentGroupsInclude = "observed_users"
	// AssignmentGroupsListAssignmentGroupsIncludeCanEdit enum value ("can_edit")
	AssignmentGroupsListAssignmentGroupsIncludeCanEdit AssignmentGroupsListAssignmentGroupsInclude = "can_edit"
)

// AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypes enumeration
type AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypes string

const (
	// AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypesOnlineQuiz enum value ("online_quiz")
	AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypesOnlineQuiz AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypes = "online_quiz"
	// AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypesDiscussionTopic enum value ("discussion_topic")
	AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypesDiscussionTopic AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypes = "discussion_topic"
	// AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypesWikiPage enum value ("wiki_page")
	AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypesWikiPage AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypes = "wiki_page"
	// AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypesExternalTool enum value ("external_tool")
	AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypesExternalTool AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypes = "external_tool"
)

// AssignmentsListAssignmentsInclude enumeration
type AssignmentsListAssignmentsInclude string

const (
	// AssignmentsListAssignmentsIncludeSubmission enum value ("submission")
	AssignmentsListAssignmentsIncludeSubmission AssignmentsListAssignmentsInclude = "submission"
	// AssignmentsListAssignmentsIncludeAssignmentVisibility enum value ("assignment_visibility")
	AssignmentsListAssignmentsIncludeAssignmentVisibility AssignmentsListAssignmentsInclude = "assignment_visibility"
	// AssignmentsListAssignmentsIncludeAllDates enum value ("all_dates")
	AssignmentsListAssignmentsIncludeAllDates AssignmentsListAssignmentsInclude = "all_dates"
	// AssignmentsListAssignmentsIncludeOverrides enum value ("overrides")
	AssignmentsListAssignmentsIncludeOverrides AssignmentsListAssignmentsInclude = "overrides"
	// AssignmentsListAssignmentsIncludeObservedUsers enum value ("observed_users")
	AssignmentsListAssignmentsIncludeObservedUsers AssignmentsListAssignmentsInclude = "observed_users"
	// AssignmentsListAssignmentsIncludeCanEdit enum value ("can_edit")
	AssignmentsListAssignmentsIncludeCanEdit AssignmentsListAssignmentsInclude = "can_edit"
)

// AssignmentsListAssignmentsBucket enumeration
type AssignmentsListAssignmentsBucket string

const (
	// AssignmentsListAssignmentsBucketPast enum value ("past")
	AssignmentsListAssignmentsBucketPast AssignmentsListAssignmentsBucket = "past"
	// AssignmentsListAssignmentsBucketOverdue enum value ("overdue")
	AssignmentsListAssignmentsBucketOverdue AssignmentsListAssignmentsBucket = "overdue"
	// AssignmentsListAssignmentsBucketUndated enum value ("undated")
	AssignmentsListAssignmentsBucketUndated AssignmentsListAssignmentsBucket = "undated"
	// AssignmentsListAssignmentsBucketUngraded enum value ("ungraded")
	AssignmentsListAssignmentsBucketUngraded AssignmentsListAssignmentsBucket = "ungraded"
	// AssignmentsListAssignmentsBucketUnsubmitted enum value ("unsubmitted")
	AssignmentsListAssignmentsBucketUnsubmitted AssignmentsListAssignmentsBucket = "unsubmitted"
	// AssignmentsListAssignmentsBucketUpcoming enum value ("upcoming")
	AssignmentsListAssignmentsBucketUpcoming AssignmentsListAssignmentsBucket = "upcoming"
	// AssignmentsListAssignmentsBucketFuture enum value ("future")
	AssignmentsListAssignmentsBucketFuture AssignmentsListAssignmentsBucket = "future"
)

// AssignmentsListAssignmentsOrderBy enumeration
type AssignmentsListAssignmentsOrderBy string

const (
	// AssignmentsListAssignmentsOrderByPosition enum value ("position")
	AssignmentsListAssignmentsOrderByPosition AssignmentsListAssignmentsOrderBy = "position"
	// AssignmentsListAssignmentsOrderByName enum value ("name")
	AssignmentsListAssignmentsOrderByName AssignmentsListAssignmentsOrderBy = "name"
	// AssignmentsListAssignmentsOrderByDueAt enum value ("due_at")
	AssignmentsListAssignmentsOrderByDueAt AssignmentsListAssignmentsOrderBy = "due_at"
)

// AssignmentsGetASingleAssignmentInclude enumeration
type AssignmentsGetASingleAssignmentInclude string

const (
	// AssignmentsGetASingleAssignmentIncludeSubmission enum value ("submission")
	AssignmentsGetASingleAssignmentIncludeSubmission AssignmentsGetASingleAssignmentInclude = "submission"
	// AssignmentsGetASingleAssignmentIncludeAssignmentVisibility enum value ("assignment_visibility")
	AssignmentsGetASingleAssignmentIncludeAssignmentVisibility AssignmentsGetASingleAssignmentInclude = "assignment_visibility"
	// AssignmentsGetASingleAssignmentIncludeOverrides enum value ("overrides")
	AssignmentsGetASingleAssignmentIncludeOverrides AssignmentsGetASingleAssignmentInclude = "overrides"
	// AssignmentsGetASingleAssignmentIncludeObservedUsers enum value ("observed_users")
	AssignmentsGetASingleAssignmentIncludeObservedUsers AssignmentsGetASingleAssignmentInclude = "observed_users"
	// AssignmentsGetASingleAssignmentIncludeCanEdit enum value ("can_edit")
	AssignmentsGetASingleAssignmentIncludeCanEdit AssignmentsGetASingleAssignmentInclude = "can_edit"
)

// CollaborationsListMembersOfACollaborationInclude enumeration
type CollaborationsListMembersOfACollaborationInclude string

const (
	// CollaborationsListMembersOfACollaborationIncludeCollaboratorLtiID enum value ("collaborator_lti_id")
	CollaborationsListMembersOfACollaborationIncludeCollaboratorLtiID CollaborationsListMembersOfACollaborationInclude = "collaborator_lti_id"
	// CollaborationsListMembersOfACollaborationIncludeAvatarImageURL enum value ("avatar_image_url")
	CollaborationsListMembersOfACollaborationIncludeAvatarImageURL CollaborationsListMembersOfACollaborationInclude = "avatar_image_url"
)

// ContentExportsExportContentExportType enumeration
type ContentExportsExportContentExportType string

const (
	// ContentExportsExportContentExportTypeCommonCartridge enum value ("common_cartridge")
	ContentExportsExportContentExportTypeCommonCartridge ContentExportsExportContentExportType = "common_cartridge"
	// ContentExportsExportContentExportTypeQti enum value ("qti")
	ContentExportsExportContentExportTypeQti ContentExportsExportContentExportType = "qti"
	// ContentExportsExportContentExportTypeZip enum value ("zip")
	ContentExportsExportContentExportTypeZip ContentExportsExportContentExportType = "zip"
)

// ContentExportsExportContentSelect enumeration
type ContentExportsExportContentSelect string

const (
	// ContentExportsExportContentSelectFolders enum value ("folders")
	ContentExportsExportContentSelectFolders ContentExportsExportContentSelect = "folders"
	// ContentExportsExportContentSelectFiles enum value ("files")
	ContentExportsExportContentSelectFiles ContentExportsExportContentSelect = "files"
	// ContentExportsExportContentSelectAttachments enum value ("attachments")
	ContentExportsExportContentSelectAttachments ContentExportsExportContentSelect = "attachments"
	// ContentExportsExportContentSelectQuizzes enum value ("quizzes")
	ContentExportsExportContentSelectQuizzes ContentExportsExportContentSelect = "quizzes"
	// ContentExportsExportContentSelectAssignments enum value ("assignments")
	ContentExportsExportContentSelectAssignments ContentExportsExportContentSelect = "assignments"
	// ContentExportsExportContentSelectAnnouncements enum value ("announcements")
	ContentExportsExportContentSelectAnnouncements ContentExportsExportContentSelect = "announcements"
	// ContentExportsExportContentSelectCalendarEvents enum value ("calendar_events")
	ContentExportsExportContentSelectCalendarEvents ContentExportsExportContentSelect = "calendar_events"
	// ContentExportsExportContentSelectDiscussionTopics enum value ("discussion_topics")
	ContentExportsExportContentSelectDiscussionTopics ContentExportsExportContentSelect = "discussion_topics"
	// ContentExportsExportContentSelectModules enum value ("modules")
	ContentExportsExportContentSelectModules ContentExportsExportContentSelect = "modules"
	// ContentExportsExportContentSelectModuleItems enum value ("module_items")
	ContentExportsExportContentSelectModuleItems ContentExportsExportContentSelect = "module_items"
	// ContentExportsExportContentSelectPages enum value ("pages")
	ContentExportsExportContentSelectPages ContentExportsExportContentSelect = "pages"
	// ContentExportsExportContentSelectRubrics enum value ("rubrics")
	ContentExportsExportContentSelectRubrics ContentExportsExportContentSelect = "rubrics"
)

// CoursesCopyCourseContentExcept enumeration
type CoursesCopyCourseContentExcept string

const (
	// CoursesCopyCourseContentExceptCourseSettings enum value ("course_settings")
	CoursesCopyCourseContentExceptCourseSettings CoursesCopyCourseContentExcept = "course_settings"
	// CoursesCopyCourseContentExceptAssignments enum value ("assignments")
	CoursesCopyCourseContentExceptAssignments CoursesCopyCourseContentExcept = "assignments"
	// CoursesCopyCourseContentExceptExternalTools enum value ("external_tools")
	CoursesCopyCourseContentExceptExternalTools CoursesCopyCourseContentExcept = "external_tools"
	// CoursesCopyCourseContentExceptFiles enum value ("files")
	CoursesCopyCourseContentExceptFiles CoursesCopyCourseContentExcept = "files"
	// CoursesCopyCourseContentExceptTopics enum value ("topics")
	CoursesCopyCourseContentExceptTopics CoursesCopyCourseContentExcept = "topics"
	// CoursesCopyCourseContentExceptCalendarEvents enum value ("calendar_events")
	CoursesCopyCourseContentExceptCalendarEvents CoursesCopyCourseContentExcept = "calendar_events"
	// CoursesCopyCourseContentExceptQuizzes enum value ("quizzes")
	CoursesCopyCourseContentExceptQuizzes CoursesCopyCourseContentExcept = "quizzes"
	// CoursesCopyCourseContentExceptWikiPages enum value ("wiki_pages")
	CoursesCopyCourseContentExceptWikiPages CoursesCopyCourseContentExcept = "wiki_pages"
	// CoursesCopyCourseContentExceptModules enum value ("modules")
	CoursesCopyCourseContentExceptModules CoursesCopyCourseContentExcept = "modules"
	// CoursesCopyCourseContentExceptOutcomes enum value ("outcomes")
	CoursesCopyCourseContentExceptOutcomes CoursesCopyCourseContentExcept = "outcomes"
)

// CoursesCopyCourseContentOnly enumeration
type CoursesCopyCourseContentOnly string

const (
	// CoursesCopyCourseContentOnlyCourseSettings enum value ("course_settings")
	CoursesCopyCourseContentOnlyCourseSettings CoursesCopyCourseContentOnly = "course_settings"
	// CoursesCopyCourseContentOnlyAssignments enum value ("assignments")
	CoursesCopyCourseContentOnlyAssignments CoursesCopyCourseContentOnly = "assignments"
	// CoursesCopyCourseContentOnlyExternalTools enum value ("external_tools")
	CoursesCopyCourseContentOnlyExternalTools CoursesCopyCourseContentOnly = "external_tools"
	// CoursesCopyCourseContentOnlyFiles enum value ("files")
	CoursesCopyCourseContentOnlyFiles CoursesCopyCourseContentOnly = "files"
	// CoursesCopyCourseContentOnlyTopics enum value ("topics")
	CoursesCopyCourseContentOnlyTopics CoursesCopyCourseContentOnly = "topics"
	// CoursesCopyCourseContentOnlyCalendarEvents enum value ("calendar_events")
	CoursesCopyCourseContentOnlyCalendarEvents CoursesCopyCourseContentOnly = "calendar_events"
	// CoursesCopyCourseContentOnlyQuizzes enum value ("quizzes")
	CoursesCopyCourseContentOnlyQuizzes CoursesCopyCourseContentOnly = "quizzes"
	// CoursesCopyCourseContentOnlyWikiPages enum value ("wiki_pages")
	CoursesCopyCourseContentOnlyWikiPages CoursesCopyCourseContentOnly = "wiki_pages"
	// CoursesCopyCourseContentOnlyModules enum value ("modules")
	CoursesCopyCourseContentOnlyModules CoursesCopyCourseContentOnly = "modules"
	// CoursesCopyCourseContentOnlyOutcomes enum value ("outcomes")
	CoursesCopyCourseContentOnlyOutcomes CoursesCopyCourseContentOnly = "outcomes"
)

// ContentMigrationsCreateAContentMigrationSelect enumeration
type ContentMigrationsCreateAContentMigrationSelect string

const (
	// ContentMigrationsCreateAContentMigrationSelectFolders enum value ("folders")
	ContentMigrationsCreateAContentMigrationSelectFolders ContentMigrationsCreateAContentMigrationSelect = "folders"
	// ContentMigrationsCreateAContentMigrationSelectFiles enum value ("files")
	ContentMigrationsCreateAContentMigrationSelectFiles ContentMigrationsCreateAContentMigrationSelect = "files"
	// ContentMigrationsCreateAContentMigrationSelectAttachments enum value ("attachments")
	ContentMigrationsCreateAContentMigrationSelectAttachments ContentMigrationsCreateAContentMigrationSelect = "attachments"
	// ContentMigrationsCreateAContentMigrationSelectQuizzes enum value ("quizzes")
	ContentMigrationsCreateAContentMigrationSelectQuizzes ContentMigrationsCreateAContentMigrationSelect = "quizzes"
	// ContentMigrationsCreateAContentMigrationSelectAssignments enum value ("assignments")
	ContentMigrationsCreateAContentMigrationSelectAssignments ContentMigrationsCreateAContentMigrationSelect = "assignments"
	// ContentMigrationsCreateAContentMigrationSelectAnnouncements enum value ("announcements")
	ContentMigrationsCreateAContentMigrationSelectAnnouncements ContentMigrationsCreateAContentMigrationSelect = "announcements"
	// ContentMigrationsCreateAContentMigrationSelectCalendarEvents enum value ("calendar_events")
	ContentMigrationsCreateAContentMigrationSelectCalendarEvents ContentMigrationsCreateAContentMigrationSelect = "calendar_events"
	// ContentMigrationsCreateAContentMigrationSelectDiscussionTopics enum value ("discussion_topics")
	ContentMigrationsCreateAContentMigrationSelectDiscussionTopics ContentMigrationsCreateAContentMigrationSelect = "discussion_topics"
	// ContentMigrationsCreateAContentMigrationSelectModules enum value ("modules")
	ContentMigrationsCreateAContentMigrationSelectModules ContentMigrationsCreateAContentMigrationSelect = "modules"
	// ContentMigrationsCreateAContentMigrationSelectModuleItems enum value ("module_items")
	ContentMigrationsCreateAContentMigrationSelectModuleItems ContentMigrationsCreateAContentMigrationSelect = "module_items"
	// ContentMigrationsCreateAContentMigrationSelectPages enum value ("pages")
	ContentMigrationsCreateAContentMigrationSelectPages ContentMigrationsCreateAContentMigrationSelect = "pages"
	// ContentMigrationsCreateAContentMigrationSelectRubrics enum value ("rubrics")
	ContentMigrationsCreateAContentMigrationSelectRubrics ContentMigrationsCreateAContentMigrationSelect = "rubrics"
)

// ContentMigrationsListItemsForSelectiveImportType enumeration
type ContentMigrationsListItemsForSelectiveImportType string

const (
	// ContentMigrationsListItemsForSelectiveImportTypeContextModules enum value ("context_modules")
	ContentMigrationsListItemsForSelectiveImportTypeContextModules ContentMigrationsListItemsForSelectiveImportType = "context_modules"
	// ContentMigrationsListItemsForSelectiveImportTypeAssignments enum value ("assignments")
	ContentMigrationsListItemsForSelectiveImportTypeAssignments ContentMigrationsListItemsForSelectiveImportType = "assignments"
	// ContentMigrationsListItemsForSelectiveImportTypeQuizzes enum value ("quizzes")
	ContentMigrationsListItemsForSelectiveImportTypeQuizzes ContentMigrationsListItemsForSelectiveImportType = "quizzes"
	// ContentMigrationsListItemsForSelectiveImportTypeAssessmentQuestionBanks enum value ("assessment_question_banks")
	ContentMigrationsListItemsForSelectiveImportTypeAssessmentQuestionBanks ContentMigrationsListItemsForSelectiveImportType = "assessment_question_banks"
	// ContentMigrationsListItemsForSelectiveImportTypeDiscussionTopics enum value ("discussion_topics")
	ContentMigrationsListItemsForSelectiveImportTypeDiscussionTopics ContentMigrationsListItemsForSelectiveImportType = "discussion_topics"
	// ContentMigrationsListItemsForSelectiveImportTypeWikiPages enum value ("wiki_pages")
	ContentMigrationsListItemsForSelectiveImportTypeWikiPages ContentMigrationsListItemsForSelectiveImportType = "wiki_pages"
	// ContentMigrationsListItemsForSelectiveImportTypeContextExternalTools enum value ("context_external_tools")
	ContentMigrationsListItemsForSelectiveImportTypeContextExternalTools ContentMigrationsListItemsForSelectiveImportType = "context_external_tools"
	// ContentMigrationsListItemsForSelectiveImportTypeToolProfiles enum value ("tool_profiles")
	ContentMigrationsListItemsForSelectiveImportTypeToolProfiles ContentMigrationsListItemsForSelectiveImportType = "tool_profiles"
	// ContentMigrationsListItemsForSelectiveImportTypeAnnouncements enum value ("announcements")
	ContentMigrationsListItemsForSelectiveImportTypeAnnouncements ContentMigrationsListItemsForSelectiveImportType = "announcements"
	// ContentMigrationsListItemsForSelectiveImportTypeCalendarEvents enum value ("calendar_events")
	ContentMigrationsListItemsForSelectiveImportTypeCalendarEvents ContentMigrationsListItemsForSelectiveImportType = "calendar_events"
	// ContentMigrationsListItemsForSelectiveImportTypeRubrics enum value ("rubrics")
	ContentMigrationsListItemsForSelectiveImportTypeRubrics ContentMigrationsListItemsForSelectiveImportType = "rubrics"
	// ContentMigrationsListItemsForSelectiveImportTypeGroups enum value ("groups")
	ContentMigrationsListItemsForSelectiveImportTypeGroups ContentMigrationsListItemsForSelectiveImportType = "groups"
	// ContentMigrationsListItemsForSelectiveImportTypeLearningOutcomes enum value ("learning_outcomes")
	ContentMigrationsListItemsForSelectiveImportTypeLearningOutcomes ContentMigrationsListItemsForSelectiveImportType = "learning_outcomes"
	// ContentMigrationsListItemsForSelectiveImportTypeAttachments enum value ("attachments")
	ContentMigrationsListItemsForSelectiveImportTypeAttachments ContentMigrationsListItemsForSelectiveImportType = "attachments"
)

// ContentSharesCreateAContentShareContentType enumeration
type ContentSharesCreateAContentShareContentType string

const (
	// ContentSharesCreateAContentShareContentTypeAssignment enum value ("assignment")
	ContentSharesCreateAContentShareContentTypeAssignment ContentSharesCreateAContentShareContentType = "assignment"
	// ContentSharesCreateAContentShareContentTypeDiscussionTopic enum value ("discussion_topic")
	ContentSharesCreateAContentShareContentTypeDiscussionTopic ContentSharesCreateAContentShareContentType = "discussion_topic"
	// ContentSharesCreateAContentShareContentTypePage enum value ("page")
	ContentSharesCreateAContentShareContentTypePage ContentSharesCreateAContentShareContentType = "page"
	// ContentSharesCreateAContentShareContentTypeQuiz enum value ("quiz")
	ContentSharesCreateAContentShareContentTypeQuiz ContentSharesCreateAContentShareContentType = "quiz"
	// ContentSharesCreateAContentShareContentTypeModule enum value ("module")
	ContentSharesCreateAContentShareContentTypeModule ContentSharesCreateAContentShareContentType = "module"
	// ContentSharesCreateAContentShareContentTypeModuleItem enum value ("module_item")
	ContentSharesCreateAContentShareContentTypeModuleItem ContentSharesCreateAContentShareContentType = "module_item"
)

// ContentSharesUpdateAContentShareReadState enumeration
type ContentSharesUpdateAContentShareReadState string

const (
	// ContentSharesUpdateAContentShareReadStateRead enum value ("read")
	ContentSharesUpdateAContentShareReadStateRead ContentSharesUpdateAContentShareReadState = "read"
	// ContentSharesUpdateAContentShareReadStateUnread enum value ("unread")
	ContentSharesUpdateAContentShareReadStateUnread ContentSharesUpdateAContentShareReadState = "unread"
)

// ModulesListModuleItemsInclude enumeration
type ModulesListModuleItemsInclude string

const (
	// ModulesListModuleItemsIncludeContentDetails enum value ("content_details")
	ModulesListModuleItemsIncludeContentDetails ModulesListModuleItemsInclude = "content_details"
)

// ModulesShowModuleItemInclude enumeration
type ModulesShowModuleItemInclude string

const (
	// ModulesShowModuleItemIncludeContentDetails enum value ("content_details")
	ModulesShowModuleItemIncludeContentDetails ModulesShowModuleItemInclude = "content_details"
)

// ModulesGetModuleItemSequenceAssetType enumeration
type ModulesGetModuleItemSequenceAssetType string

const (
	// ModulesGetModuleItemSequenceAssetTypeModuleItem enum value ("ModuleItem")
	ModulesGetModuleItemSequenceAssetTypeModuleItem ModulesGetModuleItemSequenceAssetType = "ModuleItem"
	// ModulesGetModuleItemSequenceAssetTypeFile enum value ("File")
	ModulesGetModuleItemSequenceAssetTypeFile ModulesGetModuleItemSequenceAssetType = "File"
	// ModulesGetModuleItemSequenceAssetTypePage enum value ("Page")
	ModulesGetModuleItemSequenceAssetTypePage ModulesGetModuleItemSequenceAssetType = "Page"
	// ModulesGetModuleItemSequenceAssetTypeDiscussion enum value ("Discussion")
	ModulesGetModuleItemSequenceAssetTypeDiscussion ModulesGetModuleItemSequenceAssetType = "Discussion"
	// ModulesGetModuleItemSequenceAssetTypeAssignment enum value ("Assignment")
	ModulesGetModuleItemSequenceAssetTypeAssignment ModulesGetModuleItemSequenceAssetType = "Assignment"
	// ModulesGetModuleItemSequenceAssetTypeQuiz enum value ("Quiz")
	ModulesGetModuleItemSequenceAssetTypeQuiz ModulesGetModuleItemSequenceAssetType = "Quiz"
	// ModulesGetModuleItemSequenceAssetTypeExternalTool enum value ("ExternalTool")
	ModulesGetModuleItemSequenceAssetTypeExternalTool ModulesGetModuleItemSequenceAssetType = "ExternalTool"
)

// ModulesListModulesInclude enumeration
type ModulesListModulesInclude string

const (
	// ModulesListModulesIncludeItems enum value ("items")
	ModulesListModulesIncludeItems ModulesListModulesInclude = "items"
	// ModulesListModulesIncludeContentDetails enum value ("content_details")
	ModulesListModulesIncludeContentDetails ModulesListModulesInclude = "content_details"
)

// ModulesShowModuleInclude enumeration
type ModulesShowModuleInclude string

const (
	// ModulesShowModuleIncludeItems enum value ("items")
	ModulesShowModuleIncludeItems ModulesShowModuleInclude = "items"
	// ModulesShowModuleIncludeContentDetails enum value ("content_details")
	ModulesShowModuleIncludeContentDetails ModulesShowModuleInclude = "content_details"
)

// ConversationsListConversationsScope enumeration
type ConversationsListConversationsScope string

const (
	// ConversationsListConversationsScopeUnread enum value ("unread")
	ConversationsListConversationsScopeUnread ConversationsListConversationsScope = "unread"
	// ConversationsListConversationsScopeStarred enum value ("starred")
	ConversationsListConversationsScopeStarred ConversationsListConversationsScope = "starred"
	// ConversationsListConversationsScopeArchived enum value ("archived")
	ConversationsListConversationsScopeArchived ConversationsListConversationsScope = "archived"
)

// ConversationsListConversationsFilter enumeration
type ConversationsListConversationsFilter string

const (
	// ConversationsListConversationsFilterCourseID enum value ("course_id")
	ConversationsListConversationsFilterCourseID ConversationsListConversationsFilter = "course_id"
	// ConversationsListConversationsFilterGroupID enum value ("group_id")
	ConversationsListConversationsFilterGroupID ConversationsListConversationsFilter = "group_id"
	// ConversationsListConversationsFilterUserID enum value ("user_id")
	ConversationsListConversationsFilterUserID ConversationsListConversationsFilter = "user_id"
)

// ConversationsListConversationsFilterMode enumeration
type ConversationsListConversationsFilterMode string

const (
	// ConversationsListConversationsFilterModeAnd enum value ("and")
	ConversationsListConversationsFilterModeAnd ConversationsListConversationsFilterMode = "and"
	// ConversationsListConversationsFilterModeOr enum value ("or")
	ConversationsListConversationsFilterModeOr ConversationsListConversationsFilterMode = "or"
)

// ConversationsListConversationsInclude enumeration
type ConversationsListConversationsInclude string

const (
	// ConversationsListConversationsIncludeParticipantAvatars enum value ("participant_avatars")
	ConversationsListConversationsIncludeParticipantAvatars ConversationsListConversationsInclude = "participant_avatars"
)

// ConversationsCreateAConversationMediaCommentType enumeration
type ConversationsCreateAConversationMediaCommentType string

const (
	// ConversationsCreateAConversationMediaCommentTypeAudio enum value ("audio")
	ConversationsCreateAConversationMediaCommentTypeAudio ConversationsCreateAConversationMediaCommentType = "audio"
	// ConversationsCreateAConversationMediaCommentTypeVideo enum value ("video")
	ConversationsCreateAConversationMediaCommentTypeVideo ConversationsCreateAConversationMediaCommentType = "video"
)

// ConversationsCreateAConversationMode enumeration
type ConversationsCreateAConversationMode string

const (
	// ConversationsCreateAConversationModeSync enum value ("sync")
	ConversationsCreateAConversationModeSync ConversationsCreateAConversationMode = "sync"
	// ConversationsCreateAConversationModeAsync enum value ("async")
	ConversationsCreateAConversationModeAsync ConversationsCreateAConversationMode = "async"
)

// ConversationsCreateAConversationScope enumeration
type ConversationsCreateAConversationScope string

const (
	// ConversationsCreateAConversationScopeUnread enum value ("unread")
	ConversationsCreateAConversationScopeUnread ConversationsCreateAConversationScope = "unread"
	// ConversationsCreateAConversationScopeStarred enum value ("starred")
	ConversationsCreateAConversationScopeStarred ConversationsCreateAConversationScope = "starred"
	// ConversationsCreateAConversationScopeArchived enum value ("archived")
	ConversationsCreateAConversationScopeArchived ConversationsCreateAConversationScope = "archived"
)

// ConversationsCreateAConversationFilter enumeration
type ConversationsCreateAConversationFilter string

const (
	// ConversationsCreateAConversationFilterCourseID enum value ("course_id")
	ConversationsCreateAConversationFilterCourseID ConversationsCreateAConversationFilter = "course_id"
	// ConversationsCreateAConversationFilterGroupID enum value ("group_id")
	ConversationsCreateAConversationFilterGroupID ConversationsCreateAConversationFilter = "group_id"
	// ConversationsCreateAConversationFilterUserID enum value ("user_id")
	ConversationsCreateAConversationFilterUserID ConversationsCreateAConversationFilter = "user_id"
)

// ConversationsCreateAConversationFilterMode enumeration
type ConversationsCreateAConversationFilterMode string

const (
	// ConversationsCreateAConversationFilterModeAnd enum value ("and")
	ConversationsCreateAConversationFilterModeAnd ConversationsCreateAConversationFilterMode = "and"
	// ConversationsCreateAConversationFilterModeOr enum value ("or")
	ConversationsCreateAConversationFilterModeOr ConversationsCreateAConversationFilterMode = "or"
)

// ConversationsGetASingleConversationScope enumeration
type ConversationsGetASingleConversationScope string

const (
	// ConversationsGetASingleConversationScopeUnread enum value ("unread")
	ConversationsGetASingleConversationScopeUnread ConversationsGetASingleConversationScope = "unread"
	// ConversationsGetASingleConversationScopeStarred enum value ("starred")
	ConversationsGetASingleConversationScopeStarred ConversationsGetASingleConversationScope = "starred"
	// ConversationsGetASingleConversationScopeArchived enum value ("archived")
	ConversationsGetASingleConversationScopeArchived ConversationsGetASingleConversationScope = "archived"
)

// ConversationsGetASingleConversationFilter enumeration
type ConversationsGetASingleConversationFilter string

const (
	// ConversationsGetASingleConversationFilterCourseID enum value ("course_id")
	ConversationsGetASingleConversationFilterCourseID ConversationsGetASingleConversationFilter = "course_id"
	// ConversationsGetASingleConversationFilterGroupID enum value ("group_id")
	ConversationsGetASingleConversationFilterGroupID ConversationsGetASingleConversationFilter = "group_id"
	// ConversationsGetASingleConversationFilterUserID enum value ("user_id")
	ConversationsGetASingleConversationFilterUserID ConversationsGetASingleConversationFilter = "user_id"
)

// ConversationsGetASingleConversationFilterMode enumeration
type ConversationsGetASingleConversationFilterMode string

const (
	// ConversationsGetASingleConversationFilterModeAnd enum value ("and")
	ConversationsGetASingleConversationFilterModeAnd ConversationsGetASingleConversationFilterMode = "and"
	// ConversationsGetASingleConversationFilterModeOr enum value ("or")
	ConversationsGetASingleConversationFilterModeOr ConversationsGetASingleConversationFilterMode = "or"
)

// ConversationsEditAConversationConversation enumeration
type ConversationsEditAConversationConversation string

const (
	// ConversationsEditAConversationConversationRead enum value ("read")
	ConversationsEditAConversationConversationRead ConversationsEditAConversationConversation = "read"
	// ConversationsEditAConversationConversationUnread enum value ("unread")
	ConversationsEditAConversationConversationUnread ConversationsEditAConversationConversation = "unread"
	// ConversationsEditAConversationConversationArchived enum value ("archived")
	ConversationsEditAConversationConversationArchived ConversationsEditAConversationConversation = "archived"
)

// ConversationsEditAConversationScope enumeration
type ConversationsEditAConversationScope string

const (
	// ConversationsEditAConversationScopeUnread enum value ("unread")
	ConversationsEditAConversationScopeUnread ConversationsEditAConversationScope = "unread"
	// ConversationsEditAConversationScopeStarred enum value ("starred")
	ConversationsEditAConversationScopeStarred ConversationsEditAConversationScope = "starred"
	// ConversationsEditAConversationScopeArchived enum value ("archived")
	ConversationsEditAConversationScopeArchived ConversationsEditAConversationScope = "archived"
)

// ConversationsEditAConversationFilter enumeration
type ConversationsEditAConversationFilter string

const (
	// ConversationsEditAConversationFilterCourseID enum value ("course_id")
	ConversationsEditAConversationFilterCourseID ConversationsEditAConversationFilter = "course_id"
	// ConversationsEditAConversationFilterGroupID enum value ("group_id")
	ConversationsEditAConversationFilterGroupID ConversationsEditAConversationFilter = "group_id"
	// ConversationsEditAConversationFilterUserID enum value ("user_id")
	ConversationsEditAConversationFilterUserID ConversationsEditAConversationFilter = "user_id"
)

// ConversationsEditAConversationFilterMode enumeration
type ConversationsEditAConversationFilterMode string

const (
	// ConversationsEditAConversationFilterModeAnd enum value ("and")
	ConversationsEditAConversationFilterModeAnd ConversationsEditAConversationFilterMode = "and"
	// ConversationsEditAConversationFilterModeOr enum value ("or")
	ConversationsEditAConversationFilterModeOr ConversationsEditAConversationFilterMode = "or"
)

// ConversationsAddAMessageMediaCommentType enumeration
type ConversationsAddAMessageMediaCommentType string

const (
	// ConversationsAddAMessageMediaCommentTypeAudio enum value ("audio")
	ConversationsAddAMessageMediaCommentTypeAudio ConversationsAddAMessageMediaCommentType = "audio"
	// ConversationsAddAMessageMediaCommentTypeVideo enum value ("video")
	ConversationsAddAMessageMediaCommentTypeVideo ConversationsAddAMessageMediaCommentType = "video"
)

// ConversationsBatchUpdateConversationsEvent enumeration
type ConversationsBatchUpdateConversationsEvent string

const (
	// ConversationsBatchUpdateConversationsEventMarkAsRead enum value ("mark_as_read")
	ConversationsBatchUpdateConversationsEventMarkAsRead ConversationsBatchUpdateConversationsEvent = "mark_as_read"
	// ConversationsBatchUpdateConversationsEventMarkAsUnread enum value ("mark_as_unread")
	ConversationsBatchUpdateConversationsEventMarkAsUnread ConversationsBatchUpdateConversationsEvent = "mark_as_unread"
	// ConversationsBatchUpdateConversationsEventStar enum value ("star")
	ConversationsBatchUpdateConversationsEventStar ConversationsBatchUpdateConversationsEvent = "star"
	// ConversationsBatchUpdateConversationsEventUnstar enum value ("unstar")
	ConversationsBatchUpdateConversationsEventUnstar ConversationsBatchUpdateConversationsEvent = "unstar"
	// ConversationsBatchUpdateConversationsEventArchive enum value ("archive")
	ConversationsBatchUpdateConversationsEventArchive ConversationsBatchUpdateConversationsEvent = "archive"
	// ConversationsBatchUpdateConversationsEventDestroy enum value ("destroy")
	ConversationsBatchUpdateConversationsEventDestroy ConversationsBatchUpdateConversationsEvent = "destroy"
)

// CoursesListYourCoursesEnrollmentType enumeration
type CoursesListYourCoursesEnrollmentType string

const (
	// CoursesListYourCoursesEnrollmentTypeTeacher enum value ("teacher")
	CoursesListYourCoursesEnrollmentTypeTeacher CoursesListYourCoursesEnrollmentType = "teacher"
	// CoursesListYourCoursesEnrollmentTypeStudent enum value ("student")
	CoursesListYourCoursesEnrollmentTypeStudent CoursesListYourCoursesEnrollmentType = "student"
	// CoursesListYourCoursesEnrollmentTypeTa enum value ("ta")
	CoursesListYourCoursesEnrollmentTypeTa CoursesListYourCoursesEnrollmentType = "ta"
	// CoursesListYourCoursesEnrollmentTypeObserver enum value ("observer")
	CoursesListYourCoursesEnrollmentTypeObserver CoursesListYourCoursesEnrollmentType = "observer"
	// CoursesListYourCoursesEnrollmentTypeDesigner enum value ("designer")
	CoursesListYourCoursesEnrollmentTypeDesigner CoursesListYourCoursesEnrollmentType = "designer"
)

// CoursesListYourCoursesEnrollmentState enumeration
type CoursesListYourCoursesEnrollmentState string

const (
	// CoursesListYourCoursesEnrollmentStateActive enum value ("active")
	CoursesListYourCoursesEnrollmentStateActive CoursesListYourCoursesEnrollmentState = "active"
	// CoursesListYourCoursesEnrollmentStateInvitedOrPending enum value ("invited_or_pending")
	CoursesListYourCoursesEnrollmentStateInvitedOrPending CoursesListYourCoursesEnrollmentState = "invited_or_pending"
	// CoursesListYourCoursesEnrollmentStateCompleted enum value ("completed")
	CoursesListYourCoursesEnrollmentStateCompleted CoursesListYourCoursesEnrollmentState = "completed"
)

// CoursesListYourCoursesInclude enumeration
type CoursesListYourCoursesInclude string

const (
	// CoursesListYourCoursesIncludeNeedsGradingCount enum value ("needs_grading_count")
	CoursesListYourCoursesIncludeNeedsGradingCount CoursesListYourCoursesInclude = "needs_grading_count"
	// CoursesListYourCoursesIncludeSyllabusBody enum value ("syllabus_body")
	CoursesListYourCoursesIncludeSyllabusBody CoursesListYourCoursesInclude = "syllabus_body"
	// CoursesListYourCoursesIncludePublicDescription enum value ("public_description")
	CoursesListYourCoursesIncludePublicDescription CoursesListYourCoursesInclude = "public_description"
	// CoursesListYourCoursesIncludeTotalScores enum value ("total_scores")
	CoursesListYourCoursesIncludeTotalScores CoursesListYourCoursesInclude = "total_scores"
	// CoursesListYourCoursesIncludeCurrentGradingPeriodScores enum value ("current_grading_period_scores")
	CoursesListYourCoursesIncludeCurrentGradingPeriodScores CoursesListYourCoursesInclude = "current_grading_period_scores"
	// CoursesListYourCoursesIncludeTerm enum value ("term")
	CoursesListYourCoursesIncludeTerm CoursesListYourCoursesInclude = "term"
	// CoursesListYourCoursesIncludeAccount enum value ("account")
	CoursesListYourCoursesIncludeAccount CoursesListYourCoursesInclude = "account"
	// CoursesListYourCoursesIncludeCourseProgress enum value ("course_progress")
	CoursesListYourCoursesIncludeCourseProgress CoursesListYourCoursesInclude = "course_progress"
	// CoursesListYourCoursesIncludeSections enum value ("sections")
	CoursesListYourCoursesIncludeSections CoursesListYourCoursesInclude = "sections"
	// CoursesListYourCoursesIncludeStorageQuotaUsedMb enum value ("storage_quota_used_mb")
	CoursesListYourCoursesIncludeStorageQuotaUsedMb CoursesListYourCoursesInclude = "storage_quota_used_mb"
	// CoursesListYourCoursesIncludeTotalStudents enum value ("total_students")
	CoursesListYourCoursesIncludeTotalStudents CoursesListYourCoursesInclude = "total_students"
	// CoursesListYourCoursesIncludePassbackStatus enum value ("passback_status")
	CoursesListYourCoursesIncludePassbackStatus CoursesListYourCoursesInclude = "passback_status"
	// CoursesListYourCoursesIncludeFavorites enum value ("favorites")
	CoursesListYourCoursesIncludeFavorites CoursesListYourCoursesInclude = "favorites"
	// CoursesListYourCoursesIncludeTeachers enum value ("teachers")
	CoursesListYourCoursesIncludeTeachers CoursesListYourCoursesInclude = "teachers"
	// CoursesListYourCoursesIncludeObservedUsers enum value ("observed_users")
	CoursesListYourCoursesIncludeObservedUsers CoursesListYourCoursesInclude = "observed_users"
	// CoursesListYourCoursesIncludeCourseImage enum value ("course_image")
	CoursesListYourCoursesIncludeCourseImage CoursesListYourCoursesInclude = "course_image"
	// CoursesListYourCoursesIncludeConcluded enum value ("concluded")
	CoursesListYourCoursesIncludeConcluded CoursesListYourCoursesInclude = "concluded"
)

// CoursesListYourCoursesState enumeration
type CoursesListYourCoursesState string

const (
	// CoursesListYourCoursesStateUnpublished enum value ("unpublished")
	CoursesListYourCoursesStateUnpublished CoursesListYourCoursesState = "unpublished"
	// CoursesListYourCoursesStateAvailable enum value ("available")
	CoursesListYourCoursesStateAvailable CoursesListYourCoursesState = "available"
	// CoursesListYourCoursesStateCompleted enum value ("completed")
	CoursesListYourCoursesStateCompleted CoursesListYourCoursesState = "completed"
	// CoursesListYourCoursesStateDeleted enum value ("deleted")
	CoursesListYourCoursesStateDeleted CoursesListYourCoursesState = "deleted"
)

// CoursesListCoursesForAUserInclude enumeration
type CoursesListCoursesForAUserInclude string

const (
	// CoursesListCoursesForAUserIncludeNeedsGradingCount enum value ("needs_grading_count")
	CoursesListCoursesForAUserIncludeNeedsGradingCount CoursesListCoursesForAUserInclude = "needs_grading_count"
	// CoursesListCoursesForAUserIncludeSyllabusBody enum value ("syllabus_body")
	CoursesListCoursesForAUserIncludeSyllabusBody CoursesListCoursesForAUserInclude = "syllabus_body"
	// CoursesListCoursesForAUserIncludePublicDescription enum value ("public_description")
	CoursesListCoursesForAUserIncludePublicDescription CoursesListCoursesForAUserInclude = "public_description"
	// CoursesListCoursesForAUserIncludeTotalScores enum value ("total_scores")
	CoursesListCoursesForAUserIncludeTotalScores CoursesListCoursesForAUserInclude = "total_scores"
	// CoursesListCoursesForAUserIncludeCurrentGradingPeriodScores enum value ("current_grading_period_scores")
	CoursesListCoursesForAUserIncludeCurrentGradingPeriodScores CoursesListCoursesForAUserInclude = "current_grading_period_scores"
	// CoursesListCoursesForAUserIncludeTerm enum value ("term")
	CoursesListCoursesForAUserIncludeTerm CoursesListCoursesForAUserInclude = "term"
	// CoursesListCoursesForAUserIncludeAccount enum value ("account")
	CoursesListCoursesForAUserIncludeAccount CoursesListCoursesForAUserInclude = "account"
	// CoursesListCoursesForAUserIncludeCourseProgress enum value ("course_progress")
	CoursesListCoursesForAUserIncludeCourseProgress CoursesListCoursesForAUserInclude = "course_progress"
	// CoursesListCoursesForAUserIncludeSections enum value ("sections")
	CoursesListCoursesForAUserIncludeSections CoursesListCoursesForAUserInclude = "sections"
	// CoursesListCoursesForAUserIncludeStorageQuotaUsedMb enum value ("storage_quota_used_mb")
	CoursesListCoursesForAUserIncludeStorageQuotaUsedMb CoursesListCoursesForAUserInclude = "storage_quota_used_mb"
	// CoursesListCoursesForAUserIncludeTotalStudents enum value ("total_students")
	CoursesListCoursesForAUserIncludeTotalStudents CoursesListCoursesForAUserInclude = "total_students"
	// CoursesListCoursesForAUserIncludePassbackStatus enum value ("passback_status")
	CoursesListCoursesForAUserIncludePassbackStatus CoursesListCoursesForAUserInclude = "passback_status"
	// CoursesListCoursesForAUserIncludeFavorites enum value ("favorites")
	CoursesListCoursesForAUserIncludeFavorites CoursesListCoursesForAUserInclude = "favorites"
	// CoursesListCoursesForAUserIncludeTeachers enum value ("teachers")
	CoursesListCoursesForAUserIncludeTeachers CoursesListCoursesForAUserInclude = "teachers"
	// CoursesListCoursesForAUserIncludeObservedUsers enum value ("observed_users")
	CoursesListCoursesForAUserIncludeObservedUsers CoursesListCoursesForAUserInclude = "observed_users"
	// CoursesListCoursesForAUserIncludeCourseImage enum value ("course_image")
	CoursesListCoursesForAUserIncludeCourseImage CoursesListCoursesForAUserInclude = "course_image"
	// CoursesListCoursesForAUserIncludeConcluded enum value ("concluded")
	CoursesListCoursesForAUserIncludeConcluded CoursesListCoursesForAUserInclude = "concluded"
)

// CoursesListCoursesForAUserState enumeration
type CoursesListCoursesForAUserState string

const (
	// CoursesListCoursesForAUserStateUnpublished enum value ("unpublished")
	CoursesListCoursesForAUserStateUnpublished CoursesListCoursesForAUserState = "unpublished"
	// CoursesListCoursesForAUserStateAvailable enum value ("available")
	CoursesListCoursesForAUserStateAvailable CoursesListCoursesForAUserState = "available"
	// CoursesListCoursesForAUserStateCompleted enum value ("completed")
	CoursesListCoursesForAUserStateCompleted CoursesListCoursesForAUserState = "completed"
	// CoursesListCoursesForAUserStateDeleted enum value ("deleted")
	CoursesListCoursesForAUserStateDeleted CoursesListCoursesForAUserState = "deleted"
)

// CoursesListCoursesForAUserEnrollmentState enumeration
type CoursesListCoursesForAUserEnrollmentState string

const (
	// CoursesListCoursesForAUserEnrollmentStateActive enum value ("active")
	CoursesListCoursesForAUserEnrollmentStateActive CoursesListCoursesForAUserEnrollmentState = "active"
	// CoursesListCoursesForAUserEnrollmentStateInvitedOrPending enum value ("invited_or_pending")
	CoursesListCoursesForAUserEnrollmentStateInvitedOrPending CoursesListCoursesForAUserEnrollmentState = "invited_or_pending"
	// CoursesListCoursesForAUserEnrollmentStateCompleted enum value ("completed")
	CoursesListCoursesForAUserEnrollmentStateCompleted CoursesListCoursesForAUserEnrollmentState = "completed"
)

// CoursesListUsersInCourseSort enumeration
type CoursesListUsersInCourseSort string

const (
	// CoursesListUsersInCourseSortUsername enum value ("username")
	CoursesListUsersInCourseSortUsername CoursesListUsersInCourseSort = "username"
	// CoursesListUsersInCourseSortLastLogin enum value ("last_login")
	CoursesListUsersInCourseSortLastLogin CoursesListUsersInCourseSort = "last_login"
	// CoursesListUsersInCourseSortEmail enum value ("email")
	CoursesListUsersInCourseSortEmail CoursesListUsersInCourseSort = "email"
	// CoursesListUsersInCourseSortSisID enum value ("sis_id")
	CoursesListUsersInCourseSortSisID CoursesListUsersInCourseSort = "sis_id"
)

// CoursesListUsersInCourseEnrollmentType enumeration
type CoursesListUsersInCourseEnrollmentType string

const (
	// CoursesListUsersInCourseEnrollmentTypeTeacher enum value ("teacher")
	CoursesListUsersInCourseEnrollmentTypeTeacher CoursesListUsersInCourseEnrollmentType = "teacher"
	// CoursesListUsersInCourseEnrollmentTypeStudent enum value ("student")
	CoursesListUsersInCourseEnrollmentTypeStudent CoursesListUsersInCourseEnrollmentType = "student"
	// CoursesListUsersInCourseEnrollmentTypeStudentView enum value ("student_view")
	CoursesListUsersInCourseEnrollmentTypeStudentView CoursesListUsersInCourseEnrollmentType = "student_view"
	// CoursesListUsersInCourseEnrollmentTypeTa enum value ("ta")
	CoursesListUsersInCourseEnrollmentTypeTa CoursesListUsersInCourseEnrollmentType = "ta"
	// CoursesListUsersInCourseEnrollmentTypeObserver enum value ("observer")
	CoursesListUsersInCourseEnrollmentTypeObserver CoursesListUsersInCourseEnrollmentType = "observer"
	// CoursesListUsersInCourseEnrollmentTypeDesigner enum value ("designer")
	CoursesListUsersInCourseEnrollmentTypeDesigner CoursesListUsersInCourseEnrollmentType = "designer"
)

// CoursesListUsersInCourseInclude enumeration
type CoursesListUsersInCourseInclude string

const (
	// CoursesListUsersInCourseIncludeEnrollments enum value ("enrollments")
	CoursesListUsersInCourseIncludeEnrollments CoursesListUsersInCourseInclude = "enrollments"
	// CoursesListUsersInCourseIncludeLocked enum value ("locked")
	CoursesListUsersInCourseIncludeLocked CoursesListUsersInCourseInclude = "locked"
	// CoursesListUsersInCourseIncludeAvatarURL enum value ("avatar_url")
	CoursesListUsersInCourseIncludeAvatarURL CoursesListUsersInCourseInclude = "avatar_url"
	// CoursesListUsersInCourseIncludeTestStudent enum value ("test_student")
	CoursesListUsersInCourseIncludeTestStudent CoursesListUsersInCourseInclude = "test_student"
	// CoursesListUsersInCourseIncludeBio enum value ("bio")
	CoursesListUsersInCourseIncludeBio CoursesListUsersInCourseInclude = "bio"
	// CoursesListUsersInCourseIncludeCustomLinks enum value ("custom_links")
	CoursesListUsersInCourseIncludeCustomLinks CoursesListUsersInCourseInclude = "custom_links"
	// CoursesListUsersInCourseIncludeCurrentGradingPeriodScores enum value ("current_grading_period_scores")
	CoursesListUsersInCourseIncludeCurrentGradingPeriodScores CoursesListUsersInCourseInclude = "current_grading_period_scores"
	// CoursesListUsersInCourseIncludeUUID enum value ("uuid")
	CoursesListUsersInCourseIncludeUUID CoursesListUsersInCourseInclude = "uuid"
)

// CoursesListUsersInCourseEnrollmentState enumeration
type CoursesListUsersInCourseEnrollmentState string

const (
	// CoursesListUsersInCourseEnrollmentStateActive enum value ("active")
	CoursesListUsersInCourseEnrollmentStateActive CoursesListUsersInCourseEnrollmentState = "active"
	// CoursesListUsersInCourseEnrollmentStateInvited enum value ("invited")
	CoursesListUsersInCourseEnrollmentStateInvited CoursesListUsersInCourseEnrollmentState = "invited"
	// CoursesListUsersInCourseEnrollmentStateRejected enum value ("rejected")
	CoursesListUsersInCourseEnrollmentStateRejected CoursesListUsersInCourseEnrollmentState = "rejected"
	// CoursesListUsersInCourseEnrollmentStateCompleted enum value ("completed")
	CoursesListUsersInCourseEnrollmentStateCompleted CoursesListUsersInCourseEnrollmentState = "completed"
	// CoursesListUsersInCourseEnrollmentStateInactive enum value ("inactive")
	CoursesListUsersInCourseEnrollmentStateInactive CoursesListUsersInCourseEnrollmentState = "inactive"
)

// CoursesDeleteConcludeACourseEvent enumeration
type CoursesDeleteConcludeACourseEvent string

const (
	// CoursesDeleteConcludeACourseEventDelete enum value ("delete")
	CoursesDeleteConcludeACourseEventDelete CoursesDeleteConcludeACourseEvent = "delete"
	// CoursesDeleteConcludeACourseEventConclude enum value ("conclude")
	CoursesDeleteConcludeACourseEventConclude CoursesDeleteConcludeACourseEvent = "conclude"
)

// CoursesGetASingleCourseInclude enumeration
type CoursesGetASingleCourseInclude string

const (
	// CoursesGetASingleCourseIncludeNeedsGradingCount enum value ("needs_grading_count")
	CoursesGetASingleCourseIncludeNeedsGradingCount CoursesGetASingleCourseInclude = "needs_grading_count"
	// CoursesGetASingleCourseIncludeSyllabusBody enum value ("syllabus_body")
	CoursesGetASingleCourseIncludeSyllabusBody CoursesGetASingleCourseInclude = "syllabus_body"
	// CoursesGetASingleCourseIncludePublicDescription enum value ("public_description")
	CoursesGetASingleCourseIncludePublicDescription CoursesGetASingleCourseInclude = "public_description"
	// CoursesGetASingleCourseIncludeTotalScores enum value ("total_scores")
	CoursesGetASingleCourseIncludeTotalScores CoursesGetASingleCourseInclude = "total_scores"
	// CoursesGetASingleCourseIncludeCurrentGradingPeriodScores enum value ("current_grading_period_scores")
	CoursesGetASingleCourseIncludeCurrentGradingPeriodScores CoursesGetASingleCourseInclude = "current_grading_period_scores"
	// CoursesGetASingleCourseIncludeTerm enum value ("term")
	CoursesGetASingleCourseIncludeTerm CoursesGetASingleCourseInclude = "term"
	// CoursesGetASingleCourseIncludeAccount enum value ("account")
	CoursesGetASingleCourseIncludeAccount CoursesGetASingleCourseInclude = "account"
	// CoursesGetASingleCourseIncludeCourseProgress enum value ("course_progress")
	CoursesGetASingleCourseIncludeCourseProgress CoursesGetASingleCourseInclude = "course_progress"
	// CoursesGetASingleCourseIncludeSections enum value ("sections")
	CoursesGetASingleCourseIncludeSections CoursesGetASingleCourseInclude = "sections"
	// CoursesGetASingleCourseIncludeStorageQuotaUsedMb enum value ("storage_quota_used_mb")
	CoursesGetASingleCourseIncludeStorageQuotaUsedMb CoursesGetASingleCourseInclude = "storage_quota_used_mb"
	// CoursesGetASingleCourseIncludeTotalStudents enum value ("total_students")
	CoursesGetASingleCourseIncludeTotalStudents CoursesGetASingleCourseInclude = "total_students"
	// CoursesGetASingleCourseIncludePassbackStatus enum value ("passback_status")
	CoursesGetASingleCourseIncludePassbackStatus CoursesGetASingleCourseInclude = "passback_status"
	// CoursesGetASingleCourseIncludeFavorites enum value ("favorites")
	CoursesGetASingleCourseIncludeFavorites CoursesGetASingleCourseInclude = "favorites"
	// CoursesGetASingleCourseIncludeTeachers enum value ("teachers")
	CoursesGetASingleCourseIncludeTeachers CoursesGetASingleCourseInclude = "teachers"
	// CoursesGetASingleCourseIncludeAllCourses enum value ("all_courses")
	CoursesGetASingleCourseIncludeAllCourses CoursesGetASingleCourseInclude = "all_courses"
	// CoursesGetASingleCourseIncludePermissions enum value ("permissions")
	CoursesGetASingleCourseIncludePermissions CoursesGetASingleCourseInclude = "permissions"
	// CoursesGetASingleCourseIncludeObservedUsers enum value ("observed_users")
	CoursesGetASingleCourseIncludeObservedUsers CoursesGetASingleCourseInclude = "observed_users"
	// CoursesGetASingleCourseIncludeCourseImage enum value ("course_image")
	CoursesGetASingleCourseIncludeCourseImage CoursesGetASingleCourseInclude = "course_image"
	// CoursesGetASingleCourseIncludeConcluded enum value ("concluded")
	CoursesGetASingleCourseIncludeConcluded CoursesGetASingleCourseInclude = "concluded"
)

// CoursesUpdateCoursesEvent enumeration
type CoursesUpdateCoursesEvent string

const (
	// CoursesUpdateCoursesEventOffer enum value ("offer")
	CoursesUpdateCoursesEventOffer CoursesUpdateCoursesEvent = "offer"
	// CoursesUpdateCoursesEventConclude enum value ("conclude")
	CoursesUpdateCoursesEventConclude CoursesUpdateCoursesEvent = "conclude"
	// CoursesUpdateCoursesEventDelete enum value ("delete")
	CoursesUpdateCoursesEventDelete CoursesUpdateCoursesEvent = "delete"
	// CoursesUpdateCoursesEventUndelete enum value ("undelete")
	CoursesUpdateCoursesEventUndelete CoursesUpdateCoursesEvent = "undelete"
)

// ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatus enumeration
type ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatus string

const (
	// ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatusEnabled enum value ("enabled")
	ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatusEnabled ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatus = "enabled"
	// ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatusDisabled enum value ("disabled")
	ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatusDisabled ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatus = "disabled"
	// ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatusInherited enum value ("inherited")
	ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatusInherited ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatus = "inherited"
)

// DiscussionTopicsGetASingleTopicInclude enumeration
type DiscussionTopicsGetASingleTopicInclude string

const (
	// DiscussionTopicsGetASingleTopicIncludeOverrides enum value ("overrides")
	DiscussionTopicsGetASingleTopicIncludeOverrides DiscussionTopicsGetASingleTopicInclude = "overrides"
)

// DiscussionTopicsListDiscussionTopicsInclude enumeration
type DiscussionTopicsListDiscussionTopicsInclude string

const (
	// DiscussionTopicsListDiscussionTopicsIncludeOverrides enum value ("overrides")
	DiscussionTopicsListDiscussionTopicsIncludeOverrides DiscussionTopicsListDiscussionTopicsInclude = "overrides"
)

// DiscussionTopicsListDiscussionTopicsOrderBy enumeration
type DiscussionTopicsListDiscussionTopicsOrderBy string

const (
	// DiscussionTopicsListDiscussionTopicsOrderByPosition enum value ("position")
	DiscussionTopicsListDiscussionTopicsOrderByPosition DiscussionTopicsListDiscussionTopicsOrderBy = "position"
	// DiscussionTopicsListDiscussionTopicsOrderByRecentActivity enum value ("recent_activity")
	DiscussionTopicsListDiscussionTopicsOrderByRecentActivity DiscussionTopicsListDiscussionTopicsOrderBy = "recent_activity"
	// DiscussionTopicsListDiscussionTopicsOrderByTitle enum value ("title")
	DiscussionTopicsListDiscussionTopicsOrderByTitle DiscussionTopicsListDiscussionTopicsOrderBy = "title"
)

// DiscussionTopicsListDiscussionTopicsScope enumeration
type DiscussionTopicsListDiscussionTopicsScope string

const (
	// DiscussionTopicsListDiscussionTopicsScopeLocked enum value ("locked")
	DiscussionTopicsListDiscussionTopicsScopeLocked DiscussionTopicsListDiscussionTopicsScope = "locked"
	// DiscussionTopicsListDiscussionTopicsScopeUnlocked enum value ("unlocked")
	DiscussionTopicsListDiscussionTopicsScopeUnlocked DiscussionTopicsListDiscussionTopicsScope = "unlocked"
	// DiscussionTopicsListDiscussionTopicsScopePinned enum value ("pinned")
	DiscussionTopicsListDiscussionTopicsScopePinned DiscussionTopicsListDiscussionTopicsScope = "pinned"
	// DiscussionTopicsListDiscussionTopicsScopeUnpinned enum value ("unpinned")
	DiscussionTopicsListDiscussionTopicsScopeUnpinned DiscussionTopicsListDiscussionTopicsScope = "unpinned"
)

// DiscussionTopicsListDiscussionTopicsFilterBy enumeration
type DiscussionTopicsListDiscussionTopicsFilterBy string

const (
	// DiscussionTopicsListDiscussionTopicsFilterByUnread enum value ("unread")
	DiscussionTopicsListDiscussionTopicsFilterByUnread DiscussionTopicsListDiscussionTopicsFilterBy = "unread"
)

// DiscussionTopicsCreateANewDiscussionTopicDiscussionType enumeration
type DiscussionTopicsCreateANewDiscussionTopicDiscussionType string

const (
	// DiscussionTopicsCreateANewDiscussionTopicDiscussionTypeSideComment enum value ("side_comment")
	DiscussionTopicsCreateANewDiscussionTopicDiscussionTypeSideComment DiscussionTopicsCreateANewDiscussionTopicDiscussionType = "side_comment"
	// DiscussionTopicsCreateANewDiscussionTopicDiscussionTypeThreaded enum value ("threaded")
	DiscussionTopicsCreateANewDiscussionTopicDiscussionTypeThreaded DiscussionTopicsCreateANewDiscussionTopicDiscussionType = "threaded"
)

// DiscussionTopicsUpdateATopicDiscussionType enumeration
type DiscussionTopicsUpdateATopicDiscussionType string

const (
	// DiscussionTopicsUpdateATopicDiscussionTypeSideComment enum value ("side_comment")
	DiscussionTopicsUpdateATopicDiscussionTypeSideComment DiscussionTopicsUpdateATopicDiscussionType = "side_comment"
	// DiscussionTopicsUpdateATopicDiscussionTypeThreaded enum value ("threaded")
	DiscussionTopicsUpdateATopicDiscussionTypeThreaded DiscussionTopicsUpdateATopicDiscussionType = "threaded"
)

// EnrollmentsListEnrollmentsState enumeration
type EnrollmentsListEnrollmentsState string

const (
	// EnrollmentsListEnrollmentsStateActive enum value ("active")
	EnrollmentsListEnrollmentsStateActive EnrollmentsListEnrollmentsState = "active"
	// EnrollmentsListEnrollmentsStateInvited enum value ("invited")
	EnrollmentsListEnrollmentsStateInvited EnrollmentsListEnrollmentsState = "invited"
	// EnrollmentsListEnrollmentsStateCreationPending enum value ("creation_pending")
	EnrollmentsListEnrollmentsStateCreationPending EnrollmentsListEnrollmentsState = "creation_pending"
	// EnrollmentsListEnrollmentsStateDeleted enum value ("deleted")
	EnrollmentsListEnrollmentsStateDeleted EnrollmentsListEnrollmentsState = "deleted"
	// EnrollmentsListEnrollmentsStateRejected enum value ("rejected")
	EnrollmentsListEnrollmentsStateRejected EnrollmentsListEnrollmentsState = "rejected"
	// EnrollmentsListEnrollmentsStateCompleted enum value ("completed")
	EnrollmentsListEnrollmentsStateCompleted EnrollmentsListEnrollmentsState = "completed"
	// EnrollmentsListEnrollmentsStateInactive enum value ("inactive")
	EnrollmentsListEnrollmentsStateInactive EnrollmentsListEnrollmentsState = "inactive"
	// EnrollmentsListEnrollmentsStateCurrentAndInvited enum value ("current_and_invited")
	EnrollmentsListEnrollmentsStateCurrentAndInvited EnrollmentsListEnrollmentsState = "current_and_invited"
	// EnrollmentsListEnrollmentsStateCurrentAndFuture enum value ("current_and_future")
	EnrollmentsListEnrollmentsStateCurrentAndFuture EnrollmentsListEnrollmentsState = "current_and_future"
	// EnrollmentsListEnrollmentsStateCurrentAndConcluded enum value ("current_and_concluded")
	EnrollmentsListEnrollmentsStateCurrentAndConcluded EnrollmentsListEnrollmentsState = "current_and_concluded"
)

// EnrollmentsListEnrollmentsInclude enumeration
type EnrollmentsListEnrollmentsInclude string

const (
	// EnrollmentsListEnrollmentsIncludeAvatarURL enum value ("avatar_url")
	EnrollmentsListEnrollmentsIncludeAvatarURL EnrollmentsListEnrollmentsInclude = "avatar_url"
	// EnrollmentsListEnrollmentsIncludeGroupIds enum value ("group_ids")
	EnrollmentsListEnrollmentsIncludeGroupIds EnrollmentsListEnrollmentsInclude = "group_ids"
	// EnrollmentsListEnrollmentsIncludeLocked enum value ("locked")
	EnrollmentsListEnrollmentsIncludeLocked EnrollmentsListEnrollmentsInclude = "locked"
	// EnrollmentsListEnrollmentsIncludeObservedUsers enum value ("observed_users")
	EnrollmentsListEnrollmentsIncludeObservedUsers EnrollmentsListEnrollmentsInclude = "observed_users"
	// EnrollmentsListEnrollmentsIncludeCanBeRemoved enum value ("can_be_removed")
	EnrollmentsListEnrollmentsIncludeCanBeRemoved EnrollmentsListEnrollmentsInclude = "can_be_removed"
	// EnrollmentsListEnrollmentsIncludeUUID enum value ("uuid")
	EnrollmentsListEnrollmentsIncludeUUID EnrollmentsListEnrollmentsInclude = "uuid"
)

// EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTask enumeration
type EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTask string

const (
	// EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTaskConclude enum value ("conclude")
	EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTaskConclude EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTask = "conclude"
	// EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTaskDelete enum value ("delete")
	EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTaskDelete EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTask = "delete"
	// EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTaskInactivate enum value ("inactivate")
	EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTaskInactivate EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTask = "inactivate"
	// EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTaskDeactivate enum value ("deactivate")
	EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTaskDeactivate EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTask = "deactivate"
)

// AnnouncementExternalFeedsCreateAnExternalFeedVerbosity enumeration
type AnnouncementExternalFeedsCreateAnExternalFeedVerbosity string

const (
	// AnnouncementExternalFeedsCreateAnExternalFeedVerbosityFull enum value ("full")
	AnnouncementExternalFeedsCreateAnExternalFeedVerbosityFull AnnouncementExternalFeedsCreateAnExternalFeedVerbosity = "full"
	// AnnouncementExternalFeedsCreateAnExternalFeedVerbosityTruncate enum value ("truncate")
	AnnouncementExternalFeedsCreateAnExternalFeedVerbosityTruncate AnnouncementExternalFeedsCreateAnExternalFeedVerbosity = "truncate"
	// AnnouncementExternalFeedsCreateAnExternalFeedVerbosityLinkOnly enum value ("link_only")
	AnnouncementExternalFeedsCreateAnExternalFeedVerbosityLinkOnly AnnouncementExternalFeedsCreateAnExternalFeedVerbosity = "link_only"
)

// ExternalToolsGetASessionlessLaunchURLForAnExternalToolLaunchType enumeration
type ExternalToolsGetASessionlessLaunchURLForAnExternalToolLaunchType string

const (
	// ExternalToolsGetASessionlessLaunchURLForAnExternalToolLaunchTypeAssessment enum value ("assessment")
	ExternalToolsGetASessionlessLaunchURLForAnExternalToolLaunchTypeAssessment ExternalToolsGetASessionlessLaunchURLForAnExternalToolLaunchType = "assessment"
	// ExternalToolsGetASessionlessLaunchURLForAnExternalToolLaunchTypeModuleItem enum value ("module_item")
	ExternalToolsGetASessionlessLaunchURLForAnExternalToolLaunchTypeModuleItem ExternalToolsGetASessionlessLaunchURLForAnExternalToolLaunchType = "module_item"
)

// ExternalToolsCreateAnExternalToolPrivacyLevel enumeration
type ExternalToolsCreateAnExternalToolPrivacyLevel string

const (
	// ExternalToolsCreateAnExternalToolPrivacyLevelAnonymous enum value ("anonymous")
	ExternalToolsCreateAnExternalToolPrivacyLevelAnonymous ExternalToolsCreateAnExternalToolPrivacyLevel = "anonymous"
	// ExternalToolsCreateAnExternalToolPrivacyLevelNameOnly enum value ("name_only")
	ExternalToolsCreateAnExternalToolPrivacyLevelNameOnly ExternalToolsCreateAnExternalToolPrivacyLevel = "name_only"
	// ExternalToolsCreateAnExternalToolPrivacyLevelPublic enum value ("public")
	ExternalToolsCreateAnExternalToolPrivacyLevelPublic ExternalToolsCreateAnExternalToolPrivacyLevel = "public"
)

// FeatureFlagsSetFeatureFlagState enumeration
type FeatureFlagsSetFeatureFlagState string

const (
	// FeatureFlagsSetFeatureFlagStateOff enum value ("off")
	FeatureFlagsSetFeatureFlagStateOff FeatureFlagsSetFeatureFlagState = "off"
	// FeatureFlagsSetFeatureFlagStateAllowed enum value ("allowed")
	FeatureFlagsSetFeatureFlagStateAllowed FeatureFlagsSetFeatureFlagState = "allowed"
	// FeatureFlagsSetFeatureFlagStateOn enum value ("on")
	FeatureFlagsSetFeatureFlagStateOn FeatureFlagsSetFeatureFlagState = "on"
)

// FilesListFilesInclude enumeration
type FilesListFilesInclude string

const (
	// FilesListFilesIncludeUser enum value ("user")
	FilesListFilesIncludeUser FilesListFilesInclude = "user"
)

// FilesListFilesSort enumeration
type FilesListFilesSort string

const (
	// FilesListFilesSortName enum value ("name")
	FilesListFilesSortName FilesListFilesSort = "name"
	// FilesListFilesSortSize enum value ("size")
	FilesListFilesSortSize FilesListFilesSort = "size"
	// FilesListFilesSortCreatedAt enum value ("created_at")
	FilesListFilesSortCreatedAt FilesListFilesSort = "created_at"
	// FilesListFilesSortUpdatedAt enum value ("updated_at")
	FilesListFilesSortUpdatedAt FilesListFilesSort = "updated_at"
	// FilesListFilesSortContentType enum value ("content_type")
	FilesListFilesSortContentType FilesListFilesSort = "content_type"
	// FilesListFilesSortUser enum value ("user")
	FilesListFilesSortUser FilesListFilesSort = "user"
)

// FilesListFilesOrder enumeration
type FilesListFilesOrder string

const (
	// FilesListFilesOrderAsc enum value ("asc")
	FilesListFilesOrderAsc FilesListFilesOrder = "asc"
	// FilesListFilesOrderDesc enum value ("desc")
	FilesListFilesOrderDesc FilesListFilesOrder = "desc"
)

// FilesGetFileInclude enumeration
type FilesGetFileInclude string

const (
	// FilesGetFileIncludeUser enum value ("user")
	FilesGetFileIncludeUser FilesGetFileInclude = "user"
)

// FilesUpdateFileOnDuplicate enumeration
type FilesUpdateFileOnDuplicate string

const (
	// FilesUpdateFileOnDuplicateOverwrite enum value ("overwrite")
	FilesUpdateFileOnDuplicateOverwrite FilesUpdateFileOnDuplicate = "overwrite"
	// FilesUpdateFileOnDuplicateRename enum value ("rename")
	FilesUpdateFileOnDuplicateRename FilesUpdateFileOnDuplicate = "rename"
)

// FilesCopyAFileOnDuplicate enumeration
type FilesCopyAFileOnDuplicate string

const (
	// FilesCopyAFileOnDuplicateOverwrite enum value ("overwrite")
	FilesCopyAFileOnDuplicateOverwrite FilesCopyAFileOnDuplicate = "overwrite"
	// FilesCopyAFileOnDuplicateRename enum value ("rename")
	FilesCopyAFileOnDuplicateRename FilesCopyAFileOnDuplicate = "rename"
)

// GroupCategoriesCreateAGroupCategorySelfSignup enumeration
type GroupCategoriesCreateAGroupCategorySelfSignup string

const (
	// GroupCategoriesCreateAGroupCategorySelfSignupEnabled enum value ("enabled")
	GroupCategoriesCreateAGroupCategorySelfSignupEnabled GroupCategoriesCreateAGroupCategorySelfSignup = "enabled"
	// GroupCategoriesCreateAGroupCategorySelfSignupRestricted enum value ("restricted")
	GroupCategoriesCreateAGroupCategorySelfSignupRestricted GroupCategoriesCreateAGroupCategorySelfSignup = "restricted"
)

// GroupCategoriesCreateAGroupCategoryAutoLeader enumeration
type GroupCategoriesCreateAGroupCategoryAutoLeader string

const (
	// GroupCategoriesCreateAGroupCategoryAutoLeaderFirst enum value ("first")
	GroupCategoriesCreateAGroupCategoryAutoLeaderFirst GroupCategoriesCreateAGroupCategoryAutoLeader = "first"
	// GroupCategoriesCreateAGroupCategoryAutoLeaderRandom enum value ("random")
	GroupCategoriesCreateAGroupCategoryAutoLeaderRandom GroupCategoriesCreateAGroupCategoryAutoLeader = "random"
)

// GroupCategoriesUpdateAGroupCategorySelfSignup enumeration
type GroupCategoriesUpdateAGroupCategorySelfSignup string

const (
	// GroupCategoriesUpdateAGroupCategorySelfSignupEnabled enum value ("enabled")
	GroupCategoriesUpdateAGroupCategorySelfSignupEnabled GroupCategoriesUpdateAGroupCategorySelfSignup = "enabled"
	// GroupCategoriesUpdateAGroupCategorySelfSignupRestricted enum value ("restricted")
	GroupCategoriesUpdateAGroupCategorySelfSignupRestricted GroupCategoriesUpdateAGroupCategorySelfSignup = "restricted"
)

// GroupCategoriesUpdateAGroupCategoryAutoLeader enumeration
type GroupCategoriesUpdateAGroupCategoryAutoLeader string

const (
	// GroupCategoriesUpdateAGroupCategoryAutoLeaderFirst enum value ("first")
	GroupCategoriesUpdateAGroupCategoryAutoLeaderFirst GroupCategoriesUpdateAGroupCategoryAutoLeader = "first"
	// GroupCategoriesUpdateAGroupCategoryAutoLeaderRandom enum value ("random")
	GroupCategoriesUpdateAGroupCategoryAutoLeaderRandom GroupCategoriesUpdateAGroupCategoryAutoLeader = "random"
)

// GroupsListGroupMembershipsFilterStates enumeration
type GroupsListGroupMembershipsFilterStates string

const (
	// GroupsListGroupMembershipsFilterStatesAccepted enum value ("accepted")
	GroupsListGroupMembershipsFilterStatesAccepted GroupsListGroupMembershipsFilterStates = "accepted"
	// GroupsListGroupMembershipsFilterStatesInvited enum value ("invited")
	GroupsListGroupMembershipsFilterStatesInvited GroupsListGroupMembershipsFilterStates = "invited"
	// GroupsListGroupMembershipsFilterStatesRequested enum value ("requested")
	GroupsListGroupMembershipsFilterStatesRequested GroupsListGroupMembershipsFilterStates = "requested"
)

// GroupsUpdateAMembershipWorkflowState enumeration
type GroupsUpdateAMembershipWorkflowState string

const (
	// GroupsUpdateAMembershipWorkflowStateAccepted enum value ("accepted")
	GroupsUpdateAMembershipWorkflowStateAccepted GroupsUpdateAMembershipWorkflowState = "accepted"
)

// GroupsListYourGroupsContextType enumeration
type GroupsListYourGroupsContextType string

const (
	// GroupsListYourGroupsContextTypeAccount enum value ("Account")
	GroupsListYourGroupsContextTypeAccount GroupsListYourGroupsContextType = "Account"
	// GroupsListYourGroupsContextTypeCourse enum value ("Course")
	GroupsListYourGroupsContextTypeCourse GroupsListYourGroupsContextType = "Course"
)

// GroupsListYourGroupsInclude enumeration
type GroupsListYourGroupsInclude string

const (
	// GroupsListYourGroupsIncludeTabs enum value ("tabs")
	GroupsListYourGroupsIncludeTabs GroupsListYourGroupsInclude = "tabs"
)

// GroupsListTheGroupsAvailableInAContextInclude enumeration
type GroupsListTheGroupsAvailableInAContextInclude string

const (
	// GroupsListTheGroupsAvailableInAContextIncludeTabs enum value ("tabs")
	GroupsListTheGroupsAvailableInAContextIncludeTabs GroupsListTheGroupsAvailableInAContextInclude = "tabs"
)

// GroupsGetASingleGroupInclude enumeration
type GroupsGetASingleGroupInclude string

const (
	// GroupsGetASingleGroupIncludeTabs enum value ("tabs")
	GroupsGetASingleGroupIncludeTabs GroupsGetASingleGroupInclude = "tabs"
)

// GroupsCreateAGroupJoinLevel enumeration
type GroupsCreateAGroupJoinLevel string

const (
	// GroupsCreateAGroupJoinLevelParentContextAutoJoin enum value ("parent_context_auto_join")
	GroupsCreateAGroupJoinLevelParentContextAutoJoin GroupsCreateAGroupJoinLevel = "parent_context_auto_join"
	// GroupsCreateAGroupJoinLevelParentContextRequest enum value ("parent_context_request")
	GroupsCreateAGroupJoinLevelParentContextRequest GroupsCreateAGroupJoinLevel = "parent_context_request"
	// GroupsCreateAGroupJoinLevelInvitationOnly enum value ("invitation_only")
	GroupsCreateAGroupJoinLevelInvitationOnly GroupsCreateAGroupJoinLevel = "invitation_only"
)

// GroupsEditAGroupJoinLevel enumeration
type GroupsEditAGroupJoinLevel string

const (
	// GroupsEditAGroupJoinLevelParentContextAutoJoin enum value ("parent_context_auto_join")
	GroupsEditAGroupJoinLevelParentContextAutoJoin GroupsEditAGroupJoinLevel = "parent_context_auto_join"
	// GroupsEditAGroupJoinLevelParentContextRequest enum value ("parent_context_request")
	GroupsEditAGroupJoinLevelParentContextRequest GroupsEditAGroupJoinLevel = "parent_context_request"
	// GroupsEditAGroupJoinLevelInvitationOnly enum value ("invitation_only")
	GroupsEditAGroupJoinLevelInvitationOnly GroupsEditAGroupJoinLevel = "invitation_only"
)

// GroupsListGroupSUsersInclude enumeration
type GroupsListGroupSUsersInclude string

const (
	// GroupsListGroupSUsersIncludeAvatarURL enum value ("avatar_url")
	GroupsListGroupSUsersIncludeAvatarURL GroupsListGroupSUsersInclude = "avatar_url"
)

// MediaObjectsListMediaObjectsSort enumeration
type MediaObjectsListMediaObjectsSort string

const (
	// MediaObjectsListMediaObjectsSortTitle enum value ("title")
	MediaObjectsListMediaObjectsSortTitle MediaObjectsListMediaObjectsSort = "title"
	// MediaObjectsListMediaObjectsSortCreatedAt enum value ("created_at")
	MediaObjectsListMediaObjectsSortCreatedAt MediaObjectsListMediaObjectsSort = "created_at"
)

// MediaObjectsListMediaObjectsOrder enumeration
type MediaObjectsListMediaObjectsOrder string

const (
	// MediaObjectsListMediaObjectsOrderAsc enum value ("asc")
	MediaObjectsListMediaObjectsOrderAsc MediaObjectsListMediaObjectsOrder = "asc"
	// MediaObjectsListMediaObjectsOrderDesc enum value ("desc")
	MediaObjectsListMediaObjectsOrderDesc MediaObjectsListMediaObjectsOrder = "desc"
)

// MediaObjectsListMediaObjectsExclude enumeration
type MediaObjectsListMediaObjectsExclude string

const (
	// MediaObjectsListMediaObjectsExcludeSources enum value ("sources")
	MediaObjectsListMediaObjectsExcludeSources MediaObjectsListMediaObjectsExclude = "sources"
	// MediaObjectsListMediaObjectsExcludeTracks enum value ("tracks")
	MediaObjectsListMediaObjectsExcludeTracks MediaObjectsListMediaObjectsExclude = "tracks"
)

// MediaObjectsListMediaTracksForAMediaObjectInclude enumeration
type MediaObjectsListMediaTracksForAMediaObjectInclude string

const (
	// MediaObjectsListMediaTracksForAMediaObjectIncludeContent enum value ("content")
	MediaObjectsListMediaTracksForAMediaObjectIncludeContent MediaObjectsListMediaTracksForAMediaObjectInclude = "content"
	// MediaObjectsListMediaTracksForAMediaObjectIncludeWebvttContent enum value ("webvtt_content")
	MediaObjectsListMediaTracksForAMediaObjectIncludeWebvttContent MediaObjectsListMediaTracksForAMediaObjectInclude = "webvtt_content"
	// MediaObjectsListMediaTracksForAMediaObjectIncludeUpdatedAt enum value ("updated_at")
	MediaObjectsListMediaTracksForAMediaObjectIncludeUpdatedAt MediaObjectsListMediaTracksForAMediaObjectInclude = "updated_at"
	// MediaObjectsListMediaTracksForAMediaObjectIncludeCreatedAt enum value ("created_at")
	MediaObjectsListMediaTracksForAMediaObjectIncludeCreatedAt MediaObjectsListMediaTracksForAMediaObjectInclude = "created_at"
)

// ContentMigrationsUpdateAMigrationIssueWorkflowState enumeration
type ContentMigrationsUpdateAMigrationIssueWorkflowState string

const (
	// ContentMigrationsUpdateAMigrationIssueWorkflowStateActive enum value ("active")
	ContentMigrationsUpdateAMigrationIssueWorkflowStateActive ContentMigrationsUpdateAMigrationIssueWorkflowState = "active"
	// ContentMigrationsUpdateAMigrationIssueWorkflowStateResolved enum value ("resolved")
	ContentMigrationsUpdateAMigrationIssueWorkflowStateResolved ContentMigrationsUpdateAMigrationIssueWorkflowState = "resolved"
)

// OutcomeGroupsCreateLinkAnOutcomeCalculationMethod enumeration
type OutcomeGroupsCreateLinkAnOutcomeCalculationMethod string

const (
	// OutcomeGroupsCreateLinkAnOutcomeCalculationMethodDecayingAverage enum value ("decaying_average")
	OutcomeGroupsCreateLinkAnOutcomeCalculationMethodDecayingAverage OutcomeGroupsCreateLinkAnOutcomeCalculationMethod = "decaying_average"
	// OutcomeGroupsCreateLinkAnOutcomeCalculationMethodNMastery enum value ("n_mastery")
	OutcomeGroupsCreateLinkAnOutcomeCalculationMethodNMastery OutcomeGroupsCreateLinkAnOutcomeCalculationMethod = "n_mastery"
	// OutcomeGroupsCreateLinkAnOutcomeCalculationMethodLatest enum value ("latest")
	OutcomeGroupsCreateLinkAnOutcomeCalculationMethodLatest OutcomeGroupsCreateLinkAnOutcomeCalculationMethod = "latest"
	// OutcomeGroupsCreateLinkAnOutcomeCalculationMethodHighest enum value ("highest")
	OutcomeGroupsCreateLinkAnOutcomeCalculationMethodHighest OutcomeGroupsCreateLinkAnOutcomeCalculationMethod = "highest"
)

// OutcomeResultsGetOutcomeResultsInclude enumeration
type OutcomeResultsGetOutcomeResultsInclude string

const (
	// OutcomeResultsGetOutcomeResultsIncludeAlignments enum value ("alignments")
	OutcomeResultsGetOutcomeResultsIncludeAlignments OutcomeResultsGetOutcomeResultsInclude = "alignments"
	// OutcomeResultsGetOutcomeResultsIncludeOutcomes enum value ("outcomes")
	OutcomeResultsGetOutcomeResultsIncludeOutcomes OutcomeResultsGetOutcomeResultsInclude = "outcomes"
	// OutcomeResultsGetOutcomeResultsIncludeOutcomesAlignments enum value ("outcomes.alignments")
	OutcomeResultsGetOutcomeResultsIncludeOutcomesAlignments OutcomeResultsGetOutcomeResultsInclude = "outcomes.alignments"
	// OutcomeResultsGetOutcomeResultsIncludeOutcomeGroups enum value ("outcome_groups")
	OutcomeResultsGetOutcomeResultsIncludeOutcomeGroups OutcomeResultsGetOutcomeResultsInclude = "outcome_groups"
	// OutcomeResultsGetOutcomeResultsIncludeOutcomeLinks enum value ("outcome_links")
	OutcomeResultsGetOutcomeResultsIncludeOutcomeLinks OutcomeResultsGetOutcomeResultsInclude = "outcome_links"
	// OutcomeResultsGetOutcomeResultsIncludeOutcomePaths enum value ("outcome_paths")
	OutcomeResultsGetOutcomeResultsIncludeOutcomePaths OutcomeResultsGetOutcomeResultsInclude = "outcome_paths"
	// OutcomeResultsGetOutcomeResultsIncludeUsers enum value ("users")
	OutcomeResultsGetOutcomeResultsIncludeUsers OutcomeResultsGetOutcomeResultsInclude = "users"
)

// OutcomeResultsGetOutcomeResultRollupsAggregate enumeration
type OutcomeResultsGetOutcomeResultRollupsAggregate string

const (
	// OutcomeResultsGetOutcomeResultRollupsAggregateCourse enum value ("course")
	OutcomeResultsGetOutcomeResultRollupsAggregateCourse OutcomeResultsGetOutcomeResultRollupsAggregate = "course"
)

// OutcomeResultsGetOutcomeResultRollupsAggregateStat enumeration
type OutcomeResultsGetOutcomeResultRollupsAggregateStat string

const (
	// OutcomeResultsGetOutcomeResultRollupsAggregateStatMean enum value ("mean")
	OutcomeResultsGetOutcomeResultRollupsAggregateStatMean OutcomeResultsGetOutcomeResultRollupsAggregateStat = "mean"
	// OutcomeResultsGetOutcomeResultRollupsAggregateStatMedian enum value ("median")
	OutcomeResultsGetOutcomeResultRollupsAggregateStatMedian OutcomeResultsGetOutcomeResultRollupsAggregateStat = "median"
)

// OutcomeResultsGetOutcomeResultRollupsInclude enumeration
type OutcomeResultsGetOutcomeResultRollupsInclude string

const (
	// OutcomeResultsGetOutcomeResultRollupsIncludeCourses enum value ("courses")
	OutcomeResultsGetOutcomeResultRollupsIncludeCourses OutcomeResultsGetOutcomeResultRollupsInclude = "courses"
	// OutcomeResultsGetOutcomeResultRollupsIncludeOutcomes enum value ("outcomes")
	OutcomeResultsGetOutcomeResultRollupsIncludeOutcomes OutcomeResultsGetOutcomeResultRollupsInclude = "outcomes"
	// OutcomeResultsGetOutcomeResultRollupsIncludeOutcomesAlignments enum value ("outcomes.alignments")
	OutcomeResultsGetOutcomeResultRollupsIncludeOutcomesAlignments OutcomeResultsGetOutcomeResultRollupsInclude = "outcomes.alignments"
	// OutcomeResultsGetOutcomeResultRollupsIncludeOutcomeGroups enum value ("outcome_groups")
	OutcomeResultsGetOutcomeResultRollupsIncludeOutcomeGroups OutcomeResultsGetOutcomeResultRollupsInclude = "outcome_groups"
	// OutcomeResultsGetOutcomeResultRollupsIncludeOutcomeLinks enum value ("outcome_links")
	OutcomeResultsGetOutcomeResultRollupsIncludeOutcomeLinks OutcomeResultsGetOutcomeResultRollupsInclude = "outcome_links"
	// OutcomeResultsGetOutcomeResultRollupsIncludeOutcomePaths enum value ("outcome_paths")
	OutcomeResultsGetOutcomeResultRollupsIncludeOutcomePaths OutcomeResultsGetOutcomeResultRollupsInclude = "outcome_paths"
	// OutcomeResultsGetOutcomeResultRollupsIncludeUsers enum value ("users")
	OutcomeResultsGetOutcomeResultRollupsIncludeUsers OutcomeResultsGetOutcomeResultRollupsInclude = "users"
)

// OutcomeResultsGetOutcomeResultRollupsExclude enumeration
type OutcomeResultsGetOutcomeResultRollupsExclude string

const (
	// OutcomeResultsGetOutcomeResultRollupsExcludeMissingUserRollups enum value ("missing_user_rollups")
	OutcomeResultsGetOutcomeResultRollupsExcludeMissingUserRollups OutcomeResultsGetOutcomeResultRollupsExclude = "missing_user_rollups"
)

// OutcomeResultsGetOutcomeResultRollupsSortBy enumeration
type OutcomeResultsGetOutcomeResultRollupsSortBy string

const (
	// OutcomeResultsGetOutcomeResultRollupsSortByStudent enum value ("student")
	OutcomeResultsGetOutcomeResultRollupsSortByStudent OutcomeResultsGetOutcomeResultRollupsSortBy = "student"
	// OutcomeResultsGetOutcomeResultRollupsSortByOutcome enum value ("outcome")
	OutcomeResultsGetOutcomeResultRollupsSortByOutcome OutcomeResultsGetOutcomeResultRollupsSortBy = "outcome"
)

// OutcomeResultsGetOutcomeResultRollupsSortOrder enumeration
type OutcomeResultsGetOutcomeResultRollupsSortOrder string

const (
	// OutcomeResultsGetOutcomeResultRollupsSortOrderDesc enum value ("desc")
	OutcomeResultsGetOutcomeResultRollupsSortOrderDesc OutcomeResultsGetOutcomeResultRollupsSortOrder = "desc"
)

// OutcomesUpdateAnOutcomeCalculationMethod enumeration
type OutcomesUpdateAnOutcomeCalculationMethod string

const (
	// OutcomesUpdateAnOutcomeCalculationMethodDecayingAverage enum value ("decaying_average")
	OutcomesUpdateAnOutcomeCalculationMethodDecayingAverage OutcomesUpdateAnOutcomeCalculationMethod = "decaying_average"
	// OutcomesUpdateAnOutcomeCalculationMethodNMastery enum value ("n_mastery")
	OutcomesUpdateAnOutcomeCalculationMethodNMastery OutcomesUpdateAnOutcomeCalculationMethod = "n_mastery"
	// OutcomesUpdateAnOutcomeCalculationMethodLatest enum value ("latest")
	OutcomesUpdateAnOutcomeCalculationMethodLatest OutcomesUpdateAnOutcomeCalculationMethod = "latest"
	// OutcomesUpdateAnOutcomeCalculationMethodHighest enum value ("highest")
	OutcomesUpdateAnOutcomeCalculationMethodHighest OutcomesUpdateAnOutcomeCalculationMethod = "highest"
)

// PeerReviewsGetAllPeerReviewsInclude enumeration
type PeerReviewsGetAllPeerReviewsInclude string

const (
	// PeerReviewsGetAllPeerReviewsIncludeSubmissionComments enum value ("submission_comments")
	PeerReviewsGetAllPeerReviewsIncludeSubmissionComments PeerReviewsGetAllPeerReviewsInclude = "submission_comments"
	// PeerReviewsGetAllPeerReviewsIncludeUser enum value ("user")
	PeerReviewsGetAllPeerReviewsIncludeUser PeerReviewsGetAllPeerReviewsInclude = "user"
)

// PlannerListPlannerItemsFilter enumeration
type PlannerListPlannerItemsFilter string

const (
	// PlannerListPlannerItemsFilterNewActivity enum value ("new_activity")
	PlannerListPlannerItemsFilterNewActivity PlannerListPlannerItemsFilter = "new_activity"
)

// PlannerCreateAPlannerOverridePlannableType enumeration
type PlannerCreateAPlannerOverridePlannableType string

const (
	// PlannerCreateAPlannerOverridePlannableTypeAnnouncement enum value ("announcement")
	PlannerCreateAPlannerOverridePlannableTypeAnnouncement PlannerCreateAPlannerOverridePlannableType = "announcement"
	// PlannerCreateAPlannerOverridePlannableTypeAssignment enum value ("assignment")
	PlannerCreateAPlannerOverridePlannableTypeAssignment PlannerCreateAPlannerOverridePlannableType = "assignment"
	// PlannerCreateAPlannerOverridePlannableTypeDiscussionTopic enum value ("discussion_topic")
	PlannerCreateAPlannerOverridePlannableTypeDiscussionTopic PlannerCreateAPlannerOverridePlannableType = "discussion_topic"
	// PlannerCreateAPlannerOverridePlannableTypeQuiz enum value ("quiz")
	PlannerCreateAPlannerOverridePlannableTypeQuiz PlannerCreateAPlannerOverridePlannableType = "quiz"
	// PlannerCreateAPlannerOverridePlannableTypeWikiPage enum value ("wiki_page")
	PlannerCreateAPlannerOverridePlannableTypeWikiPage PlannerCreateAPlannerOverridePlannableType = "wiki_page"
	// PlannerCreateAPlannerOverridePlannableTypePlannerNote enum value ("planner_note")
	PlannerCreateAPlannerOverridePlannableTypePlannerNote PlannerCreateAPlannerOverridePlannableType = "planner_note"
)

// RolesListRolesState enumeration
type RolesListRolesState string

const (
	// RolesListRolesStateActive enum value ("active")
	RolesListRolesStateActive RolesListRolesState = "active"
	// RolesListRolesStateInactive enum value ("inactive")
	RolesListRolesStateInactive RolesListRolesState = "inactive"
)

// RolesCreateANewRoleBaseRoleType enumeration
type RolesCreateANewRoleBaseRoleType string

const (
	// RolesCreateANewRoleBaseRoleTypeAccountMembership enum value ("AccountMembership")
	RolesCreateANewRoleBaseRoleTypeAccountMembership RolesCreateANewRoleBaseRoleType = "AccountMembership"
	// RolesCreateANewRoleBaseRoleTypeStudentEnrollment enum value ("StudentEnrollment")
	RolesCreateANewRoleBaseRoleTypeStudentEnrollment RolesCreateANewRoleBaseRoleType = "StudentEnrollment"
	// RolesCreateANewRoleBaseRoleTypeTeacherEnrollment enum value ("TeacherEnrollment")
	RolesCreateANewRoleBaseRoleTypeTeacherEnrollment RolesCreateANewRoleBaseRoleType = "TeacherEnrollment"
	// RolesCreateANewRoleBaseRoleTypeTaEnrollment enum value ("TaEnrollment")
	RolesCreateANewRoleBaseRoleTypeTaEnrollment RolesCreateANewRoleBaseRoleType = "TaEnrollment"
	// RolesCreateANewRoleBaseRoleTypeObserverEnrollment enum value ("ObserverEnrollment")
	RolesCreateANewRoleBaseRoleTypeObserverEnrollment RolesCreateANewRoleBaseRoleType = "ObserverEnrollment"
	// RolesCreateANewRoleBaseRoleTypeDesignerEnrollment enum value ("DesignerEnrollment")
	RolesCreateANewRoleBaseRoleTypeDesignerEnrollment RolesCreateANewRoleBaseRoleType = "DesignerEnrollment"
)

// RubricsGetASingleRubricInclude enumeration
type RubricsGetASingleRubricInclude string

const (
	// RubricsGetASingleRubricIncludeAssessments enum value ("assessments")
	RubricsGetASingleRubricIncludeAssessments RubricsGetASingleRubricInclude = "assessments"
	// RubricsGetASingleRubricIncludeGradedAssessments enum value ("graded_assessments")
	RubricsGetASingleRubricIncludeGradedAssessments RubricsGetASingleRubricInclude = "graded_assessments"
	// RubricsGetASingleRubricIncludePeerAssessments enum value ("peer_assessments")
	RubricsGetASingleRubricIncludePeerAssessments RubricsGetASingleRubricInclude = "peer_assessments"
	// RubricsGetASingleRubricIncludeAssociations enum value ("associations")
	RubricsGetASingleRubricIncludeAssociations RubricsGetASingleRubricInclude = "associations"
	// RubricsGetASingleRubricIncludeAssignmentAssociations enum value ("assignment_associations")
	RubricsGetASingleRubricIncludeAssignmentAssociations RubricsGetASingleRubricInclude = "assignment_associations"
	// RubricsGetASingleRubricIncludeCourseAssociations enum value ("course_associations")
	RubricsGetASingleRubricIncludeCourseAssociations RubricsGetASingleRubricInclude = "course_associations"
	// RubricsGetASingleRubricIncludeAccountAssociations enum value ("account_associations")
	RubricsGetASingleRubricIncludeAccountAssociations RubricsGetASingleRubricInclude = "account_associations"
)

// RubricsGetASingleRubricStyle enumeration
type RubricsGetASingleRubricStyle string

const (
	// RubricsGetASingleRubricStyleFull enum value ("full")
	RubricsGetASingleRubricStyleFull RubricsGetASingleRubricStyle = "full"
	// RubricsGetASingleRubricStyleCommentsOnly enum value ("comments_only")
	RubricsGetASingleRubricStyleCommentsOnly RubricsGetASingleRubricStyle = "comments_only"
)

// APITokenScopesListScopesGroupBy enumeration
type APITokenScopesListScopesGroupBy string

const (
	// APITokenScopesListScopesGroupByResourceName enum value ("resource_name")
	APITokenScopesListScopesGroupByResourceName APITokenScopesListScopesGroupBy = "resource_name"
)

// SearchFindRecipientsType enumeration
type SearchFindRecipientsType string

const (
	// SearchFindRecipientsTypeUser enum value ("user")
	SearchFindRecipientsTypeUser SearchFindRecipientsType = "user"
	// SearchFindRecipientsTypeContext enum value ("context")
	SearchFindRecipientsTypeContext SearchFindRecipientsType = "context"
)

// SectionsListCourseSectionsInclude enumeration
type SectionsListCourseSectionsInclude string

const (
	// SectionsListCourseSectionsIncludeStudents enum value ("students")
	SectionsListCourseSectionsIncludeStudents SectionsListCourseSectionsInclude = "students"
	// SectionsListCourseSectionsIncludeAvatarURL enum value ("avatar_url")
	SectionsListCourseSectionsIncludeAvatarURL SectionsListCourseSectionsInclude = "avatar_url"
	// SectionsListCourseSectionsIncludeEnrollments enum value ("enrollments")
	SectionsListCourseSectionsIncludeEnrollments SectionsListCourseSectionsInclude = "enrollments"
	// SectionsListCourseSectionsIncludeTotalStudents enum value ("total_students")
	SectionsListCourseSectionsIncludeTotalStudents SectionsListCourseSectionsInclude = "total_students"
	// SectionsListCourseSectionsIncludePassbackStatus enum value ("passback_status")
	SectionsListCourseSectionsIncludePassbackStatus SectionsListCourseSectionsInclude = "passback_status"
)

// SectionsGetSectionInformationInclude enumeration
type SectionsGetSectionInformationInclude string

const (
	// SectionsGetSectionInformationIncludeStudents enum value ("students")
	SectionsGetSectionInformationIncludeStudents SectionsGetSectionInformationInclude = "students"
	// SectionsGetSectionInformationIncludeAvatarURL enum value ("avatar_url")
	SectionsGetSectionInformationIncludeAvatarURL SectionsGetSectionInformationInclude = "avatar_url"
	// SectionsGetSectionInformationIncludeEnrollments enum value ("enrollments")
	SectionsGetSectionInformationIncludeEnrollments SectionsGetSectionInformationInclude = "enrollments"
	// SectionsGetSectionInformationIncludeTotalStudents enum value ("total_students")
	SectionsGetSectionInformationIncludeTotalStudents SectionsGetSectionInformationInclude = "total_students"
	// SectionsGetSectionInformationIncludePassbackStatus enum value ("passback_status")
	SectionsGetSectionInformationIncludePassbackStatus SectionsGetSectionInformationInclude = "passback_status"
)

// SISImportsGetSISImportListWorkflowState enumeration
type SISImportsGetSISImportListWorkflowState string

const (
	// SISImportsGetSISImportListWorkflowStateInitializing enum value ("initializing")
	SISImportsGetSISImportListWorkflowStateInitializing SISImportsGetSISImportListWorkflowState = "initializing"
	// SISImportsGetSISImportListWorkflowStateCreated enum value ("created")
	SISImportsGetSISImportListWorkflowStateCreated SISImportsGetSISImportListWorkflowState = "created"
	// SISImportsGetSISImportListWorkflowStateImporting enum value ("importing")
	SISImportsGetSISImportListWorkflowStateImporting SISImportsGetSISImportListWorkflowState = "importing"
	// SISImportsGetSISImportListWorkflowStateCleanupBatch enum value ("cleanup_batch")
	SISImportsGetSISImportListWorkflowStateCleanupBatch SISImportsGetSISImportListWorkflowState = "cleanup_batch"
	// SISImportsGetSISImportListWorkflowStateImported enum value ("imported")
	SISImportsGetSISImportListWorkflowStateImported SISImportsGetSISImportListWorkflowState = "imported"
	// SISImportsGetSISImportListWorkflowStateImportedWithMessages enum value ("imported_with_messages")
	SISImportsGetSISImportListWorkflowStateImportedWithMessages SISImportsGetSISImportListWorkflowState = "imported_with_messages"
	// SISImportsGetSISImportListWorkflowStateAborted enum value ("aborted")
	SISImportsGetSISImportListWorkflowStateAborted SISImportsGetSISImportListWorkflowState = "aborted"
	// SISImportsGetSISImportListWorkflowStateFailed enum value ("failed")
	SISImportsGetSISImportListWorkflowStateFailed SISImportsGetSISImportListWorkflowState = "failed"
	// SISImportsGetSISImportListWorkflowStateFailedWithMessages enum value ("failed_with_messages")
	SISImportsGetSISImportListWorkflowStateFailedWithMessages SISImportsGetSISImportListWorkflowState = "failed_with_messages"
	// SISImportsGetSISImportListWorkflowStateRestoring enum value ("restoring")
	SISImportsGetSISImportListWorkflowStateRestoring SISImportsGetSISImportListWorkflowState = "restoring"
	// SISImportsGetSISImportListWorkflowStatePartiallyRestored enum value ("partially_restored")
	SISImportsGetSISImportListWorkflowStatePartiallyRestored SISImportsGetSISImportListWorkflowState = "partially_restored"
	// SISImportsGetSISImportListWorkflowStateRestored enum value ("restored")
	SISImportsGetSISImportListWorkflowStateRestored SISImportsGetSISImportListWorkflowState = "restored"
)

// SISImportsImportSISDataDiffingDropStatus enumeration
type SISImportsImportSISDataDiffingDropStatus string

const (
	// SISImportsImportSISDataDiffingDropStatusDeleted enum value ("deleted")
	SISImportsImportSISDataDiffingDropStatusDeleted SISImportsImportSISDataDiffingDropStatus = "deleted"
	// SISImportsImportSISDataDiffingDropStatusCompleted enum value ("completed")
	SISImportsImportSISDataDiffingDropStatusCompleted SISImportsImportSISDataDiffingDropStatus = "completed"
	// SISImportsImportSISDataDiffingDropStatusInactive enum value ("inactive")
	SISImportsImportSISDataDiffingDropStatusInactive SISImportsImportSISDataDiffingDropStatus = "inactive"
)

// SubmissionsListAssignmentSubmissionsInclude enumeration
type SubmissionsListAssignmentSubmissionsInclude string

const (
	// SubmissionsListAssignmentSubmissionsIncludeSubmissionHistory enum value ("submission_history")
	SubmissionsListAssignmentSubmissionsIncludeSubmissionHistory SubmissionsListAssignmentSubmissionsInclude = "submission_history"
	// SubmissionsListAssignmentSubmissionsIncludeSubmissionComments enum value ("submission_comments")
	SubmissionsListAssignmentSubmissionsIncludeSubmissionComments SubmissionsListAssignmentSubmissionsInclude = "submission_comments"
	// SubmissionsListAssignmentSubmissionsIncludeRubricAssessment enum value ("rubric_assessment")
	SubmissionsListAssignmentSubmissionsIncludeRubricAssessment SubmissionsListAssignmentSubmissionsInclude = "rubric_assessment"
	// SubmissionsListAssignmentSubmissionsIncludeAssignment enum value ("assignment")
	SubmissionsListAssignmentSubmissionsIncludeAssignment SubmissionsListAssignmentSubmissionsInclude = "assignment"
	// SubmissionsListAssignmentSubmissionsIncludeVisibility enum value ("visibility")
	SubmissionsListAssignmentSubmissionsIncludeVisibility SubmissionsListAssignmentSubmissionsInclude = "visibility"
	// SubmissionsListAssignmentSubmissionsIncludeCourse enum value ("course")
	SubmissionsListAssignmentSubmissionsIncludeCourse SubmissionsListAssignmentSubmissionsInclude = "course"
	// SubmissionsListAssignmentSubmissionsIncludeUser enum value ("user")
	SubmissionsListAssignmentSubmissionsIncludeUser SubmissionsListAssignmentSubmissionsInclude = "user"
	// SubmissionsListAssignmentSubmissionsIncludeGroup enum value ("group")
	SubmissionsListAssignmentSubmissionsIncludeGroup SubmissionsListAssignmentSubmissionsInclude = "group"
)

// SubmissionsListSubmissionsForMultipleAssignmentsWorkflowState enumeration
type SubmissionsListSubmissionsForMultipleAssignmentsWorkflowState string

const (
	// SubmissionsListSubmissionsForMultipleAssignmentsWorkflowStateSubmitted enum value ("submitted")
	SubmissionsListSubmissionsForMultipleAssignmentsWorkflowStateSubmitted SubmissionsListSubmissionsForMultipleAssignmentsWorkflowState = "submitted"
	// SubmissionsListSubmissionsForMultipleAssignmentsWorkflowStateUnsubmitted enum value ("unsubmitted")
	SubmissionsListSubmissionsForMultipleAssignmentsWorkflowStateUnsubmitted SubmissionsListSubmissionsForMultipleAssignmentsWorkflowState = "unsubmitted"
	// SubmissionsListSubmissionsForMultipleAssignmentsWorkflowStateGraded enum value ("graded")
	SubmissionsListSubmissionsForMultipleAssignmentsWorkflowStateGraded SubmissionsListSubmissionsForMultipleAssignmentsWorkflowState = "graded"
	// SubmissionsListSubmissionsForMultipleAssignmentsWorkflowStatePendingReview enum value ("pending_review")
	SubmissionsListSubmissionsForMultipleAssignmentsWorkflowStatePendingReview SubmissionsListSubmissionsForMultipleAssignmentsWorkflowState = "pending_review"
)

// SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentState enumeration
type SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentState string

const (
	// SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentStateActive enum value ("active")
	SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentStateActive SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentState = "active"
	// SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentStateConcluded enum value ("concluded")
	SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentStateConcluded SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentState = "concluded"
	// SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentStateNone enum value ("")
	SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentStateNone SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentState = ""
)

// SubmissionsListSubmissionsForMultipleAssignmentsOrder enumeration
type SubmissionsListSubmissionsForMultipleAssignmentsOrder string

const (
	// SubmissionsListSubmissionsForMultipleAssignmentsOrderID enum value ("id")
	SubmissionsListSubmissionsForMultipleAssignmentsOrderID SubmissionsListSubmissionsForMultipleAssignmentsOrder = "id"
	// SubmissionsListSubmissionsForMultipleAssignmentsOrderGradedAt enum value ("graded_at")
	SubmissionsListSubmissionsForMultipleAssignmentsOrderGradedAt SubmissionsListSubmissionsForMultipleAssignmentsOrder = "graded_at"
)

// SubmissionsListSubmissionsForMultipleAssignmentsOrderDirection enumeration
type SubmissionsListSubmissionsForMultipleAssignmentsOrderDirection string

const (
	// SubmissionsListSubmissionsForMultipleAssignmentsOrderDirectionAscending enum value ("ascending")
	SubmissionsListSubmissionsForMultipleAssignmentsOrderDirectionAscending SubmissionsListSubmissionsForMultipleAssignmentsOrderDirection = "ascending"
	// SubmissionsListSubmissionsForMultipleAssignmentsOrderDirectionDescending enum value ("descending")
	SubmissionsListSubmissionsForMultipleAssignmentsOrderDirectionDescending SubmissionsListSubmissionsForMultipleAssignmentsOrderDirection = "descending"
)

// SubmissionsListSubmissionsForMultipleAssignmentsInclude enumeration
type SubmissionsListSubmissionsForMultipleAssignmentsInclude string

const (
	// SubmissionsListSubmissionsForMultipleAssignmentsIncludeSubmissionHistory enum value ("submission_history")
	SubmissionsListSubmissionsForMultipleAssignmentsIncludeSubmissionHistory SubmissionsListSubmissionsForMultipleAssignmentsInclude = "submission_history"
	// SubmissionsListSubmissionsForMultipleAssignmentsIncludeSubmissionComments enum value ("submission_comments")
	SubmissionsListSubmissionsForMultipleAssignmentsIncludeSubmissionComments SubmissionsListSubmissionsForMultipleAssignmentsInclude = "submission_comments"
	// SubmissionsListSubmissionsForMultipleAssignmentsIncludeRubricAssessment enum value ("rubric_assessment")
	SubmissionsListSubmissionsForMultipleAssignmentsIncludeRubricAssessment SubmissionsListSubmissionsForMultipleAssignmentsInclude = "rubric_assessment"
	// SubmissionsListSubmissionsForMultipleAssignmentsIncludeAssignment enum value ("assignment")
	SubmissionsListSubmissionsForMultipleAssignmentsIncludeAssignment SubmissionsListSubmissionsForMultipleAssignmentsInclude = "assignment"
	// SubmissionsListSubmissionsForMultipleAssignmentsIncludeTotalScores enum value ("total_scores")
	SubmissionsListSubmissionsForMultipleAssignmentsIncludeTotalScores SubmissionsListSubmissionsForMultipleAssignmentsInclude = "total_scores"
	// SubmissionsListSubmissionsForMultipleAssignmentsIncludeVisibility enum value ("visibility")
	SubmissionsListSubmissionsForMultipleAssignmentsIncludeVisibility SubmissionsListSubmissionsForMultipleAssignmentsInclude = "visibility"
	// SubmissionsListSubmissionsForMultipleAssignmentsIncludeCourse enum value ("course")
	SubmissionsListSubmissionsForMultipleAssignmentsIncludeCourse SubmissionsListSubmissionsForMultipleAssignmentsInclude = "course"
	// SubmissionsListSubmissionsForMultipleAssignmentsIncludeUser enum value ("user")
	SubmissionsListSubmissionsForMultipleAssignmentsIncludeUser SubmissionsListSubmissionsForMultipleAssignmentsInclude = "user"
)

// SubmissionsGetASingleSubmissionInclude enumeration
type SubmissionsGetASingleSubmissionInclude string

const (
	// SubmissionsGetASingleSubmissionIncludeSubmissionHistory enum value ("submission_history")
	SubmissionsGetASingleSubmissionIncludeSubmissionHistory SubmissionsGetASingleSubmissionInclude = "submission_history"
	// SubmissionsGetASingleSubmissionIncludeSubmissionComments enum value ("submission_comments")
	SubmissionsGetASingleSubmissionIncludeSubmissionComments SubmissionsGetASingleSubmissionInclude = "submission_comments"
	// SubmissionsGetASingleSubmissionIncludeRubricAssessment enum value ("rubric_assessment")
	SubmissionsGetASingleSubmissionIncludeRubricAssessment SubmissionsGetASingleSubmissionInclude = "rubric_assessment"
	// SubmissionsGetASingleSubmissionIncludeFullRubricAssessment enum value ("full_rubric_assessment")
	SubmissionsGetASingleSubmissionIncludeFullRubricAssessment SubmissionsGetASingleSubmissionInclude = "full_rubric_assessment"
	// SubmissionsGetASingleSubmissionIncludeVisibility enum value ("visibility")
	SubmissionsGetASingleSubmissionIncludeVisibility SubmissionsGetASingleSubmissionInclude = "visibility"
	// SubmissionsGetASingleSubmissionIncludeCourse enum value ("course")
	SubmissionsGetASingleSubmissionIncludeCourse SubmissionsGetASingleSubmissionInclude = "course"
	// SubmissionsGetASingleSubmissionIncludeUser enum value ("user")
	SubmissionsGetASingleSubmissionIncludeUser SubmissionsGetASingleSubmissionInclude = "user"
)

// SubmissionsSubmitAnAssignmentSubmission enumeration
type SubmissionsSubmitAnAssignmentSubmission string

const (
	// SubmissionsSubmitAnAssignmentSubmissionOnlineTextEntry enum value ("online_text_entry")
	SubmissionsSubmitAnAssignmentSubmissionOnlineTextEntry SubmissionsSubmitAnAssignmentSubmission = "online_text_entry"
	// SubmissionsSubmitAnAssignmentSubmissionOnlineURL enum value ("online_url")
	SubmissionsSubmitAnAssignmentSubmissionOnlineURL SubmissionsSubmitAnAssignmentSubmission = "online_url"
	// SubmissionsSubmitAnAssignmentSubmissionOnlineUpload enum value ("online_upload")
	SubmissionsSubmitAnAssignmentSubmissionOnlineUpload SubmissionsSubmitAnAssignmentSubmission = "online_upload"
	// SubmissionsSubmitAnAssignmentSubmissionMediaRecording enum value ("media_recording")
	SubmissionsSubmitAnAssignmentSubmissionMediaRecording SubmissionsSubmitAnAssignmentSubmission = "media_recording"
	// SubmissionsSubmitAnAssignmentSubmissionBasicLtiLaunch enum value ("basic_lti_launch")
	SubmissionsSubmitAnAssignmentSubmissionBasicLtiLaunch SubmissionsSubmitAnAssignmentSubmission = "basic_lti_launch"
)

// EnrollmentTermsListEnrollmentTermsWorkflowState enumeration
type EnrollmentTermsListEnrollmentTermsWorkflowState string

const (
	// EnrollmentTermsListEnrollmentTermsWorkflowStateActive enum value ("active")
	EnrollmentTermsListEnrollmentTermsWorkflowStateActive EnrollmentTermsListEnrollmentTermsWorkflowState = "active"
	// EnrollmentTermsListEnrollmentTermsWorkflowStateDeleted enum value ("deleted")
	EnrollmentTermsListEnrollmentTermsWorkflowStateDeleted EnrollmentTermsListEnrollmentTermsWorkflowState = "deleted"
	// EnrollmentTermsListEnrollmentTermsWorkflowStateAll enum value ("all")
	EnrollmentTermsListEnrollmentTermsWorkflowStateAll EnrollmentTermsListEnrollmentTermsWorkflowState = "all"
)

// EnrollmentTermsListEnrollmentTermsInclude enumeration
type EnrollmentTermsListEnrollmentTermsInclude string

const (
	// EnrollmentTermsListEnrollmentTermsIncludeOverrides enum value ("overrides")
	EnrollmentTermsListEnrollmentTermsIncludeOverrides EnrollmentTermsListEnrollmentTermsInclude = "overrides"
)

// FilesSetUsageRightsUsageRights enumeration
type FilesSetUsageRightsUsageRights string

const (
	// FilesSetUsageRightsUsageRightsOwnCopyright enum value ("own_copyright")
	FilesSetUsageRightsUsageRightsOwnCopyright FilesSetUsageRightsUsageRights = "own_copyright"
	// FilesSetUsageRightsUsageRightsUsedByPermission enum value ("used_by_permission")
	FilesSetUsageRightsUsageRightsUsedByPermission FilesSetUsageRightsUsageRights = "used_by_permission"
	// FilesSetUsageRightsUsageRightsFairUse enum value ("fair_use")
	FilesSetUsageRightsUsageRightsFairUse FilesSetUsageRightsUsageRights = "fair_use"
	// FilesSetUsageRightsUsageRightsPublicDomain enum value ("public_domain")
	FilesSetUsageRightsUsageRightsPublicDomain FilesSetUsageRightsUsageRights = "public_domain"
	// FilesSetUsageRightsUsageRightsCreativeCommons enum value ("creative_commons")
	FilesSetUsageRightsUsageRightsCreativeCommons FilesSetUsageRightsUsageRights = "creative_commons"
)

// UserObserveesListObserveesInclude enumeration
type UserObserveesListObserveesInclude string

const (
	// UserObserveesListObserveesIncludeAvatarURL enum value ("avatar_url")
	UserObserveesListObserveesIncludeAvatarURL UserObserveesListObserveesInclude = "avatar_url"
)

// UsersListUsersInAccountSort enumeration
type UsersListUsersInAccountSort string

const (
	// UsersListUsersInAccountSortUsername enum value ("username")
	UsersListUsersInAccountSortUsername UsersListUsersInAccountSort = "username"
	// UsersListUsersInAccountSortEmail enum value ("email")
	UsersListUsersInAccountSortEmail UsersListUsersInAccountSort = "email"
	// UsersListUsersInAccountSortSisID enum value ("sis_id")
	UsersListUsersInAccountSortSisID UsersListUsersInAccountSort = "sis_id"
	// UsersListUsersInAccountSortLastLogin enum value ("last_login")
	UsersListUsersInAccountSortLastLogin UsersListUsersInAccountSort = "last_login"
)

// UsersListUsersInAccountOrder enumeration
type UsersListUsersInAccountOrder string

const (
	// UsersListUsersInAccountOrderAsc enum value ("asc")
	UsersListUsersInAccountOrderAsc UsersListUsersInAccountOrder = "asc"
	// UsersListUsersInAccountOrderDesc enum value ("desc")
	UsersListUsersInAccountOrderDesc UsersListUsersInAccountOrder = "desc"
)

// UsersListTheTODOItemsInclude enumeration
type UsersListTheTODOItemsInclude string

const (
	// UsersListTheTODOItemsIncludeUngradedQuizzes enum value ("ungraded_quizzes")
	UsersListTheTODOItemsIncludeUngradedQuizzes UsersListTheTODOItemsInclude = "ungraded_quizzes"
)

// UsersListCountsForTodoItemsInclude enumeration
type UsersListCountsForTodoItemsInclude string

const (
	// UsersListCountsForTodoItemsIncludeUngradedQuizzes enum value ("ungraded_quizzes")
	UsersListCountsForTodoItemsIncludeUngradedQuizzes UsersListCountsForTodoItemsInclude = "ungraded_quizzes"
)

// UsersListMissingSubmissionsInclude enumeration
type UsersListMissingSubmissionsInclude string

const (
	// UsersListMissingSubmissionsIncludePlannerOverrides enum value ("planner_overrides")
	UsersListMissingSubmissionsIncludePlannerOverrides UsersListMissingSubmissionsInclude = "planner_overrides"
	// UsersListMissingSubmissionsIncludeCourse enum value ("course")
	UsersListMissingSubmissionsIncludeCourse UsersListMissingSubmissionsInclude = "course"
)

// UsersListMissingSubmissionsFilter enumeration
type UsersListMissingSubmissionsFilter string

const (
	// UsersListMissingSubmissionsFilterSubmittable enum value ("submittable")
	UsersListMissingSubmissionsFilterSubmittable UsersListMissingSubmissionsFilter = "submittable"
)

// UsersShowUserDetailsInclude enumeration
type UsersShowUserDetailsInclude string

const (
	// UsersShowUserDetailsIncludeUUID enum value ("uuid")
	UsersShowUserDetailsIncludeUUID UsersShowUserDetailsInclude = "uuid"
)

// UsersGetAUsersMostRecentlyGradedSubmissionsInclude enumeration
type UsersGetAUsersMostRecentlyGradedSubmissionsInclude string

const (
	// UsersGetAUsersMostRecentlyGradedSubmissionsIncludeAssignment enum value ("assignment")
	UsersGetAUsersMostRecentlyGradedSubmissionsIncludeAssignment UsersGetAUsersMostRecentlyGradedSubmissionsInclude = "assignment"
)

// PagesListPagesSort enumeration
type PagesListPagesSort string

const (
	// PagesListPagesSortTitle enum value ("title")
	PagesListPagesSortTitle PagesListPagesSort = "title"
	// PagesListPagesSortCreatedAt enum value ("created_at")
	PagesListPagesSortCreatedAt PagesListPagesSort = "created_at"
	// PagesListPagesSortUpdatedAt enum value ("updated_at")
	PagesListPagesSortUpdatedAt PagesListPagesSort = "updated_at"
)

// PagesListPagesOrder enumeration
type PagesListPagesOrder string

const (
	// PagesListPagesOrderAsc enum value ("asc")
	PagesListPagesOrderAsc PagesListPagesOrder = "asc"
	// PagesListPagesOrderDesc enum value ("desc")
	PagesListPagesOrderDesc PagesListPagesOrder = "desc"
)

// AccountNotification model object
type AccountNotification struct {
	// EndAt field: When to expire the notification.
	EndAt time.Time `json:"end_at"`
	// Icon field: The icon to display with the message.  Defaults to warning.
	Icon *AccountNotificationIcon `json:"icon"`
	// Message field: The message to be sent in the notification.
	Message string `json:"message"`
	// RoleIds field: The roles to send the notification to.  If roles is not passed it defaults to all roles
	RoleIds []int `json:"role_ids"`
	// Roles field: (Deprecated) The roles to send the notification to.  If roles is not passed it defaults to all roles
	Roles []string `json:"roles"`
	// StartAt field: When to send out the notification.
	StartAt time.Time `json:"start_at"`
	// Subject field: The subject of the notifications
	Subject string `json:"subject"`
}

// Report model object
type Report struct {
	// Attachment field: The attachment api object of the report. Only available after the report has completed.
	Attachment *File `json:"attachment"`
	// CreatedAt field: The date and time the report was created.
	CreatedAt time.Time `json:"created_at"`
	// CurrentLine field: This is the current line count being written to the report. It updates every 1000 records.
	CurrentLine int `json:"current_line"`
	// EndedAt field: The date and time the report finished processing.
	EndedAt time.Time `json:"ended_at"`
	// FileURL field: The url to the report download.
	FileURL string `json:"file_url"`
	// ID field: The unique identifier for the report.
	ID int `json:"id"`
	// Parameters field: The report parameters
	Parameters *ReportParameters `json:"parameters"`
	// Progress field: The progress of the report
	Progress int `json:"progress"`
	// Report field: The type of report.
	Report string `json:"report"`
	// StartedAt field: The date and time the report started processing.
	StartedAt time.Time `json:"started_at"`
	// Status field: The status of the report
	Status string `json:"status"`
}

// ReportParameters model object: The parameters returned will vary for each report.
type ReportParameters struct {
	// Accounts field: If true, account data will be included. If false, account data will be omitted.
	Accounts bool `json:"accounts"`
	// CourseID field: The id of the course to report on
	CourseID int `json:"course_id"`
	// Courses field: If true, course data will be included. If false, course data will be omitted.
	Courses bool `json:"courses"`
	// EndAt field: The end date for submissions. Max time range is 2 weeks.
	EndAt time.Time `json:"end_at"`
	// EnrollmentState field: Include enrollment state. Defaults to 'all' Options: ['active'| 'invited'|
	// 'creation_pending'| 'deleted'| 'rejected'| 'completed'| 'inactive'| 'all']
	EnrollmentState []string `json:"enrollment_state"`
	// EnrollmentTermID field: The canvas id of the term to get grades from
	EnrollmentTermID int `json:"enrollment_term_id"`
	// Enrollments field: If true, enrollment data will be included. If false, enrollment data will be omitted.
	Enrollments bool `json:"enrollments"`
	// Groups field: If true, group data will be included. If false, group data will be omitted.
	Groups bool `json:"groups"`
	// IncludeDeleted field: If true, deleted objects will be included. If false, deleted objects will be omitted.
	IncludeDeleted bool `json:"include_deleted"`
	// IncludeEnrollmentState field: If true, enrollment state will be included. If false, enrollment state will be
	// omitted. Defaults to false.
	IncludeEnrollmentState bool `json:"include_enrollment_state"`
	// Order field: The sort order for the csv, Options: 'users', 'courses', 'outcomes'.
	Order *ReportParametersOrder `json:"order"`
	// Sections field: If true, section data will be included. If false, section data will be omitted.
	Sections bool `json:"sections"`
	// SisAccountsCsv field
	SisAccountsCsv int `json:"sis_accounts_csv"`
	// SisTermsCsv field
	SisTermsCsv int `json:"sis_terms_csv"`
	// StartAt field: The beginning date for submissions. Max time range is 2 weeks.
	StartAt time.Time `json:"start_at"`
	// Terms field: If true, term data will be included. If false, term data will be omitted.
	Terms bool `json:"terms"`
	// Users field: If true, user data will be included. If false, user data will be omitted.
	Users bool `json:"users"`
	// Xlist field: If true, data for crosslisted courses will be included. If false, data for crosslisted courses will
	// be omitted.
	Xlist bool `json:"xlist"`
}

// Account model object
type Account struct {
	// DefaultGroupStorageQuotaMb field: The storage quota for a group in the account in megabytes, if not otherwise
	// specified
	DefaultGroupStorageQuotaMb int `json:"default_group_storage_quota_mb"`
	// DefaultStorageQuotaMb field: The storage quota for the account in megabytes, if not otherwise specified
	DefaultStorageQuotaMb int `json:"default_storage_quota_mb"`
	// DefaultTimeZone field: The default time zone of the account. Allowed time zones are
	// {http://www.iana.org/time-zones IANA time zones} or friendlier
	// {http://api.rubyonrails.org/classes/ActiveSupport/TimeZone.html Ruby on Rails time zones}.
	DefaultTimeZone string `json:"default_time_zone"`
	// DefaultUserStorageQuotaMb field: The storage quota for a user in the account in megabytes, if not otherwise
	// specified
	DefaultUserStorageQuotaMb int `json:"default_user_storage_quota_mb"`
	// ID field: the ID of the Account object
	ID int `json:"id"`
	// IntegrationID field: The account's identifier in the Student Information System. Only included if the user has
	// permission to view SIS information.
	IntegrationID string `json:"integration_id"`
	// LtiGUID field: The account's identifier that is sent as context_id in LTI launches.
	LtiGUID string `json:"lti_guid"`
	// Name field: The display name of the account
	Name string `json:"name"`
	// ParentAccountID field: The account's parent ID, or null if this is the root account
	ParentAccountID int `json:"parent_account_id"`
	// RootAccountID field: The ID of the root account, or null if this is the root account
	RootAccountID int `json:"root_account_id"`
	// SisAccountID field: The account's identifier in the Student Information System. Only included if the user has
	// permission to view SIS information.
	SisAccountID string `json:"sis_account_id"`
	// SisImportID field: The id of the SIS import if created through SIS. Only included if the user has permission to
	// manage SIS information.
	SisImportID int `json:"sis_import_id"`
	// UUID field: The UUID of the account
	UUID string `json:"uuid"`
	// WorkflowState field: The state of the account. Can be 'active' or 'deleted'.
	WorkflowState string `json:"workflow_state"`
}

// TermsOfService model object
type TermsOfService struct {
	// AccountID field: The id of the root account that owns the Terms of Service
	AccountID int `json:"account_id"`
	// Content field: Content of the Terms of Service
	Content string `json:"content"`
	// ID field: Terms Of Service id
	ID int `json:"id"`
	// Passive field: Boolean dictating if the user must accept Terms of Service
	Passive bool `json:"passive"`
	// TermsType field: The given type for the Terms of Service
	TermsType string `json:"terms_type"`
}

// HelpLink model object
type HelpLink struct {
	// AvailableTo field: The roles that have access to this help link
	AvailableTo []string `json:"available_to"`
	// ID field: The ID of the help link
	ID string `json:"id"`
	// Subtext field: The description of the help link
	Subtext string `json:"subtext"`
	// Text field: The name of the help link
	Text string `json:"text"`
	// Type field: The type of the help link
	Type string `json:"type"`
	// URL field: The URL of the help link
	URL string `json:"url"`
}

// HelpLinks model object
type HelpLinks struct {
	// CustomHelpLinks field: Help links defined by the account. Could include default help links.
	CustomHelpLinks []HelpLink `json:"custom_help_links"`
	// DefaultHelpLinks field: Default help links provided when account has not set help links of their own.
	DefaultHelpLinks []HelpLink `json:"default_help_links"`
	// HelpLinkIcon field: Help link button icon
	HelpLinkIcon string `json:"help_link_icon"`
	// HelpLinkName field: Help link button title
	HelpLinkName string `json:"help_link_name"`
}

// Admin model object
type Admin struct {
	// ID field: The unique identifier for the account role/user assignment.
	ID int `json:"id"`
	// Role field: The account role assigned. This can be 'AccountAdmin' or a user-defined role created by the Roles
	// API.
	Role string `json:"role"`
	// User field: The user the role is assigned to. See the Users API for details.
	User *User `json:"user"`
	// WorkflowState field: The status of the account role/user assignment.
	WorkflowState string `json:"workflow_state"`
}

// Appointment model object: Date and time for an appointment
type Appointment struct {
	// EndAt field: End time for the appointment
	EndAt time.Time `json:"end_at"`
	// ID field: The appointment identifier.
	ID int `json:"id"`
	// StartAt field: Start time for the appointment
	StartAt time.Time `json:"start_at"`
}

// AppointmentGroup model object
type AppointmentGroup struct {
	// Appointments field: Calendar Events representing the time slots (see include[] argument) Refer to the Calendar
	// Events API for more information
	Appointments []CalendarEvent `json:"appointments"`
	// AppointmentsCount field: Number of time slots in this appointment group
	AppointmentsCount int `json:"appointments_count"`
	// ContextCodes field: The context codes (i.e. courses) this appointment group belongs to. Only people in these
	// courses will be eligible to sign up.
	ContextCodes []string `json:"context_codes"`
	// CreatedAt field: When the appointment group was created
	CreatedAt time.Time `json:"created_at"`
	// Description field: The text description of the appointment group
	Description string `json:"description"`
	// EndAt field: The end of the last time slot in the appointment group
	EndAt time.Time `json:"end_at"`
	// HTMLURL field: URL for a user to view this appointment group
	HTMLURL string `json:"html_url"`
	// ID field: The ID of the appointment group
	ID int `json:"id"`
	// LocationAddress field: The address of the appointment group's location
	LocationAddress string `json:"location_address"`
	// LocationName field: The location name of the appointment group
	LocationName string `json:"location_name"`
	// MaxAppointmentsPerParticipant field: Maximum number of time slots a user may register for, or null if no limit
	MaxAppointmentsPerParticipant int `json:"max_appointments_per_participant"`
	// MinAppointmentsPerParticipant field: Minimum number of time slots a user must register for. If not set, users do
	// not need to sign up for any time slots
	MinAppointmentsPerParticipant int `json:"min_appointments_per_participant"`
	// NewAppointments field: Newly created time slots (same format as appointments above). Only returned in
	// Create/Update responses where new time slots have been added
	NewAppointments []CalendarEvent `json:"new_appointments"`
	// ParticipantCount field: The number of participant who have reserved slots (see include[] argument)
	ParticipantCount int `json:"participant_count"`
	// ParticipantType field: Indicates how participants sign up for the appointment group, either as individuals
	// ('User') or in student groups ('Group'). Related to sub_context_codes (i.e. 'Group' signups always have a single
	// group category)
	ParticipantType *AppointmentGroupParticipantType `json:"participant_type"`
	// ParticipantVisibility field: 'private' means participants cannot see who has signed up for a particular time
	// slot, 'protected' means that they can
	ParticipantVisibility *AppointmentGroupParticipantVisibility `json:"participant_visibility"`
	// ParticipantsPerAppointment field: Maximum number of participants that may register for each time slot, or null if
	// no limit
	ParticipantsPerAppointment int `json:"participants_per_appointment"`
	// RequiringAction field: Boolean indicating whether the current user needs to sign up for this appointment group
	// (i.e. it's reservable and the min_appointments_per_participant limit has not been met by this user).
	RequiringAction bool `json:"requiring_action"`
	// ReservedTimes field: The start and end times of slots reserved by the current user as well as the id of the
	// calendar event for the reservation (see include[] argument)
	ReservedTimes []Appointment `json:"reserved_times"`
	// StartAt field: The start of the first time slot in the appointment group
	StartAt time.Time `json:"start_at"`
	// SubContextCodes field: The sub-context codes (i.e. course sections and group categories) this appointment group
	// is restricted to
	SubContextCodes []int `json:"sub_context_codes"`
	// Title field: The title of the appointment group
	Title string `json:"title"`
	// UpdatedAt field: When the appointment group was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// URL field: URL for this appointment group (to update, delete, etc.)
	URL string `json:"url"`
	// WorkflowState field: Current state of the appointment group ('pending', 'active' or 'deleted'). 'pending'
	// indicates that it has not been published yet and is invisible to participants.
	WorkflowState *AppointmentGroupWorkflowState `json:"workflow_state"`
}

// AssignmentExtension model object
type AssignmentExtension struct {
	// AssignmentID field: The ID of the Assignment the extension belongs to.
	AssignmentID int `json:"assignment_id"`
	// ExtraAttempts field: Number of times the student is allowed to re-submit the assignment
	ExtraAttempts int `json:"extra_attempts"`
	// UserID field: The ID of the Student that needs the assignment extension.
	UserID int `json:"user_id"`
}

// GradingRules model object
type GradingRules struct {
	// DropHighest field: Number of highest scores to be dropped for each user.
	DropHighest int `json:"drop_highest"`
	// DropLowest field: Number of lowest scores to be dropped for each user.
	DropLowest int `json:"drop_lowest"`
	// NeverDrop field: Assignment IDs that should never be dropped.
	NeverDrop []int `json:"never_drop"`
}

// AssignmentGroup model object
type AssignmentGroup struct {
	// Assignments field: the assignments in this Assignment Group (see the Assignment API for a detailed list of
	// fields)
	Assignments []int `json:"assignments"`
	// GroupWeight field: the weight of the Assignment Group
	GroupWeight int `json:"group_weight"`
	// ID field: the id of the Assignment Group
	ID int `json:"id"`
	// IntegrationData field: the integration data of the Assignment Group
	IntegrationData map[interface{}]interface{} `json:"integration_data"`
	// Name field: the name of the Assignment Group
	Name string `json:"name"`
	// Position field: the position of the Assignment Group
	Position int `json:"position"`
	// Rules field: the grading rules that this Assignment Group has
	Rules *GradingRules `json:"rules"`
	// SisSourceID field: the sis source id of the Assignment Group
	SisSourceID string `json:"sis_source_id"`
}

// AssignmentOverride model object
type AssignmentOverride struct {
	// AllDay field: the overridden all day flag (present if due_at is overridden)
	AllDay bool `json:"all_day"`
	// AllDayDate field: the overridden all day date (present if due_at is overridden)
	AllDayDate time.Time `json:"all_day_date"`
	// AssignmentID field: the ID of the assignment the override applies to
	AssignmentID int `json:"assignment_id"`
	// CourseSectionID field: the ID of the overrides's target section (present if the override targets a section)
	CourseSectionID int `json:"course_section_id"`
	// DueAt field: the overridden due at (present if due_at is overridden)
	DueAt time.Time `json:"due_at"`
	// GroupID field: the ID of the override's target group (present if the override targets a group and the assignment
	// is a group assignment)
	GroupID int `json:"group_id"`
	// ID field: the ID of the assignment override
	ID int `json:"id"`
	// LockAt field: the overridden lock at, if any (present if lock_at is overridden)
	LockAt time.Time `json:"lock_at"`
	// StudentIds field: the IDs of the override's target students (present if the override targets an ad-hoc set of
	// students)
	StudentIds []int `json:"student_ids"`
	// Title field: the title of the override
	Title string `json:"title"`
	// UnlockAt field: the overridden unlock at (present if unlock_at is overridden)
	UnlockAt time.Time `json:"unlock_at"`
}

// ExternalToolTagAttributes model object
type ExternalToolTagAttributes struct {
	// NewTab field: Whether or not there is a new tab for the external tool
	NewTab bool `json:"new_tab"`
	// ResourceLinkID field: the identifier for this tool_tag
	ResourceLinkID string `json:"resource_link_id"`
	// URL field: URL to the external tool
	URL string `json:"url"`
}

// LockInfo model object
type LockInfo struct {
	// AssetString field: Asset string for the object causing the lock
	AssetString string `json:"asset_string"`
	// ContextModule field: (Optional) Context module causing the lock.
	ContextModule string `json:"context_module"`
	// LockAt field: (Optional) Time at which this was/will be locked. Must be after the due date.
	LockAt time.Time `json:"lock_at"`
	// ManuallyLocked field
	ManuallyLocked bool `json:"manually_locked"`
	// UnlockAt field: (Optional) Time at which this was/will be unlocked. Must be before the due date.
	UnlockAt time.Time `json:"unlock_at"`
}

// RubricRating model object
type RubricRating struct {
	// Description field
	Description string `json:"description"`
	// ID field
	ID string `json:"id"`
	// LongDescription field
	LongDescription string `json:"long_description"`
	// Points field
	Points int `json:"points"`
	// CriterionID field
	CriterionID string `json:"criterion_id"`
}

// RubricCriteria model object
type RubricCriteria struct {
	// CriterionUseRange field
	CriterionUseRange bool `json:"criterion_use_range"`
	// Description field
	Description string `json:"description"`
	// ID field: The id of rubric criteria.
	ID string `json:"id"`
	// IgnoreForScoring field
	IgnoreForScoring bool `json:"ignore_for_scoring"`
	// LearningOutcomeID field: (Optional) The id of the learning outcome this criteria uses, if any.
	LearningOutcomeID string `json:"learning_outcome_id"`
	// LongDescription field
	LongDescription string `json:"long_description"`
	// Points field
	Points int `json:"points"`
	// Ratings field
	Ratings []RubricRating `json:"ratings"`
	// VendorGUID field: (Optional) The 3rd party vendor's GUID for the outcome this criteria references, if any.
	VendorGUID string `json:"vendor_guid"`
}

// AssignmentDate model object: Object representing a due date for an assignment or quiz. If the due date came from an
// assignment override, it will have an 'id' field.
type AssignmentDate struct {
	// Base field: (Optional, present if 'id' is missing) whether this date represents the assignment's or quiz's
	// default due date
	Base bool `json:"base"`
	// DueAt field: The due date for the assignment. Must be between the unlock date and the lock date if there are lock
	// dates
	DueAt time.Time `json:"due_at"`
	// ID field: (Optional, missing if 'base' is present) id of the assignment override this date represents
	ID int `json:"id"`
	// LockAt field: The lock date for the assignment. Must be after the due date if there is a due date.
	LockAt time.Time `json:"lock_at"`
	// Title field
	Title string `json:"title"`
	// UnlockAt field: The unlock date for the assignment. Must be before the due date if there is a due date.
	UnlockAt time.Time `json:"unlock_at"`
}

// TurnitinSettings model object
type TurnitinSettings struct {
	// ExcludeBiblio field
	ExcludeBiblio bool `json:"exclude_biblio"`
	// ExcludeQuoted field
	ExcludeQuoted bool `json:"exclude_quoted"`
	// ExcludeSmallMatchesType field
	ExcludeSmallMatchesType string `json:"exclude_small_matches_type"`
	// ExcludeSmallMatchesValue field
	ExcludeSmallMatchesValue int `json:"exclude_small_matches_value"`
	// InternetCheck field
	InternetCheck bool `json:"internet_check"`
	// JournalCheck field
	JournalCheck bool `json:"journal_check"`
	// OriginalityReportVisibility field
	OriginalityReportVisibility string `json:"originality_report_visibility"`
	// SPaperCheck field
	SPaperCheck bool `json:"s_paper_check"`
}

// NeedsGradingCount model object: Used by Assignment model
type NeedsGradingCount struct {
	// NeedsGradingCount field: Number of submissions that need grading
	NeedsGradingCount int `json:"needs_grading_count"`
	// SectionID field: The section ID
	SectionID string `json:"section_id"`
}

// Assignment model object
type Assignment struct {
	// AllDates field: (Optional) all dates associated with the assignment, if applicable
	AllDates []AssignmentDate `json:"all_dates"`
	// AllowedAttempts field: The number of submission attempts a student can make for this assignment. -1 is considered
	// unlimited.
	AllowedAttempts int `json:"allowed_attempts"`
	// AllowedExtensions field: Allowed file extensions, which take effect if submission_types includes 'online_upload'.
	AllowedExtensions []string `json:"allowed_extensions"`
	// AnonymousGrading field: Boolean indicating if the assignment is graded anonymously. If true, graders cannot see
	// student identities.
	AnonymousGrading bool `json:"anonymous_grading"`
	// AnonymousSubmissions field: (Optional) whether anonymous submissions are accepted (applies only to quiz
	// assignments)
	AnonymousSubmissions bool `json:"anonymous_submissions"`
	// AssignmentGroupID field: the ID of the assignment's group
	AssignmentGroupID int `json:"assignment_group_id"`
	// AssignmentVisibility field: (Optional) If 'assignment_visibility' is included in the 'include' parameter,
	// includes an array of student IDs who can see this assignment.
	AssignmentVisibility []int `json:"assignment_visibility"`
	// AutomaticPeerReviews field: Boolean indicating peer reviews are assigned automatically. If false, the teacher is
	// expected to manually assign peer reviews.
	AutomaticPeerReviews bool `json:"automatic_peer_reviews"`
	// CourseID field: the ID of the course the assignment belongs to
	CourseID int `json:"course_id"`
	// CreatedAt field: The time at which this assignment was originally created
	CreatedAt time.Time `json:"created_at"`
	// Description field: the assignment description, in an HTML fragment
	Description string `json:"description"`
	// DiscussionTopic field: (Optional) the DiscussionTopic associated with the assignment, if applicable
	DiscussionTopic *DiscussionTopic `json:"discussion_topic"`
	// DueAt field: the due date for the assignment. returns null if not present. NOTE: If this assignment has
	// assignment overrides, this field will be the due date as it applies to the user requesting information from the
	// API.
	DueAt time.Time `json:"due_at"`
	// DueDateRequired field: Boolean flag indicating whether the assignment requires a due date based on the account
	// level setting
	DueDateRequired bool `json:"due_date_required"`
	// ExternalToolTagAttributes field: (Optional) assignment's settings for external tools if submission_types include
	// 'external_tool'. Only url and new_tab are included (new_tab defaults to false).  Use the 'External Tools' API if
	// you need more information about an external tool.
	ExternalToolTagAttributes *ExternalToolTagAttributes `json:"external_tool_tag_attributes"`
	// FinalGraderID field: The user ID of the grader responsible for choosing final grades for this assignment. Only
	// relevant for moderated assignments.
	FinalGraderID int `json:"final_grader_id"`
	// FreezeOnCopy field: (Optional) Boolean indicating if assignment will be frozen when it is copied. NOTE: This
	// field will only be present if the AssignmentFreezer plugin is available for your account.
	FreezeOnCopy bool `json:"freeze_on_copy"`
	// Frozen field: (Optional) Boolean indicating if assignment is frozen for the calling user. NOTE: This field will
	// only be present if the AssignmentFreezer plugin is available for your account.
	Frozen bool `json:"frozen"`
	// FrozenAttributes field: (Optional) Array of frozen attributes for the assignment. Only account administrators
	// currently have permission to change an attribute in this list. Will be empty if no attributes are frozen for this
	// assignment. Possible frozen attributes are: title, description, lock_at, points_possible, grading_type,
	// submission_types, assignment_group_id, allowed_extensions, group_category_id, notify_of_update, peer_reviews
	// NOTE: This field will only be present if the AssignmentFreezer plugin is available for your account.
	FrozenAttributes []string `json:"frozen_attributes"`
	// GradeGroupStudentsIndividually field: If this is a group assignment, boolean flag indicating whether or not
	// students will be graded individually.
	GradeGroupStudentsIndividually bool `json:"grade_group_students_individually"`
	// GraderCommentsVisibleToGraders field: Boolean indicating if provisional graders' comments are visible to other
	// provisional graders. Only relevant for moderated assignments.
	GraderCommentsVisibleToGraders bool `json:"grader_comments_visible_to_graders"`
	// GraderCount field: The maximum number of provisional graders who may issue grades for this assignment. Only
	// relevant for moderated assignments. Must be a positive value, and must be set to 1 if the course has fewer than
	// two active instructors. Otherwise, the maximum value is the number of active instructors in the course minus one,
	// or 10 if the course has more than 11 active instructors.
	GraderCount int `json:"grader_count"`
	// GraderNamesVisibleToFinalGrader field: Boolean indicating if provisional grader identities are visible to the
	// final grader. Only relevant for moderated assignments.
	GraderNamesVisibleToFinalGrader bool `json:"grader_names_visible_to_final_grader"`
	// GradersAnonymousToGraders field: Boolean indicating if provisional graders' identities are hidden from other
	// provisional graders. Only relevant for moderated assignments with grader_comments_visible_to_graders set to true.
	GradersAnonymousToGraders bool `json:"graders_anonymous_to_graders"`
	// GradingStandardID field: The id of the grading standard being applied to this assignment. Valid if grading_type
	// is 'letter_grade' or 'gpa_scale'.
	GradingStandardID int `json:"grading_standard_id"`
	// GradingType field: The type of grading the assignment receives; one of 'pass_fail', 'percent', 'letter_grade',
	// 'gpa_scale', 'points'
	GradingType *AssignmentGradingType `json:"grading_type"`
	// GroupCategoryID field: The ID of the assignment’s group set, if this is a group assignment. For group
	// discussions, set group_category_id on the discussion topic, not the linked assignment.
	GroupCategoryID int `json:"group_category_id"`
	// HasOverrides field: whether this assignment has overrides
	HasOverrides bool `json:"has_overrides"`
	// HasSubmittedSubmissions field: If true, the assignment has been submitted to by at least one student
	HasSubmittedSubmissions bool `json:"has_submitted_submissions"`
	// HTMLURL field: the URL to the assignment's web page
	HTMLURL string `json:"html_url"`
	// ID field: the ID of the assignment
	ID int `json:"id"`
	// IntegrationData field: (optional, Third Party integration data for assignment)
	IntegrationData map[interface{}]interface{} `json:"integration_data"`
	// IntegrationID field: (optional, Third Party unique identifier for Assignment)
	IntegrationID string `json:"integration_id"`
	// IntraGroupPeerReviews field: Boolean representing whether or not members from within the same group on a group
	// assignment can be assigned to peer review their own group's work
	IntraGroupPeerReviews bool `json:"intra_group_peer_reviews"`
	// LockAt field: the lock date (assignment is locked after this date). returns null if not present. NOTE: If this
	// assignment has assignment overrides, this field will be the lock date as it applies to the user requesting
	// information from the API.
	LockAt time.Time `json:"lock_at"`
	// LockExplanation field: (Optional) An explanation of why this is locked for the user. Present when locked_for_user
	// is true.
	LockExplanation string `json:"lock_explanation"`
	// LockInfo field: (Optional) Information for the user about the lock. Present when locked_for_user is true.
	LockInfo *LockInfo `json:"lock_info"`
	// LockedForUser field: Whether or not this is locked for the user.
	LockedForUser bool `json:"locked_for_user"`
	// MaxNameLength field: An integer indicating the maximum length an assignment's name may be
	MaxNameLength int `json:"max_name_length"`
	// ModeratedGrading field: Boolean indicating if the assignment is moderated.
	ModeratedGrading bool `json:"moderated_grading"`
	// Muted field: For courses using Old Gradebook, indicates whether the assignment is muted. For courses using New
	// Gradebook, true if the assignment has any unposted submissions, otherwise false. To see the posted status of
	// submissions, check the 'posted_attribute' on Submission.
	Muted bool `json:"muted"`
	// Name field: the name of the assignment
	Name string `json:"name"`
	// NeedsGradingCount field: if the requesting user has grading rights, the number of submissions that need grading.
	NeedsGradingCount int `json:"needs_grading_count"`
	// NeedsGradingCountBySection field: if the requesting user has grading rights and the
	// 'needs_grading_count_by_section' flag is specified, the number of submissions that need grading split out by
	// section. NOTE: This key is NOT present unless you pass the 'needs_grading_count_by_section' argument as true.
	// ANOTHER NOTE: it's possible to be enrolled in multiple sections, and if a student is setup that way they will
	// show an assignment that needs grading in multiple sections (effectively the count will be duplicated between
	// sections)
	NeedsGradingCountBySection []NeedsGradingCount `json:"needs_grading_count_by_section"`
	// OmitFromFinalGrade field: (Optional) If true, the assignment will be omitted from the student's final grade
	OmitFromFinalGrade bool `json:"omit_from_final_grade"`
	// OnlyVisibleToOverrides field: Whether the assignment is only visible to overrides.
	OnlyVisibleToOverrides bool `json:"only_visible_to_overrides"`
	// Overrides field: (Optional) If 'overrides' is included in the 'include' parameter, includes an array of
	// assignment override objects.
	Overrides []AssignmentOverride `json:"overrides"`
	// PeerReviewCount field: Integer representing the amount of reviews each user is assigned. NOTE: This key is NOT
	// present unless you have automatic_peer_reviews set to true.
	PeerReviewCount int `json:"peer_review_count"`
	// PeerReviews field: Boolean indicating if peer reviews are required for this assignment
	PeerReviews bool `json:"peer_reviews"`
	// PeerReviewsAssignAt field: String representing a date the reviews are due by. Must be a date that occurs after
	// the default due date. If blank, or date is not after the assignment's due date, the assignment's due date will be
	// used. NOTE: This key is NOT present unless you have automatic_peer_reviews set to true.
	PeerReviewsAssignAt time.Time `json:"peer_reviews_assign_at"`
	// PointsPossible field: the maximum points possible for the assignment
	PointsPossible float64 `json:"points_possible"`
	// Position field: the sorting order of the assignment in the group
	Position int `json:"position"`
	// PostManually field: Whether the assignment has manual posting enabled. Only relevant for courses using New
	// Gradebook.
	PostManually bool `json:"post_manually"`
	// PostToSis field: (optional, present if Sync Grades to SIS feature is enabled)
	PostToSis bool `json:"post_to_sis"`
	// Published field: Whether the assignment is published
	Published bool `json:"published"`
	// QuizID field: (Optional) id of the associated quiz (applies only when submission_types is ['online_quiz'])
	QuizID int `json:"quiz_id"`
	// Rubric field: (Optional) A list of scoring criteria and ratings for each rubric criterion. Included if there is
	// an associated rubric.
	Rubric []RubricCriteria `json:"rubric"`
	// RubricSettings field: (Optional) An object describing the basic attributes of the rubric, including the point
	// total. Included if there is an associated rubric.
	RubricSettings string `json:"rubric_settings"`
	// Submission field: (Optional) If 'submission' is included in the 'include' parameter, includes a Submission object
	// that represents the current user's (user who is requesting information from the api) current submission for the
	// assignment. See the Submissions API for an example response. If the user does not have a submission, this key
	// will be absent.
	Submission *Submission `json:"submission"`
	// SubmissionTypes field: the types of submissions allowed for this assignment list containing one or more of the
	// following: 'discussion_topic', 'online_quiz', 'on_paper', 'none', 'external_tool', 'online_text_entry',
	// 'online_url', 'online_upload' 'media_recording'
	SubmissionTypes *AssignmentSubmissionTypes `json:"submission_types"`
	// SubmissionsDownloadURL field: the URL to download all submissions as a zip
	SubmissionsDownloadURL string `json:"submissions_download_url"`
	// TurnitinEnabled field: Boolean flag indicating whether or not Turnitin has been enabled for the assignment. NOTE:
	// This flag will not appear unless your account has the Turnitin plugin available
	TurnitinEnabled bool `json:"turnitin_enabled"`
	// TurnitinSettings field: Settings to pass along to turnitin to control what kinds of matches should be considered.
	// originality_report_visibility can be 'immediate', 'after_grading', 'after_due_date', or 'never'
	// exclude_small_matches_type can be null, 'percent', 'words' exclude_small_matches_value: - if type is null, this
	// will be null also - if type is 'percent', this will be a number between 0 and 100 representing match size to
	// exclude as a percentage of the document size. - if type is 'words', this will be number > 0 representing how many
	// words a match must contain for it to be considered NOTE: This flag will not appear unless your account has the
	// Turnitin plugin available
	TurnitinSettings *TurnitinSettings `json:"turnitin_settings"`
	// UnlockAt field: the unlock date (assignment is unlocked after this date) returns null if not present NOTE: If
	// this assignment has assignment overrides, this field will be the unlock date as it applies to the user requesting
	// information from the API.
	UnlockAt time.Time `json:"unlock_at"`
	// Unpublishable field: Whether the assignment's 'published' state can be changed to false. Will be false if there
	// are student submissions for the assignment.
	Unpublishable bool `json:"unpublishable"`
	// UpdatedAt field: The time at which this assignment was last modified in any way
	UpdatedAt time.Time `json:"updated_at"`
	// UseRubricForGrading field: (Optional) If true, the rubric is directly tied to grading the assignment. Otherwise,
	// it is only advisory. Included if there is an associated rubric.
	UseRubricForGrading bool `json:"use_rubric_for_grading"`
	// VericiteEnabled field: Boolean flag indicating whether or not VeriCite has been enabled for the assignment. NOTE:
	// This flag will not appear unless your account has the VeriCite plugin available
	VericiteEnabled bool `json:"vericite_enabled"`
}

// AuthenticationEvent model object
type AuthenticationEvent struct {
	// AccountID field: ID of the account associated with the event. will match the account_id in the associated
	// pseudonym.
	AccountID int `json:"account_id"`
	// CreatedAt field: timestamp of the event
	CreatedAt time.Time `json:"created_at"`
	// EventType field: authentication event type ('login' or 'logout')
	EventType *AuthenticationEventEventType `json:"event_type"`
	// PseudonymID field: ID of the pseudonym (login) associated with the event
	PseudonymID int `json:"pseudonym_id"`
	// UserID field: ID of the user associated with the event will match the user_id in the associated pseudonym.
	UserID int `json:"user_id"`
}

// AuthenticationProvider model object
type AuthenticationProvider struct {
	// AuthBase field: Valid for LDAP and CAS providers.
	AuthBase string `json:"auth_base"`
	// AuthFilter field: Valid for LDAP providers.
	AuthFilter string `json:"auth_filter"`
	// AuthHost field: Valid for LDAP providers.
	AuthHost string `json:"auth_host"`
	// AuthOverTLS field: Valid for LDAP providers.
	AuthOverTLS int `json:"auth_over_tls"`
	// AuthPort field: Valid for LDAP providers.
	AuthPort int `json:"auth_port"`
	// AuthType field: Valid for all providers.
	AuthType string `json:"auth_type"`
	// AuthUsername field: Valid for LDAP providers.
	AuthUsername string `json:"auth_username"`
	// CertificateFingerprint field: Valid for SAML providers.
	CertificateFingerprint string `json:"certificate_fingerprint"`
	// FederatedAttributes field
	FederatedAttributes *FederatedAttributesConfig `json:"federated_attributes"`
	// ID field: Valid for all providers.
	ID int `json:"id"`
	// IdentifierFormat field: Valid for SAML providers.
	IdentifierFormat string `json:"identifier_format"`
	// IdpEntityID field: Valid for SAML providers.
	IdpEntityID string `json:"idp_entity_id"`
	// JitProvisioning field: Just In Time provisioning. Valid for all providers except Canvas (which has the similar in
	// concept self_registration setting).
	JitProvisioning bool `json:"jit_provisioning"`
	// LogInURL field: Valid for SAML and CAS providers.
	LogInURL string `json:"log_in_url"`
	// LogOutURL field: Valid for SAML providers.
	LogOutURL string `json:"log_out_url"`
	// LoginAttribute field: Valid for SAML providers.
	LoginAttribute string `json:"login_attribute"`
	// Position field: Valid for all providers.
	Position int `json:"position"`
	// RequestedAuthnContext field: Valid for SAML providers.
	RequestedAuthnContext string `json:"requested_authn_context"`
	// SigAlg field: Valid for SAML providers.
	SigAlg string `json:"sig_alg"`
}

// SSOSettings model object: Settings that are applicable across an account's authentication configuration, even if
// there are multiple individual providers
type SSOSettings struct {
	// AuthDiscoveryURL field: If a discovery url is set, canvas will forward all users to that URL when they need to be
	// authenticated. That page will need to then help the user figure out where they need to go to log in. If no
	// discovery url is configured, the first configuration will be used to attempt to authenticate the user.
	AuthDiscoveryURL string `json:"auth_discovery_url"`
	// ChangePasswordURL field: The url to redirect users to for password resets. Leave blank for default Canvas
	// behavior
	ChangePasswordURL string `json:"change_password_url"`
	// LoginHandleName field: The label used for unique login identifiers.
	LoginHandleName string `json:"login_handle_name"`
	// UnknownUserURL field: If an unknown user url is set, Canvas will forward to that url when a service authenticates
	// a user, but that user does not exist in Canvas. The default behavior is to present an error.
	UnknownUserURL string `json:"unknown_user_url"`
}

// FederatedAttributesConfig model object: A mapping of Canvas attribute names to attribute names that a provider may
// send, in order to update the value of these attributes when a user logs in. The values can be a
// FederatedAttributeConfig, or a raw string corresponding to the "attribute" property of a FederatedAttributeConfig. In
// responses, full FederatedAttributeConfig objects are returned if JIT provisioning is enabled, otherwise just the
// attribute names are returned.
type FederatedAttributesConfig struct {
	// AdminRoles field: A comma separated list of role names to grant to the user. Note that these only apply at the
	// root account level, and not sub-accounts. If the attribute is not marked for provisioning only, the user will
	// also be removed from any other roles they currently hold that are not still specified by the IdP.
	AdminRoles string `json:"admin_roles"`
	// DisplayName field: The full display name of the user
	DisplayName string `json:"display_name"`
	// Email field: The user's e-mail address
	Email string `json:"email"`
	// GivenName field: The first, or given, name of the user
	GivenName string `json:"given_name"`
	// IntegrationID field: The secondary unique identifier for SIS purposes
	IntegrationID string `json:"integration_id"`
	// Locale field: The user's preferred locale/language
	Locale string `json:"locale"`
	// Name field: The full name of the user
	Name string `json:"name"`
	// SisUserID field: The unique SIS identifier
	SisUserID string `json:"sis_user_id"`
	// SortableName field: The full name of the user for sorting purposes
	SortableName string `json:"sortable_name"`
	// Surname field: The surname, or last name, of the user
	Surname string `json:"surname"`
	// Timezone field: The user's preferred time zone
	Timezone string `json:"timezone"`
}

// FederatedAttributeConfig model object: A single attribute name to be federated when a user logs in
type FederatedAttributeConfig struct {
	// Attribute field: The name of the attribute as it will be sent from the authentication provider
	Attribute string `json:"attribute"`
	// ProvisioningOnly field: If the attribute should be applied only when provisioning a new user, rather than all
	// logins
	ProvisioningOnly bool `json:"provisioning_only"`
}

// CalendarEvent model object
type CalendarEvent struct {
	// AllContextCodes field: a comma-separated list of all calendar contexts this event is part of
	AllContextCodes string `json:"all_context_codes"`
	// AllDay field: Boolean indicating whether this is an all-day event (midnight to midnight)
	AllDay bool `json:"all_day"`
	// AllDayDate field: The date of this event
	AllDayDate time.Time `json:"all_day_date"`
	// AppointmentGroupID field: Various Appointment-Group-related fields.These fields are only pertinent to time slots
	// (appointments) and reservations of those time slots. See the Appointment Groups API. The id of the appointment
	// group
	AppointmentGroupID int `json:"appointment_group_id"`
	// AppointmentGroupURL field: The API URL of the appointment group
	AppointmentGroupURL string `json:"appointment_group_url"`
	// AvailableSlots field: If the event is a time slot and it has a participant limit, an integer indicating how many
	// slots are available
	AvailableSlots int `json:"available_slots"`
	// ChildEvents field: Included by default, but may be excluded (see include[] option). If this is a time slot (see
	// the Appointment Groups API) this will be a list of any reservations. If this is a course-level event, this will
	// be a list of section-level events (if any)
	ChildEvents []int `json:"child_events"`
	// ChildEventsCount field: The number of child_events. See child_events (and parent_event_id)
	ChildEventsCount int `json:"child_events_count"`
	// ContextCode field: the context code of the calendar this event belongs to (course, user or group)
	ContextCode string `json:"context_code"`
	// CreatedAt field: When the calendar event was created
	CreatedAt time.Time `json:"created_at"`
	// Description field: The HTML description of the event
	Description string `json:"description"`
	// EffectiveContextCode field: if specified, it indicates which calendar this event should be displayed on. for
	// example, a section-level event would have the course's context code here, while the section's context code would
	// be returned above)
	EffectiveContextCode string `json:"effective_context_code"`
	// EndAt field: The end timestamp of the event
	EndAt time.Time `json:"end_at"`
	// Group field: If the event is a group-level reservation, this will contain the group participant JSON (refer to
	// the Groups API).
	Group string `json:"group"`
	// Hidden field: Whether this event should be displayed on the calendar. Only true for course-level events with
	// section-level child events.
	Hidden bool `json:"hidden"`
	// HTMLURL field: URL for a user to view this event
	HTMLURL string `json:"html_url"`
	// ID field: The ID of the calendar event
	ID int `json:"id"`
	// LocationAddress field: The address where the event is taking place
	LocationAddress string `json:"location_address"`
	// LocationName field: The location name of the event
	LocationName string `json:"location_name"`
	// OwnReservation field: If the event is a reservation, this a boolean indicating whether it is the current user's
	// reservation, or someone else's
	OwnReservation bool `json:"own_reservation"`
	// ParentEventID field: Normally null. If this is a reservation (see the Appointment Groups API), the id will
	// indicate the time slot it is for. If this is a section-level event, this will be the course-level parent event.
	ParentEventID int `json:"parent_event_id"`
	// ParticipantType field: The type of participant to sign up for a slot: 'User' or 'Group'
	ParticipantType string `json:"participant_type"`
	// ParticipantsPerAppointment field: If the event is a time slot, this is the participant limit
	ParticipantsPerAppointment int `json:"participants_per_appointment"`
	// ReserveURL field: If the event is a time slot, the API URL for reserving it
	ReserveURL string `json:"reserve_url"`
	// Reserved field: If the event is a time slot, a boolean indicating whether the user has already made a reservation
	// for it
	Reserved bool `json:"reserved"`
	// StartAt field: The start timestamp of the event
	StartAt time.Time `json:"start_at"`
	// Title field: The title of the calendar event
	Title string `json:"title"`
	// UpdatedAt field: When the calendar event was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// URL field: URL for this calendar event (to update, delete, etc.)
	URL string `json:"url"`
	// User field: If the event is a user-level reservation, this will contain the user participant JSON (refer to the
	// Users API).
	User string `json:"user"`
	// WorkflowState field: Current state of the event ('active', 'locked' or 'deleted') 'locked' indicates that
	// start_at/end_at cannot be changed (though the event could be deleted). Normally only reservations or time slots
	// with reservations are locked (see the Appointment Groups API)
	WorkflowState string `json:"workflow_state"`
}

// AssignmentEvent model object
type AssignmentEvent struct {
	// AllDay field: Boolean indicating whether this is an all-day event (e.g. assignment due at midnight)
	AllDay bool `json:"all_day"`
	// AllDayDate field: The due date of this assignment
	AllDayDate time.Time `json:"all_day_date"`
	// Assignment field: The full assignment JSON data (See the Assignments API)
	Assignment *Assignment `json:"assignment"`
	// AssignmentOverrides field: The list of AssignmentOverrides that apply to this event (See the Assignments API).
	// This information is useful for determining which students or sections this assignment-due event applies to.
	AssignmentOverrides *AssignmentOverride `json:"assignment_overrides"`
	// ContextCode field: the context code of the (course) calendar this assignment belongs to
	ContextCode string `json:"context_code"`
	// CreatedAt field: When the assignment was created
	CreatedAt time.Time `json:"created_at"`
	// Description field: The HTML description of the assignment
	Description string `json:"description"`
	// EndAt field: The due_at timestamp of the assignment
	EndAt time.Time `json:"end_at"`
	// HTMLURL field: URL for a user to view this assignment
	HTMLURL string `json:"html_url"`
	// ID field: A synthetic ID for the assignment
	ID string `json:"id"`
	// StartAt field: The due_at timestamp of the assignment
	StartAt time.Time `json:"start_at"`
	// Title field: The title of the assignment
	Title string `json:"title"`
	// UpdatedAt field: When the assignment was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// URL field: URL for this assignment (note that updating/deleting should be done via the Assignments API)
	URL string `json:"url"`
	// WorkflowState field: Current state of the assignment ('published' or 'deleted')
	WorkflowState *AssignmentEventWorkflowState `json:"workflow_state"`
}

// Collaboration model object
type Collaboration struct {
	// CollaborationType field: A name for the type of collaboration
	CollaborationType string `json:"collaboration_type"`
	// ContextID field: The canvas id of the course or group to which the collaboration belongs
	ContextID int `json:"context_id"`
	// ContextType field: The canvas type of the course or group to which the collaboration belongs
	ContextType string `json:"context_type"`
	// CreatedAt field: The timestamp when the collaboration was created
	CreatedAt time.Time `json:"created_at"`
	// Description field
	Description string `json:"description"`
	// DocumentID field: The collaboration document identifier for the collaboration provider
	DocumentID string `json:"document_id"`
	// ID field: The unique identifier for the collaboration
	ID int `json:"id"`
	// Title field
	Title string `json:"title"`
	// Type field: Another representation of the collaboration type
	Type string `json:"type"`
	// UpdateURL field: The LTI launch url to edit the collaboration
	UpdateURL string `json:"update_url"`
	// UpdatedAt field: The timestamp when the collaboration was last modified
	UpdatedAt time.Time `json:"updated_at"`
	// URL field: The LTI launch url to view collaboration.
	URL string `json:"url"`
	// UserID field: The canvas id of the user who created the collaboration
	UserID int `json:"user_id"`
	// UserName field: The name of the user who owns the collaboration
	UserName string `json:"user_name"`
}

// Collaborator model object
type Collaborator struct {
	// ID field: The unique user or group identifier for the collaborator.
	ID int `json:"id"`
	// Name field: The name of the collaborator.
	Name string `json:"name"`
	// Type field: The type of collaborator (e.g. 'user' or 'group').
	Type *CollaboratorType `json:"type"`
}

// CommMessage model object
type CommMessage struct {
	// Body field: The plain text body of the message
	Body string `json:"body"`
	// CreatedAt field: The date and time this message was created
	CreatedAt time.Time `json:"created_at"`
	// From field: The address that was put in the 'from' field of the message
	From string `json:"from"`
	// FromName field: The display name for the from address
	FromName string `json:"from_name"`
	// HTMLBody field: The HTML body of the message.
	HTMLBody string `json:"html_body"`
	// ID field: The ID of the CommMessage.
	ID int `json:"id"`
	// ReplyTo field: The reply_to header of the message
	ReplyTo string `json:"reply_to"`
	// SentAt field: The date and time this message was sent
	SentAt time.Time `json:"sent_at"`
	// Subject field: The message subject
	Subject string `json:"subject"`
	// To field: The address the message was sent to:
	To string `json:"to"`
	// WorkflowState field: The workflow state of the message. One of 'created', 'staged', 'sending', 'sent', 'bounced',
	// 'dashboard', 'cancelled', or 'closed'
	WorkflowState *CommMessageWorkflowState `json:"workflow_state"`
}

// CommunicationChannel model object
type CommunicationChannel struct {
	// Address field: The address, or path, of the communication channel.
	Address string `json:"address"`
	// ID field: The ID of the communication channel.
	ID int `json:"id"`
	// Position field: The position of this communication channel relative to the user's other channels when they are
	// ordered.
	Position int `json:"position"`
	// Type field: The type of communcation channel being described. Possible values are: 'email', 'push', 'sms', or
	// 'twitter'. This field determines the type of value seen in 'address'.
	Type *CommunicationChannelType `json:"type"`
	// UserID field: The ID of the user that owns this communication channel.
	UserID int `json:"user_id"`
	// WorkflowState field: The current state of the communication channel. Possible values are: 'unconfirmed' or
	// 'active'.
	WorkflowState *CommunicationChannelWorkflowState `json:"workflow_state"`
}

// ConferenceRecording model object
type ConferenceRecording struct {
	// CreatedAt field
	CreatedAt time.Time `json:"created_at"`
	// DurationMinutes field
	DurationMinutes int `json:"duration_minutes"`
	// PlaybackURL field
	PlaybackURL string `json:"playback_url"`
	// Title field
	Title string `json:"title"`
	// UpdatedAt field
	UpdatedAt time.Time `json:"updated_at"`
}

// Conference model object
type Conference struct {
	// ConferenceKey field: The 3rd party's ID for the conference
	ConferenceKey string `json:"conference_key"`
	// ConferenceType field: The type of conference
	ConferenceType string `json:"conference_type"`
	// ContextID field: The ID of this conference's context.
	ContextID int `json:"context_id"`
	// ContextType field: The type of this conference's context, typically 'Course' or 'Group'.
	ContextType string `json:"context_type"`
	// Description field: The description for the conference
	Description string `json:"description"`
	// Duration field: The expected duration the conference is supposed to last
	Duration int `json:"duration"`
	// EndedAt field: The date that the conference ended at, null if it hasn't ended
	EndedAt time.Time `json:"ended_at"`
	// HasAdvancedSettings field: True if the conference type has advanced settings.
	HasAdvancedSettings bool `json:"has_advanced_settings"`
	// ID field: The id of the conference
	ID int `json:"id"`
	// JoinURL field: URL to join the conference, may be null if the conference type doesn't set it
	JoinURL string `json:"join_url"`
	// LongRunning field: If true the conference is long running and has no expected end time
	LongRunning bool `json:"long_running"`
	// Recordings field: A List of recordings for the conference
	Recordings []ConferenceRecording `json:"recordings"`
	// StartedAt field: The date the conference started at, null if it hasn't started
	StartedAt time.Time `json:"started_at"`
	// Title field: The title of the conference
	Title string `json:"title"`
	// URL field: URL for the conference, may be null if the conference type doesn't set it
	URL string `json:"url"`
	// UserSettings field: A collection of settings specific to the conference type
	UserSettings map[interface{}]interface{} `json:"user_settings"`
	// Users field: Array of user ids that are participants in the conference
	Users []int `json:"users"`
}

// ContentExport model object
type ContentExport struct {
	// Attachment field: attachment api object for the export package (not present before the export completes or after
	// it becomes unavailable for download.)
	Attachment *File `json:"attachment"`
	// CreatedAt field: the date and time this export was requested
	CreatedAt time.Time `json:"created_at"`
	// ExportType field: the type of content migration: 'common_cartridge' or 'qti'
	ExportType *ContentExportExportType `json:"export_type"`
	// ID field: the unique identifier for the export
	ID int `json:"id"`
	// ProgressURL field: The api endpoint for polling the current progress
	ProgressURL string `json:"progress_url"`
	// UserID field: The ID of the user who started the export
	UserID int `json:"user_id"`
	// WorkflowState field: Current state of the content migration: created exporting exported failed
	WorkflowState *ContentExportWorkflowState `json:"workflow_state"`
}

// ContentMigration model object
type ContentMigration struct {
	// Attachment field: attachment api object for the uploaded file may not be present for all migrations
	Attachment string `json:"attachment"`
	// FinishedAt field: timestamp
	FinishedAt time.Time `json:"finished_at"`
	// ID field: the unique identifier for the migration
	ID int `json:"id"`
	// MigrationIssuesURL field: API url to the content migration's issues
	MigrationIssuesURL string `json:"migration_issues_url"`
	// MigrationType field: the type of content migration
	MigrationType string `json:"migration_type"`
	// MigrationTypeTitle field: the name of the content migration type
	MigrationTypeTitle string `json:"migration_type_title"`
	// PreAttachment field: file uploading data, see {file:file_uploads.html File Upload Documentation} for file upload
	// workflow This works a little differently in that all the file data is in the pre_attachment hash if there is no
	// upload_url then there was an attachment pre-processing error, the error message will be in the message key This
	// data will only be here after a create or update call
	PreAttachment string `json:"pre_attachment"`
	// ProgressURL field: The api endpoint for polling the current progress
	ProgressURL string `json:"progress_url"`
	// StartedAt field: timestamp
	StartedAt time.Time `json:"started_at"`
	// UserID field: The user who started the migration
	UserID int `json:"user_id"`
	// WorkflowState field: Current state of the content migration: pre_processing, pre_processed, running,
	// waiting_for_select, completed, failed
	WorkflowState *ContentMigrationWorkflowState `json:"workflow_state"`
}

// Migrator model object
type Migrator struct {
	// Name field: Description of the package type expected
	Name string `json:"name"`
	// RequiredSettings field: A list of fields this system requires
	RequiredSettings []string `json:"required_settings"`
	// RequiresFileUpload field: Whether this endpoint requires a file upload
	RequiresFileUpload bool `json:"requires_file_upload"`
	// Type field: The value to pass to the create endpoint
	Type string `json:"type"`
}

// ContentShare model object: Content shared between users
type ContentShare struct {
	// ContentExport field: The content export record associated with this content share
	ContentExport *ContentExport `json:"content_export"`
	// ContentType field: The type of content that was shared. Can be assignment, discussion_topic, page, quiz, module,
	// or module_item.
	ContentType string `json:"content_type"`
	// CreatedAt field: The datetime the content was shared with this user.
	CreatedAt time.Time `json:"created_at"`
	// ID field: The id of the content share for the current user
	ID int `json:"id"`
	// Name field: The name of the shared content
	Name string `json:"name"`
	// ReadState field: Whether the recipient has viewed the content share.
	ReadState string `json:"read_state"`
	// Receivers field: An Array of users the content is shared with.  This field is provided only to senders; an empty
	// array will be returned for the receiving users.
	Receivers []map[interface{}]interface{} `json:"receivers"`
	// Sender field: The user who shared the content. This field is provided only to receivers; it is not populated in
	// the sender's list of sent content shares.
	Sender map[interface{}]interface{} `json:"sender"`
	// SourceCourse field: The course the content was originally shared from.
	SourceCourse map[interface{}]interface{} `json:"source_course"`
	// UpdatedAt field: The datetime the content was updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UserID field: The id of the user who sent or received the content share.
	UserID int `json:"user_id"`
}

// CompletionRequirement model object
type CompletionRequirement struct {
	// Completed field: whether the calling user has met this requirement (Optional; present only if the caller is a
	// student or if the optional parameter 'student_id' is included)
	Completed bool `json:"completed"`
	// MinScore field: minimum score required to complete (only present when type == 'min_score')
	MinScore int `json:"min_score"`
	// Type field: one of 'must_view', 'must_submit', 'must_contribute', 'min_score'
	Type *CompletionRequirementType `json:"type"`
}

// ContentDetails model object
type ContentDetails struct {
	// DueAt field
	DueAt time.Time `json:"due_at"`
	// LockAt field
	LockAt time.Time `json:"lock_at"`
	// LockExplanation field
	LockExplanation string `json:"lock_explanation"`
	// LockInfo field
	LockInfo *LockInfo `json:"lock_info"`
	// LockedForUser field
	LockedForUser bool `json:"locked_for_user"`
	// PointsPossible field
	PointsPossible int `json:"points_possible"`
	// UnlockAt field
	UnlockAt time.Time `json:"unlock_at"`
}

// ModuleItem model object
type ModuleItem struct {
	// CompletionRequirement field: Completion requirement for this module item
	CompletionRequirement *CompletionRequirement `json:"completion_requirement"`
	// ContentDetails field: (Present only if requested through include[]=content_details) If applicable, returns
	// additional details specific to the associated object
	ContentDetails *ContentDetails `json:"content_details"`
	// ContentID field: the id of the object referred to applies to 'File', 'Discussion', 'Assignment', 'Quiz',
	// 'ExternalTool' types
	ContentID int `json:"content_id"`
	// ExternalURL field: (only for 'ExternalUrl' and 'ExternalTool' types) external url that the item points to
	ExternalURL string `json:"external_url"`
	// HTMLURL field: link to the item in Canvas
	HTMLURL string `json:"html_url"`
	// ID field: the unique identifier for the module item
	ID int `json:"id"`
	// Indent field: 0-based indent level; module items may be indented to show a hierarchy
	Indent int `json:"indent"`
	// ModuleID field: the id of the Module this item appears in
	ModuleID int `json:"module_id"`
	// NewTab field: (only for 'ExternalTool' type) whether the external tool opens in a new tab
	NewTab bool `json:"new_tab"`
	// PageURL field: (only for 'Page' type) unique locator for the linked wiki page
	PageURL string `json:"page_url"`
	// Position field: the position of this item in the module (1-based)
	Position int `json:"position"`
	// Published field: (Optional) Whether this module item is published. This field is present only if the caller has
	// permission to view unpublished items.
	Published bool `json:"published"`
	// Title field: the title of this item
	Title string `json:"title"`
	// Type field: the type of object referred to one of 'File', 'Page', 'Discussion', 'Assignment', 'Quiz',
	// 'SubHeader', 'ExternalUrl', 'ExternalTool'
	Type *ModuleItemType `json:"type"`
	// URL field: (Optional) link to the Canvas API object, if applicable
	URL string `json:"url"`
}

// ModuleItemSequenceNode model object
type ModuleItemSequenceNode struct {
	// Current field: The ModuleItem being queried
	Current *ModuleItem `json:"current"`
	// MasteryPath field: The conditional release rule for the module item, if applicable
	MasteryPath map[interface{}]interface{} `json:"mastery_path"`
	// Next field: The next ModuleItem in the sequence
	Next *ModuleItem `json:"next"`
	// Prev field: The previous ModuleItem in the sequence
	Prev *ModuleItem `json:"prev"`
}

// ModuleItemSequence model object
type ModuleItemSequence struct {
	// Items field: an array containing one ModuleItemSequenceNode for each appearence of the asset in the module
	// sequence (up to 10 total)
	Items []ModuleItemSequenceNode `json:"items"`
	// Modules field: an array containing each Module referenced above
	Modules []Module `json:"modules"`
}

// Module model object
type Module struct {
	// CompletedAt field: the date the calling user completed the module (Optional; present only if the caller is a
	// student or if the optional parameter 'student_id' is included)
	CompletedAt time.Time `json:"completed_at"`
	// ID field: the unique identifier for the module
	ID int `json:"id"`
	// Items field: The contents of this module, as an array of Module Items. (Present only if requested via
	// include[]=items AND the module is not deemed too large by Canvas.)
	Items []ModuleItem `json:"items"`
	// ItemsCount field: The number of items in the module
	ItemsCount int `json:"items_count"`
	// ItemsURL field: The API URL to retrive this module's items
	ItemsURL string `json:"items_url"`
	// Name field: the name of this module
	Name string `json:"name"`
	// Position field: the position of this module in the course (1-based)
	Position int `json:"position"`
	// PrerequisiteModuleIds field: IDs of Modules that must be completed before this one is unlocked
	PrerequisiteModuleIds []int `json:"prerequisite_module_ids"`
	// PublishFinalGrade field: if the student's final grade for the course should be published to the SIS upon
	// completion of this module
	PublishFinalGrade bool `json:"publish_final_grade"`
	// Published field: (Optional) Whether this module is published. This field is present only if the caller has
	// permission to view unpublished modules.
	Published bool `json:"published"`
	// RequireSequentialProgress field: Whether module items must be unlocked in order
	RequireSequentialProgress bool `json:"require_sequential_progress"`
	// State field: The state of this Module for the calling user one of 'locked', 'unlocked', 'started', 'completed'
	// (Optional; present only if the caller is a student or if the optional parameter 'student_id' is included)
	State *ModuleState `json:"state"`
	// UnlockAt field: (Optional) the date this module will unlock
	UnlockAt time.Time `json:"unlock_at"`
	// WorkflowState field: the state of the module: 'active', 'deleted'
	WorkflowState *ModuleWorkflowState `json:"workflow_state"`
}

// Conversation model object
type Conversation struct {
	// Audience field: Array of user ids who are involved in the conversation, ordered by participation level, then
	// alphabetical. Excludes current user, unless this is a monologue.
	Audience []int `json:"audience"`
	// AudienceContexts field: Most relevant shared contexts (courses and groups) between current user and other
	// participants. If there is only one participant, it will also include that user's enrollment(s)/ membership
	// type(s) in each course/group.
	AudienceContexts []string `json:"audience_contexts"`
	// AvatarURL field: URL to appropriate icon for this conversation (custom, individual or group avatar, depending on
	// audience).
	AvatarURL string `json:"avatar_url"`
	// ContextName field: Name of the course or group in which the conversation is occurring.
	ContextName string `json:"context_name"`
	// ID field: the unique identifier for the conversation.
	ID int `json:"id"`
	// LastMessage field: A <=100 character preview from the most recent message.
	LastMessage string `json:"last_message"`
	// MessageCount field: the number of messages in the conversation.
	MessageCount int `json:"message_count"`
	// Participants field: Array of users participating in the conversation. Includes current user.
	Participants []ConversationParticipant `json:"participants"`
	// Private field: whether the conversation is private.
	Private bool `json:"private"`
	// Properties field: Additional conversation flags (last_author, attachments, media_objects). Each listed property
	// means the flag is set to true (i.e. the current user is the most recent author, there are attachments, or there
	// are media objects)
	Properties []string `json:"properties"`
	// Starred field: whether the conversation is starred.
	Starred bool `json:"starred"`
	// StartAt field: the date and time at which the last message was sent.
	StartAt time.Time `json:"start_at"`
	// Subject field: the subject of the conversation.
	Subject string `json:"subject"`
	// Subscribed field: whether the current user is subscribed to the conversation.
	Subscribed bool `json:"subscribed"`
	// Visible field: indicates whether the conversation is visible under the current scope and filter. This attribute
	// is always true in the index API response, and is primarily useful in create/update responses so that you can know
	// if the record should be displayed in the UI. The default scope is assumed, unless a scope or filter is passed to
	// the create/update API call.
	Visible bool `json:"visible"`
	// WorkflowState field: The current state of the conversation (read, unread or archived).
	WorkflowState string `json:"workflow_state"`
}

// ConversationParticipant model object
type ConversationParticipant struct {
	// AvatarURL field: If requested, this field will be included and contain a url to retrieve the user's avatar.
	AvatarURL string `json:"avatar_url"`
	// FullName field: The full name of the user.
	FullName string `json:"full_name"`
	// ID field: The user ID for the participant.
	ID int `json:"id"`
	// Name field: A short name the user has selected, for use in conversations or other less formal places through the
	// site.
	Name string `json:"name"`
}

// CourseEventLink model object
type CourseEventLink struct {
	// CopiedFrom field: ID of the course that this course was copied from. This is only included if the event_type is
	// copied_from.
	CopiedFrom int `json:"copied_from"`
	// CopiedTo field: ID of the course that this course was copied to. This is only included if the event_type is
	// copied_to.
	CopiedTo int `json:"copied_to"`
	// Course field: ID of the course for the event.
	Course int `json:"course"`
	// PageView field: ID of the page view during the event if it exists.
	PageView string `json:"page_view"`
	// SisBatch field: ID of the SIS batch that triggered the event.
	SisBatch int `json:"sis_batch"`
	// User field: ID of the user for the event (who made the change).
	User int `json:"user"`
}

// CourseEvent model object
type CourseEvent struct {
	// CreatedAt field: timestamp of the event
	CreatedAt time.Time `json:"created_at"`
	// EventData field: Course event data depending on the event type.  This will return an object containing the
	// relevant event data.  An updated event type will return an UpdatedEventData object.
	EventData string `json:"event_data"`
	// EventSource field: Course event source depending on the event type.  This will return a string containing the
	// source of the event.
	EventSource string `json:"event_source"`
	// EventType field: Course event type The event type defines the type and schema of the event_data object.
	EventType string `json:"event_type"`
	// ID field: ID of the event.
	ID string `json:"id"`
	// Links field: Jsonapi.org links
	Links *CourseEventLink `json:"links"`
}

// CreatedEventData model object: The created event data object returns all the fields that were set in the format of
// the following example.  If a field does not exist it was not set. The value of each field changed is in the format of
// [:old_value, :new_value].  The created event type also includes a created_source field to specify what triggered the
// creation of the course.
type CreatedEventData struct {
	// ConcludeAt field
	ConcludeAt []time.Time `json:"conclude_at"`
	// CreatedSource field: The type of action that triggered the creation of the course.
	CreatedSource string `json:"created_source"`
	// IsPublic field
	IsPublic []bool `json:"is_public"`
	// Name field
	Name []string `json:"name"`
	// StartAt field
	StartAt []time.Time `json:"start_at"`
}

// UpdatedEventData model object: The updated event data object returns all the fields that have changed in the format
// of the following example.  If a field does not exist it was not changed.  The value is an array that contains the
// before and after values for the change as in [:old_value, :new_value].
type UpdatedEventData struct {
	// ConcludeAt field
	ConcludeAt []time.Time `json:"conclude_at"`
	// IsPublic field
	IsPublic []bool `json:"is_public"`
	// Name field
	Name []string `json:"name"`
	// StartAt field
	StartAt []time.Time `json:"start_at"`
}

// CourseNickname model object
type CourseNickname struct {
	// CourseID field: the ID of the course
	CourseID int `json:"course_id"`
	// Name field: the actual name of the course
	Name string `json:"name"`
	// Nickname field: the calling user's nickname for the course
	Nickname string `json:"nickname"`
}

// Term model object
type Term struct {
	// EndAt field
	EndAt time.Time `json:"end_at"`
	// ID field
	ID int `json:"id"`
	// Name field
	Name string `json:"name"`
	// StartAt field
	StartAt time.Time `json:"start_at"`
}

// CourseProgress model object
type CourseProgress struct {
	// CompletedAt field: date the course was completed. null if the course has not been completed by this user
	CompletedAt time.Time `json:"completed_at"`
	// NextRequirementURL field: url to next module item that has an unmet requirement. null if the user has completed
	// the course or the current module does not require sequential progress
	NextRequirementURL string `json:"next_requirement_url"`
	// RequirementCompletedCount field: total number of requirements the user has completed from all modules
	RequirementCompletedCount int `json:"requirement_completed_count"`
	// RequirementCount field: total number of requirements from all modules
	RequirementCount int `json:"requirement_count"`
}

// Course model object
type Course struct {
	// AccessRestrictedByDate field: optional: this will be true if this user is currently prevented from viewing the
	// course because of date restriction settings
	AccessRestrictedByDate bool `json:"access_restricted_by_date"`
	// AccountID field: the account associated with the course
	AccountID int `json:"account_id"`
	// AllowStudentAssignmentEdits field
	AllowStudentAssignmentEdits bool `json:"allow_student_assignment_edits"`
	// AllowStudentForumAttachments field
	AllowStudentForumAttachments bool `json:"allow_student_forum_attachments"`
	// AllowWikiComments field
	AllowWikiComments bool `json:"allow_wiki_comments"`
	// ApplyAssignmentGroupWeights field: weight final grade based on assignment group percentages
	ApplyAssignmentGroupWeights bool `json:"apply_assignment_group_weights"`
	// Blueprint field: optional: whether the course is set as a Blueprint Course (blueprint fields require the
	// Blueprint Courses feature)
	Blueprint bool `json:"blueprint"`
	// BlueprintRestrictions field: optional: Set of restrictions applied to all locked course objects
	BlueprintRestrictions map[interface{}]interface{} `json:"blueprint_restrictions"`
	// BlueprintRestrictionsByObjectType field: optional: Sets of restrictions differentiated by object type applied to
	// locked course objects
	BlueprintRestrictionsByObjectType map[interface{}]interface{} `json:"blueprint_restrictions_by_object_type"`
	// Calendar field: course calendar
	Calendar *CalendarLink `json:"calendar"`
	// CourseCode field: the course code
	CourseCode string `json:"course_code"`
	// CourseFormat field
	CourseFormat string `json:"course_format"`
	// CourseProgress field: optional: information on progress through the course returned only if
	// include[]=course_progress
	CourseProgress *CourseProgress `json:"course_progress"`
	// CreatedAt field: the date the course was created.
	CreatedAt time.Time `json:"created_at"`
	// DefaultView field: the type of page that users will see when they first visit the course - 'feed': Recent
	// Activity Dashboard - 'wiki': Wiki Front Page - 'modules': Course Modules/Sections Page - 'assignments': Course
	// Assignments List - 'syllabus': Course Syllabus Page other types may be added in the future
	DefaultView *CourseDefaultView `json:"default_view"`
	// EndAt field: the end date for the course, if applicable
	EndAt time.Time `json:"end_at"`
	// EnrollmentTermID field: the enrollment term associated with the course
	EnrollmentTermID int `json:"enrollment_term_id"`
	// Enrollments field: A list of enrollments linking the current user to the course. for student enrollments, grading
	// information may be included if include[]=total_scores
	Enrollments []Enrollment `json:"enrollments"`
	// GradePassbackSetting field: the grade_passback_setting set on the course
	GradePassbackSetting string `json:"grade_passback_setting"`
	// GradingStandardID field: the grading standard associated with the course
	GradingStandardID int `json:"grading_standard_id"`
	// HideFinalGrades field
	HideFinalGrades bool `json:"hide_final_grades"`
	// ID field: the unique identifier for the course
	ID int `json:"id"`
	// IntegrationID field: the integration identifier for the course, if defined. This field is only included if the
	// user has permission to view SIS information.
	IntegrationID string `json:"integration_id"`
	// IsPublic field
	IsPublic bool `json:"is_public"`
	// IsPublicToAuthUsers field
	IsPublicToAuthUsers bool `json:"is_public_to_auth_users"`
	// License field
	License string `json:"license"`
	// Locale field: the course-set locale, if applicable
	Locale string `json:"locale"`
	// Name field: the full name of the course
	Name string `json:"name"`
	// NeedsGradingCount field: optional: the number of submissions needing grading returned only if the current user
	// has grading rights and include[]=needs_grading_count
	NeedsGradingCount int `json:"needs_grading_count"`
	// OpenEnrollment field
	OpenEnrollment bool `json:"open_enrollment"`
	// Permissions field: optional: the permissions the user has for the course. returned only for a single course and
	// include[]=permissions
	Permissions map[string]bool `json:"permissions"`
	// PublicDescription field: optional: the public description of the course
	PublicDescription string `json:"public_description"`
	// PublicSyllabus field
	PublicSyllabus bool `json:"public_syllabus"`
	// PublicSyllabusToAuth field
	PublicSyllabusToAuth bool `json:"public_syllabus_to_auth"`
	// RestrictEnrollmentsToCourseDates field
	RestrictEnrollmentsToCourseDates bool `json:"restrict_enrollments_to_course_dates"`
	// RootAccountID field: the root account associated with the course
	RootAccountID int `json:"root_account_id"`
	// SelfEnrollment field
	SelfEnrollment bool `json:"self_enrollment"`
	// SisCourseID field: the SIS identifier for the course, if defined. This field is only included if the user has
	// permission to view SIS information.
	SisCourseID string `json:"sis_course_id"`
	// SisImportID field: the unique identifier for the SIS import. This field is only included if the user has
	// permission to manage SIS information.
	SisImportID int `json:"sis_import_id"`
	// StartAt field: the start date for the course, if applicable
	StartAt time.Time `json:"start_at"`
	// StorageQuotaMb field
	StorageQuotaMb int `json:"storage_quota_mb"`
	// StorageQuotaUsedMb field
	StorageQuotaUsedMb float64 `json:"storage_quota_used_mb"`
	// SyllabusBody field: optional: user-generated HTML for the course syllabus
	SyllabusBody string `json:"syllabus_body"`
	// Term field: optional: the enrollment term object for the course returned only if include[]=term
	Term *Term `json:"term"`
	// TimeZone field: The course's IANA time zone name.
	TimeZone string `json:"time_zone"`
	// TotalStudents field: optional: the total number of active and invited students in the course
	TotalStudents int `json:"total_students"`
	// UUID field: the UUID of the course
	UUID string `json:"uuid"`
	// WorkflowState field: the current state of the course one of 'unpublished', 'available', 'completed', or 'deleted'
	WorkflowState *CourseWorkflowState `json:"workflow_state"`
}

// CalendarLink model object
type CalendarLink struct {
	// Ics field: The URL of the calendar in ICS format
	Ics string `json:"ics"`
}

// ColumnDatum model object: ColumnDatum objects contain the entry for a column for each user.
type ColumnDatum struct {
	// Content field
	Content string `json:"content"`
	// UserID field
	UserID int `json:"user_id"`
}

// CustomColumn model object
type CustomColumn struct {
	// Hidden field: won't be displayed if hidden is true
	Hidden bool `json:"hidden"`
	// ID field: The ID of the custom gradebook column
	ID int `json:"id"`
	// Position field: column order
	Position int `json:"position"`
	// ReadOnly field: won't be editable in the gradebook UI
	ReadOnly bool `json:"read_only"`
	// TeacherNotes field: When true, this column's visibility will be toggled in the Gradebook when a user selects to
	// show or hide notes
	TeacherNotes bool `json:"teacher_notes"`
	// Title field: header text
	Title string `json:"title"`
}

// DeveloperKeyAccountBinding model object
type DeveloperKeyAccountBinding struct {
	// AccountID field: The global Canvas ID of the account in the binding
	AccountID float64 `json:"account_id"`
	// AccountOwnsBinding field: True if the requested context owns the binding
	AccountOwnsBinding bool `json:"account_owns_binding"`
	// DeveloperKeyID field: The global Canvas ID of the developer key in the binding
	DeveloperKeyID float64 `json:"developer_key_id"`
	// ID field: The Canvas ID of the binding
	ID float64 `json:"id"`
	// WorkflowState field: The workflow state of the binding. Will be one of 'on', 'off', or 'allow.'
	WorkflowState float64 `json:"workflow_state"`
}

// FileAttachment model object: A file attachment
type FileAttachment struct {
	// ContentType field
	ContentType string `json:"content-type"`
	// DisplayName field
	DisplayName string `json:"display_name"`
	// Filename field
	Filename string `json:"filename"`
	// URL field
	URL string `json:"url"`
}

// DiscussionTopic model object: A discussion topic
type DiscussionTopic struct {
	// AllowRating field: Whether or not users can rate entries in this topic.
	AllowRating bool `json:"allow_rating"`
	// AssignmentID field: The unique identifier of the assignment if the topic is for grading, otherwise null.
	AssignmentID int `json:"assignment_id"`
	// Attachments field: Array of file attachments.
	Attachments []FileAttachment `json:"attachments"`
	// DelayedPostAt field: The datetime to publish the topic (if not right away).
	DelayedPostAt time.Time `json:"delayed_post_at"`
	// DiscussionSubentryCount field: The count of entries in the topic.
	DiscussionSubentryCount int `json:"discussion_subentry_count"`
	// DiscussionType field: The type of discussion. Values are 'side_comment', for discussions that only allow one
	// level of nested comments, and 'threaded' for fully threaded discussions.
	DiscussionType *DiscussionTopicDiscussionType `json:"discussion_type"`
	// GroupCategoryID field: The unique identifier of the group category if the topic is a group discussion, otherwise
	// null.
	GroupCategoryID int `json:"group_category_id"`
	// GroupTopicChildren field: An array of group discussions the user is a part of. Fields include: id, group_id
	GroupTopicChildren []map[interface{}]interface{} `json:"group_topic_children"`
	// HTMLURL field: The URL to the discussion topic in canvas.
	HTMLURL string `json:"html_url"`
	// ID field: The ID of this topic.
	ID int `json:"id"`
	// LastReplyAt field: The datetime for when the last reply was in the topic.
	LastReplyAt time.Time `json:"last_reply_at"`
	// LockAt field: The datetime to lock the topic (if ever).
	LockAt time.Time `json:"lock_at"`
	// LockExplanation field: (Optional) An explanation of why this is locked for the user. Present when locked_for_user
	// is true.
	LockExplanation string `json:"lock_explanation"`
	// LockInfo field: (Optional) Information for the user about the lock. Present when locked_for_user is true.
	LockInfo *LockInfo `json:"lock_info"`
	// Locked field: Whether or not the discussion is 'closed for comments'.
	Locked bool `json:"locked"`
	// LockedForUser field: Whether or not this is locked for the user.
	LockedForUser bool `json:"locked_for_user"`
	// Message field: The HTML content of the message body.
	Message string `json:"message"`
	// OnlyGradersCanRate field: Whether or not grade permissions are required to rate entries.
	OnlyGradersCanRate bool `json:"only_graders_can_rate"`
	// Permissions field: The current user's permissions on this topic.
	Permissions map[string]bool `json:"permissions"`
	// Pinned field: Whether or not the discussion has been 'pinned' by an instructor
	Pinned bool `json:"pinned"`
	// PodcastURL field: If the topic is a podcast topic this is the feed url for the current user.
	PodcastURL string `json:"podcast_url"`
	// PostedAt field: The datetime the topic was posted. If it is null it hasn't been posted yet. (see delayed_post_at)
	PostedAt time.Time `json:"posted_at"`
	// Published field: Whether this discussion topic is published (true) or draft state (false)
	Published bool `json:"published"`
	// ReadState field: The read_state of the topic for the current user, 'read' or 'unread'.
	ReadState *DiscussionTopicReadState `json:"read_state"`
	// RequireInitialPost field: If true then a user may not respond to other replies until that user has made an
	// initial reply. Defaults to false.
	RequireInitialPost bool `json:"require_initial_post"`
	// RootTopicID field: If the topic is for grading and a group assignment this will point to the original topic in
	// the course.
	RootTopicID int `json:"root_topic_id"`
	// SortByRating field: Whether or not entries should be sorted by rating.
	SortByRating bool `json:"sort_by_rating"`
	// Subscribed field: Whether or not the current user is subscribed to this topic.
	Subscribed bool `json:"subscribed"`
	// SubscriptionHold field: (Optional) Why the user cannot subscribe to this topic. Only one reason will be returned
	// even if multiple apply. Can be one of: 'initial_post_required': The user must post a reply first;
	// 'not_in_group_set': The user is not in the group set for this graded group discussion; 'not_in_group': The user
	// is not in this topic's group; 'topic_is_announcement': This topic is an announcement
	SubscriptionHold *DiscussionTopicSubscriptionHold `json:"subscription_hold"`
	// Title field: The topic title.
	Title string `json:"title"`
	// TopicChildren field: DEPRECATED An array of topic_ids for the group discussions the user is a part of.
	TopicChildren []int `json:"topic_children"`
	// UnreadCount field: The count of unread entries of this topic for the current user.
	UnreadCount int `json:"unread_count"`
	// UserCanSeePosts field: Whether or not posts in this topic are visible to the user.
	UserCanSeePosts bool `json:"user_can_see_posts"`
	// UserName field: The username of the topic creator.
	UserName string `json:"user_name"`
}

// Grade model object
type Grade struct {
	// CurrentGrade field: The user's current grade in the class. Only included if user has permissions to view this
	// grade.
	CurrentGrade string `json:"current_grade"`
	// CurrentScore field: The user's current score in the class. Only included if user has permissions to view this
	// score.
	CurrentScore string `json:"current_score"`
	// FinalGrade field: The user's final grade for the class. Only included if user has permissions to view this grade.
	FinalGrade string `json:"final_grade"`
	// FinalScore field: The user's final score for the class. Only included if user has permissions to view this score.
	FinalScore string `json:"final_score"`
	// HTMLURL field: The URL to the Canvas web UI page for the user's grades, if this is a student enrollment.
	HTMLURL string `json:"html_url"`
	// UnpostedCurrentGrade field: The user's current grade in the class including muted/unposted assignments. Only
	// included if user has permissions to view this grade, typically teachers, TAs, and admins.
	UnpostedCurrentGrade string `json:"unposted_current_grade"`
	// UnpostedCurrentScore field: The user's current score in the class including muted/unposted assignments. Only
	// included if user has permissions to view this score, typically teachers, TAs, and admins..
	UnpostedCurrentScore string `json:"unposted_current_score"`
	// UnpostedFinalGrade field: The user's final grade for the class including muted/unposted assignments. Only
	// included if user has permissions to view this grade, typically teachers, TAs, and admins..
	UnpostedFinalGrade string `json:"unposted_final_grade"`
	// UnpostedFinalScore field: The user's final score for the class including muted/unposted assignments. Only
	// included if user has permissions to view this score, typically teachers, TAs, and admins..
	UnpostedFinalScore string `json:"unposted_final_score"`
}

// Enrollment model object
type Enrollment struct {
	// AssociatedUserID field: The unique id of the associated user. Will be null unless type is ObserverEnrollment.
	AssociatedUserID int `json:"associated_user_id"`
	// CourseID field: The unique id of the course.
	CourseID int `json:"course_id"`
	// CourseIntegrationID field: The Course Integration ID in which the enrollment is associated. This field is only
	// included if the user has permission to view SIS information.
	CourseIntegrationID string `json:"course_integration_id"`
	// CourseSectionID field: The unique id of the user's section.
	CourseSectionID int `json:"course_section_id"`
	// CreatedAt field: The created time of the enrollment, in ISO8601 format.
	CreatedAt time.Time `json:"created_at"`
	// CurrentGradingPeriodID field: optional: The id of the currently active grading period, if one exists. If the
	// course the enrollment belongs to does not have grading periods, or if no currently active grading period exists,
	// the value will be null. (applies only to student enrollments, and only available in course endpoints)
	CurrentGradingPeriodID int `json:"current_grading_period_id"`
	// CurrentGradingPeriodTitle field: optional: The name of the currently active grading period, if one exists. If the
	// course the enrollment belongs to does not have grading periods, or if no currently active grading period exists,
	// the value will be null. (applies only to student enrollments, and only available in course endpoints)
	CurrentGradingPeriodTitle string `json:"current_grading_period_title"`
	// CurrentPeriodOverrideGrade field: The user's override grade for the current grading period.
	CurrentPeriodOverrideGrade string `json:"current_period_override_grade"`
	// CurrentPeriodOverrideScore field: The user's override score for the current grading period.
	CurrentPeriodOverrideScore float64 `json:"current_period_override_score"`
	// CurrentPeriodUnpostedCurrentGrade field: optional: The letter grade equivalent of
	// current_period_unposted_current_score, if available. Only included if user has permission to view this grade,
	// typically teachers, TAs, and admins. If the course the enrollment belongs to does not have grading periods, or if
	// no currently active grading period exists, the value will be null. (applies only to student enrollments, and only
	// available in course endpoints)
	CurrentPeriodUnpostedCurrentGrade string `json:"current_period_unposted_current_grade"`
	// CurrentPeriodUnpostedCurrentScore field: optional: The student's score in the course for the current grading
	// period, including muted/unposted assignments. Only included if user has permission to view this score, typically
	// teachers, TAs, and admins. If the course the enrollment belongs to does not have grading periods, or if no
	// currently active grading period exists, the value will be null. (applies only to student enrollments, and only
	// available in course endpoints)
	CurrentPeriodUnpostedCurrentScore float64 `json:"current_period_unposted_current_score"`
	// CurrentPeriodUnpostedFinalGrade field: optional: The letter grade equivalent of
	// current_period_unposted_final_score, if available. Only included if user has permission to view this grade,
	// typically teachers, TAs, and admins. If the course the enrollment belongs to does not have grading periods, or if
	// no currently active grading period exists, the value will be null. (applies only to student enrollments, and only
	// available in course endpoints)
	CurrentPeriodUnpostedFinalGrade string `json:"current_period_unposted_final_grade"`
	// CurrentPeriodUnpostedFinalScore field: optional: The student's score in the course for the current grading
	// period, including muted/unposted assignments and including ungraded assignments with a score of 0. Only included
	// if user has permission to view this score, typically teachers, TAs, and admins. If the course the enrollment
	// belongs to does not have grading periods, or if no currently active grading period exists, the value will be
	// null. (applies only to student enrollments, and only available in course endpoints)
	CurrentPeriodUnpostedFinalScore float64 `json:"current_period_unposted_final_score"`
	// EndAt field: The end time of the enrollment, in ISO8601 format.
	EndAt time.Time `json:"end_at"`
	// EnrollmentState field: The state of the user's enrollment in the course.
	EnrollmentState string `json:"enrollment_state"`
	// Grades field: The URL to the Canvas web UI page containing the grades associated with this enrollment.
	Grades *Grade `json:"grades"`
	// HasGradingPeriods field: optional: Indicates whether the course the enrollment belongs to has grading periods set
	// up. (applies only to student enrollments, and only available in course endpoints)
	HasGradingPeriods bool `json:"has_grading_periods"`
	// HTMLURL field: The URL to the Canvas web UI page for this course enrollment.
	HTMLURL string `json:"html_url"`
	// ID field: The ID of the enrollment.
	ID int `json:"id"`
	// LastActivityAt field: The last activity time of the user for the enrollment, in ISO8601 format.
	LastActivityAt time.Time `json:"last_activity_at"`
	// LastAttendedAt field: The last attended date of the user for the enrollment in a course, in ISO8601 format.
	LastAttendedAt time.Time `json:"last_attended_at"`
	// LimitPrivilegesToCourseSection field: User can only access his or her own course section.
	LimitPrivilegesToCourseSection bool `json:"limit_privileges_to_course_section"`
	// OverrideGrade field: The user's override grade for the course.
	OverrideGrade string `json:"override_grade"`
	// OverrideScore field: The user's override score for the course.
	OverrideScore float64 `json:"override_score"`
	// Role field: The enrollment role, for course-level permissions. This field will match `type` if the enrollment
	// role has not been customized.
	Role string `json:"role"`
	// RoleID field: The id of the enrollment role.
	RoleID int `json:"role_id"`
	// RootAccountID field: The unique id of the user's account.
	RootAccountID int `json:"root_account_id"`
	// SectionIntegrationID field: The Section Integration ID in which the enrollment is associated. This field is only
	// included if the user has permission to view SIS information.
	SectionIntegrationID string `json:"section_integration_id"`
	// SisAccountID field: The SIS Account ID in which the enrollment is associated. Only displayed if present. This
	// field is only included if the user has permission to view SIS information.
	SisAccountID string `json:"sis_account_id"`
	// SisCourseID field: The SIS Course ID in which the enrollment is associated. Only displayed if present. This field
	// is only included if the user has permission to view SIS information.
	SisCourseID string `json:"sis_course_id"`
	// SisImportID field: The unique identifier for the SIS import. This field is only included if the user has
	// permission to manage SIS information.
	SisImportID int `json:"sis_import_id"`
	// SisSectionID field: The SIS Section ID in which the enrollment is associated. Only displayed if present. This
	// field is only included if the user has permission to view SIS information.
	SisSectionID string `json:"sis_section_id"`
	// SisUserID field: The SIS User ID in which the enrollment is associated. Only displayed if present. This field is
	// only included if the user has permission to view SIS information.
	SisUserID string `json:"sis_user_id"`
	// StartAt field: The start time of the enrollment, in ISO8601 format.
	StartAt time.Time `json:"start_at"`
	// TotalActivityTime field: The total activity time of the user for the enrollment, in seconds.
	TotalActivityTime int `json:"total_activity_time"`
	// TotalsForAllGradingPeriodsOption field: optional: Indicates whether the course the enrollment belongs to has the
	// Display Totals for 'All Grading Periods' feature enabled. (applies only to student enrollments, and only
	// available in course endpoints)
	TotalsForAllGradingPeriodsOption bool `json:"totals_for_all_grading_periods_option"`
	// Type field: The enrollment type. One of 'StudentEnrollment', 'TeacherEnrollment', 'TaEnrollment',
	// 'DesignerEnrollment', 'ObserverEnrollment'.
	Type string `json:"type"`
	// UnpostedCurrentGrade field: The user's current grade in the class including muted/unposted assignments. Only
	// included if user has permissions to view this grade, typically teachers, TAs, and admins.
	UnpostedCurrentGrade string `json:"unposted_current_grade"`
	// UnpostedCurrentScore field: The user's current score in the class including muted/unposted assignments. Only
	// included if user has permissions to view this score, typically teachers, TAs, and admins..
	UnpostedCurrentScore string `json:"unposted_current_score"`
	// UnpostedFinalGrade field: The user's final grade for the class including muted/unposted assignments. Only
	// included if user has permissions to view this grade, typically teachers, TAs, and admins..
	UnpostedFinalGrade string `json:"unposted_final_grade"`
	// UnpostedFinalScore field: The user's final score for the class including muted/unposted assignments. Only
	// included if user has permissions to view this score, typically teachers, TAs, and admins..
	UnpostedFinalScore string `json:"unposted_final_score"`
	// UpdatedAt field: The updated time of the enrollment, in ISO8601 format.
	UpdatedAt time.Time `json:"updated_at"`
	// User field: A description of the user.
	User *User `json:"user"`
	// UserID field: The unique id of the user.
	UserID int `json:"user_id"`
}

// CourseEpubExport model object: Combination of a Course & EpubExport.
type CourseEpubExport struct {
	// EpubExport field: ePub export API object
	EpubExport *EpubExport `json:"epub_export"`
	// ID field: the unique identifier for the course
	ID int `json:"id"`
	// Name field: the name for the course
	Name string `json:"name"`
}

// EpubExport model object
type EpubExport struct {
	// Attachment field: attachment api object for the export ePub (not present until the export completes)
	Attachment *File `json:"attachment"`
	// CreatedAt field: the date and time this export was requested
	CreatedAt time.Time `json:"created_at"`
	// ID field: the unique identifier for the export
	ID int `json:"id"`
	// ProgressURL field: The api endpoint for polling the current progress
	ProgressURL string `json:"progress_url"`
	// UserID field: The ID of the user who started the export
	UserID int `json:"user_id"`
	// WorkflowState field: Current state of the ePub export: created exporting exported generating generated failed
	WorkflowState *EpubExportWorkflowState `json:"workflow_state"`
}

// ErrorReport model object: A collection of information around a specific notification of a problem
type ErrorReport struct {
	// Comments field: long form documentation of what was witnessed
	Comments string `json:"comments"`
	// ContextAssetString field: string describing the asset being interacted with at the time of error.  Formatted
	// '[type]_[id]'
	ContextAssetString string `json:"context_asset_string"`
	// Email field: the email address of the reporting user
	Email string `json:"email"`
	// Subject field: The users problem summary, like an email subject line
	Subject string `json:"subject"`
	// URL field: URL of the page on which the error was reported
	URL string `json:"url"`
	// UserPerceivedSeverity field: categorization of how bad the user thinks the problem is.  Should be one of
	// [just_a_comment, not_urgent, workaround_possible, blocks_what_i_need_to_do, extreme_critical_emergency].
	UserPerceivedSeverity string `json:"user_perceived_severity"`
	// UserRoles field: comma seperated list of roles the reporting user holds.  Can be one [student], or many
	// [teacher,admin]
	UserRoles string `json:"user_roles"`
}

// ExternalFeed model object
type ExternalFeed struct {
	// CreatedAt field: When this external feed was added to Canvas
	CreatedAt time.Time `json:"created_at"`
	// DisplayName field: The title of the feed, pulled from the feed itself. If the feed hasn't yet been pulled, a
	// temporary name will be synthesized based on the URL
	DisplayName string `json:"display_name"`
	// HeaderMatch field: If not null, only feed entries whose title contains this string will trigger new posts in
	// Canvas
	HeaderMatch string `json:"header_match"`
	// ID field: The ID of the feed
	ID int `json:"id"`
	// URL field: The HTTP/HTTPS URL to the feed
	URL string `json:"url"`
	// Verbosity field: The verbosity setting determines how much of the feed's content is imported into Canvas as part
	// of the posting. 'link_only' means that only the title and a link to the item. 'truncate' means that a summary of
	// the first portion of the item body will be used. 'full' means that the full item body will be used.
	Verbosity *ExternalFeedVerbosity `json:"verbosity"`
}

// Favorite model object
type Favorite struct {
	// ContextID field: The ID of the object the Favorite refers to
	ContextID int `json:"context_id"`
	// ContextType field: The type of the object the Favorite refers to (currently, only 'Course' is supported)
	ContextType *FavoriteContextType `json:"context_type"`
}

// Feature model object
type Feature struct {
	// AppliesTo field: The type of object the feature applies to (RootAccount, Account, Course, or User):
	//  RootAccount features may only be controlled by flags on root accounts.
	//  Account features may be controlled by flags on accounts and their parent accounts.
	//  Course features may be controlled by flags on courses and their parent accounts.
	//  User features may be controlled by flags on users and site admin only.
	AppliesTo *FeatureAppliesTo `json:"applies_to"`
	// Autoexpand field: Whether the details of the feature are autoexpanded on page load vs. the user clicking to
	// expand.
	Autoexpand bool `json:"autoexpand"`
	// Beta field: Whether the feature is a beta feature. If true, the feature may not be fully polished and may be
	// subject to change in the future.
	Beta bool `json:"beta"`
	// Development field: Whether the feature is in active development. Features in this state are only visible in test
	// and beta instances and are not yet available for production use.
	Development bool `json:"development"`
	// DisplayName field: The user-visible name of the feature
	DisplayName string `json:"display_name"`
	// EnableAt field: The date this feature will be globally enabled, or null if this is not planned. (This information
	// is subject to change.)
	EnableAt time.Time `json:"enable_at"`
	// Feature field: The symbolic name of the feature, used in FeatureFlags
	Feature string `json:"feature"`
	// FeatureFlag field: The FeatureFlag that applies to the caller
	FeatureFlag *FeatureFlag `json:"feature_flag"`
	// ReleaseNotesURL field: A URL to the release notes describing the feature
	ReleaseNotesURL string `json:"release_notes_url"`
	// RootOptIn field: If true, a feature that is 'allowed' globally will be 'off' by default in root accounts.
	// Otherwise, root accounts inherit the global 'allowed' setting, which allows sub-accounts and courses to turn
	// features on with no root account action.
	RootOptIn bool `json:"root_opt_in"`
}

// FeatureFlag model object
type FeatureFlag struct {
	// ContextID field: The id of the object to which this flag applies (This field is not present if this FeatureFlag
	// represents the global Canvas default)
	ContextID int `json:"context_id"`
	// ContextType field: The type of object to which this flag applies (Account, Course, or User). (This field is not
	// present if this FeatureFlag represents the global Canvas default)
	ContextType *FeatureFlagContextType `json:"context_type"`
	// Feature field: The feature this flag controls
	Feature string `json:"feature"`
	// Locked field: If set, this feature flag cannot be changed in the caller's context because the flag is set 'off'
	// or 'on' in a higher context
	Locked bool `json:"locked"`
	// State field: The policy for the feature at this context.  can be 'off', 'allowed', or 'on'.
	State *FeatureFlagState `json:"state"`
}

// File model object
type File struct {
	// ContentType field
	ContentType string `json:"content-type"`
	// CreatedAt field
	CreatedAt time.Time `json:"created_at"`
	// DisplayName field
	DisplayName string `json:"display_name"`
	// Filename field
	Filename string `json:"filename"`
	// FolderID field
	FolderID int `json:"folder_id"`
	// Hidden field
	Hidden bool `json:"hidden"`
	// HiddenForUser field
	HiddenForUser bool `json:"hidden_for_user"`
	// ID field
	ID int `json:"id"`
	// LockAt field
	LockAt time.Time `json:"lock_at"`
	// LockExplanation field
	LockExplanation string `json:"lock_explanation"`
	// LockInfo field
	LockInfo *LockInfo `json:"lock_info"`
	// Locked field
	Locked bool `json:"locked"`
	// LockedForUser field
	LockedForUser bool `json:"locked_for_user"`
	// MediaEntryID field: identifier for file in third-party transcoding service
	MediaEntryID string `json:"media_entry_id"`
	// MimeClass field: simplified content-type mapping
	MimeClass string `json:"mime_class"`
	// ModifiedAt field
	ModifiedAt time.Time `json:"modified_at"`
	// PreviewURL field: optional: url to the document preview. This url is specific to the user making the api call.
	// Only included in submission endpoints.
	PreviewURL string `json:"preview_url"`
	// Size field: file size in bytes
	Size int `json:"size"`
	// ThumbnailURL field
	ThumbnailURL string `json:"thumbnail_url"`
	// UnlockAt field
	UnlockAt time.Time `json:"unlock_at"`
	// UpdatedAt field
	UpdatedAt time.Time `json:"updated_at"`
	// URL field
	URL string `json:"url"`
	// UUID field
	UUID string `json:"uuid"`
}

// Folder model object
type Folder struct {
	// ContextID field
	ContextID int `json:"context_id"`
	// ContextType field
	ContextType string `json:"context_type"`
	// CreatedAt field
	CreatedAt time.Time `json:"created_at"`
	// FilesCount field
	FilesCount int `json:"files_count"`
	// FilesURL field
	FilesURL string `json:"files_url"`
	// FoldersCount field
	FoldersCount int `json:"folders_count"`
	// FoldersURL field
	FoldersURL string `json:"folders_url"`
	// ForSubmissions field: If true, indicates this is a read-only folder containing files submitted to assignments
	ForSubmissions bool `json:"for_submissions"`
	// FullName field
	FullName string `json:"full_name"`
	// Hidden field
	Hidden bool `json:"hidden"`
	// HiddenForUser field
	HiddenForUser bool `json:"hidden_for_user"`
	// ID field
	ID int `json:"id"`
	// LockAt field
	LockAt time.Time `json:"lock_at"`
	// Locked field
	Locked bool `json:"locked"`
	// LockedForUser field
	LockedForUser bool `json:"locked_for_user"`
	// Name field
	Name string `json:"name"`
	// ParentFolderID field
	ParentFolderID int `json:"parent_folder_id"`
	// Position field
	Position int `json:"position"`
	// UnlockAt field
	UnlockAt time.Time `json:"unlock_at"`
	// UpdatedAt field
	UpdatedAt time.Time `json:"updated_at"`
}

// GradeChangeEventLinks model object
type GradeChangeEventLinks struct {
	// Assignment field: ID of the assignment associated with the event
	Assignment int `json:"assignment"`
	// Course field: ID of the course associated with the event. will match the context_id in the associated assignment
	// if the context type for the assignment is a course
	Course int `json:"course"`
	// Grader field: ID of the grader associated with the event. will match the grader_id in the associated submission.
	Grader int `json:"grader"`
	// PageView field: ID of the page view during the event if it exists.
	PageView string `json:"page_view"`
	// Student field: ID of the student associated with the event. will match the user_id in the associated submission.
	Student int `json:"student"`
}

// GradeChangeEvent model object
type GradeChangeEvent struct {
	// CreatedAt field: timestamp of the event
	CreatedAt time.Time `json:"created_at"`
	// EventType field: GradeChange event type
	EventType string `json:"event_type"`
	// ExcusedAfter field: Boolean indicating whether the submission was excused after the change.
	ExcusedAfter bool `json:"excused_after"`
	// ExcusedBefore field: Boolean indicating whether the submission was excused before the change.
	ExcusedBefore bool `json:"excused_before"`
	// GradeAfter field: The grade after the change.
	GradeAfter string `json:"grade_after"`
	// GradeBefore field: The grade before the change.
	GradeBefore string `json:"grade_before"`
	// GradedAnonymously field: Boolean indicating whether the student name was visible when the grade was given. Could
	// be null if the grade change record was created before this feature existed.
	GradedAnonymously bool `json:"graded_anonymously"`
	// ID field: ID of the event.
	ID string `json:"id"`
	// Links field
	Links *GradeChangeEventLinks `json:"links"`
	// RequestID field: The unique request id of the request during the grade change.
	RequestID string `json:"request_id"`
	// VersionNumber field: Version Number of the grade change submission.
	VersionNumber string `json:"version_number"`
}

// Grader model object
type Grader struct {
	// Assignments field: the assignment groups for all submissions in this response that were graded by this user.  The
	// details are not nested inside here, but the fact that an assignment is present here means that the grader did
	// grade submissions for this assignment on the contextual date. You can use the id of a grader and of an assignment
	// to make another API call to find all submissions for a grader/assignment combination on a given date.
	Assignments []int `json:"assignments"`
	// ID field: the user_id of the user who graded the contained submissions
	ID int `json:"id"`
	// Name field: the name of the user who graded the contained submissions
	Name string `json:"name"`
}

// Day model object
type Day struct {
	// Date field: the date represented by this entry
	Date time.Time `json:"date"`
	// Graders field: an array of the graders who were responsible for the submissions in this response. the submissions
	// are grouped according to the person who graded them and the assignment they were submitted for.
	Graders int `json:"graders"`
}

// SubmissionVersion model object: A SubmissionVersion object contains all the fields that a Submission object does,
// plus additional fields prefixed with current_* new_* and previous_* described below.
type SubmissionVersion struct {
	// AssignmentID field: the id of the assignment this submissions is for
	AssignmentID int `json:"assignment_id"`
	// AssignmentName field: the name of the assignment this submission is for
	AssignmentName string `json:"assignment_name"`
	// Body field: the body text of the submission
	Body string `json:"body"`
	// CurrentGrade field: the most up to date grade for the current version of this submission
	CurrentGrade string `json:"current_grade"`
	// CurrentGradedAt field: the latest time stamp for the grading of this submission
	CurrentGradedAt time.Time `json:"current_graded_at"`
	// CurrentGrader field: the name of the most recent grader for this submission
	CurrentGrader string `json:"current_grader"`
	// GradeMatchesCurrentSubmission field: boolean indicating whether the grade is equal to the current submission
	// grade
	GradeMatchesCurrentSubmission bool `json:"grade_matches_current_submission"`
	// GradedAt field: time stamp for the grading of this version of the submission
	GradedAt time.Time `json:"graded_at"`
	// Grader field: the name of the user who graded this version of the submission
	Grader string `json:"grader"`
	// GraderID field: the user id of the user who graded this version of the submission
	GraderID int `json:"grader_id"`
	// ID field: the id of the submission of which this is a version
	ID int `json:"id"`
	// NewGrade field: the updated grade provided in this version of the submission
	NewGrade string `json:"new_grade"`
	// NewGradedAt field: the timestamp for the grading of this version of the submission (alias for graded_at)
	NewGradedAt time.Time `json:"new_graded_at"`
	// NewGrader field: alias for 'grader'
	NewGrader string `json:"new_grader"`
	// PreviousGrade field: the grade for the submission version immediately preceding this one
	PreviousGrade string `json:"previous_grade"`
	// PreviousGradedAt field: the timestamp for the grading of the submission version immediately preceding this one
	PreviousGradedAt time.Time `json:"previous_graded_at"`
	// PreviousGrader field: the name of the grader who graded the version of this submission immediately preceding this
	// one
	PreviousGrader string `json:"previous_grader"`
	// Score field: the score for this version of the submission
	Score int `json:"score"`
	// SubmissionType field: the type of submission
	SubmissionType string `json:"submission_type"`
	// URL field: the url of the submission, if there is one
	URL string `json:"url"`
	// UserID field: the user ID of the student who created this submission
	UserID int `json:"user_id"`
	// UserName field: the name of the student who created this submission
	UserName string `json:"user_name"`
	// WorkflowState field: the state of the submission at this version
	WorkflowState string `json:"workflow_state"`
}

// SubmissionHistory model object
type SubmissionHistory struct {
	// SubmissionID field: the id of the submission
	SubmissionID int `json:"submission_id"`
	// Versions field: an array of all the versions of this submission
	Versions []SubmissionVersion `json:"versions"`
}

// GradingPeriod model object
type GradingPeriod struct {
	// CloseDate field: Grades can only be changed before the close date of the grading period.
	CloseDate string `json:"close_date"`
	// EndDate field: The end date of the grading period.
	EndDate string `json:"end_date"`
	// ID field: The unique identifier for the grading period.
	ID int `json:"id"`
	// IsClosed field: If true, the grading period's close_date has passed.
	IsClosed bool `json:"is_closed"`
	// StartDate field: The start date of the grading period.
	StartDate string `json:"start_date"`
	// Title field: The title for the grading period.
	Title string `json:"title"`
	// Weight field: A weight value that contributes to the overall weight of a grading period set which is used to
	// calculate how much assignments in this period contribute to the total grade
	Weight int `json:"weight"`
}

// GradingSchemeEntry model object
type GradingSchemeEntry struct {
	// Name field: The name for an entry value within a GradingStandard that describes the range of the value
	Name string `json:"name"`
	// Value field: The value for the name of the entry within a GradingStandard.  The entry represents the lower bound
	// of the range for the entry. This range includes the value up to the next entry in the GradingStandard, or 100 if
	// there is no upper bound. The lowest value will have a lower bound range of 0.
	Value int `json:"value"`
}

// GradingStandard model object
type GradingStandard struct {
	// ContextID field: the id for the context either the Account or Course id
	ContextID int `json:"context_id"`
	// ContextType field: the context this standard is associated with, either 'Account' or 'Course'
	ContextType string `json:"context_type"`
	// GradingScheme field: A list of GradingSchemeEntry that make up the Grading Standard as an array of values with
	// the scheme name and value
	GradingScheme []GradingSchemeEntry `json:"grading_scheme"`
	// ID field: the id of the grading standard
	ID int `json:"id"`
	// Title field: the title of the grading standard
	Title string `json:"title"`
}

// GroupCategory model object
type GroupCategory struct {
	// AccountID field
	AccountID int `json:"account_id"`
	// AutoLeader field: Gives instructors the ability to automatically have group leaders assigned.  Values include
	// 'random', 'first', and null; 'random' picks a student from the group at random as the leader, 'first' sets the
	// first student to be assigned to the group as the leader
	AutoLeader *GroupCategoryAutoLeader `json:"auto_leader"`
	// ContextType field: The course or account that the category group belongs to. The pattern here is that whatever
	// the context_type is, there will be an _id field named after that type. So if instead context_type was 'Course',
	// the course_id field would be replaced by an course_id field.
	ContextType string `json:"context_type"`
	// GroupLimit field: If self-signup is enabled, group_limit can be set to cap the number of users in each group. If
	// null, there is no limit.
	GroupLimit int `json:"group_limit"`
	// ID field: The ID of the group category.
	ID int `json:"id"`
	// Name field: The display name of the group category.
	Name string `json:"name"`
	// Progress field: If the group category has not yet finished a randomly student assignment request, a progress
	// object will be attached, which will contain information related to the progress of the assignment request. Refer
	// to the Progress API for more information
	Progress *Progress `json:"progress"`
	// Role field: Certain types of group categories have special role designations. Currently, these include:
	// 'communities', 'student_organized', and 'imported'. Regular course/account group categories have a role of null.
	Role string `json:"role"`
	// SelfSignup field: If the group category allows users to join a group themselves, thought they may only be a
	// member of one group per group category at a time. Values include 'restricted', 'enabled', and null 'enabled'
	// allows students to assign themselves to a group 'restricted' restricts them to only joining a group in their
	// section null disallows students from joining groups
	SelfSignup *GroupCategorySelfSignup `json:"self_signup"`
	// SisGroupCategoryID field: The SIS identifier for the group category. This field is only included if the user has
	// permission to manage or view SIS information.
	SisGroupCategoryID string `json:"sis_group_category_id"`
	// SisImportID field: The unique identifier for the SIS import. This field is only included if the user has
	// permission to manage SIS information.
	SisImportID int `json:"sis_import_id"`
}

// GroupMembership model object
type GroupMembership struct {
	// GroupID field: The id of the group object to which the membership belongs
	GroupID int `json:"group_id"`
	// ID field: The id of the membership object
	ID int `json:"id"`
	// JustCreated field: optional: whether or not the record was just created on a create call (POST), i.e. was the
	// user just added to the group, or was the user already a member
	JustCreated bool `json:"just_created"`
	// Moderator field: Whether or not the user is a moderator of the group (the must also be an active member of the
	// group to moderate)
	Moderator bool `json:"moderator"`
	// SisImportID field: The id of the SIS import if created through SIS. Only included if the user has permission to
	// manage SIS information.
	SisImportID int `json:"sis_import_id"`
	// UserID field: The id of the user object to which the membership belongs
	UserID int `json:"user_id"`
	// WorkflowState field: The current state of the membership. Current possible values are 'accepted', 'invited', and
	// 'requested'
	WorkflowState *GroupMembershipWorkflowState `json:"workflow_state"`
}

// Group model object
type Group struct {
	// AvatarURL field: The url of the group's avatar
	AvatarURL string `json:"avatar_url"`
	// ContextType field: The course or account that the group belongs to. The pattern here is that whatever the
	// context_type is, there will be an _id field named after that type. So if instead context_type was 'account', the
	// course_id field would be replaced by an account_id field.
	ContextType string `json:"context_type"`
	// CourseID field
	CourseID int `json:"course_id"`
	// Description field: A description of the group. This is plain text.
	Description string `json:"description"`
	// FollowedByUser field: Whether or not the current user is following this group.
	FollowedByUser bool `json:"followed_by_user"`
	// GroupCategoryID field: The ID of the group's category.
	GroupCategoryID int `json:"group_category_id"`
	// ID field: The ID of the group.
	ID int `json:"id"`
	// IsPublic field: Whether or not the group is public.  Currently only community groups can be made public.  Also,
	// once a group has been set to public, it cannot be changed back to private.
	IsPublic bool `json:"is_public"`
	// JoinLevel field: How people are allowed to join the group.  For all groups except for community groups, the user
	// must share the group's parent course or account.  For student organized or community groups, where a user can be
	// a member of as many or few as they want, the applicable levels are 'parent_context_auto_join',
	// 'parent_context_request', and 'invitation_only'.  For class groups, where students are divided up and should only
	// be part of one group of the category, this value will always be 'invitation_only', and is not relevant. * If
	// 'parent_context_auto_join', anyone can join and will be automatically accepted. * If 'parent_context_request',
	// anyone  can request to join, which must be approved by a group moderator. * If 'invitation_only', only those how
	// have received an invitation my join the group, by accepting that invitation.
	JoinLevel *GroupJoinLevel `json:"join_level"`
	// MembersCount field: The number of members currently in the group
	MembersCount int `json:"members_count"`
	// Name field: The display name of the group.
	Name string `json:"name"`
	// Permissions field: optional: the permissions the user has for the group. returned only for a single group and
	// include[]=permissions
	Permissions map[string]bool `json:"permissions"`
	// Role field: Certain types of groups have special role designations. Currently, these include: 'communities',
	// 'student_organized', and 'imported'. Regular course/account groups have a role of null.
	Role *GroupRole `json:"role"`
	// SisGroupID field: The SIS ID of the group. Only included if the user has permission to view SIS information.
	SisGroupID string `json:"sis_group_id"`
	// SisImportID field: The id of the SIS import if created through SIS. Only included if the user has permission to
	// manage SIS information.
	SisImportID int `json:"sis_import_id"`
	// StorageQuotaMb field: the storage quota for the group, in megabytes
	StorageQuotaMb int `json:"storage_quota_mb"`
}

// JWT model object
type JWT struct {
	// Token field: The signed, encrypted, base64 encoded JWT
	Token string `json:"token"`
}

// LatePolicy model object
type LatePolicy struct {
	// CourseID field: the unique identifier for the course
	CourseID int `json:"course_id"`
	// CreatedAt field: the time at which this late policy was originally created
	CreatedAt time.Time `json:"created_at"`
	// ID field: the unique identifier for the late policy
	ID int `json:"id"`
	// LateSubmissionDeduction field: amount of percentage points to deduct per late_submission_interval
	LateSubmissionDeduction float64 `json:"late_submission_deduction"`
	// LateSubmissionDeductionEnabled field: whether to enable late submission deductions
	LateSubmissionDeductionEnabled bool `json:"late_submission_deduction_enabled"`
	// LateSubmissionInterval field: time interval for late submission deduction
	LateSubmissionInterval string `json:"late_submission_interval"`
	// LateSubmissionMinimumPercent field: the minimum score a submission can receive in percentage points
	LateSubmissionMinimumPercent float64 `json:"late_submission_minimum_percent"`
	// LateSubmissionMinimumPercentEnabled field: whether to enable late submission minimum percent
	LateSubmissionMinimumPercentEnabled bool `json:"late_submission_minimum_percent_enabled"`
	// MissingSubmissionDeduction field: amount of percentage points to deduct
	MissingSubmissionDeduction float64 `json:"missing_submission_deduction"`
	// MissingSubmissionDeductionEnabled field: whether to enable missing submission deductions
	MissingSubmissionDeductionEnabled bool `json:"missing_submission_deduction_enabled"`
	// UpdatedAt field: the time at which this late policy was last modified in any way
	UpdatedAt time.Time `json:"updated_at"`
}

// MigrationIssue model object
type MigrationIssue struct {
	// ContentMigrationURL field: API url to the content migration
	ContentMigrationURL string `json:"content_migration_url"`
	// CreatedAt field: timestamp
	CreatedAt time.Time `json:"created_at"`
	// Description field: Description of the issue for the end-user
	Description string `json:"description"`
	// ErrorMessage field: Site administrator error message (If the requesting user has permissions)
	ErrorMessage string `json:"error_message"`
	// ErrorReportHTMLURL field: Link to a Canvas error report if present (If the requesting user has permissions)
	ErrorReportHTMLURL string `json:"error_report_html_url"`
	// FixIssueHTMLURL field: HTML Url to the Canvas page to investigate the issue
	FixIssueHTMLURL string `json:"fix_issue_html_url"`
	// ID field: the unique identifier for the issue
	ID int `json:"id"`
	// IssueType field: Severity of the issue: todo, warning, error
	IssueType *MigrationIssueIssueType `json:"issue_type"`
	// UpdatedAt field: timestamp
	UpdatedAt time.Time `json:"updated_at"`
	// WorkflowState field: Current state of the issue: active, resolved
	WorkflowState *MigrationIssueWorkflowState `json:"workflow_state"`
}

// NotificationPreference model object
type NotificationPreference struct {
	// Category field: The category of that notification
	Category string `json:"category"`
	// Frequency field: How often to send notifications to this communication channel for the given notification.
	// Possible values are 'immediately', 'daily', 'weekly', and 'never'
	Frequency *NotificationPreferenceFrequency `json:"frequency"`
	// Href field
	Href string `json:"href"`
	// Notification field: The notification this preference belongs to
	Notification string `json:"notification"`
}

// OutcomeGroup model object
type OutcomeGroup struct {
	// CanEdit field: whether the current user can update the outcome group
	CanEdit bool `json:"can_edit"`
	// ContextID field: the context owning the outcome group. may be null for global outcome groups. omitted in the
	// abbreviated form.
	ContextID int `json:"context_id"`
	// ContextType field
	ContextType string `json:"context_type"`
	// Description field: description of the outcome group. omitted in the abbreviated form.
	Description string `json:"description"`
	// ID field: the ID of the outcome group
	ID int `json:"id"`
	// ImportURL field: the URL for importing another group into this outcome group. should be treated as opaque.
	// omitted in the abbreviated form.
	ImportURL string `json:"import_url"`
	// OutcomesURL field: the URL for listing/creating outcome links under the outcome group. should be treated as
	// opaque
	OutcomesURL string `json:"outcomes_url"`
	// ParentOutcomeGroup field: an abbreviated OutcomeGroup object representing the parent group of this outcome group,
	// if any. omitted in the abbreviated form.
	ParentOutcomeGroup *OutcomeGroup `json:"parent_outcome_group"`
	// SubgroupsURL field: the URL for listing/creating subgroups under the outcome group. should be treated as opaque
	SubgroupsURL string `json:"subgroups_url"`
	// Title field: title of the outcome group
	Title string `json:"title"`
	// URL field: the URL for fetching/updating the outcome group. should be treated as opaque
	URL string `json:"url"`
	// VendorGUID field: A custom GUID for the learning standard.
	VendorGUID string `json:"vendor_guid"`
}

// OutcomeLink model object
type OutcomeLink struct {
	// Assessed field: whether this outcome has been used to assess a student in the context of this outcome link.  In
	// other words, this will be set to true if the context is a course, and a student has been assessed with this
	// outcome in that course.
	Assessed bool `json:"assessed"`
	// CanUnlink field: whether this outcome link is manageable and is not the last link to an aligned outcome
	CanUnlink bool `json:"can_unlink"`
	// ContextID field: the context owning the outcome link. will match the context owning the outcome group containing
	// the outcome link; included for convenience. may be null for links in global outcome groups.
	ContextID int `json:"context_id"`
	// ContextType field
	ContextType string `json:"context_type"`
	// Outcome field: an abbreviated Outcome object representing the outcome linked into the containing outcome group.
	Outcome *Outcome `json:"outcome"`
	// OutcomeGroup field: an abbreviated OutcomeGroup object representing the group containing the outcome link.
	OutcomeGroup *OutcomeGroup `json:"outcome_group"`
	// URL field: the URL for fetching/updating the outcome link. should be treated as opaque
	URL string `json:"url"`
}

// OutcomeImportData model object
type OutcomeImportData struct {
	// ImportType field: The type of outcome import
	ImportType string `json:"import_type"`
}

// OutcomeImport model object
type OutcomeImport struct {
	// CreatedAt field: The date the outcome import was created.
	CreatedAt time.Time `json:"created_at"`
	// Data field: See the OutcomeImportData specification above.
	Data *OutcomeImportData `json:"data"`
	// EndedAt field: The date the outcome import finished. Returns null if not finished.
	EndedAt time.Time `json:"ended_at"`
	// ID field: The unique identifier for the outcome import.
	ID int `json:"id"`
	// ProcessingErrors field: An array of row number / error message pairs. Returns the first 25 errors.
	ProcessingErrors [][]map[interface{}]interface{} `json:"processing_errors"`
	// Progress field: The progress of the outcome import.
	Progress string `json:"progress"`
	// UpdatedAt field: The date the outcome import was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// User field: The user that initiated the outcome_import. See the Users API for details.
	User *User `json:"user"`
	// WorkflowState field: The current state of the outcome import.
	//  'created': The outcome import has been created.
	//  'importing': The outcome import is currently processing.
	//  'succeeded': The outcome import has completed successfully.
	//  'failed': The outcome import failed.
	WorkflowState *OutcomeImportWorkflowState `json:"workflow_state"`
}

// ProficiencyRating model object
type ProficiencyRating struct {
	// Color field: The hex color code of the rating
	Color string `json:"color"`
	// Description field: The description of the rating
	Description string `json:"description"`
	// Mastery field: Indicates the rating where mastery is first achieved
	Mastery bool `json:"mastery"`
	// Points field: A non-negative number of points for the rating
	Points float64 `json:"points"`
}

// Proficiency model object
type Proficiency struct {
	// Ratings field: An array of proficiency ratings. See the ProficiencyRating specification above.
	Ratings []interface{} `json:"ratings"`
}

// OutcomeResult model object: A student's result for an outcome
type OutcomeResult struct {
	// ID field: A unique identifier for this result
	ID int `json:"id"`
	// Links field: Unique identifiers of objects associated with this result
	Links map[interface{}]interface{} `json:"links"`
	// Percent field: score's percent of maximum points possible for outcome, scaled to reflect any custom mastery
	// levels that differ from the learning outcome
	Percent float64 `json:"percent"`
	// Score field: The student's score
	Score int `json:"score"`
	// SubmittedOrAssessedAt field: The datetime the resulting OutcomeResult was submitted at, or absent that, when it
	// was assessed.
	SubmittedOrAssessedAt time.Time `json:"submitted_or_assessed_at"`
}

// OutcomeRollupScoreLinks model object
type OutcomeRollupScoreLinks struct {
	// Outcome field: The id of the related outcome
	Outcome int `json:"outcome"`
}

// OutcomeRollupScore model object
type OutcomeRollupScore struct {
	// Count field: The number of alignment scores included in this rollup.
	Count int `json:"count"`
	// Links field
	Links *OutcomeRollupScoreLinks `json:"links"`
	// Score field: The rollup score for the outcome, based on the student alignment scores related to the outcome. This
	// could be null if the student has no related scores.
	Score int `json:"score"`
}

// OutcomeRollupLinks model object
type OutcomeRollupLinks struct {
	// Course field: If an aggregate result was requested, the course field will be present. Otherwise, the user and
	// section field will be present (Optional) The id of the course that this rollup applies to
	Course int `json:"course"`
	// Section field: (Optional) The id of the section the user is in
	Section int `json:"section"`
	// User field: (Optional) The id of the user that this rollup applies to
	User int `json:"user"`
}

// OutcomeRollup model object
type OutcomeRollup struct {
	// Links field
	Links *OutcomeRollupLinks `json:"links"`
	// Name field: The name of the resource for this rollup. For example, the user name.
	Name string `json:"name"`
	// Scores field: an array of OutcomeRollupScore objects
	Scores *OutcomeRollupScore `json:"scores"`
}

// OutcomeAlignment model object: An asset aligned with this outcome
type OutcomeAlignment struct {
	// HTMLURL field: (Optional) A URL for details about this alignment
	HTMLURL string `json:"html_url"`
	// ID field: A unique identifier for this alignment
	ID string `json:"id"`
	// Name field: The name of this alignment
	Name string `json:"name"`
	// AssessmentID field: the id of the aligned live assessment (null for assignments).
	AssessmentID int `json:"assessment_id"`
	// AssignmentID field: the id of the aligned assignment (null for live assessments).
	AssignmentID int `json:"assignment_id"`
	// SubmissionTypes field: a string representing the different submission types of an aligned assignment.
	SubmissionTypes string `json:"submission_types"`
	// Title field: the title of the aligned assignment.
	Title string `json:"title"`
	// URL field: the URL for the aligned assignment.
	URL string `json:"url"`
}

// OutcomePath model object: The full path to an outcome
type OutcomePath struct {
	// ID field: A unique identifier for this outcome
	ID int `json:"id"`
	// Parts field: an array of OutcomePathPart objects
	Parts *OutcomePathPart `json:"parts"`
}

// OutcomePathPart model object: An outcome or outcome group
type OutcomePathPart struct {
	// Name field: The title of the outcome or outcome group
	Name string `json:"name"`
}

// Outcome model object
type Outcome struct {
	// Assessed field: whether this outcome has been used to assess a student
	Assessed bool `json:"assessed"`
	// CalculationInt field: this defines the variable value used by the calculation_method. included only if
	// calculation_method uses it
	CalculationInt int `json:"calculation_int"`
	// CalculationMethod field: the method used to calculate a students score
	CalculationMethod *OutcomeCalculationMethod `json:"calculation_method"`
	// CanEdit field: whether the current user can update the outcome
	CanEdit bool `json:"can_edit"`
	// CanUnlink field: whether the outcome can be unlinked
	CanUnlink bool `json:"can_unlink"`
	// ContextID field: the context owning the outcome. may be null for global outcomes
	ContextID int `json:"context_id"`
	// ContextType field
	ContextType string `json:"context_type"`
	// Description field: description of the outcome. omitted in the abbreviated form.
	Description string `json:"description"`
	// DisplayName field: Optional friendly name for reporting
	DisplayName string `json:"display_name"`
	// HasUpdateableRubrics field: whether updates to this outcome will propagate to unassessed rubrics that have
	// imported it
	HasUpdateableRubrics bool `json:"has_updateable_rubrics"`
	// ID field: the ID of the outcome
	ID int `json:"id"`
	// MasteryPoints field: points necessary to demonstrate mastery outcomes. included only if the outcome embeds a
	// rubric criterion. omitted in the abbreviated form.
	MasteryPoints int `json:"mastery_points"`
	// PointsPossible field: maximum points possible. included only if the outcome embeds a rubric criterion. omitted in
	// the abbreviated form.
	PointsPossible int `json:"points_possible"`
	// Ratings field: possible ratings for this outcome. included only if the outcome embeds a rubric criterion. omitted
	// in the abbreviated form.
	Ratings []RubricRating `json:"ratings"`
	// Title field: title of the outcome
	Title string `json:"title"`
	// URL field: the URL for fetching/updating the outcome. should be treated as opaque
	URL string `json:"url"`
	// VendorGUID field: A custom GUID for the learning standard.
	VendorGUID string `json:"vendor_guid"`
}

// PageView model object: The record of a user page view access in Canvas
type PageView struct {
	// Action field: The rails action that handled the request
	Action string `json:"action"`
	// AppName field: If the request is from an API request, the app that generated the access token
	AppName string `json:"app_name"`
	// AssetType field: The type of asset in the context for the request, if any
	AssetType string `json:"asset_type"`
	// ContextType field: The type of context for the request
	ContextType string `json:"context_type"`
	// Contributed field: This field is deprecated, and will always be false
	Contributed bool `json:"contributed"`
	// Controller field: The rails controller that handled the request
	Controller string `json:"controller"`
	// CreatedAt field: When the request was made
	CreatedAt time.Time `json:"created_at"`
	// HTTPMethod field: The HTTP method such as GET or POST
	HTTPMethod string `json:"http_method"`
	// ID field: A UUID representing the page view.  This is also the unique request id
	ID string `json:"id"`
	// InteractionSeconds field: An approximation of how long the user spent on the page, in seconds
	InteractionSeconds float64 `json:"interaction_seconds"`
	// Links field: The page view links to define the relationships
	Links *PageViewLinks `json:"links"`
	// Participated field: True if the request counted as participating, such as submitting homework
	Participated bool `json:"participated"`
	// RemoteIP field: The origin IP address of the request
	RemoteIP string `json:"remote_ip"`
	// RenderTime field: How long the response took to render, in seconds
	RenderTime float64 `json:"render_time"`
	// URL field: The URL requested
	URL string `json:"url"`
	// UserAgent field: The user-agent of the browser or program that made the request
	UserAgent string `json:"user_agent"`
	// UserRequest field: A flag indicating whether the request was user-initiated, or automatic (such as an AJAX call)
	UserRequest bool `json:"user_request"`
}

// PageViewLinks model object: The links of a page view access in Canvas
type PageViewLinks struct {
	// Account field: The ID of the account context for this page view
	Account int `json:"account"`
	// Asset field: The ID of the asset for the request, if any
	Asset int `json:"asset"`
	// Context field: The ID of the context for the request (course id if context_type is Course, etc)
	Context int `json:"context"`
	// RealUser field: The ID of the actual user who made this request, if the request was made by a user who was
	// masquerading
	RealUser int `json:"real_user"`
	// User field: The ID of the user for this page view
	User int `json:"user"`
}

// PeerReview model object
type PeerReview struct {
	// Assessor field: The User object for the assessor if the user include parameter is provided (see user API)
	// (optional)
	Assessor string `json:"assessor"`
	// AssessorID field: The assessors user id
	AssessorID int `json:"assessor_id"`
	// AssetID field: The id for the asset associated with this Peer Review
	AssetID int `json:"asset_id"`
	// AssetType field: The type of the asset
	AssetType string `json:"asset_type"`
	// ID field: The id of the Peer Review
	ID int `json:"id"`
	// SubmissionComments field: The submission comments associated with this Peer Review if the submission_comment
	// include parameter is provided (see submissions API) (optional)
	SubmissionComments string `json:"submission_comments"`
	// User field: the User object for the owner of the asset if the user include parameter is provided (see user API)
	// (optional)
	User string `json:"user"`
	// UserID field: The user id for the owner of the asset
	UserID int `json:"user_id"`
	// WorkflowState field: The state of the Peer Review, either 'assigned' or 'completed'
	WorkflowState string `json:"workflow_state"`
}

// PlannerNote model object: A planner note
type PlannerNote struct {
	// CourseID field: The course that the note is in relation too, if applicable
	CourseID int `json:"course_id"`
	// Description field: The description of the planner note
	Description string `json:"description"`
	// ID field: The ID of the planner note
	ID int `json:"id"`
	// LinkedObjectHTMLURL field: the Canvas web URL of the linked learning object
	LinkedObjectHTMLURL string `json:"linked_object_html_url"`
	// LinkedObjectID field: the id of the linked learning object
	LinkedObjectID int `json:"linked_object_id"`
	// LinkedObjectType field: the type of the linked learning object
	LinkedObjectType string `json:"linked_object_type"`
	// LinkedObjectURL field: the API URL of the linked learning object
	LinkedObjectURL string `json:"linked_object_url"`
	// Title field: The title for a planner note
	Title string `json:"title"`
	// TodoDate field: The datetime of when the planner note should show up on their planner
	TodoDate time.Time `json:"todo_date"`
	// UserID field: The id of the associated user creating the planner note
	UserID int `json:"user_id"`
	// WorkflowState field: The current published state of the planner note
	WorkflowState string `json:"workflow_state"`
}

// PlannerOverride model object: User-controlled setting for whether an item should be displayed on the planner or not
type PlannerOverride struct {
	// AssignmentID field: The id of the plannable's associated assignment, if it has one
	AssignmentID int `json:"assignment_id"`
	// CreatedAt field: The datetime of when the planner override was created
	CreatedAt time.Time `json:"created_at"`
	// DeletedAt field: The datetime of when the planner override was deleted, if applicable
	DeletedAt time.Time `json:"deleted_at"`
	// Dismissed field: Controls whether or not the associated plannable item shows up in the opportunities list
	Dismissed bool `json:"dismissed"`
	// ID field: The ID of the planner override
	ID int `json:"id"`
	// MarkedComplete field: Controls whether or not the associated plannable item is marked complete on the planner
	MarkedComplete bool `json:"marked_complete"`
	// PlannableID field: The id of the associated object for the planner override
	PlannableID int `json:"plannable_id"`
	// PlannableType field: The type of the associated object for the planner override
	PlannableType string `json:"plannable_type"`
	// UpdatedAt field: The datetime of when the planner override was updated
	UpdatedAt time.Time `json:"updated_at"`
	// UserID field: The id of the associated user for the planner override
	UserID int `json:"user_id"`
	// WorkflowState field: The current published state of the item, synced with the associated object
	WorkflowState string `json:"workflow_state"`
}

// Profile model object: Profile details for a Canvas user.
type Profile struct {
	// AvatarURL field: The avatar_url can change over time, so we recommend not caching it for more than a few hours
	AvatarURL string `json:"avatar_url"`
	// Bio field
	Bio string `json:"bio"`
	// Calendar field
	Calendar *CalendarLink `json:"calendar"`
	// ID field: The ID of the user.
	ID int `json:"id"`
	// Locale field: The users locale.
	Locale string `json:"locale"`
	// LoginID field: sample_user@example.com
	LoginID string `json:"login_id"`
	// LtiUserID field
	LtiUserID string `json:"lti_user_id"`
	// Name field: Sample User
	Name string `json:"name"`
	// PrimaryEmail field: sample_user@example.com
	PrimaryEmail string `json:"primary_email"`
	// ShortName field: Sample User
	ShortName string `json:"short_name"`
	// SisUserID field: sis1
	SisUserID string `json:"sis_user_id"`
	// SortableName field: user, sample
	SortableName string `json:"sortable_name"`
	// TimeZone field: Optional: This field is only returned in certain API calls, and will return the IANA time zone
	// name of the user's preferred timezone.
	TimeZone string `json:"time_zone"`
	// Title field
	Title string `json:"title"`
}

// Avatar model object: Possible avatar for a user.
type Avatar struct {
	// ContentType field: ['attachment' type only] the content-type of the attachment.
	ContentType string `json:"content-type"`
	// DisplayName field: A textual description of the avatar record.
	DisplayName string `json:"display_name"`
	// Filename field: ['attachment' type only] the filename of the attachment
	Filename string `json:"filename"`
	// ID field: ['attachment' type only] the internal id of the attachment
	ID int `json:"id"`
	// Size field: ['attachment' type only] the size of the attachment
	Size int `json:"size"`
	// Token field: A unique representation of the avatar record which can be used to set the avatar with the user
	// update endpoint. Note: this is an internal representation and is subject to change without notice. It should be
	// consumed with this api endpoint and used in the user update endpoint, and should not be constructed by the
	// client.
	Token string `json:"token"`
	// Type field: ['gravatar'|'attachment'|'no_pic'] The type of avatar record, for categorization purposes.
	Type string `json:"type"`
	// URL field: The url of the avatar
	URL string `json:"url"`
}

// Progress model object
type Progress struct {
	// Completion field: percent completed
	Completion int `json:"completion"`
	// ContextID field: the context owning the job.
	ContextID int `json:"context_id"`
	// ContextType field
	ContextType string `json:"context_type"`
	// CreatedAt field: the time the job was created
	CreatedAt time.Time `json:"created_at"`
	// ID field: the ID of the Progress object
	ID int `json:"id"`
	// Message field: optional details about the job
	Message string `json:"message"`
	// Results field: optional results of the job. omitted when job is still pending
	Results map[interface{}]interface{} `json:"results"`
	// Tag field: the type of operation
	Tag string `json:"tag"`
	// UpdatedAt field: the time the job was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// URL field: url where a progress update can be retrieved
	URL string `json:"url"`
	// UserID field: the id of the user who started the job
	UserID int `json:"user_id"`
	// WorkflowState field: the state of the job one of 'queued', 'running', 'completed', 'failed'
	WorkflowState *ProgressWorkflowState `json:"workflow_state"`
}

// ProvisionalGrade model object
type ProvisionalGrade struct {
	// Final field: Whether this is the 'final' provisional grade created by the moderator
	Final bool `json:"final"`
	// Grade field: The grade
	Grade string `json:"grade"`
	// GradeMatchesCurrentSubmission field: Whether the grade was applied to the most current submission (false if the
	// student resubmitted after grading)
	GradeMatchesCurrentSubmission bool `json:"grade_matches_current_submission"`
	// GradedAt field: When the grade was given
	GradedAt time.Time `json:"graded_at"`
	// ProvisionalGradeID field: The identifier for the provisional grade
	ProvisionalGradeID int `json:"provisional_grade_id"`
	// Score field: The numeric score
	Score int `json:"score"`
	// SpeedgraderURL field: A link to view this provisional grade in SpeedGrader™
	SpeedgraderURL string `json:"speedgrader_url"`
}

// RolePermissions model object
type RolePermissions struct {
	// AppliesToDescendants field: Whether the permission cascades down to sub accounts of the account this role is in.
	// Only present if enabled is true
	AppliesToDescendants bool `json:"applies_to_descendants"`
	// AppliesToSelf field: Whether the permission applies to the account this role is in. Only present if enabled is
	// true
	AppliesToSelf bool `json:"applies_to_self"`
	// Enabled field: Whether the role has the permission
	Enabled bool `json:"enabled"`
	// Explicit field: Whether the value of enabled is specified explicitly by this role, or inherited from an upstream
	// role.
	Explicit bool `json:"explicit"`
	// Locked field: Whether the permission is locked by this role
	Locked bool `json:"locked"`
	// PriorDefault field: The value that would have been inherited from upstream if the role had not explicitly set a
	// value. Only present if explicit is true.
	PriorDefault bool `json:"prior_default"`
	// Readonly field: Whether the permission can be modified in this role (i.e. whether the permission is locked by an
	// upstream role).
	Readonly bool `json:"readonly"`
}

// Role model object
type Role struct {
	// Account field: JSON representation of the account the role is in.
	Account *Account `json:"account"`
	// BaseRoleType field: The role type that is being used as a base for this role. For account-level roles, this is
	// 'AccountMembership'. For course-level roles, it is an enrollment type.
	BaseRoleType string `json:"base_role_type"`
	// Label field: The label of the role.
	Label string `json:"label"`
	// Permissions field: A dictionary of permissions keyed by name (see permissions input parameter in the 'Create a
	// role' API).
	Permissions map[string]RolePermissions `json:"permissions"`
	// Role field: The label of the role. (Deprecated alias for 'label')
	Role string `json:"role"`
	// WorkflowState field: The state of the role: 'active', 'inactive', or 'built_in'
	WorkflowState string `json:"workflow_state"`
}

// Rubric model object
type Rubric struct {
	// Assessments field: If an assessment type is included in the 'include' parameter, includes an array of rubric
	// assessment objects for a given rubric, based on the assessment type requested. If the user does not request an
	// assessment type this key will be absent.
	Assessments []RubricAssessment `json:"assessments"`
	// Associations field: If an association type is included in the 'include' parameter, includes an array of rubric
	// association objects for a given rubric, based on the association type requested. If the user does not request an
	// association type this key will be absent.
	Associations []RubricAssociation `json:"associations"`
	// ContextID field: the context owning the rubric
	ContextID int `json:"context_id"`
	// ContextType field
	ContextType string `json:"context_type"`
	// Data field: An array with all of this Rubric's grading Criteria
	Data []RubricCriterion `json:"data"`
	// FreeFormCriterionComments field: whether or not free-form comments are used
	FreeFormCriterionComments bool `json:"free_form_criterion_comments"`
	// HideScoreTotal field
	HideScoreTotal bool `json:"hide_score_total"`
	// ID field: the ID of the rubric
	ID int `json:"id"`
	// PointsPossible field
	PointsPossible int `json:"points_possible"`
	// ReadOnly field
	ReadOnly bool `json:"read_only"`
	// Reusable field
	Reusable bool `json:"reusable"`
	// Title field: title of the rubric
	Title string `json:"title"`
}

// RubricCriterion model object
type RubricCriterion struct {
	// CriterionUseRange field
	CriterionUseRange bool `json:"criterion_use_range"`
	// Description field
	Description string `json:"description"`
	// ID field: the ID of the criterion
	ID string `json:"id"`
	// LongDescription field
	LongDescription string `json:"long_description"`
	// Points field
	Points int `json:"points"`
	// Ratings field: the possible ratings for this Criterion
	Ratings []RubricRating `json:"ratings"`
}

// RubricAssessment model object
type RubricAssessment struct {
	// ArtifactAttempt field: the current number of attempts made on the object of the assessment
	ArtifactAttempt int `json:"artifact_attempt"`
	// ArtifactID field: the id of the object of the assessment
	ArtifactID int `json:"artifact_id"`
	// ArtifactType field: the object of the assessment
	ArtifactType string `json:"artifact_type"`
	// AssessmentType field: the type of assessment. values will be either 'grading', 'peer_review', or
	// 'provisional_grade'
	AssessmentType string `json:"assessment_type"`
	// AssessorID field: user id of the person who made the assessment
	AssessorID int `json:"assessor_id"`
	// Comments field: (Optional) If 'comments_only' is included in the 'style' parameter, returned assessments will
	// include only the comments portion of their data hash. If the user does not request a style, this key will be
	// absent.
	Comments []string `json:"comments"`
	// Data field: (Optional) If 'full' is included in the 'style' parameter, returned assessments will have their full
	// details contained in their data hash. If the user does not request a style, this key will be absent.
	Data []map[interface{}]interface{} `json:"data"`
	// ID field: the ID of the rubric
	ID int `json:"id"`
	// RubricAssociationID field
	RubricAssociationID int `json:"rubric_association_id"`
	// RubricID field: the rubric the assessment belongs to
	RubricID int `json:"rubric_id"`
	// Score field
	Score int `json:"score"`
}

// RubricAssociation model object
type RubricAssociation struct {
	// AssociationID field: the ID of the object this association links to
	AssociationID int `json:"association_id"`
	// AssociationType field: the type of object this association links to
	AssociationType string `json:"association_type"`
	// HideOutcomeResults field
	HideOutcomeResults bool `json:"hide_outcome_results"`
	// HidePoints field
	HidePoints bool `json:"hide_points"`
	// HideScoreTotal field: Whether or not the score total is displayed within the rubric. This option is only
	// available if the rubric is not used for grading.
	HideScoreTotal bool `json:"hide_score_total"`
	// ID field: the ID of the association
	ID int `json:"id"`
	// Purpose field: Whether or not the association is for grading (and thus linked to an assignment) or if it's to
	// indicate the rubric should appear in its context. Values will be grading or bookmark.
	Purpose string `json:"purpose"`
	// RubricID field: the ID of the rubric
	RubricID int `json:"rubric_id"`
	// SummaryData field
	SummaryData string `json:"summary_data"`
	// UseForGrading field: Whether or not the associated rubric is used for grade calculation
	UseForGrading bool `json:"use_for_grading"`
}

// Scope model object
type Scope struct {
	// Action field: The controller action the scope is associated to
	Action string `json:"action"`
	// Controller field: The controller the scope is associated to
	Controller string `json:"controller"`
	// Resource field: The resource the scope is associated with
	Resource string `json:"resource"`
	// ResourceName field: The localized resource name
	ResourceName string `json:"resource_name"`
	// Scope field: The identifier for the scope
	Scope string `json:"scope"`
	// Verb field: The HTTP verb for the scope
	Verb string `json:"verb"`
}

// Section model object
type Section struct {
	// CourseID field: The unique Canvas identifier for the course in which the section belongs
	CourseID int `json:"course_id"`
	// EndAt field: the end date for the section, if applicable
	EndAt time.Time `json:"end_at"`
	// ID field: The unique identifier for the section.
	ID int `json:"id"`
	// IntegrationID field: Optional: The integration ID of the section. This field is only included if the user has
	// permission to view SIS information.
	IntegrationID string `json:"integration_id"`
	// Name field: The name of the section.
	Name string `json:"name"`
	// NonxlistCourseID field: The unique identifier of the original course of a cross-listed section
	NonxlistCourseID int `json:"nonxlist_course_id"`
	// RestrictEnrollmentsToSectionDates field: Restrict user enrollments to the start and end dates of the section
	RestrictEnrollmentsToSectionDates bool `json:"restrict_enrollments_to_section_dates"`
	// SisCourseID field: The unique SIS identifier for the course in which the section belongs. This field is only
	// included if the user has permission to view SIS information.
	SisCourseID string `json:"sis_course_id"`
	// SisImportID field: The unique identifier for the SIS import if created through SIS. This field is only included
	// if the user has permission to manage SIS information.
	SisImportID int `json:"sis_import_id"`
	// SisSectionID field: The sis id of the section. This field is only included if the user has permission to view SIS
	// information.
	SisSectionID string `json:"sis_section_id"`
	// StartAt field: the start date for the section, if applicable
	StartAt time.Time `json:"start_at"`
	// TotalStudents field: optional: the total number of active and invited students in the section
	TotalStudents int `json:"total_students"`
}

// SharedBrandConfig model object
type SharedBrandConfig struct {
	// AccountID field: The id of the account it should be shared within.
	AccountID string `json:"account_id"`
	// BrandConfigMd5 field: The md5 (since BrandConfigs are identified by MD5 and not numeric id) of the BrandConfig to
	// share.
	BrandConfigMd5 string `json:"brand_config_md5"`
	// CreatedAt field: When this was created
	CreatedAt time.Time `json:"created_at"`
	// ID field: The shared_brand_config identifier.
	ID int `json:"id"`
	// Name field: The name to share this theme as
	Name string `json:"name"`
	// UpdatedAt field: When this was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

// SisAssignment model object: Assignments that have post_to_sis enabled with other objects for convenience
type SisAssignment struct {
	// AssignmentGroup field: Includes attributes of a assignment_group for convenience. For more details see
	// Assignments API.
	AssignmentGroup []AssignmentGroupAttributes `json:"assignment_group"`
	// CourseID field: The unique identifier for the course.
	CourseID int `json:"course_id"`
	// CreatedAt field: The time at which this assignment was originally created
	CreatedAt time.Time `json:"created_at"`
	// DueAt field: the due date for the assignment. returns null if not present. NOTE: If this assignment has
	// assignment overrides, this field will be the due date as it applies to the user requesting information from the
	// API.
	DueAt time.Time `json:"due_at"`
	// ID field: The unique identifier for the assignment.
	ID int `json:"id"`
	// IncludeInFinalGrade field: If false, the assignment will be omitted from the student's final grade
	IncludeInFinalGrade bool `json:"include_in_final_grade"`
	// IntegrationData field: (optional, Third Party integration data for assignment)
	IntegrationData string `json:"integration_data"`
	// IntegrationID field: Third Party integration id for assignment
	IntegrationID string `json:"integration_id"`
	// LockAt field: (Optional) Time at which this was/will be locked.
	LockAt time.Time `json:"lock_at"`
	// Name field: the name of the assignment
	Name string `json:"name"`
	// PointsPossible field: The maximum points possible for the assignment
	PointsPossible int `json:"points_possible"`
	// Sections field: Includes attributes of a section for convenience. For more details see Sections API.
	Sections []SectionAttributes `json:"sections"`
	// SubmissionTypes field: the types of submissions allowed for this assignment list containing one or more of the
	// following: 'discussion_topic', 'online_quiz', 'on_paper', 'none', 'external_tool', 'online_text_entry',
	// 'online_url', 'online_upload' 'media_recording'
	SubmissionTypes *SisAssignmentSubmissionTypes `json:"submission_types"`
	// UnlockAt field: (Optional) Time at which this was/will be unlocked.
	UnlockAt time.Time `json:"unlock_at"`
	// UserOverrides field: Includes attributes of a user assignment overrides. For more details see Assignments API.
	UserOverrides []UserAssignmentOverrideAttributes `json:"user_overrides"`
}

// AssignmentGroupAttributes model object: Some of the attributes of an Assignment Group. See Assignments API for more
// details
type AssignmentGroupAttributes struct {
	// GroupWeight field: the weight of the Assignment Group
	GroupWeight int `json:"group_weight"`
	// ID field: the id of the Assignment Group
	ID int `json:"id"`
	// IntegrationData field: the integration data of the Assignment Group
	IntegrationData map[interface{}]interface{} `json:"integration_data"`
	// Name field: the name of the Assignment Group
	Name string `json:"name"`
	// SisSourceID field: the sis source id of the Assignment Group
	SisSourceID string `json:"sis_source_id"`
}

// SectionAttributes model object: Some of the attributes of a section. For more details see Sections API.
type SectionAttributes struct {
	// ID field: The unique identifier for the section.
	ID int `json:"id"`
	// IntegrationID field: Optional: The integration ID of the section.
	IntegrationID string `json:"integration_id"`
	// Name field: The name of the section.
	Name string `json:"name"`
	// OriginCourse field: The course to which the section belongs or the course from which the section was cross-listed
	OriginCourse *CourseAttributes `json:"origin_course"`
	// Override field: Optional: Attributes of the assignment override that apply to the section. See Assignment API for
	// more details
	Override *SectionAssignmentOverrideAttributes `json:"override"`
	// SisID field: The sis id of the section.
	SisID string `json:"sis_id"`
	// XlistCourse field: Optional: Attributes of the xlist course. Only present when the section has been cross-listed.
	// See Courses API for more details
	XlistCourse *CourseAttributes `json:"xlist_course"`
}

// CourseAttributes model object: Attributes of a course object.  See Courses API for more details
type CourseAttributes struct {
	// ID field: The unique Canvas identifier for the origin course
	ID int `json:"id"`
	// IntegrationID field: The integration ID of the origin_course.
	IntegrationID string `json:"integration_id"`
	// Name field: The name of the origin course.
	Name string `json:"name"`
	// SisID field: The sis id of the origin_course.
	SisID string `json:"sis_id"`
}

// SectionAssignmentOverrideAttributes model object: Attributes of an assignment override that apply to the section
// object.  See Assignments API for more details
type SectionAssignmentOverrideAttributes struct {
	// DueAt field: the due date for the assignment. returns null if not present. NOTE: If this assignment has
	// assignment overrides, this field will be the due date as it applies to the user requesting information from the
	// API.
	DueAt time.Time `json:"due_at"`
	// LockAt field: (Optional) Time at which this was/will be locked.
	LockAt time.Time `json:"lock_at"`
	// OverrideTitle field: The title for the assignment override
	OverrideTitle string `json:"override_title"`
	// UnlockAt field: (Optional) Time at which this was/will be unlocked.
	UnlockAt time.Time `json:"unlock_at"`
}

// UserAssignmentOverrideAttributes model object: Attributes of assignment overrides that apply to users.  See
// Assignments API for more details
type UserAssignmentOverrideAttributes struct {
	// DueAt field: The time at which this assignment is due
	DueAt time.Time `json:"due_at"`
	// ID field: The unique Canvas identifier for the assignment override
	ID int `json:"id"`
	// LockAt field: (Optional) Time at which this was/will be locked.
	LockAt time.Time `json:"lock_at"`
	// Students field: Includes attributes of a student for convenience. For more details see Users API.
	Students []StudentAttributes `json:"students"`
	// Title field: The title of the assignment override.
	Title string `json:"title"`
	// UnlockAt field: (Optional) Time at which this was/will be unlocked.
	UnlockAt time.Time `json:"unlock_at"`
}

// StudentAttributes model object: Attributes of student.  See Users API for more details
type StudentAttributes struct {
	// SisUserID field: The SIS ID associated with the user.  This field is only included if the user came from a SIS
	// import and has permissions to view SIS information.
	SisUserID string `json:"sis_user_id"`
	// UserID field: The unique Canvas identifier for the user
	UserID int `json:"user_id"`
}

// SisImportError model object
type SisImportError struct {
	// File field: The file where the error message occurred.
	File string `json:"file"`
	// Message field: The error message that from the record.
	Message string `json:"message"`
	// Row field: The line number where the error occurred. Some Importers do not yet support this. This is a 1 based
	// index starting with the header row.
	Row int `json:"row"`
	// RowInfo field: The contents of the line that had the error.
	RowInfo string `json:"row_info"`
	// SisImportID field: The unique identifier for the SIS import.
	SisImportID int `json:"sis_import_id"`
}

// SisImportData model object
type SisImportData struct {
	// Counts field: The number of rows processed for each type of import
	Counts *SisImportCounts `json:"counts"`
	// ImportType field: The type of SIS import
	ImportType string `json:"import_type"`
	// SuppliedBatches field: Which files were included in the SIS import
	SuppliedBatches []string `json:"supplied_batches"`
}

// SisImportStatistic model object
type SisImportStatistic struct {
	// Concluded field: This is the number of items that marked as completed. This only applies to courses and
	// enrollments.
	Concluded int `json:"concluded"`
	// Created field: This is the number of items that were created.
	Created int `json:"created"`
	// Deactivated field: This is the number of Enrollments that were marked as 'inactive'. This only applies to
	// enrollments.
	Deactivated int `json:"deactivated"`
	// Deleted field: This is the number of items that were deleted.
	Deleted int `json:"deleted"`
	// Restored field: This is the number of items that were set to an active state from a completed, inactive, or
	// deleted state.
	Restored int `json:"restored"`
}

// SisImportStatistics model object
type SisImportStatistics struct {
	// AbstractCourse field: This contains that statistics for abstract courses.
	AbstractCourse *SisImportStatistic `json:"AbstractCourse"`
	// Account field: This contains that statistics for accounts.
	Account *SisImportStatistic `json:"Account"`
	// AccountUser field: This contains that statistics for account users.
	AccountUser *SisImportStatistic `json:"AccountUser"`
	// CommunicationChannel field: This contains that statistics for communication channels. This is an indirect effect
	// from creating or deleting a user.
	CommunicationChannel *SisImportStatistic `json:"CommunicationChannel"`
	// Course field: This contains that statistics for courses.
	Course *SisImportStatistic `json:"Course"`
	// CourseSection field: This contains that statistics for course sections.
	CourseSection *SisImportStatistic `json:"CourseSection"`
	// Enrollment field: This contains that statistics for enrollments.
	Enrollment *SisImportStatistic `json:"Enrollment"`
	// EnrollmentTerm field: This contains that statistics for terms.
	EnrollmentTerm *SisImportStatistic `json:"EnrollmentTerm"`
	// Group field: This contains that statistics for groups.
	Group *SisImportStatistic `json:"Group"`
	// GroupCategory field: This contains that statistics for group categories.
	GroupCategory *SisImportStatistic `json:"GroupCategory"`
	// GroupMembership field: This contains that statistics for group memberships. This can be a direct impact from the
	// import or indirect from an enrollment being deleted.
	GroupMembership *SisImportStatistic `json:"GroupMembership"`
	// Pseudonym field: This contains that statistics for pseudonyms. Pseudonyms are logins for users, and are the
	// object that ties an enrollment to a user. This would be impacted from the user importer.
	Pseudonym *SisImportStatistic `json:"Pseudonym"`
	// UserObserver field: This contains that statistics for user observers.
	UserObserver *SisImportStatistic `json:"UserObserver"`
	// TotalStateChanges field: This is the total number of items that were changed in the sis import. There are a few
	// caveats that can cause this number to not add up to the individual counts. There are some state changes that
	// happen that have no impact to the object. An example would be changing a course from 'created' to 'claimed'. Both
	// of these would be considered an active course, but would increment this counter. In this example the course would
	// not increment the created or restored counters for course statistic.
	TotalStateChanges int `json:"total_state_changes"`
}

// SisImportCounts model object
type SisImportCounts struct {
	// AbstractCourses field
	AbstractCourses int `json:"abstract_courses"`
	// Accounts field
	Accounts int `json:"accounts"`
	// BatchCoursesDeleted field: the number of courses that were removed because they were not included in the batch
	// for batch_mode imports. Only included if courses were deleted
	BatchCoursesDeleted int `json:"batch_courses_deleted"`
	// BatchEnrollmentsDeleted field: the number of enrollments that were removed because they were not included in the
	// batch for batch_mode imports. Only included if enrollments were deleted
	BatchEnrollmentsDeleted int `json:"batch_enrollments_deleted"`
	// BatchSectionsDeleted field: the number of sections that were removed because they were not included in the batch
	// for batch_mode imports. Only included if sections were deleted
	BatchSectionsDeleted int `json:"batch_sections_deleted"`
	// Courses field
	Courses int `json:"courses"`
	// Enrollments field
	Enrollments int `json:"enrollments"`
	// ErrorCount field
	ErrorCount int `json:"error_count"`
	// GradePublishingResults field
	GradePublishingResults int `json:"grade_publishing_results"`
	// GroupMemberships field
	GroupMemberships int `json:"group_memberships"`
	// Groups field
	Groups int `json:"groups"`
	// Sections field
	Sections int `json:"sections"`
	// Terms field
	Terms int `json:"terms"`
	// Users field
	Users int `json:"users"`
	// WarningCount field
	WarningCount int `json:"warning_count"`
	// Xlists field
	Xlists int `json:"xlists"`
}

// SisImport model object
type SisImport struct {
	// AddSisStickiness field: Whether stickiness was added to the batch changes.
	AddSisStickiness bool `json:"add_sis_stickiness"`
	// BatchMode field: Whether the import was run in batch mode.
	BatchMode bool `json:"batch_mode"`
	// BatchModeTermID field: The term the batch was limited to.
	BatchModeTermID string `json:"batch_mode_term_id"`
	// ClearSisStickiness field: Whether stickiness was cleared.
	ClearSisStickiness bool `json:"clear_sis_stickiness"`
	// CreatedAt field: The date the SIS import was created.
	CreatedAt time.Time `json:"created_at"`
	// CsvAttachments field: An array of CSV files for processing
	CsvAttachments [][]File `json:"csv_attachments"`
	// Data field: data
	Data *SisImportData `json:"data"`
	// DiffedAgainstImportID field: The ID of the SIS Import that this import was diffed against
	DiffedAgainstImportID int `json:"diffed_against_import_id"`
	// DiffingDataSetIdentifier field: The identifier of the data set that this SIS batch diffs against
	DiffingDataSetIdentifier string `json:"diffing_data_set_identifier"`
	// EndedAt field: The date the SIS import finished. Returns null if not finished.
	EndedAt time.Time `json:"ended_at"`
	// ErrorsAttachment field: The errors_attachment api object of the SIS import. Only available if there are errors or
	// warning and import has completed.
	ErrorsAttachment *File `json:"errors_attachment"`
	// ID field: The unique identifier for the SIS import.
	ID int `json:"id"`
	// MultiTermBatchMode field: Enables batch mode against all terms in term file. Requires change_threshold to be set.
	MultiTermBatchMode bool `json:"multi_term_batch_mode"`
	// OverrideSisStickiness field: Whether UI changes were overridden.
	OverrideSisStickiness bool `json:"override_sis_stickiness"`
	// ProcessingErrors field: An array of CSV_file/error_message pairs.
	ProcessingErrors [][]string `json:"processing_errors"`
	// ProcessingWarnings field: Only imports that are complete will get this data. An array of CSV_file/warning_message
	// pairs.
	ProcessingWarnings [][]string `json:"processing_warnings"`
	// Progress field: The progress of the SIS import. The progress will reset when using batch_mode and have a
	// different progress for the cleanup stage
	Progress string `json:"progress"`
	// SkipDeletes field: When set the import will skip any deletes.
	SkipDeletes bool `json:"skip_deletes"`
	// Statistics field: statistics
	Statistics *SisImportStatistics `json:"statistics"`
	// UpdatedAt field: The date the SIS import was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// User field: The user that initiated the sis_batch. See the Users API for details.
	User *User `json:"user"`
	// WorkflowState field: The current state of the SIS import.
	//  'initializing': The SIS import is being created, if this gets stuck in initializing, it will not import and will
	// continue on to next import.
	//  'created': The SIS import has been created.
	//  'importing': The SIS import is currently processing.
	//  'cleanup_batch': The SIS import is currently cleaning up courses, sections, and enrollments not included in the
	// batch for batch_mode imports.
	//  'imported': The SIS import has completed successfully.
	//  'imported_with_messages': The SIS import completed with errors or warnings.
	//  'aborted': The SIS import was aborted.
	//  'failed_with_messages': The SIS import failed with errors.
	//  'failed': The SIS import failed.
	//  'restoring': The SIS import is restoring states of imported items.
	//  'partially_restored': The SIS import is restored some of the states of imported items. This is generally due to
	// passing a param like undelete only.
	//  'restored': The SIS import is restored all of the states of imported items.
	WorkflowState *SisImportWorkflowState `json:"workflow_state"`
}

// Submission model object
type Submission struct {
	// AnonymousID field: A unique short ID identifying this submission without reference to the owning user. Only
	// included if the caller has administrator access for the current account.
	AnonymousID string `json:"anonymous_id"`
	// Assignment field: The submission's assignment (see the assignments API) (optional)
	Assignment *Assignment `json:"assignment"`
	// AssignmentID field: The submission's assignment id
	AssignmentID int `json:"assignment_id"`
	// AssignmentVisible field: Whether the assignment is visible to the user who submitted the assignment. Submissions
	// where `assignment_visible` is false no longer count towards the student's grade and the assignment can no longer
	// be accessed by the student. `assignment_visible` becomes false for submissions that do not have a grade and whose
	// assignment is no longer assigned to the student's section.
	AssignmentVisible bool `json:"assignment_visible"`
	// Attempt field: This is the submission attempt number.
	Attempt int `json:"attempt"`
	// Body field: The content of the submission, if it was submitted directly in a text field.
	Body string `json:"body"`
	// Course field: The submission's course (see the course API) (optional)
	Course *Course `json:"course"`
	// Excused field: Whether the assignment is excused.  Excused assignments have no impact on a user's grade.
	Excused bool `json:"excused"`
	// ExtraAttempts field: Extra submission attempts allowed for the given user and assignment.
	ExtraAttempts float64 `json:"extra_attempts"`
	// Grade field: The grade for the submission, translated into the assignment grading scheme (so a letter grade, for
	// example).
	Grade string `json:"grade"`
	// GradeMatchesCurrentSubmission field: A boolean flag which is false if the student has re-submitted since the
	// submission was last graded.
	GradeMatchesCurrentSubmission bool `json:"grade_matches_current_submission"`
	// GradedAt field
	GradedAt time.Time `json:"graded_at"`
	// GraderID field: The id of the user who graded the submission. This will be null for submissions that haven't been
	// graded yet. It will be a positive number if a real user has graded the submission and a negative number if the
	// submission was graded by a process (e.g. Quiz autograder and autograding LTI tools).  Specifically autograded
	// quizzes set grader_id to the negative of the quiz id.  Submissions autograded by LTI tools set grader_id to the
	// negative of the tool id.
	GraderID int `json:"grader_id"`
	// HTMLURL field: URL to the submission. This will require the user to log in.
	HTMLURL string `json:"html_url"`
	// Late field: Whether the submission was made after the applicable due date
	Late bool `json:"late"`
	// LatePolicyStatus field: The status of the submission in relation to the late policy. Can be late, missing, none,
	// or null.
	LatePolicyStatus string `json:"late_policy_status"`
	// Missing field: Whether the assignment is missing.
	Missing bool `json:"missing"`
	// PointsDeducted field: The amount of points automatically deducted from the score by the missing/late policy for a
	// late or missing assignment.
	PointsDeducted float64 `json:"points_deducted"`
	// PostedAt field: The date this submission was posted to the student, or nil if it has not been posted.
	PostedAt time.Time `json:"posted_at"`
	// PreviewURL field: URL to the submission preview. This will require the user to log in.
	PreviewURL string `json:"preview_url"`
	// Score field: The raw score
	Score float64 `json:"score"`
	// SecondsLate field: The amount of time, in seconds, that an submission is late by.
	SecondsLate float64 `json:"seconds_late"`
	// SubmissionComments field: Associated comments for a submission (optional)
	SubmissionComments []SubmissionComment `json:"submission_comments"`
	// SubmissionType field: The types of submission ex:
	// ('online_text_entry'|'online_url'|'online_upload'|'media_recording')
	SubmissionType *SubmissionSubmissionType `json:"submission_type"`
	// SubmittedAt field: The timestamp when the assignment was submitted
	SubmittedAt time.Time `json:"submitted_at"`
	// URL field: The URL of the submission (for 'online_url' submissions).
	URL string `json:"url"`
	// User field: The submissions user (see user API) (optional)
	User *User `json:"user"`
	// UserID field: The id of the user who created the submission
	UserID int `json:"user_id"`
	// WorkflowState field: The current state of the submission
	WorkflowState *SubmissionWorkflowState `json:"workflow_state"`
}

// MediaComment model object
type MediaComment struct {
	// ContentType field
	ContentType string `json:"content-type"`
	// DisplayName field
	DisplayName string `json:"display_name"`
	// MediaID field
	MediaID string `json:"media_id"`
	// MediaType field
	MediaType string `json:"media_type"`
	// URL field
	URL string `json:"url"`
}

// SubmissionComment model object
type SubmissionComment struct {
	// Author field: Abbreviated user object UserDisplay (see users API).
	Author string `json:"author"`
	// AuthorID field
	AuthorID int `json:"author_id"`
	// AuthorName field
	AuthorName string `json:"author_name"`
	// Comment field
	Comment string `json:"comment"`
	// CreatedAt field
	CreatedAt time.Time `json:"created_at"`
	// EditedAt field
	EditedAt time.Time `json:"edited_at"`
	// ID field
	ID int `json:"id"`
	// MediaComment field
	MediaComment *MediaComment `json:"media_comment"`
}

// Tab model object
type Tab struct {
	// Hidden field: only included if true
	Hidden bool `json:"hidden"`
	// HTMLURL field
	HTMLURL string `json:"html_url"`
	// ID field
	ID string `json:"id"`
	// Label field
	Label string `json:"label"`
	// Position field: 1 based
	Position int `json:"position"`
	// Type field
	Type string `json:"type"`
	// Visibility field: possible values are: public, members, admins, and none
	Visibility string `json:"visibility"`
}

// EnrollmentTerm model object
type EnrollmentTerm struct {
	// EndAt field: The datetime of the end of the term.
	EndAt time.Time `json:"end_at"`
	// ID field: The unique identifier for the enrollment term.
	ID int `json:"id"`
	// Name field: The name of the term.
	Name string `json:"name"`
	// Overrides field: Term date overrides for specific enrollment types
	Overrides map[interface{}]interface{} `json:"overrides"`
	// SisImportID field: the unique identifier for the SIS import. This field is only included if the user has
	// permission to manage SIS information.
	SisImportID int `json:"sis_import_id"`
	// SisTermID field: The SIS id of the term. Only included if the user has permission to view SIS information.
	SisTermID string `json:"sis_term_id"`
	// StartAt field: The datetime of the start of the term.
	StartAt time.Time `json:"start_at"`
	// WorkflowState field: The state of the term. Can be 'active' or 'deleted'.
	WorkflowState string `json:"workflow_state"`
}

// EnrollmentTermsList model object
type EnrollmentTermsList struct {
	// EnrollmentTerms field: a paginated list of all terms in the account
	EnrollmentTerms []EnrollmentTerm `json:"enrollment_terms"`
}

// UsageRights model object: Describes the copyright and license information for a File
type UsageRights struct {
	// FileIds field: List of ids of files that were updated
	FileIds []int `json:"file_ids"`
	// LegalCopyright field: Copyright line for the file
	LegalCopyright string `json:"legal_copyright"`
	// License field: License identifier for the file.
	License string `json:"license"`
	// LicenseName field: Readable license name
	LicenseName string `json:"license_name"`
	// Message field: Explanation of the action performed
	Message string `json:"message"`
	// UseJustification field: Justification for using the file in a Canvas course. Valid values are 'own_copyright',
	// 'public_domain', 'used_by_permission', 'fair_use', 'creative_commons'
	UseJustification string `json:"use_justification"`
}

// License model object
type License struct {
	// ID field: a short string identifying the license
	ID string `json:"id"`
	// Name field: the name of the license
	Name string `json:"name"`
	// URL field: a link to the license text
	URL string `json:"url"`
}

// UserDisplay model object: This mini-object is used for secondary user responses, when we just want to provide enough
// information to display a user.
type UserDisplay struct {
	// AvatarImageURL field: If avatars are enabled, this field will be included and contain a url to retrieve the
	// user's avatar.
	AvatarImageURL string `json:"avatar_image_url"`
	// HTMLURL field: URL to access user, either nested to a context or directly.
	HTMLURL string `json:"html_url"`
	// ID field: The ID of the user.
	ID int `json:"id"`
	// ShortName field: A short name the user has selected, for use in conversations or other less formal places through
	// the site.
	ShortName string `json:"short_name"`
}

// AnonymousUserDisplay model object: This mini-object is returned in place of UserDisplay when returning student data
// for anonymous assignments, and includes an anonymous ID to identify a user within the scope of a single assignment.
type AnonymousUserDisplay struct {
	// AnonymousID field: A unique short ID identifying this user within the scope of a particular assignment.
	AnonymousID string `json:"anonymous_id"`
	// AvatarImageURL field: A URL to retrieve a generic avatar.
	AvatarImageURL string `json:"avatar_image_url"`
}

// User model object: A Canvas user, e.g. a student, teacher, administrator, observer, etc.
type User struct {
	// AvatarURL field: If avatars are enabled, this field will be included and contain a url to retrieve the user's
	// avatar.
	AvatarURL string `json:"avatar_url"`
	// Bio field: Optional: The user's bio.
	Bio string `json:"bio"`
	// Email field: Optional: This field can be requested with certain API calls, and will return the users primary
	// email address.
	Email string `json:"email"`
	// Enrollments field: Optional: This field can be requested with certain API calls, and will return a list of the
	// users active enrollments. See the List enrollments API for more details about the format of these records.
	Enrollments []Enrollment `json:"enrollments"`
	// ID field: The ID of the user.
	ID int `json:"id"`
	// IntegrationID field: The integration_id associated with the user.  This field is only included if the user came
	// from a SIS import and has permissions to view SIS information.
	IntegrationID string `json:"integration_id"`
	// LastLogin field: Optional: This field is only returned in certain API calls, and will return a timestamp
	// representing the last time the user logged in to canvas.
	LastLogin string `json:"last_login"`
	// Locale field: Optional: This field can be requested with certain API calls, and will return the users locale in
	// RFC 5646 format.
	Locale string `json:"locale"`
	// LoginID field: The unique login id for the user.  This is what the user uses to log in to Canvas.
	LoginID string `json:"login_id"`
	// Name field: The name of the user.
	Name string `json:"name"`
	// ShortName field: A short name the user has selected, for use in conversations or other less formal places through
	// the site.
	ShortName string `json:"short_name"`
	// SisImportID field: The id of the SIS import.  This field is only included if the user came from a SIS import and
	// has permissions to manage SIS information.
	SisImportID int `json:"sis_import_id"`
	// SisUserID field: The SIS ID associated with the user.  This field is only included if the user came from a SIS
	// import and has permissions to view SIS information.
	SisUserID string `json:"sis_user_id"`
	// SortableName field: The name of the user that is should be used for sorting groups of users, such as in the
	// gradebook.
	SortableName string `json:"sortable_name"`
	// TimeZone field: Optional: This field is only returned in certain API calls, and will return the IANA time zone
	// name of the user's preferred timezone.
	TimeZone string `json:"time_zone"`
}

// Page model object
type Page struct {
	// Body field: the page content, in HTML (present when requesting a single page; omitted when listing pages)
	Body string `json:"body"`
	// CreatedAt field: the creation date for the page
	CreatedAt time.Time `json:"created_at"`
	// EditingRoles field: roles allowed to edit the page; comma-separated list comprising a combination of 'teachers',
	// 'students', 'members', and/or 'public' if not supplied, course defaults are used
	EditingRoles string `json:"editing_roles"`
	// FrontPage field: whether this page is the front page for the wiki
	FrontPage bool `json:"front_page"`
	// HideFromStudents field: (DEPRECATED) whether this page is hidden from students (note: this is always reflected as
	// the inverse of the published value)
	HideFromStudents bool `json:"hide_from_students"`
	// LastEditedBy field: the User who last edited the page (this may not be present if the page was imported from
	// another system)
	LastEditedBy *User `json:"last_edited_by"`
	// LockExplanation field: (Optional) An explanation of why this is locked for the user. Present when locked_for_user
	// is true.
	LockExplanation string `json:"lock_explanation"`
	// LockInfo field: (Optional) Information for the user about the lock. Present when locked_for_user is true.
	LockInfo *LockInfo `json:"lock_info"`
	// LockedForUser field: Whether or not this is locked for the user.
	LockedForUser bool `json:"locked_for_user"`
	// Published field: whether the page is published (true) or draft state (false).
	Published bool `json:"published"`
	// Title field: the title of the page
	Title string `json:"title"`
	// UpdatedAt field: the date the page was last updated
	UpdatedAt time.Time `json:"updated_at"`
	// URL field: the unique locator for the page
	URL string `json:"url"`
}

// PageRevision model object
type PageRevision struct {
	// Body field: the historic page contents
	Body string `json:"body"`
	// EditedBy field: the User who saved this revision, if applicable (this may not be present if the page was imported
	// from another system)
	EditedBy *User `json:"edited_by"`
	// Latest field: whether this is the latest revision or not
	Latest bool `json:"latest"`
	// RevisionID field: an identifier for this revision of the page
	RevisionID int `json:"revision_id"`
	// Title field: the historic page title
	Title string `json:"title"`
	// UpdatedAt field: the time when this revision was saved
	UpdatedAt time.Time `json:"updated_at"`
	// URL field: the following fields are not included in the index action and may be omitted from the show action via
	// summary=1 the historic url of the page
	URL string `json:"url"`
}

// AccountNotificationsIndexOfActiveGlobalNotificationForTheUser API call: Returns a list of all global notifications in
// the account for the current user Any notifications that have been closed by the user will not be returned
func (c *Canvas) AccountNotificationsIndexOfActiveGlobalNotificationForTheUser(progress *task.Progress) ([]AccountNotification, error) {
	endpoint := fmt.Sprintf("accounts/2/users/self/account_notifications")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]AccountNotification{}
	}
	var res []AccountNotification
	callback := func(obj interface{}) error {
		arr := *obj.(*[]AccountNotification)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountNotificationsShowAGlobalNotification API call: Returns a global notification for the current user A
// notification that has been closed by the user will not be returned
func (c *Canvas) AccountNotificationsShowAGlobalNotification(progress *task.Progress) (*AccountNotification, error) {
	endpoint := fmt.Sprintf("accounts/2/users/self/account_notifications/4")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &AccountNotification{}
	}
	var res *AccountNotification
	callback := func(obj interface{}) error {
		res = obj.(*AccountNotification)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountNotificationsCloseNotificationForUser API call: If the current user no long wants to see this notification it
// can be excused with this call
func (c *Canvas) AccountNotificationsCloseNotificationForUser(progress *task.Progress) (*AccountNotification, error) {
	endpoint := fmt.Sprintf("accounts/2/users/self/account_notifications/4")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &AccountNotification{}
	}
	var res *AccountNotification
	callback := func(obj interface{}) error {
		res = obj.(*AccountNotification)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountNotificationsCreateAGlobalNotification API call: Create and return a new global notification for an account.
func (c *Canvas) AccountNotificationsCreateAGlobalNotification(progress *task.Progress, accountNotification *string, accountNotificationRoles *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("accounts/2/account_notifications")
	params := map[string]interface{}{}
	if accountNotification != nil {
		params["account_notification"] = *accountNotification
	}
	if accountNotificationRoles != nil {
		params["account_notification_roles"] = *accountNotificationRoles
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountNotificationsUpdateAGlobalNotification API call: Update global notification for an account.
func (c *Canvas) AccountNotificationsUpdateAGlobalNotification(progress *task.Progress, accountNotification *string, accountNotificationRoles *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("accounts/2/account_notifications/1")
	params := map[string]interface{}{}
	if accountNotification != nil {
		params["account_notification"] = *accountNotification
	}
	if accountNotificationRoles != nil {
		params["account_notification_roles"] = *accountNotificationRoles
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountReportsListAvailableReports API call: Returns a paginated list of reports for the current context.
func (c *Canvas) AccountReportsListAvailableReports(progress *task.Progress, accountID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("accounts/%s/reports/", accountID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountReportsStartAReport API call: Generates a report instance for the account. Note that "report" in the request
// must match one of the available report names. To fetch a list of available report names and parameters for each
// report (including whether or not those parameters are required), see {api:AccountReportsController#available_reports
// List Available Reports}.
func (c *Canvas) AccountReportsStartAReport(progress *task.Progress, parameters *interface{}) (*Report, error) {
	endpoint := fmt.Sprintf("accounts/1/reports/provisioning_csv")
	params := map[string]interface{}{}
	if parameters != nil {
		params["parameters"] = *parameters
	}
	responseCtor := func() interface{} {
		return &Report{}
	}
	var res *Report
	callback := func(obj interface{}) error {
		res = obj.(*Report)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountReportsIndexOfReports API call: Shows all reports that have been run for the account of a specific type.
func (c *Canvas) AccountReportsIndexOfReports(progress *task.Progress, accountID string, reportType string) ([]Report, error) {
	endpoint := fmt.Sprintf("accounts/%s/reports/%s", accountID, reportType)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Report{}
	}
	var res []Report
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Report)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountReportsStatusOfAReport API call: Returns the status of a report.
func (c *Canvas) AccountReportsStatusOfAReport(progress *task.Progress, accountID string, reportType string, reportID string) (*Report, error) {
	endpoint := fmt.Sprintf("accounts/%s/reports/%s/%s", accountID, reportType, reportID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Report{}
	}
	var res *Report
	callback := func(obj interface{}) error {
		res = obj.(*Report)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountReportsDeleteAReport API call: Deletes a generated report instance.
func (c *Canvas) AccountReportsDeleteAReport(progress *task.Progress, accountID string, reportType string, id string) (*Report, error) {
	endpoint := fmt.Sprintf("accounts/%s/reports/%s/%s", accountID, reportType, id)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Report{}
	}
	var res *Report
	callback := func(obj interface{}) error {
		res = obj.(*Report)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsListAccounts API call: A paginated list of accounts that the current user can view or manage. Typically,
// students and even teachers will get an empty list in response, only account admins can view the accounts that they
// are in.
func (c *Canvas) AccountsListAccounts(progress *task.Progress, include *AccountsListAccountsInclude) ([]Account, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]Account{}
	}
	var res []Account
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Account)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsListAccountsForCourseAdmins API call: A paginated list of accounts that the current user can view through
// their admin course enrollments. (Teacher, TA, or designer enrollments). Only returns "id", "name", "workflow_state",
// "root_account_id" and "parent_account_id"
func (c *Canvas) AccountsListAccountsForCourseAdmins(progress *task.Progress) ([]Account, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Account{}
	}
	var res []Account
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Account)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsGetASingleAccount API call: Retrieve information on an individual account, given by id or sis sis_account_id.
func (c *Canvas) AccountsGetASingleAccount(progress *task.Progress) (*Account, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Account{}
	}
	var res *Account
	callback := func(obj interface{}) error {
		res = obj.(*Account)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsPermissions API call: Returns permission information for the calling user and the given account. You may use
// `self` as the account id to check permissions against the domain root account. The caller must have an account role
// or admin (teacher/TA/designer) enrollment in a course in the account. See also the {api:CoursesController#permissions
// Course} and {api:GroupsController#permissions Group} counterparts.
func (c *Canvas) AccountsPermissions(progress *task.Progress, permissions *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("accounts/self/permissions")
	params := map[string]interface{}{}
	if permissions != nil {
		params["permissions"] = *permissions
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsGetTheSubAccountsOfAnAccount API call: List accounts that are sub-accounts of the given account.
func (c *Canvas) AccountsGetTheSubAccountsOfAnAccount(progress *task.Progress, recursive *interface{}, accountID string) ([]Account, error) {
	endpoint := fmt.Sprintf("accounts/%s/sub_accounts", accountID)
	params := map[string]interface{}{}
	if recursive != nil {
		params["recursive"] = *recursive
	}
	responseCtor := func() interface{} {
		return &[]Account{}
	}
	var res []Account
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Account)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsGetTheTermsOfService API call: Returns the terms of service for that account
func (c *Canvas) AccountsGetTheTermsOfService(progress *task.Progress) (*TermsOfService, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &TermsOfService{}
	}
	var res *TermsOfService
	callback := func(obj interface{}) error {
		res = obj.(*TermsOfService)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsGetHelpLinks API call: Returns the help links for that account
func (c *Canvas) AccountsGetHelpLinks(progress *task.Progress) (*HelpLinks, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &HelpLinks{}
	}
	var res *HelpLinks
	callback := func(obj interface{}) error {
		res = obj.(*HelpLinks)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsListActiveCoursesInAnAccount API call: Retrieve a paginated list of courses in this account.
func (c *Canvas) AccountsListActiveCoursesInAnAccount(progress *task.Progress, withEnrollments *bool, enrollmentType *AccountsListActiveCoursesInAnAccountEnrollmentType, published *bool, completed *bool, blueprint *bool, blueprintAssociated *bool, byTeachers *int, bySubaccounts *int, hideEnrollmentlessCourses *bool, state *AccountsListActiveCoursesInAnAccountState, enrollmentTermID *int, searchTerm *string, include *AccountsListActiveCoursesInAnAccountInclude, sort *AccountsListActiveCoursesInAnAccountSort, order *AccountsListActiveCoursesInAnAccountOrder, searchBy *AccountsListActiveCoursesInAnAccountSearchBy, startsBefore *time.Time, endsAfter *time.Time) ([]Course, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if withEnrollments != nil {
		params["with_enrollments"] = *withEnrollments
	}
	if enrollmentType != nil {
		params["enrollment_type"] = *enrollmentType
	}
	if published != nil {
		params["published"] = *published
	}
	if completed != nil {
		params["completed"] = *completed
	}
	if blueprint != nil {
		params["blueprint"] = *blueprint
	}
	if blueprintAssociated != nil {
		params["blueprint_associated"] = *blueprintAssociated
	}
	if byTeachers != nil {
		params["by_teachers"] = *byTeachers
	}
	if bySubaccounts != nil {
		params["by_subaccounts"] = *bySubaccounts
	}
	if hideEnrollmentlessCourses != nil {
		params["hide_enrollmentless_courses"] = *hideEnrollmentlessCourses
	}
	if state != nil {
		params["state"] = *state
	}
	if enrollmentTermID != nil {
		params["enrollment_term_id"] = *enrollmentTermID
	}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if include != nil {
		params["include"] = *include
	}
	if sort != nil {
		params["sort"] = *sort
	}
	if order != nil {
		params["order"] = *order
	}
	if searchBy != nil {
		params["search_by"] = *searchBy
	}
	if startsBefore != nil {
		params["starts_before"] = *startsBefore
	}
	if endsAfter != nil {
		params["ends_after"] = *endsAfter
	}
	responseCtor := func() interface{} {
		return &[]Course{}
	}
	var res []Course
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Course)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsUpdateAnAccount API call: Update an existing account.
func (c *Canvas) AccountsUpdateAnAccount(progress *task.Progress, account *string, accountID string) (*Account, error) {
	endpoint := fmt.Sprintf("accounts/%s", accountID)
	params := map[string]interface{}{}
	if account != nil {
		params["account"] = *account
	}
	responseCtor := func() interface{} {
		return &Account{}
	}
	var res *Account
	callback := func(obj interface{}) error {
		res = obj.(*Account)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsDeleteAUserFromTheRootAccount API call: Delete a user record from a Canvas root account. If a user is
// associated with multiple root accounts (in a multi-tenant instance of Canvas), this action will NOT remove them from
// the other accounts. WARNING: This API will allow a user to remove themselves from the account. If they do this, they
// won't be able to make API calls or log into Canvas at that account.
func (c *Canvas) AccountsDeleteAUserFromTheRootAccount(progress *task.Progress) (*User, error) {
	endpoint := fmt.Sprintf("accounts/3/users/5")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AdminsMakeAnAccountAdmin API call: Flag an existing user as an admin within the account.
func (c *Canvas) AdminsMakeAnAccountAdmin(progress *task.Progress, userID *int, role *string, roleID *int, sendConfirmation *bool) (*Admin, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if userID != nil {
		params["user_id"] = *userID
	}
	if role != nil {
		params["role"] = *role
	}
	if roleID != nil {
		params["role_id"] = *roleID
	}
	if sendConfirmation != nil {
		params["send_confirmation"] = *sendConfirmation
	}
	responseCtor := func() interface{} {
		return &Admin{}
	}
	var res *Admin
	callback := func(obj interface{}) error {
		res = obj.(*Admin)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AdminsRemoveAccountAdmin API call: Remove the rights associated with an account admin role from a user.
func (c *Canvas) AdminsRemoveAccountAdmin(progress *task.Progress, role *string, roleID *int) (*Admin, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if role != nil {
		params["role"] = *role
	}
	if roleID != nil {
		params["role_id"] = *roleID
	}
	responseCtor := func() interface{} {
		return &Admin{}
	}
	var res *Admin
	callback := func(obj interface{}) error {
		res = obj.(*Admin)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AdminsListAccountAdmins API call: A paginated list of the admins in the account
func (c *Canvas) AdminsListAccountAdmins(progress *task.Progress, userID *interface{}) ([]Admin, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if userID != nil {
		params["user_id"] = *userID
	}
	responseCtor := func() interface{} {
		return &[]Admin{}
	}
	var res []Admin
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Admin)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AnnouncementsListAnnouncements API call: Returns the paginated list of announcements for the given courses and date
// range.  Note that a +context_code+ field is added to the responses so you can tell which course each announcement
// belongs to.
func (c *Canvas) AnnouncementsListAnnouncements(progress *task.Progress, contextCodes []string, startDate *time.Time, endDate *time.Time, activeOnly *bool, include []interface{}) ([]DiscussionTopic, error) {
	endpoint := fmt.Sprintf("announcements")
	params := map[string]interface{}{}
	if contextCodes != nil && len(contextCodes) > 0 {
		params["context_codes"] = contextCodes
	}
	if startDate != nil {
		params["start_date"] = *startDate
	}
	if endDate != nil {
		params["end_date"] = *endDate
	}
	if activeOnly != nil {
		params["active_only"] = *activeOnly
	}
	if include != nil && len(include) > 0 {
		params["include"] = include
	}
	responseCtor := func() interface{} {
		return &[]DiscussionTopic{}
	}
	var res []DiscussionTopic
	callback := func(obj interface{}) error {
		arr := *obj.(*[]DiscussionTopic)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModeratedGradingShowProvisionalGradeStatusForAStudent API call: Determine whether or not the student's submission
// needs one or more provisional grades.
func (c *Canvas) ModeratedGradingShowProvisionalGradeStatusForAStudent(progress *task.Progress, anonymousID *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/1/assignments/2/anonymous_provisional_grades/status")
	params := map[string]interface{}{}
	if anonymousID != nil {
		params["anonymous_id"] = *anonymousID
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AppointmentGroupsListAppointmentGroups API call: Retrieve the paginated list of appointment groups that can be
// reserved or managed by the current user.
func (c *Canvas) AppointmentGroupsListAppointmentGroups(progress *task.Progress, scope *AppointmentGroupsListAppointmentGroupsScope, contextCodes *string, includePastAppointments *bool, include *AppointmentGroupsListAppointmentGroupsInclude) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if scope != nil {
		params["scope"] = *scope
	}
	if contextCodes != nil {
		params["context_codes"] = *contextCodes
	}
	if includePastAppointments != nil {
		params["include_past_appointments"] = *includePastAppointments
	}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AppointmentGroupsCreateAnAppointmentGroup API call: Create and return a new appointment group. If new_appointments
// are specified, the response will return a new_appointments array (same format as appointments array, see "List
// appointment groups" action)
func (c *Canvas) AppointmentGroupsCreateAnAppointmentGroup(progress *task.Progress, appointmentGroup *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("appointment_groups.json")
	params := map[string]interface{}{}
	if appointmentGroup != nil {
		params["appointment_group"] = *appointmentGroup
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AppointmentGroupsGetASingleAppointmentGroup API call: Returns information for a single appointment group
func (c *Canvas) AppointmentGroupsGetASingleAppointmentGroup(progress *task.Progress, include *AppointmentGroupsGetASingleAppointmentGroupInclude) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AppointmentGroupsUpdateAnAppointmentGroup API call: Update and return an appointment group. If new_appointments are
// specified, the response will return a new_appointments array (same format as appointments array, see "List
// appointment groups" action).
func (c *Canvas) AppointmentGroupsUpdateAnAppointmentGroup(progress *task.Progress, appointmentGroup *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("appointment_groups/543.json")
	params := map[string]interface{}{}
	if appointmentGroup != nil {
		params["appointment_group"] = *appointmentGroup
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AppointmentGroupsDeleteAnAppointmentGroup API call: Delete an appointment group (and associated time slots and
// reservations) and return the deleted group
func (c *Canvas) AppointmentGroupsDeleteAnAppointmentGroup(progress *task.Progress, cancelReason *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("appointment_groups/543.json")
	params := map[string]interface{}{}
	if cancelReason != nil {
		params["cancel_reason"] = *cancelReason
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AppointmentGroupsListUserParticipants API call: A paginated list of users that are (or may be) participating in this
// appointment group.  Refer to the Users API for the response fields. Returns no results for appointment groups with
// the "Group" participant_type.
func (c *Canvas) AppointmentGroupsListUserParticipants(progress *task.Progress, registrationStatus *AppointmentGroupsListUserParticipantsRegistrationStatus) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if registrationStatus != nil {
		params["registration_status"] = *registrationStatus
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AppointmentGroupsListStudentGroupParticipants API call: A paginated list of student groups that are (or may be)
// participating in this appointment group. Refer to the Groups API for the response fields. Returns no results for
// appointment groups with the "User" participant_type.
func (c *Canvas) AppointmentGroupsListStudentGroupParticipants(progress *task.Progress, registrationStatus *AppointmentGroupsListStudentGroupParticipantsRegistrationStatus) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if registrationStatus != nil {
		params["registration_status"] = *registrationStatus
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AppointmentGroupsGetNextAppointment API call: Return the next appointment available to sign up for. The appointment
// is returned in a one-element array. If no future appointments are available, an empty array is returned.
func (c *Canvas) AppointmentGroupsGetNextAppointment(progress *task.Progress, appointmentGroupIds *string) ([]CalendarEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if appointmentGroupIds != nil {
		params["appointment_group_ids"] = *appointmentGroupIds
	}
	responseCtor := func() interface{} {
		return &[]CalendarEvent{}
	}
	var res []CalendarEvent
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CalendarEvent)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentExtensionsSetExtensionsForStudentAssignmentSubmissions API call
func (c *Canvas) AssignmentExtensionsSetExtensionsForStudentAssignmentSubmissions(progress *task.Progress, assignmentExtensions *int) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if assignmentExtensions != nil {
		params["assignment_extensions"] = *assignmentExtensions
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentGroupsGetAnAssignmentGroup API call: Returns the assignment group with the given id.
func (c *Canvas) AssignmentGroupsGetAnAssignmentGroup(progress *task.Progress, include *AssignmentGroupsGetAnAssignmentGroupInclude, overrideAssignmentDates *bool, gradingPeriodID *int) (*AssignmentGroup, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if overrideAssignmentDates != nil {
		params["override_assignment_dates"] = *overrideAssignmentDates
	}
	if gradingPeriodID != nil {
		params["grading_period_id"] = *gradingPeriodID
	}
	responseCtor := func() interface{} {
		return &AssignmentGroup{}
	}
	var res *AssignmentGroup
	callback := func(obj interface{}) error {
		res = obj.(*AssignmentGroup)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentGroupsCreateAnAssignmentGroup API call: Create a new assignment group for this course.
func (c *Canvas) AssignmentGroupsCreateAnAssignmentGroup(progress *task.Progress, name *string, position *int, groupWeight *float64, sisSourceID *string, integrationData map[string]interface{}, rules *interface{}) (*AssignmentGroup, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if name != nil {
		params["name"] = *name
	}
	if position != nil {
		params["position"] = *position
	}
	if groupWeight != nil {
		params["group_weight"] = *groupWeight
	}
	if sisSourceID != nil {
		params["sis_source_id"] = *sisSourceID
	}
	if integrationData != nil && len(integrationData) > 0 {
		params["integration_data"] = integrationData
	}
	if rules != nil {
		params["rules"] = *rules
	}
	responseCtor := func() interface{} {
		return &AssignmentGroup{}
	}
	var res *AssignmentGroup
	callback := func(obj interface{}) error {
		res = obj.(*AssignmentGroup)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentGroupsEditAnAssignmentGroup API call: Modify an existing Assignment Group. Accepts the same parameters as
// Assignment Group creation
func (c *Canvas) AssignmentGroupsEditAnAssignmentGroup(progress *task.Progress) (*AssignmentGroup, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &AssignmentGroup{}
	}
	var res *AssignmentGroup
	callback := func(obj interface{}) error {
		res = obj.(*AssignmentGroup)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentGroupsDestroyAnAssignmentGroup API call: Deletes the assignment group with the given id.
func (c *Canvas) AssignmentGroupsDestroyAnAssignmentGroup(progress *task.Progress, moveAssignmentsTo *int) (*AssignmentGroup, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if moveAssignmentsTo != nil {
		params["move_assignments_to"] = *moveAssignmentsTo
	}
	responseCtor := func() interface{} {
		return &AssignmentGroup{}
	}
	var res *AssignmentGroup
	callback := func(obj interface{}) error {
		res = obj.(*AssignmentGroup)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentGroupsListAssignmentGroups API call: Returns the paginated list of assignment groups for the current
// context. The returned groups are sorted by their position field.
func (c *Canvas) AssignmentGroupsListAssignmentGroups(progress *task.Progress, include *AssignmentGroupsListAssignmentGroupsInclude, excludeAssignmentSubmissionTypes *AssignmentGroupsListAssignmentGroupsExcludeAssignmentSubmissionTypes, overrideAssignmentDates *bool, gradingPeriodID *int, scopeAssignmentsToStudent *bool) ([]AssignmentGroup, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if excludeAssignmentSubmissionTypes != nil {
		params["exclude_assignment_submission_types"] = *excludeAssignmentSubmissionTypes
	}
	if overrideAssignmentDates != nil {
		params["override_assignment_dates"] = *overrideAssignmentDates
	}
	if gradingPeriodID != nil {
		params["grading_period_id"] = *gradingPeriodID
	}
	if scopeAssignmentsToStudent != nil {
		params["scope_assignments_to_student"] = *scopeAssignmentsToStudent
	}
	responseCtor := func() interface{} {
		return &[]AssignmentGroup{}
	}
	var res []AssignmentGroup
	callback := func(obj interface{}) error {
		arr := *obj.(*[]AssignmentGroup)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsListAssignmentOverrides API call: Returns the paginated list of overrides for this assignment that target
// sections/groups/students visible to the current user.
func (c *Canvas) AssignmentsListAssignmentOverrides(progress *task.Progress) ([]AssignmentOverride, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]AssignmentOverride{}
	}
	var res []AssignmentOverride
	callback := func(obj interface{}) error {
		arr := *obj.(*[]AssignmentOverride)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsGetASingleAssignmentOverride API call: Returns details of the the override with the given id.
func (c *Canvas) AssignmentsGetASingleAssignmentOverride(progress *task.Progress) (*AssignmentOverride, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &AssignmentOverride{}
	}
	var res *AssignmentOverride
	callback := func(obj interface{}) error {
		res = obj.(*AssignmentOverride)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsRedirectToTheAssignmentOverrideForAGroup API call: Responds with a redirect to the override for the given
// group, if any (404 otherwise).
func (c *Canvas) AssignmentsRedirectToTheAssignmentOverrideForAGroup(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsRedirectToTheAssignmentOverrideForASection API call: Responds with a redirect to the override for the
// given section, if any (404 otherwise).
func (c *Canvas) AssignmentsRedirectToTheAssignmentOverrideForASection(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsCreateAnAssignmentOverride API call
func (c *Canvas) AssignmentsCreateAnAssignmentOverride(progress *task.Progress, assignmentOverride *interface{}) (*AssignmentOverride, error) {
	endpoint := fmt.Sprintf("courses/1/assignments/2/overrides.json")
	params := map[string]interface{}{}
	if assignmentOverride != nil {
		params["assignment_override"] = *assignmentOverride
	}
	responseCtor := func() interface{} {
		return &AssignmentOverride{}
	}
	var res *AssignmentOverride
	callback := func(obj interface{}) error {
		res = obj.(*AssignmentOverride)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsUpdateAnAssignmentOverride API call
func (c *Canvas) AssignmentsUpdateAnAssignmentOverride(progress *task.Progress, assignmentOverride *interface{}) (*AssignmentOverride, error) {
	endpoint := fmt.Sprintf("courses/1/assignments/2/overrides/3.json")
	params := map[string]interface{}{}
	if assignmentOverride != nil {
		params["assignment_override"] = *assignmentOverride
	}
	responseCtor := func() interface{} {
		return &AssignmentOverride{}
	}
	var res *AssignmentOverride
	callback := func(obj interface{}) error {
		res = obj.(*AssignmentOverride)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsDeleteAnAssignmentOverride API call: Deletes an override and returns its former details.
func (c *Canvas) AssignmentsDeleteAnAssignmentOverride(progress *task.Progress) (*AssignmentOverride, error) {
	endpoint := fmt.Sprintf("courses/1/assignments/2/overrides/3.json")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &AssignmentOverride{}
	}
	var res *AssignmentOverride
	callback := func(obj interface{}) error {
		res = obj.(*AssignmentOverride)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsBatchRetrieveOverridesInACourse API call: Returns a list of specified overrides in this course, providing
// they target sections/groups/students visible to the current user. Returns null elements in the list for requests that
// were not found.
func (c *Canvas) AssignmentsBatchRetrieveOverridesInACourse(progress *task.Progress, assignmentOverrides *interface{}) ([]AssignmentOverride, error) {
	endpoint := fmt.Sprintf("courses/12/assignments/overrides.json")
	params := map[string]interface{}{}
	if assignmentOverrides != nil {
		params["assignment_overrides"] = *assignmentOverrides
	}
	responseCtor := func() interface{} {
		return &[]AssignmentOverride{}
	}
	var res []AssignmentOverride
	callback := func(obj interface{}) error {
		arr := *obj.(*[]AssignmentOverride)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsBatchCreateOverridesInACourse API call: Creates the specified overrides for each assignment.  Handles
// creation in a transaction, so all records are created or none are. One of student_ids, group_id, or course_section_id
// must be present. At most one should be present; if multiple are present only the most specific (student_ids first,
// then group_id, then course_section_id) is used and any others are ignored. Errors are reported in an errors
// attribute, an array of errors corresponding to inputs.  Global errors will be reported as a single element errors
// array
func (c *Canvas) AssignmentsBatchCreateOverridesInACourse(progress *task.Progress, assignmentOverrides *interface{}) ([]AssignmentOverride, error) {
	endpoint := fmt.Sprintf("courses/12/assignments/overrides.json")
	params := map[string]interface{}{}
	if assignmentOverrides != nil {
		params["assignment_overrides"] = *assignmentOverrides
	}
	responseCtor := func() interface{} {
		return &[]AssignmentOverride{}
	}
	var res []AssignmentOverride
	callback := func(obj interface{}) error {
		arr := *obj.(*[]AssignmentOverride)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsBatchUpdateOverridesInACourse API call: Updates a list of specified overrides for each assignment.
// Handles overrides in a transaction, so either all updates are applied or none. See
// {api:AssignmentOverridesController#update Update an assignment override} for available attributes. All current
// overridden values must be supplied if they are to be retained; e.g. if due_at was overridden, but this PUT omits a
// value for due_at, due_at will no longer be overridden. If the override is adhoc and student_ids is not supplied, the
// target override set is unchanged. Target override sets cannot be changed for group or section overrides. Errors are
// reported in an errors attribute, an array of errors corresponding to inputs.  Global errors will be reported as a
// single element errors array
func (c *Canvas) AssignmentsBatchUpdateOverridesInACourse(progress *task.Progress, assignmentOverrides *interface{}) ([]AssignmentOverride, error) {
	endpoint := fmt.Sprintf("courses/12/assignments/overrides.json")
	params := map[string]interface{}{}
	if assignmentOverrides != nil {
		params["assignment_overrides"] = *assignmentOverrides
	}
	responseCtor := func() interface{} {
		return &[]AssignmentOverride{}
	}
	var res []AssignmentOverride
	callback := func(obj interface{}) error {
		arr := *obj.(*[]AssignmentOverride)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsListAssignments API call: Returns the paginated list of assignments for the current course or assignment
// group.
func (c *Canvas) AssignmentsListAssignments(progress *task.Progress, include *AssignmentsListAssignmentsInclude, searchTerm *string, overrideAssignmentDates *bool, needsGradingCountBySection *bool, bucket *AssignmentsListAssignmentsBucket, assignmentIds *interface{}, orderBy *AssignmentsListAssignmentsOrderBy, postToSis *bool) ([]Assignment, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if overrideAssignmentDates != nil {
		params["override_assignment_dates"] = *overrideAssignmentDates
	}
	if needsGradingCountBySection != nil {
		params["needs_grading_count_by_section"] = *needsGradingCountBySection
	}
	if bucket != nil {
		params["bucket"] = *bucket
	}
	if assignmentIds != nil {
		params["assignment_ids"] = *assignmentIds
	}
	if orderBy != nil {
		params["order_by"] = *orderBy
	}
	if postToSis != nil {
		params["post_to_sis"] = *postToSis
	}
	responseCtor := func() interface{} {
		return &[]Assignment{}
	}
	var res []Assignment
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Assignment)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsListAssignmentsForUser API call: Returns the paginated list of assignments for the specified user if the
// current user has rights to view. See {api:AssignmentsApiController#index List assignments} for valid arguments.
func (c *Canvas) AssignmentsListAssignmentsForUser(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsGetASingleAssignment API call: Returns the assignment with the given id.
func (c *Canvas) AssignmentsGetASingleAssignment(progress *task.Progress, include *AssignmentsGetASingleAssignmentInclude, overrideAssignmentDates *bool, needsGradingCountBySection *bool, allDates *bool) (*Assignment, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if overrideAssignmentDates != nil {
		params["override_assignment_dates"] = *overrideAssignmentDates
	}
	if needsGradingCountBySection != nil {
		params["needs_grading_count_by_section"] = *needsGradingCountBySection
	}
	if allDates != nil {
		params["all_dates"] = *allDates
	}
	responseCtor := func() interface{} {
		return &Assignment{}
	}
	var res *Assignment
	callback := func(obj interface{}) error {
		res = obj.(*Assignment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsCreateAnAssignment API call: Create a new assignment for this course. The assignment is created in the
// active state.
func (c *Canvas) AssignmentsCreateAnAssignment(progress *task.Progress, assignment *interface{}) (*Assignment, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if assignment != nil {
		params["assignment"] = *assignment
	}
	responseCtor := func() interface{} {
		return &Assignment{}
	}
	var res *Assignment
	callback := func(obj interface{}) error {
		res = obj.(*Assignment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsEditAnAssignment API call: Modify an existing assignment.
func (c *Canvas) AssignmentsEditAnAssignment(progress *task.Progress, assignment *interface{}) (*Assignment, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if assignment != nil {
		params["assignment"] = *assignment
	}
	responseCtor := func() interface{} {
		return &Assignment{}
	}
	var res *Assignment
	callback := func(obj interface{}) error {
		res = obj.(*Assignment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsBulkUpdateAssignmentDates API call: Update due dates and availability dates for multiple assignments in a
// course. Accepts a JSON array of objects containing two keys each: +id+, the assignment id, and +all_dates+, an array
// of +AssignmentDate+ structures containing the base and/or override dates for the assignment, as returned from the
// {api:AssignmentsApiController#index List assignments} endpoint with +include[]=all_dates+. This endpoint cannot
// create or destroy assignment overrides; any existing assignment overrides that are not referenced in the arguments
// will be left alone. If an override is given, any dates that are not supplied with it will be defaulted. To clear a
// date, specify null explicitly. All referenced assignments will be validated before any are saved. A list of errors
// will be returned if any provided dates are invalid, and no changes will be saved. The bulk update is performed in a
// background job, use the {api:ProgressController#show Progress API} to check its status.
func (c *Canvas) AssignmentsBulkUpdateAssignmentDates(progress *task.Progress) (*Progress, error) {
	endpoint := fmt.Sprintf("courses/1/assignments/bulk_update")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Progress{}
	}
	var res *Progress
	callback := func(obj interface{}) error {
		res = obj.(*Progress)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignmentsDeleteAnAssignment API call: Delete the given assignment.
func (c *Canvas) AssignmentsDeleteAnAssignment(progress *task.Progress, courseID string, assignmentID string) (*Assignment, error) {
	endpoint := fmt.Sprintf("courses/%s/assignments/%s", courseID, assignmentID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Assignment{}
	}
	var res *Assignment
	callback := func(obj interface{}) error {
		res = obj.(*Assignment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationsLogQueryByLogin API call: List authentication events for a given login.
func (c *Canvas) AuthenticationsLogQueryByLogin(progress *task.Progress, startTime *time.Time, endTime *time.Time) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationsLogQueryByAccount API call: List authentication events for a given account.
func (c *Canvas) AuthenticationsLogQueryByAccount(progress *task.Progress, startTime *time.Time, endTime *time.Time) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationsLogQueryByUser API call: List authentication events for a given user.
func (c *Canvas) AuthenticationsLogQueryByUser(progress *task.Progress, startTime *time.Time, endTime *time.Time) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationProvidersListAuthenticationProviders API call: Returns a paginated list of authentication providers
func (c *Canvas) AuthenticationProvidersListAuthenticationProviders(progress *task.Progress, accountID string) ([]AuthenticationProvider, error) {
	endpoint := fmt.Sprintf("accounts/%s/authentication_providers", accountID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]AuthenticationProvider{}
	}
	var res []AuthenticationProvider
	callback := func(obj interface{}) error {
		arr := *obj.(*[]AuthenticationProvider)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationProvidersAddAuthenticationProvider API call: Add external authentication provider(s) for the account.
// Services may be CAS, Facebook, GitHub, Google, LDAP, LinkedIn, Microsoft, OpenID Connect, SAML, or Twitter. Each
// authentication provider is specified as a set of parameters as described below. A provider specification must include
// an 'auth_type' parameter with a value of 'canvas', 'cas', 'clever', 'facebook', 'github', 'google', 'ldap',
// 'linkedin', 'microsoft', 'openid_connect', 'saml', or 'twitter'. The other recognized parameters depend on this
// auth_type; unrecognized parameters are discarded. Provider specifications not specifying a valid auth_type are
// ignored. You can set the 'position' for any configuration. The config in the 1st position is considered the default.
// You can set 'jit_provisioning' for any configuration besides Canvas. For Canvas, the additional recognized parameter
// is: - self_registration 'all', 'none', or 'observer' - who is allowed to register as a new user For CAS, the
// additional recognized parameters are: - auth_base The CAS server's URL. - log_in_url [Optional] An alternate SSO URL
// for logging into CAS. You probably should not set this. For Clever, the additional recognized parameters are: -
// client_id [Required] The Clever application's Client ID. Not available if configured globally for Canvas. -
// client_secret [Required] The Clever application's Client Secret. Not available if configured globally for Canvas. -
// district_id [Optional] A district's Clever ID. Leave this blank to let Clever handle the details with its District
// Picker. This is required for Clever Instant Login to work in a multi-tenant environment. - login_attribute [Optional]
// The attribute to use to look up the user's login in Canvas. Either 'id' (the default), 'sis_id', 'email',
// 'student_number', or 'teacher_number'. Note that some fields may not be populated for all users at Clever. -
// federated_attributes [Optional] See FederatedAttributesConfig. Valid provider attributes are 'id', 'sis_id', 'email',
// 'student_number', and 'teacher_number'. For Facebook, the additional recognized parameters are: - app_id [Required]
// The Facebook App ID. Not available if configured globally for Canvas. - app_secret [Required] The Facebook App
// Secret. Not available if configured globally for Canvas. - login_attribute [Optional] The attribute to use to look up
// the user's login in Canvas. Either 'id' (the default), or 'email' - federated_attributes [Optional] See
// FederatedAttributesConfig. Valid provider attributes are 'email', 'first_name', 'id', 'last_name', 'locale', and
// 'name'. For GitHub, the additional recognized parameters are: - domain [Optional] The domain of a GitHub Enterprise
// installation. I.e. github.mycompany.com. If not set, it will default to the public github.com. - client_id [Required]
// The GitHub application's Client ID. Not available if configured globally for Canvas. - client_secret [Required] The
// GitHub application's Client Secret. Not available if configured globally for Canvas. - login_attribute [Optional] The
// attribute to use to look up the user's login in Canvas. Either 'id' (the default), or 'login' - federated_attributes
// [Optional] See FederatedAttributesConfig. Valid provider attributes are 'email', 'id', 'login', and 'name'. For
// Google, the additional recognized parameters are: - client_id [Required] The Google application's Client ID. Not
// available if configured globally for Canvas. - client_secret [Required] The Google application's Client Secret. Not
// available if configured globally for Canvas. - hosted_domain [Optional] A Google Apps domain to restrict logins to.
// See https://developers.google.com/identity/protocols/OpenIDConnect?hl=en#hd-param - login_attribute [Optional] The
// attribute to use to look up the user's login in Canvas. Either 'sub' (the default), or 'email' - federated_attributes
// [Optional] See FederatedAttributesConfig. Valid provider attributes are 'email', 'family_name', 'given_name',
// 'locale', 'name', and 'sub'. For LDAP, the additional recognized parameters are: - auth_host The LDAP server's URL. -
// auth_port [Optional, Integer] The LDAP server's TCP port. (default: 389) - auth_over_tls [Optional] Whether to use
// TLS. Can be 'simple_tls', or 'start_tls'. For backwards compatibility, booleans are also accepted, with true meaning
// simple_tls. If not provided, it will default to start_tls. - auth_base [Optional] A default treebase parameter for
// searches performed against the LDAP server. - auth_filter LDAP search filter. Use !{{login}} as a placeholder for the
// username supplied by the user. For example: "(sAMAccountName=!{{login}})". - identifier_format [Optional] The LDAP
// attribute to use to look up the Canvas login. Omit to use the username supplied by the user. - auth_username Username
// - auth_password Password For LinkedIn, the additional recognized parameters are: - client_id [Required] The LinkedIn
// application's Client ID. Not available if configured globally for Canvas. - client_secret [Required] The LinkedIn
// application's Client Secret. Not available if configured globally for Canvas. - login_attribute [Optional] The
// attribute to use to look up the user's login in Canvas. Either 'id' (the default), or 'emailAddress' -
// federated_attributes [Optional] See FederatedAttributesConfig. Valid provider attributes are 'emailAddress',
// 'firstName', 'id', 'formattedName', and 'lastName'. For Microsoft, the additional recognized parameters are: -
// application_id [Required] The application's ID. - application_secret [Required] The application's Client Secret
// (Password) - tenant [Optional] See
// https://azure.microsoft.com/en-us/documentation/articles/active-directory-v2-protocols/ Valid values are 'common',
// 'organizations', 'consumers', or an Azure Active Directory Tenant (as either a UUID or domain, such as
// contoso.onmicrosoft.com). Defaults to 'common' - login_attribute [Optional] See
// https://azure.microsoft.com/en-us/documentation/articles/active-directory-v2-tokens/#idtokens Valid values are 'sub',
// 'email', 'oid', or 'preferred_username'. Note that email may not always be populated in the user's profile at
// Microsoft. Oid will not be populated for personal Microsoft accounts. Defaults to 'sub' - federated_attributes
// [Optional] See FederatedAttributesConfig. Valid provider attributes are 'email', 'name', 'preferred_username', 'oid',
// and 'sub'. For OpenID Connect, the additional recognized parameters are: - client_id [Required] The application's
// Client ID. - client_secret [Required] The application's Client Secret. - authorize_url [Required] The URL for getting
// starting the OAuth 2.0 web flow - token_url [Required] The URL for exchanging the OAuth 2.0 authorization code for an
// Access Token and ID Token - scope [Optional] Space separated additional scopes to request for the token. Note that
// you need not specify the 'openid' scope, or any scopes that can be automatically inferred by the rules defined at
// http://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims - end_session_endpoint [Optional] URL to send the
// end user to after logging out of Canvas. See https://openid.net/specs/openid-connect-session-1_0.html#RPLogout -
// userinfo_endpoint [Optional] URL to request additional claims from. If the initial ID Token received from the
// provider cannot be used to satisfy the login_attribute and all federated_attributes, this endpoint will be queried
// for additional information. - login_attribute [Optional] The attribute of the ID Token to look up the user's login in
// Canvas. Defaults to 'sub'. - federated_attributes [Optional] See FederatedAttributesConfig. Any value is allowed for
// the provider attribute names, but standard claims are listed at
// http://openid.net/specs/openid-connect-core-1_0.html#StandardClaims For SAML, the additional recognized parameters
// are: - metadata [Optional] An XML document to parse as SAML metadata, and automatically populate idp_entity_id,
// log_in_url, log_out_url, certificate_fingerprint, and identifier_format - metadata_uri [Optional] A URI to download
// the SAML metadata from, and automatically populate idp_entity_id, log_in_url, log_out_url, certificate_fingerprint,
// and identifier_format. This URI will also be saved, and the metadata periodically refreshed, automatically. If the
// metadata contains multiple entities, also supply idp_entity_id to distinguish which one you want (otherwise the only
// entity in the metadata will be inferred). If you provide the URI 'urn:mace:incommon' or 'http://ukfederation.org.uk',
// the InCommon or UK Access Management Federation metadata aggregate, respectively, will be used instead, and
// additional validation checks will happen (including validating that the metadata has been properly signed with the
// appropriate key). - idp_entity_id The SAML IdP's entity ID - log_in_url The SAML service's SSO target URL -
// log_out_url [Optional] The SAML service's SLO target URL - certificate_fingerprint The SAML service's certificate
// fingerprint. - identifier_format The SAML service's identifier format. Must be one of: -
// urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress - urn:oasis:names:tc:SAML:2.0:nameid-format:entity -
// urn:oasis:names:tc:SAML:2.0:nameid-format:kerberos - urn:oasis:names:tc:SAML:2.0:nameid-format:persistent -
// urn:oasis:names:tc:SAML:2.0:nameid-format:transient - urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified -
// urn:oasis:names:tc:SAML:1.1:nameid-format:WindowsDomainQualifiedName -
// urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName - requested_authn_context [Optional] The SAML AuthnContext
// - sig_alg [Optional] If set, +AuthnRequest+, +LogoutRequest+, and +LogoutResponse+ messages are signed with the
// corresponding algorithm. Supported algorithms are: - {http://www.w3.org/2000/09/xmldsig#rsa-sha1} -
// {http://www.w3.org/2001/04/xmldsig-more#rsa-sha256} RSA-SHA1 and RSA-SHA256 are acceptable aliases. -
// federated_attributes [Optional] See FederatedAttributesConfig. Any value is allowed for the provider attribute names.
// For Twitter, the additional recognized parameters are: - consumer_key [Required] The Twitter Consumer Key. Not
// available if configured globally for Canvas. - consumer_secret [Required] The Twitter Consumer Secret. Not available
// if configured globally for Canvas. - login_attribute [Optional] The attribute to use to look up the user's login in
// Canvas. Either 'user_id' (the default), or 'screen_name' - parent_registration [Optional] - DEPRECATED 2017-11-03
// Accepts a boolean value, true designates the authentication service for use on parent registrations.  Only one
// service can be selected at a time so if set to true all others will be set to false - federated_attributes [Optional]
// See FederatedAttributesConfig. Valid provider attributes are 'name', 'screen_name', 'time_zone', and 'user_id'.
func (c *Canvas) AuthenticationProvidersAddAuthenticationProvider(progress *task.Progress, accountID string) (*AuthenticationProvider, error) {
	endpoint := fmt.Sprintf("accounts/%s/authentication_providers", accountID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &AuthenticationProvider{}
	}
	var res *AuthenticationProvider
	callback := func(obj interface{}) error {
		res = obj.(*AuthenticationProvider)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationProvidersUpdateAuthenticationProvider API call: Update an authentication provider using the same
// options as the create endpoint. You can not update an existing provider to a new authentication type.
func (c *Canvas) AuthenticationProvidersUpdateAuthenticationProvider(progress *task.Progress, accountID string, id string) (*AuthenticationProvider, error) {
	endpoint := fmt.Sprintf("accounts/%s/authentication_providers/%s", accountID, id)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &AuthenticationProvider{}
	}
	var res *AuthenticationProvider
	callback := func(obj interface{}) error {
		res = obj.(*AuthenticationProvider)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationProvidersGetAuthenticationProvider API call: Get the specified authentication provider
func (c *Canvas) AuthenticationProvidersGetAuthenticationProvider(progress *task.Progress, accountID string, id string) (*AuthenticationProvider, error) {
	endpoint := fmt.Sprintf("accounts/%s/authentication_providers/%s", accountID, id)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &AuthenticationProvider{}
	}
	var res *AuthenticationProvider
	callback := func(obj interface{}) error {
		res = obj.(*AuthenticationProvider)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationProvidersDeleteAuthenticationProvider API call: Delete the config
func (c *Canvas) AuthenticationProvidersDeleteAuthenticationProvider(progress *task.Progress, accountID string, id string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("accounts/%s/authentication_providers/%s", accountID, id)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationProvidersShowAccountAuthSettings API call: The way to get the current state of each account level
// setting that's relevant to Single Sign On configuration You can list the current state of each setting with
// "update_sso_settings"
func (c *Canvas) AuthenticationProvidersShowAccountAuthSettings(progress *task.Progress, accountID string) (*SSOSettings, error) {
	endpoint := fmt.Sprintf("accounts/%s/sso_settings", accountID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &SSOSettings{}
	}
	var res *SSOSettings
	callback := func(obj interface{}) error {
		res = obj.(*SSOSettings)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AuthenticationProvidersUpdateAccountAuthSettings API call: For various cases of mixed SSO configurations, you may
// need to set some configuration at the account level to handle the particulars of your setup. This endpoint accepts a
// PUT request to set several possible account settings. All setting are optional on each request, any that are not
// provided at all are simply retained as is.  Any that provide the key but a null-ish value (blank string, null,
// undefined) will be UN-set. You can list the current state of each setting with "show_sso_settings"
func (c *Canvas) AuthenticationProvidersUpdateAccountAuthSettings(progress *task.Progress, accountID string) (*SSOSettings, error) {
	endpoint := fmt.Sprintf("accounts/%s/sso_settings", accountID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &SSOSettings{}
	}
	var res *SSOSettings
	callback := func(obj interface{}) error {
		res = obj.(*SSOSettings)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// BrandConfigsGetTheBrandConfigVariablesThatShouldBeUsedForThisDomain API call: Will redirect to a static json file
// that has all of the brand variables used by this account. Even though this is a redirect, do not store the redirected
// url since if the account makes any changes it will redirect to a new url. Needs no authentication.
func (c *Canvas) BrandConfigsGetTheBrandConfigVariablesThatShouldBeUsedForThisDomain(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("brand_variables")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsListCalendarEvents API call: Retrieve the paginated list of calendar events or assignments for the
// current user
func (c *Canvas) CalendarEventsListCalendarEvents(progress *task.Progress, typeName *interface{}, startDate *time.Time, endDate *time.Time, undated *bool, allEvents *bool, contextCodes *string, excludes []interface{}) ([]CalendarEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if typeName != nil {
		params["type"] = *typeName
	}
	if startDate != nil {
		params["start_date"] = *startDate
	}
	if endDate != nil {
		params["end_date"] = *endDate
	}
	if undated != nil {
		params["undated"] = *undated
	}
	if allEvents != nil {
		params["all_events"] = *allEvents
	}
	if contextCodes != nil {
		params["context_codes"] = *contextCodes
	}
	if excludes != nil && len(excludes) > 0 {
		params["excludes"] = excludes
	}
	responseCtor := func() interface{} {
		return &[]CalendarEvent{}
	}
	var res []CalendarEvent
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CalendarEvent)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsListCalendarEventsForAUser API call: Retrieve the paginated list of calendar events or assignments for
// the specified user. To view calendar events for a user other than yourself, you must either be an observer of that
// user or an administrator.
func (c *Canvas) CalendarEventsListCalendarEventsForAUser(progress *task.Progress, typeName *interface{}, startDate *time.Time, endDate *time.Time, undated *bool, allEvents *bool, contextCodes *string, excludes []interface{}, submissionTypes []interface{}, excludeSubmissionTypes []interface{}) ([]CalendarEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if typeName != nil {
		params["type"] = *typeName
	}
	if startDate != nil {
		params["start_date"] = *startDate
	}
	if endDate != nil {
		params["end_date"] = *endDate
	}
	if undated != nil {
		params["undated"] = *undated
	}
	if allEvents != nil {
		params["all_events"] = *allEvents
	}
	if contextCodes != nil {
		params["context_codes"] = *contextCodes
	}
	if excludes != nil && len(excludes) > 0 {
		params["excludes"] = excludes
	}
	if submissionTypes != nil && len(submissionTypes) > 0 {
		params["submission_types"] = submissionTypes
	}
	if excludeSubmissionTypes != nil && len(excludeSubmissionTypes) > 0 {
		params["exclude_submission_types"] = excludeSubmissionTypes
	}
	responseCtor := func() interface{} {
		return &[]CalendarEvent{}
	}
	var res []CalendarEvent
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CalendarEvent)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsCreateACalendarEvent API call: Create and return a new calendar event
func (c *Canvas) CalendarEventsCreateACalendarEvent(progress *task.Progress, calendarEvent *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("calendar_events.json")
	params := map[string]interface{}{}
	if calendarEvent != nil {
		params["calendar_event"] = *calendarEvent
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsGetASingleCalendarEventOrAssignment API call
func (c *Canvas) CalendarEventsGetASingleCalendarEventOrAssignment(progress *task.Progress) (*CalendarEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &CalendarEvent{}
	}
	var res *CalendarEvent
	callback := func(obj interface{}) error {
		res = obj.(*CalendarEvent)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsReserveATimeSlot API call: Reserves a particular time slot and return the new reservation
func (c *Canvas) CalendarEventsReserveATimeSlot(progress *task.Progress, participantID *string, comments *string, cancelExisting *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("calendar_events/345/reservations.json")
	params := map[string]interface{}{}
	if participantID != nil {
		params["participant_id"] = *participantID
	}
	if comments != nil {
		params["comments"] = *comments
	}
	if cancelExisting != nil {
		params["cancel_existing"] = *cancelExisting
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsUpdateACalendarEvent API call: Update and return a calendar event
func (c *Canvas) CalendarEventsUpdateACalendarEvent(progress *task.Progress, calendarEvent *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("calendar_events/234.json")
	params := map[string]interface{}{}
	if calendarEvent != nil {
		params["calendar_event"] = *calendarEvent
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsDeleteACalendarEvent API call: Delete an event from the calendar and return the deleted event
func (c *Canvas) CalendarEventsDeleteACalendarEvent(progress *task.Progress, cancelReason *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("calendar_events/234.json")
	params := map[string]interface{}{}
	if cancelReason != nil {
		params["cancel_reason"] = *cancelReason
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsSetACourseTimetable API call: Creates and updates "timetable" events for a course. Can automaticaly
// generate a series of calendar events based on simple schedules (e.g. "Monday and Wednesday at 2:00pm" ) Existing
// timetable events for the course and course sections will be updated if they still are part of the timetable.
// Otherwise, they will be deleted.
func (c *Canvas) CalendarEventsSetACourseTimetable(progress *task.Progress, timetables []interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("calendar_events/timetable")
	params := map[string]interface{}{}
	if timetables != nil && len(timetables) > 0 {
		params["timetables"] = timetables
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsGetCourseTimetable API call: Returns the last timetable set by the
// {api:CalendarEventsApiController#set_course_timetable Set a course timetable} endpoint
func (c *Canvas) CalendarEventsGetCourseTimetable(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CalendarEventsCreateOrUpdateEventsDirectlyForACourseTimetable API call: Creates and updates "timetable" events for a
// course or course section. Similar to {api:CalendarEventsApiController#set_course_timetable setting a course
// timetable}, but instead of generating a list of events based on a timetable schedule, this endpoint expects a
// complete list of events.
func (c *Canvas) CalendarEventsCreateOrUpdateEventsDirectlyForACourseTimetable(progress *task.Progress, courseSectionID *string, events []interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseSectionID != nil {
		params["course_section_id"] = *courseSectionID
	}
	if events != nil && len(events) > 0 {
		params["events"] = events
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CollaborationsListCollaborations API call: A paginated list of collaborations the current user has access to in the
// context of the course provided in the url. NOTE: this only returns ExternalToolCollaboration type collaborations.
// curl https://<canvas>/api/v1/courses/1/collaborations/
func (c *Canvas) CollaborationsListCollaborations(progress *task.Progress) ([]Collaboration, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Collaboration{}
	}
	var res []Collaboration
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Collaboration)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CollaborationsListMembersOfACollaboration API call: A paginated list of the collaborators of a given collaboration
func (c *Canvas) CollaborationsListMembersOfACollaboration(progress *task.Progress, include *CollaborationsListMembersOfACollaborationInclude) ([]Collaborator, error) {
	endpoint := fmt.Sprintf("courses/1/collaborations/1/members")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]Collaborator{}
	}
	var res []Collaborator
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Collaborator)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CollaborationsListPotentialMembers API call: A paginated list of the users who can potentially be added to a
// collaboration in the given context. For courses, this consists of all enrolled users.  For groups, it is comprised of
// the group members plus the admins of the course containing the group.
func (c *Canvas) CollaborationsListPotentialMembers(progress *task.Progress) ([]User, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CommMessagesListOfCommMessagesForAUser API call: Retrieve a paginated list of messages sent to a user.
func (c *Canvas) CommMessagesListOfCommMessagesForAUser(progress *task.Progress, userID *string, startTime *time.Time, endTime *time.Time) ([]CommMessage, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if userID != nil {
		params["user_id"] = *userID
	}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &[]CommMessage{}
	}
	var res []CommMessage
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CommMessage)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CommunicationChannelsListUserCommunicationChannels API call: Returns a paginated list of communication channels for
// the specified user, sorted by position.
func (c *Canvas) CommunicationChannelsListUserCommunicationChannels(progress *task.Progress) ([]CommunicationChannel, error) {
	endpoint := fmt.Sprintf("users/12345/communication_channels")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]CommunicationChannel{}
	}
	var res []CommunicationChannel
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CommunicationChannel)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CommunicationChannelsCreateACommunicationChannel API call: Creates a new communication channel for the specified
// user.
func (c *Canvas) CommunicationChannelsCreateACommunicationChannel(progress *task.Progress, communicationChannel *string, skipConfirmation *bool) (*CommunicationChannel, error) {
	endpoint := fmt.Sprintf("users/1/communication_channels")
	params := map[string]interface{}{}
	if communicationChannel != nil {
		params["communication_channel"] = *communicationChannel
	}
	if skipConfirmation != nil {
		params["skip_confirmation"] = *skipConfirmation
	}
	responseCtor := func() interface{} {
		return &CommunicationChannel{}
	}
	var res *CommunicationChannel
	callback := func(obj interface{}) error {
		res = obj.(*CommunicationChannel)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CommunicationChannelsDeleteACommunicationChannel API call: Delete an existing communication channel.
func (c *Canvas) CommunicationChannelsDeleteACommunicationChannel(progress *task.Progress) (*CommunicationChannel, error) {
	endpoint := fmt.Sprintf("users/5/communication_channels/3")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &CommunicationChannel{}
	}
	var res *CommunicationChannel
	callback := func(obj interface{}) error {
		res = obj.(*CommunicationChannel)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CommunicationChannelsDeleteAPushNotificationEndpoint API call
func (c *Canvas) CommunicationChannelsDeleteAPushNotificationEndpoint(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/self/communication_channels/push")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConferencesListConferences API call: Retrieve the paginated list of conferences for this context This API returns a
// JSON object containing the list of conferences, the key for the list of conferences is "conferences"
func (c *Canvas) ConferencesListConferences(progress *task.Progress, courseID string) ([]Conference, error) {
	endpoint := fmt.Sprintf("courses/%s/conferences", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Conference{}
	}
	var res []Conference
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Conference)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConferencesListConferencesForTheCurrentUser API call: Retrieve the paginated list of conferences for all courses and
// groups the current user belongs to This API returns a JSON object containing the list of conferences. The key for the
// list of conferences is "conferences".
func (c *Canvas) ConferencesListConferencesForTheCurrentUser(progress *task.Progress, state *string) ([]Conference, error) {
	endpoint := fmt.Sprintf("conferences")
	params := map[string]interface{}{}
	if state != nil {
		params["state"] = *state
	}
	responseCtor := func() interface{} {
		return &[]Conference{}
	}
	var res []Conference
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Conference)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentExportsListContentExports API call: A paginated list of the past and pending content export jobs for a course,
// group, or user. Exports are returned newest first.
func (c *Canvas) ContentExportsListContentExports(progress *task.Progress) ([]ContentExport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]ContentExport{}
	}
	var res []ContentExport
	callback := func(obj interface{}) error {
		arr := *obj.(*[]ContentExport)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentExportsShowContentExport API call: Get information about a single content export.
func (c *Canvas) ContentExportsShowContentExport(progress *task.Progress) (*ContentExport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &ContentExport{}
	}
	var res *ContentExport
	callback := func(obj interface{}) error {
		res = obj.(*ContentExport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentExportsExportContent API call: Begin a content export job for a course, group, or user. You can use the
// {api:ProgressController#show Progress API} to track the progress of the export. The migration's progress is linked to
// with the _progress_url_ value. When the export completes, use the {api:ContentExportsApiController#show Show content
// export} endpoint to retrieve a download URL for the exported content.
func (c *Canvas) ContentExportsExportContent(progress *task.Progress, exportType *ContentExportsExportContentExportType, skipNotifications *bool, selectField *ContentExportsExportContentSelect) (*ContentExport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if exportType != nil {
		params["export_type"] = *exportType
	}
	if skipNotifications != nil {
		params["skip_notifications"] = *skipNotifications
	}
	if selectField != nil {
		params["select"] = *selectField
	}
	responseCtor := func() interface{} {
		return &ContentExport{}
	}
	var res *ContentExport
	callback := func(obj interface{}) error {
		res = obj.(*ContentExport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesGetCourseCopyStatus API call: DEPRECATED: Please use the {api:ContentMigrationsController#create Content
// Migrations API} Retrieve the status of a course copy
func (c *Canvas) CoursesGetCourseCopyStatus(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesCopyCourseContent API call: DEPRECATED: Please use the {api:ContentMigrationsController#create Content
// Migrations API} Copies content from one course into another. The default is to copy all course content. You can
// control specific types to copy by using either the 'except' option or the 'only' option.
func (c *Canvas) CoursesCopyCourseContent(progress *task.Progress, sourceCourse *string, except *CoursesCopyCourseContentExcept, only *CoursesCopyCourseContentOnly) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if sourceCourse != nil {
		params["source_course"] = *sourceCourse
	}
	if except != nil {
		params["except"] = *except
	}
	if only != nil {
		params["only"] = *only
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentMigrationsListContentMigrations API call: Returns paginated content migrations
func (c *Canvas) ContentMigrationsListContentMigrations(progress *task.Progress, courseID string) ([]ContentMigration, error) {
	endpoint := fmt.Sprintf("courses/%s/content_migrations", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]ContentMigration{}
	}
	var res []ContentMigration
	callback := func(obj interface{}) error {
		arr := *obj.(*[]ContentMigration)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentMigrationsGetAContentMigration API call: Returns data on an individual content migration
func (c *Canvas) ContentMigrationsGetAContentMigration(progress *task.Progress, courseID string, id string) (*ContentMigration, error) {
	endpoint := fmt.Sprintf("courses/%s/content_migrations/%s", courseID, id)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &ContentMigration{}
	}
	var res *ContentMigration
	callback := func(obj interface{}) error {
		res = obj.(*ContentMigration)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentMigrationsCreateAContentMigration API call: Create a content migration. If the migration requires a file to be
// uploaded the actual processing of the file will start once the file upload process is completed. File uploading works
// as described in the {file:file_uploads.html File Upload Documentation} except that the values are set on a
// *pre_attachment* sub-hash. For migrations that don't require a file to be uploaded, like course copy, the processing
// will begin as soon as the migration is created. You can use the {api:ProgressController#show Progress API} to track
// the progress of the migration. The migration's progress is linked to with the _progress_url_ value. The two general
// workflows are: If no file upload is needed: 1. POST to create 2. Use the {api:ProgressController#show Progress}
// specified in _progress_url_ to monitor progress For file uploading: 1. POST to create with file info in
// *pre_attachment* 2. Do {file:file_uploads.html file upload processing} using the data in the *pre_attachment* data 3.
// {api:ContentMigrationsController#show GET} the ContentMigration 4. Use the {api:ProgressController#show Progress}
// specified in _progress_url_ to monitor progress
func (c *Canvas) ContentMigrationsCreateAContentMigration(progress *task.Progress, migrationType *string, preAttachment *string, settings *interface{}, dateShiftOptions *bool, selectiveImport *bool, selectField *ContentMigrationsCreateAContentMigrationSelect, courseID string) (*ContentMigration, error) {
	endpoint := fmt.Sprintf("courses/%s/content_migrations", courseID)
	params := map[string]interface{}{}
	if migrationType != nil {
		params["migration_type"] = *migrationType
	}
	if preAttachment != nil {
		params["pre_attachment"] = *preAttachment
	}
	if settings != nil {
		params["settings"] = *settings
	}
	if dateShiftOptions != nil {
		params["date_shift_options"] = *dateShiftOptions
	}
	if selectiveImport != nil {
		params["selective_import"] = *selectiveImport
	}
	if selectField != nil {
		params["select"] = *selectField
	}
	responseCtor := func() interface{} {
		return &ContentMigration{}
	}
	var res *ContentMigration
	callback := func(obj interface{}) error {
		res = obj.(*ContentMigration)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentMigrationsUpdateAContentMigration API call: Update a content migration. Takes same arguments as
// {api:ContentMigrationsController#create create} except that you can't change the migration type. However, changing
// most settings after the migration process has started will not do anything. Generally updating the content migration
// will be used when there is a file upload problem, or when importing content selectively. If the first upload has a
// problem you can supply new _pre_attachment_ values to start the process again.
func (c *Canvas) ContentMigrationsUpdateAContentMigration(progress *task.Progress) (*ContentMigration, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &ContentMigration{}
	}
	var res *ContentMigration
	callback := func(obj interface{}) error {
		res = obj.(*ContentMigration)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentMigrationsListMigrationSystems API call: Lists the currently available migration types. These values may
// change.
func (c *Canvas) ContentMigrationsListMigrationSystems(progress *task.Progress) ([]Migrator, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Migrator{}
	}
	var res []Migrator
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Migrator)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentMigrationsListItemsForSelectiveImport API call: Enumerates the content available for selective import in a
// tree structure. Each node provides a +property+ copy argument that can be supplied to the
// {api:ContentMigrationsController#update Update endpoint} to selectively copy the content associated with that tree
// node and its children. Each node may also provide a +sub_items_url+ or an array of +sub_items+ which you can use to
// obtain copy parameters for a subset of the resources in a given node. If no +type+ is sent you will get a list of the
// top-level sections in the content. It will look something like this: [{ "type": "course_settings", "property":
// "copy[all_course_settings]", "title": "Course Settings" },
func (c *Canvas) ContentMigrationsListItemsForSelectiveImport(progress *task.Progress, typeName *ContentMigrationsListItemsForSelectiveImportType) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if typeName != nil {
		params["type"] = *typeName
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSharesCreateAContentShare API call: Share content directly between two or more users
func (c *Canvas) ContentSharesCreateAContentShare(progress *task.Progress, receiverIds []interface{}, contentType *ContentSharesCreateAContentShareContentType, contentID *int) (*ContentShare, error) {
	endpoint := fmt.Sprintf("users/self/content_shares")
	params := map[string]interface{}{}
	if receiverIds != nil && len(receiverIds) > 0 {
		params["receiver_ids"] = receiverIds
	}
	if contentType != nil {
		params["content_type"] = *contentType
	}
	if contentID != nil {
		params["content_id"] = *contentID
	}
	responseCtor := func() interface{} {
		return &ContentShare{}
	}
	var res *ContentShare
	callback := func(obj interface{}) error {
		res = obj.(*ContentShare)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSharesListContentShares API call: Return a paginated list of content shares a user has sent or received. Use
// +self+ as the user_id to retrieve your own content shares. Only linked observers and administrators may view other
// users' content shares.
func (c *Canvas) ContentSharesListContentShares(progress *task.Progress) ([]ContentShare, error) {
	endpoint := fmt.Sprintf("users/self/content_shares/received")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]ContentShare{}
	}
	var res []ContentShare
	callback := func(obj interface{}) error {
		arr := *obj.(*[]ContentShare)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSharesGetUnreadSharesCount API call: Return the number of content shares a user has received that have not yet
// been read. Use +self+ as the user_id to retrieve your own content shares. Only linked observers and administrators
// may view other users' content shares.
func (c *Canvas) ContentSharesGetUnreadSharesCount(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/self/content_shares/unread_count")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSharesGetContentShare API call: Return information about a single content share. You may use +self+ as the
// user_id to retrieve your own content share.
func (c *Canvas) ContentSharesGetContentShare(progress *task.Progress) (*ContentShare, error) {
	endpoint := fmt.Sprintf("users/self/content_shares/123")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &ContentShare{}
	}
	var res *ContentShare
	callback := func(obj interface{}) error {
		res = obj.(*ContentShare)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSharesRemoveContentShare API call: Remove a content share from your list. Use +self+ as the user_id. Note that
// this endpoint does not delete other users' copies of the content share.
func (c *Canvas) ContentSharesRemoveContentShare(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/self/content_shares/123")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSharesAddUsersToContentShare API call: Send a previously created content share to additional users
func (c *Canvas) ContentSharesAddUsersToContentShare(progress *task.Progress, receiverIds []interface{}) (*ContentShare, error) {
	endpoint := fmt.Sprintf("users/self/content_shares/123/add_users")
	params := map[string]interface{}{}
	if receiverIds != nil && len(receiverIds) > 0 {
		params["receiver_ids"] = receiverIds
	}
	responseCtor := func() interface{} {
		return &ContentShare{}
	}
	var res *ContentShare
	callback := func(obj interface{}) error {
		res = obj.(*ContentShare)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSharesUpdateAContentShare API call: Mark a content share read or unread
func (c *Canvas) ContentSharesUpdateAContentShare(progress *task.Progress, readState *ContentSharesUpdateAContentShareReadState) (*ContentShare, error) {
	endpoint := fmt.Sprintf("users/self/content_shares/123")
	params := map[string]interface{}{}
	if readState != nil {
		params["read_state"] = *readState
	}
	responseCtor := func() interface{} {
		return &ContentShare{}
	}
	var res *ContentShare
	callback := func(obj interface{}) error {
		res = obj.(*ContentShare)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesListModuleItems API call: A paginated list of the items in a module
func (c *Canvas) ModulesListModuleItems(progress *task.Progress, include *ModulesListModuleItemsInclude, searchTerm *string, studentID *interface{}) ([]ModuleItem, error) {
	endpoint := fmt.Sprintf("courses/222/modules/123/items")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if studentID != nil {
		params["student_id"] = *studentID
	}
	responseCtor := func() interface{} {
		return &[]ModuleItem{}
	}
	var res []ModuleItem
	callback := func(obj interface{}) error {
		arr := *obj.(*[]ModuleItem)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesShowModuleItem API call: Get information about a single module item
func (c *Canvas) ModulesShowModuleItem(progress *task.Progress, include *ModulesShowModuleItemInclude, studentID *interface{}) (*ModuleItem, error) {
	endpoint := fmt.Sprintf("courses/222/modules/123/items/768")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if studentID != nil {
		params["student_id"] = *studentID
	}
	responseCtor := func() interface{} {
		return &ModuleItem{}
	}
	var res *ModuleItem
	callback := func(obj interface{}) error {
		res = obj.(*ModuleItem)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesCreateAModuleItem API call: Create and return a new module item
func (c *Canvas) ModulesCreateAModuleItem(progress *task.Progress, moduleItem *string, courseID string, moduleID string) (*ModuleItem, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/%s/items", courseID, moduleID)
	params := map[string]interface{}{}
	if moduleItem != nil {
		params["module_item"] = *moduleItem
	}
	responseCtor := func() interface{} {
		return &ModuleItem{}
	}
	var res *ModuleItem
	callback := func(obj interface{}) error {
		res = obj.(*ModuleItem)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesUpdateAModuleItem API call: Update and return an existing module item
func (c *Canvas) ModulesUpdateAModuleItem(progress *task.Progress, moduleItem *string, courseID string, moduleID string, itemID string) (*ModuleItem, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/%s/items/%s", courseID, moduleID, itemID)
	params := map[string]interface{}{}
	if moduleItem != nil {
		params["module_item"] = *moduleItem
	}
	responseCtor := func() interface{} {
		return &ModuleItem{}
	}
	var res *ModuleItem
	callback := func(obj interface{}) error {
		res = obj.(*ModuleItem)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesSelectAMasteryPath API call: Select a mastery path when module item includes several possible paths. Requires
// Mastery Paths feature to be enabled.  Returns a compound document with the assignments included in the given path and
// any module items related to those assignments
func (c *Canvas) ModulesSelectAMasteryPath(progress *task.Progress, assignmentSetID *interface{}, studentID *interface{}, courseID string, moduleID string, itemID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/%s/items/%s/select_master_path", courseID, moduleID, itemID)
	params := map[string]interface{}{}
	if assignmentSetID != nil {
		params["assignment_set_id"] = *assignmentSetID
	}
	if studentID != nil {
		params["student_id"] = *studentID
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesDeleteModuleItem API call: Delete a module item
func (c *Canvas) ModulesDeleteModuleItem(progress *task.Progress, courseID string, moduleID string, itemID string) (*ModuleItem, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/%s/items/%s", courseID, moduleID, itemID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &ModuleItem{}
	}
	var res *ModuleItem
	callback := func(obj interface{}) error {
		res = obj.(*ModuleItem)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesMarkModuleItemAsDoneNotDone API call: Mark a module item as done/not done. Use HTTP method PUT to mark as
// done, and DELETE to mark as not done.
func (c *Canvas) ModulesMarkModuleItemAsDoneNotDone(progress *task.Progress, courseID string, moduleID string, itemID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/%s/items/%s/done", courseID, moduleID, itemID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesGetModuleItemSequence API call: Given an asset in a course, find the ModuleItem it belongs to, the previous
// and next Module Items in the course sequence, and also any applicable mastery path rules
func (c *Canvas) ModulesGetModuleItemSequence(progress *task.Progress, assetType *ModulesGetModuleItemSequenceAssetType, assetID *int, courseID string) (*ModuleItemSequence, error) {
	endpoint := fmt.Sprintf("courses/%s/module_item_sequence", courseID)
	params := map[string]interface{}{}
	if assetType != nil {
		params["asset_type"] = *assetType
	}
	if assetID != nil {
		params["asset_id"] = *assetID
	}
	responseCtor := func() interface{} {
		return &ModuleItemSequence{}
	}
	var res *ModuleItemSequence
	callback := func(obj interface{}) error {
		res = obj.(*ModuleItemSequence)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesMarkModuleItemRead API call: Fulfills "must view" requirement for a module item. It is generally not necessary
// to do this explicitly, but it is provided for applications that need to access external content directly (bypassing
// the html_url redirect that normally allows Canvas to fulfill "must view" requirements). This endpoint cannot be used
// to complete requirements on locked or unpublished module items.
func (c *Canvas) ModulesMarkModuleItemRead(progress *task.Progress, courseID string, moduleID string, itemID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/%s/items/%s/mark_read", courseID, moduleID, itemID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesDuplicateModuleItem API call: Makes a copy of an assignment, discussion or wiki page module item, within the
// same module. It also creates a duplicate copy of the assignment, discussion, or wiki page.
func (c *Canvas) ModulesDuplicateModuleItem(progress *task.Progress, courseID string, itemID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/items/%s/duplicate", courseID, itemID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesListModules API call: A paginated list of the modules in a course
func (c *Canvas) ModulesListModules(progress *task.Progress, include *ModulesListModulesInclude, searchTerm *string, studentID *interface{}) ([]Module, error) {
	endpoint := fmt.Sprintf("courses/222/modules")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if studentID != nil {
		params["student_id"] = *studentID
	}
	responseCtor := func() interface{} {
		return &[]Module{}
	}
	var res []Module
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Module)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesShowModule API call: Get information about a single module
func (c *Canvas) ModulesShowModule(progress *task.Progress, include *ModulesShowModuleInclude, studentID *interface{}) (*Module, error) {
	endpoint := fmt.Sprintf("courses/222/modules/123")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if studentID != nil {
		params["student_id"] = *studentID
	}
	responseCtor := func() interface{} {
		return &Module{}
	}
	var res *Module
	callback := func(obj interface{}) error {
		res = obj.(*Module)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesCreateAModule API call: Create and return a new module
func (c *Canvas) ModulesCreateAModule(progress *task.Progress, module *string, courseID string) (*Module, error) {
	endpoint := fmt.Sprintf("courses/%s/modules", courseID)
	params := map[string]interface{}{}
	if module != nil {
		params["module"] = *module
	}
	responseCtor := func() interface{} {
		return &Module{}
	}
	var res *Module
	callback := func(obj interface{}) error {
		res = obj.(*Module)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesUpdateAModule API call: Update and return an existing module
func (c *Canvas) ModulesUpdateAModule(progress *task.Progress, module *string, courseID string, moduleID string) (*Module, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/%s", courseID, moduleID)
	params := map[string]interface{}{}
	if module != nil {
		params["module"] = *module
	}
	responseCtor := func() interface{} {
		return &Module{}
	}
	var res *Module
	callback := func(obj interface{}) error {
		res = obj.(*Module)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesDeleteModule API call: Delete a module
func (c *Canvas) ModulesDeleteModule(progress *task.Progress, courseID string, moduleID string) (*Module, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/%s", courseID, moduleID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Module{}
	}
	var res *Module
	callback := func(obj interface{}) error {
		res = obj.(*Module)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModulesReLockModuleProgressions API call: Resets module progressions to their default locked state and recalculates
// them based on the current requirements. Adding progression requirements to an active course will not lock students
// out of modules they have already unlocked unless this action is called.
func (c *Canvas) ModulesReLockModuleProgressions(progress *task.Progress, courseID string, moduleID string) (*Module, error) {
	endpoint := fmt.Sprintf("courses/%s/modules/%s/relock", courseID, moduleID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Module{}
	}
	var res *Module
	callback := func(obj interface{}) error {
		res = obj.(*Module)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsListConversations API call: Returns the paginated list of conversations for the current user, most
// recent ones first.
func (c *Canvas) ConversationsListConversations(progress *task.Progress, scope *ConversationsListConversationsScope, filter *ConversationsListConversationsFilter, filterMode *ConversationsListConversationsFilterMode, interleaveSubmissions *interface{}, includeAllConversationIds *interface{}, include *ConversationsListConversationsInclude) ([]Conversation, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if scope != nil {
		params["scope"] = *scope
	}
	if filter != nil {
		params["filter"] = *filter
	}
	if filterMode != nil {
		params["filter_mode"] = *filterMode
	}
	if interleaveSubmissions != nil {
		params["interleave_submissions"] = *interleaveSubmissions
	}
	if includeAllConversationIds != nil {
		params["include_all_conversation_ids"] = *includeAllConversationIds
	}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]Conversation{}
	}
	var res []Conversation
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Conversation)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsCreateAConversation API call: Create a new conversation with one or more recipients. If there is already
// an existing private conversation with the given recipients, it will be reused.
func (c *Canvas) ConversationsCreateAConversation(progress *task.Progress, recipients *string, subject *string, body *string, forceNew *bool, groupConversation *bool, attachmentIds *string, mediaCommentID *string, mediaCommentType *ConversationsCreateAConversationMediaCommentType, userNote *bool, mode *ConversationsCreateAConversationMode, scope *ConversationsCreateAConversationScope, filter *ConversationsCreateAConversationFilter, filterMode *ConversationsCreateAConversationFilterMode, contextCode *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if recipients != nil {
		params["recipients"] = *recipients
	}
	if subject != nil {
		params["subject"] = *subject
	}
	if body != nil {
		params["body"] = *body
	}
	if forceNew != nil {
		params["force_new"] = *forceNew
	}
	if groupConversation != nil {
		params["group_conversation"] = *groupConversation
	}
	if attachmentIds != nil {
		params["attachment_ids"] = *attachmentIds
	}
	if mediaCommentID != nil {
		params["media_comment_id"] = *mediaCommentID
	}
	if mediaCommentType != nil {
		params["media_comment_type"] = *mediaCommentType
	}
	if userNote != nil {
		params["user_note"] = *userNote
	}
	if mode != nil {
		params["mode"] = *mode
	}
	if scope != nil {
		params["scope"] = *scope
	}
	if filter != nil {
		params["filter"] = *filter
	}
	if filterMode != nil {
		params["filter_mode"] = *filterMode
	}
	if contextCode != nil {
		params["context_code"] = *contextCode
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsGetRunningBatches API call: Returns any currently running conversation batches for the current user.
// Conversation batches are created when a bulk private message is sent asynchronously (see the mode argument to the
// {api:ConversationsController#create create API action}).
func (c *Canvas) ConversationsGetRunningBatches(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsGetASingleConversation API call: Returns information for a single conversation for the current user.
// Response includes all fields that are present in the list/index action as well as messages and extended participant
// information.
func (c *Canvas) ConversationsGetASingleConversation(progress *task.Progress, interleaveSubmissions *interface{}, scope *ConversationsGetASingleConversationScope, filter *ConversationsGetASingleConversationFilter, filterMode *ConversationsGetASingleConversationFilterMode, autoMarkAsRead *interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if interleaveSubmissions != nil {
		params["interleave_submissions"] = *interleaveSubmissions
	}
	if scope != nil {
		params["scope"] = *scope
	}
	if filter != nil {
		params["filter"] = *filter
	}
	if filterMode != nil {
		params["filter_mode"] = *filterMode
	}
	if autoMarkAsRead != nil {
		params["auto_mark_as_read"] = *autoMarkAsRead
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsEditAConversation API call: Updates attributes for a single conversation.
func (c *Canvas) ConversationsEditAConversation(progress *task.Progress, conversation *ConversationsEditAConversationConversation, scope *ConversationsEditAConversationScope, filter *ConversationsEditAConversationFilter, filterMode *ConversationsEditAConversationFilterMode) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if conversation != nil {
		params["conversation"] = *conversation
	}
	if scope != nil {
		params["scope"] = *scope
	}
	if filter != nil {
		params["filter"] = *filter
	}
	if filterMode != nil {
		params["filter_mode"] = *filterMode
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsMarkAllAsRead API call: Mark all conversations as read.
func (c *Canvas) ConversationsMarkAllAsRead(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsDeleteAConversation API call: Delete this conversation and its messages. Note that this only deletes
// this user's view of the conversation. Response includes same fields as UPDATE action
func (c *Canvas) ConversationsDeleteAConversation(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsAddRecipients API call: Add recipients to an existing group conversation. Response is similar to the
// GET/show action, except that only includes the latest message (e.g. "joe was added to the conversation by bob")
func (c *Canvas) ConversationsAddRecipients(progress *task.Progress, recipients *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if recipients != nil {
		params["recipients"] = *recipients
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsAddAMessage API call: Add a message to an existing conversation. Response is similar to the GET/show
// action, except that only includes the latest message (i.e. what we just sent)
func (c *Canvas) ConversationsAddAMessage(progress *task.Progress, body *string, attachmentIds *string, mediaCommentID *string, mediaCommentType *ConversationsAddAMessageMediaCommentType, recipients *string, includedMessages *string, userNote *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if body != nil {
		params["body"] = *body
	}
	if attachmentIds != nil {
		params["attachment_ids"] = *attachmentIds
	}
	if mediaCommentID != nil {
		params["media_comment_id"] = *mediaCommentID
	}
	if mediaCommentType != nil {
		params["media_comment_type"] = *mediaCommentType
	}
	if recipients != nil {
		params["recipients"] = *recipients
	}
	if includedMessages != nil {
		params["included_messages"] = *includedMessages
	}
	if userNote != nil {
		params["user_note"] = *userNote
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsDeleteAMessage API call: Delete messages from this conversation. Note that this only affects this user's
// view of the conversation. If all messages are deleted, the conversation will be as well (equivalent to DELETE)
func (c *Canvas) ConversationsDeleteAMessage(progress *task.Progress, remove *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if remove != nil {
		params["remove"] = *remove
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsBatchUpdateConversations API call: Perform a change on a set of conversations. Operates asynchronously;
// use the {api:ProgressController#show progress endpoint} to query the status of an operation.
func (c *Canvas) ConversationsBatchUpdateConversations(progress *task.Progress, conversationIds *string, event *ConversationsBatchUpdateConversationsEvent) (*Progress, error) {
	endpoint := fmt.Sprintf("conversations")
	params := map[string]interface{}{}
	if conversationIds != nil {
		params["conversation_ids"] = *conversationIds
	}
	if event != nil {
		params["event"] = *event
	}
	responseCtor := func() interface{} {
		return &Progress{}
	}
	var res *Progress
	callback := func(obj interface{}) error {
		res = obj.(*Progress)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsFindRecipients API call: Deprecated, see the {api:SearchController#recipients Find recipients endpoint}
// in the Search API
func (c *Canvas) ConversationsFindRecipients(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ConversationsUnreadCount API call: Get the number of unread conversations for the current user
func (c *Canvas) ConversationsUnreadCount(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CourseAuditLogQueryByCourse API call: List course change events for a given course.
func (c *Canvas) CourseAuditLogQueryByCourse(progress *task.Progress, startTime *time.Time, endTime *time.Time) ([]CourseEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &[]CourseEvent{}
	}
	var res []CourseEvent
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CourseEvent)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CourseAuditLogQueryByAccount API call: List course change events for a given account.
func (c *Canvas) CourseAuditLogQueryByAccount(progress *task.Progress, startTime *time.Time, endTime *time.Time) ([]CourseEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &[]CourseEvent{}
	}
	var res []CourseEvent
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CourseEvent)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersListCourseNicknames API call: Returns all course nicknames you have set.
func (c *Canvas) UsersListCourseNicknames(progress *task.Progress) ([]CourseNickname, error) {
	endpoint := fmt.Sprintf("users/self/course_nicknames")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]CourseNickname{}
	}
	var res []CourseNickname
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CourseNickname)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersGetCourseNickname API call: Returns the nickname for a specific course.
func (c *Canvas) UsersGetCourseNickname(progress *task.Progress, courseID string) (*CourseNickname, error) {
	endpoint := fmt.Sprintf("users/self/course_nicknames/%s", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &CourseNickname{}
	}
	var res *CourseNickname
	callback := func(obj interface{}) error {
		res = obj.(*CourseNickname)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersSetCourseNickname API call: Set a nickname for the given course. This will replace the course's name in output
// of API calls you make subsequently, as well as in selected places in the Canvas web user interface.
func (c *Canvas) UsersSetCourseNickname(progress *task.Progress, nickname *string, courseID string) (*CourseNickname, error) {
	endpoint := fmt.Sprintf("users/self/course_nicknames/%s", courseID)
	params := map[string]interface{}{}
	if nickname != nil {
		params["nickname"] = *nickname
	}
	responseCtor := func() interface{} {
		return &CourseNickname{}
	}
	var res *CourseNickname
	callback := func(obj interface{}) error {
		res = obj.(*CourseNickname)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersRemoveCourseNickname API call: Remove the nickname for the given course. Subsequent course API calls will return
// the actual name for the course.
func (c *Canvas) UsersRemoveCourseNickname(progress *task.Progress, courseID string) (*CourseNickname, error) {
	endpoint := fmt.Sprintf("users/self/course_nicknames/%s", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &CourseNickname{}
	}
	var res *CourseNickname
	callback := func(obj interface{}) error {
		res = obj.(*CourseNickname)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersClearCourseNicknames API call: Remove all stored course nicknames.
func (c *Canvas) UsersClearCourseNicknames(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/self/course_nicknames")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesListYourCourses API call: Returns the paginated list of active courses for the current user.
func (c *Canvas) CoursesListYourCourses(progress *task.Progress, enrollmentType *CoursesListYourCoursesEnrollmentType, enrollmentRole *interface{}, enrollmentRoleID *int, enrollmentState *CoursesListYourCoursesEnrollmentState, excludeBlueprintCourses *bool, include *CoursesListYourCoursesInclude, state *CoursesListYourCoursesState) ([]Course, error) {
	endpoint := fmt.Sprintf("courses")
	params := map[string]interface{}{}
	if enrollmentType != nil {
		params["enrollment_type"] = *enrollmentType
	}
	if enrollmentRole != nil {
		params["enrollment_role"] = *enrollmentRole
	}
	if enrollmentRoleID != nil {
		params["enrollment_role_id"] = *enrollmentRoleID
	}
	if enrollmentState != nil {
		params["enrollment_state"] = *enrollmentState
	}
	if excludeBlueprintCourses != nil {
		params["exclude_blueprint_courses"] = *excludeBlueprintCourses
	}
	if include != nil {
		params["include"] = *include
	}
	if state != nil {
		params["state"] = *state
	}
	responseCtor := func() interface{} {
		return &[]Course{}
	}
	var res []Course
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Course)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesListCoursesForAUser API call: Returns a paginated list of active courses for this user. To view the course
// list for a user other than yourself, you must be either an observer of that user or an administrator.
func (c *Canvas) CoursesListCoursesForAUser(progress *task.Progress, include *CoursesListCoursesForAUserInclude, state *CoursesListCoursesForAUserState, enrollmentState *CoursesListCoursesForAUserEnrollmentState) ([]Course, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if state != nil {
		params["state"] = *state
	}
	if enrollmentState != nil {
		params["enrollment_state"] = *enrollmentState
	}
	responseCtor := func() interface{} {
		return &[]Course{}
	}
	var res []Course
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Course)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesCreateANewCourse API call: Create a new course
func (c *Canvas) CoursesCreateANewCourse(progress *task.Progress, course *string, offer *bool, enrollMe *bool, enableSisReactivation *bool) (*Course, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if course != nil {
		params["course"] = *course
	}
	if offer != nil {
		params["offer"] = *offer
	}
	if enrollMe != nil {
		params["enroll_me"] = *enrollMe
	}
	if enableSisReactivation != nil {
		params["enable_sis_reactivation"] = *enableSisReactivation
	}
	responseCtor := func() interface{} {
		return &Course{}
	}
	var res *Course
	callback := func(obj interface{}) error {
		res = obj.(*Course)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesUploadAFile API call: Upload a file to the course. This API endpoint is the first step in uploading a file to
// a course. See the {file:file_uploads.html File Upload Documentation} for details on the file upload workflow. Only
// those with the "Manage Files" permission on a course can upload files to the course. By default, this is Teachers,
// TAs and Designers.
func (c *Canvas) CoursesUploadAFile(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesListStudents API call: Returns the paginated list of students enrolled in this course. DEPRECATED: Please use
// the {api:CoursesController#users course users} endpoint and pass "student" as the enrollment_type.
func (c *Canvas) CoursesListStudents(progress *task.Progress) ([]User, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesListUsersInCourse API call: Returns the paginated list of users in this course. And optionally the user's
// enrollments in the course.
func (c *Canvas) CoursesListUsersInCourse(progress *task.Progress, searchTerm *string, sort *CoursesListUsersInCourseSort, enrollmentType *CoursesListUsersInCourseEnrollmentType, enrollmentRole *interface{}, enrollmentRoleID *int, include *CoursesListUsersInCourseInclude, userID *string, userIds *int, enrollmentState *CoursesListUsersInCourseEnrollmentState) ([]User, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if sort != nil {
		params["sort"] = *sort
	}
	if enrollmentType != nil {
		params["enrollment_type"] = *enrollmentType
	}
	if enrollmentRole != nil {
		params["enrollment_role"] = *enrollmentRole
	}
	if enrollmentRoleID != nil {
		params["enrollment_role_id"] = *enrollmentRoleID
	}
	if include != nil {
		params["include"] = *include
	}
	if userID != nil {
		params["user_id"] = *userID
	}
	if userIds != nil {
		params["user_ids"] = *userIds
	}
	if enrollmentState != nil {
		params["enrollment_state"] = *enrollmentState
	}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesListRecentlyLoggedInStudents API call: Returns the paginated list of users in this course, ordered by how
// recently they have logged in. The records include the 'last_login' field which contains a timestamp of the last time
// that user logged into canvas.  The querying user must have the 'View usage reports' permission.
func (c *Canvas) CoursesListRecentlyLoggedInStudents(progress *task.Progress, courseID string) ([]User, error) {
	endpoint := fmt.Sprintf("courses/%s/recent_users", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesGetSingleUser API call: Return information on a single user. Accepts the same include[] parameters as the
// :users: action, and returns a single user with the same fields as that action.
func (c *Canvas) CoursesGetSingleUser(progress *task.Progress) (*User, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesSearchForContentShareUsers API call: Returns a paginated list of users you can share content with.  Requires
// the content share feature and the user must have the manage content permission for the course.
func (c *Canvas) CoursesSearchForContentShareUsers(progress *task.Progress, searchTerm *string, courseID string) ([]User, error) {
	endpoint := fmt.Sprintf("courses/%s/content_share_users", courseID)
	params := map[string]interface{}{}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesPreviewProcessedHTML API call: Preview html content processed for this course
func (c *Canvas) CoursesPreviewProcessedHTML(progress *task.Progress, html *interface{}, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/preview_html", courseID)
	params := map[string]interface{}{}
	if html != nil {
		params["html"] = *html
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesCourseActivityStream API call: Returns the current user's course-specific activity stream, paginated. For full
// documentation, see the API documentation for the user activity stream, in the user api.
func (c *Canvas) CoursesCourseActivityStream(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesCourseActivityStreamSummary API call: Returns a summary of the current user's course-specific activity stream.
// For full documentation, see the API documentation for the user activity stream summary, in the user api.
func (c *Canvas) CoursesCourseActivityStreamSummary(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesCourseTODOItems API call: Returns the current user's course-specific todo items. For full documentation, see
// the API documentation for the user todo items, in the user api.
func (c *Canvas) CoursesCourseTODOItems(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesDeleteConcludeACourse API call: Delete or conclude an existing course
func (c *Canvas) CoursesDeleteConcludeACourse(progress *task.Progress, event *CoursesDeleteConcludeACourseEvent) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if event != nil {
		params["event"] = *event
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesGetCourseSettings API call: Returns some of a course's settings.
func (c *Canvas) CoursesGetCourseSettings(progress *task.Progress, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/settings", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesUpdateCourseSettings API call: Can update the following course settings:
func (c *Canvas) CoursesUpdateCourseSettings(progress *task.Progress, allowStudentDiscussionTopics *bool, allowStudentForumAttachments *bool, allowStudentDiscussionEditing *bool, allowStudentOrganizedGroups *bool, filterSpeedGraderByStudentGroup *bool, hideFinalGrades *bool, hideDistributionGraphs *bool, lockAllAnnouncements *bool, usageRightsRequired *bool, restrictStudentPastView *bool, restrictStudentFutureView *bool, showAnnouncementsOnHomePage *bool, homePageAnnouncementLimit *int, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/settings", courseID)
	params := map[string]interface{}{}
	if allowStudentDiscussionTopics != nil {
		params["allow_student_discussion_topics"] = *allowStudentDiscussionTopics
	}
	if allowStudentForumAttachments != nil {
		params["allow_student_forum_attachments"] = *allowStudentForumAttachments
	}
	if allowStudentDiscussionEditing != nil {
		params["allow_student_discussion_editing"] = *allowStudentDiscussionEditing
	}
	if allowStudentOrganizedGroups != nil {
		params["allow_student_organized_groups"] = *allowStudentOrganizedGroups
	}
	if filterSpeedGraderByStudentGroup != nil {
		params["filter_speed_grader_by_student_group"] = *filterSpeedGraderByStudentGroup
	}
	if hideFinalGrades != nil {
		params["hide_final_grades"] = *hideFinalGrades
	}
	if hideDistributionGraphs != nil {
		params["hide_distribution_graphs"] = *hideDistributionGraphs
	}
	if lockAllAnnouncements != nil {
		params["lock_all_announcements"] = *lockAllAnnouncements
	}
	if usageRightsRequired != nil {
		params["usage_rights_required"] = *usageRightsRequired
	}
	if restrictStudentPastView != nil {
		params["restrict_student_past_view"] = *restrictStudentPastView
	}
	if restrictStudentFutureView != nil {
		params["restrict_student_future_view"] = *restrictStudentFutureView
	}
	if showAnnouncementsOnHomePage != nil {
		params["show_announcements_on_home_page"] = *showAnnouncementsOnHomePage
	}
	if homePageAnnouncementLimit != nil {
		params["home_page_announcement_limit"] = *homePageAnnouncementLimit
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesReturnTestStudentForCourse API call: Returns information for a test student in this course. Creates a test
// student if one does not already exist for the course. The caller must have permission to access the course's student
// view.
func (c *Canvas) CoursesReturnTestStudentForCourse(progress *task.Progress, courseID string) (*User, error) {
	endpoint := fmt.Sprintf("courses/%s/student_view_student", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesGetASingleCourse API call: Return information on a single course. Accepts the same include[] parameters as the
// list action plus:
func (c *Canvas) CoursesGetASingleCourse(progress *task.Progress, include *CoursesGetASingleCourseInclude, teacherLimit *int) (*Course, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if teacherLimit != nil {
		params["teacher_limit"] = *teacherLimit
	}
	responseCtor := func() interface{} {
		return &Course{}
	}
	var res *Course
	callback := func(obj interface{}) error {
		res = obj.(*Course)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesUpdateACourse API call: Update an existing course. Arguments are the same as Courses#create, with a few
// exceptions (enroll_me). If a user has content management rights, but not full course editing rights, the only
// attribute editable through this endpoint will be "syllabus_body"
func (c *Canvas) CoursesUpdateACourse(progress *task.Progress, course *int, offer *bool, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s", courseID)
	params := map[string]interface{}{}
	if course != nil {
		params["course"] = *course
	}
	if offer != nil {
		params["offer"] = *offer
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesUpdateCourses API call: Update multiple courses in an account.  Operates asynchronously; use the
// {api:ProgressController#show progress endpoint} to query the status of an operation.
func (c *Canvas) CoursesUpdateCourses(progress *task.Progress, courseIds *interface{}, event *CoursesUpdateCoursesEvent, accountID string) (*Progress, error) {
	endpoint := fmt.Sprintf("accounts/%s/courses", accountID)
	params := map[string]interface{}{}
	if courseIds != nil {
		params["course_ids"] = *courseIds
	}
	if event != nil {
		params["event"] = *event
	}
	responseCtor := func() interface{} {
		return &Progress{}
	}
	var res *Progress
	callback := func(obj interface{}) error {
		res = obj.(*Progress)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesResetACourse API call: Deletes the current course, and creates a new equivalent course with no content, but
// all sections and users moved over.
func (c *Canvas) CoursesResetACourse(progress *task.Progress) (*Course, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Course{}
	}
	var res *Course
	callback := func(obj interface{}) error {
		res = obj.(*Course)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesGetEffectiveDueDates API call: For each assignment in the course, returns each assigned student's ID and their
// corresponding due date along with some grading period data. Returns a collection with keys representing assignment
// IDs and values as a collection containing keys representing student IDs and values representing the student's
// effective due_at, the grading_period_id of which the due_at falls in, and whether or not the grading period is closed
// (in_closed_grading_period)
func (c *Canvas) CoursesGetEffectiveDueDates(progress *task.Progress, assignmentIds *string, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/effective_due_dates", courseID)
	params := map[string]interface{}{}
	if assignmentIds != nil {
		params["assignment_ids"] = *assignmentIds
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CoursesPermissions API call: Returns permission information for the calling user in the given course. See also the
// {api:AccountsController#permissions Account} and {api:GroupsController#permissions Group} counterparts.
func (c *Canvas) CoursesPermissions(progress *task.Progress, permissions *string, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/permissions", courseID)
	params := map[string]interface{}{}
	if permissions != nil {
		params["permissions"] = *permissions
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSecurityPolicySettingsGetCurrentSettingsForAccountOrCourse API call: Update multiple modules in an account.
func (c *Canvas) ContentSecurityPolicySettingsGetCurrentSettingsForAccountOrCourse(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSetting API call: Either explicitly sets CSP to be on or
// off for courses and sub-accounts, or clear the explicit settings to default to those set by a parent account Note: If
// "inherited" and "settings_locked" are both true for this account or course, then the CSP setting cannot be modified.
func (c *Canvas) ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSetting(progress *task.Progress, status *ContentSecurityPolicySettingsEnableDisableOrClearExplicitCSPSettingStatus) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if status != nil {
		params["status"] = *status
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSecurityPolicySettingsLockOrUnlockCurrentCSPSettingsForSubAccountsAndCourses API call: Can only be set if CSP
// is explicitly enabled or disabled on this account (i.e. "inherited" is false).
func (c *Canvas) ContentSecurityPolicySettingsLockOrUnlockCurrentCSPSettingsForSubAccountsAndCourses(progress *task.Progress, settingsLocked *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if settingsLocked != nil {
		params["settings_locked"] = *settingsLocked
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSecurityPolicySettingsAddADomainToAccountWhitelist API call: Adds a domain to the whitelist for the current
// account. Note: this will not take effect unless CSP is explicitly enabled on this account.
func (c *Canvas) ContentSecurityPolicySettingsAddADomainToAccountWhitelist(progress *task.Progress, domain *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if domain != nil {
		params["domain"] = *domain
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSecurityPolicySettingsAddMultipleDomainsToAccountWhitelist API call: Adds multiple domains to the whitelist
// for the current account. Note: this will not take effect unless CSP is explicitly enabled on this account.
func (c *Canvas) ContentSecurityPolicySettingsAddMultipleDomainsToAccountWhitelist(progress *task.Progress, domains []interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if domains != nil && len(domains) > 0 {
		params["domains"] = domains
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSecurityPolicySettingsRetrieveReportedCSPViolationsForAccount API call: Must be called on a root account.
func (c *Canvas) ContentSecurityPolicySettingsRetrieveReportedCSPViolationsForAccount(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentSecurityPolicySettingsRemoveADomainFromAccountWhitelist API call: Removes a domain from the whitelist for the
// current account.
func (c *Canvas) ContentSecurityPolicySettingsRemoveADomainFromAccountWhitelist(progress *task.Progress, domain *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if domain != nil {
		params["domain"] = *domain
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersStoreCustomData API call: Store arbitrary user data as JSON. Arbitrary JSON data can be stored for a User. A
// typical scenario would be an external site/service that registers users in Canvas and wants to capture additional
// info about them.  The part of the URL that follows +/custom_data/+ defines the scope of the request, and it reflects
// the structure of the JSON data to be stored or retrieved. The value +self+ may be used for +user_id+ to store data
// associated with the calling user. In order to access another user's custom data, you must be an account administrator
// with permission to manage users. A namespace parameter, +ns+, is used to prevent custom_data collisions between
// different apps.  This parameter is required for all custom_data requests. A request with Content-Type
// multipart/form-data or Content-Type application/x-www-form-urlencoded can only be used to store strings. Example PUT
// with multipart/form-data data: curl 'https://<canvas>/api/v1/users/<user_id>/custom_data/telephone' \ -X PUT \ -F
// 'ns=com.my-organization.canvas-app' \ -F 'data=555-1234' \ -H 'Authorization: Bearer <token>' Response: !!!javascript
func (c *Canvas) UsersStoreCustomData(progress *task.Progress, ns *string, data map[string]interface{}, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/%s/custom_data/food_app", userID)
	params := map[string]interface{}{}
	if ns != nil {
		params["ns"] = *ns
	}
	if data != nil && len(data) > 0 {
		params["data"] = data
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersLoadCustomData API call: Load custom user data. Arbitrary JSON data can be stored for a User.  This API call
// retrieves that data for a (optional) given scope. See {api:UsersController#set_custom_data Store Custom Data} for
// details and examples. On success, this endpoint returns an object containing the data that was requested. Responds
// with status code 400 if the namespace parameter, +ns+, is missing or invalid, or if the specified scope does not
// contain any data.
func (c *Canvas) UsersLoadCustomData(progress *task.Progress, ns *string, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/%s/custom_data/food_app/favorites/dessert", userID)
	params := map[string]interface{}{}
	if ns != nil {
		params["ns"] = *ns
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersDeleteCustomData API call: Delete custom user data. Arbitrary JSON data can be stored for a User.  This API call
// deletes that data for a given scope.  Without a scope, all custom_data is deleted. See
// {api:UsersController#set_custom_data Store Custom Data} for details and examples of storage and retrieval. As an
// example, we'll store some data, then delete a subset of it. Example {api:UsersController#set_custom_data PUT} with
// valid JSON data: curl 'https://<canvas>/api/v1/users/<user_id>/custom_data' \ -X PUT \ -F
// 'ns=com.my-organization.canvas-app' \ -F 'data[fruit][apple]=so tasty' \ -F 'data[fruit][kiwi]=a bit sour' \ -F
// 'data[veggies][root][onion]=tear-jerking' \ -H 'Authorization: Bearer <token>' Response: !!!javascript
func (c *Canvas) UsersDeleteCustomData(progress *task.Progress, ns *string, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/%s/custom_data/fruit/kiwi", userID)
	params := map[string]interface{}{}
	if ns != nil {
		params["ns"] = *ns
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CustomGradebookColumnsListEntriesForAColumn API call: This does not list entries for students without associated
// data.
func (c *Canvas) CustomGradebookColumnsListEntriesForAColumn(progress *task.Progress, includeHidden *bool) ([]ColumnDatum, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if includeHidden != nil {
		params["include_hidden"] = *includeHidden
	}
	responseCtor := func() interface{} {
		return &[]ColumnDatum{}
	}
	var res []ColumnDatum
	callback := func(obj interface{}) error {
		arr := *obj.(*[]ColumnDatum)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CustomGradebookColumnsUpdateColumnData API call: Set the content of a custom column
func (c *Canvas) CustomGradebookColumnsUpdateColumnData(progress *task.Progress, columnData *string) (*ColumnDatum, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if columnData != nil {
		params["column_data"] = *columnData
	}
	responseCtor := func() interface{} {
		return &ColumnDatum{}
	}
	var res *ColumnDatum
	callback := func(obj interface{}) error {
		res = obj.(*ColumnDatum)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CustomGradebookColumnsBulkUpdateColumnData API call: Set the content of custom columns
func (c *Canvas) CustomGradebookColumnsBulkUpdateColumnData(progress *task.Progress, columnData []interface{}) (*Progress, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if columnData != nil && len(columnData) > 0 {
		params["column_data"] = columnData
	}
	responseCtor := func() interface{} {
		return &Progress{}
	}
	var res *Progress
	callback := func(obj interface{}) error {
		res = obj.(*Progress)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CustomGradebookColumnsListCustomGradebookColumns API call: A paginated list of all custom gradebook columns for a
// course
func (c *Canvas) CustomGradebookColumnsListCustomGradebookColumns(progress *task.Progress, includeHidden *bool) ([]CustomColumn, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if includeHidden != nil {
		params["include_hidden"] = *includeHidden
	}
	responseCtor := func() interface{} {
		return &[]CustomColumn{}
	}
	var res []CustomColumn
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CustomColumn)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CustomGradebookColumnsCreateACustomGradebookColumn API call: Create a custom gradebook column
func (c *Canvas) CustomGradebookColumnsCreateACustomGradebookColumn(progress *task.Progress, column *string) (*CustomColumn, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if column != nil {
		params["column"] = *column
	}
	responseCtor := func() interface{} {
		return &CustomColumn{}
	}
	var res *CustomColumn
	callback := func(obj interface{}) error {
		res = obj.(*CustomColumn)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CustomGradebookColumnsUpdateACustomGradebookColumn API call: Accepts the same parameters as custom gradebook column
// creation
func (c *Canvas) CustomGradebookColumnsUpdateACustomGradebookColumn(progress *task.Progress) (*CustomColumn, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &CustomColumn{}
	}
	var res *CustomColumn
	callback := func(obj interface{}) error {
		res = obj.(*CustomColumn)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CustomGradebookColumnsDeleteACustomGradebookColumn API call: Permanently deletes a custom column and its associated
// data
func (c *Canvas) CustomGradebookColumnsDeleteACustomGradebookColumn(progress *task.Progress) (*CustomColumn, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &CustomColumn{}
	}
	var res *CustomColumn
	callback := func(obj interface{}) error {
		res = obj.(*CustomColumn)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// CustomGradebookColumnsReorderCustomColumns API call: Puts the given columns in the specified order
func (c *Canvas) CustomGradebookColumnsReorderCustomColumns(progress *task.Progress, order *int) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if order != nil {
		params["order"] = *order
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DeveloperKeyAccountBindingsCreateADeveloperKeyAccountBinding API call: Create a new Developer Key Account Binding.
// The developer key specified in the request URL must be available in the requested account or the requeted account's
// account chain. If the binding already exists for the specified account/key combination it will be updated.
func (c *Canvas) DeveloperKeyAccountBindingsCreateADeveloperKeyAccountBinding(progress *task.Progress, workflowState *string) (*DeveloperKeyAccountBinding, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if workflowState != nil {
		params["workflow_state"] = *workflowState
	}
	responseCtor := func() interface{} {
		return &DeveloperKeyAccountBinding{}
	}
	var res *DeveloperKeyAccountBinding
	callback := func(obj interface{}) error {
		res = obj.(*DeveloperKeyAccountBinding)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DeveloperKeyAccountBindingsListDeveloperKeyAccountBinding API call: List all Developer Key Account Bindings in the
// requested account
func (c *Canvas) DeveloperKeyAccountBindingsListDeveloperKeyAccountBinding(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISIntegrationDisableAssignmentsCurrentlyEnabledForGradeExportToSIS API call: Disable all assignments flagged as
// "post_to_sis", with the option of making it specific to a grading period, in a course.
func (c *Canvas) SISIntegrationDisableAssignmentsCurrentlyEnabledForGradeExportToSIS(progress *task.Progress, courseID *interface{}, gradingPeriodID *interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	if gradingPeriodID != nil {
		params["grading_period_id"] = *gradingPeriodID
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsUpdateAnEntry API call: Update an existing discussion entry. The entry must have been created by the
// current user, or the current user must have admin rights to the discussion. If the edit is not allowed, a 401 will be
// returned.
func (c *Canvas) DiscussionTopicsUpdateAnEntry(progress *task.Progress, message *interface{}, courseID string, topicID string, entryID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/entries/%s", courseID, topicID, entryID)
	params := map[string]interface{}{}
	if message != nil {
		params["message"] = *message
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsDeleteAnEntry API call: Delete a discussion entry. The entry must have been created by the current
// user, or the current user must have admin rights to the discussion. If the delete is not allowed, a 401 will be
// returned. The discussion will be marked deleted, and the user_id and message will be cleared out.
func (c *Canvas) DiscussionTopicsDeleteAnEntry(progress *task.Progress, courseID string, topicID string, entryID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/entries/%s", courseID, topicID, entryID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsGetASingleTopic API call: Returns data on an individual discussion topic. See the List action for the
// response formatting.
func (c *Canvas) DiscussionTopicsGetASingleTopic(progress *task.Progress, include *DiscussionTopicsGetASingleTopicInclude, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s", courseID, topicID)
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsGetTheFullTopic API call: Return a cached structure of the discussion topic, containing all entries,
// their authors, and their message bodies. May require (depending on the topic) that the user has posted in the topic.
// If it is required, and the user has not posted, will respond with a 403 Forbidden status and the body
// 'require_initial_post'. In some rare situations, this cached structure may not be available yet. In that case, the
// server will respond with a 503 error, and the caller should try again soon. The response is an object containing the
// following keys: * "participants": A list of summary information on users who have posted to the discussion. Each
// value is an object containing their id, display_name, and avatar_url. * "unread_entries": A list of entry ids that
// are unread by the current user. this implies that any entry not in this list is read. * "entry_ratings": A map of
// entry ids to ratings by the current user. Entries not in this list have no rating. Only populated if rating is
// enabled. * "forced_entries": A list of entry ids that have forced_read_state set to true. This flag is meant to
// indicate the entry's read_state has been manually set to 'unread' by the user, so the entry should not be
// automatically marked as read. * "view": A threaded view of all the entries in the discussion, containing the id,
// user_id, and message. * "new_entries": Because this view is eventually consistent, it's possible that newly created
// or updated entries won't yet be reflected in the view. If the application wants to also get a flat list of all
// entries not yet reflected in the view, pass include_new_entries=1 to the request and this array of entries will be
// returned. These entries are returned in a flat array, in ascending created_at order.
func (c *Canvas) DiscussionTopicsGetTheFullTopic(progress *task.Progress, courseID string, topicID string) (*DiscussionTopicFullView, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/view", courseID, topicID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &DiscussionTopicFullView{}
	}
	var res *DiscussionTopicFullView
	callback := func(obj interface{}) error {
		res = obj.(*DiscussionTopicFullView)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsPostAnEntry API call: Create a new entry in a discussion topic. Returns a json representation of the
// created entry (see documentation for 'entries' method) on success.
func (c *Canvas) DiscussionTopicsPostAnEntry(progress *task.Progress, message *interface{}, attachment *interface{}, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/entries.json", courseID, topicID)
	params := map[string]interface{}{}
	if message != nil {
		params["message"] = *message
	}
	if attachment != nil {
		params["attachment"] = *attachment
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsListTopicEntries API call: Retrieve the (paginated) top-level entries in a discussion topic. May
// require (depending on the topic) that the user has posted in the topic. If it is required, and the user has not
// posted, will respond with a 403 Forbidden status and the body 'require_initial_post'. Will include the 10 most recent
// replies, if any, for each entry returned. If the topic is a root topic with children corresponding to groups of a
// group assignment, entries from those subtopics for which the user belongs to the corresponding group will be
// returned. Ordering of returned entries is newest-first by posting timestamp (reply activity is ignored).
func (c *Canvas) DiscussionTopicsListTopicEntries(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsPostAReply API call: Add a reply to an entry in a discussion topic. Returns a json representation of
// the created reply (see documentation for 'replies' method) on success. May require (depending on the topic) that the
// user has posted in the topic. If it is required, and the user has not posted, will respond with a 403 Forbidden
// status and the body 'require_initial_post'.
func (c *Canvas) DiscussionTopicsPostAReply(progress *task.Progress, message *interface{}, attachment *interface{}, courseID string, topicID string, entryID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/entries/%s/replies.json", courseID, topicID, entryID)
	params := map[string]interface{}{}
	if message != nil {
		params["message"] = *message
	}
	if attachment != nil {
		params["attachment"] = *attachment
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsListEntryReplies API call: Retrieve the (paginated) replies to a top-level entry in a discussion
// topic. May require (depending on the topic) that the user has posted in the topic. If it is required, and the user
// has not posted, will respond with a 403 Forbidden status and the body 'require_initial_post'. Ordering of returned
// entries is newest-first by creation timestamp.
func (c *Canvas) DiscussionTopicsListEntryReplies(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsListEntries API call: Retrieve a paginated list of discussion entries, given a list of ids. May
// require (depending on the topic) that the user has posted in the topic. If it is required, and the user has not
// posted, will respond with a 403 Forbidden status and the body 'require_initial_post'.
func (c *Canvas) DiscussionTopicsListEntries(progress *task.Progress, ids *string, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/entry_list", courseID, topicID)
	params := map[string]interface{}{}
	if ids != nil {
		params["ids"] = *ids
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsMarkTopicAsRead API call: Mark the initial text of the discussion topic as read. No request fields
// are necessary. On success, the response will be 204 No Content with an empty body.
func (c *Canvas) DiscussionTopicsMarkTopicAsRead(progress *task.Progress, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/read.json", courseID, topicID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsMarkTopicAsUnread API call: Mark the initial text of the discussion topic as unread. No request
// fields are necessary. On success, the response will be 204 No Content with an empty body.
func (c *Canvas) DiscussionTopicsMarkTopicAsUnread(progress *task.Progress, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/read.json", courseID, topicID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsMarkAllEntriesAsRead API call: Mark the discussion topic and all its entries as read. No request
// fields are necessary.
func (c *Canvas) DiscussionTopicsMarkAllEntriesAsRead(progress *task.Progress, forcedReadState *bool, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/read_all.json", courseID, topicID)
	params := map[string]interface{}{}
	if forcedReadState != nil {
		params["forced_read_state"] = *forcedReadState
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsMarkAllEntriesAsUnread API call: Mark the discussion topic and all its entries as unread. No request
// fields are necessary.
func (c *Canvas) DiscussionTopicsMarkAllEntriesAsUnread(progress *task.Progress, forcedReadState *bool, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/read_all.json", courseID, topicID)
	params := map[string]interface{}{}
	if forcedReadState != nil {
		params["forced_read_state"] = *forcedReadState
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsMarkEntryAsRead API call: Mark a discussion entry as read. No request fields are necessary.
func (c *Canvas) DiscussionTopicsMarkEntryAsRead(progress *task.Progress, forcedReadState *bool, courseID string, topicID string, entryID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/entries/%s/read.json", courseID, topicID, entryID)
	params := map[string]interface{}{}
	if forcedReadState != nil {
		params["forced_read_state"] = *forcedReadState
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsMarkEntryAsUnread API call: Mark a discussion entry as unread. No request fields are necessary.
func (c *Canvas) DiscussionTopicsMarkEntryAsUnread(progress *task.Progress, forcedReadState *bool, courseID string, topicID string, entryID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/entries/%s/read.json", courseID, topicID, entryID)
	params := map[string]interface{}{}
	if forcedReadState != nil {
		params["forced_read_state"] = *forcedReadState
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsRateEntry API call: Rate a discussion entry.
func (c *Canvas) DiscussionTopicsRateEntry(progress *task.Progress, rating *int, courseID string, topicID string, entryID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/entries/%s/rating.json", courseID, topicID, entryID)
	params := map[string]interface{}{}
	if rating != nil {
		params["rating"] = *rating
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsSubscribeToATopic API call: Subscribe to a topic to receive notifications about new entries On
// success, the response will be 204 No Content with an empty body
func (c *Canvas) DiscussionTopicsSubscribeToATopic(progress *task.Progress, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/subscribed.json", courseID, topicID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsUnsubscribeFromATopic API call: Unsubscribe from a topic to stop receiving notifications about new
// entries On success, the response will be 204 No Content with an empty body
func (c *Canvas) DiscussionTopicsUnsubscribeFromATopic(progress *task.Progress, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s/subscribed.json", courseID, topicID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsListDiscussionTopics API call: Returns the paginated list of discussion topics for this course or
// group.
func (c *Canvas) DiscussionTopicsListDiscussionTopics(progress *task.Progress, include *DiscussionTopicsListDiscussionTopicsInclude, orderBy *DiscussionTopicsListDiscussionTopicsOrderBy, scope *DiscussionTopicsListDiscussionTopicsScope, onlyAnnouncements *bool, filterBy *DiscussionTopicsListDiscussionTopicsFilterBy, searchTerm *string, excludeContextModuleLockedTopics *bool, courseID string) ([]DiscussionTopic, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics", courseID)
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if orderBy != nil {
		params["order_by"] = *orderBy
	}
	if scope != nil {
		params["scope"] = *scope
	}
	if onlyAnnouncements != nil {
		params["only_announcements"] = *onlyAnnouncements
	}
	if filterBy != nil {
		params["filter_by"] = *filterBy
	}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if excludeContextModuleLockedTopics != nil {
		params["exclude_context_module_locked_topics"] = *excludeContextModuleLockedTopics
	}
	responseCtor := func() interface{} {
		return &[]DiscussionTopic{}
	}
	var res []DiscussionTopic
	callback := func(obj interface{}) error {
		arr := *obj.(*[]DiscussionTopic)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsCreateANewDiscussionTopic API call: Create an new discussion topic for the course or group.
func (c *Canvas) DiscussionTopicsCreateANewDiscussionTopic(progress *task.Progress, title *string, message *string, discussionType *DiscussionTopicsCreateANewDiscussionTopicDiscussionType, published *bool, delayedPostAt *time.Time, allowRating *bool, lockAt *time.Time, podcastEnabled *bool, podcastHasStudentPosts *bool, requireInitialPost *bool, assignment *Assignment, isAnnouncement *bool, pinned *bool, positionAfter *string, groupCategoryID *int, onlyGradersCanRate *bool, sortByRating *bool, attachment *File, specificSections *string, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics", courseID)
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if message != nil {
		params["message"] = *message
	}
	if discussionType != nil {
		params["discussion_type"] = *discussionType
	}
	if published != nil {
		params["published"] = *published
	}
	if delayedPostAt != nil {
		params["delayed_post_at"] = *delayedPostAt
	}
	if allowRating != nil {
		params["allow_rating"] = *allowRating
	}
	if lockAt != nil {
		params["lock_at"] = *lockAt
	}
	if podcastEnabled != nil {
		params["podcast_enabled"] = *podcastEnabled
	}
	if podcastHasStudentPosts != nil {
		params["podcast_has_student_posts"] = *podcastHasStudentPosts
	}
	if requireInitialPost != nil {
		params["require_initial_post"] = *requireInitialPost
	}
	if assignment != nil {
		params["assignment"] = *assignment
	}
	if isAnnouncement != nil {
		params["is_announcement"] = *isAnnouncement
	}
	if pinned != nil {
		params["pinned"] = *pinned
	}
	if positionAfter != nil {
		params["position_after"] = *positionAfter
	}
	if groupCategoryID != nil {
		params["group_category_id"] = *groupCategoryID
	}
	if onlyGradersCanRate != nil {
		params["only_graders_can_rate"] = *onlyGradersCanRate
	}
	if sortByRating != nil {
		params["sort_by_rating"] = *sortByRating
	}
	if attachment != nil {
		params["attachment"] = *attachment
	}
	if specificSections != nil {
		params["specific_sections"] = *specificSections
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsUpdateATopic API call: Update an existing discussion topic for the course or group.
func (c *Canvas) DiscussionTopicsUpdateATopic(progress *task.Progress, title *string, message *string, discussionType *DiscussionTopicsUpdateATopicDiscussionType, published *bool, delayedPostAt *time.Time, lockAt *time.Time, podcastEnabled *bool, podcastHasStudentPosts *bool, requireInitialPost *bool, assignment *Assignment, isAnnouncement *bool, pinned *bool, positionAfter *string, groupCategoryID *int, allowRating *bool, onlyGradersCanRate *bool, sortByRating *bool, specificSections *string, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s", courseID, topicID)
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if message != nil {
		params["message"] = *message
	}
	if discussionType != nil {
		params["discussion_type"] = *discussionType
	}
	if published != nil {
		params["published"] = *published
	}
	if delayedPostAt != nil {
		params["delayed_post_at"] = *delayedPostAt
	}
	if lockAt != nil {
		params["lock_at"] = *lockAt
	}
	if podcastEnabled != nil {
		params["podcast_enabled"] = *podcastEnabled
	}
	if podcastHasStudentPosts != nil {
		params["podcast_has_student_posts"] = *podcastHasStudentPosts
	}
	if requireInitialPost != nil {
		params["require_initial_post"] = *requireInitialPost
	}
	if assignment != nil {
		params["assignment"] = *assignment
	}
	if isAnnouncement != nil {
		params["is_announcement"] = *isAnnouncement
	}
	if pinned != nil {
		params["pinned"] = *pinned
	}
	if positionAfter != nil {
		params["position_after"] = *positionAfter
	}
	if groupCategoryID != nil {
		params["group_category_id"] = *groupCategoryID
	}
	if allowRating != nil {
		params["allow_rating"] = *allowRating
	}
	if onlyGradersCanRate != nil {
		params["only_graders_can_rate"] = *onlyGradersCanRate
	}
	if sortByRating != nil {
		params["sort_by_rating"] = *sortByRating
	}
	if specificSections != nil {
		params["specific_sections"] = *specificSections
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsDeleteATopic API call: Deletes the discussion topic. This will also delete the assignment, if it's an
// assignment discussion.
func (c *Canvas) DiscussionTopicsDeleteATopic(progress *task.Progress, courseID string, topicID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/discussion_topics/%s", courseID, topicID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// DiscussionTopicsReorderPinnedTopics API call: Puts the pinned discussion topics in the specified order. All pinned
// topics should be included.
func (c *Canvas) DiscussionTopicsReorderPinnedTopics(progress *task.Progress, order *int) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if order != nil {
		params["order"] = *order
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentsListEnrollments API call: Depending on the URL given, return a paginated list of either (1) all of the
// enrollments in a course, (2) all of the enrollments in a section or (3) all of a user's enrollments. This includes
// student, teacher, TA, and observer enrollments. If a user has multiple enrollments in a context (e.g. as a teacher
// and a student or in multiple course sections), each enrollment will be listed separately. note: Currently, only a
// root level admin user can return other users' enrollments. A user can, however, return his/her own enrollments.
func (c *Canvas) EnrollmentsListEnrollments(progress *task.Progress, typeName *string, role *string, state *EnrollmentsListEnrollmentsState, include *EnrollmentsListEnrollmentsInclude, userID *string, gradingPeriodID *int, enrollmentTermID *int, sisAccountID *string, sisCourseID *string, sisSectionID *string, sisUserID *string, createdForSisID *bool) ([]Enrollment, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if typeName != nil {
		params["type"] = *typeName
	}
	if role != nil {
		params["role"] = *role
	}
	if state != nil {
		params["state"] = *state
	}
	if include != nil {
		params["include"] = *include
	}
	if userID != nil {
		params["user_id"] = *userID
	}
	if gradingPeriodID != nil {
		params["grading_period_id"] = *gradingPeriodID
	}
	if enrollmentTermID != nil {
		params["enrollment_term_id"] = *enrollmentTermID
	}
	if sisAccountID != nil {
		params["sis_account_id"] = *sisAccountID
	}
	if sisCourseID != nil {
		params["sis_course_id"] = *sisCourseID
	}
	if sisSectionID != nil {
		params["sis_section_id"] = *sisSectionID
	}
	if sisUserID != nil {
		params["sis_user_id"] = *sisUserID
	}
	if createdForSisID != nil {
		params["created_for_sis_id"] = *createdForSisID
	}
	responseCtor := func() interface{} {
		return &[]Enrollment{}
	}
	var res []Enrollment
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Enrollment)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentsEnrollmentByID API call: Get an Enrollment object by Enrollment ID
func (c *Canvas) EnrollmentsEnrollmentByID(progress *task.Progress, id *int) (*Enrollment, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	responseCtor := func() interface{} {
		return &Enrollment{}
	}
	var res *Enrollment
	callback := func(obj interface{}) error {
		res = obj.(*Enrollment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentsEnrollAUser API call: Create a new user enrollment for a course or section.
func (c *Canvas) EnrollmentsEnrollAUser(progress *task.Progress, enrollment *string) (*Enrollment, error) {
	endpoint := fmt.Sprintf("courses/:course_id/enrollments")
	params := map[string]interface{}{}
	if enrollment != nil {
		params["enrollment"] = *enrollment
	}
	responseCtor := func() interface{} {
		return &Enrollment{}
	}
	var res *Enrollment
	callback := func(obj interface{}) error {
		res = obj.(*Enrollment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentsConcludeDeactivateOrDeleteAnEnrollment API call: Conclude, deactivate, or delete an enrollment. If the
// +task+ argument isn't given, the enrollment will be concluded.
func (c *Canvas) EnrollmentsConcludeDeactivateOrDeleteAnEnrollment(progress *task.Progress, task *EnrollmentsConcludeDeactivateOrDeleteAnEnrollmentTask) (*Enrollment, error) {
	endpoint := fmt.Sprintf("courses/:course_id/enrollments/:enrollment_id")
	params := map[string]interface{}{}
	if task != nil {
		params["task"] = *task
	}
	responseCtor := func() interface{} {
		return &Enrollment{}
	}
	var res *Enrollment
	callback := func(obj interface{}) error {
		res = obj.(*Enrollment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentsAcceptCourseInvitation API call: accepts a pending course invitation for the current user
func (c *Canvas) EnrollmentsAcceptCourseInvitation(progress *task.Progress, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/enrollments/:id/accept", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentsRejectCourseInvitation API call: rejects a pending course invitation for the current user
func (c *Canvas) EnrollmentsRejectCourseInvitation(progress *task.Progress, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/enrollments/:id/reject", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentsReActivateAnEnrollment API call: Activates an inactive enrollment
func (c *Canvas) EnrollmentsReActivateAnEnrollment(progress *task.Progress) (*Enrollment, error) {
	endpoint := fmt.Sprintf("courses/:course_id/enrollments/:enrollment_id/reactivate")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Enrollment{}
	}
	var res *Enrollment
	callback := func(obj interface{}) error {
		res = obj.(*Enrollment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentsAddsLastAttendedDateToStudentEnrollmentInCourse API call
func (c *Canvas) EnrollmentsAddsLastAttendedDateToStudentEnrollmentInCourse(progress *task.Progress) (*Enrollment, error) {
	endpoint := fmt.Sprintf("courses/:course_id/user/:user_id/last_attended")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Enrollment{}
	}
	var res *Enrollment
	callback := func(obj interface{}) error {
		res = obj.(*Enrollment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EPubExportsListCoursesWithTheirLatestEPubExport API call: A paginated list of all courses a user is actively
// participating in, and the latest ePub export associated with the user & course.
func (c *Canvas) EPubExportsListCoursesWithTheirLatestEPubExport(progress *task.Progress) ([]CourseEpubExport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]CourseEpubExport{}
	}
	var res []CourseEpubExport
	callback := func(obj interface{}) error {
		arr := *obj.(*[]CourseEpubExport)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EPubExportsCreateEPubExport API call: Begin an ePub export for a course. You can use the {api:ProgressController#show
// Progress API} to track the progress of the export. The export's progress is linked to with the _progress_url_ value.
// When the export completes, use the {api:EpubExportsController#show Show content export} endpoint to retrieve a
// download URL for the exported content.
func (c *Canvas) EPubExportsCreateEPubExport(progress *task.Progress) (*EpubExport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &EpubExport{}
	}
	var res *EpubExport
	callback := func(obj interface{}) error {
		res = obj.(*EpubExport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EPubExportsShowEPubExport API call: Get information about a single ePub export.
func (c *Canvas) EPubExportsShowEPubExport(progress *task.Progress) (*EpubExport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &EpubExport{}
	}
	var res *EpubExport
	callback := func(obj interface{}) error {
		res = obj.(*EpubExport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ErrorReportsCreateErrorReport API call: Create a new error report documenting an experienced problem Performs the
// same action as when a user uses the "help -> report a problem" dialog.
func (c *Canvas) ErrorReportsCreateErrorReport(progress *task.Progress, err *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("error_reports")
	params := map[string]interface{}{}
	if err != nil {
		params["error"] = *err
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AnnouncementExternalFeedsListExternalFeeds API call: Returns the paginated list of External Feeds this course or
// group.
func (c *Canvas) AnnouncementExternalFeedsListExternalFeeds(progress *task.Progress, courseID string) ([]ExternalFeed, error) {
	endpoint := fmt.Sprintf("courses/%s/external_feeds", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]ExternalFeed{}
	}
	var res []ExternalFeed
	callback := func(obj interface{}) error {
		arr := *obj.(*[]ExternalFeed)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AnnouncementExternalFeedsCreateAnExternalFeed API call: Create a new external feed for the course or group.
func (c *Canvas) AnnouncementExternalFeedsCreateAnExternalFeed(progress *task.Progress, url *string, headerMatch *bool, verbosity *AnnouncementExternalFeedsCreateAnExternalFeedVerbosity, courseID string) (*ExternalFeed, error) {
	endpoint := fmt.Sprintf("courses/%s/external_feeds", courseID)
	params := map[string]interface{}{}
	if url != nil {
		params["url"] = *url
	}
	if headerMatch != nil {
		params["header_match"] = *headerMatch
	}
	if verbosity != nil {
		params["verbosity"] = *verbosity
	}
	responseCtor := func() interface{} {
		return &ExternalFeed{}
	}
	var res *ExternalFeed
	callback := func(obj interface{}) error {
		res = obj.(*ExternalFeed)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AnnouncementExternalFeedsDeleteAnExternalFeed API call: Deletes the external feed.
func (c *Canvas) AnnouncementExternalFeedsDeleteAnExternalFeed(progress *task.Progress, courseID string, feedID string) (*ExternalFeed, error) {
	endpoint := fmt.Sprintf("courses/%s/external_feeds/%s", courseID, feedID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &ExternalFeed{}
	}
	var res *ExternalFeed
	callback := func(obj interface{}) error {
		res = obj.(*ExternalFeed)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ExternalToolsListExternalTools API call: Returns the paginated list of external tools for the current context. See
// the get request docs for a single tool for a list of properties on an external tool.
func (c *Canvas) ExternalToolsListExternalTools(progress *task.Progress, searchTerm *string, selectable *bool, includeParents *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if selectable != nil {
		params["selectable"] = *selectable
	}
	if includeParents != nil {
		params["include_parents"] = *includeParents
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ExternalToolsGetASessionlessLaunchURLForAnExternalTool API call: Returns a sessionless launch url for an external
// tool. NOTE: Either the id or url must be provided unless launch_type is assessment or module_item.
func (c *Canvas) ExternalToolsGetASessionlessLaunchURLForAnExternalTool(progress *task.Progress, id *string, url *string, assignmentID *string, moduleItemID *string, launchType *ExternalToolsGetASessionlessLaunchURLForAnExternalToolLaunchType, courseID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/external_tools/sessionless_launch", courseID)
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	if url != nil {
		params["url"] = *url
	}
	if assignmentID != nil {
		params["assignment_id"] = *assignmentID
	}
	if moduleItemID != nil {
		params["module_item_id"] = *moduleItemID
	}
	if launchType != nil {
		params["launch_type"] = *launchType
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ExternalToolsGetASingleExternalTool API call: Returns the specified external tool.
func (c *Canvas) ExternalToolsGetASingleExternalTool(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ExternalToolsCreateAnExternalTool API call: Create an external tool in the specified course/account. The created tool
// will be returned, see the "show" endpoint for an example. If a client ID is supplied canvas will attempt to create a
// context external tool using the LTI 1.3 standard.
func (c *Canvas) ExternalToolsCreateAnExternalTool(progress *task.Progress, clientID *string, name *string, privacyLevel *ExternalToolsCreateAnExternalToolPrivacyLevel, consumerKey *string, sharedSecret *string, description *string, url *string, domain *string, iconURL *string, text *string, customFields *string, accountNavigation *string, userNavigation *string, courseHomeSubNavigation *string, courseNavigation *bool, editorButton *string, homeworkSubmission *string, linkSelection *string, migrationSelection *string, toolConfiguration *string, resourceSelection *string, configType *string, configXML *string, configURL *string, notSelectable *bool, oauthCompliant *bool, accountID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("accounts/%s/external_tools", accountID)
	params := map[string]interface{}{}
	if clientID != nil {
		params["client_id"] = *clientID
	}
	if name != nil {
		params["name"] = *name
	}
	if privacyLevel != nil {
		params["privacy_level"] = *privacyLevel
	}
	if consumerKey != nil {
		params["consumer_key"] = *consumerKey
	}
	if sharedSecret != nil {
		params["shared_secret"] = *sharedSecret
	}
	if description != nil {
		params["description"] = *description
	}
	if url != nil {
		params["url"] = *url
	}
	if domain != nil {
		params["domain"] = *domain
	}
	if iconURL != nil {
		params["icon_url"] = *iconURL
	}
	if text != nil {
		params["text"] = *text
	}
	if customFields != nil {
		params["custom_fields"] = *customFields
	}
	if accountNavigation != nil {
		params["account_navigation"] = *accountNavigation
	}
	if userNavigation != nil {
		params["user_navigation"] = *userNavigation
	}
	if courseHomeSubNavigation != nil {
		params["course_home_sub_navigation"] = *courseHomeSubNavigation
	}
	if courseNavigation != nil {
		params["course_navigation"] = *courseNavigation
	}
	if editorButton != nil {
		params["editor_button"] = *editorButton
	}
	if homeworkSubmission != nil {
		params["homework_submission"] = *homeworkSubmission
	}
	if linkSelection != nil {
		params["link_selection"] = *linkSelection
	}
	if migrationSelection != nil {
		params["migration_selection"] = *migrationSelection
	}
	if toolConfiguration != nil {
		params["tool_configuration"] = *toolConfiguration
	}
	if resourceSelection != nil {
		params["resource_selection"] = *resourceSelection
	}
	if configType != nil {
		params["config_type"] = *configType
	}
	if configXML != nil {
		params["config_xml"] = *configXML
	}
	if configURL != nil {
		params["config_url"] = *configURL
	}
	if notSelectable != nil {
		params["not_selectable"] = *notSelectable
	}
	if oauthCompliant != nil {
		params["oauth_compliant"] = *oauthCompliant
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ExternalToolsEditAnExternalTool API call: Update the specified external tool. Uses same parameters as create
func (c *Canvas) ExternalToolsEditAnExternalTool(progress *task.Progress, courseID string, externalToolID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/external_tools/%s", courseID, externalToolID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ExternalToolsDeleteAnExternalTool API call: Remove the specified external tool
func (c *Canvas) ExternalToolsDeleteAnExternalTool(progress *task.Progress, courseID string, externalToolID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/external_tools/%s", courseID, externalToolID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FavoritesListFavoriteCourses API call: Retrieve the paginated list of favorite courses for the current user. If the
// user has not chosen any favorites, then a selection of currently enrolled courses will be returned.
func (c *Canvas) FavoritesListFavoriteCourses(progress *task.Progress, excludeBlueprintCourses *bool) ([]Course, error) {
	endpoint := fmt.Sprintf("users/self/favorites/courses")
	params := map[string]interface{}{}
	if excludeBlueprintCourses != nil {
		params["exclude_blueprint_courses"] = *excludeBlueprintCourses
	}
	responseCtor := func() interface{} {
		return &[]Course{}
	}
	var res []Course
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Course)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FavoritesListFavoriteGroups API call: Retrieve the paginated list of favorite groups for the current user. If the
// user has not chosen any favorites, then a selection of groups that the user is a member of will be returned.
func (c *Canvas) FavoritesListFavoriteGroups(progress *task.Progress) ([]Group, error) {
	endpoint := fmt.Sprintf("users/self/favorites/groups")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Group{}
	}
	var res []Group
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Group)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FavoritesAddCourseToFavorites API call: Add a course to the current user's favorites.  If the course is already in
// the user's favorites, nothing happens.
func (c *Canvas) FavoritesAddCourseToFavorites(progress *task.Progress, id *string) (*Favorite, error) {
	endpoint := fmt.Sprintf("users/self/favorites/courses/1170")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	responseCtor := func() interface{} {
		return &Favorite{}
	}
	var res *Favorite
	callback := func(obj interface{}) error {
		res = obj.(*Favorite)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FavoritesAddGroupToFavorites API call: Add a group to the current user's favorites.  If the group is already in the
// user's favorites, nothing happens.
func (c *Canvas) FavoritesAddGroupToFavorites(progress *task.Progress, id *string) (*Favorite, error) {
	endpoint := fmt.Sprintf("users/self/favorites/group/1170")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	responseCtor := func() interface{} {
		return &Favorite{}
	}
	var res *Favorite
	callback := func(obj interface{}) error {
		res = obj.(*Favorite)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FavoritesRemoveCourseFromFavorites API call: Remove a course from the current user's favorites.
func (c *Canvas) FavoritesRemoveCourseFromFavorites(progress *task.Progress, id *string) (*Favorite, error) {
	endpoint := fmt.Sprintf("users/self/favorites/courses/1170")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	responseCtor := func() interface{} {
		return &Favorite{}
	}
	var res *Favorite
	callback := func(obj interface{}) error {
		res = obj.(*Favorite)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FavoritesRemoveGroupFromFavorites API call: Remove a group from the current user's favorites.
func (c *Canvas) FavoritesRemoveGroupFromFavorites(progress *task.Progress, id *string) (*Favorite, error) {
	endpoint := fmt.Sprintf("users/self/favorites/groups/1170")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	responseCtor := func() interface{} {
		return &Favorite{}
	}
	var res *Favorite
	callback := func(obj interface{}) error {
		res = obj.(*Favorite)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FavoritesResetCourseFavorites API call: Reset the current user's course favorites to the default automatically
// generated list of enrolled courses
func (c *Canvas) FavoritesResetCourseFavorites(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/self/favorites/courses")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FavoritesResetGroupFavorites API call: Reset the current user's group favorites to the default automatically
// generated list of enrolled group
func (c *Canvas) FavoritesResetGroupFavorites(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/self/favorites/group")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FeatureFlagsListFeatures API call: A paginated list of all features that apply to a given Account, Course, or User.
func (c *Canvas) FeatureFlagsListFeatures(progress *task.Progress) ([]Feature, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Feature{}
	}
	var res []Feature
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Feature)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FeatureFlagsListEnabledFeatures API call: A paginated list of all features that are enabled on a given Account,
// Course, or User. Only the feature names are returned.
func (c *Canvas) FeatureFlagsListEnabledFeatures(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FeatureFlagsGetFeatureFlag API call: Get the feature flag that applies to a given Account, Course, or User. The flag
// may be defined on the object, or it may be inherited from a parent account. You can look at the context_id and
// context_type of the returned object to determine which is the case. If these fields are missing, then the object is
// the global Canvas default.
func (c *Canvas) FeatureFlagsGetFeatureFlag(progress *task.Progress) (*FeatureFlag, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &FeatureFlag{}
	}
	var res *FeatureFlag
	callback := func(obj interface{}) error {
		res = obj.(*FeatureFlag)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FeatureFlagsSetFeatureFlag API call: Set a feature flag for a given Account, Course, or User. This call will fail if
// a parent account sets a feature flag for the same feature in any state other than "allowed".
func (c *Canvas) FeatureFlagsSetFeatureFlag(progress *task.Progress, state *FeatureFlagsSetFeatureFlagState) (*FeatureFlag, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if state != nil {
		params["state"] = *state
	}
	responseCtor := func() interface{} {
		return &FeatureFlag{}
	}
	var res *FeatureFlag
	callback := func(obj interface{}) error {
		res = obj.(*FeatureFlag)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FeatureFlagsRemoveFeatureFlag API call: Remove feature flag for a given Account, Course, or User.  (Note that the
// flag must be defined on the Account, Course, or User directly.)  The object will then inherit the feature flags from
// a higher account, if any exist.  If this flag was 'on' or 'off', then lower-level account flags that were masked by
// this one will apply again.
func (c *Canvas) FeatureFlagsRemoveFeatureFlag(progress *task.Progress) (*FeatureFlag, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &FeatureFlag{}
	}
	var res *FeatureFlag
	callback := func(obj interface{}) error {
		res = obj.(*FeatureFlag)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesGetQuotaInformation API call: Returns the total and used storage quota for the course, group, or user.
func (c *Canvas) FilesGetQuotaInformation(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/1/files/quota")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesListFiles API call: Returns the paginated list of files for the folder or course.
func (c *Canvas) FilesListFiles(progress *task.Progress, contentTypes *string, excludeContentTypes *string, searchTerm *string, include *FilesListFilesInclude, only []interface{}, sort *FilesListFilesSort, order *FilesListFilesOrder, folderID string) ([]File, error) {
	endpoint := fmt.Sprintf("folders/%s/files", folderID)
	params := map[string]interface{}{}
	if contentTypes != nil {
		params["content_types"] = *contentTypes
	}
	if excludeContentTypes != nil {
		params["exclude_content_types"] = *excludeContentTypes
	}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if include != nil {
		params["include"] = *include
	}
	if only != nil && len(only) > 0 {
		params["only"] = only
	}
	if sort != nil {
		params["sort"] = *sort
	}
	if order != nil {
		params["order"] = *order
	}
	responseCtor := func() interface{} {
		return &[]File{}
	}
	var res []File
	callback := func(obj interface{}) error {
		arr := *obj.(*[]File)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesGetPublicInlinePreviewURL API call: Determine the URL that should be used for inline preview of the file.
func (c *Canvas) FilesGetPublicInlinePreviewURL(progress *task.Progress, submissionID *int) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("files/1/public_url")
	params := map[string]interface{}{}
	if submissionID != nil {
		params["submission_id"] = *submissionID
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesGetFile API call: Returns the standard attachment json object
func (c *Canvas) FilesGetFile(progress *task.Progress, include *FilesGetFileInclude, fileID string) (*File, error) {
	endpoint := fmt.Sprintf("files/%s", fileID)
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &File{}
	}
	var res *File
	callback := func(obj interface{}) error {
		res = obj.(*File)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesUpdateFile API call: Update some settings on the specified file
func (c *Canvas) FilesUpdateFile(progress *task.Progress, name *string, parentFolderID *string, onDuplicate *FilesUpdateFileOnDuplicate, lockAt *time.Time, unlockAt *time.Time, locked *bool, hidden *bool, fileID string) (*File, error) {
	endpoint := fmt.Sprintf("files/%s", fileID)
	params := map[string]interface{}{}
	if name != nil {
		params["name"] = *name
	}
	if parentFolderID != nil {
		params["parent_folder_id"] = *parentFolderID
	}
	if onDuplicate != nil {
		params["on_duplicate"] = *onDuplicate
	}
	if lockAt != nil {
		params["lock_at"] = *lockAt
	}
	if unlockAt != nil {
		params["unlock_at"] = *unlockAt
	}
	if locked != nil {
		params["locked"] = *locked
	}
	if hidden != nil {
		params["hidden"] = *hidden
	}
	responseCtor := func() interface{} {
		return &File{}
	}
	var res *File
	callback := func(obj interface{}) error {
		res = obj.(*File)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesDeleteFile API call: Remove the specified file. Unlike most other DELETE endpoints, using this endpoint will
// result in comprehensive, irretrievable destruction of the file. It should be used with the `replace` parameter set to
// true in cases where the file preview also needs to be destroyed (such as to remove files that violate privacy laws).
func (c *Canvas) FilesDeleteFile(progress *task.Progress, replace *bool, fileID string) (*File, error) {
	endpoint := fmt.Sprintf("files/%s", fileID)
	params := map[string]interface{}{}
	if replace != nil {
		params["replace"] = *replace
	}
	responseCtor := func() interface{} {
		return &File{}
	}
	var res *File
	callback := func(obj interface{}) error {
		res = obj.(*File)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesResetLinkVerifier API call: Resets the link verifier. Any existing links to the file using the previous
// hard-coded "verifier" parameter will no longer automatically grant access. Must have manage files and become other
// users permissions
func (c *Canvas) FilesResetLinkVerifier(progress *task.Progress, fileID string) (*File, error) {
	endpoint := fmt.Sprintf("files/%s/reset_verifier", fileID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &File{}
	}
	var res *File
	callback := func(obj interface{}) error {
		res = obj.(*File)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesListFolders API call
func (c *Canvas) FilesListFolders(progress *task.Progress, folderID string) ([]Folder, error) {
	endpoint := fmt.Sprintf("folders/%s/folders", folderID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Folder{}
	}
	var res []Folder
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Folder)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesListAllFolders API call
func (c *Canvas) FilesListAllFolders(progress *task.Progress, courseID string) ([]Folder, error) {
	endpoint := fmt.Sprintf("courses/%s/folders", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Folder{}
	}
	var res []Folder
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Folder)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesResolvePath API call
func (c *Canvas) FilesResolvePath(progress *task.Progress, courseID string) ([]Folder, error) {
	endpoint := fmt.Sprintf("courses/%s/folders/by_path/foo/bar/baz", courseID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Folder{}
	}
	var res []Folder
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Folder)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesGetFolder API call
func (c *Canvas) FilesGetFolder(progress *task.Progress, folderID string) (*Folder, error) {
	endpoint := fmt.Sprintf("folders/%s", folderID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Folder{}
	}
	var res *Folder
	callback := func(obj interface{}) error {
		res = obj.(*Folder)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesUpdateFolder API call
func (c *Canvas) FilesUpdateFolder(progress *task.Progress, name *string, parentFolderID *string, lockAt *time.Time, unlockAt *time.Time, locked *bool, hidden *bool, position *int, folderID string) (*Folder, error) {
	endpoint := fmt.Sprintf("folders/%s", folderID)
	params := map[string]interface{}{}
	if name != nil {
		params["name"] = *name
	}
	if parentFolderID != nil {
		params["parent_folder_id"] = *parentFolderID
	}
	if lockAt != nil {
		params["lock_at"] = *lockAt
	}
	if unlockAt != nil {
		params["unlock_at"] = *unlockAt
	}
	if locked != nil {
		params["locked"] = *locked
	}
	if hidden != nil {
		params["hidden"] = *hidden
	}
	if position != nil {
		params["position"] = *position
	}
	responseCtor := func() interface{} {
		return &Folder{}
	}
	var res *Folder
	callback := func(obj interface{}) error {
		res = obj.(*Folder)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesCreateFolder API call
func (c *Canvas) FilesCreateFolder(progress *task.Progress, name *string, parentFolderID *string, parentFolderPath *string, lockAt *time.Time, unlockAt *time.Time, locked *bool, hidden *bool, position *int, courseID string) (*Folder, error) {
	endpoint := fmt.Sprintf("courses/%s/folders", courseID)
	params := map[string]interface{}{}
	if name != nil {
		params["name"] = *name
	}
	if parentFolderID != nil {
		params["parent_folder_id"] = *parentFolderID
	}
	if parentFolderPath != nil {
		params["parent_folder_path"] = *parentFolderPath
	}
	if lockAt != nil {
		params["lock_at"] = *lockAt
	}
	if unlockAt != nil {
		params["unlock_at"] = *unlockAt
	}
	if locked != nil {
		params["locked"] = *locked
	}
	if hidden != nil {
		params["hidden"] = *hidden
	}
	if position != nil {
		params["position"] = *position
	}
	responseCtor := func() interface{} {
		return &Folder{}
	}
	var res *Folder
	callback := func(obj interface{}) error {
		res = obj.(*Folder)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesDeleteFolder API call
func (c *Canvas) FilesDeleteFolder(progress *task.Progress, force *bool, folderID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("folders/%s", folderID)
	params := map[string]interface{}{}
	if force != nil {
		params["force"] = *force
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesUploadAFile API call: Upload a file to a folder. This API endpoint is the first step in uploading a file. See
// the {file:file_uploads.html File Upload Documentation} for details on the file upload workflow. Only those with the
// "Manage Files" permission on a course or group can upload files to a folder in that course or group.
func (c *Canvas) FilesUploadAFile(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesCopyAFile API call: Copy a file from elsewhere in Canvas into a folder. Copying a file across contexts (between
// courses and users) is permitted, but the source and destination must belong to the same institution.
func (c *Canvas) FilesCopyAFile(progress *task.Progress, sourceFileID *string, onDuplicate *FilesCopyAFileOnDuplicate) (*File, error) {
	endpoint := fmt.Sprintf("folders/123/copy_file")
	params := map[string]interface{}{}
	if sourceFileID != nil {
		params["source_file_id"] = *sourceFileID
	}
	if onDuplicate != nil {
		params["on_duplicate"] = *onDuplicate
	}
	responseCtor := func() interface{} {
		return &File{}
	}
	var res *File
	callback := func(obj interface{}) error {
		res = obj.(*File)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesCopyAFolder API call: Copy a folder (and its contents) from elsewhere in Canvas into a folder. Copying a folder
// across contexts (between courses and users) is permitted, but the source and destination must belong to the same
// institution. If the source and destination folders are in the same context, the source folder may not contain the
// destination folder. A folder will be renamed at its destination if another folder with the same name already exists.
func (c *Canvas) FilesCopyAFolder(progress *task.Progress, sourceFolderID *string) (*Folder, error) {
	endpoint := fmt.Sprintf("folders/123/copy_folder")
	params := map[string]interface{}{}
	if sourceFolderID != nil {
		params["source_folder_id"] = *sourceFolderID
	}
	responseCtor := func() interface{} {
		return &Folder{}
	}
	var res *Folder
	callback := func(obj interface{}) error {
		res = obj.(*Folder)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesGetUploadedMediaFolderForUser API call
func (c *Canvas) FilesGetUploadedMediaFolderForUser(progress *task.Progress) (*Folder, error) {
	endpoint := fmt.Sprintf("courses/1337/folders/media")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Folder{}
	}
	var res *Folder
	callback := func(obj interface{}) error {
		res = obj.(*Folder)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradeChangeLogQueryByAssignment API call: List grade change events for a given assignment.
func (c *Canvas) GradeChangeLogQueryByAssignment(progress *task.Progress, startTime *time.Time, endTime *time.Time) ([]GradeChangeEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &[]GradeChangeEvent{}
	}
	var res []GradeChangeEvent
	callback := func(obj interface{}) error {
		arr := *obj.(*[]GradeChangeEvent)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradeChangeLogQueryByCourse API call: List grade change events for a given course.
func (c *Canvas) GradeChangeLogQueryByCourse(progress *task.Progress, startTime *time.Time, endTime *time.Time) ([]GradeChangeEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &[]GradeChangeEvent{}
	}
	var res []GradeChangeEvent
	callback := func(obj interface{}) error {
		arr := *obj.(*[]GradeChangeEvent)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradeChangeLogQueryByStudent API call: List grade change events for a given student.
func (c *Canvas) GradeChangeLogQueryByStudent(progress *task.Progress, startTime *time.Time, endTime *time.Time) ([]GradeChangeEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &[]GradeChangeEvent{}
	}
	var res []GradeChangeEvent
	callback := func(obj interface{}) error {
		arr := *obj.(*[]GradeChangeEvent)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradeChangeLogQueryByGrader API call: List grade change events for a given grader.
func (c *Canvas) GradeChangeLogQueryByGrader(progress *task.Progress, startTime *time.Time, endTime *time.Time) ([]GradeChangeEvent, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &[]GradeChangeEvent{}
	}
	var res []GradeChangeEvent
	callback := func(obj interface{}) error {
		arr := *obj.(*[]GradeChangeEvent)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradebookHistoryDaysInGradebookHistoryForThisCourse API call: Returns a map of dates to grader/assignment groups
func (c *Canvas) GradebookHistoryDaysInGradebookHistoryForThisCourse(progress *task.Progress, courseID *int) ([]Day, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	responseCtor := func() interface{} {
		return &[]Day{}
	}
	var res []Day
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Day)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradebookHistoryDetailsForAGivenDateInGradebookHistoryForThisCourse API call: Returns the graders who worked on this
// day, along with the assignments they worked on. More details can be obtained by selecting a grader and assignment and
// calling the 'submissions' api endpoint for a given date.
func (c *Canvas) GradebookHistoryDetailsForAGivenDateInGradebookHistoryForThisCourse(progress *task.Progress, courseID *int, date *string) ([]Grader, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	if date != nil {
		params["date"] = *date
	}
	responseCtor := func() interface{} {
		return &[]Grader{}
	}
	var res []Grader
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Grader)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradebookHistoryListsSubmissions API call: Gives a nested list of submission versions
func (c *Canvas) GradebookHistoryListsSubmissions(progress *task.Progress, courseID *int, date *string, graderID *int, assignmentID *int) ([]SubmissionHistory, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	if date != nil {
		params["date"] = *date
	}
	if graderID != nil {
		params["grader_id"] = *graderID
	}
	if assignmentID != nil {
		params["assignment_id"] = *assignmentID
	}
	responseCtor := func() interface{} {
		return &[]SubmissionHistory{}
	}
	var res []SubmissionHistory
	callback := func(obj interface{}) error {
		arr := *obj.(*[]SubmissionHistory)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradebookHistoryListUncollatedSubmissionVersions API call: Gives a paginated, uncollated list of submission versions
// for all matching submissions in the context. This SubmissionVersion objects will not include the +new_grade+ or
// +previous_grade+ keys, only the +grade+; same for +graded_at+ and +grader+.
func (c *Canvas) GradebookHistoryListUncollatedSubmissionVersions(progress *task.Progress, courseID *int, assignmentID *int, userID *int, ascending *bool) ([]SubmissionVersion, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	if assignmentID != nil {
		params["assignment_id"] = *assignmentID
	}
	if userID != nil {
		params["user_id"] = *userID
	}
	if ascending != nil {
		params["ascending"] = *ascending
	}
	responseCtor := func() interface{} {
		return &[]SubmissionVersion{}
	}
	var res []SubmissionVersion
	callback := func(obj interface{}) error {
		arr := *obj.(*[]SubmissionVersion)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradingPeriodsListGradingPeriods API call: Returns the paginated list of grading periods for the current course.
func (c *Canvas) GradingPeriodsListGradingPeriods(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradingPeriodsGetASingleGradingPeriod API call: Returns the grading period with the given id
func (c *Canvas) GradingPeriodsGetASingleGradingPeriod(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradingPeriodsUpdateASingleGradingPeriod API call: Update an existing grading period.
func (c *Canvas) GradingPeriodsUpdateASingleGradingPeriod(progress *task.Progress, gradingPeriods *time.Time) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if gradingPeriods != nil {
		params["grading_periods"] = *gradingPeriods
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradingPeriodsDeleteAGradingPeriod API call: <b>204 No Content</b> response code is returned if the deletion was
// successful.
func (c *Canvas) GradingPeriodsDeleteAGradingPeriod(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradingStandardsCreateANewGradingStandard API call: Create a new grading standard If grading_scheme_entry arguments
// are omitted, then a default grading scheme will be set. The default scheme is as follows: "A" : 94, "A-" : 90, "B+" :
// 87, "B" : 84, "B-" : 80, "C+" : 77, "C" : 74, "C-" : 70, "D+" : 67, "D" : 64, "D-" : 61, "F" : 0,
func (c *Canvas) GradingStandardsCreateANewGradingStandard(progress *task.Progress, title *string, gradingSchemeEntry *string, courseID string) (*GradingStandard, error) {
	endpoint := fmt.Sprintf("courses/%s/grading_standards", courseID)
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if gradingSchemeEntry != nil {
		params["grading_scheme_entry"] = *gradingSchemeEntry
	}
	responseCtor := func() interface{} {
		return &GradingStandard{}
	}
	var res *GradingStandard
	callback := func(obj interface{}) error {
		res = obj.(*GradingStandard)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradingStandardsListTheGradingStandardsAvailableInAContext API call: Returns the paginated list of grading standards
// for the given context that are visible to the user.
func (c *Canvas) GradingStandardsListTheGradingStandardsAvailableInAContext(progress *task.Progress) ([]GradingStandard, error) {
	endpoint := fmt.Sprintf("courses/1/grading_standards")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]GradingStandard{}
	}
	var res []GradingStandard
	callback := func(obj interface{}) error {
		arr := *obj.(*[]GradingStandard)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GradingStandardsGetASingleGradingStandardInAContext API call: Returns a grading standard for the given context that
// is visible to the user.
func (c *Canvas) GradingStandardsGetASingleGradingStandardInAContext(progress *task.Progress) (*GradingStandard, error) {
	endpoint := fmt.Sprintf("courses/1/grading_standards/5")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &GradingStandard{}
	}
	var res *GradingStandard
	callback := func(obj interface{}) error {
		res = obj.(*GradingStandard)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupCategoriesListGroupCategoriesForAContext API call: Returns a paginated list of group categories in a context
func (c *Canvas) GroupCategoriesListGroupCategoriesForAContext(progress *task.Progress, accountID string) ([]GroupCategory, error) {
	endpoint := fmt.Sprintf("accounts/%s/group_categories", accountID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]GroupCategory{}
	}
	var res []GroupCategory
	callback := func(obj interface{}) error {
		arr := *obj.(*[]GroupCategory)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupCategoriesGetASingleGroupCategory API call: Returns the data for a single group category, or a 401 if the caller
// doesn't have the rights to see it.
func (c *Canvas) GroupCategoriesGetASingleGroupCategory(progress *task.Progress, groupCategoryID string) (*GroupCategory, error) {
	endpoint := fmt.Sprintf("group_categories/%s", groupCategoryID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &GroupCategory{}
	}
	var res *GroupCategory
	callback := func(obj interface{}) error {
		res = obj.(*GroupCategory)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupCategoriesCreateAGroupCategory API call: Create a new group category
func (c *Canvas) GroupCategoriesCreateAGroupCategory(progress *task.Progress, name *string, selfSignup *GroupCategoriesCreateAGroupCategorySelfSignup, autoLeader *GroupCategoriesCreateAGroupCategoryAutoLeader, groupLimit *int, sisGroupCategoryID *string, createGroupCount *int, splitGroupCount *interface{}) (*GroupCategory, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if name != nil {
		params["name"] = *name
	}
	if selfSignup != nil {
		params["self_signup"] = *selfSignup
	}
	if autoLeader != nil {
		params["auto_leader"] = *autoLeader
	}
	if groupLimit != nil {
		params["group_limit"] = *groupLimit
	}
	if sisGroupCategoryID != nil {
		params["sis_group_category_id"] = *sisGroupCategoryID
	}
	if createGroupCount != nil {
		params["create_group_count"] = *createGroupCount
	}
	if splitGroupCount != nil {
		params["split_group_count"] = *splitGroupCount
	}
	responseCtor := func() interface{} {
		return &GroupCategory{}
	}
	var res *GroupCategory
	callback := func(obj interface{}) error {
		res = obj.(*GroupCategory)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupCategoriesUpdateAGroupCategory API call: Modifies an existing group category.
func (c *Canvas) GroupCategoriesUpdateAGroupCategory(progress *task.Progress, name *string, selfSignup *GroupCategoriesUpdateAGroupCategorySelfSignup, autoLeader *GroupCategoriesUpdateAGroupCategoryAutoLeader, groupLimit *int, sisGroupCategoryID *string, createGroupCount *int, splitGroupCount *interface{}, groupCategoryID string) (*GroupCategory, error) {
	endpoint := fmt.Sprintf("group_categories/%s", groupCategoryID)
	params := map[string]interface{}{}
	if name != nil {
		params["name"] = *name
	}
	if selfSignup != nil {
		params["self_signup"] = *selfSignup
	}
	if autoLeader != nil {
		params["auto_leader"] = *autoLeader
	}
	if groupLimit != nil {
		params["group_limit"] = *groupLimit
	}
	if sisGroupCategoryID != nil {
		params["sis_group_category_id"] = *sisGroupCategoryID
	}
	if createGroupCount != nil {
		params["create_group_count"] = *createGroupCount
	}
	if splitGroupCount != nil {
		params["split_group_count"] = *splitGroupCount
	}
	responseCtor := func() interface{} {
		return &GroupCategory{}
	}
	var res *GroupCategory
	callback := func(obj interface{}) error {
		res = obj.(*GroupCategory)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupCategoriesDeleteAGroupCategory API call: Deletes a group category and all groups under it. Protected group
// categories can not be deleted, i.e. "communities" and "student_organized".
func (c *Canvas) GroupCategoriesDeleteAGroupCategory(progress *task.Progress, groupCategoryID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("group_categories/%s", groupCategoryID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupCategoriesListGroupsInGroupCategory API call: Returns a paginated list of groups in a group category
func (c *Canvas) GroupCategoriesListGroupsInGroupCategory(progress *task.Progress, groupCateogryID string) ([]Group, error) {
	endpoint := fmt.Sprintf("group_categories/%s/groups", groupCateogryID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Group{}
	}
	var res []Group
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Group)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupCategoriesListUsersInGroupCategory API call: Returns a paginated list of users in the group category.
func (c *Canvas) GroupCategoriesListUsersInGroupCategory(progress *task.Progress, searchTerm *string, unassigned *bool) ([]User, error) {
	endpoint := fmt.Sprintf("group_categories/1/users")
	params := map[string]interface{}{}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if unassigned != nil {
		params["unassigned"] = *unassigned
	}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupCategoriesAssignUnassignedMembers API call: Assign all unassigned members as evenly as possible among the
// existing student groups.
func (c *Canvas) GroupCategoriesAssignUnassignedMembers(progress *task.Progress, sync *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("group_categories/1/assign_unassigned_members")
	params := map[string]interface{}{}
	if sync != nil {
		params["sync"] = *sync
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsListGroupMemberships API call
func (c *Canvas) GroupsListGroupMemberships(progress *task.Progress, filterStates *GroupsListGroupMembershipsFilterStates, groupID string) ([]GroupMembership, error) {
	endpoint := fmt.Sprintf("groups/%s/memberships", groupID)
	params := map[string]interface{}{}
	if filterStates != nil {
		params["filter_states"] = *filterStates
	}
	responseCtor := func() interface{} {
		return &[]GroupMembership{}
	}
	var res []GroupMembership
	callback := func(obj interface{}) error {
		arr := *obj.(*[]GroupMembership)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsGetASingleGroupMembership API call
func (c *Canvas) GroupsGetASingleGroupMembership(progress *task.Progress, groupID string, userID string) (*GroupMembership, error) {
	endpoint := fmt.Sprintf("groups/%s/users/%s", groupID, userID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &GroupMembership{}
	}
	var res *GroupMembership
	callback := func(obj interface{}) error {
		res = obj.(*GroupMembership)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsCreateAMembership API call
func (c *Canvas) GroupsCreateAMembership(progress *task.Progress, userID *string, groupID string) (*GroupMembership, error) {
	endpoint := fmt.Sprintf("groups/%s/memberships", groupID)
	params := map[string]interface{}{}
	if userID != nil {
		params["user_id"] = *userID
	}
	responseCtor := func() interface{} {
		return &GroupMembership{}
	}
	var res *GroupMembership
	callback := func(obj interface{}) error {
		res = obj.(*GroupMembership)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsUpdateAMembership API call
func (c *Canvas) GroupsUpdateAMembership(progress *task.Progress, workflowState *GroupsUpdateAMembershipWorkflowState, moderator *interface{}, groupID string, userID string) (*GroupMembership, error) {
	endpoint := fmt.Sprintf("groups/%s/users/%s", groupID, userID)
	params := map[string]interface{}{}
	if workflowState != nil {
		params["workflow_state"] = *workflowState
	}
	if moderator != nil {
		params["moderator"] = *moderator
	}
	responseCtor := func() interface{} {
		return &GroupMembership{}
	}
	var res *GroupMembership
	callback := func(obj interface{}) error {
		res = obj.(*GroupMembership)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsLeaveAGroup API call
func (c *Canvas) GroupsLeaveAGroup(progress *task.Progress, groupID string, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("groups/%s/users/%s", groupID, userID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsListYourGroups API call: Returns a paginated list of active groups for the current user.
func (c *Canvas) GroupsListYourGroups(progress *task.Progress, contextType *GroupsListYourGroupsContextType, include *GroupsListYourGroupsInclude) ([]Group, error) {
	endpoint := fmt.Sprintf("users/self/groups")
	params := map[string]interface{}{}
	if contextType != nil {
		params["context_type"] = *contextType
	}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]Group{}
	}
	var res []Group
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Group)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsListTheGroupsAvailableInAContext API call: Returns the paginated list of active groups in the given context
// that are visible to user.
func (c *Canvas) GroupsListTheGroupsAvailableInAContext(progress *task.Progress, onlyOwnGroups *bool, include *GroupsListTheGroupsAvailableInAContextInclude) ([]Group, error) {
	endpoint := fmt.Sprintf("courses/1/groups")
	params := map[string]interface{}{}
	if onlyOwnGroups != nil {
		params["only_own_groups"] = *onlyOwnGroups
	}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]Group{}
	}
	var res []Group
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Group)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsGetASingleGroup API call: Returns the data for a single group, or a 401 if the caller doesn't have the rights
// to see it.
func (c *Canvas) GroupsGetASingleGroup(progress *task.Progress, include *GroupsGetASingleGroupInclude, groupID string) (*Group, error) {
	endpoint := fmt.Sprintf("groups/%s", groupID)
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &Group{}
	}
	var res *Group
	callback := func(obj interface{}) error {
		res = obj.(*Group)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsCreateAGroup API call: Creates a new group. Groups created using the "/api/v1/groups/" endpoint will be
// community groups.
func (c *Canvas) GroupsCreateAGroup(progress *task.Progress, name *string, description *string, isPublic *bool, joinLevel *GroupsCreateAGroupJoinLevel, storageQuotaMb *int, sisGroupID *string) (*Group, error) {
	endpoint := fmt.Sprintf("groups")
	params := map[string]interface{}{}
	if name != nil {
		params["name"] = *name
	}
	if description != nil {
		params["description"] = *description
	}
	if isPublic != nil {
		params["is_public"] = *isPublic
	}
	if joinLevel != nil {
		params["join_level"] = *joinLevel
	}
	if storageQuotaMb != nil {
		params["storage_quota_mb"] = *storageQuotaMb
	}
	if sisGroupID != nil {
		params["sis_group_id"] = *sisGroupID
	}
	responseCtor := func() interface{} {
		return &Group{}
	}
	var res *Group
	callback := func(obj interface{}) error {
		res = obj.(*Group)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsEditAGroup API call: Modifies an existing group.  Note that to set an avatar image for the group, you must
// first upload the image file to the group, and the use the id in the response as the argument to this function.  See
// the {file:file_uploads.html File Upload Documentation} for details on the file upload workflow.
func (c *Canvas) GroupsEditAGroup(progress *task.Progress, name *string, description *string, isPublic *bool, joinLevel *GroupsEditAGroupJoinLevel, avatarID *int, storageQuotaMb *int, members *string, sisGroupID *string, groupID string) (*Group, error) {
	endpoint := fmt.Sprintf("groups/%s", groupID)
	params := map[string]interface{}{}
	if name != nil {
		params["name"] = *name
	}
	if description != nil {
		params["description"] = *description
	}
	if isPublic != nil {
		params["is_public"] = *isPublic
	}
	if joinLevel != nil {
		params["join_level"] = *joinLevel
	}
	if avatarID != nil {
		params["avatar_id"] = *avatarID
	}
	if storageQuotaMb != nil {
		params["storage_quota_mb"] = *storageQuotaMb
	}
	if members != nil {
		params["members"] = *members
	}
	if sisGroupID != nil {
		params["sis_group_id"] = *sisGroupID
	}
	responseCtor := func() interface{} {
		return &Group{}
	}
	var res *Group
	callback := func(obj interface{}) error {
		res = obj.(*Group)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsDeleteAGroup API call: Deletes a group and removes all members.
func (c *Canvas) GroupsDeleteAGroup(progress *task.Progress, groupID string) (*Group, error) {
	endpoint := fmt.Sprintf("groups/%s", groupID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Group{}
	}
	var res *Group
	callback := func(obj interface{}) error {
		res = obj.(*Group)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsInviteOthersToAGroup API call
func (c *Canvas) GroupsInviteOthersToAGroup(progress *task.Progress, invitees *string, groupID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("groups/%s/invite", groupID)
	params := map[string]interface{}{}
	if invitees != nil {
		params["invitees"] = *invitees
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsListGroupSUsers API call: Returns a paginated list of users in the group.
func (c *Canvas) GroupsListGroupSUsers(progress *task.Progress, searchTerm *string, include *GroupsListGroupSUsersInclude, excludeInactive *bool) ([]User, error) {
	endpoint := fmt.Sprintf("groups/1/users")
	params := map[string]interface{}{}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if include != nil {
		params["include"] = *include
	}
	if excludeInactive != nil {
		params["exclude_inactive"] = *excludeInactive
	}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsUploadAFile API call: Upload a file to the group. This API endpoint is the first step in uploading a file to a
// group. See the {file:file_uploads.html File Upload Documentation} for details on the file upload workflow. Only those
// with the "Manage Files" permission on a group can upload files to the group. By default, this is anybody
// participating in the group, or any admin over the group.
func (c *Canvas) GroupsUploadAFile(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsPreviewProcessedHTML API call: Preview html content processed for this group
func (c *Canvas) GroupsPreviewProcessedHTML(progress *task.Progress, html *string, groupID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("groups/%s/preview_html", groupID)
	params := map[string]interface{}{}
	if html != nil {
		params["html"] = *html
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsGroupActivityStream API call: Returns the current user's group-specific activity stream, paginated. For full
// documentation, see the API documentation for the user activity stream, in the user api.
func (c *Canvas) GroupsGroupActivityStream(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsGroupActivityStreamSummary API call: Returns a summary of the current user's group-specific activity stream.
// For full documentation, see the API documentation for the user activity stream summary, in the user api.
func (c *Canvas) GroupsGroupActivityStreamSummary(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// GroupsPermissions API call: Returns permission information for the calling user in the given group. See also the
// {api:AccountsController#permissions Account} and {api:CoursesController#permissions Course} counterparts.
func (c *Canvas) GroupsPermissions(progress *task.Progress, permissions *string, groupID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("groups/%s/permissions", groupID)
	params := map[string]interface{}{}
	if permissions != nil {
		params["permissions"] = *permissions
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ImageSearchFindImages API call: Find public domain images for use in courses and user content.  If you select an
// image using this API, please use the {api:InternetImageController#image_selection Confirm image selection API} to
// indicate photo usage to the server.
func (c *Canvas) ImageSearchFindImages(progress *task.Progress, query *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if query != nil {
		params["query"] = *query
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ImageSearchConfirmImageSelection API call: After you have used the search API, you should hit this API to indicate
// photo usage to the server.
func (c *Canvas) ImageSearchConfirmImageSelection(progress *task.Progress, id *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// JWTsCreateJWT API call: Create a unique jwt for using with other canvas services Generates a different JWT each time
// it's called, each one expires after a short window (1 hour)
func (c *Canvas) JWTsCreateJWT(progress *task.Progress) (*JWT, error) {
	endpoint := fmt.Sprintf("jwts")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &JWT{}
	}
	var res *JWT
	callback := func(obj interface{}) error {
		res = obj.(*JWT)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// JWTsRefreshJWT API call: Refresh a JWT for use with other canvas services Generates a different JWT each time it's
// called, each one expires after a short window (1 hour).
func (c *Canvas) JWTsRefreshJWT(progress *task.Progress, jwt *string) (*JWT, error) {
	endpoint := fmt.Sprintf("jwts/refresh")
	params := map[string]interface{}{}
	if jwt != nil {
		params["jwt"] = *jwt
	}
	responseCtor := func() interface{} {
		return &JWT{}
	}
	var res *JWT
	callback := func(obj interface{}) error {
		res = obj.(*JWT)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// LatePolicyGetALatePolicy API call: Returns the late policy for a course.
func (c *Canvas) LatePolicyGetALatePolicy(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// LatePolicyCreateALatePolicy API call: Create a late policy. If the course already has a late policy, a bad_request is
// returned since there can only be one late policy per course.
func (c *Canvas) LatePolicyCreateALatePolicy(progress *task.Progress, latePolicy *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if latePolicy != nil {
		params["late_policy"] = *latePolicy
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// LatePolicyPatchALatePolicy API call: Patch a late policy. No body is returned upon success.
func (c *Canvas) LatePolicyPatchALatePolicy(progress *task.Progress, latePolicy *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if latePolicy != nil {
		params["late_policy"] = *latePolicy
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// MediaObjectsListMediaObjects API call: Returns Media Objects created by the user making the request. When using the
// second version, returns only those Media Objects associated with the given course.
func (c *Canvas) MediaObjectsListMediaObjects(progress *task.Progress, sort *MediaObjectsListMediaObjectsSort, order *MediaObjectsListMediaObjectsOrder, exclude *MediaObjectsListMediaObjectsExclude) ([]map[string]interface{}, error) {
	endpoint := fmt.Sprintf("media_objects")
	params := map[string]interface{}{}
	if sort != nil {
		params["sort"] = *sort
	}
	if order != nil {
		params["order"] = *order
	}
	if exclude != nil {
		params["exclude"] = *exclude
	}
	responseCtor := func() interface{} {
		return &[]map[string]interface{}{}
	}
	var res []map[string]interface{}
	callback := func(obj interface{}) error {
		arr := *obj.(*[]map[string]interface{})
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// MediaObjectsUpdateMediaObject API call
func (c *Canvas) MediaObjectsUpdateMediaObject(progress *task.Progress, userEnteredTitle *interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if userEnteredTitle != nil {
		params["user_entered_title"] = *userEnteredTitle
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// MediaObjectsListMediaTracksForAMediaObject API call: List the media tracks associated with a media object
func (c *Canvas) MediaObjectsListMediaTracksForAMediaObject(progress *task.Progress, include *MediaObjectsListMediaTracksForAMediaObjectInclude, mediaObjectID string) ([]map[string]interface{}, error) {
	endpoint := fmt.Sprintf("media_objects/%s/media_tracks", mediaObjectID)
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]map[string]interface{}{}
	}
	var res []map[string]interface{}
	callback := func(obj interface{}) error {
		arr := *obj.(*[]map[string]interface{})
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// MediaObjectsUpdateMediaTracks API call: Replace the media tracks associated with a media object with the array of
// tracks provided in the body. Update will delete any existing tracks not listed, leave untouched any tracks with no
// content field, and update or create tracks with a content field.
func (c *Canvas) MediaObjectsUpdateMediaTracks(progress *task.Progress, include *string, mediaObjectID string) ([]map[string]interface{}, error) {
	endpoint := fmt.Sprintf("media_objects/%s/mediatracksinclude[]=content", mediaObjectID)
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]map[string]interface{}{}
	}
	var res []map[string]interface{}
	callback := func(obj interface{}) error {
		arr := *obj.(*[]map[string]interface{})
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentMigrationsListMigrationIssues API call: Returns paginated migration issues
func (c *Canvas) ContentMigrationsListMigrationIssues(progress *task.Progress, courseID string, contentMigrationID string) ([]MigrationIssue, error) {
	endpoint := fmt.Sprintf("courses/%s/content_migrations/%s/migration_issues", courseID, contentMigrationID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]MigrationIssue{}
	}
	var res []MigrationIssue
	callback := func(obj interface{}) error {
		arr := *obj.(*[]MigrationIssue)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentMigrationsGetAMigrationIssue API call: Returns data on an individual migration issue
func (c *Canvas) ContentMigrationsGetAMigrationIssue(progress *task.Progress, courseID string, contentMigrationID string, id string) (*MigrationIssue, error) {
	endpoint := fmt.Sprintf("courses/%s/content_migrations/%s/migration_issues/%s", courseID, contentMigrationID, id)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &MigrationIssue{}
	}
	var res *MigrationIssue
	callback := func(obj interface{}) error {
		res = obj.(*MigrationIssue)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ContentMigrationsUpdateAMigrationIssue API call: Update the workflow_state of a migration issue
func (c *Canvas) ContentMigrationsUpdateAMigrationIssue(progress *task.Progress, workflowState *ContentMigrationsUpdateAMigrationIssueWorkflowState, courseID string, contentMigrationID string, id string) (*MigrationIssue, error) {
	endpoint := fmt.Sprintf("courses/%s/content_migrations/%s/migration_issues/%s", courseID, contentMigrationID, id)
	params := map[string]interface{}{}
	if workflowState != nil {
		params["workflow_state"] = *workflowState
	}
	responseCtor := func() interface{} {
		return &MigrationIssue{}
	}
	var res *MigrationIssue
	callback := func(obj interface{}) error {
		res = obj.(*MigrationIssue)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModeratedGradingListStudentsSelectedForModeration API call: Returns a paginated list of students selected for
// moderation
func (c *Canvas) ModeratedGradingListStudentsSelectedForModeration(progress *task.Progress) ([]User, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModeratedGradingSelectStudentsForModeration API call: Returns an array of users that were selected for moderation
func (c *Canvas) ModeratedGradingSelectStudentsForModeration(progress *task.Progress, studentIds *float64) ([]User, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if studentIds != nil {
		params["student_ids"] = *studentIds
	}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// NotificationPreferencesListPreferences API call: Fetch all preferences for the given communication channel
func (c *Canvas) NotificationPreferencesListPreferences(progress *task.Progress) ([]NotificationPreference, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]NotificationPreference{}
	}
	var res []NotificationPreference
	callback := func(obj interface{}) error {
		arr := *obj.(*[]NotificationPreference)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// NotificationPreferencesListOfPreferenceCategories API call: Fetch all notification preference categories for the
// given communication channel
func (c *Canvas) NotificationPreferencesListOfPreferenceCategories(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// NotificationPreferencesGetAPreference API call: Fetch the preference for the given notification for the given
// communication channel
func (c *Canvas) NotificationPreferencesGetAPreference(progress *task.Progress) (*NotificationPreference, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &NotificationPreference{}
	}
	var res *NotificationPreference
	callback := func(obj interface{}) error {
		res = obj.(*NotificationPreference)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// NotificationPreferencesUpdateAPreference API call: Change the preference for a single notification for a single
// communication channel
func (c *Canvas) NotificationPreferencesUpdateAPreference(progress *task.Progress, notificationPreferences *interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if notificationPreferences != nil {
		params["notification_preferences"] = *notificationPreferences
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// NotificationPreferencesUpdatePreferencesByCategory API call: Change the preferences for multiple notifications based
// on the category for a single communication channel
func (c *Canvas) NotificationPreferencesUpdatePreferencesByCategory(progress *task.Progress, category *interface{}, notificationPreferences *interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if category != nil {
		params["category"] = *category
	}
	if notificationPreferences != nil {
		params["notification_preferences"] = *notificationPreferences
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// NotificationPreferencesUpdateMultiplePreferences API call: Change the preferences for multiple notifications for a
// single communication channel at once
func (c *Canvas) NotificationPreferencesUpdateMultiplePreferences(progress *task.Progress, notificationPreferences *interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if notificationPreferences != nil {
		params["notification_preferences"] = *notificationPreferences
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsRedirectToRootOutcomeGroupForContext API call: Convenience redirect to find the root outcome group for a
// particular context. Will redirect to the appropriate outcome group's URL.
func (c *Canvas) OutcomeGroupsRedirectToRootOutcomeGroupForContext(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsGetAllOutcomeGroupsForContext API call
func (c *Canvas) OutcomeGroupsGetAllOutcomeGroupsForContext(progress *task.Progress) ([]OutcomeGroup, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]OutcomeGroup{}
	}
	var res []OutcomeGroup
	callback := func(obj interface{}) error {
		arr := *obj.(*[]OutcomeGroup)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsGetAllOutcomeLinksForContext API call
func (c *Canvas) OutcomeGroupsGetAllOutcomeLinksForContext(progress *task.Progress, outcomeStyle *string, outcomeGroupStyle *string) ([]OutcomeLink, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if outcomeStyle != nil {
		params["outcome_style"] = *outcomeStyle
	}
	if outcomeGroupStyle != nil {
		params["outcome_group_style"] = *outcomeGroupStyle
	}
	responseCtor := func() interface{} {
		return &[]OutcomeLink{}
	}
	var res []OutcomeLink
	callback := func(obj interface{}) error {
		arr := *obj.(*[]OutcomeLink)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsShowAnOutcomeGroup API call
func (c *Canvas) OutcomeGroupsShowAnOutcomeGroup(progress *task.Progress) (*OutcomeGroup, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &OutcomeGroup{}
	}
	var res *OutcomeGroup
	callback := func(obj interface{}) error {
		res = obj.(*OutcomeGroup)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsUpdateAnOutcomeGroup API call: Modify an existing outcome group. Fields not provided are left as is;
// unrecognized fields are ignored. When changing the parent outcome group, the new parent group must belong to the same
// context as this outcome group, and must not be a descendant of this outcome group (i.e. no cycles allowed).
func (c *Canvas) OutcomeGroupsUpdateAnOutcomeGroup(progress *task.Progress, title *string, description *string, vendorGUID *string, parentOutcomeGroupID *int) (*OutcomeGroup, error) {
	endpoint := fmt.Sprintf("accounts/1/outcome_groups/2.json")
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if description != nil {
		params["description"] = *description
	}
	if vendorGUID != nil {
		params["vendor_guid"] = *vendorGUID
	}
	if parentOutcomeGroupID != nil {
		params["parent_outcome_group_id"] = *parentOutcomeGroupID
	}
	responseCtor := func() interface{} {
		return &OutcomeGroup{}
	}
	var res *OutcomeGroup
	callback := func(obj interface{}) error {
		res = obj.(*OutcomeGroup)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsDeleteAnOutcomeGroup API call: Deleting an outcome group deletes descendant outcome groups and outcome
// links. The linked outcomes themselves are only deleted if all links to the outcome were deleted. Aligned outcomes
// cannot be deleted; as such, if all remaining links to an aligned outcome are included in this group's descendants,
// the group deletion will fail.
func (c *Canvas) OutcomeGroupsDeleteAnOutcomeGroup(progress *task.Progress) (*OutcomeGroup, error) {
	endpoint := fmt.Sprintf("accounts/1/outcome_groups/2.json")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &OutcomeGroup{}
	}
	var res *OutcomeGroup
	callback := func(obj interface{}) error {
		res = obj.(*OutcomeGroup)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsListLinkedOutcomes API call: A paginated list of the immediate OutcomeLink children of the outcome
// group.
func (c *Canvas) OutcomeGroupsListLinkedOutcomes(progress *task.Progress, outcomeStyle *string) ([]OutcomeLink, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if outcomeStyle != nil {
		params["outcome_style"] = *outcomeStyle
	}
	responseCtor := func() interface{} {
		return &[]OutcomeLink{}
	}
	var res []OutcomeLink
	callback := func(obj interface{}) error {
		arr := *obj.(*[]OutcomeLink)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsCreateLinkAnOutcome API call: Link an outcome into the outcome group. The outcome to link can either be
// specified by a PUT to the link URL for a specific outcome (the outcome_id in the PUT URLs) or by supplying the
// information for a new outcome (title, description, ratings, mastery_points) in a POST to the collection. If linking
// an existing outcome, the outcome_id must identify an outcome available to this context; i.e. an outcome owned by this
// group's context, an outcome owned by an associated account, or a global outcome. With outcome_id present, any other
// parameters (except move_from) are ignored. If defining a new outcome, the outcome is created in the outcome group's
// context using the provided title, description, ratings, and mastery points; the title is required but all other
// fields are optional. The new outcome is then linked into the outcome group. If ratings are provided when creating a
// new outcome, an embedded rubric criterion is included in the new outcome. This criterion's mastery_points default to
// the maximum points in the highest rating if not specified in the mastery_points parameter. Any ratings lacking a
// description are given a default of "No description". Any ratings lacking a point value are given a default of 0. If
// no ratings are provided, the mastery_points parameter is ignored.
func (c *Canvas) OutcomeGroupsCreateLinkAnOutcome(progress *task.Progress, outcomeID *int, moveFrom *int, title *string, displayName *string, description *string, vendorGUID *string, masteryPoints *int, ratings *string, calculationMethod *OutcomeGroupsCreateLinkAnOutcomeCalculationMethod, calculationInt *int) (*OutcomeLink, error) {
	endpoint := fmt.Sprintf("accounts/1/outcome_groups/1/outcomes.json")
	params := map[string]interface{}{}
	if outcomeID != nil {
		params["outcome_id"] = *outcomeID
	}
	if moveFrom != nil {
		params["move_from"] = *moveFrom
	}
	if title != nil {
		params["title"] = *title
	}
	if displayName != nil {
		params["display_name"] = *displayName
	}
	if description != nil {
		params["description"] = *description
	}
	if vendorGUID != nil {
		params["vendor_guid"] = *vendorGUID
	}
	if masteryPoints != nil {
		params["mastery_points"] = *masteryPoints
	}
	if ratings != nil {
		params["ratings"] = *ratings
	}
	if calculationMethod != nil {
		params["calculation_method"] = *calculationMethod
	}
	if calculationInt != nil {
		params["calculation_int"] = *calculationInt
	}
	responseCtor := func() interface{} {
		return &OutcomeLink{}
	}
	var res *OutcomeLink
	callback := func(obj interface{}) error {
		res = obj.(*OutcomeLink)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsUnlinkAnOutcome API call: Unlinking an outcome only deletes the outcome itself if this was the last link
// to the outcome in any group in any context. Aligned outcomes cannot be deleted; as such, if this is the last link to
// an aligned outcome, the unlinking will fail.
func (c *Canvas) OutcomeGroupsUnlinkAnOutcome(progress *task.Progress) (*OutcomeLink, error) {
	endpoint := fmt.Sprintf("accounts/1/outcome_groups/1/outcomes/1.json")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &OutcomeLink{}
	}
	var res *OutcomeLink
	callback := func(obj interface{}) error {
		res = obj.(*OutcomeLink)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsListSubgroups API call: A paginated list of the immediate OutcomeGroup children of the outcome group.
func (c *Canvas) OutcomeGroupsListSubgroups(progress *task.Progress) ([]OutcomeGroup, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]OutcomeGroup{}
	}
	var res []OutcomeGroup
	callback := func(obj interface{}) error {
		arr := *obj.(*[]OutcomeGroup)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsCreateASubgroup API call: Creates a new empty subgroup under the outcome group with the given title and
// description.
func (c *Canvas) OutcomeGroupsCreateASubgroup(progress *task.Progress, title *string, description *string, vendorGUID *string) (*OutcomeGroup, error) {
	endpoint := fmt.Sprintf("accounts/1/outcome_groups/1/subgroups.json")
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if description != nil {
		params["description"] = *description
	}
	if vendorGUID != nil {
		params["vendor_guid"] = *vendorGUID
	}
	responseCtor := func() interface{} {
		return &OutcomeGroup{}
	}
	var res *OutcomeGroup
	callback := func(obj interface{}) error {
		res = obj.(*OutcomeGroup)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeGroupsImportAnOutcomeGroup API call: Creates a new subgroup of the outcome group with the same title and
// description as the source group, then creates links in that new subgroup to the same outcomes that are linked in the
// source group. Recurses on the subgroups of the source group, importing them each in turn into the new subgroup.
// Allows you to copy organizational structure, but does not create copies of the outcomes themselves, only new links.
// The source group must be either global, from the same context as this outcome group, or from an associated account.
// The source group cannot be the root outcome group of its context.
func (c *Canvas) OutcomeGroupsImportAnOutcomeGroup(progress *task.Progress, sourceOutcomeGroupID *int, async *bool) (*OutcomeGroup, error) {
	endpoint := fmt.Sprintf("accounts/2/outcome_groups/3/import.json")
	params := map[string]interface{}{}
	if sourceOutcomeGroupID != nil {
		params["source_outcome_group_id"] = *sourceOutcomeGroupID
	}
	if async != nil {
		params["async"] = *async
	}
	responseCtor := func() interface{} {
		return &OutcomeGroup{}
	}
	var res *OutcomeGroup
	callback := func(obj interface{}) error {
		res = obj.(*OutcomeGroup)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeImportsImportOutcomes API call: Import outcomes into Canvas. For more information on the format that's
// expected here, please see the "Outcomes CSV" section in the API docs.
func (c *Canvas) OutcomeImportsImportOutcomes(progress *task.Progress, importType *string, attachment *interface{}, extension *string) (*OutcomeImport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if importType != nil {
		params["import_type"] = *importType
	}
	if attachment != nil {
		params["attachment"] = *attachment
	}
	if extension != nil {
		params["extension"] = *extension
	}
	responseCtor := func() interface{} {
		return &OutcomeImport{}
	}
	var res *OutcomeImport
	callback := func(obj interface{}) error {
		res = obj.(*OutcomeImport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeImportsGetOutcomeImportStatus API call: Get the status of an already created Outcome import. Pass 'latest' for
// the outcome import id for the latest import. Examples: curl
// 'https://<canvas>/api/v1/accounts/<account_id>/outcome_imports/<outcome_import_id>' \ -H "Authorization: Bearer
// <token>" curl 'https://<canvas>/api/v1/courses/<course_id>/outcome_imports/<outcome_import_id>' \ -H "Authorization:
// Bearer <token>"
func (c *Canvas) OutcomeImportsGetOutcomeImportStatus(progress *task.Progress) (*OutcomeImport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &OutcomeImport{}
	}
	var res *OutcomeImport
	callback := func(obj interface{}) error {
		res = obj.(*OutcomeImport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ProficiencyRatingsCreateUpdateProficiencyRatings API call: Create or update account-level proficiency ratings. These
// ratings will apply to all sub-accounts, unless they have their own account-level proficiency ratings defined.
func (c *Canvas) ProficiencyRatingsCreateUpdateProficiencyRatings(progress *task.Progress, ratings *string, accountID string) (*Proficiency, error) {
	endpoint := fmt.Sprintf("accounts/%s/outcome_proficiency", accountID)
	params := map[string]interface{}{}
	if ratings != nil {
		params["ratings"] = *ratings
	}
	responseCtor := func() interface{} {
		return &Proficiency{}
	}
	var res *Proficiency
	callback := func(obj interface{}) error {
		res = obj.(*Proficiency)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ProficiencyRatingsGetProficiencyRatings API call: Get account-level proficiency ratings. If not defined for this
// account, it will return proficiency ratings for the nearest super-account with ratings defined. Will return 404 if
// none found. Examples: curl https://<canvas>/api/v1/accounts/<account_id>/outcome_proficiency \ -H 'Authorization:
// Bearer <token>'
func (c *Canvas) ProficiencyRatingsGetProficiencyRatings(progress *task.Progress) (*Proficiency, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Proficiency{}
	}
	var res *Proficiency
	callback := func(obj interface{}) error {
		res = obj.(*Proficiency)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeResultsGetOutcomeResults API call: Gets the outcome results for users and outcomes in the specified context.
func (c *Canvas) OutcomeResultsGetOutcomeResults(progress *task.Progress, userIds *int, outcomeIds *int, include *OutcomeResultsGetOutcomeResultsInclude, includeHidden *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if userIds != nil {
		params["user_ids"] = *userIds
	}
	if outcomeIds != nil {
		params["outcome_ids"] = *outcomeIds
	}
	if include != nil {
		params["include"] = *include
	}
	if includeHidden != nil {
		params["include_hidden"] = *includeHidden
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomeResultsGetOutcomeResultRollups API call: Gets the outcome rollups for the users and outcomes in the specified
// context.
func (c *Canvas) OutcomeResultsGetOutcomeResultRollups(progress *task.Progress, aggregate *OutcomeResultsGetOutcomeResultRollupsAggregate, aggregateStat *OutcomeResultsGetOutcomeResultRollupsAggregateStat, userIds *int, outcomeIds *int, include *OutcomeResultsGetOutcomeResultRollupsInclude, exclude *OutcomeResultsGetOutcomeResultRollupsExclude, sortBy *OutcomeResultsGetOutcomeResultRollupsSortBy, sortOutcomeID *int, sortOrder *OutcomeResultsGetOutcomeResultRollupsSortOrder) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if aggregate != nil {
		params["aggregate"] = *aggregate
	}
	if aggregateStat != nil {
		params["aggregate_stat"] = *aggregateStat
	}
	if userIds != nil {
		params["user_ids"] = *userIds
	}
	if outcomeIds != nil {
		params["outcome_ids"] = *outcomeIds
	}
	if include != nil {
		params["include"] = *include
	}
	if exclude != nil {
		params["exclude"] = *exclude
	}
	if sortBy != nil {
		params["sort_by"] = *sortBy
	}
	if sortOutcomeID != nil {
		params["sort_outcome_id"] = *sortOutcomeID
	}
	if sortOrder != nil {
		params["sort_order"] = *sortOrder
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomesShowAnOutcome API call: Returns the details of the outcome with the given id.
func (c *Canvas) OutcomesShowAnOutcome(progress *task.Progress) (*Outcome, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Outcome{}
	}
	var res *Outcome
	callback := func(obj interface{}) error {
		res = obj.(*Outcome)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomesUpdateAnOutcome API call: Modify an existing outcome. Fields not provided are left as is; unrecognized fields
// are ignored. If any new ratings are provided, the combination of all new ratings provided completely replace any
// existing embedded rubric criterion; it is not possible to tweak the ratings of the embedded rubric criterion. A new
// embedded rubric criterion's mastery_points default to the maximum points in the highest rating if not specified in
// the mastery_points parameter. Any new ratings lacking a description are given a default of "No description". Any new
// ratings lacking a point value are given a default of 0.
func (c *Canvas) OutcomesUpdateAnOutcome(progress *task.Progress, title *string, displayName *string, description *string, vendorGUID *string, masteryPoints *int, ratings *string, calculationMethod *OutcomesUpdateAnOutcomeCalculationMethod, calculationInt *int) (*Outcome, error) {
	endpoint := fmt.Sprintf("outcomes/1.json")
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if displayName != nil {
		params["display_name"] = *displayName
	}
	if description != nil {
		params["description"] = *description
	}
	if vendorGUID != nil {
		params["vendor_guid"] = *vendorGUID
	}
	if masteryPoints != nil {
		params["mastery_points"] = *masteryPoints
	}
	if ratings != nil {
		params["ratings"] = *ratings
	}
	if calculationMethod != nil {
		params["calculation_method"] = *calculationMethod
	}
	if calculationInt != nil {
		params["calculation_int"] = *calculationInt
	}
	responseCtor := func() interface{} {
		return &Outcome{}
	}
	var res *Outcome
	callback := func(obj interface{}) error {
		res = obj.(*Outcome)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// OutcomesGetAlignedAssignmentsForAnOutcomeInACourseForAParticularStudent API call
func (c *Canvas) OutcomesGetAlignedAssignmentsForAnOutcomeInACourseForAParticularStudent(progress *task.Progress, courseID *int, studentID *int) ([]OutcomeAlignment, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	if studentID != nil {
		params["student_id"] = *studentID
	}
	responseCtor := func() interface{} {
		return &[]OutcomeAlignment{}
	}
	var res []OutcomeAlignment
	callback := func(obj interface{}) error {
		arr := *obj.(*[]OutcomeAlignment)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersListUserPageViews API call: Return a paginated list of the user's page view history in json format, similar to
// the available CSV download. Page views are returned in descending order, newest to oldest.
func (c *Canvas) UsersListUserPageViews(progress *task.Progress, startTime *time.Time, endTime *time.Time) ([]PageView, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startTime != nil {
		params["start_time"] = *startTime
	}
	if endTime != nil {
		params["end_time"] = *endTime
	}
	responseCtor := func() interface{} {
		return &[]PageView{}
	}
	var res []PageView
	callback := func(obj interface{}) error {
		arr := *obj.(*[]PageView)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PeerReviewsGetAllPeerReviews API call: Get a list of all Peer Reviews for this assignment
func (c *Canvas) PeerReviewsGetAllPeerReviews(progress *task.Progress, include *PeerReviewsGetAllPeerReviewsInclude) ([]PeerReview, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]PeerReview{}
	}
	var res []PeerReview
	callback := func(obj interface{}) error {
		arr := *obj.(*[]PeerReview)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PeerReviewsCreatePeerReview API call: Create a peer review for the assignment
func (c *Canvas) PeerReviewsCreatePeerReview(progress *task.Progress, userID *int) (*PeerReview, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if userID != nil {
		params["user_id"] = *userID
	}
	responseCtor := func() interface{} {
		return &PeerReview{}
	}
	var res *PeerReview
	callback := func(obj interface{}) error {
		res = obj.(*PeerReview)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PeerReviewsDeletePeerReview API call: Delete a peer review for the assignment
func (c *Canvas) PeerReviewsDeletePeerReview(progress *task.Progress, userID *int) (*PeerReview, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if userID != nil {
		params["user_id"] = *userID
	}
	responseCtor := func() interface{} {
		return &PeerReview{}
	}
	var res *PeerReview
	callback := func(obj interface{}) error {
		res = obj.(*PeerReview)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerListPlannerItems API call: Retrieve the paginated list of objects to be shown on the planner for the current
// user with the associated planner override to override an item's visibility if set. Planner items for a student may
// also be retrieved by a linked observer. Use the path that accepts a user_id and supply the student's id.
func (c *Canvas) PlannerListPlannerItems(progress *task.Progress, startDate *time.Time, endDate *time.Time, contextCodes *string, filter *PlannerListPlannerItemsFilter) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startDate != nil {
		params["start_date"] = *startDate
	}
	if endDate != nil {
		params["end_date"] = *endDate
	}
	if contextCodes != nil {
		params["context_codes"] = *contextCodes
	}
	if filter != nil {
		params["filter"] = *filter
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerListPlannerNotes API call: Retrieve planner note for a user
func (c *Canvas) PlannerListPlannerNotes(progress *task.Progress, startDate *time.Time, endDate *time.Time, contextCodes *string) ([]PlannerNote, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if startDate != nil {
		params["start_date"] = *startDate
	}
	if endDate != nil {
		params["end_date"] = *endDate
	}
	if contextCodes != nil {
		params["context_codes"] = *contextCodes
	}
	responseCtor := func() interface{} {
		return &[]PlannerNote{}
	}
	var res []PlannerNote
	callback := func(obj interface{}) error {
		arr := *obj.(*[]PlannerNote)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerShowAPlannerNote API call: Retrieve a planner note for the current user
func (c *Canvas) PlannerShowAPlannerNote(progress *task.Progress) (*PlannerNote, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &PlannerNote{}
	}
	var res *PlannerNote
	callback := func(obj interface{}) error {
		res = obj.(*PlannerNote)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerUpdateAPlannerNote API call: Update a planner note for the current user
func (c *Canvas) PlannerUpdateAPlannerNote(progress *task.Progress, title *string, details *string, todoDate *time.Time, courseID *int) (*PlannerNote, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if details != nil {
		params["details"] = *details
	}
	if todoDate != nil {
		params["todo_date"] = *todoDate
	}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	responseCtor := func() interface{} {
		return &PlannerNote{}
	}
	var res *PlannerNote
	callback := func(obj interface{}) error {
		res = obj.(*PlannerNote)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerCreateAPlannerNote API call: Create a planner note for the current user
func (c *Canvas) PlannerCreateAPlannerNote(progress *task.Progress, title *string, details *string, todoDate *time.Time, courseID *int, linkedObjectType *string, linkedObjectID *int) (*PlannerNote, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if details != nil {
		params["details"] = *details
	}
	if todoDate != nil {
		params["todo_date"] = *todoDate
	}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	if linkedObjectType != nil {
		params["linked_object_type"] = *linkedObjectType
	}
	if linkedObjectID != nil {
		params["linked_object_id"] = *linkedObjectID
	}
	responseCtor := func() interface{} {
		return &PlannerNote{}
	}
	var res *PlannerNote
	callback := func(obj interface{}) error {
		res = obj.(*PlannerNote)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerDeleteAPlannerNote API call: Delete a planner note for the current user
func (c *Canvas) PlannerDeleteAPlannerNote(progress *task.Progress) (*PlannerNote, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &PlannerNote{}
	}
	var res *PlannerNote
	callback := func(obj interface{}) error {
		res = obj.(*PlannerNote)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerListPlannerOverrides API call: Retrieve a planner override for the current user
func (c *Canvas) PlannerListPlannerOverrides(progress *task.Progress) ([]PlannerOverride, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]PlannerOverride{}
	}
	var res []PlannerOverride
	callback := func(obj interface{}) error {
		arr := *obj.(*[]PlannerOverride)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerShowAPlannerOverride API call: Retrieve a planner override for the current user
func (c *Canvas) PlannerShowAPlannerOverride(progress *task.Progress) (*PlannerOverride, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &PlannerOverride{}
	}
	var res *PlannerOverride
	callback := func(obj interface{}) error {
		res = obj.(*PlannerOverride)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerUpdateAPlannerOverride API call: Update a planner override's visibilty for the current user
func (c *Canvas) PlannerUpdateAPlannerOverride(progress *task.Progress, markedComplete *interface{}, dismissed *interface{}) (*PlannerOverride, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if markedComplete != nil {
		params["marked_complete"] = *markedComplete
	}
	if dismissed != nil {
		params["dismissed"] = *dismissed
	}
	responseCtor := func() interface{} {
		return &PlannerOverride{}
	}
	var res *PlannerOverride
	callback := func(obj interface{}) error {
		res = obj.(*PlannerOverride)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerCreateAPlannerOverride API call: Create a planner override for the current user
func (c *Canvas) PlannerCreateAPlannerOverride(progress *task.Progress, plannableType *PlannerCreateAPlannerOverridePlannableType, plannableID *int, markedComplete *bool, dismissed *bool) (*PlannerOverride, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if plannableType != nil {
		params["plannable_type"] = *plannableType
	}
	if plannableID != nil {
		params["plannable_id"] = *plannableID
	}
	if markedComplete != nil {
		params["marked_complete"] = *markedComplete
	}
	if dismissed != nil {
		params["dismissed"] = *dismissed
	}
	responseCtor := func() interface{} {
		return &PlannerOverride{}
	}
	var res *PlannerOverride
	callback := func(obj interface{}) error {
		res = obj.(*PlannerOverride)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PlannerDeleteAPlannerOverride API call: Delete a planner override for the current user
func (c *Canvas) PlannerDeleteAPlannerOverride(progress *task.Progress) (*PlannerOverride, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &PlannerOverride{}
	}
	var res *PlannerOverride
	callback := func(obj interface{}) error {
		res = obj.(*PlannerOverride)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersGetUserProfile API call: Returns user profile data, including user id, name, and profile pic. When requesting
// the profile for the user accessing the API, the user's calendar feed URL and LTI user id will be returned as well.
func (c *Canvas) UsersGetUserProfile(progress *task.Progress) (*Profile, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Profile{}
	}
	var res *Profile
	callback := func(obj interface{}) error {
		res = obj.(*Profile)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersListAvatarOptions API call: A paginated list of the possible user avatar options that can be set with the user
// update endpoint. The response will be an array of avatar records. If the 'type' field is 'attachment', the record
// will include all the normal attachment json fields; otherwise it will include only the 'url' and 'display_name'
// fields. Additionally, all records will include a 'type' field and a 'token' field. The following explains each field
// in more detail type:: ["gravatar"|"attachment"|"no_pic"] The type of avatar record, for categorization purposes.
// url:: The url of the avatar token:: A unique representation of the avatar record which can be used to set the avatar
// with the user update endpoint. Note: this is an internal representation and is subject to change without notice. It
// should be consumed with this api endpoint and used in the user update endpoint, and should not be constructed by the
// client. display_name:: A textual description of the avatar record id:: ['attachment' type only] the internal id of
// the attachment content-type:: ['attachment' type only] the content-type of the attachment filename:: ['attachment'
// type only] the filename of the attachment size:: ['attachment' type only] the size of the attachment
func (c *Canvas) UsersListAvatarOptions(progress *task.Progress) ([]Avatar, error) {
	endpoint := fmt.Sprintf("users/1/avatars.json")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]Avatar{}
	}
	var res []Avatar
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Avatar)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ProgressQueryProgress API call: Return completion and status information about an asynchronous job
func (c *Canvas) ProgressQueryProgress(progress *task.Progress) (*Progress, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Progress{}
	}
	var res *Progress
	callback := func(obj interface{}) error {
		res = obj.(*Progress)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModeratedGradingBulkSelectProvisionalGrades API call: Choose which provisional grades will be received by associated
// students for an assignment. The caller must be the final grader for the assignment or an admin with
// :select_final_grade rights.
func (c *Canvas) ModeratedGradingBulkSelectProvisionalGrades(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModeratedGradingShowProvisionalGradeStatusForAStudent2 API call: Tell whether the student's submission needs one or
// more provisional grades.
func (c *Canvas) ModeratedGradingShowProvisionalGradeStatusForAStudent2(progress *task.Progress, studentID *int) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/1/assignments/2/provisional_grades/status")
	params := map[string]interface{}{}
	if studentID != nil {
		params["student_id"] = *studentID
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModeratedGradingSelectProvisionalGrade API call: Choose which provisional grade the student should receive for a
// submission. The caller must be the final grader for the assignment or an admin with :select_final_grade rights.
func (c *Canvas) ModeratedGradingSelectProvisionalGrade(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ModeratedGradingPublishProvisionalGradesForAnAssignment API call: Publish the selected provisional grade for all
// submissions to an assignment. Use the "Select provisional grade" endpoint to choose which provisional grade to
// publish for a particular submission. Students not in the moderation set will have their one and only provisional
// grade published. WARNING: This is irreversible. This will overwrite existing grades in the gradebook.
func (c *Canvas) ModeratedGradingPublishProvisionalGradesForAnAssignment(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/1/assignments/2/provisional_grades/publish")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// LoginsListUserLogins API call: Given a user ID, return a paginated list of that user's logins for the given account.
func (c *Canvas) LoginsListUserLogins(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// LoginsCreateAUserLogin API call: Create a new login for an existing user in the given account.
func (c *Canvas) LoginsCreateAUserLogin(progress *task.Progress, user *string, login *string, accountID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("accounts/%s/logins", accountID)
	params := map[string]interface{}{}
	if user != nil {
		params["user"] = *user
	}
	if login != nil {
		params["login"] = *login
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// LoginsEditAUserLogin API call: Update an existing login for a user in the given account.
func (c *Canvas) LoginsEditAUserLogin(progress *task.Progress, login *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if login != nil {
		params["login"] = *login
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// LoginsDeleteAUserLogin API call: Delete an existing login.
func (c *Canvas) LoginsDeleteAUserLogin(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/:user_id/logins/:login_id")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RolesListRoles API call: A paginated list of the roles available to an account.
func (c *Canvas) RolesListRoles(progress *task.Progress, accountID *string, state *RolesListRolesState, showInherited *bool) ([]Role, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if accountID != nil {
		params["account_id"] = *accountID
	}
	if state != nil {
		params["state"] = *state
	}
	if showInherited != nil {
		params["show_inherited"] = *showInherited
	}
	responseCtor := func() interface{} {
		return &[]Role{}
	}
	var res []Role
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Role)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RolesGetASingleRole API call: Retrieve information about a single role
func (c *Canvas) RolesGetASingleRole(progress *task.Progress, accountID *string, roleID *int, role *string) (*Role, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if accountID != nil {
		params["account_id"] = *accountID
	}
	if roleID != nil {
		params["role_id"] = *roleID
	}
	if role != nil {
		params["role"] = *role
	}
	responseCtor := func() interface{} {
		return &Role{}
	}
	var res *Role
	callback := func(obj interface{}) error {
		res = obj.(*Role)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RolesCreateANewRole API call: Create a new course-level or account-level role.
func (c *Canvas) RolesCreateANewRole(progress *task.Progress, label *string, role *string, baseRoleType *RolesCreateANewRoleBaseRoleType, permissions *bool, accountID string) (*Role, error) {
	endpoint := fmt.Sprintf("accounts/%s/roles.json", accountID)
	params := map[string]interface{}{}
	if label != nil {
		params["label"] = *label
	}
	if role != nil {
		params["role"] = *role
	}
	if baseRoleType != nil {
		params["base_role_type"] = *baseRoleType
	}
	if permissions != nil {
		params["permissions"] = *permissions
	}
	responseCtor := func() interface{} {
		return &Role{}
	}
	var res *Role
	callback := func(obj interface{}) error {
		res = obj.(*Role)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RolesDeactivateARole API call: Deactivates a custom role.  This hides it in the user interface and prevents it from
// being assigned to new users.  Existing users assigned to the role will continue to function with the same permissions
// they had previously. Built-in roles cannot be deactivated.
func (c *Canvas) RolesDeactivateARole(progress *task.Progress, roleID *int, role *string) (*Role, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if roleID != nil {
		params["role_id"] = *roleID
	}
	if role != nil {
		params["role"] = *role
	}
	responseCtor := func() interface{} {
		return &Role{}
	}
	var res *Role
	callback := func(obj interface{}) error {
		res = obj.(*Role)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RolesActivateARole API call: Re-activates an inactive role (allowing it to be assigned to new users)
func (c *Canvas) RolesActivateARole(progress *task.Progress, roleID *int, role *string) (*Role, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if roleID != nil {
		params["role_id"] = *roleID
	}
	if role != nil {
		params["role"] = *role
	}
	responseCtor := func() interface{} {
		return &Role{}
	}
	var res *Role
	callback := func(obj interface{}) error {
		res = obj.(*Role)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RolesUpdateARole API call: Update permissions for an existing role. Recognized roles are: * TeacherEnrollment *
// StudentEnrollment * TaEnrollment * ObserverEnrollment * DesignerEnrollment * AccountAdmin * Any previously created
// custom role
func (c *Canvas) RolesUpdateARole(progress *task.Progress, label *string, permissions *bool) (*Role, error) {
	endpoint := fmt.Sprintf("accounts/:account_id/roles/2")
	params := map[string]interface{}{}
	if label != nil {
		params["label"] = *label
	}
	if permissions != nil {
		params["permissions"] = *permissions
	}
	responseCtor := func() interface{} {
		return &Role{}
	}
	var res *Role
	callback := func(obj interface{}) error {
		res = obj.(*Role)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsCreateASingleRubricAssessment API call: Returns the rubric assessment with the given id. The returned object
// also provides the information of :ratings, :assessor_name, :related_group_submissions_and_assessments, :artifact
func (c *Canvas) RubricsCreateASingleRubricAssessment(progress *task.Progress, courseID *int, rubricAssociationID *int, provisional *string, final *string, gradedAnonymously *bool, rubricAssessment *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	if rubricAssociationID != nil {
		params["rubric_association_id"] = *rubricAssociationID
	}
	if provisional != nil {
		params["provisional"] = *provisional
	}
	if final != nil {
		params["final"] = *final
	}
	if gradedAnonymously != nil {
		params["graded_anonymously"] = *gradedAnonymously
	}
	if rubricAssessment != nil {
		params["rubric_assessment"] = *rubricAssessment
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsUpdateASingleRubricAssessment API call: Returns the rubric assessment with the given id. The returned object
// also provides the information of :ratings, :assessor_name, :related_group_submissions_and_assessments, :artifact
func (c *Canvas) RubricsUpdateASingleRubricAssessment(progress *task.Progress, id *int, courseID *int, rubricAssociationID *int, provisional *string, final *string, gradedAnonymously *bool, rubricAssessment *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	if rubricAssociationID != nil {
		params["rubric_association_id"] = *rubricAssociationID
	}
	if provisional != nil {
		params["provisional"] = *provisional
	}
	if final != nil {
		params["final"] = *final
	}
	if gradedAnonymously != nil {
		params["graded_anonymously"] = *gradedAnonymously
	}
	if rubricAssessment != nil {
		params["rubric_assessment"] = *rubricAssessment
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsDeleteASingleRubricAssessment API call: Deletes a rubric assessment
func (c *Canvas) RubricsDeleteASingleRubricAssessment(progress *task.Progress) (*RubricAssessment, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &RubricAssessment{}
	}
	var res *RubricAssessment
	callback := func(obj interface{}) error {
		res = obj.(*RubricAssessment)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsCreateARubricAssociation API call: Returns the rubric with the given id.
func (c *Canvas) RubricsCreateARubricAssociation(progress *task.Progress, rubricAssociation *int) (*RubricAssociation, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if rubricAssociation != nil {
		params["rubric_association"] = *rubricAssociation
	}
	responseCtor := func() interface{} {
		return &RubricAssociation{}
	}
	var res *RubricAssociation
	callback := func(obj interface{}) error {
		res = obj.(*RubricAssociation)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsUpdateARubricAssociation API call: Returns the rubric with the given id.
func (c *Canvas) RubricsUpdateARubricAssociation(progress *task.Progress, id *int, rubricAssociation *int) (*RubricAssociation, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	if rubricAssociation != nil {
		params["rubric_association"] = *rubricAssociation
	}
	responseCtor := func() interface{} {
		return &RubricAssociation{}
	}
	var res *RubricAssociation
	callback := func(obj interface{}) error {
		res = obj.(*RubricAssociation)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsDeleteARubricAssociation API call: Delete the RubricAssociation with the given ID
func (c *Canvas) RubricsDeleteARubricAssociation(progress *task.Progress) (*RubricAssociation, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &RubricAssociation{}
	}
	var res *RubricAssociation
	callback := func(obj interface{}) error {
		res = obj.(*RubricAssociation)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsListRubrics API call: Returns the paginated list of active rubrics for the current context.
func (c *Canvas) RubricsListRubrics(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsGetASingleRubric API call: Returns the rubric with the given id.
func (c *Canvas) RubricsGetASingleRubric(progress *task.Progress, include *RubricsGetASingleRubricInclude, style *RubricsGetASingleRubricStyle) (*Rubric, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if style != nil {
		params["style"] = *style
	}
	responseCtor := func() interface{} {
		return &Rubric{}
	}
	var res *Rubric
	callback := func(obj interface{}) error {
		res = obj.(*Rubric)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsCreateASingleRubric API call: Returns the rubric with the given id. Unfortuantely this endpoint does not
// return a standard Rubric object, instead it returns a hash that looks like { 'rubric': Rubric, 'rubric_association':
// RubricAssociation } This may eventually be deprecated in favor of a more standardized return value, but that is not
// currently planned.
func (c *Canvas) RubricsCreateASingleRubric(progress *task.Progress, id *int, rubricAssociationID *int, rubric *string, rubricAssociation *int) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	if rubricAssociationID != nil {
		params["rubric_association_id"] = *rubricAssociationID
	}
	if rubric != nil {
		params["rubric"] = *rubric
	}
	if rubricAssociation != nil {
		params["rubric_association"] = *rubricAssociation
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsUpdateASingleRubric API call: Returns the rubric with the given id. Unfortuantely this endpoint does not
// return a standard Rubric object, instead it returns a hash that looks like { 'rubric': Rubric, 'rubric_association':
// RubricAssociation } This may eventually be deprecated in favor of a more standardized return value, but that is not
// currently planned.
func (c *Canvas) RubricsUpdateASingleRubric(progress *task.Progress, id *int, rubricAssociationID *int, rubric *string, rubricAssociation *int) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if id != nil {
		params["id"] = *id
	}
	if rubricAssociationID != nil {
		params["rubric_association_id"] = *rubricAssociationID
	}
	if rubric != nil {
		params["rubric"] = *rubric
	}
	if rubricAssociation != nil {
		params["rubric_association"] = *rubricAssociation
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// RubricsDeleteASingleRubric API call: Deletes a Rubric and removes all RubricAssociations.
func (c *Canvas) RubricsDeleteASingleRubric(progress *task.Progress) (*Rubric, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Rubric{}
	}
	var res *Rubric
	callback := func(obj interface{}) error {
		res = obj.(*Rubric)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// APITokenScopesListScopes API call: A list of scopes that can be applied to developer keys and access tokens.
func (c *Canvas) APITokenScopesListScopes(progress *task.Progress, groupBy *APITokenScopesListScopesGroupBy) ([]Scope, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if groupBy != nil {
		params["group_by"] = *groupBy
	}
	responseCtor := func() interface{} {
		return &[]Scope{}
	}
	var res []Scope
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Scope)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SearchFindRecipients API call: Find valid recipients (users, courses and groups) that the current user can send
// messages to. The /api/v1/search/recipients path is the preferred endpoint, /api/v1/conversations/find_recipients is
// deprecated. Pagination is supported.
func (c *Canvas) SearchFindRecipients(progress *task.Progress, search *string, context *string, exclude *string, typeName *SearchFindRecipientsType, userID *int, fromConversationID *int, permissions *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if search != nil {
		params["search"] = *search
	}
	if context != nil {
		params["context"] = *context
	}
	if exclude != nil {
		params["exclude"] = *exclude
	}
	if typeName != nil {
		params["type"] = *typeName
	}
	if userID != nil {
		params["user_id"] = *userID
	}
	if fromConversationID != nil {
		params["from_conversation_id"] = *fromConversationID
	}
	if permissions != nil {
		params["permissions"] = *permissions
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SearchListAllCourses API call: A paginated list of all courses visible in the public index
func (c *Canvas) SearchListAllCourses(progress *task.Progress, search *string, publicOnly *bool, openEnrollmentOnly *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if search != nil {
		params["search"] = *search
	}
	if publicOnly != nil {
		params["public_only"] = *publicOnly
	}
	if openEnrollmentOnly != nil {
		params["open_enrollment_only"] = *openEnrollmentOnly
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SectionsListCourseSections API call: A paginated list of the list of sections for this course.
func (c *Canvas) SectionsListCourseSections(progress *task.Progress, include *SectionsListCourseSectionsInclude) ([]Section, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]Section{}
	}
	var res []Section
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Section)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SectionsCreateCourseSection API call: Creates a new section for this course.
func (c *Canvas) SectionsCreateCourseSection(progress *task.Progress, courseSection *string, enableSisReactivation *bool) (*Section, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseSection != nil {
		params["course_section"] = *courseSection
	}
	if enableSisReactivation != nil {
		params["enable_sis_reactivation"] = *enableSisReactivation
	}
	responseCtor := func() interface{} {
		return &Section{}
	}
	var res *Section
	callback := func(obj interface{}) error {
		res = obj.(*Section)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SectionsCrossListASection API call: Move the Section to another course.  The new course may be in a different account
// (department), but must belong to the same root account (institution).
func (c *Canvas) SectionsCrossListASection(progress *task.Progress) (*Section, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Section{}
	}
	var res *Section
	callback := func(obj interface{}) error {
		res = obj.(*Section)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SectionsDeCrossListASection API call: Undo cross-listing of a Section, returning it to its original course.
func (c *Canvas) SectionsDeCrossListASection(progress *task.Progress) (*Section, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Section{}
	}
	var res *Section
	callback := func(obj interface{}) error {
		res = obj.(*Section)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SectionsEditASection API call: Modify an existing section.
func (c *Canvas) SectionsEditASection(progress *task.Progress, courseSection *string) (*Section, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if courseSection != nil {
		params["course_section"] = *courseSection
	}
	responseCtor := func() interface{} {
		return &Section{}
	}
	var res *Section
	callback := func(obj interface{}) error {
		res = obj.(*Section)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SectionsGetSectionInformation API call: Gets details about a specific section
func (c *Canvas) SectionsGetSectionInformation(progress *task.Progress, include *SectionsGetSectionInformationInclude) (*Section, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &Section{}
	}
	var res *Section
	callback := func(obj interface{}) error {
		res = obj.(*Section)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SectionsDeleteASection API call: Delete an existing section.  Returns the former Section.
func (c *Canvas) SectionsDeleteASection(progress *task.Progress) (*Section, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Section{}
	}
	var res *Section
	callback := func(obj interface{}) error {
		res = obj.(*Section)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ServicesGetKalturaConfig API call: Return the config information for the Kaltura plugin in json format.
func (c *Canvas) ServicesGetKalturaConfig(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// ServicesStartKalturaSession API call: Start a new Kaltura session, so that new media can be recorded and uploaded to
// this Canvas instance's Kaltura instance.
func (c *Canvas) ServicesStartKalturaSession(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SharedBrandConfigsShareABrandConfigTheme API call: Create a SharedBrandConfig, which will give the given brand_config
// a name and make it available to other users of this account.
func (c *Canvas) SharedBrandConfigsShareABrandConfigTheme(progress *task.Progress, sharedBrandConfig *string, accountID string) (*SharedBrandConfig, error) {
	endpoint := fmt.Sprintf("accounts/%s/shared_brand_configs", accountID)
	params := map[string]interface{}{}
	if sharedBrandConfig != nil {
		params["shared_brand_config"] = *sharedBrandConfig
	}
	responseCtor := func() interface{} {
		return &SharedBrandConfig{}
	}
	var res *SharedBrandConfig
	callback := func(obj interface{}) error {
		res = obj.(*SharedBrandConfig)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SharedBrandConfigsUpdateASharedTheme API call: Update the specified shared_brand_config with a new name or to point
// to a new brand_config. Uses same parameters as create.
func (c *Canvas) SharedBrandConfigsUpdateASharedTheme(progress *task.Progress, accountID string, sharedBrandConfigID string) (*SharedBrandConfig, error) {
	endpoint := fmt.Sprintf("accounts/%s/shared_brand_configs/%s", accountID, sharedBrandConfigID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &SharedBrandConfig{}
	}
	var res *SharedBrandConfig
	callback := func(obj interface{}) error {
		res = obj.(*SharedBrandConfig)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SharedBrandConfigsUnShareABrandConfigTheme API call: Delete a SharedBrandConfig, which will unshare it so you nor
// anyone else in your account will see it as an option to pick from.
func (c *Canvas) SharedBrandConfigsUnShareABrandConfigTheme(progress *task.Progress, id string) (*SharedBrandConfig, error) {
	endpoint := fmt.Sprintf("shared_brand_configs/%s", id)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &SharedBrandConfig{}
	}
	var res *SharedBrandConfig
	callback := func(obj interface{}) error {
		res = obj.(*SharedBrandConfig)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISIntegrationRetrieveAssignmentsEnabledForGradeExportToSIS API call: Retrieve a list of published assignments
// flagged as "post_to_sis". See the Assignments API for more details on assignments. Assignment group and section
// information are included for convenience. Each section includes course information for the origin course and the
// cross-listed course, if applicable. The `origin_course` is the course to which the section belongs or the course from
// which the section was cross-listed. Generally, the `origin_course` should be preferred when performing integration
// work. The `xlist_course` is provided for consistency and is only present when the section has been cross-listed. See
// Sections API and Courses Api for me details. The `override` is only provided if the Differentiated Assignments course
// feature is turned on and the assignment has an override for that section. When there is an override for the
// assignment the override object's keys/values can be merged with the top level assignment object to create a view of
// the assignment object specific to that section. See Assignments api for more information on assignment overrides.
func (c *Canvas) SISIntegrationRetrieveAssignmentsEnabledForGradeExportToSIS(progress *task.Progress, accountID *interface{}, courseID *interface{}, startsBefore *interface{}, endsAfter *interface{}, include *interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if accountID != nil {
		params["account_id"] = *accountID
	}
	if courseID != nil {
		params["course_id"] = *courseID
	}
	if startsBefore != nil {
		params["starts_before"] = *startsBefore
	}
	if endsAfter != nil {
		params["ends_after"] = *endsAfter
	}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISImportErrorsGetSISImportErrorList API call: Returns the list of SIS import errors for an account or a SIS import.
// Import errors are only stored for 30 days.
func (c *Canvas) SISImportErrorsGetSISImportErrorList(progress *task.Progress, failure *bool) ([]SisImportError, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if failure != nil {
		params["failure"] = *failure
	}
	responseCtor := func() interface{} {
		return &[]SisImportError{}
	}
	var res []SisImportError
	callback := func(obj interface{}) error {
		arr := *obj.(*[]SisImportError)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISImportsGetSISImportList API call: Returns the list of SIS imports for an account
func (c *Canvas) SISImportsGetSISImportList(progress *task.Progress, createdSince *time.Time, workflowState *SISImportsGetSISImportListWorkflowState) ([]SisImport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if createdSince != nil {
		params["created_since"] = *createdSince
	}
	if workflowState != nil {
		params["workflow_state"] = *workflowState
	}
	responseCtor := func() interface{} {
		return &[]SisImport{}
	}
	var res []SisImport
	callback := func(obj interface{}) error {
		arr := *obj.(*[]SisImport)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISImportsGetTheCurrentImportingSISImport API call: Returns the SIS imports that are currently processing for an
// account. If no imports are running, will return an empty array. Example: curl
// https://<canvas>/api/v1/accounts/<account_id>/sis_imports/importing \ -H 'Authorization: Bearer <token>'
func (c *Canvas) SISImportsGetTheCurrentImportingSISImport(progress *task.Progress) (*SisImport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &SisImport{}
	}
	var res *SisImport
	callback := func(obj interface{}) error {
		res = obj.(*SisImport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISImportsImportSISData API call: Import SIS data into Canvas. Must be on a root account with SIS imports enabled.
// For more information on the format that's expected here, please see the "SIS CSV" section in the API docs.
func (c *Canvas) SISImportsImportSISData(progress *task.Progress, importType *string, attachment *interface{}, extension *string, batchMode *bool, batchModeTermID *string, multiTermBatchMode *bool, skipDeletes *bool, overrideSisStickiness *bool, addSisStickiness *bool, clearSisStickiness *bool, diffingDataSetIdentifier *string, diffingRemasterDataSet *bool, diffingDropStatus *SISImportsImportSISDataDiffingDropStatus, changeThreshold *int, diffRowCountThreshold *int) (*SisImport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if importType != nil {
		params["import_type"] = *importType
	}
	if attachment != nil {
		params["attachment"] = *attachment
	}
	if extension != nil {
		params["extension"] = *extension
	}
	if batchMode != nil {
		params["batch_mode"] = *batchMode
	}
	if batchModeTermID != nil {
		params["batch_mode_term_id"] = *batchModeTermID
	}
	if multiTermBatchMode != nil {
		params["multi_term_batch_mode"] = *multiTermBatchMode
	}
	if skipDeletes != nil {
		params["skip_deletes"] = *skipDeletes
	}
	if overrideSisStickiness != nil {
		params["override_sis_stickiness"] = *overrideSisStickiness
	}
	if addSisStickiness != nil {
		params["add_sis_stickiness"] = *addSisStickiness
	}
	if clearSisStickiness != nil {
		params["clear_sis_stickiness"] = *clearSisStickiness
	}
	if diffingDataSetIdentifier != nil {
		params["diffing_data_set_identifier"] = *diffingDataSetIdentifier
	}
	if diffingRemasterDataSet != nil {
		params["diffing_remaster_data_set"] = *diffingRemasterDataSet
	}
	if diffingDropStatus != nil {
		params["diffing_drop_status"] = *diffingDropStatus
	}
	if changeThreshold != nil {
		params["change_threshold"] = *changeThreshold
	}
	if diffRowCountThreshold != nil {
		params["diff_row_count_threshold"] = *diffRowCountThreshold
	}
	responseCtor := func() interface{} {
		return &SisImport{}
	}
	var res *SisImport
	callback := func(obj interface{}) error {
		res = obj.(*SisImport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISImportsGetSISImportStatus API call: Get the status of an already created SIS import. Examples: curl
// https://<canvas>/api/v1/accounts/<account_id>/sis_imports/<sis_import_id> \ -H 'Authorization: Bearer <token>'
func (c *Canvas) SISImportsGetSISImportStatus(progress *task.Progress) (*SisImport, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &SisImport{}
	}
	var res *SisImport
	callback := func(obj interface{}) error {
		res = obj.(*SisImport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISImportsRestoreWorkflowStatesOfSISImportedItems API call: This will restore the the workflow_state for all the
// items that changed their workflow_state during the import being restored. This will restore states for items imported
// with the following importers: accounts.csv terms.csv courses.csv sections.csv group_categories.csv groups.csv
// users.csv admins.csv This also restores states for other items that changed during the import. An example would be if
// an enrollment was deleted from a sis import and the group_membership was also deleted as a result of the enrollment
// deletion, both items would be restored when the sis batch is restored.
func (c *Canvas) SISImportsRestoreWorkflowStatesOfSISImportedItems(progress *task.Progress, batchMode *bool, undeleteOnly *bool, unconcludeOnly *bool, accountID string, sisImportID string) (*Progress, error) {
	endpoint := fmt.Sprintf("accounts/%s/sis_imports/%s/restore_states", accountID, sisImportID)
	params := map[string]interface{}{}
	if batchMode != nil {
		params["batch_mode"] = *batchMode
	}
	if undeleteOnly != nil {
		params["undelete_only"] = *undeleteOnly
	}
	if unconcludeOnly != nil {
		params["unconclude_only"] = *unconcludeOnly
	}
	responseCtor := func() interface{} {
		return &Progress{}
	}
	var res *Progress
	callback := func(obj interface{}) error {
		res = obj.(*Progress)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISImportsAbortSISImport API call: Abort a SIS import that has not completed. Aborting a sis batch that is running
// can take some time for every process to see the abort event. Subsequent sis batches begin to process 10 minutes after
// the abort to allow each process to clean up properly.
func (c *Canvas) SISImportsAbortSISImport(progress *task.Progress, accountID string, sisImportID string) (*SisImport, error) {
	endpoint := fmt.Sprintf("accounts/%s/sis_imports/%s/abort", accountID, sisImportID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &SisImport{}
	}
	var res *SisImport
	callback := func(obj interface{}) error {
		res = obj.(*SisImport)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SISImportsAbortAllPendingSISImports API call: Abort already created but not processed or processing SIS imports.
func (c *Canvas) SISImportsAbortAllPendingSISImports(progress *task.Progress, accountID string) (*bool, error) {
	endpoint := fmt.Sprintf("accounts/%s/sis_imports/abort_all_pending", accountID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		var tmp bool
		return &tmp
	}
	var res *bool
	callback := func(obj interface{}) error {
		res = obj.(*bool)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsCreateANewSubAccount API call: Add a new sub-account to a given account.
func (c *Canvas) AccountsCreateANewSubAccount(progress *task.Progress, account *string) (*Account, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if account != nil {
		params["account"] = *account
	}
	responseCtor := func() interface{} {
		return &Account{}
	}
	var res *Account
	callback := func(obj interface{}) error {
		res = obj.(*Account)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// AccountsDeleteASubAccount API call: Cannot delete an account with active courses or active sub_accounts. Cannot
// delete a root_account
func (c *Canvas) AccountsDeleteASubAccount(progress *task.Progress) (*Account, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Account{}
	}
	var res *Account
	callback := func(obj interface{}) error {
		res = obj.(*Account)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionCommentsUploadAFile API call: Upload a file to attach to a submission comment See the
// {file:file_uploads.html File Upload Documentation} for details on the file upload workflow. The final step of the
// file upload workflow will return the attachment data, including the new file id. The caller can then PUT the file_id
// to the submission API to attach it to a comment
func (c *Canvas) SubmissionCommentsUploadAFile(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsListAssignmentSubmissions API call: A paginated list of all existing submissions for an assignment.
func (c *Canvas) SubmissionsListAssignmentSubmissions(progress *task.Progress, include *SubmissionsListAssignmentSubmissionsInclude, grouped *bool) ([]Submission, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if grouped != nil {
		params["grouped"] = *grouped
	}
	responseCtor := func() interface{} {
		return &[]Submission{}
	}
	var res []Submission
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Submission)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsListSubmissionsForMultipleAssignments API call: A paginated list of all existing submissions for a given
// set of students and assignments.
func (c *Canvas) SubmissionsListSubmissionsForMultipleAssignments(progress *task.Progress, studentIds *string, assignmentIds *string, grouped *bool, postToSis *bool, submittedSince *time.Time, gradedSince *time.Time, gradingPeriodID *int, workflowState *SubmissionsListSubmissionsForMultipleAssignmentsWorkflowState, enrollmentState *SubmissionsListSubmissionsForMultipleAssignmentsEnrollmentState, stateBasedOnDate *bool, order *SubmissionsListSubmissionsForMultipleAssignmentsOrder, orderDirection *SubmissionsListSubmissionsForMultipleAssignmentsOrderDirection, include *SubmissionsListSubmissionsForMultipleAssignmentsInclude) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if studentIds != nil {
		params["student_ids"] = *studentIds
	}
	if assignmentIds != nil {
		params["assignment_ids"] = *assignmentIds
	}
	if grouped != nil {
		params["grouped"] = *grouped
	}
	if postToSis != nil {
		params["post_to_sis"] = *postToSis
	}
	if submittedSince != nil {
		params["submitted_since"] = *submittedSince
	}
	if gradedSince != nil {
		params["graded_since"] = *gradedSince
	}
	if gradingPeriodID != nil {
		params["grading_period_id"] = *gradingPeriodID
	}
	if workflowState != nil {
		params["workflow_state"] = *workflowState
	}
	if enrollmentState != nil {
		params["enrollment_state"] = *enrollmentState
	}
	if stateBasedOnDate != nil {
		params["state_based_on_date"] = *stateBasedOnDate
	}
	if order != nil {
		params["order"] = *order
	}
	if orderDirection != nil {
		params["order_direction"] = *orderDirection
	}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsGetASingleSubmission API call: Get a single submission, based on user id.
func (c *Canvas) SubmissionsGetASingleSubmission(progress *task.Progress, include *SubmissionsGetASingleSubmissionInclude) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsUploadAFile API call: Upload a file to a submission. This API endpoint is the first step in uploading a
// file to a submission as a student. See the {file:file_uploads.html File Upload Documentation} for details on the file
// upload workflow. The final step of the file upload workflow will return the attachment data, including the new file
// id. The caller can then POST to submit the +online_upload+ assignment with these file ids.
func (c *Canvas) SubmissionsUploadAFile(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsListGradeableStudents API call: A paginated list of students eligible to submit the assignment. The caller
// must have permission to view grades. If anonymous grading is enabled for the current assignment and the
// allow_new_anonymous_id parameter is passed, the returned data will not include any values identifying the student,
// but will instead include an assignment-specific anonymous ID for each student. Section-limited instructors will only
// see students in their own sections.
func (c *Canvas) SubmissionsListGradeableStudents(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsListMultipleAssignmentsGradeableStudents API call
func (c *Canvas) SubmissionsListMultipleAssignmentsGradeableStudents(progress *task.Progress, assignmentIds *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if assignmentIds != nil {
		params["assignment_ids"] = *assignmentIds
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsGradeOrCommentOnMultipleSubmissions API call: Update the grading and comments on multiple student's
// assignment submissions in an asynchronous job. The user must have permission to manage grades in the appropriate
// context (course or section).
func (c *Canvas) SubmissionsGradeOrCommentOnMultipleSubmissions(progress *task.Progress, gradeData *string) (*Progress, error) {
	endpoint := fmt.Sprintf("courses/1/assignments/2/submissions/update_grades")
	params := map[string]interface{}{}
	if gradeData != nil {
		params["grade_data"] = *gradeData
	}
	responseCtor := func() interface{} {
		return &Progress{}
	}
	var res *Progress
	callback := func(obj interface{}) error {
		res = obj.(*Progress)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsMarkSubmissionAsRead API call: No request fields are necessary. On success, the response will be 204 No
// Content with an empty body.
func (c *Canvas) SubmissionsMarkSubmissionAsRead(progress *task.Progress, courseID string, assignmentID string, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/assignments/%s/submissions/%s/read.json", courseID, assignmentID, userID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsMarkSubmissionAsUnread API call: No request fields are necessary. On success, the response will be 204 No
// Content with an empty body.
func (c *Canvas) SubmissionsMarkSubmissionAsUnread(progress *task.Progress, courseID string, assignmentID string, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("courses/%s/assignments/%s/submissions/%s/read.json", courseID, assignmentID, userID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsSubmissionSummary API call: Returns the number of submissions for the given assignment based on gradeable
// students that fall into three categories: graded, ungraded, not submitted.
func (c *Canvas) SubmissionsSubmissionSummary(progress *task.Progress, grouped *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if grouped != nil {
		params["grouped"] = *grouped
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionsSubmitAnAssignment API call: Make a submission for an assignment. You must be enrolled as a student in the
// course/section to do this. All online turn-in submission types are supported in this API. However, there are a few
// things that are not yet supported: * Files can be submitted based on a file ID of a user or group file or through the
// {api:SubmissionsApiController#create_file file upload API}. However, there is no API yet for listing the user and
// group files. * Media comments can be submitted, however, there is no API yet for creating a media comment to submit.
// * Integration with Google Docs is not yet supported.
func (c *Canvas) SubmissionsSubmitAnAssignment(progress *task.Progress, comment *string, submission *SubmissionsSubmitAnAssignmentSubmission) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if comment != nil {
		params["comment"] = *comment
	}
	if submission != nil {
		params["submission"] = *submission
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// TabsListAvailableTabsForACourseOrGroup API call: Returns a paginated list of navigation tabs available in the current
// context.
func (c *Canvas) TabsListAvailableTabsForACourseOrGroup(progress *task.Progress, groupID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("groups/%s/tabs", groupID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// TabsUpdateATabForACourse API call: Home and Settings tabs are not manageable, and can't be hidden or moved Returns a
// tab object
func (c *Canvas) TabsUpdateATabForACourse(progress *task.Progress, position *interface{}, hidden *bool, courseID string) (*Tab, error) {
	endpoint := fmt.Sprintf("courses/%s/tabs/tab_id", courseID)
	params := map[string]interface{}{}
	if position != nil {
		params["position"] = *position
	}
	if hidden != nil {
		params["hidden"] = *hidden
	}
	responseCtor := func() interface{} {
		return &Tab{}
	}
	var res *Tab
	callback := func(obj interface{}) error {
		res = obj.(*Tab)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentTermsListEnrollmentTerms API call: An object with a paginated list of all of the terms in the account.
func (c *Canvas) EnrollmentTermsListEnrollmentTerms(progress *task.Progress, workflowState *EnrollmentTermsListEnrollmentTermsWorkflowState, include *EnrollmentTermsListEnrollmentTermsInclude) (*EnrollmentTermsList, error) {
	endpoint := fmt.Sprintf("accounts/1/terms")
	params := map[string]interface{}{}
	if workflowState != nil {
		params["workflow_state"] = *workflowState
	}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &EnrollmentTermsList{}
	}
	var res *EnrollmentTermsList
	callback := func(obj interface{}) error {
		res = obj.(*EnrollmentTermsList)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentTermsRetrieveEnrollmentTerm API call: Retrieves the details for an enrollment term in the account. Includes
// overrides by default.
func (c *Canvas) EnrollmentTermsRetrieveEnrollmentTerm(progress *task.Progress) (*EnrollmentTerm, error) {
	endpoint := fmt.Sprintf("accounts/1/terms/2")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &EnrollmentTerm{}
	}
	var res *EnrollmentTerm
	callback := func(obj interface{}) error {
		res = obj.(*EnrollmentTerm)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentTermsCreateEnrollmentTerm API call: Create a new enrollment term for the specified account.
func (c *Canvas) EnrollmentTermsCreateEnrollmentTerm(progress *task.Progress, enrollmentTerm *string) (*EnrollmentTerm, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if enrollmentTerm != nil {
		params["enrollment_term"] = *enrollmentTerm
	}
	responseCtor := func() interface{} {
		return &EnrollmentTerm{}
	}
	var res *EnrollmentTerm
	callback := func(obj interface{}) error {
		res = obj.(*EnrollmentTerm)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentTermsUpdateEnrollmentTerm API call: Update an existing enrollment term for the specified account.
func (c *Canvas) EnrollmentTermsUpdateEnrollmentTerm(progress *task.Progress, enrollmentTerm *string) (*EnrollmentTerm, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if enrollmentTerm != nil {
		params["enrollment_term"] = *enrollmentTerm
	}
	responseCtor := func() interface{} {
		return &EnrollmentTerm{}
	}
	var res *EnrollmentTerm
	callback := func(obj interface{}) error {
		res = obj.(*EnrollmentTerm)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// EnrollmentTermsDeleteEnrollmentTerm API call: Delete the specified enrollment term.
func (c *Canvas) EnrollmentTermsDeleteEnrollmentTerm(progress *task.Progress) (*EnrollmentTerm, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &EnrollmentTerm{}
	}
	var res *EnrollmentTerm
	callback := func(obj interface{}) error {
		res = obj.(*EnrollmentTerm)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesSetUsageRights API call: Sets copyright and license information for one or more files
func (c *Canvas) FilesSetUsageRights(progress *task.Progress, fileIds *interface{}, folderIds *interface{}, publish *bool, usageRights *FilesSetUsageRightsUsageRights) (*UsageRights, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if fileIds != nil {
		params["file_ids"] = *fileIds
	}
	if folderIds != nil {
		params["folder_ids"] = *folderIds
	}
	if publish != nil {
		params["publish"] = *publish
	}
	if usageRights != nil {
		params["usage_rights"] = *usageRights
	}
	responseCtor := func() interface{} {
		return &UsageRights{}
	}
	var res *UsageRights
	callback := func(obj interface{}) error {
		res = obj.(*UsageRights)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesRemoveUsageRights API call: Removes copyright and license information associated with one or more files
func (c *Canvas) FilesRemoveUsageRights(progress *task.Progress, fileIds *interface{}, folderIds *interface{}) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if fileIds != nil {
		params["file_ids"] = *fileIds
	}
	if folderIds != nil {
		params["folder_ids"] = *folderIds
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// FilesListLicenses API call: A paginated list of licenses that can be applied
func (c *Canvas) FilesListLicenses(progress *task.Progress) ([]License, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]License{}
	}
	var res []License
	callback := func(obj interface{}) error {
		arr := *obj.(*[]License)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UserObserveesListObservees API call: A paginated list of the users that the given user is observing. *Note:* all
// users are allowed to list their own observees. Administrators can list other users' observees. The returned observees
// will include an attribute "observation_link_root_account_ids", a list of ids for the root accounts the observer and
// observee are linked on. The observer will only be able to observe in courses associated with these root accounts.
func (c *Canvas) UserObserveesListObservees(progress *task.Progress, include *UserObserveesListObserveesInclude, userID string) ([]User, error) {
	endpoint := fmt.Sprintf("users/%s/observees", userID)
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UserObserveesAddAnObserveeWithCredentials API call: Register the given user to observe another user, given the
// observee's credentials. *Note:* all users are allowed to add their own observees, given the observee's credentials or
// access token are provided. Administrators can add observees given credentials, access token or the
// {api:UserObserveesController#update observee's id}.
func (c *Canvas) UserObserveesAddAnObserveeWithCredentials(progress *task.Progress, observee *string, accessToken *string, pairingCode *string, rootAccountID *int, userID string) (*User, error) {
	endpoint := fmt.Sprintf("users/%s/observees", userID)
	params := map[string]interface{}{}
	if observee != nil {
		params["observee"] = *observee
	}
	if accessToken != nil {
		params["access_token"] = *accessToken
	}
	if pairingCode != nil {
		params["pairing_code"] = *pairingCode
	}
	if rootAccountID != nil {
		params["root_account_id"] = *rootAccountID
	}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UserObserveesShowAnObservee API call: Gets information about an observed user. *Note:* all users are allowed to view
// their own observees.
func (c *Canvas) UserObserveesShowAnObservee(progress *task.Progress, userID string, observeeID string) (*User, error) {
	endpoint := fmt.Sprintf("users/%s/observees/%s", userID, observeeID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UserObserveesAddAnObservee API call: Registers a user as being observed by the given user.
func (c *Canvas) UserObserveesAddAnObservee(progress *task.Progress, rootAccountID *int, userID string, observeeID string) (*User, error) {
	endpoint := fmt.Sprintf("users/%s/observees/%s", userID, observeeID)
	params := map[string]interface{}{}
	if rootAccountID != nil {
		params["root_account_id"] = *rootAccountID
	}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UserObserveesRemoveAnObservee API call: Unregisters a user as being observed by the given user.
func (c *Canvas) UserObserveesRemoveAnObservee(progress *task.Progress, rootAccountID *int, userID string, observeeID string) (*User, error) {
	endpoint := fmt.Sprintf("users/%s/observees/%s", userID, observeeID)
	params := map[string]interface{}{}
	if rootAccountID != nil {
		params["root_account_id"] = *rootAccountID
	}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersListUsersInAccount API call: A paginated list of of users associated with this account.
func (c *Canvas) UsersListUsersInAccount(progress *task.Progress, searchTerm *string, enrollmentType *string, sort *UsersListUsersInAccountSort, order *UsersListUsersInAccountOrder) ([]User, error) {
	endpoint := fmt.Sprintf("accounts/self/users")
	params := map[string]interface{}{}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if enrollmentType != nil {
		params["enrollment_type"] = *enrollmentType
	}
	if sort != nil {
		params["sort"] = *sort
	}
	if order != nil {
		params["order"] = *order
	}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersListTheActivityStream API call: Returns the current user's global activity stream, paginated.
func (c *Canvas) UsersListTheActivityStream(progress *task.Progress, onlyActiveCourses *bool) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if onlyActiveCourses != nil {
		params["only_active_courses"] = *onlyActiveCourses
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersActivityStreamSummary API call: Returns a summary of the current user's global activity stream.
func (c *Canvas) UsersActivityStreamSummary(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersListTheTODOItems API call: A paginated list of the current user's list of todo items.
func (c *Canvas) UsersListTheTODOItems(progress *task.Progress, include *UsersListTheTODOItemsInclude) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersListCountsForTodoItems API call: Counts of different todo items such as the number of assignments needing
// grading as well as the number of assignments needing submitting.
func (c *Canvas) UsersListCountsForTodoItems(progress *task.Progress, include *UsersListCountsForTodoItemsInclude) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersListUpcomingAssignmentsCalendarEvents API call: A paginated list of the current user's upcoming events.
func (c *Canvas) UsersListUpcomingAssignmentsCalendarEvents(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersListMissingSubmissions API call: A paginated list of past-due assignments for which the student does not have a
// submission. The user sending the request must either be the student, an admin or a parent observer using the parent
// app
func (c *Canvas) UsersListMissingSubmissions(progress *task.Progress, userID *interface{}, include *UsersListMissingSubmissionsInclude, filter *UsersListMissingSubmissionsFilter) ([]Assignment, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if userID != nil {
		params["user_id"] = *userID
	}
	if include != nil {
		params["include"] = *include
	}
	if filter != nil {
		params["filter"] = *filter
	}
	responseCtor := func() interface{} {
		return &[]Assignment{}
	}
	var res []Assignment
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Assignment)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersHideAStreamItem API call: Hide the given stream item.
func (c *Canvas) UsersHideAStreamItem(progress *task.Progress, streamItemID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/self/activity_stream/%s", streamItemID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersHideAllStreamItems API call: Hide all stream items for the user
func (c *Canvas) UsersHideAllStreamItems(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/self/activity_stream")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersUploadAFile API call: Upload a file to the user's personal files section. This API endpoint is the first step in
// uploading a file to a user's files. See the {file:file_uploads.html File Upload Documentation} for details on the
// file upload workflow. Note that typically users will only be able to upload files to their own files section. Passing
// a user_id of +self+ is an easy shortcut to specify the current user.
func (c *Canvas) UsersUploadAFile(progress *task.Progress) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersShowUserDetails API call: Shows details for user. Also includes an attribute "permissions", a non-comprehensive
// list of permissions for the user. Example: !!!javascript "permissions": { "can_update_name": true, // Whether the
// user can update their name. "can_update_avatar": false, // Whether the user can update their avatar.
// "limit_parent_app_web_access": false // Whether the user can interact with Canvas web from the Canvas Parent app.
func (c *Canvas) UsersShowUserDetails(progress *task.Progress, include *UsersShowUserDetailsInclude) (*User, error) {
	endpoint := fmt.Sprintf("users/self")
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersCreateAUser API call: Create and return a new user and pseudonym for an account. If you don't have the "Modify
// login details for users" permission, but self-registration is enabled on the account, you can still use this endpoint
// to register new users. Certain fields will be required, and others will be ignored (see below).
func (c *Canvas) UsersCreateAUser(progress *task.Progress, user *string, pseudonym *string, communicationChannel *string, forceValidations *bool, enableSisReactivation *bool, destination *string) (*User, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if user != nil {
		params["user"] = *user
	}
	if pseudonym != nil {
		params["pseudonym"] = *pseudonym
	}
	if communicationChannel != nil {
		params["communication_channel"] = *communicationChannel
	}
	if forceValidations != nil {
		params["force_validations"] = *forceValidations
	}
	if enableSisReactivation != nil {
		params["enable_sis_reactivation"] = *enableSisReactivation
	}
	if destination != nil {
		params["destination"] = *destination
	}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersSelfRegisterAUser API call: Self register and return a new user and pseudonym for an account. If
// self-registration is enabled on the account, you can use this endpoint to self register new users.
func (c *Canvas) UsersSelfRegisterAUser(progress *task.Progress, user *string, pseudonym *string, communicationChannel *string) (*User, error) {
	endpoint := fmt.Sprintf("")
	params := map[string]interface{}{}
	if user != nil {
		params["user"] = *user
	}
	if pseudonym != nil {
		params["pseudonym"] = *pseudonym
	}
	if communicationChannel != nil {
		params["communication_channel"] = *communicationChannel
	}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersUpdateUserSettings API call: Update an existing user's settings.
func (c *Canvas) UsersUpdateUserSettings(progress *task.Progress, manualMarkAsRead *bool, collapseGlobalNav *bool, hideDashcardColorOverlays *bool, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/%s/settings", userID)
	params := map[string]interface{}{}
	if manualMarkAsRead != nil {
		params["manual_mark_as_read"] = *manualMarkAsRead
	}
	if collapseGlobalNav != nil {
		params["collapse_global_nav"] = *collapseGlobalNav
	}
	if hideDashcardColorOverlays != nil {
		params["hide_dashcard_color_overlays"] = *hideDashcardColorOverlays
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersGetCustomColors API call: Returns all custom colors that have been saved for a user.
func (c *Canvas) UsersGetCustomColors(progress *task.Progress, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/%s/colors/", userID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersGetCustomColor API call: Returns the custom colors that have been saved for a user for a given context. The
// asset_string parameter should be in the format 'context_id', for example 'course_42'.
func (c *Canvas) UsersGetCustomColor(progress *task.Progress, userID string, assetString string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/%s/colors/%s", userID, assetString)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersUpdateCustomColor API call: Updates a custom color for a user for a given context.  This allows colors for the
// calendar and elsewhere to be customized on a user basis. The asset string parameter should be in the format
// 'context_id', for example 'course_42'
func (c *Canvas) UsersUpdateCustomColor(progress *task.Progress, hexcode *string, userID string, assetString string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/%s/colors/%s", userID, assetString)
	params := map[string]interface{}{}
	if hexcode != nil {
		params["hexcode"] = *hexcode
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersGetDashboardPositions API call: Returns all dashboard positions that have been saved for a user.
func (c *Canvas) UsersGetDashboardPositions(progress *task.Progress, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/%s/dashboard_positions/", userID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersUpdateDashboardPositions API call: Updates the dashboard positions for a user for a given context.  This allows
// positions for the dashboard cards and elsewhere to be customized on a per user basis. The asset string parameter
// should be in the format 'context_id', for example 'course_42'
func (c *Canvas) UsersUpdateDashboardPositions(progress *task.Progress, userID string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/%s/dashboard_positions/", userID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersEditAUser API call: Modify an existing user. To modify a user's login, see the documentation for logins.
func (c *Canvas) UsersEditAUser(progress *task.Progress, user *string) (*User, error) {
	endpoint := fmt.Sprintf("users/133.json")
	params := map[string]interface{}{}
	if user != nil {
		params["user"] = *user
	}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersMergeUserIntoAnotherUser API call: Merge a user into another user. To merge users, the caller must have
// permissions to manage both users. This should be considered irreversible. This will delete the user and move all the
// data into the destination user. User merge details and caveats: The from_user is the user that was deleted in the
// user_merge process. The destination_user is the user that remains, that is being split. Avatars: When both users have
// avatars, only the destination_users avatar will remain. When one user has an avatar, will it will end up on the
// destination_user. Terms of Use: If either user has accepted terms of use, it will be be left as accepted.
// Communication Channels: All unique communication channels moved to the destination_user. All notification preferences
// are moved to the destination_user. Enrollments: All unique enrollments are moved to the destination_user. When there
// is an enrollment that would end up making it so that a user would be observing themselves, the enrollment is not
// moved over. Everything that is tied to the from_user at the course level relating to the enrollment is also moved to
// the destination_user. Submissions: All submissions are moved to the destination_user. If there are enrollments for
// both users in the same course, we prefer submissions that have grades then submissions that have work in them, and if
// there are no grades or no work, they are not moved. Other notes: Access Tokens are moved on merge. Conversations are
// moved on merge. Favorites are moved on merge. Courses will commonly use LTI tools. LTI tools reference the user with
// IDs that are stored on a user object. Merging users deletes one user and moves all records from the deleted user to
// the destination_user. These IDs are kept for all enrollments, group_membership, and account_users for the from_user
// at the time of the merge. When the destination_user launches an LTI tool from a course that used to be the
// from_user's, it doesn't appear as a new user to the tool provider. Instead it will send the stored ids. The
// destination_user's LTI IDs remain as they were for the courses that they originally had. Future enrollments for the
// destination_user will use the IDs that are on the destination_user object. LTI IDs that are kept and tracked per
// context include lti_context_id, lti_id and uuid. APIs that return the LTI ids will return the one for the context
// that it is called for, except for the user uuid. The user UUID will display the destination_users uuid, and when
// getting the uuid from an api that is in a context that was recorded from a merge event, an additional attribute is
// added as past_uuid. When finding users by SIS ids in different accounts the destination_account_id is required. The
// account can also be identified by passing the domain in destination_account_id.
func (c *Canvas) UsersMergeUserIntoAnotherUser(progress *task.Progress, userID string, destinationAccountID string, destinationUserID string) (*User, error) {
	endpoint := fmt.Sprintf("users/%s/merge_into/accounts/%s/users/%s", userID, destinationAccountID, destinationUserID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &User{}
	}
	var res *User
	callback := func(obj interface{}) error {
		res = obj.(*User)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersSplitMergedUsersIntoSeparateUsers API call: Merged users cannot be fully restored to their previous state, but
// this will attempt to split as much as possible to the previous state. To split a merged user, the caller must have
// permissions to manage all of the users logins. If there are multiple users that have been merged into one user it
// will split each merge into a separate user. A split can only happen within 180 days of a user merge. A user merge
// deletes the previous user and may be permanently deleted. In this scenario we create a new user object and proceed to
// move as much as possible to the new user. The user object will not have preserved the name or settings from the
// previous user. Some items may have been deleted during a user_merge that cannot be restored, and/or the data has
// become stale because of other changes to the objects since the time of the user_merge. Split users details and
// caveats: The from_user is the user that was deleted in the user_merge process. The destination_user is the user that
// remains, that is being split. Avatars: When both users had avatars, both will be remain. When from_user had an avatar
// and destination_user did not have an avatar, the destination_user's avatar will be deleted if it still matches what
// was there are the time of the merge. If the destination_user's avatar was changed at anytime after the merge, it will
// remain on the destination user. If the from_user had an avatar it will be there after split. Terms of Use: If
// from_user had not accepted terms of use, they will be prompted again to accept terms of use after the split. If the
// destination_user had not accepted terms of use, hey will be prompted again to accept terms of use after the split. If
// neither user had accepted the terms of use, but since the time of the merge had accepted, both will be prompted to
// accept terms of use. If both had accepted terms of use, this will remain. Communication Channels: All communication
// channels are restored to what they were prior to the merge. If a communication channel was added after the merge, it
// will remain on the destination_user. Notification preferences remain with the communication channels. Enrollments:
// All enrollments from the time of the merge will be moved back to where they were. Enrollments created since the time
// of the merge that were created by sis_import will go to the user that owns that sis_id used for the import. Other new
// enrollments will remain on the destination_user. Everything that is tied to the destination_user at the course level
// relating to an enrollment is moved to the from_user. When both users are in the same course prior to merge this can
// cause some unexpected items to move. Submissions: Unlike other items tied to a course, submissions are explicitly
// recorded to avoid problems with grades. All submissions were moved are restored to the spot prior to merge. All
// submission that were created in a course that was moved in enrollments are moved over to the from_user. Other notes:
// Access Tokens are moved back on split. Conversations are moved back on split. Favorites that existing at the time of
// merge are moved back on split. LTI ids are restored to how they were prior to merge.
func (c *Canvas) UsersSplitMergedUsersIntoSeparateUsers(progress *task.Progress, userID string) ([]User, error) {
	endpoint := fmt.Sprintf("users/%s/split", userID)
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]User{}
	}
	var res []User
	callback := func(obj interface{}) error {
		arr := *obj.(*[]User)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersGetAPandataEventsJwtTokenAndItsExpirationDate API call: Returns a jwt auth and props token that can be used to
// send events to Pandata. NOTE: This is currently only available to the mobile developer keys.
func (c *Canvas) UsersGetAPandataEventsJwtTokenAndItsExpirationDate(progress *task.Progress, appKey *string) (*map[string]interface{}, error) {
	endpoint := fmt.Sprintf("users/self/pandata_events_token")
	params := map[string]interface{}{}
	if appKey != nil {
		params["app_key"] = *appKey
	}
	responseCtor := func() interface{} {
		return &map[string]interface{}{}
	}
	var res *map[string]interface{}
	callback := func(obj interface{}) error {
		res = obj.(*map[string]interface{})
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// UsersGetAUsersMostRecentlyGradedSubmissions API call
func (c *Canvas) UsersGetAUsersMostRecentlyGradedSubmissions(progress *task.Progress, include *UsersGetAUsersMostRecentlyGradedSubmissionsInclude, onlyCurrentEnrollments *bool, onlyPublishedAssignments *bool, userID string) ([]Submission, error) {
	endpoint := fmt.Sprintf("users/%s/graded_submissions", userID)
	params := map[string]interface{}{}
	if include != nil {
		params["include"] = *include
	}
	if onlyCurrentEnrollments != nil {
		params["only_current_enrollments"] = *onlyCurrentEnrollments
	}
	if onlyPublishedAssignments != nil {
		params["only_published_assignments"] = *onlyPublishedAssignments
	}
	responseCtor := func() interface{} {
		return &[]Submission{}
	}
	var res []Submission
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Submission)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesShowFrontPage API call: Retrieve the content of the front page
func (c *Canvas) PagesShowFrontPage(progress *task.Progress) (*Page, error) {
	endpoint := fmt.Sprintf("courses/123/front_page")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Page{}
	}
	var res *Page
	callback := func(obj interface{}) error {
		res = obj.(*Page)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesDuplicatePage API call: Duplicate a wiki page
func (c *Canvas) PagesDuplicatePage(progress *task.Progress) (*Page, error) {
	endpoint := fmt.Sprintf("courses/123/pages/14/duplicate")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Page{}
	}
	var res *Page
	callback := func(obj interface{}) error {
		res = obj.(*Page)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesUpdateCreateFrontPage API call: Update the title or contents of the front page
func (c *Canvas) PagesUpdateCreateFrontPage(progress *task.Progress, wikiPage *string) (*Page, error) {
	endpoint := fmt.Sprintf("courses/123/front_page")
	params := map[string]interface{}{}
	if wikiPage != nil {
		params["wiki_page"] = *wikiPage
	}
	responseCtor := func() interface{} {
		return &Page{}
	}
	var res *Page
	callback := func(obj interface{}) error {
		res = obj.(*Page)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesListPages API call: A paginated list of the wiki pages associated with a course or group
func (c *Canvas) PagesListPages(progress *task.Progress, sort *PagesListPagesSort, order *PagesListPagesOrder, searchTerm *string, published *bool) ([]Page, error) {
	endpoint := fmt.Sprintf("courses/123/pages")
	params := map[string]interface{}{}
	if sort != nil {
		params["sort"] = *sort
	}
	if order != nil {
		params["order"] = *order
	}
	if searchTerm != nil {
		params["search_term"] = *searchTerm
	}
	if published != nil {
		params["published"] = *published
	}
	responseCtor := func() interface{} {
		return &[]Page{}
	}
	var res []Page
	callback := func(obj interface{}) error {
		arr := *obj.(*[]Page)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesCreatePage API call: Create a new wiki page
func (c *Canvas) PagesCreatePage(progress *task.Progress, wikiPage *string) (*Page, error) {
	endpoint := fmt.Sprintf("courses/123/pages")
	params := map[string]interface{}{}
	if wikiPage != nil {
		params["wiki_page"] = *wikiPage
	}
	responseCtor := func() interface{} {
		return &Page{}
	}
	var res *Page
	callback := func(obj interface{}) error {
		res = obj.(*Page)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesShowPage API call: Retrieve the content of a wiki page
func (c *Canvas) PagesShowPage(progress *task.Progress) (*Page, error) {
	endpoint := fmt.Sprintf("courses/123/pages/my-page-url")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Page{}
	}
	var res *Page
	callback := func(obj interface{}) error {
		res = obj.(*Page)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesUpdateCreatePage API call: Update the title or contents of a wiki page
func (c *Canvas) PagesUpdateCreatePage(progress *task.Progress, wikiPage *string) (*Page, error) {
	endpoint := fmt.Sprintf("courses/123/pages/the-page-url")
	params := map[string]interface{}{}
	if wikiPage != nil {
		params["wiki_page"] = *wikiPage
	}
	responseCtor := func() interface{} {
		return &Page{}
	}
	var res *Page
	callback := func(obj interface{}) error {
		res = obj.(*Page)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesDeletePage API call: Delete a wiki page
func (c *Canvas) PagesDeletePage(progress *task.Progress) (*Page, error) {
	endpoint := fmt.Sprintf("courses/123/pages/the-page-url")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &Page{}
	}
	var res *Page
	callback := func(obj interface{}) error {
		res = obj.(*Page)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesListRevisions API call: A paginated list of the revisions of a page. Callers must have update rights on the page
// in order to see page history.
func (c *Canvas) PagesListRevisions(progress *task.Progress) ([]PageRevision, error) {
	endpoint := fmt.Sprintf("courses/123/pages/the-page-url/revisions")
	params := map[string]interface{}{}
	responseCtor := func() interface{} {
		return &[]PageRevision{}
	}
	var res []PageRevision
	callback := func(obj interface{}) error {
		arr := *obj.(*[]PageRevision)
		res = append(res, arr...)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesShowRevision API call: Retrieve the metadata and optionally content of a revision of the page. Note that
// retrieving historic versions of pages requires edit rights.
func (c *Canvas) PagesShowRevision(progress *task.Progress, summary *bool) (*PageRevision, error) {
	endpoint := fmt.Sprintf("courses/123/pages/the-page-url/revisions/4")
	params := map[string]interface{}{}
	if summary != nil {
		params["summary"] = *summary
	}
	responseCtor := func() interface{} {
		return &PageRevision{}
	}
	var res *PageRevision
	callback := func(obj interface{}) error {
		res = obj.(*PageRevision)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}

// PagesRevertToRevision API call: Revert a page to a prior revision.
func (c *Canvas) PagesRevertToRevision(progress *task.Progress, revisionID *int) (*PageRevision, error) {
	endpoint := fmt.Sprintf("courses/123/pages/the-page-url/revisions/6")
	params := map[string]interface{}{}
	if revisionID != nil {
		params["revision_id"] = *revisionID
	}
	responseCtor := func() interface{} {
		return &PageRevision{}
	}
	var res *PageRevision
	callback := func(obj interface{}) error {
		res = obj.(*PageRevision)
		return nil
	}
	if err := c.Request(endpoint, params, progress, responseCtor, callback); err != nil {
		return nil, err
	}
	return res, nil
}
