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

var revisionScoreTestErrors = []error{io.EOF, io.EOF, context.Canceled}
var revisionScoreTestSince = time.Now().UTC()
var revisionScoreTestResponse = map[int]struct {
	Topic     string
	PageTitle string
	RevID     int
}{
	66132507: {
		Topic:     "eqiad.mediawiki.revision-score",
		PageTitle: "Q66533108",
		RevID:     1316923273,
	},
	13731303: {
		Topic:     "eqiad.mediawiki.revision-score",
		PageTitle: "Utilisateur:Denvis1/NCAA-Squelette_équipe",
		RevID:     177205614,
	},
}

const revisionScoreTestExecURL = "/revision-score-exec"
const revisionScoreTestSubURL = "/revision-score-sub"

func createRevisionScoreServer(t *testing.T, since *time.Time) (http.Handler, error) {
	router := http.NewServeMux()
	stubs, err := readStub("revision-score.json")

	if err != nil {
		return router, err
	}

	router.HandleFunc(revisionScoreTestExecURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			w.Write(stub)
			f.Flush()
		}
	})

	router.HandleFunc(revisionScoreTestSubURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))

		f := w.(http.Flusher)

		for _, stub := range stubs {
			w.Write(stub)
			f.Flush()
		}
	})

	return router, nil
}

func testRevisionScoreEvent(t *testing.T, evt *RevisionScore) {
	expected := revisionScoreTestResponse[evt.Data.PageID]
	assert.NotNil(t, expected)
	assert.Equal(t, expected.Topic, evt.ID[0].Topic)
	assert.Equal(t, expected.PageTitle, evt.Data.PageTitle)
	assert.Equal(t, expected.RevID, evt.Data.RevID)
}

func TestRevisionScoreExec(t *testing.T) {
	router, err := createRevisionScoreServer(t, &revisionScoreTestSince)
	assert.NoError(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionScoreURL: revisionScoreTestExecURL,
		}).
		Build()

	stream := client.RevisionScore(context.Background(), time.Now().UTC(), func(evt *RevisionScore) {
		testRevisionScoreEvent(t, evt)
	})

	assert.Equal(t, io.EOF, stream.Exec())
}

func TestRevisionScoreSub(t *testing.T) {
	since := time.Now().UTC()
	router, err := createRevisionScoreServer(t, &since)

	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			RevisionScoreURL: revisionScoreTestSubURL,
		}).
		Build()

	msgs := 0
	stream := client.RevisionScore(ctx, revisionScoreTestSince, func(evt *RevisionScore) {
		testRevisionScoreEvent(t, evt)
		since = evt.Data.Meta.Dt
		msgs++

		if msgs > 3 {
			cancel()
		}
	})

	errs := 0
	for err := range stream.Sub() {
		assert.Contains(t, err.Error(), revisionScoreTestErrors[errs].Error())
		errs++
	}

	assert.Equal(t, 4, msgs)
}
