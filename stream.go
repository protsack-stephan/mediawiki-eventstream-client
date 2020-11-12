package eventstream

import (
	"time"

	"github.com/protsack-stephan/mediawiki-eventstream-client/utils"
)

// NewStream create new result instance
func NewStream(reconnectTime time.Duration, errors chan error, subscriber func() error) *Stream {
	if reconnectTime == 0 {
		reconnectTime = time.Second * 1
	}

	return &Stream{
		reconnectTime,
		errors,
		subscriber,
	}
}

// Stream stream execution result
type Stream struct {
	reconnectTime time.Duration
	errors        chan error
	subscriber    func() error
}

// Exec blocking execution stream
func (str *Stream) Exec() error {
	return str.subscriber()
}

// Sub non blocking execution stream
func (str *Stream) Sub() chan error {
	go utils.KeepAlive(str.subscriber, str.reconnectTime, str.errors)
	return str.errors
}
