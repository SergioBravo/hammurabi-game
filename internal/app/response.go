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

	weather, icon := utils.GetWeatherIcon(r.Weather)

	return fmt.Sprintf(`
	Город %v.
	В данный момент %v %v.
	Температура воздуха 🌡%v градусов Цельсия.
	Ощущается как 🌡%v градусов Цельсия.
	Влажность %v процентов.
	Атмосферное давление %v гПа.
	🌬Ветер %v. Скорость ветра %v м/с.
	Восход %v
	Закат %v`,
		r.Name,
		weather,
		icon,
		r.Main.Temp,
		r.Main.FeelsLike,
		r.Main.Humidity,
		r.Main.Pressure,
		utils.GetWindDirection(r.Wind.Deg),
		r.Wind.Speed,
		utils.FormatTime(time.Unix(int64(r.Sys.Sunrise+r.Timezone), 0)),
		utils.FormatTime(time.Unix(int64(r.Sys.Sunset+r.Timezone), 0)))
}
