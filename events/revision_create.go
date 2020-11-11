package events

import "time"

// RevisionCreate event scheme struct
type RevisionCreate struct {
	baseEvent
	Data struct {
		Schema            string    `json:"$schema"`
		Meta              Meta      `json:"meta"`
		Database          string    `json:"database"`
		PageID            int       `json:"page_id"`
		PageTitle         string    `json:"page_title"`
		PageNamespace     int       `json:"page_namespace"`
		RevID             int       `json:"rev_id"`
		RevTimestamp      time.Time `json:"rev_timestamp"`
		RevSha1           string    `json:"rev_sha1"`
		RevMinorEdit      bool      `json:"rev_minor_edit"`
		RevLen            int       `json:"rev_len"`
		RevContentModel   string    `json:"rev_content_model"`
		RevContentFormat  string    `json:"rev_content_format"`
		Performer         Performer `json:"performer"`
		PageIsRedirect    bool      `json:"page_is_redirect"`
		Comment           string    `json:"comment"`
		ChronologyID      string    `json:"chronology_id"`
		Parsedcomment     string    `json:"parsedcomment"`
		RevParentID       int       `json:"rev_parent_id"`
		RevContentChanged bool      `json:"rev_content_changed"`
		RevIsRevert       bool      `json:"rev_is_revert"`
	}
}
