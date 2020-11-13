package main

import (
	"context"
	"fmt"
	"time"

	eventstream "github.com/protsack-stephan/mediawiki-eventstream-client"
	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

func main() {
	client := eventstream.NewClient()

	stream := client.RevisionVisibilityChange(context.Background(), time.Now().UTC(), func(evt *events.RevisionVisibilityChange) {
		fmt.Println(evt.Data.Schema)
		fmt.Println(evt.Data.Meta.Dt)
	})

	for err := range stream.Sub() {
		fmt.Println(err)
	}
}
