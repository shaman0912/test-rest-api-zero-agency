package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
)

var NewsModels reform.Table // Это ваша таблица

// Ручка для получения списка новостей
func ListNews(c *fiber.Ctx, db *reform.DB, logger *logrus.Logger) error {

	// Find records by IDs.
	// Find records by IDs.
	newsData, err := db.FindAllFrom(NewsModels, "id")
	if err != nil {
		logger.Error(err)
		return err
	}
	fmt.Println(newsData)
	// for _, p := range newsData {
	// 	fmt.Println(p)
	// }
	// // Преобразуйте новости в нужный формат для ответа
	//var response []fiber.Map
	// for _, n := range newsData {
	// 	response = append(response, fiber.Map{
	// 		"Id":      n.ID,
	// 		"Title":   n.Title,
	// 		"Content": n.Content,
	// 	})
	// }

	// return c.JSON(fiber.Map{
	// 	"Success": true,
	// 	"News":    response,
	// })
	return c.SendString("Welcome api")
}
