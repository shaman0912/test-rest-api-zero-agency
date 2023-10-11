package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/shaman0912/test-rest-api-zero-agency/internal/database"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
)

type Repository struct {
	db  *reform.DB
	log *logrus.Entry
}

func (r *Repository) EditNews(context *fiber.Ctx) error {
	// Получаем параметр ID из URL
	newsID := context.Params("id")

	// Получаем новые данные новости из тела запроса
	var updatedNews database.News

	if err := context.BodyParser(&updatedNews); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// Создайте объект reform.DB с использованием r.db.
	reformDBInstance := reform.NewDB(r.db, reform.Dialect, reform.Logger)

	// Используйте метод FindByPrimaryKeyTo для поиска новости по ID
	var idTemp database.News

	err := reformDBInstance.FindByPrimaryKeyTo(&idTemp, newsID)
	if err != nil {
		if err == reform.ErrNoRows {
			context.Status(http.StatusInternalServerError).JSON(
				&fiber.Map{"message": "ID not found in DB"})
			return err
			// Запись не найдена
		} else {
			// Произошла другая ошибка при выполнении запроса
			context.Status(http.StatusInternalServerError).JSON(
				&fiber.Map{"message": "Internal Server Error"})
			return err
		}
	}

	// Обновляем данные новости
	idTemp.Title = updatedNews.Title
	idTemp.Content = updatedNews.Content

	// Сохраняем обновленную запись в базе данных
	if err := reformDBInstance.Save(&idTemp); err != nil {
		context.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{"message": "could not update news"})
		return err
	}

	// Возвращаем ответ с кодом статуса 200 OK
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "news has been updated",
	})
	return nil
}
