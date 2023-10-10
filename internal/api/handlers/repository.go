package handlers

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
)

// type Repository struct {
// 	db  *reform.DB
// 	log *logrus.Entry
// }

func NewRepository(db *reform.DB, log *logrus.Entry) *Repository {
	return &Repository{db, log}
}

func (r *Repository) EditNewsByID(id int64, title, content string, categories []int64) error {
	// Ваш код для изменения новости в базе данных
	// Используйте вашу библиотеку базы данных (например, reform) для выполнения запроса
	// Например, если вы используете reform, это может выглядеть следующим образом:

	// news := &News{
	//     ID:         id,
	//     Title:      title,
	//     Content:    content,
	//     Categories: categories,
	// }
	// _, err := r.db.Update(news)

	// Верните ошибку, если что-то пошло не так
	// return err

	return nil
}

func (r *Repository) GetNewsList() ([]*News, error) {
	// Ваш код для получения списка новостей из базы данных
	return nil, nil
}

func (r *Repository) HandleEditNews(requestData struct {
	ID         int64
	Title      string
	Content    string
	Categories []int64
}) error {
	// Ваш код для изменения новости в базе данных на основе requestData

	return nil
}
