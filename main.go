package main

import (
	"log"
	"net/http"
	"pravelin/events"
	"pravelin/logging"
)

func main() {
	logging.InfoLog.Println("Application Loaded")
	store := events.NewEventStore()

	http.HandleFunc("/api/v1/events", store.EventRequestHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
