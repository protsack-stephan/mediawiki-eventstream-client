package eventstream

import (
	"context"
	"time"
)

func keepAlive(cb func() error, errors chan error) {
	for {
		err := cb()

		errors <- err
		if err == context.Canceled {
			close(errors)
			break
		} else {
			time.Sleep(1 * time.Second)
		}
	}
}
