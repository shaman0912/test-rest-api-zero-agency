package handlers

import "github.com/gofiber/fiber/v2"

// Ручка для получения списка новостей
func ListNews(c *fiber.Ctx) error {
	// Здесь вы можете запросить список новостей из базы данных
	// Пример: var news []News
	// err := db.Model(&news).Select()
	// Не забудьте обработать ошибки
	return c.JSON(fiber.Map{"news": nil})
}
