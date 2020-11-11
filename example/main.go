package main

import (
	"context"
	"fmt"
	"time"

	eventstream "github.com/protsack-stephan/mediawiki-eventstream-client"
	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	client := eventstream.NewClient()
	stream := client.RevisionCreate(ctx, time.Now(), func(evt *events.RevisionCreate) {
		fmt.Println(evt)
	})

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for err := range stream.Sub() {
		fmt.Println(err)
	}
}
