package models

//go:generate reform

// reform:NewsCategories
type NewsCategories struct {
	NewsID     int64 `reform:"News_id"`
	Categories int64 `reform:"Categories_id"`
}
