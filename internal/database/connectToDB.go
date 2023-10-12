package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Функция для установления соединения с базой данных
func ConnectToDB() (*sql.DB, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	// Получите параметры подключения из переменных окружения
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Остальной код остается без изменений
	// ...// Формируйте строку подключения к PostgreSQL
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	// Открываем соединение с PostgreSQL
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Устанавливаем максимальное количество соединений в пуле
	db.SetMaxOpenConns(10)

	// Устанавливаем максимальное время жизни соединения в пуле
	db.SetConnMaxLifetime(time.Minute * 5)

	// Устанавливаем максимальное время простоя соединения в пуле
	db.SetMaxIdleConns(5)

	// Проверка соединения
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Успешно подключено к базе данных PostgreSQL!")

	return db, nil

}
