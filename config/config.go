package config

import "os"

// App struct for store application configs
type App struct {
	Bot        BotConfig
	WeatherAPI WeatherAPI
}

// BotConfig store telegram bot configs
type BotConfig struct {
	Token string
	URL   string
	Port  string
}

// WeatherAPI ...
type WeatherAPI struct {
	Token  string
	URL    string
	CityID string
}

// New returns a new App struct
func New() *App {
	return &App{
		Bot: BotConfig{
			Token: getEnv("TOKEN", ""),
			URL:   getEnv("URL", ""),
			Port:  getEnv("PORT", ""),
		},

		WeatherAPI: WeatherAPI{
			Token:  getEnv("OPEN_WEATHER_TOKEN", ""),
			URL:    getEnv("OPEN_WEATHER_URL", ""),
			CityID: getEnv("CITY_ID", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
