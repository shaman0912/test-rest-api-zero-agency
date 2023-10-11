package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB
var log = logrus.New()

func Init() {
	//get params from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	//for connect Postgres
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	//connection pool
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.WithError(err).Fatal("Error when opening a database connect")
	}

	//check connect witn DB
	if err = db.Ping(); err != nil {
		log.WithError(err).Fatal("Error when check connect whit DB")

	}
	// Выполнение SQL-запроса для создания таблицы
	filename := "сreateTable.sql" // Замените на имя вашего файла
	if err := ExecuteSQLFromFile(db, filename); err != nil {
		log.Fatal(err)
	}
}

// provides access to the database connection
func NewDB() *sql.DB {
	return db
}

func ExecuteSQLFromFile(db *sql.DB, filename string) error {
	sqlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	sqlQuery := string(sqlFile)

	_, err = db.Exec(sqlQuery)
	return err
}
