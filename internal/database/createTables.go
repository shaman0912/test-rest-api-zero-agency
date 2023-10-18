package database

import (
	"database/sql"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func CreateTables(db *sql.DB, filePath string) error {
	// Чтение содержимого SQL-файла
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Выполнение SQL-запросов
	_, err = db.Exec(string(contents))
	if err != nil {
		logrus.WithError(err).Error("Ошибка при выполнении SQL-запросов из файла")
		return err
	}

	logrus.Info("TABLES: News and NewsCategories  successfully created")

	return nil
}
