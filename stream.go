package eventstream

import (
	"context"
	"time"
)

// NewStream create new result instance
func NewStream(store *storage, handler func(since time.Time) error) *Stream {
	return &Stream{
		store,
		handler,
	}
}

// Stream stream execution result
type Stream struct {
	store   *storage
	handler func(since time.Time) error
}

// Exec blocking execution stream
func (sm *Stream) Exec() error {
	return sm.handler(sm.store.getSince())
}

// Sub non blocking execution stream
func (sm *Stream) Sub() chan error {
	go sm.keepAlive()
	return sm.store.getErrors()
}

func (sm *Stream) keepAlive() {
	for {
		err := sm.handler(sm.store.getSince())
		sm.store.setError(err)

		if err == context.Canceled {
			sm.store.closeErrors()
			return
		}

		time.Sleep(sm.store.getBackoff())
	}
}
