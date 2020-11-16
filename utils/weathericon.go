package utils

import "hammurabi-game/internal/models"

// GetWeatherIcon return emoji for corrent wwether condition
func GetWeatherIcon(weatherList []models.Weather) (string, string) {
	w, icon := "", ""
	for _, weather := range weatherList {
		if len(w) > 0 && len(icon) > 0 {
			w += " и "
			icon += " "
		}
		w += weather.Description
		switch weather.Main {
		case "Clouds":
			icon += "⛅️"
		case "Clear":
			icon += "☀️"
		case "Mist":
		case "Fog":
		case "Smoke":
		case "Haze":
		case "Dust":
		case "Sand":
			icon += "🌫"
		case "Squall":
			icon += "💨"
		case "Thunderstorm":
			icon += "⛈"
		case "Drizzle":
		case "Rain":
			icon += "🌧"
		case "Snow":
			icon += "🌨"
		}
	}

	return w, icon
}
