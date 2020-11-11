package events

import (
	"time"

	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
)

type baseEvent struct {
	ID   []subscriber.Info
	Data interface{}
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
