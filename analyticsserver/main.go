package main

import (
	"log"

	"github.com/AndrewSerra/thalamus/analyticsserver/internal/analytics"
)

func main() {

	analyticsq := analytics.NewAnalyticsQueue()

	for {
		event, err := analyticsq.PopRequestEventQueue()

		if err != nil {
			log.Printf("Error popping request event to queue: %s", err)
			continue
		}

		log.Printf("Received event: %+v", event)
	}

}
