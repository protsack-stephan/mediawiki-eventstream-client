package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Info for the topic
type Info struct {
	Topic     string `json:"topic"`
	Partition int    `json:"partition"`
	Timestamp int    `json:"timestamp"`
	Offset    int    `json:"offset"`
}

// Event streams sse eevent
type Event struct {
	ID   []Info
	Data []byte
}

// SetID set id from string
func (evt *Event) SetID(body string) error {
	if strings.HasPrefix(body, "id:") {
		return json.Unmarshal([]byte(strings.TrimSpace(strings.TrimPrefix(body, "id:"))), &evt.ID)
	}

	return fmt.Errorf("wrong body format")
}

//SetData set data interface from string
func (evt *Event) SetData(body string) error {
	if strings.HasPrefix(body, "data:") {
		evt.Data = []byte(strings.TrimSpace(strings.TrimPrefix(body, "data:")))
	}

	return fmt.Errorf("wrong body fromat")
}
