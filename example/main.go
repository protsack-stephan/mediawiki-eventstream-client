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
	stream := client.PageDelete(ctx, time.Now(), func(evt *events.PageDelete) {
		fmt.Println(evt.Data)
	})

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for err := range stream.Sub() {
		fmt.Println(err)
	}
}
