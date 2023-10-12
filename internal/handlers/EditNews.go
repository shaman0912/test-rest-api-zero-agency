package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Ручка для изменения новости
func EditNews(c *fiber.Ctx) error {
	// Здесь вы можете получить параметры из URL и тела запроса
	id := c.Params("Id")
	// Выполнить изменение новости в базе данных
	// Не забудьте обработать ошибки
	//_, err := db.Model((*News)(nil)).Where("id = ?", id).Update(news)
	fmt.Println(id)
	return c.JSON(fiber.Map{"message": "News updated successfully"})
}
