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

			reply = r.Name
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
	urlPath := cfg.WeatherAPI.URL + "?" + "id=" + cfg.WeatherAPI.CityID + "&appid=" + cfg.WeatherAPI.Token + "&lang=ru"
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
	Weather []struct {
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
