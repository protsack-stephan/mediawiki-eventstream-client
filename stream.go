package eventstream

import (
	"context"
	"time"
)

// NewStream create new result instance
func NewStream(reconnectTime time.Duration, errors chan error, handler func() error) *Stream {
	if reconnectTime == 0 {
		reconnectTime = time.Second * 1
	}

	return &Stream{
		reconnectTime,
		errors,
		handler,
	}
}

// Stream stream execution result
type Stream struct {
	reconnectTime time.Duration
	errors        chan error
	handler       func() error
}

// Exec blocking execution stream
func (str *Stream) Exec() error {
	return str.handler()
}

// Sub non blocking execution stream
func (str *Stream) Sub() chan error {
	go str.keepAlive()
	return str.errors
}

func (str *Stream) keepAlive() {
	for {
		err := str.handler()

		str.errors <- err
		if err == context.Canceled {
			close(str.errors)
			break
		} else {
			time.Sleep(str.reconnectTime)
		}
	}
}
