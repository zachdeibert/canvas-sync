package model

// File object
type File struct {
    Contenttype     string      `json:"content-type"`
    CreatedAt       string      `json:"created_at"`
    DisplayName     string      `json:"display_name"`
    Filename        string      `json:"filename"`
    FolderID        int         `json:"folder_id"`
    Hidden          bool        `json:"hidden"`
    HiddenForUser   bool        `json:"hidden_for_user"`
    ID              int         `json:"id"`
    LockAt          string      `json:"lock_at"`
    LockExplanation string      `json:"lock_explanation"`
    LockInfo        interface{} `json:"lock_info"`
    Locked          bool        `json:"locked"`
    LockedForUser   bool        `json:"locked_for_user"`
    MediaEntryID    string      `json:"media_entry_id"`
    MimeClass       string      `json:"mime_class"`
    ModifiedAt      string      `json:"modified_at"`
    PreviewURL      interface{} `json:"preview_url"`
    Size            int         `json:"size"`
    ThumbnailURL    interface{} `json:"thumbnail_url"`
    URL             string      `json:"url"`
    UUID            string      `json:"uuid"`
    UnlockAt        string      `json:"unlock_at"`
    UpdatedAt       string      `json:"updated_at"`
}
