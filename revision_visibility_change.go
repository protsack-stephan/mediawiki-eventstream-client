package eventstream

import (
	"context"
	"encoding/json"

	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

func revisionVisibilityChange(ctx context.Context, url string, handler func(evt *events.RevisionVisibilityChange)) error {
	return subscriber.Subscribe(ctx, url, func(msg *subscriber.Event) {
		evt := new(events.RevisionVisibilityChange)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}

// RevisionVisibilityChange stream from mediawiki
func RevisionVisibilityChange(ctx context.Context, url string, handler func(evt *events.RevisionVisibilityChange)) error {
	return revisionVisibilityChange(ctx, revisionVisibilityChangeURL, handler)
}
