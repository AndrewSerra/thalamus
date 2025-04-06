package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/AndrewSerra/thalamus/proxyserver/internal/lookup"
)

var lock = sync.Mutex{}

type RequestBody struct {
	ServiceName string `json:"service_name"`
	AvailableAt string `json:"available_at"`
}

func handleRegistration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		registry := lookup.NewLookupWorker()
		log.Printf("Received registration request from client %s", r.URL.String())

		lock.Lock()
		defer lock.Unlock()

		var body RequestBody
		err := json.NewDecoder(r.Body).Decode(&body)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		registry.SetAddress(body.ServiceName, body.AvailableAt)
		w.WriteHeader(http.StatusOK)
	default:
		log.Printf("Unknown method request from client %s: %s", r.URL.String(), r.Method)
		w.WriteHeader(http.StatusNotFound)
	}
}

func handleUnregistration(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		registry := lookup.NewLookupWorker()
		log.Printf("Received registration request from client %s", r.URL.String())

		lock.Lock()
		defer lock.Unlock()

		var body RequestBody
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		registry.DeleteAddress(body.ServiceName, body.AvailableAt)
		w.WriteHeader(http.StatusOK)
	default:
		log.Printf("Unknown method request from client %s: %s", r.URL.String(), r.Method)
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func main() {
	http.HandleFunc("/register", handleRegistration)
	http.HandleFunc("/unregister", handleUnregistration)

	log.Println("Registration server listening on 127.0.0.1:8081")
	http.ListenAndServe(":8081", nil)
}
