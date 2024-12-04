package main

import (
	"log"
	"net/http"

	"../take-home/handler"
)

func main() {
	http.HandleFunc("/", handler.CombinedHandler)
	port := "8080"
	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
