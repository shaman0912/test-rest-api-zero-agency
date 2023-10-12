package internal

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/shaman0912/test-rest-api-zero-agency/internal/handlers"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome api")
}

func SetupRoutes(app *fiber.App) {
	// Создание группы маршрутов для API
	api := app.Group("/")

	// Регистрация маршрутов
	api.Get("/wel", welcome)
	api.Post("/edit/:Id", handlers.EditNews)
	api.Get("/list", handlers.ListNews)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // По умолчанию используем порт 3000
	}

	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(app.Listen(":" + port))
}
