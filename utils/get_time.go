package utils

import (
	"time"
)

// GetTimeMilliseconds convert time to milliseconds timstamp
func GetTimeMilliseconds(date time.Time) int {
	return int(time.Now().UnixNano() / int64(time.Millisecond))
}
