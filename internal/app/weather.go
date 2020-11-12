package app

import (
	"encoding/json"
	"fmt"
	"hammurabi-game/config"
	"hammurabi-game/internal/models"
	"log"
	"net/http"
)

// GetCurrentWeather ...
func GetCurrentWeather(cfg *config.App) (*models.WeatherAPIResponse, error) {
	urlPath := cfg.WeatherAPI.URL + "?" + "id=" + cfg.WeatherAPI.CityID + "&appid=" + cfg.WeatherAPI.Token + "&lang=ru&units=metric"
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %s",
			http.StatusText(resp.StatusCode))
	}

	var data *models.WeatherAPIResponse

	err = json.NewDecoder(resp.Body).Decode(data)
	if err != nil {
		log.Fatalf("Error decoding response: %s", err)
		return nil, err
	}

	return data, nil
}
