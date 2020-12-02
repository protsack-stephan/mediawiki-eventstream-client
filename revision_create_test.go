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

var revisionCreateTestErrors = []error{io.EOF, io.EOF, context.Canceled}
var revisionCreateTestSince = time.Now().UTC()
var revisionCreateTestResponse = map[int]struct {
	Topic     string
	PageTitle string
	RevID     int
}{
	21512239: {
		Topic:     "eqiad.mediawiki.revision-create",
		PageTitle: "Category:Cyprian_Dylczyński",
		RevID:     516364180,
	},
	99305888: {
		Topic:     "eqiad.mediawiki.revision-create",
		PageTitle: "Q103437718",
		RevID:     1316829186,
	},
}

const revisionCreateTestExecURL = "/revision-create-exec"
const revisionCreateTestSubURL = "/revision-create-sub"

func createRevisionCreateServer(t *testing.T, since *time.Time) (http.Handler, error) {
	router := http.NewServeMux()
	stubs, err := readStub("revision-create.json")

	if err != nil {
		return router, err
	}

	router.HandleFunc(revisionCreateTestExecURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			w.Write(stub)
			f.Flush()
		}
	})

	router.HandleFunc(revisionCreateTestSubURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			w.Write(stub)
			f.Flush()
		}
	})

	return router, nil
}

func testRevisionCreateEvent(t *testing.T, evt *RevisionCreate) {
	expected := revisionCreateTestResponse[evt.Data.PageID]
	assert.NotNil(t, expected)
	assert.Equal(t, expected.Topic, evt.ID[0].Topic)
	assert.Equal(t, expected.PageTitle, evt.Data.PageTitle)
	assert.Equal(t, expected.RevID, evt.Data.RevID)
}

func TestRevisionCreateExec(t *testing.T) {
	router, err := createRevisionCreateServer(t, &revisionCreateTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionCreateURL: revisionCreateTestExecURL,
		}).
		Build()

	stream := client.RevisionCreate(context.Background(), time.Now().UTC(), func(evt *RevisionCreate) {
		testRevisionCreateEvent(t, evt)
	})

	assert.Equal(t, io.EOF, stream.Exec())
}

func TestRevisionCreateSub(t *testing.T) {
	since := time.Now().UTC()
	router, err := createRevisionCreateServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionCreateURL: revisionCreateTestSubURL,
		}).
		Build()

	msgs := 0
	stream := client.RevisionCreate(ctx, revisionCreateTestSince, func(evt *RevisionCreate) {
		testRevisionCreateEvent(t, evt)
		since = evt.Data.Meta.Dt
		msgs++

		if msgs > 3 {
			cancel()
		}
	})

	errs := 0
	for err := range stream.Sub() {
		assert.Contains(t, err.Error(), revisionCreateTestErrors[errs].Error())
		errs++
	}

	assert.Equal(t, 4, msgs)
}
