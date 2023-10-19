package handlers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shaman0912/test-rest-api-zero-agency/internal/models"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
)

func ListNews(c *fiber.Ctx, db *reform.DB, logger *logrus.Logger) error {
	// экземпляр таблицы News
	newsTable := models.NewsTable

	// Find records by IDs.
	newsData, err := db.FindAllFrom(newsTable, "Id", 1, 2)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range newsData {
		fmt.Println(p)
	}
	return nil
}
