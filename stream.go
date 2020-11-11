package eventstream

import (
	"time"

	"github.com/protsack-stephan/mediawiki-eventstream-client/utils"
)

// NewStream create new result instance
func NewStream(reconnectTime time.Duration, errors chan error, stream func() error) *Stream {
	if reconnectTime == 0 {
		reconnectTime = time.Second * 1
	}

	return &Stream{
		reconnectTime,
		errors,
		stream,
	}
}

// Stream stream execution result
type Stream struct {
	reconnectTime time.Duration
	errors        chan error
	stream        func() error
}

// Exec blocking execution stream
func (r *Stream) Exec() error {
	return r.stream()
}

// Sub non blocking execution stream
func (r *Stream) Sub() chan error {
	go utils.KeepAlive(r.stream, r.reconnectTime, r.errors)
	return r.errors
}
