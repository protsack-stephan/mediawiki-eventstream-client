package eventstream

import (
	"context"
	"net/http"
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
		new(http.Client),
		backoffTime,
		&Options{
			pageDeleteURL,
			pageMoveURL,
			revisionCreateURL,
			revisionScoreURL,
			revisionVisibilityChangeURL,
		},
	}
}

// Client request client
type Client struct {
	httpClient  *http.Client
	backoffTime time.Duration
	options     *Options
}

// PageDelete connect to page delete stream
func (cl *Client) PageDelete(ctx context.Context, since time.Time, handler func(evt *PageDelete)) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.options.PageDeleteURL, store.getSince(), func(msg *Event) {
			evt := new(PageDelete)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}

// PageMove connect to page move stream
func (cl *Client) PageMove(ctx context.Context, since time.Time, handler func(evt *PageMove)) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.options.PageMoveURL, store.getSince(), func(msg *Event) {
			evt := new(PageMove)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}

// RevisionCreate connect to revision create stream
func (cl *Client) RevisionCreate(ctx context.Context, since time.Time, handler func(evt *RevisionCreate)) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.options.RevisionCreateURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionCreate)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}

// RevisionScore connect to revision score stream
func (cl *Client) RevisionScore(ctx context.Context, since time.Time, handler func(evt *RevisionScore)) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.options.RevisionScoreURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionScore)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}

// RevisionVisibilityChange connext to revision visibility change stream
func (cl *Client) RevisionVisibilityChange(ctx context.Context, since time.Time, handler func(evt *RevisionVisibilityChange)) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.options.RevisionVisibilityChangeURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionVisibilityChange)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}
