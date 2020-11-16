package main

import (
	"hammurabi-game/config"
	"hammurabi-game/internal/app"
	"log"
	"net/http"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
)

// ChatID ...
const ChatID = 191155356

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

	go func() {
		for {
			<-time.After(5 * time.Minute)
			go sendWeather(bot, cfg)
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
			reply = app.WeatherResponse(cfg)
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

func sendWeather(bot tgbotapi.Bot, cfg *config.App) {
	reply := app.WeatherResponse(cfg)
	msg := tgbotapi.NewMessage(ChatID, reply)
	// отправляем
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
}
