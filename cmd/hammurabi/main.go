package main

import (
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

	_, err = bot.SetWebhook(tgbotapi.NewWebhook("https://hammurabi-bot-game.herokuapp.com/" + cfg.Bot.Token))
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
		log.Printf("%+v\n", update)
	}
}
