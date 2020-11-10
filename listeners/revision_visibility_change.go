package listeners

import (
	"context"
	"encoding/json"
	"time"

	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
	"github.com/protsack-stephan/mediawiki-eventstream-client/utils"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

// RevisionVisibilityChange event listener
func RevisionVisibilityChange(ctx context.Context, url string, since time.Time, handler func(evt *events.RevisionVisibilityChange)) error {
	return subscriber.Subscribe(ctx, utils.FormatURL(url, since), func(msg *subscriber.Event) {
		evt := new(events.RevisionVisibilityChange)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		}
	})
}
