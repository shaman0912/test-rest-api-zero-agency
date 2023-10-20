package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Ручка для изменения новости
func EditNews(c *fiber.Ctx) error {
	id := c.Params("Id")

	fmt.Println(id)
	return c.JSON(fiber.Map{"message": "News updated successfully"})
}
