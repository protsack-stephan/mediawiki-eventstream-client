package utils

import (
	"context"
	"time"
)

// KeepAlive keep the connection alive forever
func KeepAlive(cb func() error, reconnectTime time.Duration, errors chan error) {
	for {
		err := cb()

		errors <- err
		if err == context.Canceled {
			close(errors)
			break
		} else {
			time.Sleep(reconnectTime)
		}
	}
}
