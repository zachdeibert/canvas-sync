package canvas

import "time"

// DiscussionTopicFullViewUser object
type DiscussionTopicFullViewUser struct {
	ID          int    `json:"id"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
}

// DiscussionTopicFullViewEntry object
type DiscussionTopicFullViewEntry struct {
	ID        int                            `json:"id"`
	UserID    int                            `json:"user_id"`
	Message   string                         `json:"message"`
	Replies   []DiscussionTopicFullViewEntry `json:"replies"`
	UpdatedAt time.Time                      `json:"updated_at"`
}

// DiscussionTopicFullView object
type DiscussionTopicFullView struct {
	Participants  []DiscussionTopicFullViewUser  `json:"participants"`
	UnreadEntries []int                          `json:"unread_entries"`
	EntryRatings  map[int]int                    `json:"entry_ratings"`
	ForcedEntries []int                          `json:"forced_entries"`
	View          []DiscussionTopicFullViewEntry `json:"view"`
}
