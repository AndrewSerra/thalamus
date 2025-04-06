package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/AndrewSerra/thalamus/proxyserver/internal/analytics"
	"github.com/AndrewSerra/thalamus/proxyserver/internal/lookup"
)

func extractRequestInfo(r *http.Request, serviceName string, pathStartIdx uint64) analytics.RequestInfo {
	return analytics.RequestInfo{
		ServiceName: serviceName,
		Path:        r.URL.Path[pathStartIdx:],
		Method:      r.Method,
		Sender:      r.RemoteAddr,
		Timestamp:   time.Now().Format(time.RFC3339),
	}
}

func extractServiceName(path string) (string, uint64, error) {
	if path == "/" {
		return "", 0, errors.New("invalid path, no service defined")
	}

	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			return path[1:i], uint64(i), nil
		}
	}
	return path[1:], uint64(len(path)), nil
}

func getForwardAddress(servName string) (string, error) {
	forwardLookup := lookup.NewLookupWorker()
	addr := forwardLookup.GetAddresses(servName)
	if len(addr) == 0 {
		return "", errors.New("no available workers")
	}
	return addr[0], nil
}

func getHandler(eventChan chan analytics.RequestInfo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		servName, pathStartIdx, err := extractServiceName(r.URL.Path)
		if err != nil {
			http.Error(w, "No target for this path", http.StatusBadRequest)
			return
		}
		target, err := getForwardAddress(servName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if target == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		eventChan <- extractRequestInfo(r, servName, pathStartIdx)

		targetURL := target + r.URL.Path[pathStartIdx:]
		log.Println("Forwarding request to: ", targetURL)
		proxyReq, err := http.NewRequest(r.Method, targetURL, r.Body)
		if err != nil {
			http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
			return
		}

		for name, values := range r.Header {
			for _, value := range values {
				proxyReq.Header.Add(name, value)
			}
		}

		resp, err := http.DefaultClient.Do(proxyReq)
		if err != nil {
			http.Error(w, "Error sending proxy request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		for name, values := range resp.Header {
			for _, value := range values {
				w.Header().Add(name, value)
			}
		}

		w.WriteHeader(resp.StatusCode)

		io.Copy(w, resp.Body)
	}
}

func main() {

	events := make(chan analytics.RequestInfo)

	// Analytics event sender
	// go func() {
	// 	aq := analytics.NewAnalyticsQueue()

	// 	for {
	// 		event := <-events
	// 		log.Printf("Sending event: %+v", event)

	// 		aq.PushRequestEventQueue(event)
	// 	}
	// }()

	log.Println("Proxy server listening on 127.0.0.1:8080")
	http.HandleFunc("/", getHandler(events))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
