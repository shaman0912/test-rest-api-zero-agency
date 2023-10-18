package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DBConfig содержит конфигурацию базы данных.
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Sslmode  string
}

// AppConfig содержит всю конфигурацию приложения, включая DBConfig.
type AppConfig struct {
	DBConfig DBConfig
}

// LoadConfig загружает конфигурацию из файла .env и возвращает объект AppConfig.
func LoadConfig() *AppConfig {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	return &AppConfig{
		DBConfig: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASS"),
			DBName:   os.Getenv("DB_NAME"),
			Sslmode:  os.Getenv("DB_SSLMODE"),
		},
	}
}
