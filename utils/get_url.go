package utils

import (
	"strconv"
	"time"
)

// FormatURL update url to contain since param
func FormatURL(url string, since time.Time) string {
	return url + "?since=" + strconv.Itoa(GetTimeMilliseconds(since))
}
