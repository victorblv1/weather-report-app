package weather

import (
	"github.com/victorblv1/weather-report-app/internal/weather/domain"
)

type Weather struct {
	City        string
	Temperature float64
	Condition   string
}

type WeatherService struct {
	repo domain.WeatherRepository
}

func NewWeatherService(repo domain.WeatherRepository) *WeatherService {
	return &WeatherService{repo: repo}
}

func (s *WeatherService) GetWeather(city string) (domain.Weather, error) {
	weather, err := s.repo.FetchWeather(city)
	if err != nil {
		return domain.Weather{}, err
	}
	return weather, nil
}

func (s *WeatherService) RefreshWeather(city string) error {
	weather, err := s.repo.FetchWeather(city)
	if err != nil {
		return err
	}
	return s.repo.SaveWeather(weather)
}
