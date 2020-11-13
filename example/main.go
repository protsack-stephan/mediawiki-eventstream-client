package main

import (
	"context"
	"fmt"
	"time"

	eventstream "github.com/protsack-stephan/mediawiki-eventstream-client"
)

func main() {
	client := eventstream.NewClient()

	stream := client.RevisionCreate(context.Background(), time.Now().UTC(), func(evt *eventstream.RevisionCreate) {
		fmt.Println(evt.Data.Schema)
		fmt.Println(evt.Data.Meta.Dt)
	})

	for err := range stream.Sub() {
		fmt.Println(err)
	}
}
