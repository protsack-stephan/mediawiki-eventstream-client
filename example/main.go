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
	stream := client.PageDelete(context.Background(), time.Now(), func(evt *events.PageDelete) {
		fmt.Println(evt.Data)
	})

	errs := stream.Sub()

	for err := range errs {
		fmt.Println(err)
	}
}
