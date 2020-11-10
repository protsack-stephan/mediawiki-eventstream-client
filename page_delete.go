package eventstream

import (
	"context"
	"encoding/json"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
)

func pageDelete(ctx context.Context, url string, handler func(evt *events.PageDelete)) error {
	return subscriber.Subscribe(ctx, url, func(msg *subscriber.Event) {
		evt := new(events.PageDelete)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}

// PageDelete stream from mediawiki
func PageDelete(ctx context.Context, handler func(evt *events.PageDelete)) error {
	return pageDelete(ctx, pageDeleteURL, handler)
}
