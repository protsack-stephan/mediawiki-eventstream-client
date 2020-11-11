package events

import "time"

// RevisionVisibilityChange event scheme struct
type RevisionVisibilityChange struct {
	baseEvent
	Data struct {
		Schema           string    `json:"$schema"`
		Meta             Meta      `json:"meta"`
		Database         string    `json:"database"`
		PageID           int       `json:"page_id"`
		PageTitle        string    `json:"page_title"`
		PageNamespace    int       `json:"page_namespace"`
		RevID            int       `json:"rev_id"`
		RevTimestamp     time.Time `json:"rev_timestamp"`
		RevSha1          string    `json:"rev_sha1"`
		RevMinorEdit     bool      `json:"rev_minor_edit"`
		RevLen           int       `json:"rev_len"`
		RevContentModel  string    `json:"rev_content_model"`
		RevContentFormat string    `json:"rev_content_format"`
		Performer        Performer `json:"performer"`
		PageIsRedirect   bool      `json:"page_is_redirect"`
		Comment          string    `json:"comment"`
		Parsedcomment    string    `json:"parsedcomment"`
		RevParentID      int       `json:"rev_parent_id"`
		Visibility       struct {
			Text    bool `json:"text"`
			User    bool `json:"user"`
			Comment bool `json:"comment"`
		} `json:"visibility"`
		PriorState struct {
			Visibility struct {
				Text    bool `json:"text"`
				User    bool `json:"user"`
				Comment bool `json:"comment"`
			} `json:"visibility"`
		} `json:"prior_state"`
	}
}
