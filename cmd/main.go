package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/shaman0912/test-rest-api-zero-agency/internal"
	"github.com/shaman0912/test-rest-api-zero-agency/internal/database"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	db, err := database.ConnectToDB()
	if err != nil {
		logger.WithError(err).Error("Ошибка при выполнении connect  к БД")
		os.Exit(1)
	}

	err = database.CreateTables(db, "createTabless.sql")
	if err != nil {
		logger.WithError(err).Error("Ошибка при выполнении SQL-запросов из файла")
		os.Exit(1)
	}
	defer db.Close()

	dialect := postgresql.Dialect
	reformDB := reform.NewDB(db, dialect, nil)

	app := fiber.New()
	internal.SetupRoutes(app, reformDB, logger)

}
