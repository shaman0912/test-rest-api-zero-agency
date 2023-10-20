package database

import (
	"database/sql"
	"os"

	"github.com/sirupsen/logrus"
)

func CreateTables(db *sql.DB, filePath string) error {
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(contents))
	if err != nil {
		logrus.WithError(err).Error("Ошибка при выполнении SQL-запросов из файла")
		return err
	}

	logrus.Info("TABLES: News and NewsCategories  successfully created")

	return nil
}
