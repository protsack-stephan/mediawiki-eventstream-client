package eventstream

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var pageMoveTestErrors = []error{io.EOF, io.EOF, context.Canceled}
var pageMoveTestSince = time.Now().UTC()
var pageMoveTestResponse = map[int]struct {
	Topic     string
	PageTitle string
	RevID     int
}{
	504089: {
		Topic:     "eqiad.mediawiki.page-move",
		PageTitle: "Coulon_(gemeente)",
		RevID:     57655779,
	},
	8394444: {
		Topic:     "eqiad.mediawiki.page-move",
		PageTitle: "Olson_(Rapper)",
		RevID:     206122560,
	},
}

const pageMoveTestExecURL = "/page-move-exec"
const pageMoveTestSubURL = "/page-move-sub"

func createPageMoveServer(t *testing.T, since *time.Time) (http.Handler, error) {
	router := http.NewServeMux()
	stubs, err := readStub("page-move.json")

	if err != nil {
		return router, err
	}

	router.HandleFunc(pageMoveTestExecURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			w.Write(stub)
			f.Flush()
		}
	})

	router.HandleFunc(pageMoveTestSubURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			w.Write(stub)
			f.Flush()
		}
	})

	return router, nil
}

func testPageMoveEvent(t *testing.T, evt *PageMove) {
	expected := pageMoveTestResponse[evt.Data.PageID]
	assert.NotNil(t, expected)
	assert.Equal(t, expected.Topic, evt.ID[0].Topic)
	assert.Equal(t, expected.PageTitle, evt.Data.PageTitle)
	assert.Equal(t, expected.RevID, evt.Data.RevID)
}

func TestPageMoveExec(t *testing.T) {
	router, err := createPageMoveServer(t, &pageMoveTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageMoveURL: pageMoveTestExecURL,
		}).
		Build()

	stream := client.PageMove(context.Background(), time.Now().UTC(), func(evt *PageMove) {
		testPageMoveEvent(t, evt)
	})

	assert.Equal(t, io.EOF, stream.Exec())
}

func TestPageMoveSub(t *testing.T) {
	since := time.Now().UTC()
	router, err := createPageMoveServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageMoveURL: pageMoveTestSubURL,
		}).
		Build()

	msgs := 0
	stream := client.PageMove(ctx, pageMoveTestSince, func(evt *PageMove) {
		testPageMoveEvent(t, evt)
		since = evt.Data.Meta.Dt
		msgs++

		if msgs > 3 {
			cancel()
		}
	})

	errs := 0
	for err := range stream.Sub() {
		assert.Contains(t, err.Error(), pageMoveTestErrors[errs].Error())
		errs++
	}

	assert.Equal(t, 4, msgs)
}