package listeners

import (
	"context"
	"encoding/json"
	"time"

	"github.com/protsack-stephan/mediawiki-eventstream-client/utils"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
)

// PageDelete event listener
func PageDelete(ctx context.Context, url string, since time.Time, handler func(evt *events.PageDelete), errors ...chan error) error {
	return subscriber.Subscribe(ctx, utils.FormatURL(url, since), func(msg *subscriber.Event) {
		evt := new(events.PageDelete)
		evt.ID = msg.ID

		err := json.Unmarshal(msg.Data, &evt.Data)
		if err == nil {
			handler(evt)
		} else {
			for _, errChanel := range errors {
				errChanel <- err
			}
		}
	})
}
