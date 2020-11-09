package events

import "github.com/protsack-stephan/mediawiki-eventstream-client/client"

type baseEvent struct {
	ID   []client.Info
	Data interface{}
}
