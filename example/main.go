package main

import (
	"fmt"
	"log"

	eventstream "github.com/protsack-stephan/mediawiki-eventstream-client"
	"github.com/protsack-stephan/mediawiki-eventstream-client/events"
)

func main() {
	err := eventstream.RevisionScore(func(evt *events.RevisionScore) {
		fmt.Println(evt)
	})

	if err == nil {
		log.Panic(err)
	}
}
