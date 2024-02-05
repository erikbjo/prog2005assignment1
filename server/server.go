package server

import (
	"log"
	"net/http"
	"os"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Set up handler endpoints
	http.HandleFunc(LOCATION_PATH, LocationHandler)
	http.HandleFunc(COLLECTION_PATH, CollectionHandler)

	// Start server
	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
