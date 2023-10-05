package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBEngine string

	APP_PORT      string
	APP_MODE      string
	GIN_MODE      string
	CLIENT_ORIGIN string

	TOKEN_SECRET  string
	TOKEN_MAX_AGE int

	TelebotToken string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	maxAge, _ := strconv.Atoi(os.Getenv("TOKEN_MAX_AGE"))
	cfg := &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		DBUser:        os.Getenv("DB_USERNAME"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBEngine:        os.Getenv("DB_ENGINE"),
		APP_PORT:      os.Getenv("PORT"),
		APP_MODE:      os.Getenv("APP_MODE"),
		CLIENT_ORIGIN: os.Getenv("CLIENT_ORIGIN"),
		TOKEN_SECRET:  os.Getenv("TOKEN_SECRET"),
		TOKEN_MAX_AGE: maxAge,
		TelebotToken:  os.Getenv("TELEBOT_TOKEN"),
	}

	return cfg, nil
}
