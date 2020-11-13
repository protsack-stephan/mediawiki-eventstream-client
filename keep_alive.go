package eventstream

import (
	"context"
	"time"
)

func keepAlive(handler func(since time.Time) error, store *storage) {
	for {
		err := handler(store.getSince())
		store.setError(err)

		if err == context.Canceled {
			store.closeErrors()
			return
		}

		time.Sleep(store.getBackoff())
	}
}
