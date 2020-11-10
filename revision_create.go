package eventstream

import (
	"context"
	"encoding/json"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
)

func revisionCreate(ctx context.Context, url string, handler func(evt *events.RevisionCreate)) error {
	return subscriber.Subscribe(ctx, url, func(msg *subscriber.Event) {
		evt := new(events.RevisionCreate)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}

// RevisionCreate stream from mediawiki
func RevisionCreate(ctx context.Context, handler func(evt *events.RevisionCreate)) error {
	return revisionCreate(ctx, revisionCreateURL, handler)
}
