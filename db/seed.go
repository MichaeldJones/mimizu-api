package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	m "../models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Temp struct {
	ID     int    `json:"id" `
	Title  string `json:"title"`
	Author string `json:"author"`
	Link   string `json:"link"`
	Total  int    `json:"total"`
	Unique int    `json:"unique"`
	Kanji  []rune `json:"kanji"`
}

func main() {

	var temps []Temp

	connStr := "user= dbname= password= host=localhost sslmode=disable"

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	query := `insert into books(id, title, author, link, total, uniq, kanji)
						 values(:id, :title, :author, :link, :total, :uniq, :kanji)`

	file, err := ioutil.ReadFile("aozora.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(file, &temps)

	for _, i := range temps {

		var tmp m.Book

		tmp.ID = i.ID
		tmp.Title = i.Title
		tmp.Author = i.Author
		tmp.Link = i.Link
		tmp.Total = i.Total
		tmp.Unique = i.Unique
		tmp.Kanji = string(i.Kanji)

		_, err = db.NamedExec(query, tmp)
		if err != nil {
			fmt.Println(err)
		}

	}

}
