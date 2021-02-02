package models

type Book struct {
	ID     int    `json:"_id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
	Link   string `json:"link" db:"link"`
	Total  int    `json:"total" db:"total"`
	Unique int    `json:"unique" db:"uniq"`
	Kanji  string `json:"kanji" db:"kanji"`
}

type BookSub struct {
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
	Link   string `json:"link" db:"link"`
	Total  int    `json:"total" db:"total"`
	Unique int    `json:"unique" db:"uniq"`
}
