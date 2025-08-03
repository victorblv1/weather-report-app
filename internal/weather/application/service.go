package weather

import (
	"errors"
)

type Weather struct {
	City        string
	Temperature float64
	Condition   string
}

type WeatherService struct {
	repo WeatherRepository
}

func NewWeatherService(repo WeatherRepository) *WeatherService {
	return &WeatherService{repo: repo}
}

func (s *WeatherService) GetWeather(city string) (*Weather, error) {
	weather, err := s.repo.FetchWeather(city)
	if err != nil {
		return nil, err
	}
	return weather, nil
}

func (s *WeatherService) RefreshWeather(city string) error {
	weather, err := s.repo.FetchWeather(city)
	if err != nil {
		return err
	}
	if weather == nil {
		return errors.New("no weather data found")
	}
	return s.repo.SaveWeather(*weather)
}