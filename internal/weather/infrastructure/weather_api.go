package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/victorblv1/weather-report-app/internal/weather/domain"
)

type WeatherAPI struct {
	APIKey  string
	BaseURL string
}

func NewWeatherAPI(apiKey string) *WeatherAPI {
	return &WeatherAPI{
		APIKey:  apiKey,
		BaseURL: "https://api.openweathermap.org/data/2.5/weather",
	}
}

func (w *WeatherAPI) FetchWeather(city string) (domain.Weather, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", w.BaseURL, city, w.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return domain.Weather{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Weather{}, fmt.Errorf("failed to fetch weather: %s", resp.Status)
	}

	var result struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return domain.Weather{}, err
	}

	return domain.Weather{
		City:        city,
		Temperature: result.Main.Temp,
		Condition:   result.Weather[0].Description,
	}, nil
}

func (w *WeatherAPI) SaveWeather(data domain.Weather) error {
	// This method can be implemented to save weather data if needed
	return nil
}
