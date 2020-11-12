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
	stream := client.RevisionVisibilityChange(context.Background(), time.Now(), func(evt *events.RevisionVisibilityChange) {
		fmt.Println(evt.ID)
	})

	errs := stream.Sub()

	for err := range errs {
		fmt.Println(err)
	}
}
