package model

// User object
type User struct {
	AvatarURL     string      `json:"avatar_url"`
	Bio           string      `json:"bio"`
	Email         string      `json:"email"`
	Enrollments   interface{} `json:"enrollments"`
	ID            int         `json:"id"`
	IntegrationID string      `json:"integration_id"`
	LastLogin     string      `json:"last_login"`
	Locale        string      `json:"locale"`
	LoginID       string      `json:"login_id"`
	Name          string      `json:"name"`
	ShortName     string      `json:"short_name"`
	SisImportID   int         `json:"sis_import_id"`
	SisUserID     string      `json:"sis_user_id"`
	SortableName  string      `json:"sortable_name"`
	TimeZone      string      `json:"time_zone"`
}
