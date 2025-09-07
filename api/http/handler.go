package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	weather "github.com/victorblv1/weather-report-app/internal/weather/application"
)

type Handler struct {
	weatherService *weather.WeatherService
}

func NewHandler(weatherService *weather.WeatherService) *Handler {
	return &Handler{weatherService: weatherService}
}

func (h *Handler) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]
	weather, err := h.weatherService.GetWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Respond with weather data (this should be formatted as JSON)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weather)
}

func (h *Handler) RefreshWeatherHandler(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["city"]
	err := h.weatherService.RefreshWeather(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "refreshed",
		"city":   city,
	})
}
