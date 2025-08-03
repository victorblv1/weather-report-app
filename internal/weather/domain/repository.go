package domain

type Weather struct {
    City        string  `json:"city"`
    Temperature float64 `json:"temperature"`
    Condition   string  `json:"condition"`
}

type WeatherRepository interface {
    FetchWeather(city string) (Weather, error)
    SaveWeather(data Weather) error
}