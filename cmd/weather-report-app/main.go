package main

import (
	"log"
	"net/http"
	"os"

	apihttp "github.com/victorblv1/weather-report-app/api/http"
	"github.com/victorblv1/weather-report-app/internal/app"
	user "github.com/victorblv1/weather-report-app/internal/user/application"
	weather "github.com/victorblv1/weather-report-app/internal/weather/application"
	"github.com/victorblv1/weather-report-app/internal/weather/infrastructure"
)

func main() {

	apiKey := os.Getenv("WEATHER_API_KEY")
	weatherRepo := infrastructure.NewWeatherAPI(apiKey)

	// Initialize the weather service
	weatherService := weather.NewWeatherService(weatherRepo)

	// Initialize the HTTP handler with the weather service
	handler := apihttp.NewHandler(weatherService)

	// Set up the HTTP router
	router := apihttp.NewRouter(handler)

	// Initialize the application
	userService := user.NewUserService()
	application := app.NewApp(weatherService, *userService)
	application.Start()

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
