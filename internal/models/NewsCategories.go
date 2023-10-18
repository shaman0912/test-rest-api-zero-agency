package models

type NewsCategories struct {
	NewsID     int64 `reform:"news_id"`
	Categories int64 `reform:"category_id"`
}
