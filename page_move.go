package eventstream

import (
	"context"
	"encoding/json"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
)

func pageMove(ctx context.Context, url string, handler func(evt *events.PageMove)) error {
	return subscriber.Subscribe(ctx, url, func(msg *subscriber.Event) {
		evt := new(events.PageMove)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}

// PageMove stream from mediawiki
func PageMove(ctx context.Context, handler func(evt *events.PageMove)) error {
	return pageMove(ctx, pageMoveURL, handler)
}
