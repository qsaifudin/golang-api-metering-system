package main

import (
	"api_metering_system/handlers"
	"api_metering_system/metering"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the API metering system...")

	// Create new instances for API and storage metrics
	apiMetrics := metering.NewApiMetrics()
	storageMetrics := metering.NewStorageMetrics()

	// Create handlers for API and storage
	apiHandler := handlers.NewApiHandler(apiMetrics)
	storageHandler := handlers.NewStorageHandler(storageMetrics)

	log.Println("Registering HTTP handlers...")

	// Set up HTTP routes and handlers
	http.HandleFunc("/api/endpoint1", apiHandler.TrackEndpoint1)
	http.HandleFunc("/api/metrics", apiHandler.GetMetrics)
	http.HandleFunc("/upload", storageHandler.UploadFile)
	http.HandleFunc("/storage", storageHandler.GetStorage)

	log.Println("HTTP server is listening on http://localhost:8080")

	// Start the HTTP server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
