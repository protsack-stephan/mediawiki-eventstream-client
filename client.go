package eventstream

import (
	"context"

	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

const baseURL = "https://stream.wikimedia.org/v2/stream/"

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
		PageDeleteURL:               pageDeleteURL,
		PageMoveURL:                 pageMoveURL,
		RevisionCreateURL:           revisionCreateURL,
		RevisionScoreURL:            revisionScoreURL,
		RevisionVisibilityChangeURL: revisionVisibilityChangeURL,
	}
}

// Client request client
type Client struct {
	PageDeleteURL               string
	PageMoveURL                 string
	RevisionCreateURL           string
	RevisionScoreURL            string
	RevisionVisibilityChangeURL string
}

// PageDelete connect to page delete stream
func (cl *Client) PageDelete(ctx context.Context, handler func(evt *events.PageDelete)) error {
	return pageDelete(ctx, cl.PageDeleteURL, handler)
}

// PageMove connect to page move stream
func (cl *Client) PageMove(ctx context.Context, handler func(evt *events.PageMove)) error {
	return pageMove(ctx, cl.PageMoveURL, handler)
}

// RevisionCreate connect to revision create stream
func (cl *Client) RevisionCreate(ctx context.Context, handler func(evt *events.RevisionCreate)) error {
	return revisionCreate(ctx, cl.RevisionCreateURL, handler)
}

// RevisionScore connect to revision score stream
func (cl *Client) RevisionScore(ctx context.Context, handler func(evt *events.RevisionScore)) error {
	return revisionScore(ctx, cl.RevisionScoreURL, handler)
}

// RevisionVisibilityChange connext to revision visibility change stream
func (cl *Client) RevisionVisibilityChange(ctx context.Context, handler func(evt *events.RevisionVisibilityChange)) error {
	return revisionVisibilityChange(ctx, cl.RevisionVisibilityChangeURL, handler)
}
