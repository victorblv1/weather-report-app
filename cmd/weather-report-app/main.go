package main

import (
    "log"
    "net/http"
    "weather-report-app/api/http"
    "weather-report-app/internal/app"
)

func main() {
    // Initialize the application
    application := app.NewApplication()

    // Set up the HTTP router
    router := http.NewRouter(application)

    // Start the HTTP server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Could not start server: %s\n", err)
    }
}