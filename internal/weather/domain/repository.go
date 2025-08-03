package domain

type WeatherRepository interface {
	FetchWeather(city string) (Weather, error)
	SaveWeather(data Weather) error
}
