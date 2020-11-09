package eventstream

import (
	"encoding/json"

	"github.com/protsack-stephan/mediawiki-eventstream-client/client"
	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

// RevisionScore stream from mediawiki
func RevisionScore(handler func(evt *events.RevisionScore)) error {
	return client.Subscribe(revisionScore, func(msg *client.Event) {
		evt := new(events.RevisionScore)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}
