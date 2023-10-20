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
	// newsTable := models.NewsTable

	// newsCategories := models.NewsCategoriesView

	// fmt.Println("newsTable: ", newsTable)
	// fmt.Println("newscATEGORIES: ", newsCategories)

	// Find records by IDs.
	newsData, err := db.FindAllFrom(models.NewsTable, "Id", 1, 2)
	if err != nil {
		log.Fatal(err)
	}
	for _, p := range newsData {
		fmt.Println(p)
	}
	return nil

}
