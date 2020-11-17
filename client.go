package eventstream

import (
	"context"
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
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}

// PageMove connect to page move stream
func (cl *Client) PageMove(ctx context.Context, since time.Time, handler func(evt *PageMove)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.PageMoveURL, store.getSince(), func(msg *Event) {
			evt := new(PageMove)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}

// RevisionCreate connect to revision create stream
func (cl *Client) RevisionCreate(ctx context.Context, since time.Time, handler func(evt *RevisionCreate)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.RevisionCreateURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionCreate)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}

// RevisionScore connect to revision score stream
func (cl *Client) RevisionScore(ctx context.Context, since time.Time, handler func(evt *RevisionScore)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.RevisionScoreURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionScore)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}

// RevisionVisibilityChange connext to revision visibility change stream
func (cl *Client) RevisionVisibilityChange(ctx context.Context, since time.Time, handler func(evt *RevisionVisibilityChange)) *Stream {
	store := newStorage(since, cl.BackoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.RevisionVisibilityChangeURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionVisibilityChange)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}
