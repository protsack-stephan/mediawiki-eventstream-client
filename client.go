package eventstream

import (
	"context"
	"encoding/json"
	"time"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
	"github.com/protsack-stephan/mediawiki-eventstream-client/subscriber"
)

const baseURL = "https://stream.wikimedia.org/v2/stream/"

const backoffTime = time.Second * 1

// All the available streams
const (
	pageDeleteURL               = baseURL + "page-delete"
	pageMoveURL                 = baseURL + "page-move"
	revisionCreateURL           = baseURL + "revision-create"
	revisionScoreURL            = baseURL + "revision-score"
	revisionVisibilityChangeURL = baseURL + "revision-visibility-change"
)

// NewClient creating new connection client
func NewClient() *Client {
	return &Client{
		backoffTime,
		pageDeleteURL,
		pageMoveURL,
		revisionCreateURL,
		revisionScoreURL,
		revisionVisibilityChangeURL,
	}
}

// Client request client
type Client struct {
	BackoffTime                 time.Duration
	PageDeleteURL               string
	PageMoveURL                 string
	RevisionCreateURL           string
	RevisionScoreURL            string
	RevisionVisibilityChangeURL string
}

// PageDelete connect to page delete stream
func (cl *Client) PageDelete(ctx context.Context, since time.Time, handler func(evt *events.PageDelete)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscriber.Subscribe(ctx, cl.PageDeleteURL, store.getSince(), func(msg *subscriber.Event) {
			evt := new(events.PageDelete)
			evt.ID = msg.ID
			err := json.Unmarshal(msg.Data, &evt.Data)

			if err == nil {
				store.setSince(evt.Data.Meta.Dt)
				handler(evt)
			} else {
				store.setError(err)
			}
		})
	})
}

// PageMove connect to page move stream
func (cl *Client) PageMove(ctx context.Context, since time.Time, handler func(evt *events.PageMove)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscriber.Subscribe(ctx, cl.PageMoveURL, store.getSince(), func(msg *subscriber.Event) {
			evt := new(events.PageMove)
			evt.ID = msg.ID
			err := json.Unmarshal(msg.Data, &evt.Data)

			if err == nil {
				store.setSince(evt.Data.Meta.Dt)
				handler(evt)
			} else {
				store.setError(err)
			}
		})
	})
}

// RevisionCreate connect to revision create stream
func (cl *Client) RevisionCreate(ctx context.Context, since time.Time, handler func(evt *events.RevisionCreate)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscriber.Subscribe(ctx, cl.RevisionCreateURL, store.getSince(), func(msg *subscriber.Event) {
			evt := new(events.RevisionCreate)
			evt.ID = msg.ID
			err := json.Unmarshal(msg.Data, &evt.Data)

			if err == nil {
				store.setSince(evt.Data.Meta.Dt)
				handler(evt)
			} else {
				store.setError(err)
			}
		})
	})
}

// RevisionScore connect to revision score stream
func (cl *Client) RevisionScore(ctx context.Context, since time.Time, handler func(evt *events.RevisionScore)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscriber.Subscribe(ctx, cl.RevisionScoreURL, store.getSince(), func(msg *subscriber.Event) {
			evt := new(events.RevisionScore)
			evt.ID = msg.ID
			err := json.Unmarshal(msg.Data, &evt.Data)

			if err == nil {
				store.setSince(evt.Data.Meta.Dt)
				handler(evt)
			} else {
				store.setError(err)
			}
		})
	})
}

// RevisionVisibilityChange connext to revision visibility change stream
func (cl *Client) RevisionVisibilityChange(ctx context.Context, since time.Time, handler func(evt *events.RevisionVisibilityChange)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscriber.Subscribe(ctx, cl.RevisionVisibilityChangeURL, store.getSince(), func(msg *subscriber.Event) {
			evt := new(events.RevisionVisibilityChange)
			evt.ID = msg.ID
			err := json.Unmarshal(msg.Data, &evt.Data)

			if err == nil {
				store.setSince(evt.Data.Meta.Dt)
				handler(evt)
			} else {
				store.setError(err)
			}
		})
	})
}
