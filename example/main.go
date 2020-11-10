package main

import (
	"context"
	"fmt"
	"log"

	eventstream "github.com/protsack-stephan/mediawiki-eventstream-client"
	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

func main() {
	ctx := context.Background()
	client := eventstream.NewClient()

	err := client.RevisionScore(ctx, func(evt *events.RevisionScore) {
		fmt.Println(evt)
	})

	if err != nil {
		log.Panic(err)
	}
}
