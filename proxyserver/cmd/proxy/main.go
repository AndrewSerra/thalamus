package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/AndrewSerra/thalamus/internal/lookup"
)

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

func handler(w http.ResponseWriter, r *http.Request) {
	servName, pathStartIdx, err := extractServiceName(r.URL.Path)
	if err != nil {
		http.Error(w, "No target for this path", http.StatusBadRequest)
		return
	}
	target, err := getForwardAddress(servName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if target == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	targetURL := target + r.URL.Path[pathStartIdx:]

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

func main() {
	fmt.Println("Proxy server listening on 127.0.0.1:8080")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}
