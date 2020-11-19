package eventstream

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var pageDeleteTestSince = time.Now().UTC()

const pageDeleteTestExecURL = "/page-delete-exec"
const pageDeleteTestSubURL = "/page-delete-sub"

func createPageDeleteServer(ctx context.Context, t *testing.T, since *time.Time) (http.Handler, error) {
	router := http.NewServeMux()
	stubs, err := readStub("page-delete.json")

	if err != nil {
		return router, err
	}

	router.HandleFunc(pageDeleteTestExecURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))
		f := w.(http.Flusher)

		for _, stub := range stubs {
			w.Write(stub)
			f.Flush()
		}
	})

	router.HandleFunc(pageDeleteTestSubURL, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, since.Format(time.RFC3339), r.URL.Query().Get("since"))
		fmt.Println(r.URL.Query().Get("since"), "srv")
		f := w.(http.Flusher)

		for {
			if ctx.Err() != nil {
				break
			}

			for _, stub := range stubs {
				time.Sleep(1 * time.Second)
				w.Write(stub)
				f.Flush()
			}
		}
	})

	return router, nil
}

func TestPageDeleteExec(t *testing.T) {
	ctx := context.Background()
	router, err := createPageDeleteServer(ctx, t, &pageDeleteTestSince)
	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageDeleteURL: pageDeleteTestExecURL,
		}).
		Build()

	type msg struct {
		Topic     string
		PageTitle string
		RevID     int
	}

	expected := map[int]msg{
		4656021: {
			Topic:     "eqiad.mediawiki.page-delete",
			PageTitle: "réduisit",
			RevID:     22058660,
		},
		4656283: {
			Topic:     "eqiad.mediawiki.page-delete",
			PageTitle: "récupéreraient",
			RevID:     22110162,
		},
	}

	stream := client.PageDelete(context.Background(), time.Now().UTC(), func(evt *PageDelete) {
		expectedItem := expected[evt.Data.PageID]

		assert.NotNil(t, expectedItem)
		assert.Equal(t, expectedItem.Topic, evt.ID[0].Topic)
		assert.Equal(t, expectedItem.PageTitle, evt.Data.PageTitle)
		assert.Equal(t, expectedItem.RevID, evt.Data.RevID)
	})

	assert.Equal(t, io.EOF, stream.Exec())
}

func TestPageDeleteSub(t *testing.T) {
	since := time.Now().UTC()
	srvCtx, srvCancel := context.WithCancel(context.Background())
	router, err := createPageDeleteServer(srvCtx, t, &since)
	assert.Nil(t, err)

	srv := httptest.NewServer(router)
	defer srv.Close()

	clientCtx, clientCancel := context.WithCancel(context.Background())
	client := NewBuilder().
		URL(srv.URL).
		Options(&Options{
			PageDeleteURL: pageDeleteTestSubURL,
		}).
		Build()

	stream := client.PageDelete(clientCtx, pageDeleteTestSince, func(evt *PageDelete) {
		since = evt.Data.Meta.Dt
		fmt.Println(evt.Data.Meta.Dt)
	})

	go func() {
		time.Sleep(5 * time.Second)
		srvCancel()
		srv.CloseClientConnections()
		time.Sleep(5 * time.Second)
		clientCancel()
	}()

	for err := range stream.Sub() {
		fmt.Println(err, "test")
	}
}

// 1, 2
// -> /page-delete-sub -> for -> msg -> 1,2 -> 1,2 -> 1,2

// -> /page-delete-sub -> for -> msg -> 1,2 -> 1,2 -> 1,2
