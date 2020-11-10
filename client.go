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
		ReconnectTime:               reconnectTime,
		PageDeleteURL:               pageDeleteURL,
		PageMoveURL:                 pageMoveURL,
		RevisionCreateURL:           revisionCreateURL,
		RevisionScoreURL:            revisionScoreURL,
		RevisionVisibilityChangeURL: revisionVisibilityChangeURL,
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
func (cl *Client) PageDelete(ctx context.Context, since time.Time, handler func(evt *events.PageDelete)) error {
	return listeners.PageDelete(ctx, cl.PageDeleteURL, since, handler)
}

// PageMove connect to page move stream
func (cl *Client) PageMove(ctx context.Context, since time.Time, handler func(evt *events.PageMove)) error {
	return listeners.PageMove(ctx, cl.PageMoveURL, since, handler)
}

// RevisionCreate connect to revision create stream
func (cl *Client) RevisionCreate(ctx context.Context, since time.Time, handler func(evt *events.RevisionCreate)) error {
	return listeners.RevisionCreate(ctx, cl.RevisionCreateURL, since, handler)
}

// RevisionCreateKeepAlive connect to revision create stream with instant reconnect
func (cl *Client) RevisionCreateKeepAlive(ctx context.Context, since time.Time, handler func(evt *events.RevisionCreate)) chan error {
	errors := make(chan error)

	go cl.keepAlive(func() error {
		return cl.RevisionCreate(ctx, since, handler)
	}, errors)

	return errors
}

// RevisionScore connect to revision score stream
func (cl *Client) RevisionScore(ctx context.Context, since time.Time, handler func(evt *events.RevisionScore)) error {
	return listeners.RevisionScore(ctx, cl.RevisionScoreURL, since, handler)
}

// RevisionVisibilityChange connext to revision visibility change stream
func (cl *Client) RevisionVisibilityChange(ctx context.Context, since time.Time, handler func(evt *events.RevisionVisibilityChange)) error {
	return listeners.RevisionVisibilityChange(ctx, cl.RevisionVisibilityChangeURL, since, handler)
}

func (cl *Client) keepAlive(cb func() error, errors chan error) {
	if cl.ReconnectTime == 0 {
		cl.ReconnectTime = reconnectTime
	}

	for {
		err := cb()

		errors <- err
		if err == context.Canceled {
			close(errors)
			break
		} else {
			time.Sleep(cl.ReconnectTime)
		}
	}
}
