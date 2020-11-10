package eventstream

import (
	"context"
	"encoding/json"

	"github.com/protsack-stephan/mediawiki-eventstream-client/client"
	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

// RevisionCreate stream from mediawiki
func RevisionCreate(ctx context.Context, handler func(evt *events.RevisionCreate)) error {
	return client.Subscribe(ctx, revisionCreate, func(msg *client.Event) {
		evt := new(events.RevisionCreate)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}
