package config

import "os"

// App struct for store application configs
type App struct {
	Bot BotConfig
}

// BotConfig store telegram bot configs
type BotConfig struct {
	Token string
}

// New returns a new App struct
func New() *App {
	return &App{
		Bot: BotConfig{
			Token: getEnv("TOKEN", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
