package eventstream

import (
	"context"
	"net/http"
	"time"
)

const url = "https://stream.wikimedia.org"

const backoffTime = time.Second * 1

// All the available streams
const (
	pageDeleteURL               = "/v2/stream/page-delete"
	pageMoveURL                 = "/v2/stream/page-move"
	revisionCreateURL           = "/v2/stream/revision-create"
	revisionScoreURL            = "/v2/stream/revision-score"
	revisionVisibilityChangeURL = "/v2/stream/mediawiki.revision-visibility-change"
)

// NewClient creating new connection client
func NewClient() *Client {
	return &Client{
		url,
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
	url         string
	httpClient  *http.Client
	backoffTime time.Duration
	options     *Options
}

// PageDelete connect to page delete stream
func (cl *Client) PageDelete(ctx context.Context, since time.Time, handler func(evt *PageDelete)) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageDeleteURL, store.getSince(), func(msg *Event) {
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
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.PageMoveURL, store.getSince(), func(msg *Event) {
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
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.RevisionCreateURL, store.getSince(), func(msg *Event) {
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
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.RevisionScoreURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionScore)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}

// RevisionVisibilityChange connect to revision visibility change stream
func (cl *Client) RevisionVisibilityChange(ctx context.Context, since time.Time, handler func(evt *RevisionVisibilityChange)) *Stream {
	store := newStorage(since, cl.backoffTime)

	return NewStream(store, func(since time.Time) error {
		return subscribe(ctx, cl.httpClient, cl.url+cl.options.RevisionVisibilityChangeURL, store.getSince(), func(msg *Event) {
			evt := new(RevisionVisibilityChange)
			parseSchema(evt, msg, store)
			handler(evt)
		})
	})
}
