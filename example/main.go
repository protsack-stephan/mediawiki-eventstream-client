package main

import (
	"context"
	"fmt"
	"time"

	eventstream "github.com/protsack-stephan/mediawiki-eventstream-client"
)

func main() {
	client := eventstream.NewClient()

	stream := client.RevisionScore(context.Background(), time.Now().UTC(), func(evt *eventstream.RevisionScore) {
		fmt.Println(evt.Data.Schema)
		fmt.Println(evt.Data.Meta.Dt)
		fmt.Println(evt.Data.Scores.Damaging)
	})

	for err := range stream.Sub() {
		fmt.Println(err)
	}
}
