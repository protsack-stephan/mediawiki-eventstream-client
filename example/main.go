package main

import (
	"context"
	"fmt"
	"log"
	"time"

	eventstream "github.com/protsack-stephan/mediawiki-eventstream-client"
	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

func main() {
	client := eventstream.NewClient()
	stream := client.PageDelete(context.Background(), time.Now(), func(evt *events.PageDelete) {
		fmt.Println(evt.Data)
	})

	err := stream.Exec()

	if err != nil {
		log.Panic(err)
	}
}
