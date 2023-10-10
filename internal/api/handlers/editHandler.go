package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
)

type News struct {
	ID      int    `reform:"id"`
	Title   string `reform:"title"`
	Content string `reform:"content"`
}
type Repository struct {
	db  *reform.DB
	log *logrus.Entry
}

// HasPK определяет, есть ли у записи первичный ключ
func (b *News) HasPK() bool {
	return b.ID != 0
}

func (r *Repository) UpdateNews(context *fiber.Ctx) error {
	// Получаем параметр ID из URL
	id := context.Params("id")

	// Получаем новые данные новости из тела запроса
	var updatedNews News
	if err := context.BodyParser(&updatedNews); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	// Ищем запись по ID
	var existingNews News
	err := r.db.FindByPrimaryKeyTo(&existingNews, id)
	if err != nil {
		// Запись с указанным ID не найдена
		context.Status(http.StatusNotFound).JSON(
			&fiber.Map{"message": "news not found"})
		return err
	}

	// Обновляем данные новости
	existingNews.Title = updatedNews.Title
	existingNews.Content = updatedNews.Content

	// Сохраняем обновленную запись в базе данных
	err = r.db.Save(&existingNews)
	if err != nil {
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
