package eventstream

import (
	"encoding/json"

	"github.com/protsack-stephan/mediawiki-eventstream-client/client"
	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

// PageMove stream from mediawiki
func PageMove(handler func(evt *events.PageMove)) error {
	return client.Subscribe(pageMove, func(msg *client.Event) {
		evt := new(events.PageMove)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}
