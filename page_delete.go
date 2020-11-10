package eventstream

import (
	"context"
	"encoding/json"

	"github.com/protsack-stephan/mediawiki-eventstream-client/client"
	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

// PageDelete stream from mediawiki
func PageDelete(ctx context.Context, handler func(evt *events.PageDelete)) error {
	return client.Subscribe(ctx, pageDelete, func(msg *client.Event) {
		evt := new(events.PageDelete)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}
