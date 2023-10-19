package internal

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/shaman0912/test-rest-api-zero-agency/internal/handlers"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
)

type NewsHandlers struct {
	db     *reform.DB
	logger *logrus.Logger
}

func (listHandler *NewsHandlers) listHandler(c *fiber.Ctx) error {
	return handlers.ListNews(c, listHandler.db, listHandler.logger)
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome api")
}

func SetupRoutes(app *fiber.App, db *reform.DB, logger *logrus.Logger) {
	// Создание группы маршрутов для API
	api := app.Group("/")

	listNewHandler := &NewsHandlers{db, logger}

	// Регистрация маршрутов
	api.Get("/wel", welcome)
	api.Post("/edit/:Id", handlers.EditNews)
	api.Get("/list", listNewHandler.listHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // По умолчанию используем порт 3000
	}

	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(app.Listen(":" + port))
}
