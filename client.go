package eventstream

import (
	"context"
	"encoding/json"
	"time"
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
func (cl *Client) PageDelete(ctx context.Context, since time.Time, handler func(evt *PageDelete)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.PageDeleteURL, store.getSince(), func(msg *Event) {
			evt := new(PageDelete)
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
func (cl *Client) PageMove(ctx context.Context, since time.Time, handler func(evt *PageMove)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.PageMoveURL, store.getSince(), func(msg *Event) {
			evt := new(PageMove)
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
func (cl *Client) RevisionCreate(ctx context.Context, since time.Time, handler func(evt *RevisionCreate)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.RevisionCreateURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionCreate)
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
func (cl *Client) RevisionScore(ctx context.Context, since time.Time, handler func(evt *RevisionScore)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.RevisionScoreURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionScore)
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
func (cl *Client) RevisionVisibilityChange(ctx context.Context, since time.Time, handler func(evt *RevisionVisibilityChange)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.RevisionVisibilityChangeURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionVisibilityChange)
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
