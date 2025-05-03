package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Token string
}

func Load() *Config {
	// Загружаем .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Token: os.Getenv("BOT_TOKEN"), // Читаем переменную
	}
}
