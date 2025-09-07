package app

import (
	user "github.com/victorblv1/weather-report-app/internal/user/application"
	weather "github.com/victorblv1/weather-report-app/internal/weather/application"
)

type App struct {
	WeatherService *weather.WeatherService
	UserService    user.UserService
}

func NewApp(weatherService *weather.WeatherService, userService user.UserService) *App {
	return &App{
		WeatherService: weatherService,
		UserService:    userService,
	}
}

func (a *App) Start() {
	// Initialize application logic and start services
}
