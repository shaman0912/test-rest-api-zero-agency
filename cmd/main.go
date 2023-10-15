package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/shaman0912/test-rest-api-zero-agency/internal"
	"github.com/shaman0912/test-rest-api-zero-agency/internal/database"

	"github.com/sirupsen/logrus"
)

func main() {

	// Инициализация логгера
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Подключение к базе данных Postgres с использованием connection pool
	db, err := database.ConnectToDB()
	if err != nil {
		logger.WithError(err).Error("Ошибка при выполнении connect  к БД")
		os.Exit(1)
	}

	// создаем таблицы в БД
	err = database.CreateTables(db, "createTabless.sql")
	if err != nil {
		logger.WithError(err).Error("Ошибка при выполнении SQL-запросов из файла")
		os.Exit(1)
	}
	defer db.Close()
	app := fiber.New()

	internal.SetupRoutes(app)

}
