package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	api "github.com/shaman0912/test-rest-api-zero-agency/internal"
	"github.com/shaman0912/test-rest-api-zero-agency/internal/database"
	"github.com/sirupsen/logrus"
)

func main() {
	// Создание экземпляра Fiber
	app := fiber.New()

	// Инициализация соединения с базой данных с использованием connection pool
	db := database.NewDB()
	defer db.Close()

	// Настройка логирования с Logrus для вывода в терминал
	logrus.SetOutput(os.Stdout)

	// Передача экземпляра Fiber, базы данных и логгера в настройку маршрутов
	api.SetupRoutes(app)
}
