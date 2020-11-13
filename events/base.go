package events

import (
	"time"

	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
)

type baseEvent struct {
	ID   []subscriber.Meta
	Data struct {
		baseData
	}
}

type baseData struct {
	Schema         string    `json:"schema"`
	Database       string    `json:"database"`
	Meta           Meta      `json:"meta"`
	Performer      Performer `json:"performer"`
	PageID         int       `json:"page_id"`
	PageTitle      string    `json:"page_title"`
	PageNamespace  int       `json:"page_namespace"`
	PageIsRedirect bool      `json:"page_is_redirect"`
	RevID          int       `json:"rev_id"`
}

// Meta event meta data
type Meta struct {
	URI       string    `json:"uri"`
	RequestID string    `json:"request_id"`
	ID        string    `json:"id"`
	Dt        time.Time `json:"dt"`
	Domain    string    `json:"domain"`
	Stream    string    `json:"stream"`
	Topic     string    `json:"topic"`
	Partition int       `json:"partition"`
	Offset    int       `json:"offset"`
}

// Performer user that is responsible for event
type Performer struct {
	UserText           string    `json:"user_text"`
	UserGroups         []string  `json:"user_groups"`
	UserIsBot          bool      `json:"user_is_bot"`
	UserID             int       `json:"user_id"`
	UserRegistrationDt time.Time `json:"user_registration_dt"`
	UserEditCount      int       `json:"user_edit_count"`
}
