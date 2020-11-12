package main

import (
	"encoding/json"
	"fmt"
	"hammurabi-game/config"
	"log"
	"net/http"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load("env/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cfg := config.New()

	bot, err := tgbotapi.NewBotAPI(cfg.Bot.Token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(cfg.Bot.URL + cfg.Bot.Token))
	if err != nil {
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + cfg.Bot.Token)
	go func() {
		if err := http.ListenAndServe(":"+cfg.Bot.Port, nil); err != nil {
			log.Fatalf("error: %s", err)
		}
	}()

	for update := range updates {
		reply := "Не знаю что сказать"
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "start":
			reply = "Привет. Я телеграм-бот"
		case "hello":
			reply = "world"
		case "weather":
			r, err := makeRequest(cfg)
			if err != nil {
				log.Fatalf("error: %s", err)
			}

			reply = fmt.Sprintf(`
			Город %v. \n
			В данный момент %v ⛅️
			Температура воздуха 🌡%v градусов Цельсия.\nОщущается как 🌡%v градусов Цельсия.\n
			Влажность %v \%. Атмосферное давление %v мм ртутного столба.\n
			🌬Ветер %v. Скорость ветра %v метров в секунду.\n
			`, r.Name, r.Weather.Description, r.Main.Temp, r.Main.FeelsLike, r.Main.Humidity, r.Main.Pressure, getWindDirection(r.Wind.Deg), r.Wind.Speed)
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		_, err := bot.Send(msg)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		log.Printf("%+v\n", update)
	}
}

func makeRequest(cfg *config.App) (*WeatherAPIResponse, error) {
	urlPath := cfg.WeatherAPI.URL + "?" + "id=" + cfg.WeatherAPI.CityID + "&appid=" + cfg.WeatherAPI.Token + "&lang=ru&units=metric"
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request status: %s",
			http.StatusText(resp.StatusCode))
	}

	var data WeatherAPIResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, nil
	}

	return &data, nil
}

// WeatherAPIResponse ...
type WeatherAPIResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func getWindDirection(deg int) string {

	switch {
	case 11 < deg && deg <= 33:
		return "северо северо восточный"
	case 33 < deg && deg <= 56:
		return "северо восточный"
	case 56 < deg && deg <= 76:
		return "восточно северо восточный"
	case 76 < deg && deg <= 101:
		return "Восточный"
	case 101 < deg && deg <= 123:
		return "восточно юго восточный"
	case 123 < deg && deg <= 146:
		return "юго восточный"
	case 146 < deg && deg <= 168:
		return "юго юго восточный"
	case 168 < deg && deg <= 191:
		return "южный"
	case 191 < deg && deg <= 213:
		return "юго юго западный"
	case 213 < deg && deg <= 236:
		return "югозападный"
	case 236 < deg && deg <= 258:
		return "западно юго западный"
	case 258 < deg && deg <= 281:
		return "западный"
	case 281 < deg && deg <= 303:
		return "западно северо западный"
	case 303 < deg && deg <= 326:
		return "северо западный"
	case 326 < deg && deg <= 348:
		return "северо сверо западный"
	case 348 < deg || deg <= 11:
		return "северный"
	}

	return ""
}
