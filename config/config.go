package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBName     string
	Port       string
}

func NewConfig() (*Config, error) {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Инициализация полей конфигурации из переменных окружения
	config := &Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		Port:       os.Getenv("PORT"),
	}

	return config, nil
}
