package eventstream

import (
	"context"
	"encoding/json"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
)

func revisionScore(ctx context.Context, url string, handler func(evt *events.RevisionScore)) error {
	return subscriber.Subscribe(ctx, url, func(msg *subscriber.Event) {
		evt := new(events.RevisionScore)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}

// RevisionScore stream from mediawiki
func RevisionScore(ctx context.Context, handler func(evt *events.RevisionScore)) error {
	return revisionScore(ctx, revisionScoreURL, handler)
}
