package app

import (
	"github.com/yourusername/weather-report-app/internal/weather/application"
	"github.com/yourusername/weather-report-app/internal/user/application"
)

type App struct {
	WeatherService application.WeatherService
	UserService    application.UserService
}

func NewApp(weatherService application.WeatherService, userService application.UserService) *App {
	return &App{
		WeatherService: weatherService,
		UserService:    userService,
	}
}

func (a *App) Start() {
	// Initialize application logic and start services
}