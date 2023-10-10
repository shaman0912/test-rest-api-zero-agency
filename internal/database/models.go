package database

type News struct {
	ID         int64   `json:"Id"`
	Title      string  `json:"Title"`
	Content    string  `json:"Content"`
	Categories []int64 `json:"Categories"`
}

type NewsCategory struct {
	NewsID     int64 `reform:"NewsId"`
	CategoryID int64 `reform:"CategoryId"`
}
