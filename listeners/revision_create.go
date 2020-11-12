package listeners

import (
	"context"
	"encoding/json"
	"time"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
	"github.com/protsack-stephan/mediawiki-eventstream-client/utils"
)

// RevisionCreate event listener
func RevisionCreate(ctx context.Context, url string, since time.Time, handler func(evt *events.RevisionCreate), errors ...chan error) error {
	return subscriber.Subscribe(ctx, utils.FormatURL(url, since), func(msg *subscriber.Event) {
		evt := new(events.RevisionCreate)
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
