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
			icon += ":partly_sunny:"
		case "Clear":
			icon += ":sunny:"
		case "Mist":
		case "Fog":
		case "Smoke":
		case "Haze":
		case "Dust":
		case "Sand":
			icon += ":fog:"
		case "Squall":
			icon += ":dash:"
		case "Thunderstorm":
			icon += ":thunder_cloud_and_rain:"
		case "Drizzle":
		case "Rain":
			icon += ":cloud_with_rain:"
		case "Snow":
			icon += ":cloud_with_snow:"
		}
	}

	return w, icon
}
