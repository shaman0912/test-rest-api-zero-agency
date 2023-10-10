package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var db *sql.DB
var log = logrus.New()

func init() {
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
}

// provides access to the database connection
func NewDB() *sql.DB {
	return db
}
