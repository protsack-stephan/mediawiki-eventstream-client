package eventstream

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const keepAliveTestBackoffTime = time.Millisecond * 1
const keepAliveNumberOfErrors = 5

func TestKeepAlive(t *testing.T) {
	storageTime := time.Now().UTC()
	thrownErrs := 0
	caughtErrs := 0
	storage := newStorage(storageTime, keepAliveTestBackoffTime)

	assert.NotNil(t, storage)

	handler := func(since time.Time) error {
		assert.Equal(t, storageTime, since)
		thrownErrs++

		if thrownErrs < keepAliveNumberOfErrors {
			return fmt.Errorf("test error")
		}

		return context.Canceled
	}

	go keepAlive(handler, storage)

	for err := range storage.getErrors() {
		caughtErrs++

		if thrownErrs >= keepAliveNumberOfErrors {
			assert.Equal(t, context.Canceled, err)
		} else {
			storageTime = storageTime.Add(time.Hour * 1)
			storage.setSince(storageTime)
		}
	}

	assert.Equal(t, keepAliveNumberOfErrors, thrownErrs)
	assert.Equal(t, keepAliveNumberOfErrors, caughtErrs)
}
