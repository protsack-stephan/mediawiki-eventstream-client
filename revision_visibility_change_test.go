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

var revVisibilityChangeTestErrors = []error{io.EOF, io.EOF, context.Canceled}
var revVisibilityChangeTestSince = time.Now().UTC()
var revVisibilityChangeTestResponse = map[int]struct {
	Topic     string
	PageTitle string
	RevID     int
}{
	123: {
		Topic:     "eqiad.mediawiki.revision-visibility-change",
		PageTitle: "TestPage10",
		RevID:     123,
	},
	66132507: {
		Topic:     "eqiad.mediawiki.revision-visibility-change",
		PageTitle: "Ytilisateur:Denvis1/NCAA-Squelette_éqoio",
		RevID:     177205614,
	},
}

const revVisibilityChangeTestExecURL = "/revision-visibility-change-exec"
const revVisibilityChangeTestSubURL = "/revision-visibility-change-sub"

func createRevVisibilityChangeServer(t *testing.T, since *time.Time) (http.Handler, error) {
	router := http.NewServeMux()
	stubs, err := readStub("revision-visibility-change.json")

	if err != nil {
		return router, err
	}

	router.HandleFunc(revVisibilityChangeTestExecURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			w.Write(stub)
			f.Flush()
		}
	})

	router.HandleFunc(revVisibilityChangeTestSubURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			w.Write(stub)
			f.Flush()
		}
	})

	return router, nil
}

func testRevVisibilityChangeEvent(t *testing.T, evt *RevisionVisibilityChange) {
	expected := revVisibilityChangeTestResponse[evt.Data.PageID]
	assert.NotNil(t, expected)
	assert.Equal(t, expected.Topic, evt.ID[0].Topic)
	assert.Equal(t, expected.PageTitle, evt.Data.PageTitle)
	assert.Equal(t, expected.RevID, evt.Data.RevID)
}

func TestRevVisibilityChangeExec(t *testing.T) {
	router, err := createRevVisibilityChangeServer(t, &revVisibilityChangeTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionVisibilityChangeURL: revVisibilityChangeTestExecURL,
		}).
		Build()

	stream := client.RevisionVisibilityChange(context.Background(), revVisibilityChangeTestSince, func(evt *RevisionVisibilityChange) {
		testRevVisibilityChangeEvent(t, evt)
	})

	assert.Equal(t, io.EOF, stream.Exec())
}

func TestRevVisibilityChangeSub(t *testing.T) {
	since := revVisibilityChangeTestSince
	router, err := createRevVisibilityChangeServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionVisibilityChangeURL: revVisibilityChangeTestSubURL,
		}).
		Build()

	msgs := 0
	stream := client.RevisionVisibilityChange(ctx, revVisibilityChangeTestSince, func(evt *RevisionVisibilityChange) {
		testRevVisibilityChangeEvent(t, evt)
		since = evt.Data.Meta.Dt
		msgs++

		if msgs > 3 {
			cancel()
		}
	})

	errs := 0
	for err := range stream.Sub() {
		assert.Contains(t, err.Error(), revVisibilityChangeTestErrors[errs].Error())
		errs++
	}

	assert.Equal(t, 4, msgs)
}
