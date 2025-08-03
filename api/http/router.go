package http

import (
	"github.com/gorilla/mux"
)

func NewRouter(handler *Handler) *mux.Router {
	router := mux.NewRouter()

	// Define routes for weather-related endpoints
	router.HandleFunc("/weather/{city}", handler.GetWeatherHandler).Methods("GET")
	router.HandleFunc("/weather/{city}/refresh", handler.RefreshWeatherHandler).Methods("POST")

	return router
}
