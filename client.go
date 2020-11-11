package eventstream

import (
	"context"
	"time"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
	"github.com/protsack-stephan/mediawiki-eventstream-client/listeners"
)

const baseURL = "https://stream.wikimedia.org/v2/stream/"

const reconnectTime = time.Second * 1

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
		reconnectTime,
		pageDeleteURL,
		pageMoveURL,
		revisionCreateURL,
		revisionScoreURL,
		revisionVisibilityChangeURL,
	}
}

// Client request client
type Client struct {
	ReconnectTime               time.Duration
	PageDeleteURL               string
	PageMoveURL                 string
	RevisionCreateURL           string
	RevisionScoreURL            string
	RevisionVisibilityChangeURL string
}

// PageDelete connect to page delete stream
func (cl *Client) PageDelete(ctx context.Context, since time.Time, handler func(evt *events.PageDelete)) *Stream {
	errors := make(chan error)
	return NewStream(cl.ReconnectTime, errors, func() error {
		return listeners.PageDelete(ctx, cl.PageDeleteURL, since, handler, errors)
	})
}

// PageMove connect to page move stream
func (cl *Client) PageMove(ctx context.Context, since time.Time, handler func(evt *events.PageMove)) *Stream {
	errors := make(chan error)
	return NewStream(cl.ReconnectTime, errors, func() error {
		return listeners.PageMove(ctx, cl.PageMoveURL, since, handler, errors)
	})
}

// RevisionCreate connect to revision create stream
func (cl *Client) RevisionCreate(ctx context.Context, since time.Time, handler func(evt *events.RevisionCreate)) *Stream {
	errors := make(chan error)
	return NewStream(cl.ReconnectTime, errors, func() error {
		return listeners.RevisionCreate(ctx, cl.RevisionCreateURL, since, handler, errors)
	})
}

// RevisionScore connect to revision score stream
func (cl *Client) RevisionScore(ctx context.Context, since time.Time, handler func(evt *events.RevisionScore)) *Stream {
	errors := make(chan error)
	return NewStream(cl.ReconnectTime, errors, func() error {
		return listeners.RevisionScore(ctx, cl.RevisionScoreURL, since, handler, errors)
	})
}

// RevisionVisibilityChange connext to revision visibility change stream
func (cl *Client) RevisionVisibilityChange(ctx context.Context, since time.Time, handler func(evt *events.RevisionVisibilityChange)) *Stream {
	errors := make(chan error)
	return NewStream(cl.ReconnectTime, errors, func() error {
		return listeners.RevisionVisibilityChange(ctx, cl.RevisionVisibilityChangeURL, since, handler, errors)
	})
}
