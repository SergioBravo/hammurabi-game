package app

import (
	"fmt"
	"hammurabi-game/config"
	"hammurabi-game/utils"
	"log"
)

// WeatherResponse ...
func WeatherResponse(cfg *config.App) string {
	r, err := GetCurrentWeather(cfg)
	if err != nil {
		log.Fatalf("error: %s", err)
		return ""
	}

	return fmt.Sprintf(`Город %v. \\n
			В данный момент %v ⛅️
			Температура воздуха 🌡%v градусов Цельсия.\n
			Ощущается как 🌡%v градусов Цельсия.\n
			Влажность %v процента. Атмосферное давление %v мм ртутного столба.\n
			🌬Ветер %v. Скорость ветра %v метров в секунду.\n`,
		r.Name,
		r.Weather[0].Description,
		r.Main.Temp,
		r.Main.FeelsLike,
		r.Main.Humidity,
		r.Main.Pressure,
		utils.GetWindDirection(r.Wind.Deg),
		r.Wind.Speed)
}
