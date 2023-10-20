package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func ConnectToDB() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)
	//fmt.Println("THIS is connStr", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Устанавливаем максимальное количество соединений в пуле
	db.SetMaxOpenConns(10)

	// Устанавливаем максимальное время жизни соединения в пуле
	db.SetConnMaxLifetime(time.Minute * 5)

	// Устанавливаем максимальное время простоя соединения в пуле
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Успешно подключено к базе данных PostgreSQL!")

	return db, nil

}
