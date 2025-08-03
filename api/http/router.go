package http

import (
    "github.com/gorilla/mux"
    "net/http"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()

    // Define routes for weather-related endpoints
    router.HandleFunc("/weather/{city}", GetWeatherHandler).Methods("GET")
    router.HandleFunc("/weather/{city}/refresh", RefreshWeatherHandler).Methods("POST")

    return router
}