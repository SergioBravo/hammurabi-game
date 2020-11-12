package app

import (
	"fmt"
	"hammurabi-game/config"
	"hammurabi-game/utils"
	"log"
	"time"
)

// WeatherResponse ...
func WeatherResponse(cfg *config.App) string {
	r, err := GetCurrentWeather(cfg)
	if err != nil {
		log.Fatalf("error: %s", err)
		return ""
	}

	return fmt.Sprintf(`
	Город %v.
	В данный момент %v ⛅️
	Температура воздуха 🌡%v градусов Цельсия.
	Ощущается как 🌡%v градусов Цельсия.
	Влажность %v процентов. Атмосферное давление %v мм ртутного столба.
	🌬Ветер %v. Скорость ветра %v м/с.
	Восход %v
	Закат %v`,
		r.Name,
		r.Weather[0].Description,
		r.Main.Temp,
		r.Main.FeelsLike,
		r.Main.Humidity,
		r.Main.Pressure,
		utils.GetWindDirection(r.Wind.Deg),
		r.Wind.Speed,
		time.Unix(int64(r.Sys.Sunrise+r.Timezone), 0).Format("hh:mm:ss"),
		time.Unix(int64(r.Sys.Sunset+r.Timezone), 0).Format("hh:mm:ss"))
}
