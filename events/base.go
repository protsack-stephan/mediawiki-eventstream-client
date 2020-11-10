package events

import "github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"

type baseEvent struct {
	ID   []subscriber.Info
	Data interface{}
}
