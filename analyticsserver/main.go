package main

import (
	"log"
	"net/http"
	"time"

	"github.com/AndrewSerra/thalamus/analyticsserver/internal/analytics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "incoming_requests_total",
		Help: "Count of incoming requests grouped by service, path, and method.",
	},
	[]string{"service", "path", "method", "sender"},
)

func init() {
	prometheus.MustRegister(requestCounter)
}

func recordEvent(e analytics.RequestInfo) {
	requestCounter.WithLabelValues(e.ServiceName, e.Path, e.Method, e.Sender).Inc()
}

func main() {

	go func() {
		analyticsq := analytics.NewAnalyticsQueue()

		for {
			event, err := analyticsq.PopRequestEventQueue()

			if err != nil {
				log.Printf("Error popping request event to queue: %s", err)
				time.Sleep(5 * time.Second)
				continue
			}

			log.Printf("Received event: %+v", event)
			recordEvent(event)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9091", nil)
}
